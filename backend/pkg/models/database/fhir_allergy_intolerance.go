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

type FhirAllergyIntolerance struct {
	models.OriginBase
	// The raw resource content in JSON format
	ResourceRaw datatypes.JSON `gorm:"column:resource_raw;type:text;serializer:json" json:"resource_raw,omitempty"`
	// Source of the information about the allergy
	// https://hl7.org/fhir/r4/search.html#reference
	Asserter datatypes.JSON `gorm:"column:asserter;type:text;serializer:json" json:"asserter,omitempty"`
	// food | medication | environment | biologic
	// https://hl7.org/fhir/r4/search.html#token
	Category datatypes.JSON `gorm:"column:category;type:text;serializer:json" json:"category,omitempty"`
	// active | inactive | resolved
	// https://hl7.org/fhir/r4/search.html#token
	ClinicalStatus datatypes.JSON `gorm:"column:clinicalStatus;type:text;serializer:json" json:"clinicalStatus,omitempty"`
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
	// low | high | unable-to-assess
	// https://hl7.org/fhir/r4/search.html#token
	Criticality datatypes.JSON `gorm:"column:criticality;type:text;serializer:json" json:"criticality,omitempty"`
	/*
	   Multiple Resources:

	   * [AllergyIntolerance](allergyintolerance.html): Date first version of the resource instance was recorded
	   * [CarePlan](careplan.html): Time period plan covers
	   * [CareTeam](careteam.html): Time period team covers
	   * [ClinicalImpression](clinicalimpression.html): When the assessment was documented
	   * [Composition](composition.html): Composition editing time
	   * [Consent](consent.html): When this Consent was created or indexed
	   * [DiagnosticReport](diagnosticreport.html): The clinically relevant time of the report
	   * [Encounter](encounter.html): A date within the period the Encounter lasted
	   * [EpisodeOfCare](episodeofcare.html): The provided date search value falls within the episode of care's period
	   * [FamilyMemberHistory](familymemberhistory.html): When history was recorded or last updated
	   * [Flag](flag.html): Time period when flag is active
	   * [Immunization](immunization.html): Vaccination  (non)-Administration Date
	   * [List](list.html): When the list was prepared
	   * [Observation](observation.html): Obtained date/time. If the obtained element is a period, a date that falls in the period
	   * [Procedure](procedure.html): When the procedure was performed
	   * [RiskAssessment](riskassessment.html): When was assessment made?
	   * [SupplyRequest](supplyrequest.html): When the request was made
	*/
	// https://hl7.org/fhir/r4/search.html#date
	Date time.Time `gorm:"column:date;type:datetime" json:"date,omitempty"`
	/*
	   Multiple Resources:

	   * [AllergyIntolerance](allergyintolerance.html): External ids for this item
	   * [CarePlan](careplan.html): External Ids for this plan
	   * [CareTeam](careteam.html): External Ids for this team
	   * [Composition](composition.html): Version-independent identifier for the Composition
	   * [Condition](condition.html): A unique identifier of the condition record
	   * [Consent](consent.html): Identifier for this record (external references)
	   * [DetectedIssue](detectedissue.html): Unique id for the detected issue
	   * [DeviceRequest](devicerequest.html): Business identifier for request/order
	   * [DiagnosticReport](diagnosticreport.html): An identifier for the report
	   * [DocumentManifest](documentmanifest.html): Unique Identifier for the set of documents
	   * [DocumentReference](documentreference.html): Master Version Specific Identifier
	   * [Encounter](encounter.html): Identifier(s) by which this encounter is known
	   * [EpisodeOfCare](episodeofcare.html): Business Identifier(s) relevant for this EpisodeOfCare
	   * [FamilyMemberHistory](familymemberhistory.html): A search by a record identifier
	   * [Goal](goal.html): External Ids for this goal
	   * [ImagingStudy](imagingstudy.html): Identifiers for the Study, such as DICOM Study Instance UID and Accession number
	   * [Immunization](immunization.html): Business identifier
	   * [List](list.html): Business identifier
	   * [MedicationAdministration](medicationadministration.html): Return administrations with this external identifier
	   * [MedicationDispense](medicationdispense.html): Returns dispenses with this external identifier
	   * [MedicationRequest](medicationrequest.html): Return prescriptions with this external identifier
	   * [MedicationStatement](medicationstatement.html): Return statements with this external identifier
	   * [NutritionOrder](nutritionorder.html): Return nutrition orders with this external identifier
	   * [Observation](observation.html): The unique id for a particular observation
	   * [Procedure](procedure.html): A unique identifier for a procedure
	   * [RiskAssessment](riskassessment.html): Unique identifier for the assessment
	   * [ServiceRequest](servicerequest.html): Identifiers assigned to this order
	   * [SupplyDelivery](supplydelivery.html): External identifier
	   * [SupplyRequest](supplyrequest.html): Business Identifier for SupplyRequest
	   * [VisionPrescription](visionprescription.html): Return prescriptions with this external identifier
	*/
	// https://hl7.org/fhir/r4/search.html#token
	Identifier datatypes.JSON `gorm:"column:identifier;type:text;serializer:json" json:"identifier,omitempty"`
	// Language of the resource content
	// https://hl7.org/fhir/r4/search.html#token
	Language datatypes.JSON `gorm:"column:language;type:text;serializer:json" json:"language,omitempty"`
	// Date(/time) of last known occurrence of a reaction
	// https://hl7.org/fhir/r4/search.html#date
	LastDate time.Time `gorm:"column:lastDate;type:datetime" json:"lastDate,omitempty"`
	// When the resource version last changed
	// https://hl7.org/fhir/r4/search.html#date
	LastUpdated time.Time `gorm:"column:lastUpdated;type:datetime" json:"lastUpdated,omitempty"`
	// Clinical symptoms/signs associated with the Event
	// https://hl7.org/fhir/r4/search.html#token
	Manifestation datatypes.JSON `gorm:"column:manifestation;type:text;serializer:json" json:"manifestation,omitempty"`
	// Date(/time) when manifestations showed
	// https://hl7.org/fhir/r4/search.html#date
	Onset time.Time `gorm:"column:onset;type:datetime" json:"onset,omitempty"`
	// Profiles this resource claims to conform to
	// https://hl7.org/fhir/r4/search.html#reference
	Profile datatypes.JSON `gorm:"column:profile;type:text;serializer:json" json:"profile,omitempty"`
	// Who recorded the sensitivity
	// https://hl7.org/fhir/r4/search.html#reference
	Recorder datatypes.JSON `gorm:"column:recorder;type:text;serializer:json" json:"recorder,omitempty"`
	// How the subject was exposed to the substance
	// https://hl7.org/fhir/r4/search.html#token
	Route datatypes.JSON `gorm:"column:route;type:text;serializer:json" json:"route,omitempty"`
	// mild | moderate | severe (of event as a whole)
	// https://hl7.org/fhir/r4/search.html#token
	Severity datatypes.JSON `gorm:"column:severity;type:text;serializer:json" json:"severity,omitempty"`
	// Identifies where the resource comes from
	// https://hl7.org/fhir/r4/search.html#uri
	SourceUri string `gorm:"column:sourceUri;type:text" json:"sourceUri,omitempty"`
	// Tags applied to this resource
	// https://hl7.org/fhir/r4/search.html#token
	Tag datatypes.JSON `gorm:"column:tag;type:text;serializer:json" json:"tag,omitempty"`
	// Text search against the narrative
	// https://hl7.org/fhir/r4/search.html#string
	Text string `gorm:"column:text;type:text" json:"text,omitempty"`
	// A resource type filter
	// https://hl7.org/fhir/r4/search.html#special
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
	// unconfirmed | confirmed | refuted | entered-in-error
	// https://hl7.org/fhir/r4/search.html#token
	VerificationStatus datatypes.JSON `gorm:"column:verificationStatus;type:text;serializer:json" json:"verificationStatus,omitempty"`
}

