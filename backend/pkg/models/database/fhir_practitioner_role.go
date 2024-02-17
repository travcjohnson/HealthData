// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-onprem/blob/main/backend/pkg/models/database/generate.go
// PLEASE DO NOT EDIT BY HAND

package database

import (
	"encoding/json"
	"fmt"
	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown"
	goja "github.com/dop251/goja"
	models "github.com/fastenhealth/fasten-onprem/backend/pkg/models"
	datatypes "gorm.io/datatypes"
	"time"
)

type FhirPractitionerRole struct {
	models.ResourceBase
	// Whether this practitioner role record is in active use
	// https://hl7.org/fhir/r4/search.html#token
	Active datatypes.JSON `gorm:"column:active;type:text;serializer:json" json:"active,omitempty"`
	// The period during which the practitioner is authorized to perform in these role(s)
	// https://hl7.org/fhir/r4/search.html#date
	Date *time.Time `gorm:"column:date;type:datetime" json:"date,omitempty"`
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
	// Technical endpoints providing access to services operated for the practitioner with this role
	// https://hl7.org/fhir/r4/search.html#reference
	Endpoint datatypes.JSON `gorm:"column:endpoint;type:text;serializer:json" json:"endpoint,omitempty"`
	// A practitioner's Identifier
	// https://hl7.org/fhir/r4/search.html#token
	Identifier datatypes.JSON `gorm:"column:identifier;type:text;serializer:json" json:"identifier,omitempty"`
	// Language of the resource content
	// https://hl7.org/fhir/r4/search.html#token
	Language datatypes.JSON `gorm:"column:language;type:text;serializer:json" json:"language,omitempty"`
	// One of the locations at which this practitioner provides care
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
	// The identity of the organization the practitioner represents / acts on behalf of
	// https://hl7.org/fhir/r4/search.html#reference
	Organization datatypes.JSON `gorm:"column:organization;type:text;serializer:json" json:"organization,omitempty"`
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
	// Practitioner that is able to provide the defined services for the organization
	// https://hl7.org/fhir/r4/search.html#reference
	Practitioner datatypes.JSON `gorm:"column:practitioner;type:text;serializer:json" json:"practitioner,omitempty"`
	// The practitioner can perform this role at for the organization
	// https://hl7.org/fhir/r4/search.html#token
	Role datatypes.JSON `gorm:"column:role;type:text;serializer:json" json:"role,omitempty"`
	// The list of healthcare services that this worker provides for this role's Organization/Location(s)
	// https://hl7.org/fhir/r4/search.html#reference
	Service datatypes.JSON `gorm:"column:service;type:text;serializer:json" json:"service,omitempty"`
	// The practitioner has this specialty at an organization
	// https://hl7.org/fhir/r4/search.html#token
	Specialty datatypes.JSON `gorm:"column:specialty;type:text;serializer:json" json:"specialty,omitempty"`
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
	// This is a primitive string literal (`keyword` type). It is not a recognized SearchParameter type from https://hl7.org/fhir/r4/search.html, it's Fasten Health-specific
	Text string `gorm:"column:text;type:text" json:"text,omitempty"`
	// A resource type filter
	// https://hl7.org/fhir/r4/search.html#special
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
}

