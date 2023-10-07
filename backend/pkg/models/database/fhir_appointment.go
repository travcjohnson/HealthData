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

type FhirAppointment struct {
	models.ResourceBase
	// Any one of the individuals participating in the appointment
	// https://hl7.org/fhir/r4/search.html#reference
	Actor datatypes.JSON `gorm:"column:actor;type:text;serializer:json" json:"actor,omitempty"`
	// The style of appointment or patient that has been booked in the slot (not service type)
	// https://hl7.org/fhir/r4/search.html#token
	AppointmentType datatypes.JSON `gorm:"column:appointmentType;type:text;serializer:json" json:"appointmentType,omitempty"`
	// The service request this appointment is allocated to assess
	// https://hl7.org/fhir/r4/search.html#reference
	BasedOn datatypes.JSON `gorm:"column:basedOn;type:text;serializer:json" json:"basedOn,omitempty"`
	// Appointment date/time.
	// https://hl7.org/fhir/r4/search.html#date
	Date *time.Time `gorm:"column:date;type:datetime" json:"date,omitempty"`
	// An Identifier of the Appointment
	// https://hl7.org/fhir/r4/search.html#token
	Identifier datatypes.JSON `gorm:"column:identifier;type:text;serializer:json" json:"identifier,omitempty"`
	// Language of the resource content
	// https://hl7.org/fhir/r4/search.html#token
	Language datatypes.JSON `gorm:"column:language;type:text;serializer:json" json:"language,omitempty"`
	// When the resource version last changed
	// https://hl7.org/fhir/r4/search.html#date
	LastUpdated *time.Time `gorm:"column:lastUpdated;type:datetime" json:"lastUpdated,omitempty"`
	// This location is listed in the participants of the appointment
	// https://hl7.org/fhir/r4/search.html#reference
	Location datatypes.JSON `gorm:"column:location;type:text;serializer:json" json:"location,omitempty"`
	// The Participation status of the subject, or other participant on the appointment. Can be used to locate participants that have not responded to meeting requests.
	// https://hl7.org/fhir/r4/search.html#token
	PartStatus datatypes.JSON `gorm:"column:partStatus;type:text;serializer:json" json:"partStatus,omitempty"`
	// One of the individuals of the appointment is this practitioner
	// https://hl7.org/fhir/r4/search.html#reference
	Practitioner datatypes.JSON `gorm:"column:practitioner;type:text;serializer:json" json:"practitioner,omitempty"`
	// Profiles this resource claims to conform to
	// https://hl7.org/fhir/r4/search.html#reference
	Profile datatypes.JSON `gorm:"column:profile;type:text;serializer:json" json:"profile,omitempty"`
	// Coded reason this appointment is scheduled
	// https://hl7.org/fhir/r4/search.html#token
	ReasonCode datatypes.JSON `gorm:"column:reasonCode;type:text;serializer:json" json:"reasonCode,omitempty"`
	// Reason the appointment is to take place (resource)
	// https://hl7.org/fhir/r4/search.html#reference
	ReasonReference datatypes.JSON `gorm:"column:reasonReference;type:text;serializer:json" json:"reasonReference,omitempty"`
	// A broad categorization of the service that is to be performed during this appointment
	// https://hl7.org/fhir/r4/search.html#token
	ServiceCategory datatypes.JSON `gorm:"column:serviceCategory;type:text;serializer:json" json:"serviceCategory,omitempty"`
	// The specific service that is to be performed during this appointment
	// https://hl7.org/fhir/r4/search.html#token
	ServiceType datatypes.JSON `gorm:"column:serviceType;type:text;serializer:json" json:"serviceType,omitempty"`
	// The slots that this appointment is filling
	// https://hl7.org/fhir/r4/search.html#reference
	Slot datatypes.JSON `gorm:"column:slot;type:text;serializer:json" json:"slot,omitempty"`
	// The specialty of a practitioner that would be required to perform the service requested in this appointment
	// https://hl7.org/fhir/r4/search.html#token
	Specialty datatypes.JSON `gorm:"column:specialty;type:text;serializer:json" json:"specialty,omitempty"`
	// The overall status of the appointment
	// https://hl7.org/fhir/r4/search.html#token
	Status datatypes.JSON `gorm:"column:status;type:text;serializer:json" json:"status,omitempty"`
	// Additional information to support the appointment
	// https://hl7.org/fhir/r4/search.html#reference
	SupportingInfo datatypes.JSON `gorm:"column:supportingInfo;type:text;serializer:json" json:"supportingInfo,omitempty"`
	// Tags applied to this resource
	// https://hl7.org/fhir/r4/search.html#token
	Tag datatypes.JSON `gorm:"column:tag;type:text;serializer:json" json:"tag,omitempty"`
	// Text search against the narrative
	// https://hl7.org/fhir/r4/search.html#string
	Text datatypes.JSON `gorm:"column:text;type:text;serializer:json" json:"text,omitempty"`
	// A resource type filter
	// https://hl7.org/fhir/r4/search.html#special
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
}

