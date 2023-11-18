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

type FhirMedicationRequest struct {
	models.ResourceBase
	// Return prescriptions written on this date
	// https://hl7.org/fhir/r4/search.html#date
	Authoredon *time.Time `gorm:"column:authoredon;type:datetime" json:"authoredon,omitempty"`
	// Returns prescriptions with different categories
	// https://hl7.org/fhir/r4/search.html#token
	Category datatypes.JSON `gorm:"column:category;type:text;serializer:json" json:"category,omitempty"`
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
	/*
	   Multiple Resources:

	   * [MedicationRequest](medicationrequest.html): Returns medication request to be administered on a specific date
	*/
	// https://hl7.org/fhir/r4/search.html#date
	Date *time.Time `gorm:"column:date;type:datetime" json:"date,omitempty"`
	/*
	   Multiple Resources:

	   * [MedicationRequest](medicationrequest.html): Return prescriptions with this encounter identifier
	*/
	// https://hl7.org/fhir/r4/search.html#reference
	Encounter datatypes.JSON `gorm:"column:encounter;type:text;serializer:json" json:"encounter,omitempty"`
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
	// Returns prescriptions intended to be dispensed by this Organization
	// https://hl7.org/fhir/r4/search.html#reference
	IntendedDispenser datatypes.JSON `gorm:"column:intendedDispenser;type:text;serializer:json" json:"intendedDispenser,omitempty"`
	// Returns the intended performer of the administration of the medication request
	// https://hl7.org/fhir/r4/search.html#reference
	IntendedPerformer datatypes.JSON `gorm:"column:intendedPerformer;type:text;serializer:json" json:"intendedPerformer,omitempty"`
	// Returns requests for a specific type of performer
	// https://hl7.org/fhir/r4/search.html#token
	IntendedPerformertype datatypes.JSON `gorm:"column:intendedPerformertype;type:text;serializer:json" json:"intendedPerformertype,omitempty"`
	// Returns prescriptions with different intents
	// https://hl7.org/fhir/r4/search.html#token
	Intent datatypes.JSON `gorm:"column:intent;type:text;serializer:json" json:"intent,omitempty"`
	// Language of the resource content
	// https://hl7.org/fhir/r4/search.html#token
	Language datatypes.JSON `gorm:"column:language;type:text;serializer:json" json:"language,omitempty"`
	/*
	   Multiple Resources:

	   * [MedicationAdministration](medicationadministration.html): Return administrations of this medication resource
	   * [MedicationDispense](medicationdispense.html): Returns dispenses of this medicine resource
	   * [MedicationRequest](medicationrequest.html): Return prescriptions for this medication reference
	   * [MedicationStatement](medicationstatement.html): Return statements of this medication reference
	*/
	// https://hl7.org/fhir/r4/search.html#reference
	Medication datatypes.JSON `gorm:"column:medication;type:text;serializer:json" json:"medication,omitempty"`
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
	// Returns prescriptions with different priorities
	// https://hl7.org/fhir/r4/search.html#token
	Priority datatypes.JSON `gorm:"column:priority;type:text;serializer:json" json:"priority,omitempty"`
	// Returns prescriptions prescribed by this prescriber
	// https://hl7.org/fhir/r4/search.html#reference
	Requester datatypes.JSON `gorm:"column:requester;type:text;serializer:json" json:"requester,omitempty"`
	/*
	   Multiple Resources:

	   * [MedicationAdministration](medicationadministration.html): MedicationAdministration event status (for example one of active/paused/completed/nullified)
	   * [MedicationDispense](medicationdispense.html): Returns dispenses with a specified dispense status
	   * [MedicationRequest](medicationrequest.html): Status of the prescription
	   * [MedicationStatement](medicationstatement.html): Return statements that match the given status
	*/
	// https://hl7.org/fhir/r4/search.html#token
	Status datatypes.JSON `gorm:"column:status;type:text;serializer:json" json:"status,omitempty"`
	// The identity of a patient to list orders  for
	// https://hl7.org/fhir/r4/search.html#reference
	Subject datatypes.JSON `gorm:"column:subject;type:text;serializer:json" json:"subject,omitempty"`
	// Text search against the narrative
	// This is a primitive string literal (`keyword` type). It is not a recognized SearchParameter type from https://hl7.org/fhir/r4/search.html, it's Fasten Health-specific
	Text string `gorm:"column:text;type:text" json:"text,omitempty"`
	// A resource type filter
	// https://hl7.org/fhir/r4/search.html#special
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
}

