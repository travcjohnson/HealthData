// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-onprem/blob/main/backend/pkg/models/database/generate.go
// PLEASE DO NOT EDIT BY HAND

package database

import (
	"encoding/json"
	"fmt"
	goja "github.com/dop251/goja"
	models "github.com/fastenhealth/fasten-onprem/backend/pkg/models"
	datatypes "gorm.io/datatypes"
	"time"
)

type FhirProvenance struct {
	models.ResourceBase
	// Who participated
	// https://hl7.org/fhir/r4/search.html#reference
	Agent datatypes.JSON `gorm:"column:agent;type:text;serializer:json" json:"agent,omitempty"`
	// What the agents role was
	// https://hl7.org/fhir/r4/search.html#token
	AgentRole datatypes.JSON `gorm:"column:agentRole;type:text;serializer:json" json:"agentRole,omitempty"`
	// How the agent participated
	// https://hl7.org/fhir/r4/search.html#token
	AgentType datatypes.JSON `gorm:"column:agentType;type:text;serializer:json" json:"agentType,omitempty"`
	// Identity of entity
	// https://hl7.org/fhir/r4/search.html#reference
	Entity datatypes.JSON `gorm:"column:entity;type:text;serializer:json" json:"entity,omitempty"`
	// Language of the resource content
	// https://hl7.org/fhir/r4/search.html#token
	Language datatypes.JSON `gorm:"column:language;type:text;serializer:json" json:"language,omitempty"`
	// Where the activity occurred, if relevant
	// https://hl7.org/fhir/r4/search.html#reference
	Location datatypes.JSON `gorm:"column:location;type:text;serializer:json" json:"location,omitempty"`
	// When the resource version last changed
	// https://hl7.org/fhir/r4/search.html#date
	MetaLastUpdated *time.Time `gorm:"column:metaLastUpdated;type:datetime" json:"metaLastUpdated,omitempty"`
	// Profiles this resource claims to conform to
	// https://hl7.org/fhir/r4/search.html#reference
	MetaProfile datatypes.JSON `gorm:"column:metaProfile;type:text;serializer:json" json:"metaProfile,omitempty"`
	// Tags applied to this resource
	// https://hl7.org/fhir/r4/search.html#token
	MetaTag datatypes.JSON `gorm:"column:metaTag;type:text;serializer:json" json:"metaTag,omitempty"`
	// Tags applied to this resource
	// This is a primitive string literal (`keyword` type). It is not a recognized SearchParameter type from https://hl7.org/fhir/r4/search.html, it's Fasten Health-specific
	MetaVersionId string `gorm:"column:metaVersionId;type:text" json:"metaVersionId,omitempty"`
	// When the activity was recorded / updated
	// https://hl7.org/fhir/r4/search.html#date
	Recorded *time.Time `gorm:"column:recorded;type:datetime" json:"recorded,omitempty"`
	// Indication of the reason the entity signed the object(s)
	// https://hl7.org/fhir/r4/search.html#token
	SignatureType datatypes.JSON `gorm:"column:signatureType;type:text;serializer:json" json:"signatureType,omitempty"`
	// Target Reference(s) (usually version specific)
	// https://hl7.org/fhir/r4/search.html#reference
	Target datatypes.JSON `gorm:"column:target;type:text;serializer:json" json:"target,omitempty"`
	// Text search against the narrative
	// This is a primitive string literal (`keyword` type). It is not a recognized SearchParameter type from https://hl7.org/fhir/r4/search.html, it's Fasten Health-specific
	Text string `gorm:"column:text;type:text" json:"text,omitempty"`
	// A resource type filter
	// https://hl7.org/fhir/r4/search.html#special
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
	// When the activity occurred
	// https://hl7.org/fhir/r4/search.html#date
	When *time.Time `gorm:"column:when;type:datetime" json:"when,omitempty"`
}

