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

type FhirPractitioner struct {
	models.ResourceBase
	// Whether the practitioner record is active
	// https://hl7.org/fhir/r4/search.html#token
	Active datatypes.JSON `gorm:"column:active;type:text;serializer:json" json:"active,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): A server defined search that may match any of the string fields in the Address, including line, city, district, state, country, postalCode, and/or text
	   * [Person](person.html): A server defined search that may match any of the string fields in the Address, including line, city, district, state, country, postalCode, and/or text
	   * [Practitioner](practitioner.html): A server defined search that may match any of the string fields in the Address, including line, city, district, state, country, postalCode, and/or text
	   * [RelatedPerson](relatedperson.html): A server defined search that may match any of the string fields in the Address, including line, city, district, state, country, postalCode, and/or text
	*/
	// https://hl7.org/fhir/r4/search.html#string
	Address datatypes.JSON `gorm:"column:address;type:text;serializer:json" json:"address,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): A city specified in an address
	   * [Person](person.html): A city specified in an address
	   * [Practitioner](practitioner.html): A city specified in an address
	   * [RelatedPerson](relatedperson.html): A city specified in an address
	*/
	// https://hl7.org/fhir/r4/search.html#string
	AddressCity datatypes.JSON `gorm:"column:addressCity;type:text;serializer:json" json:"addressCity,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): A country specified in an address
	   * [Person](person.html): A country specified in an address
	   * [Practitioner](practitioner.html): A country specified in an address
	   * [RelatedPerson](relatedperson.html): A country specified in an address
	*/
	// https://hl7.org/fhir/r4/search.html#string
	AddressCountry datatypes.JSON `gorm:"column:addressCountry;type:text;serializer:json" json:"addressCountry,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): A postalCode specified in an address
	   * [Person](person.html): A postal code specified in an address
	   * [Practitioner](practitioner.html): A postalCode specified in an address
	   * [RelatedPerson](relatedperson.html): A postal code specified in an address
	*/
	// https://hl7.org/fhir/r4/search.html#string
	AddressPostalcode datatypes.JSON `gorm:"column:addressPostalcode;type:text;serializer:json" json:"addressPostalcode,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): A state specified in an address
	   * [Person](person.html): A state specified in an address
	   * [Practitioner](practitioner.html): A state specified in an address
	   * [RelatedPerson](relatedperson.html): A state specified in an address
	*/
	// https://hl7.org/fhir/r4/search.html#string
	AddressState datatypes.JSON `gorm:"column:addressState;type:text;serializer:json" json:"addressState,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): A use code specified in an address
	   * [Person](person.html): A use code specified in an address
	   * [Practitioner](practitioner.html): A use code specified in an address
	   * [RelatedPerson](relatedperson.html): A use code specified in an address
	*/
	// https://hl7.org/fhir/r4/search.html#token
	AddressUse datatypes.JSON `gorm:"column:addressUse;type:text;serializer:json" json:"addressUse,omitempty"`
	// One of the languages that the practitioner can communicate with
	// https://hl7.org/fhir/r4/search.html#token
	Communication datatypes.JSON `gorm:"column:communication;type:text;serializer:json" json:"communication,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): A value in an email contact
	   * [Person](person.html): A value in an email contact
	   * [Practitioner](practitioner.html): A value in an email contact
	   * [PractitionerRole](practitionerrole.html): A value in an email contact
	   * [RelatedPerson](relatedperson.html): A value in an email contact
	*/
	// https://hl7.org/fhir/r4/search.html#token
	Email datatypes.JSON `gorm:"column:email;type:text;serializer:json" json:"email,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): A portion of the family name of the patient
	   * [Practitioner](practitioner.html): A portion of the family name
	*/
	// https://hl7.org/fhir/r4/search.html#string
	Family datatypes.JSON `gorm:"column:family;type:text;serializer:json" json:"family,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): Gender of the patient
	   * [Person](person.html): The gender of the person
	   * [Practitioner](practitioner.html): Gender of the practitioner
	   * [RelatedPerson](relatedperson.html): Gender of the related person
	*/
	// https://hl7.org/fhir/r4/search.html#token
	Gender datatypes.JSON `gorm:"column:gender;type:text;serializer:json" json:"gender,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): A portion of the given name of the patient
	   * [Practitioner](practitioner.html): A portion of the given name
	*/
	// https://hl7.org/fhir/r4/search.html#string
	Given datatypes.JSON `gorm:"column:given;type:text;serializer:json" json:"given,omitempty"`
	// A practitioner's Identifier
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
	// A server defined search that may match any of the string fields in the HumanName, including family, give, prefix, suffix, suffix, and/or text
	// https://hl7.org/fhir/r4/search.html#string
	Name datatypes.JSON `gorm:"column:name;type:text;serializer:json" json:"name,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): A value in a phone contact
	   * [Person](person.html): A value in a phone contact
	   * [Practitioner](practitioner.html): A value in a phone contact
	   * [PractitionerRole](practitionerrole.html): A value in a phone contact
	   * [RelatedPerson](relatedperson.html): A value in a phone contact
	*/
	// https://hl7.org/fhir/r4/search.html#token
	Phone datatypes.JSON `gorm:"column:phone;type:text;serializer:json" json:"phone,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): A portion of either family or given name using some kind of phonetic matching algorithm
	   * [Person](person.html): A portion of name using some kind of phonetic matching algorithm
	   * [Practitioner](practitioner.html): A portion of either family or given name using some kind of phonetic matching algorithm
	   * [RelatedPerson](relatedperson.html): A portion of name using some kind of phonetic matching algorithm
	*/
	// https://hl7.org/fhir/r4/search.html#string
	Phonetic datatypes.JSON `gorm:"column:phonetic;type:text;serializer:json" json:"phonetic,omitempty"`
	/*
	   Multiple Resources:

	   * [Patient](patient.html): The value in any kind of telecom details of the patient
	   * [Person](person.html): The value in any kind of contact
	   * [Practitioner](practitioner.html): The value in any kind of contact
	   * [PractitionerRole](practitionerrole.html): The value in any kind of contact
	   * [RelatedPerson](relatedperson.html): The value in any kind of contact
	*/
	// https://hl7.org/fhir/r4/search.html#token
	Telecom datatypes.JSON `gorm:"column:telecom;type:text;serializer:json" json:"telecom,omitempty"`
	// Text search against the narrative
	// https://hl7.org/fhir/r4/search.html#string
	Text datatypes.JSON `gorm:"column:text;type:text;serializer:json" json:"text,omitempty"`
}

