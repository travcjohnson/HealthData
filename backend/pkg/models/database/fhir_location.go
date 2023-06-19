// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-onprem/blob/main/backend/pkg/models/database/generate.go
// PLEASE DO NOT EDIT BY HAND

package database

import (
	"encoding/json"
	"fmt"
	goja "github.com/dop251/goja"
	models "github.com/fastenhealth/fastenhealth-onprem/backend/pkg/models"
	datatypes "gorm.io/datatypes"
	"time"
)

type FhirLocation struct {
	models.OriginBase
	// The raw resource content in JSON format
	ResourceRaw datatypes.JSON `gorm:"column:resource_raw;type:text;serializer:json" json:"resource_raw,omitempty"`
	// A (part of the) address of the location
	// https://hl7.org/fhir/r4/search.html#string
	Address string `gorm:"column:address;type:text" json:"address,omitempty"`
	// A city specified in an address
	// https://hl7.org/fhir/r4/search.html#string
	AddressCity string `gorm:"column:addressCity;type:text" json:"addressCity,omitempty"`
	// A country specified in an address
	// https://hl7.org/fhir/r4/search.html#string
	AddressCountry string `gorm:"column:addressCountry;type:text" json:"addressCountry,omitempty"`
	// A postal code specified in an address
	// https://hl7.org/fhir/r4/search.html#string
	AddressPostalcode string `gorm:"column:addressPostalcode;type:text" json:"addressPostalcode,omitempty"`
	// A state specified in an address
	// https://hl7.org/fhir/r4/search.html#string
	AddressState string `gorm:"column:addressState;type:text" json:"addressState,omitempty"`
	// A use code specified in an address
	// https://hl7.org/fhir/r4/search.html#token
	AddressUse datatypes.JSON `gorm:"column:addressUse;type:text;serializer:json" json:"addressUse,omitempty"`
	// Technical endpoints providing access to services operated for the location
	// https://hl7.org/fhir/r4/search.html#reference
	Endpoint datatypes.JSON `gorm:"column:endpoint;type:text;serializer:json" json:"endpoint,omitempty"`
	// An identifier for the location
	// https://hl7.org/fhir/r4/search.html#token
	Identifier datatypes.JSON `gorm:"column:identifier;type:text;serializer:json" json:"identifier,omitempty"`
	// Language of the resource content
	// https://hl7.org/fhir/r4/search.html#token
	Language datatypes.JSON `gorm:"column:language;type:text;serializer:json" json:"language,omitempty"`
	// When the resource version last changed
	// https://hl7.org/fhir/r4/search.html#date
	LastUpdated time.Time `gorm:"column:lastUpdated;type:datetime" json:"lastUpdated,omitempty"`
	// A portion of the location's name or alias
	// https://hl7.org/fhir/r4/search.html#string
	Name string `gorm:"column:name;type:text" json:"name,omitempty"`
	// Searches for locations (typically bed/room) that have an operational status (e.g. contaminated, housekeeping)
	// https://hl7.org/fhir/r4/search.html#token
	OperationalStatus datatypes.JSON `gorm:"column:operationalStatus;type:text;serializer:json" json:"operationalStatus,omitempty"`
	// Searches for locations that are managed by the provided organization
	// https://hl7.org/fhir/r4/search.html#reference
	Organization datatypes.JSON `gorm:"column:organization;type:text;serializer:json" json:"organization,omitempty"`
	// A location of which this location is a part
	// https://hl7.org/fhir/r4/search.html#reference
	Partof datatypes.JSON `gorm:"column:partof;type:text;serializer:json" json:"partof,omitempty"`
	// Profiles this resource claims to conform to
	// https://hl7.org/fhir/r4/search.html#reference
	Profile datatypes.JSON `gorm:"column:profile;type:text;serializer:json" json:"profile,omitempty"`
	// Identifies where the resource comes from
	// https://hl7.org/fhir/r4/search.html#uri
	SourceUri string `gorm:"column:sourceUri;type:text" json:"sourceUri,omitempty"`
	// Searches for locations with a specific kind of status
	// https://hl7.org/fhir/r4/search.html#token
	Status datatypes.JSON `gorm:"column:status;type:text;serializer:json" json:"status,omitempty"`
	// Tags applied to this resource
	// https://hl7.org/fhir/r4/search.html#token
	Tag datatypes.JSON `gorm:"column:tag;type:text;serializer:json" json:"tag,omitempty"`
	// Text search against the narrative
	// https://hl7.org/fhir/r4/search.html#string
	Text string `gorm:"column:text;type:text" json:"text,omitempty"`
	// A resource type filter
	// https://hl7.org/fhir/r4/search.html#special
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
}

