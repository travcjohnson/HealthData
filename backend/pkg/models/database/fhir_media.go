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

type FhirMedia struct {
	models.ResourceBase
	// Procedure that caused this media to be created
	// https://hl7.org/fhir/r4/search.html#reference
	BasedOn datatypes.JSON `gorm:"column:basedOn;type:text;serializer:json" json:"basedOn,omitempty"`
	// When Media was collected
	// https://hl7.org/fhir/r4/search.html#date
	Created *time.Time `gorm:"column:created;type:datetime" json:"created,omitempty"`
	// Observing Device
	// https://hl7.org/fhir/r4/search.html#reference
	Device datatypes.JSON `gorm:"column:device;type:text;serializer:json" json:"device,omitempty"`
	// Encounter associated with media
	// https://hl7.org/fhir/r4/search.html#reference
	Encounter datatypes.JSON `gorm:"column:encounter;type:text;serializer:json" json:"encounter,omitempty"`
	// Identifier(s) for the image
	// https://hl7.org/fhir/r4/search.html#token
	Identifier datatypes.JSON `gorm:"column:identifier;type:text;serializer:json" json:"identifier,omitempty"`
	// Language of the resource content
	// https://hl7.org/fhir/r4/search.html#token
	Language datatypes.JSON `gorm:"column:language;type:text;serializer:json" json:"language,omitempty"`
	// When the resource version last changed
	// https://hl7.org/fhir/r4/search.html#date
	LastUpdated *time.Time `gorm:"column:lastUpdated;type:datetime" json:"lastUpdated,omitempty"`
	// The type of acquisition equipment/process
	// https://hl7.org/fhir/r4/search.html#token
	Modality datatypes.JSON `gorm:"column:modality;type:text;serializer:json" json:"modality,omitempty"`
	// The person who generated the image
	// https://hl7.org/fhir/r4/search.html#reference
	Operator datatypes.JSON `gorm:"column:operator;type:text;serializer:json" json:"operator,omitempty"`
	// Profiles this resource claims to conform to
	// https://hl7.org/fhir/r4/search.html#reference
	Profile datatypes.JSON `gorm:"column:profile;type:text;serializer:json" json:"profile,omitempty"`
	// Observed body part
	// https://hl7.org/fhir/r4/search.html#token
	Site datatypes.JSON `gorm:"column:site;type:text;serializer:json" json:"site,omitempty"`
	// preparation | in-progress | not-done | on-hold | stopped | completed | entered-in-error | unknown
	// https://hl7.org/fhir/r4/search.html#token
	Status datatypes.JSON `gorm:"column:status;type:text;serializer:json" json:"status,omitempty"`
	// Who/What this Media is a record of
	// https://hl7.org/fhir/r4/search.html#reference
	Subject datatypes.JSON `gorm:"column:subject;type:text;serializer:json" json:"subject,omitempty"`
	// Tags applied to this resource
	// https://hl7.org/fhir/r4/search.html#token
	Tag datatypes.JSON `gorm:"column:tag;type:text;serializer:json" json:"tag,omitempty"`
	// Text search against the narrative
	// https://hl7.org/fhir/r4/search.html#string
	Text datatypes.JSON `gorm:"column:text;type:text;serializer:json" json:"text,omitempty"`
	// A resource type filter
	// https://hl7.org/fhir/r4/search.html#special
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
	// Imaging view, e.g. Lateral or Antero-posterior
	// https://hl7.org/fhir/r4/search.html#token
	View datatypes.JSON `gorm:"column:view;type:text;serializer:json" json:"view,omitempty"`
}

