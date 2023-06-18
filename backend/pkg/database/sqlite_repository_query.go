package database

import (
	"context"
	"fmt"
	"github.com/fastenhealth/fastenhealth-onprem/backend/pkg/models"
	databaseModel "github.com/fastenhealth/fastenhealth-onprem/backend/pkg/models/database"
	"golang.org/x/exp/maps"
	"sort"
	"strconv"
	"strings"
	"time"
)

type SearchParameterType string

const (
	SearchParameterTypeNumber    SearchParameterType = "number"
	SearchParameterTypeDate      SearchParameterType = "date"
	SearchParameterTypeString    SearchParameterType = "string"
	SearchParameterTypeToken     SearchParameterType = "token"
	SearchParameterTypeReference SearchParameterType = "reference"
	SearchParameterTypeUri       SearchParameterType = "uri"
	SearchParameterTypeComposite SearchParameterType = "composite"
	SearchParameterTypeQuantity  SearchParameterType = "quantity"
	SearchParameterTypeSpecial   SearchParameterType = "special"
)

func (sr *SqliteRepository) QueryResources(ctx context.Context, query models.QueryResource) (interface{}, error) {
	//todo, until we actually parse the select statement, we will just return all resources based on "from"

	//find the associated Gorm Model for this query
	queryModel, err := databaseModel.NewFhirResourceModelByType(query.From)
	if err != nil {
		return nil, err
	}

	whereClauses := []string{}
	namedParameters := map[string]interface{}{}

	//find the FHIR search types associated with each where clause. Any unknown parameters will be ignored.
	searchCodeToTypeLookup := queryModel.GetSearchParameters()
	for searchParamCodeWithModifier, searchParamCodeValue := range query.Where {
		searchParameter, err := ProcessSearchParameter(searchParamCodeWithModifier, searchCodeToTypeLookup)
		if err != nil {
			return nil, err
		}
		searchParameterValue, err := ProcessSearchParameterValue(searchParameter, searchParamCodeValue)
		if err != nil {
			return nil, err
		}

		whereClause, clauseNamedParameters, err := SearchCodeToWhereClause(searchParameter, searchParameterValue)
		if err != nil {
			return nil, err
		}
		//add generated where clause to the list, and add the named parameters to the map of existing named parameters
		whereClauses = append(whereClauses, whereClause)
		maps.Copy(namedParameters, clauseNamedParameters)
	}

	//for safety, we will always add/override the current user_id to the where clause. This is to ensure that the user doesnt attempt to override this value in their own where clause
	currentUser, currentUserErr := sr.GetCurrentUser(ctx)
	if currentUserErr != nil {
		return nil, currentUserErr
	}
	namedParameters["user_id"] = currentUser.ID.String()
	whereClauses = append(whereClauses, "user_id = @user_id")

	results := []map[string]interface{}{}
	clientResp := sr.GormClient.Where(strings.Join(whereClauses, " AND "), namedParameters).Model(queryModel).Find(&results)

	return results, clientResp.Error
}

/// INTERNAL functionality. These functions are exported for testing, but are not available in the Interface

type SearchParameter struct {
	Name     string
	Type     SearchParameterType
	Modifier string
}

type SearchParameterValue struct {
	Prefix          string
	Value           interface{}
	SecondaryValues map[string]interface{}
}

func ProcessSearchParameter(searchCodeWithModifier string, searchParamTypeLookup map[string]string) (SearchParameter, error) {
	searchParameter := SearchParameter{}

	//determine the searchCode searchCodeModifier
	//TODO: this is only applicable to string, token, reference and uri type (however unknown names & modifiers are ignored)
	if searchCodeParts := strings.SplitN(searchCodeWithModifier, ":", 2); len(searchCodeParts) == 2 {
		searchParameter.Name = searchCodeParts[0]
		searchParameter.Modifier = searchCodeParts[1]
	} else {
		searchParameter.Name = searchCodeParts[0]
		searchParameter.Modifier = ""
	}

	//next, determine the searchCodeType for this Resource (or throw an error if it is unknown)
	searchParamTypeStr, searchParamTypeOk := searchParamTypeLookup[searchParameter.Name]
	if !searchParamTypeOk {
		return searchParameter, fmt.Errorf("unknown search parameter: %s", searchParameter.Name)
	} else {
		searchParameter.Type = SearchParameterType(searchParamTypeStr)
	}
	return searchParameter, nil
}