func (s *FhirProvenance) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"agent":                "reference",
		"agentRole":            "token",
		"agentType":            "token",
		"entity":               "reference",
		"id":                   "keyword",
		"language":             "token",
		"location":             "reference",
		"metaLastUpdated":      "date",
		"metaProfile":          "reference",
		"metaTag":              "token",
		"metaVersionId":        "keyword",
		"recorded":             "date",
		"signatureType":        "token",
		"sort_date":            "date",
		"source_id":            "keyword",
		"source_resource_id":   "keyword",
		"source_resource_type": "keyword",
		"source_uri":           "keyword",
		"target":               "reference",
		"text":                 "keyword",
		"type":                 "special",
		"when":                 "date",
	}
	return searchParameters
}
func (s *FhirProvenance) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
	s.ResourceRaw = datatypes.JSON(resourceRaw)
	// unmarshal the raw resource (bytes) into a map
	var resourceRawMap map[string]interface{}
	err := json.Unmarshal(resourceRaw, &resourceRawMap)
	if err != nil {
		return err
	}
	if len(fhirPathJs) == 0 {
		return fmt.Errorf("fhirPathJs script is empty")
	}
	vm := goja.New()
	// setup the global window object
	vm.Set("window", vm.NewObject())
	// set the global FHIR Resource object
	vm.Set("fhirResource", resourceRawMap)
	// compile the fhirpath library
	fhirPathJsProgram, err := goja.Compile("fhirpath.min.js", fhirPathJs, true)
	if err != nil {
		return err
	}
	// compile the searchParametersExtractor library
	searchParametersExtractorJsProgram, err := goja.Compile("searchParameterExtractor.js", searchParameterExtractorJs, true)
	if err != nil {
		return err
	}
	// add the fhirpath library in the goja vm
	_, err = vm.RunProgram(fhirPathJsProgram)
	if err != nil {
		return err
	}
	// add the searchParametersExtractor library in the goja vm
	_, err = vm.RunProgram(searchParametersExtractorJsProgram)
	if err != nil {
		return err
	}
	// execute the fhirpath expression for each search parameter
	// extracting Agent
	agentResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Provenance.agent.who')")
	if err == nil && agentResult.String() != "undefined" {
		s.Agent = []byte(agentResult.String())
	}
	// extracting AgentRole
	agentRoleResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Provenance.agent.role')")
	if err == nil && agentRoleResult.String() != "undefined" {
		s.AgentRole = []byte(agentRoleResult.String())
	}
	// extracting AgentType
	agentTypeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Provenance.agent.type')")
	if err == nil && agentTypeResult.String() != "undefined" {
		s.AgentType = []byte(agentTypeResult.String())
	}
	// extracting Entity
	entityResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Provenance.entity.what')")
	if err == nil && entityResult.String() != "undefined" {
		s.Entity = []byte(entityResult.String())
	}
	// extracting Language
	languageResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'language')")
	if err == nil && languageResult.String() != "undefined" {
		s.Language = []byte(languageResult.String())
	}
	// extracting Location
	locationResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Provenance.location')")
	if err == nil && locationResult.String() != "undefined" {
		s.Location = []byte(locationResult.String())
	}
	// extracting MetaLastUpdated
	metaLastUpdatedResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'meta.lastUpdated')")
	if err == nil && metaLastUpdatedResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, metaLastUpdatedResult.String())
		if err == nil {
			s.MetaLastUpdated = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", metaLastUpdatedResult.String())
			if err == nil {
				s.MetaLastUpdated = &d
			}
		}
	}
	// extracting MetaProfile
	metaProfileResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'meta.profile')")
	if err == nil && metaProfileResult.String() != "undefined" {
		s.MetaProfile = []byte(metaProfileResult.String())
	}
	// extracting MetaTag
	metaTagResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'meta.tag')")
	if err == nil && metaTagResult.String() != "undefined" {
		s.MetaTag = []byte(metaTagResult.String())
	}
	// extracting MetaVersionId
	metaVersionIdResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'meta.versionId')")
	if err == nil && metaVersionIdResult.String() != "undefined" {
		s.MetaVersionId = metaVersionIdResult.String()
	}
	// extracting Recorded
	recordedResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'Provenance.recorded')")
	if err == nil && recordedResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, recordedResult.String())
		if err == nil {
			s.Recorded = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", recordedResult.String())
			if err == nil {
				s.Recorded = &d
			}
		}
	}
	// extracting SignatureType
	signatureTypeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Provenance.signature.type')")
	if err == nil && signatureTypeResult.String() != "undefined" {
		s.SignatureType = []byte(signatureTypeResult.String())
	}
	// extracting Target
	targetResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Provenance.target')")
	if err == nil && targetResult.String() != "undefined" {
		s.Target = []byte(targetResult.String())
	}
	// extracting Text
	textResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'text')")
	if err == nil && textResult.String() != "undefined" {
		s.Text = textResult.String()
	}
	// extracting When
	whenResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, '(Provenance.occurredDateTime)')")
	if err == nil && whenResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, whenResult.String())
		if err == nil {
			s.When = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", whenResult.String())
			if err == nil {
				s.When = &d
			}
		}
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirProvenance) TableName() string {
	return "fhir_provenance"
}
