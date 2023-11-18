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

type FhirQuestionnaireResponse struct {
	models.ResourceBase
	// The author of the questionnaire response
	// https://hl7.org/fhir/r4/search.html#reference
	Author datatypes.JSON `gorm:"column:author;type:text;serializer:json" json:"author,omitempty"`
	// When the questionnaire response was last changed
	// https://hl7.org/fhir/r4/search.html#date
	Authored *time.Time `gorm:"column:authored;type:datetime" json:"authored,omitempty"`
	// Plan/proposal/order fulfilled by this questionnaire response
	// https://hl7.org/fhir/r4/search.html#reference
	BasedOn datatypes.JSON `gorm:"column:basedOn;type:text;serializer:json" json:"basedOn,omitempty"`
	// Encounter associated with the questionnaire response
	// https://hl7.org/fhir/r4/search.html#reference
	Encounter datatypes.JSON `gorm:"column:encounter;type:text;serializer:json" json:"encounter,omitempty"`
	// The unique identifier for the questionnaire response
	// https://hl7.org/fhir/r4/search.html#token
	Identifier datatypes.JSON `gorm:"column:identifier;type:text;serializer:json" json:"identifier,omitempty"`
	// Language of the resource content
	// https://hl7.org/fhir/r4/search.html#token
	Language datatypes.JSON `gorm:"column:language;type:text;serializer:json" json:"language,omitempty"`
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
	// Procedure or observation this questionnaire response was performed as a part of
	// https://hl7.org/fhir/r4/search.html#reference
	PartOf datatypes.JSON `gorm:"column:partOf;type:text;serializer:json" json:"partOf,omitempty"`
	// The questionnaire the answers are provided for
	// https://hl7.org/fhir/r4/search.html#reference
	Questionnaire datatypes.JSON `gorm:"column:questionnaire;type:text;serializer:json" json:"questionnaire,omitempty"`
	// The individual providing the information reflected in the questionnaire respose
	// https://hl7.org/fhir/r4/search.html#reference
	Source datatypes.JSON `gorm:"column:source;type:text;serializer:json" json:"source,omitempty"`
	// The status of the questionnaire response
	// https://hl7.org/fhir/r4/search.html#token
	Status datatypes.JSON `gorm:"column:status;type:text;serializer:json" json:"status,omitempty"`
	// The subject of the questionnaire response
	// https://hl7.org/fhir/r4/search.html#reference
	Subject datatypes.JSON `gorm:"column:subject;type:text;serializer:json" json:"subject,omitempty"`
	// Text search against the narrative
	// This is a primitive string literal (`keyword` type). It is not a recognized SearchParameter type from https://hl7.org/fhir/r4/search.html, it's Fasten Health-specific
	Text string `gorm:"column:text;type:text" json:"text,omitempty"`
	// A resource type filter
	// https://hl7.org/fhir/r4/search.html#special
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
}

func (s *FhirQuestionnaireResponse) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"author":               "reference",
		"authored":             "date",
		"basedOn":              "reference",
		"encounter":            "reference",
		"id":                   "keyword",
		"identifier":           "token",
		"language":             "token",
		"metaLastUpdated":      "date",
		"metaProfile":          "reference",
		"metaTag":              "token",
		"metaVersionId":        "keyword",
		"partOf":               "reference",
		"questionnaire":        "reference",
		"sort_date":            "date",
		"source":               "reference",
		"source_id":            "keyword",
		"source_resource_id":   "keyword",
		"source_resource_type": "keyword",
		"source_uri":           "keyword",
		"status":               "token",
		"subject":              "reference",
		"text":                 "keyword",
		"type":                 "special",
	}
	return searchParameters
}
func (s *FhirQuestionnaireResponse) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
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
	// extracting Author
	authorResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'QuestionnaireResponse.author')")
	if err == nil && authorResult.String() != "undefined" {
		s.Author = []byte(authorResult.String())
	}
	// extracting Authored
	authoredResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'QuestionnaireResponse.authored')")
	if err == nil && authoredResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, authoredResult.String())
		if err == nil {
			s.Authored = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", authoredResult.String())
			if err == nil {
				s.Authored = &d
			}
		}
	}
	// extracting BasedOn
	basedOnResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'QuestionnaireResponse.basedOn')")
	if err == nil && basedOnResult.String() != "undefined" {
		s.BasedOn = []byte(basedOnResult.String())
	}
	// extracting Encounter
	encounterResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'QuestionnaireResponse.encounter')")
	if err == nil && encounterResult.String() != "undefined" {
		s.Encounter = []byte(encounterResult.String())
	}
	// extracting Identifier
	identifierResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'QuestionnaireResponse.identifier')")
	if err == nil && identifierResult.String() != "undefined" {
		s.Identifier = []byte(identifierResult.String())
	}
	// extracting Language
	languageResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'language')")
	if err == nil && languageResult.String() != "undefined" {
		s.Language = []byte(languageResult.String())
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
	// extracting PartOf
	partOfResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'QuestionnaireResponse.partOf')")
	if err == nil && partOfResult.String() != "undefined" {
		s.PartOf = []byte(partOfResult.String())
	}
	// extracting Questionnaire
	questionnaireResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'QuestionnaireResponse.questionnaire')")
	if err == nil && questionnaireResult.String() != "undefined" {
		s.Questionnaire = []byte(questionnaireResult.String())
	}
	// extracting Source
	sourceResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'QuestionnaireResponse.source')")
	if err == nil && sourceResult.String() != "undefined" {
		s.Source = []byte(sourceResult.String())
	}
	// extracting Status
	statusResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'QuestionnaireResponse.status')")
	if err == nil && statusResult.String() != "undefined" {
		s.Status = []byte(statusResult.String())
	}
	// extracting Subject
	subjectResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'QuestionnaireResponse.subject')")
	if err == nil && subjectResult.String() != "undefined" {
		s.Subject = []byte(subjectResult.String())
	}
	// extracting Text
	textResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'text')")
	if err == nil && textResult.String() != "undefined" {
		s.Text = textResult.String()
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirQuestionnaireResponse) TableName() string {
	return "fhir_questionnaire_response"
}