func ProcessSearchParameterValue(searchParameter SearchParameter, searchValueWithPrefix string) (SearchParameterValue, error) {
	searchParameterValue := SearchParameterValue{
		SecondaryValues: map[string]interface{}{},
		Value:           searchValueWithPrefix,
	}
	if (searchParameter.Type == SearchParameterTypeString || searchParameter.Type == SearchParameterTypeUri) && len(searchParameterValue.Value.(string)) == 0 {
		return searchParameterValue, fmt.Errorf("invalid search parameter value: (%s=%s)", searchParameter.Name, searchParameterValue.Value)
	}

	//certain types (like number,date and quanitty have a prefix that needs to be parsed)
	if searchParameter.Type == SearchParameterTypeNumber || searchParameter.Type == SearchParameterTypeDate || searchParameter.Type == SearchParameterTypeQuantity {
		//loop though all known/allowed prefixes, and determine if the searchValueWithPrefix starts with one of them
		allowedPrefixes := []string{"eq", "ne", "gt", "lt", "ge", "le", "sa", "eb", "ap"}
		for _, allowedPrefix := range allowedPrefixes {
			if strings.HasPrefix(searchValueWithPrefix, allowedPrefix) {
				searchParameterValue.Prefix = allowedPrefix
				searchParameterValue.Value = strings.TrimPrefix(searchValueWithPrefix, allowedPrefix)
				break
			}
		}
	}

	//certain Types (like token, quantity, reference) have secondary query values that need to be parsed (delimited by "|") value
	if searchParameter.Type == SearchParameterTypeQuantity {
		if searchParameterValueParts := strings.SplitN(searchParameterValue.Value.(string), "|", 3); len(searchParameterValueParts) == 1 {
			searchParameterValue.Value = searchParameterValueParts[0]
		} else if len(searchParameterValueParts) == 2 {
			searchParameterValue.Value = searchParameterValueParts[0]
			if len(searchParameterValueParts[1]) > 0 {
				searchParameterValue.SecondaryValues[searchParameter.Name+"System"] = searchParameterValueParts[1]
			}
		} else if len(searchParameterValueParts) == 3 {
			searchParameterValue.Value = searchParameterValueParts[0]
			if len(searchParameterValueParts[1]) > 0 {
				searchParameterValue.SecondaryValues[searchParameter.Name+"System"] = searchParameterValueParts[1]
			}
			if len(searchParameterValueParts[2]) > 0 {
				searchParameterValue.SecondaryValues[searchParameter.Name+"Code"] = searchParameterValueParts[2]
			}
		}
	} else if searchParameter.Type == SearchParameterTypeToken {
		if searchParameterValueParts := strings.SplitN(searchParameterValue.Value.(string), "|", 2); len(searchParameterValueParts) == 1 {
			searchParameterValue.Value = searchParameterValueParts[0] //this is a code
			if len(searchParameterValue.Value.(string)) == 0 {
				return searchParameterValue, fmt.Errorf("invalid search parameter value: (%s=%s)", searchParameter.Name, searchParameterValue.Value)
			}
		} else if len(searchParameterValueParts) == 2 {
			//if theres 2 parts, first is always system, second is always the code. Either one may be emty. If both are emty this is invalid.
			if len(searchParameterValueParts[0]) > 0 {
				searchParameterValue.SecondaryValues[searchParameter.Name+"System"] = searchParameterValueParts[0]
			}
			if len(searchParameterValueParts[1]) > 0 {
				searchParameterValue.Value = searchParameterValueParts[1]
			}
			if len(searchParameterValueParts[0]) == 0 && len(searchParameterValueParts[1]) == 0 {
				return searchParameterValue, fmt.Errorf("invalid search parameter value: (%s=%s)", searchParameter.Name, searchParameterValue.Value)
			}
		}
	} else if searchParameter.Type == SearchParameterTypeReference {
		//todo
		return searchParameterValue, fmt.Errorf("search parameter type not yet implemented: %s", searchParameter.Type)
	}

	//certain types (Quantity and Number) need to be converted to Float64
	if searchParameter.Type == SearchParameterTypeQuantity || searchParameter.Type == SearchParameterTypeNumber {
		if conv, err := strconv.ParseFloat(searchParameterValue.Value.(string), 64); err == nil {
			searchParameterValue.Value = conv
		} else {
			return searchParameterValue, fmt.Errorf("invalid search parameter value (NaN): (%s=%s)", searchParameter.Name, searchParameterValue.Value)
		}
	} else if searchParameter.Type == SearchParameterTypeDate {
		//other types (like date) need to be converted to a time.Time
		if conv, err := time.Parse(time.RFC3339, searchParameterValue.Value.(string)); err == nil {
			searchParameterValue.Value = conv
		} else {
			// fallback to parsing just a date (without time)
			if conv, err := time.Parse("2006-01-02", searchParameterValue.Value.(string)); err == nil {
				searchParameterValue.Value = conv
			} else {
				return searchParameterValue, fmt.Errorf("invalid search parameter value (invalid date): (%s=%s)", searchParameter.Name, searchParameterValue.Value)
			}
		}
	}
	return searchParameterValue, nil
}