func (s *FhirMedicationRequest) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"authoredon":            "date",
		"category":              "token",
		"code":                  "token",
		"date":                  "date",
		"encounter":             "reference",
		"id":                    "keyword",
		"identifier":            "token",
		"intendedDispenser":     "reference",
		"intendedPerformer":     "reference",
		"intendedPerformertype": "token",
		"intent":                "token",
		"language":              "token",
		"medication":            "reference",
		"metaLastUpdated":       "date",
		"metaProfile":           "reference",
		"metaTag":               "token",
		"metaVersionId":         "keyword",
		"priority":              "token",
		"requester":             "reference",
		"sort_date":             "date",
		"source_id":             "keyword",
		"source_resource_id":    "keyword",
		"source_resource_type":  "keyword",
		"source_uri":            "keyword",
		"status":                "token",
		"subject":               "reference",
		"text":                  "keyword",
		"type":                  "special",
	}
	return searchParameters
}
func (s *FhirMedicationRequest) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
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
	// extracting Authoredon
	authoredonResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'MedicationRequest.authoredOn')")
	if err == nil && authoredonResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, authoredonResult.String())
		if err == nil {
			s.Authoredon = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", authoredonResult.String())
			if err == nil {
				s.Authoredon = &d
			}
		}
	}
	// extracting Category
	categoryResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'MedicationRequest.category')")
	if err == nil && categoryResult.String() != "undefined" {
		s.Category = []byte(categoryResult.String())
	}
	// extracting Code
	codeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'AllergyIntolerance.code | AllergyIntolerance.reaction.substance | Condition.code | (DeviceRequest.codeCodeableConcept) | DiagnosticReport.code | FamilyMemberHistory.condition.code | List.code | Medication.code | (MedicationAdministration.medicationCodeableConcept) | (MedicationDispense.medicationCodeableConcept) | (MedicationRequest.medicationCodeableConcept) | (MedicationStatement.medicationCodeableConcept) | Observation.code | Procedure.code | ServiceRequest.code')")
	if err == nil && codeResult.String() != "undefined" {
		s.Code = []byte(codeResult.String())
	}
	// extracting Date
	dateResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'MedicationRequest.dosageInstruction.timing.event')")
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
	// extracting Encounter
	encounterResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'MedicationRequest.encounter')")
	if err == nil && encounterResult.String() != "undefined" {
		s.Encounter = []byte(encounterResult.String())
	}
	// extracting Identifier
	identifierResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'AllergyIntolerance.identifier | CarePlan.identifier | CareTeam.identifier | Composition.identifier | Condition.identifier | Consent.identifier | DetectedIssue.identifier | DeviceRequest.identifier | DiagnosticReport.identifier | DocumentManifest.masterIdentifier | DocumentManifest.identifier | DocumentReference.masterIdentifier | DocumentReference.identifier | Encounter.identifier | EpisodeOfCare.identifier | FamilyMemberHistory.identifier | Goal.identifier | ImagingStudy.identifier | Immunization.identifier | List.identifier | MedicationAdministration.identifier | MedicationDispense.identifier | MedicationRequest.identifier | MedicationStatement.identifier | NutritionOrder.identifier | Observation.identifier | Procedure.identifier | RiskAssessment.identifier | ServiceRequest.identifier | SupplyDelivery.identifier | SupplyRequest.identifier | VisionPrescription.identifier')")
	if err == nil && identifierResult.String() != "undefined" {
		s.Identifier = []byte(identifierResult.String())
	}
	// extracting IntendedDispenser
	intendedDispenserResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'MedicationRequest.dispenseRequest.performer')")
	if err == nil && intendedDispenserResult.String() != "undefined" {
		s.IntendedDispenser = []byte(intendedDispenserResult.String())
	}
	// extracting IntendedPerformer
	intendedPerformerResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'MedicationRequest.performer')")
	if err == nil && intendedPerformerResult.String() != "undefined" {
		s.IntendedPerformer = []byte(intendedPerformerResult.String())
	}
	// extracting IntendedPerformertype
	intendedPerformertypeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'MedicationRequest.performerType')")
	if err == nil && intendedPerformertypeResult.String() != "undefined" {
		s.IntendedPerformertype = []byte(intendedPerformertypeResult.String())
	}
	// extracting Intent
	intentResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'MedicationRequest.intent')")
	if err == nil && intentResult.String() != "undefined" {
		s.Intent = []byte(intentResult.String())
	}
	// extracting Language
	languageResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'language')")
	if err == nil && languageResult.String() != "undefined" {
		s.Language = []byte(languageResult.String())
	}
	// extracting Medication
	medicationResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, '(MedicationAdministration.medicationReference) | (MedicationDispense.medicationReference) | (MedicationRequest.medicationReference) | (MedicationStatement.medicationReference)')")
	if err == nil && medicationResult.String() != "undefined" {
		s.Medication = []byte(medicationResult.String())
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
	// extracting Priority
	priorityResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'MedicationRequest.priority')")
	if err == nil && priorityResult.String() != "undefined" {
		s.Priority = []byte(priorityResult.String())
	}
	// extracting Requester
	requesterResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'MedicationRequest.requester')")
	if err == nil && requesterResult.String() != "undefined" {
		s.Requester = []byte(requesterResult.String())
	}
	// extracting Status
	statusResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'MedicationAdministration.status | MedicationDispense.status | MedicationRequest.status | MedicationStatement.status')")
	if err == nil && statusResult.String() != "undefined" {
		s.Status = []byte(statusResult.String())
	}
	// extracting Subject
	subjectResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'MedicationRequest.subject')")
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
func (s *FhirMedicationRequest) TableName() string {
	return "fhir_medication_request"
}