func (s *FhirAllergyIntolerance) SetOriginBase(originBase models.OriginBase) {
	s.OriginBase = originBase
}
func (s *FhirAllergyIntolerance) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"asserter":           "reference",
		"category":           "token",
		"clinicalStatus":     "token",
		"code":               "token",
		"criticality":        "token",
		"date":               "date",
		"identifier":         "token",
		"language":           "token",
		"lastDate":           "date",
		"lastUpdated":        "date",
		"manifestation":      "token",
		"onset":              "date",
		"profile":            "reference",
		"recorder":           "reference",
		"route":              "token",
		"severity":           "token",
		"sourceUri":          "uri",
		"tag":                "token",
		"text":               "string",
		"type":               "special",
		"verificationStatus": "token",
	}
	return searchParameters
}
func (s *FhirAllergyIntolerance) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
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
	// extracting Asserter
	asserterResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'AllergyIntolerance.asserter'))")
	if err == nil && asserterResult.String() != "undefined" {
		s.Asserter = []byte(asserterResult.String())
	}
	// extracting Category
	categoryResult, err := vm.RunString(` 
							CategoryResult = window.fhirpath.evaluate(fhirResource, 'AllergyIntolerance.category')
							CategoryProcessed = CategoryResult.reduce((accumulator, currentValue) => {
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
						
				
							JSON.stringify(CategoryProcessed)
						 `)
	if err == nil && categoryResult.String() != "undefined" {
		s.Category = []byte(categoryResult.String())
	}
	// extracting ClinicalStatus
	clinicalStatusResult, err := vm.RunString(` 
							ClinicalStatusResult = window.fhirpath.evaluate(fhirResource, 'AllergyIntolerance.clinicalStatus')
							ClinicalStatusProcessed = ClinicalStatusResult.reduce((accumulator, currentValue) => {
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
						
				
							JSON.stringify(ClinicalStatusProcessed)
						 `)
	if err == nil && clinicalStatusResult.String() != "undefined" {
		s.ClinicalStatus = []byte(clinicalStatusResult.String())
	}
	// extracting Code
	codeResult, err := vm.RunString(` 
							CodeResult = window.fhirpath.evaluate(fhirResource, 'AllergyIntolerance.code | AllergyIntolerance.reaction.substance | Condition.code | (DeviceRequest.codeCodeableConcept) | DiagnosticReport.code | FamilyMemberHistory.condition.code | List.code | Medication.code | (MedicationAdministration.medicationCodeableConcept) | (MedicationDispense.medicationCodeableConcept) | (MedicationRequest.medicationCodeableConcept) | (MedicationStatement.medicationCodeableConcept) | Observation.code | Procedure.code | ServiceRequest.code')
							CodeProcessed = CodeResult.reduce((accumulator, currentValue) => {
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
						
				
							JSON.stringify(CodeProcessed)
						 `)
	if err == nil && codeResult.String() != "undefined" {
		s.Code = []byte(codeResult.String())
	}
	// extracting Criticality
	criticalityResult, err := vm.RunString(` 
							CriticalityResult = window.fhirpath.evaluate(fhirResource, 'AllergyIntolerance.criticality')
							CriticalityProcessed = CriticalityResult.reduce((accumulator, currentValue) => {
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
						
				
							JSON.stringify(CriticalityProcessed)
						 `)
	if err == nil && criticalityResult.String() != "undefined" {
		s.Criticality = []byte(criticalityResult.String())
	}
	// extracting Date
	dateResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'AllergyIntolerance.recordedDate | CarePlan.period | CareTeam.period | ClinicalImpression.date | Composition.date | Consent.dateTime | DiagnosticReport.effective | Encounter.period | EpisodeOfCare.period | FamilyMemberHistory.date | Flag.period | (Immunization.occurrenceDateTime) | List.date | Observation.effective | Procedure.performed | (RiskAssessment.occurrenceDateTime) | SupplyRequest.authoredOn')[0]")
	if err == nil && dateResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, dateResult.String())
		if err == nil {
			s.Date = t
		}
	}
	// extracting Identifier
	identifierResult, err := vm.RunString(` 
							IdentifierResult = window.fhirpath.evaluate(fhirResource, 'AllergyIntolerance.identifier | CarePlan.identifier | CareTeam.identifier | Composition.identifier | Condition.identifier | Consent.identifier | DetectedIssue.identifier | DeviceRequest.identifier | DiagnosticReport.identifier | DocumentManifest.masterIdentifier | DocumentManifest.identifier | DocumentReference.masterIdentifier | DocumentReference.identifier | Encounter.identifier | EpisodeOfCare.identifier | FamilyMemberHistory.identifier | Goal.identifier | ImagingStudy.identifier | Immunization.identifier | List.identifier | MedicationAdministration.identifier | MedicationDispense.identifier | MedicationRequest.identifier | MedicationStatement.identifier | NutritionOrder.identifier | Observation.identifier | Procedure.identifier | RiskAssessment.identifier | ServiceRequest.identifier | SupplyDelivery.identifier | SupplyRequest.identifier | VisionPrescription.identifier')
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
	// extracting LastDate
	lastDateResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'AllergyIntolerance.lastOccurrence')[0]")
	if err == nil && lastDateResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, lastDateResult.String())
		if err == nil {
			s.LastDate = t
		}
	}
	// extracting LastUpdated
	lastUpdatedResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Resource.meta.lastUpdated')[0]")
	if err == nil && lastUpdatedResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, lastUpdatedResult.String())
		if err == nil {
			s.LastUpdated = t
		}
	}
	// extracting Manifestation
	manifestationResult, err := vm.RunString(` 
							ManifestationResult = window.fhirpath.evaluate(fhirResource, 'AllergyIntolerance.reaction.manifestation')
							ManifestationProcessed = ManifestationResult.reduce((accumulator, currentValue) => {
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
						
				
							JSON.stringify(ManifestationProcessed)
						 `)
	if err == nil && manifestationResult.String() != "undefined" {
		s.Manifestation = []byte(manifestationResult.String())
	}
	// extracting Onset
	onsetResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'AllergyIntolerance.reaction.onset')[0]")
	if err == nil && onsetResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, onsetResult.String())
		if err == nil {
			s.Onset = t
		}
	}
	// extracting Profile
	profileResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Resource.meta.profile'))")
	if err == nil && profileResult.String() != "undefined" {
		s.Profile = []byte(profileResult.String())
	}
	// extracting Recorder
	recorderResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'AllergyIntolerance.recorder'))")
	if err == nil && recorderResult.String() != "undefined" {
		s.Recorder = []byte(recorderResult.String())
	}
	// extracting Route
	routeResult, err := vm.RunString(` 
							RouteResult = window.fhirpath.evaluate(fhirResource, 'AllergyIntolerance.reaction.exposureRoute')
							RouteProcessed = RouteResult.reduce((accumulator, currentValue) => {
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
						
				
							JSON.stringify(RouteProcessed)
						 `)
	if err == nil && routeResult.String() != "undefined" {
		s.Route = []byte(routeResult.String())
	}
	// extracting Severity
	severityResult, err := vm.RunString(` 
							SeverityResult = window.fhirpath.evaluate(fhirResource, 'AllergyIntolerance.reaction.severity')
							SeverityProcessed = SeverityResult.reduce((accumulator, currentValue) => {
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
						
				
							JSON.stringify(SeverityProcessed)
						 `)
	if err == nil && severityResult.String() != "undefined" {
		s.Severity = []byte(severityResult.String())
	}
	// extracting SourceUri
	sourceUriResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Resource.meta.source')[0]")
	if err == nil && sourceUriResult.String() != "undefined" {
		s.SourceUri = sourceUriResult.String()
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
	// extracting VerificationStatus
	verificationStatusResult, err := vm.RunString(` 
							VerificationStatusResult = window.fhirpath.evaluate(fhirResource, 'AllergyIntolerance.verificationStatus')
							VerificationStatusProcessed = VerificationStatusResult.reduce((accumulator, currentValue) => {
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
						
				
							JSON.stringify(VerificationStatusProcessed)
						 `)
	if err == nil && verificationStatusResult.String() != "undefined" {
		s.VerificationStatus = []byte(verificationStatusResult.String())
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirAllergyIntolerance) TableName() string {
	return "fhir_allergy_intolerance"
}