func (s *FhirPractitioner) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"active":               "token",
		"address":              "string",
		"addressCity":          "string",
		"addressCountry":       "string",
		"addressPostalcode":    "string",
		"addressState":         "string",
		"addressUse":           "token",
		"communication":        "token",
		"email":                "token",
		"family":               "string",
		"gender":               "token",
		"given":                "string",
		"id":                   "keyword",
		"identifier":           "token",
		"language":             "token",
		"metaLastUpdated":      "date",
		"metaProfile":          "reference",
		"metaTag":              "token",
		"metaVersionId":        "keyword",
		"name":                 "string",
		"phone":                "token",
		"phonetic":             "string",
		"sort_date":            "date",
		"source_id":            "keyword",
		"source_resource_id":   "keyword",
		"source_resource_type": "keyword",
		"source_uri":           "keyword",
		"telecom":              "token",
		"text":                 "string",
	}
	return searchParameters
}
func (s *FhirPractitioner) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
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
	// extracting Active
	activeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Practitioner.active')")
	if err == nil && activeResult.String() != "undefined" {
		s.Active = []byte(activeResult.String())
	}
	// extracting Address
	addressResult, err := vm.RunString("extractStringSearchParameters(fhirResource, 'Patient.address | Person.address | Practitioner.address | RelatedPerson.address')")
	if err == nil && addressResult.String() != "undefined" {
		s.Address = []byte(addressResult.String())
	}
	// extracting AddressCity
	addressCityResult, err := vm.RunString("extractStringSearchParameters(fhirResource, 'Patient.address.city | Person.address.city | Practitioner.address.city | RelatedPerson.address.city')")
	if err == nil && addressCityResult.String() != "undefined" {
		s.AddressCity = []byte(addressCityResult.String())
	}
	// extracting AddressCountry
	addressCountryResult, err := vm.RunString("extractStringSearchParameters(fhirResource, 'Patient.address.country | Person.address.country | Practitioner.address.country | RelatedPerson.address.country')")
	if err == nil && addressCountryResult.String() != "undefined" {
		s.AddressCountry = []byte(addressCountryResult.String())
	}
	// extracting AddressPostalcode
	addressPostalcodeResult, err := vm.RunString("extractStringSearchParameters(fhirResource, 'Patient.address.postalCode | Person.address.postalCode | Practitioner.address.postalCode | RelatedPerson.address.postalCode')")
	if err == nil && addressPostalcodeResult.String() != "undefined" {
		s.AddressPostalcode = []byte(addressPostalcodeResult.String())
	}
	// extracting AddressState
	addressStateResult, err := vm.RunString("extractStringSearchParameters(fhirResource, 'Patient.address.state | Person.address.state | Practitioner.address.state | RelatedPerson.address.state')")
	if err == nil && addressStateResult.String() != "undefined" {
		s.AddressState = []byte(addressStateResult.String())
	}
	// extracting AddressUse
	addressUseResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Patient.address.use | Person.address.use | Practitioner.address.use | RelatedPerson.address.use')")
	if err == nil && addressUseResult.String() != "undefined" {
		s.AddressUse = []byte(addressUseResult.String())
	}
	// extracting Communication
	communicationResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Practitioner.communication')")
	if err == nil && communicationResult.String() != "undefined" {
		s.Communication = []byte(communicationResult.String())
	}
	// extracting Email
	emailResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Patient.telecom.where(system='email') | Person.telecom.where(system='email') | Practitioner.telecom.where(system='email') | PractitionerRole.telecom.where(system='email') | RelatedPerson.telecom.where(system='email')')")
	if err == nil && emailResult.String() != "undefined" {
		s.Email = []byte(emailResult.String())
	}
	// extracting Family
	familyResult, err := vm.RunString("extractStringSearchParameters(fhirResource, 'Patient.name.family | Practitioner.name.family')")
	if err == nil && familyResult.String() != "undefined" {
		s.Family = []byte(familyResult.String())
	}
	// extracting Gender
	genderResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Patient.gender | Person.gender | Practitioner.gender | RelatedPerson.gender')")
	if err == nil && genderResult.String() != "undefined" {
		s.Gender = []byte(genderResult.String())
	}
	// extracting Given
	givenResult, err := vm.RunString("extractStringSearchParameters(fhirResource, 'Patient.name.given | Practitioner.name.given')")
	if err == nil && givenResult.String() != "undefined" {
		s.Given = []byte(givenResult.String())
	}
	// extracting Identifier
	identifierResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Practitioner.identifier')")
	if err == nil && identifierResult.String() != "undefined" {
		s.Identifier = []byte(identifierResult.String())
	}
	// extracting Language
	languageResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'language')")
	if err == nil && languageResult.String() != "undefined" {
		s.Language = []byte(languageResult.String())
	}
	// extracting MetaLastUpdated
	metaLastUpdatedResult, err := vm.RunString("extractDateSearchParameters(fhirResource, 'meta.lastUpdated')")
	if err == nil && metaLastUpdatedResult.String() != "undefined" {
		if t, err := time.Parse(time.RFC3339, metaLastUpdatedResult.String()); err == nil {
			s.MetaLastUpdated = &t
		} else if t, err = time.Parse("2006-01-02", metaLastUpdatedResult.String()); err == nil {
			s.MetaLastUpdated = &t
		} else if t, err = time.Parse("2006-01", metaLastUpdatedResult.String()); err == nil {
			s.MetaLastUpdated = &t
		} else if t, err = time.Parse("2006", metaLastUpdatedResult.String()); err == nil {
			s.MetaLastUpdated = &t
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
	// extracting Name
	nameResult, err := vm.RunString("extractStringSearchParameters(fhirResource, 'Practitioner.name')")
	if err == nil && nameResult.String() != "undefined" {
		s.Name = []byte(nameResult.String())
	}
	// extracting Phone
	phoneResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Patient.telecom.where(system='phone') | Person.telecom.where(system='phone') | Practitioner.telecom.where(system='phone') | PractitionerRole.telecom.where(system='phone') | RelatedPerson.telecom.where(system='phone')')")
	if err == nil && phoneResult.String() != "undefined" {
		s.Phone = []byte(phoneResult.String())
	}
	// extracting Phonetic
	phoneticResult, err := vm.RunString("extractStringSearchParameters(fhirResource, 'Patient.name | Person.name | Practitioner.name | RelatedPerson.name')")
	if err == nil && phoneticResult.String() != "undefined" {
		s.Phonetic = []byte(phoneticResult.String())
	}
	// extracting Telecom
	telecomResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Patient.telecom | Person.telecom | Practitioner.telecom | PractitionerRole.telecom | RelatedPerson.telecom')")
	if err == nil && telecomResult.String() != "undefined" {
		s.Telecom = []byte(telecomResult.String())
	}
	// extracting Text
	textResult, err := vm.RunString("extractStringSearchParameters(fhirResource, 'text')")
	if err == nil && textResult.String() != "undefined" {
		s.Text = []byte(textResult.String())
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirPractitioner) TableName() string {
	return "fhir_practitioner"
}