func (s *FhirMedia) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"basedOn":              "reference",
		"created":              "date",
		"device":               "reference",
		"encounter":            "reference",
		"id":                   "keyword",
		"identifier":           "token",
		"language":             "token",
		"lastUpdated":          "date",
		"modality":             "token",
		"operator":             "reference",
		"profile":              "reference",
		"site":                 "token",
		"sort_date":            "date",
		"source_id":            "keyword",
		"source_resource_id":   "keyword",
		"source_resource_type": "keyword",
		"source_uri":           "keyword",
		"status":               "token",
		"subject":              "reference",
		"tag":                  "token",
		"text":                 "string",
		"type":                 "special",
		"view":                 "token",
	}
	return searchParameters
}
func (s *FhirMedia) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
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
	// extracting BasedOn
	basedOnResult, err := vm.RunString(` 
							BasedOnResult = window.fhirpath.evaluate(fhirResource, 'Media.basedOn')
						
							if(BasedOnResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(BasedOnResult)
							}
						 `)
	if err == nil && basedOnResult.String() != "undefined" {
		s.BasedOn = []byte(basedOnResult.String())
	}
	// extracting Created
	createdResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Media.created')[0]")
	if err == nil && createdResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, createdResult.String())
		if err == nil {
			s.Created = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", createdResult.String())
			if err == nil {
				s.Created = &d
			}
		}
	}
	// extracting Device
	deviceResult, err := vm.RunString(` 
							DeviceResult = window.fhirpath.evaluate(fhirResource, 'Media.device')
						
							if(DeviceResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(DeviceResult)
							}
						 `)
	if err == nil && deviceResult.String() != "undefined" {
		s.Device = []byte(deviceResult.String())
	}
	// extracting Encounter
	encounterResult, err := vm.RunString(` 
							EncounterResult = window.fhirpath.evaluate(fhirResource, 'Media.encounter')
						
							if(EncounterResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(EncounterResult)
							}
						 `)
	if err == nil && encounterResult.String() != "undefined" {
		s.Encounter = []byte(encounterResult.String())
	}
	// extracting Identifier
	identifierResult, err := vm.RunString(` 
							IdentifierResult = window.fhirpath.evaluate(fhirResource, 'Media.identifier')
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
						
				
							if(IdentifierProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(IdentifierProcessed)
							}
						 `)
	if err == nil && identifierResult.String() != "undefined" {
		s.Identifier = []byte(identifierResult.String())
	}
	// extracting Language
	languageResult, err := vm.RunString(` 
							LanguageResult = window.fhirpath.evaluate(fhirResource, 'language')
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
						
				
							if(LanguageProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(LanguageProcessed)
							}
						 `)
	if err == nil && languageResult.String() != "undefined" {
		s.Language = []byte(languageResult.String())
	}
	// extracting LastUpdated
	lastUpdatedResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'meta.lastUpdated')[0]")
	if err == nil && lastUpdatedResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, lastUpdatedResult.String())
		if err == nil {
			s.LastUpdated = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", lastUpdatedResult.String())
			if err == nil {
				s.LastUpdated = &d
			}
		}
	}
	// extracting Modality
	modalityResult, err := vm.RunString(` 
							ModalityResult = window.fhirpath.evaluate(fhirResource, 'Media.modality')
							ModalityProcessed = ModalityResult.reduce((accumulator, currentValue) => {
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
						
				
							if(ModalityProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ModalityProcessed)
							}
						 `)
	if err == nil && modalityResult.String() != "undefined" {
		s.Modality = []byte(modalityResult.String())
	}
	// extracting Operator
	operatorResult, err := vm.RunString(` 
							OperatorResult = window.fhirpath.evaluate(fhirResource, 'Media.operator')
						
							if(OperatorResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(OperatorResult)
							}
						 `)
	if err == nil && operatorResult.String() != "undefined" {
		s.Operator = []byte(operatorResult.String())
	}
	// extracting Profile
	profileResult, err := vm.RunString(` 
							ProfileResult = window.fhirpath.evaluate(fhirResource, 'meta.profile')
						
							if(ProfileResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ProfileResult)
							}
						 `)
	if err == nil && profileResult.String() != "undefined" {
		s.Profile = []byte(profileResult.String())
	}
	// extracting Site
	siteResult, err := vm.RunString(` 
							SiteResult = window.fhirpath.evaluate(fhirResource, 'Media.bodySite')
							SiteProcessed = SiteResult.reduce((accumulator, currentValue) => {
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
						
				
							if(SiteProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(SiteProcessed)
							}
						 `)
	if err == nil && siteResult.String() != "undefined" {
		s.Site = []byte(siteResult.String())
	}
	// extracting JobStatus
	statusResult, err := vm.RunString(` 
							StatusResult = window.fhirpath.evaluate(fhirResource, 'Media.status')
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
						
				
							if(StatusProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(StatusProcessed)
							}
						 `)
	if err == nil && statusResult.String() != "undefined" {
		s.Status = []byte(statusResult.String())
	}
	// extracting Subject
	subjectResult, err := vm.RunString(` 
							SubjectResult = window.fhirpath.evaluate(fhirResource, 'Media.subject')
						
							if(SubjectResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(SubjectResult)
							}
						 `)
	if err == nil && subjectResult.String() != "undefined" {
		s.Subject = []byte(subjectResult.String())
	}
	// extracting Tag
	tagResult, err := vm.RunString(` 
							TagResult = window.fhirpath.evaluate(fhirResource, 'meta.tag')
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
						
				
							if(TagProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(TagProcessed)
							}
						 `)
	if err == nil && tagResult.String() != "undefined" {
		s.Tag = []byte(tagResult.String())
	}
	// extracting View
	viewResult, err := vm.RunString(` 
							ViewResult = window.fhirpath.evaluate(fhirResource, 'Media.view')
							ViewProcessed = ViewResult.reduce((accumulator, currentValue) => {
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
						
				
							if(ViewProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ViewProcessed)
							}
						 `)
	if err == nil && viewResult.String() != "undefined" {
		s.View = []byte(viewResult.String())
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirMedia) TableName() string {
	return "fhir_media"
}
