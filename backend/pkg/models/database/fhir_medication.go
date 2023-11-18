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

type FhirMedication struct {
	models.ResourceBase
	/*
	   Multiple Resources:

	   * [AllergyIntolerance](allergyintolerance.html): Code that identifies the allergy or intolerance
	   * [Condition](condition.html): Code for the condition
	   * [DeviceRequest](devicerequest.html): Code for what is being requested/ordered
	   * [DiagnosticReport](diagnosticreport.html): The code for the report, as opposed to codes for the atomic results, which are the names on the observation resource referred to from the result
	   * [FamilyMemberHistory](familymemberhistory.html): A search by a condition code
	   * [List](list.html): What the purpose of this list is
	   * [Medication](medication.html): Returns medications for a specific code
	   * [MedicationAdministration](medicationadministration.html): Return administrations of this medication code
	   * [MedicationDispense](medicationdispense.html): Returns dispenses of this medicine code
	   * [MedicationRequest](medicationrequest.html): Return prescriptions of this medication code
	   * [MedicationStatement](medicationstatement.html): Return statements of this medication code
	   * [Observation](observation.html): The code of the observation type
	   * [Procedure](procedure.html): A code to identify a  procedure
	   * [ServiceRequest](servicerequest.html): What is being requested/ordered
	*/
	// https://hl7.org/fhir/r4/search.html#token
	Code datatypes.JSON `gorm:"column:code;type:text;serializer:json" json:"code,omitempty"`
	// Returns medications in a batch with this expiration date
	// https://hl7.org/fhir/r4/search.html#date
	ExpirationDate *time.Time `gorm:"column:expirationDate;type:datetime" json:"expirationDate,omitempty"`
	// Returns medications for a specific dose form
	// https://hl7.org/fhir/r4/search.html#token
	Form datatypes.JSON `gorm:"column:form;type:text;serializer:json" json:"form,omitempty"`
	// Returns medications with this external identifier
	// https://hl7.org/fhir/r4/search.html#token
	Identifier datatypes.JSON `gorm:"column:identifier;type:text;serializer:json" json:"identifier,omitempty"`
	// Returns medications for this ingredient reference
	// https://hl7.org/fhir/r4/search.html#reference
	Ingredient datatypes.JSON `gorm:"column:ingredient;type:text;serializer:json" json:"ingredient,omitempty"`
	// Returns medications for this ingredient code
	// https://hl7.org/fhir/r4/search.html#token
	IngredientCode datatypes.JSON `gorm:"column:ingredientCode;type:text;serializer:json" json:"ingredientCode,omitempty"`
	// Language of the resource content
	// https://hl7.org/fhir/r4/search.html#token
	Language datatypes.JSON `gorm:"column:language;type:text;serializer:json" json:"language,omitempty"`
	// Returns medications in a batch with this lot number
	// https://hl7.org/fhir/r4/search.html#token
	LotNumber datatypes.JSON `gorm:"column:lotNumber;type:text;serializer:json" json:"lotNumber,omitempty"`
	// Returns medications made or sold for this manufacturer
	// https://hl7.org/fhir/r4/search.html#reference
	Manufacturer datatypes.JSON `gorm:"column:manufacturer;type:text;serializer:json" json:"manufacturer,omitempty"`
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
	// Returns medications for this status
	// https://hl7.org/fhir/r4/search.html#token
	Status datatypes.JSON `gorm:"column:status;type:text;serializer:json" json:"status,omitempty"`
	// Text search against the narrative
	// This is a primitive string literal (`keyword` type). It is not a recognized SearchParameter type from https://hl7.org/fhir/r4/search.html, it's Fasten Health-specific
	Text string `gorm:"column:text;type:text" json:"text,omitempty"`
	// A resource type filter
	// https://hl7.org/fhir/r4/search.html#special
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
}

func (s *FhirMedication) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"code":                 "token",
		"expirationDate":       "date",
		"form":                 "token",
		"id":                   "keyword",
		"identifier":           "token",
		"ingredient":           "reference",
		"ingredientCode":       "token",
		"language":             "token",
		"lotNumber":            "token",
		"manufacturer":         "reference",
		"metaLastUpdated":      "date",
		"metaProfile":          "reference",
		"metaTag":              "token",
		"metaVersionId":        "keyword",
		"sort_date":            "date",
		"source_id":            "keyword",
		"source_resource_id":   "keyword",
		"source_resource_type": "keyword",
		"source_uri":           "keyword",
		"status":               "token",
		"text":                 "keyword",
		"type":                 "special",
	}
	return searchParameters
}
func (s *FhirMedication) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
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
	// extracting Code
	codeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'AllergyIntolerance.code | AllergyIntolerance.reaction.substance | Condition.code | (DeviceRequest.codeCodeableConcept) | DiagnosticReport.code | FamilyMemberHistory.condition.code | List.code | Medication.code | (MedicationAdministration.medicationCodeableConcept) | (MedicationDispense.medicationCodeableConcept) | (MedicationRequest.medicationCodeableConcept) | (MedicationStatement.medicationCodeableConcept) | Observation.code | Procedure.code | ServiceRequest.code')")
	if err == nil && codeResult.String() != "undefined" {
		s.Code = []byte(codeResult.String())
	}
	// extracting ExpirationDate
	expirationDateResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'Medication.batch.expirationDate')")
	if err == nil && expirationDateResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, expirationDateResult.String())
		if err == nil {
			s.ExpirationDate = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", expirationDateResult.String())
			if err == nil {
				s.ExpirationDate = &d
			}
		}
	}
	// extracting Form
	formResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Medication.form')")
	if err == nil && formResult.String() != "undefined" {
		s.Form = []byte(formResult.String())
	}
	// extracting Identifier
	identifierResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Medication.identifier')")
	if err == nil && identifierResult.String() != "undefined" {
		s.Identifier = []byte(identifierResult.String())
	}
	// extracting Ingredient
	ingredientResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, '(Medication.ingredient.itemReference)')")
	if err == nil && ingredientResult.String() != "undefined" {
		s.Ingredient = []byte(ingredientResult.String())
	}
	// extracting IngredientCode
	ingredientCodeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, '(Medication.ingredient.itemCodeableConcept)')")
	if err == nil && ingredientCodeResult.String() != "undefined" {
		s.IngredientCode = []byte(ingredientCodeResult.String())
	}
	// extracting Language
	languageResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'language')")
	if err == nil && languageResult.String() != "undefined" {
		s.Language = []byte(languageResult.String())
	}
	// extracting LotNumber
	lotNumberResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Medication.batch.lotNumber')")
	if err == nil && lotNumberResult.String() != "undefined" {
		s.LotNumber = []byte(lotNumberResult.String())
	}
	// extracting Manufacturer
	manufacturerResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Medication.manufacturer')")
	if err == nil && manufacturerResult.String() != "undefined" {
		s.Manufacturer = []byte(manufacturerResult.String())
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
	// extracting Status
	statusResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Medication.status')")
	if err == nil && statusResult.String() != "undefined" {
		s.Status = []byte(statusResult.String())
	}
	// extracting Text
	textResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'text')")
	if err == nil && textResult.String() != "undefined" {
		s.Text = textResult.String()
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirMedication) TableName() string {
	return "fhir_medication"
}