func (s *FhirPractitionerRole) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"active":               "token",
		"date":                 "date",
		"email":                "token",
		"endpoint":             "reference",
		"id":                   "keyword",
		"identifier":           "token",
		"language":             "token",
		"location":             "reference",
		"metaLastUpdated":      "date",
		"metaProfile":          "reference",
		"metaTag":              "token",
		"metaVersionId":        "keyword",
		"organization":         "reference",
		"phone":                "token",
		"practitioner":         "reference",
		"role":                 "token",
		"service":              "reference",
		"sort_date":            "date",
		"source_id":            "keyword",
		"source_resource_id":   "keyword",
		"source_resource_type": "keyword",
		"source_uri":           "keyword",
		"specialty":            "token",
		"telecom":              "token",
		"text":                 "keyword",
		"type":                 "special",
	}
	return searchParameters
}
func (s *FhirPractitionerRole) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
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
	activeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'PractitionerRole.active')")
	if err == nil && activeResult.String() != "undefined" {
		s.Active = []byte(activeResult.String())
	}
	// extracting Date
	dateResult, err := vm.RunString("extractDateSearchParameters(fhirResource, 'PractitionerRole.period')")
	if err == nil && dateResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, dateResult.String())
		if err == nil {
			s.Date = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", dateResult.String())
			if err == nil {
				s.Date = &d
			}
		}
	}
	// extracting Email
	emailResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Patient.telecom.where(system='email') | Person.telecom.where(system='email') | Practitioner.telecom.where(system='email') | PractitionerRole.telecom.where(system='email') | RelatedPerson.telecom.where(system='email')')")
	if err == nil && emailResult.String() != "undefined" {
		s.Email = []byte(emailResult.String())
	}
	// extracting Endpoint
	endpointResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'PractitionerRole.endpoint')")
	if err == nil && endpointResult.String() != "undefined" {
		s.Endpoint = []byte(endpointResult.String())
	}
	// extracting Identifier
	identifierResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'PractitionerRole.identifier')")
	if err == nil && identifierResult.String() != "undefined" {
		s.Identifier = []byte(identifierResult.String())
	}
	// extracting Language
	languageResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'language')")
	if err == nil && languageResult.String() != "undefined" {
		s.Language = []byte(languageResult.String())
	}
	// extracting Location
	locationResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'PractitionerRole.location')")
	if err == nil && locationResult.String() != "undefined" {
		s.Location = []byte(locationResult.String())
	}
	// extracting MetaLastUpdated
	metaLastUpdatedResult, err := vm.RunString("extractDateSearchParameters(fhirResource, 'meta.lastUpdated')")
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
	// extracting Organization
	organizationResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'PractitionerRole.organization')")
	if err == nil && organizationResult.String() != "undefined" {
		s.Organization = []byte(organizationResult.String())
	}
	// extracting Phone
	phoneResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Patient.telecom.where(system='phone') | Person.telecom.where(system='phone') | Practitioner.telecom.where(system='phone') | PractitionerRole.telecom.where(system='phone') | RelatedPerson.telecom.where(system='phone')')")
	if err == nil && phoneResult.String() != "undefined" {
		s.Phone = []byte(phoneResult.String())
	}
	// extracting Practitioner
	practitionerResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'PractitionerRole.practitioner')")
	if err == nil && practitionerResult.String() != "undefined" {
		s.Practitioner = []byte(practitionerResult.String())
	}
	// extracting Role
	roleResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'PractitionerRole.code')")
	if err == nil && roleResult.String() != "undefined" {
		s.Role = []byte(roleResult.String())
	}
	// extracting Service
	serviceResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'PractitionerRole.healthcareService')")
	if err == nil && serviceResult.String() != "undefined" {
		s.Service = []byte(serviceResult.String())
	}
	// extracting Specialty
	specialtyResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'PractitionerRole.specialty')")
	if err == nil && specialtyResult.String() != "undefined" {
		s.Specialty = []byte(specialtyResult.String())
	}
	// extracting Telecom
	telecomResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Patient.telecom | Person.telecom | Practitioner.telecom | PractitionerRole.telecom | RelatedPerson.telecom')")
	if err == nil && telecomResult.String() != "undefined" {
		s.Telecom = []byte(telecomResult.String())
	}
	// extracting Text
	textResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'text.div')")
	if err == nil && textResult.String() != "undefined" {
		s.Text = textResult.String()
		converter := htmltomarkdown.NewConverter("", true, nil)
		markdown, err := converter.ConvertString(s.Text)
		if err == nil {
			s.Text = markdown
		}
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirPractitionerRole) TableName() string {
	return "fhir_practitioner_role"
}