func (s *FhirAppointment) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"actor":                "reference",
		"appointmentType":      "token",
		"basedOn":              "reference",
		"date":                 "date",
		"id":                   "keyword",
		"identifier":           "token",
		"language":             "token",
		"lastUpdated":          "date",
		"location":             "reference",
		"partStatus":           "token",
		"practitioner":         "reference",
		"profile":              "reference",
		"reasonCode":           "token",
		"reasonReference":      "reference",
		"serviceCategory":      "token",
		"serviceType":          "token",
		"slot":                 "reference",
		"sort_date":            "date",
		"source_id":            "keyword",
		"source_resource_id":   "keyword",
		"source_resource_type": "keyword",
		"source_uri":           "keyword",
		"specialty":            "token",
		"status":               "token",
		"supportingInfo":       "reference",
		"tag":                  "token",
		"text":                 "string",
		"type":                 "special",
	}
	return searchParameters
}
func (s *FhirAppointment) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
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
	// extracting Actor
	actorResult, err := vm.RunString(` 
							ActorResult = window.fhirpath.evaluate(fhirResource, 'Appointment.participant.actor')
						
							if(ActorResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ActorResult)
							}
						 `)
	if err == nil && actorResult.String() != "undefined" {
		s.Actor = []byte(actorResult.String())
	}
	// extracting AppointmentType
	appointmentTypeResult, err := vm.RunString(` 
							AppointmentTypeResult = window.fhirpath.evaluate(fhirResource, 'Appointment.appointmentType')
							AppointmentTypeProcessed = AppointmentTypeResult.reduce((accumulator, currentValue) => {
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
						
				
							if(AppointmentTypeProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(AppointmentTypeProcessed)
							}
						 `)
	if err == nil && appointmentTypeResult.String() != "undefined" {
		s.AppointmentType = []byte(appointmentTypeResult.String())
	}
	// extracting BasedOn
	basedOnResult, err := vm.RunString(` 
							BasedOnResult = window.fhirpath.evaluate(fhirResource, 'Appointment.basedOn')
						
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
	// extracting Date
	dateResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Appointment.start')[0]")
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
	// extracting Identifier
	identifierResult, err := vm.RunString(` 
							IdentifierResult = window.fhirpath.evaluate(fhirResource, 'Appointment.identifier')
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
	// extracting Location
	locationResult, err := vm.RunString(` 
							LocationResult = window.fhirpath.evaluate(fhirResource, 'Appointment.participant.actor.where(resolve() is Location)')
						
							if(LocationResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(LocationResult)
							}
						 `)
	if err == nil && locationResult.String() != "undefined" {
		s.Location = []byte(locationResult.String())
	}
	// extracting PartStatus
	partStatusResult, err := vm.RunString(` 
							PartStatusResult = window.fhirpath.evaluate(fhirResource, 'Appointment.participant.status')
							PartStatusProcessed = PartStatusResult.reduce((accumulator, currentValue) => {
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
						
				
							if(PartStatusProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(PartStatusProcessed)
							}
						 `)
	if err == nil && partStatusResult.String() != "undefined" {
		s.PartStatus = []byte(partStatusResult.String())
	}
	// extracting Practitioner
	practitionerResult, err := vm.RunString(` 
							PractitionerResult = window.fhirpath.evaluate(fhirResource, 'Appointment.participant.actor.where(resolve() is Practitioner)')
						
							if(PractitionerResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(PractitionerResult)
							}
						 `)
	if err == nil && practitionerResult.String() != "undefined" {
		s.Practitioner = []byte(practitionerResult.String())
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
	// extracting ReasonCode
	reasonCodeResult, err := vm.RunString(` 
							ReasonCodeResult = window.fhirpath.evaluate(fhirResource, 'Appointment.reasonCode')
							ReasonCodeProcessed = ReasonCodeResult.reduce((accumulator, currentValue) => {
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
						
				
							if(ReasonCodeProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ReasonCodeProcessed)
							}
						 `)
	if err == nil && reasonCodeResult.String() != "undefined" {
		s.ReasonCode = []byte(reasonCodeResult.String())
	}
	// extracting ReasonReference
	reasonReferenceResult, err := vm.RunString(` 
							ReasonReferenceResult = window.fhirpath.evaluate(fhirResource, 'Appointment.reasonReference')
						
							if(ReasonReferenceResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ReasonReferenceResult)
							}
						 `)
	if err == nil && reasonReferenceResult.String() != "undefined" {
		s.ReasonReference = []byte(reasonReferenceResult.String())
	}
	// extracting ServiceCategory
	serviceCategoryResult, err := vm.RunString(` 
							ServiceCategoryResult = window.fhirpath.evaluate(fhirResource, 'Appointment.serviceCategory')
							ServiceCategoryProcessed = ServiceCategoryResult.reduce((accumulator, currentValue) => {
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
						
				
							if(ServiceCategoryProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ServiceCategoryProcessed)
							}
						 `)
	if err == nil && serviceCategoryResult.String() != "undefined" {
		s.ServiceCategory = []byte(serviceCategoryResult.String())
	}
	// extracting ServiceType
	serviceTypeResult, err := vm.RunString(` 
							ServiceTypeResult = window.fhirpath.evaluate(fhirResource, 'Appointment.serviceType')
							ServiceTypeProcessed = ServiceTypeResult.reduce((accumulator, currentValue) => {
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
						
				
							if(ServiceTypeProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ServiceTypeProcessed)
							}
						 `)
	if err == nil && serviceTypeResult.String() != "undefined" {
		s.ServiceType = []byte(serviceTypeResult.String())
	}
	// extracting Slot
	slotResult, err := vm.RunString(` 
							SlotResult = window.fhirpath.evaluate(fhirResource, 'Appointment.slot')
						
							if(SlotResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(SlotResult)
							}
						 `)
	if err == nil && slotResult.String() != "undefined" {
		s.Slot = []byte(slotResult.String())
	}
	// extracting Specialty
	specialtyResult, err := vm.RunString(` 
							SpecialtyResult = window.fhirpath.evaluate(fhirResource, 'Appointment.specialty')
							SpecialtyProcessed = SpecialtyResult.reduce((accumulator, currentValue) => {
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
						
				
							if(SpecialtyProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(SpecialtyProcessed)
							}
						 `)
	if err == nil && specialtyResult.String() != "undefined" {
		s.Specialty = []byte(specialtyResult.String())
	}
	// extracting JobStatus
	statusResult, err := vm.RunString(` 
							StatusResult = window.fhirpath.evaluate(fhirResource, 'Appointment.status')
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
	// extracting SupportingInfo
	supportingInfoResult, err := vm.RunString(` 
							SupportingInfoResult = window.fhirpath.evaluate(fhirResource, 'Appointment.supportingInformation')
						
							if(SupportingInfoResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(SupportingInfoResult)
							}
						 `)
	if err == nil && supportingInfoResult.String() != "undefined" {
		s.SupportingInfo = []byte(supportingInfoResult.String())
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
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirAppointment) TableName() string {
	return "fhir_appointment"
}