func SearchCodeToWhereClause(searchParam SearchParameter, searchParamValue SearchParameterValue) (string, map[string]interface{}, error) {

	//add named parameters to the lookup map. Basically, this is a map of all the named parameters that will be used in the where clause we're generating
	searchClauseNamedParams := map[string]interface{}{
		searchParam.Name: searchParamValue.Value,
	}
	for k, v := range searchParamValue.SecondaryValues {
		searchClauseNamedParams[k] = v
	}

	//parse the searchCode and searchCodeValue to determine the correct where clause
	switch searchParam.Type {
	case SearchParameterTypeNumber, SearchParameterTypeDate:

		if searchParamValue.Prefix == "" || searchParamValue.Prefix == "eq" {
			return fmt.Sprintf("%s = @%s", searchParam.Name, searchParam.Name), searchClauseNamedParams, nil
		} else if searchParamValue.Prefix == "lt" || searchParamValue.Prefix == "eb" {
			return fmt.Sprintf("%s < @%s", searchParam.Name, searchParam.Name), searchClauseNamedParams, nil
		} else if searchParamValue.Prefix == "le" {
			return fmt.Sprintf("%s <= @%s", searchParam.Name, searchParam.Name), searchClauseNamedParams, nil
		} else if searchParamValue.Prefix == "gt" || searchParamValue.Prefix == "sa" {
			return fmt.Sprintf("%s > @%s", searchParam.Name, searchParam.Name), searchClauseNamedParams, nil
		} else if searchParamValue.Prefix == "ge" {
			return fmt.Sprintf("%s >= @%s", searchParam.Name, searchParam.Name), searchClauseNamedParams, nil
		} else if searchParamValue.Prefix == "ne" {
			return fmt.Sprintf("%s <> @%s", searchParam.Name, searchParam.Name), searchClauseNamedParams, nil
		} else if searchParam.Modifier == "ap" {
			return "", nil, fmt.Errorf("search modifier 'ap' not supported for search parameter type %s (%s=%s)", searchParam.Type, searchParam.Name, searchParamValue.Value)
		}
	case SearchParameterTypeString:
		if searchParam.Modifier == "" {
			searchClauseNamedParams[searchParam.Name] = searchParamValue.Value.(string) + "%" // "eve" matches "Eve" and "Evelyn"
			return fmt.Sprintf("%s LIKE @%s", searchParam.Name, searchParam.Name), searchClauseNamedParams, nil
		} else if searchParam.Modifier == "exact" {
			// "eve" matches "eve" (not "Eve" or "EVE")
			return fmt.Sprintf("%s = @%s", searchParam.Name, searchParam.Name), searchClauseNamedParams, nil
		} else if searchParam.Modifier == "contains" {
			searchClauseNamedParams[searchParam.Name] = "%" + searchParamValue.Value.(string) + "%" // "eve" matches "Eve", "Evelyn" and "Severine"
			return fmt.Sprintf("%s LIKE @%s", searchParam.Name, searchParam.Name), searchClauseNamedParams, nil
		}
	case SearchParameterTypeUri:
		if searchParam.Modifier == "" {
			return fmt.Sprintf("%s = @%s", searchParam.Name, searchParam.Name), searchClauseNamedParams, nil
		} else if searchParam.Modifier == "below" {
			searchClauseNamedParams[searchParam.Name] = searchParamValue.Value.(string) + "%" // column starts with "http://example.com"
			return fmt.Sprintf("%s LIKE @%s", searchParam.Name, searchParam.Name), searchClauseNamedParams, nil
		} else if searchParam.Modifier == "above" {
			return "", nil, fmt.Errorf("search modifier 'above' not supported for search parameter type %s (%s=%s)", searchParam.Type, searchParam.Name, searchParamValue.Value)
		}
	case SearchParameterTypeQuantity:

		//setup the clause
		var clause string
		if searchParamValue.Prefix == "" || searchParamValue.Prefix == "eq" {
			clause = fmt.Sprintf("%s = @%s", searchParam.Name, searchParam.Name)
		} else if searchParamValue.Prefix == "lt" || searchParamValue.Prefix == "eb" {
			clause = fmt.Sprintf("%s < @%s", searchParam.Name, searchParam.Name)
		} else if searchParamValue.Prefix == "le" {
			clause = fmt.Sprintf("%s <= @%s", searchParam.Name, searchParam.Name)
		} else if searchParamValue.Prefix == "gt" || searchParamValue.Prefix == "sa" {
			clause = fmt.Sprintf("%s > @%s", searchParam.Name, searchParam.Name)
		} else if searchParamValue.Prefix == "ge" {
			clause = fmt.Sprintf("%s >= @%s", searchParam.Name, searchParam.Name)
		} else if searchParamValue.Prefix == "ne" {
			clause = fmt.Sprintf("%s <> @%s", searchParam.Name, searchParam.Name)
		} else if searchParamValue.Prefix == "ap" {
			return "", nil, fmt.Errorf("search modifier 'ap' not supported for search parameter type %s (%s=%s)", searchParam.Type, searchParam.Name, searchParamValue.Value)
		}

		//append the code and/or system clauses (if required)
		//this looks like unnecessary code, however its required to ensure consistent tests
		keys := make([]string, 0, len(searchParamValue.SecondaryValues))
		for k := range searchParamValue.SecondaryValues {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			clause += fmt.Sprintf(` AND %s = @%s`, k, k)
		}

		return clause, searchClauseNamedParams, nil
	case SearchParameterTypeToken, SearchParameterTypeReference:
		return "", nil, fmt.Errorf("search parameter type %s not supported", searchParam.Type)
	}
	return "", searchClauseNamedParams, nil
}