func (s *FhirLocation) SetOriginBase(originBase models.OriginBase) {
	s.OriginBase = originBase
}
func (s *FhirLocation) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"address":           "string",
		"addressCity":       "string",
		"addressCountry":    "string",
		"addressPostalcode": "string",
		"addressState":      "string",
		"addressUse":        "token",
		"endpoint":          "reference",
		"identifier":        "token",
		"language":          "token",
		"lastUpdated":       "date",
		"name":              "string",
		"operationalStatus": "token",
		"organization":      "reference",
		"partof":            "reference",
		"profile":           "reference",
		"sourceUri":         "uri",
		"status":            "token",
		"tag":               "token",
		"text":              "string",
		"type":              "special",
	}
	return searchParameters
}
func (s *FhirLocation) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
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
	// add the fhirpath library in the goja vm
	_, err = vm.RunProgram(fhirPathJsProgram)
	if err != nil {
		return err
	}
	// execute the fhirpath expression for each search parameter
	// extracting Address
	addressResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Location.address')[0]")
	if err == nil && addressResult.String() != "undefined" {
		s.Address = addressResult.String()
	}
	// extracting AddressCity
	addressCityResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Location.address.city')[0]")
	if err == nil && addressCityResult.String() != "undefined" {
		s.AddressCity = addressCityResult.String()
	}
	// extracting AddressCountry
	addressCountryResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Location.address.country')[0]")
	if err == nil && addressCountryResult.String() != "undefined" {
		s.AddressCountry = addressCountryResult.String()
	}
	// extracting AddressPostalcode
	addressPostalcodeResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Location.address.postalCode')[0]")
	if err == nil && addressPostalcodeResult.String() != "undefined" {
		s.AddressPostalcode = addressPostalcodeResult.String()
	}
	// extracting AddressState
	addressStateResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Location.address.state')[0]")
	if err == nil && addressStateResult.String() != "undefined" {
		s.AddressState = addressStateResult.String()
	}
	// extracting AddressUse
	addressUseResult, err := vm.RunString(` 
							AddressUseResult = window.fhirpath.evaluate(fhirResource, 'Location.address.use')
							AddressUseProcessed = AddressUseResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							JSON.stringify(AddressUseProcessed)
						 `)
	if err == nil && addressUseResult.String() != "undefined" {
		s.AddressUse = []byte(addressUseResult.String())
	}
	// extracting Endpoint
	endpointResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Location.endpoint'))")
	if err == nil && endpointResult.String() != "undefined" {
		s.Endpoint = []byte(endpointResult.String())
	}
	// extracting Identifier
	identifierResult, err := vm.RunString(` 
							IdentifierResult = window.fhirpath.evaluate(fhirResource, 'Location.identifier')
							IdentifierProcessed = IdentifierResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							JSON.stringify(IdentifierProcessed)
						 `)
	if err == nil && identifierResult.String() != "undefined" {
		s.Identifier = []byte(identifierResult.String())
	}
	// extracting Language
	languageResult, err := vm.RunString(` 
							LanguageResult = window.fhirpath.evaluate(fhirResource, 'Resource.language')
							LanguageProcessed = LanguageResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							JSON.stringify(LanguageProcessed)
						 `)
	if err == nil && languageResult.String() != "undefined" {
		s.Language = []byte(languageResult.String())
	}
	// extracting LastUpdated
	lastUpdatedResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Resource.meta.lastUpdated')[0]")
	if err == nil && lastUpdatedResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, lastUpdatedResult.String())
		if err == nil {
			s.LastUpdated = t
		}
	}
	// extracting Name
	nameResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Location.name | Location.alias')[0]")
	if err == nil && nameResult.String() != "undefined" {
		s.Name = nameResult.String()
	}
	// extracting OperationalStatus
	operationalStatusResult, err := vm.RunString(` 
							OperationalStatusResult = window.fhirpath.evaluate(fhirResource, 'Location.operationalStatus')
							OperationalStatusProcessed = OperationalStatusResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							JSON.stringify(OperationalStatusProcessed)
						 `)
	if err == nil && operationalStatusResult.String() != "undefined" {
		s.OperationalStatus = []byte(operationalStatusResult.String())
	}
	// extracting Organization
	organizationResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Location.managingOrganization'))")
	if err == nil && organizationResult.String() != "undefined" {
		s.Organization = []byte(organizationResult.String())
	}
	// extracting Partof
	partofResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Location.partOf'))")
	if err == nil && partofResult.String() != "undefined" {
		s.Partof = []byte(partofResult.String())
	}
	// extracting Profile
	profileResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Resource.meta.profile'))")
	if err == nil && profileResult.String() != "undefined" {
		s.Profile = []byte(profileResult.String())
	}
	// extracting SourceUri
	sourceUriResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Resource.meta.source')[0]")
	if err == nil && sourceUriResult.String() != "undefined" {
		s.SourceUri = sourceUriResult.String()
	}
	// extracting Status
	statusResult, err := vm.RunString(` 
							StatusResult = window.fhirpath.evaluate(fhirResource, 'Location.status')
							StatusProcessed = StatusResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							JSON.stringify(StatusProcessed)
						 `)
	if err == nil && statusResult.String() != "undefined" {
		s.Status = []byte(statusResult.String())
	}
	// extracting Tag
	tagResult, err := vm.RunString(` 
							TagResult = window.fhirpath.evaluate(fhirResource, 'Resource.meta.tag')
							TagProcessed = TagResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							JSON.stringify(TagProcessed)
						 `)
	if err == nil && tagResult.String() != "undefined" {
		s.Tag = []byte(tagResult.String())
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirLocation) TableName() string {
	return "fhir_location"
}
