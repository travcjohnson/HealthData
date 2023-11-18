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

type FhirEncounter struct {
	models.ResourceBase
	// The set of accounts that may be used for billing for this Encounter
	// https://hl7.org/fhir/r4/search.html#reference
	Account datatypes.JSON `gorm:"column:account;type:text;serializer:json" json:"account,omitempty"`
	// The appointment that scheduled this encounter
	// https://hl7.org/fhir/r4/search.html#reference
	Appointment datatypes.JSON `gorm:"column:appointment;type:text;serializer:json" json:"appointment,omitempty"`
	// The ServiceRequest that initiated this encounter
	// https://hl7.org/fhir/r4/search.html#reference
	BasedOn datatypes.JSON `gorm:"column:basedOn;type:text;serializer:json" json:"basedOn,omitempty"`
	// Classification of patient encounter
	// https://hl7.org/fhir/r4/search.html#token
	Class datatypes.JSON `gorm:"column:class;type:text;serializer:json" json:"class,omitempty"`
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
	Date *time.Time `gorm:"column:date;type:datetime" json:"date,omitempty"`
	// The diagnosis or procedure relevant to the encounter
	// https://hl7.org/fhir/r4/search.html#reference
	Diagnosis datatypes.JSON `gorm:"column:diagnosis;type:text;serializer:json" json:"diagnosis,omitempty"`
	// Episode(s) of care that this encounter should be recorded against
	// https://hl7.org/fhir/r4/search.html#reference
	EpisodeOfCare datatypes.JSON `gorm:"column:episodeOfCare;type:text;serializer:json" json:"episodeOfCare,omitempty"`
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
	// Length of encounter in days
	// https://hl7.org/fhir/r4/search.html#quantity
	Length datatypes.JSON `gorm:"column:length;type:text;serializer:json" json:"length,omitempty"`
	// Location the encounter takes place
	// https://hl7.org/fhir/r4/search.html#reference
	Location datatypes.JSON `gorm:"column:location;type:text;serializer:json" json:"location,omitempty"`
	// Time period during which the patient was present at the location
	// https://hl7.org/fhir/r4/search.html#date
	LocationPeriod *time.Time `gorm:"column:locationPeriod;type:datetime" json:"locationPeriod,omitempty"`
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
	// https://hl7.org/fhir/r4/search.html#keyword
	MetaVersionId string `gorm:"column:metaVersionId;type:text" json:"metaVersionId,omitempty"`
	// Another Encounter this encounter is part of
	// https://hl7.org/fhir/r4/search.html#reference
	PartOf datatypes.JSON `gorm:"column:partOf;type:text;serializer:json" json:"partOf,omitempty"`
	// Persons involved in the encounter other than the patient
	// https://hl7.org/fhir/r4/search.html#reference
	Participant datatypes.JSON `gorm:"column:participant;type:text;serializer:json" json:"participant,omitempty"`
	// Role of participant in encounter
	// https://hl7.org/fhir/r4/search.html#token
	ParticipantType datatypes.JSON `gorm:"column:participantType;type:text;serializer:json" json:"participantType,omitempty"`
	// Persons involved in the encounter other than the patient
	// https://hl7.org/fhir/r4/search.html#reference
	Practitioner datatypes.JSON `gorm:"column:practitioner;type:text;serializer:json" json:"practitioner,omitempty"`
	// Coded reason the encounter takes place
	// https://hl7.org/fhir/r4/search.html#token
	ReasonCode datatypes.JSON `gorm:"column:reasonCode;type:text;serializer:json" json:"reasonCode,omitempty"`
	// Reason the encounter takes place (reference)
	// https://hl7.org/fhir/r4/search.html#reference
	ReasonReference datatypes.JSON `gorm:"column:reasonReference;type:text;serializer:json" json:"reasonReference,omitempty"`
	// The organization (facility) responsible for this encounter
	// https://hl7.org/fhir/r4/search.html#reference
	ServiceProvider datatypes.JSON `gorm:"column:serviceProvider;type:text;serializer:json" json:"serviceProvider,omitempty"`
	// Wheelchair, translator, stretcher, etc.
	// https://hl7.org/fhir/r4/search.html#token
	SpecialArrangement datatypes.JSON `gorm:"column:specialArrangement;type:text;serializer:json" json:"specialArrangement,omitempty"`
	// planned | arrived | triaged | in-progress | onleave | finished | cancelled +
	// https://hl7.org/fhir/r4/search.html#token
	Status datatypes.JSON `gorm:"column:status;type:text;serializer:json" json:"status,omitempty"`
	// The patient or group present at the encounter
	// https://hl7.org/fhir/r4/search.html#reference
	Subject datatypes.JSON `gorm:"column:subject;type:text;serializer:json" json:"subject,omitempty"`
	// Text search against the narrative
	// https://hl7.org/fhir/r4/search.html#string
	Text datatypes.JSON `gorm:"column:text;type:text;serializer:json" json:"text,omitempty"`
	// A resource type filter
	// https://hl7.org/fhir/r4/search.html#special
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
}

func (s *FhirEncounter) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"account":              "reference",
		"appointment":          "reference",
		"basedOn":              "reference",
		"class":                "token",
		"date":                 "date",
		"diagnosis":            "reference",
		"episodeOfCare":        "reference",
		"id":                   "keyword",
		"identifier":           "token",
		"language":             "token",
		"length":               "quantity",
		"location":             "reference",
		"locationPeriod":       "date",
		"metaLastUpdated":      "date",
		"metaProfile":          "reference",
		"metaTag":              "token",
		"metaVersionId":        "keyword",
		"partOf":               "reference",
		"participant":          "reference",
		"participantType":      "token",
		"practitioner":         "reference",
		"reasonCode":           "token",
		"reasonReference":      "reference",
		"serviceProvider":      "reference",
		"sort_date":            "date",
		"source_id":            "keyword",
		"source_resource_id":   "keyword",
		"source_resource_type": "keyword",
		"source_uri":           "keyword",
		"specialArrangement":   "token",
		"status":               "token",
		"subject":              "reference",
		"text":                 "string",
		"type":                 "special",
	}
	return searchParameters
}
func (s *FhirEncounter) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
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
	// extracting Account
	accountResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Encounter.account')")
	if err == nil && accountResult.String() != "undefined" {
		s.Account = []byte(accountResult.String())
	}
	// extracting Appointment
	appointmentResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Encounter.appointment')")
	if err == nil && appointmentResult.String() != "undefined" {
		s.Appointment = []byte(appointmentResult.String())
	}
	// extracting BasedOn
	basedOnResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Encounter.basedOn')")
	if err == nil && basedOnResult.String() != "undefined" {
		s.BasedOn = []byte(basedOnResult.String())
	}
	// extracting Class
	classResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Encounter.class')")
	if err == nil && classResult.String() != "undefined" {
		s.Class = []byte(classResult.String())
	}
	// extracting Date
	dateResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'AllergyIntolerance.recordedDate | CarePlan.period | CareTeam.period | ClinicalImpression.date | Composition.date | Consent.dateTime | DiagnosticReport.effectiveDateTime | DiagnosticReport.effectivePeriod | Encounter.period | EpisodeOfCare.period | FamilyMemberHistory.date | Flag.period | (Immunization.occurrenceDateTime) | List.date | Observation.effectiveDateTime | Observation.effectivePeriod | Observation.effectiveTiming | Observation.effectiveInstant | Procedure.performedDateTime | Procedure.performedPeriod | Procedure.performedString | Procedure.performedAge | Procedure.performedRange | (RiskAssessment.occurrenceDateTime) | SupplyRequest.authoredOn')")
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
	// extracting Diagnosis
	diagnosisResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Encounter.diagnosis.condition')")
	if err == nil && diagnosisResult.String() != "undefined" {
		s.Diagnosis = []byte(diagnosisResult.String())
	}
	// extracting EpisodeOfCare
	episodeOfCareResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Encounter.episodeOfCare')")
	if err == nil && episodeOfCareResult.String() != "undefined" {
		s.EpisodeOfCare = []byte(episodeOfCareResult.String())
	}
	// extracting Identifier
	identifierResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'AllergyIntolerance.identifier | CarePlan.identifier | CareTeam.identifier | Composition.identifier | Condition.identifier | Consent.identifier | DetectedIssue.identifier | DeviceRequest.identifier | DiagnosticReport.identifier | DocumentManifest.masterIdentifier | DocumentManifest.identifier | DocumentReference.masterIdentifier | DocumentReference.identifier | Encounter.identifier | EpisodeOfCare.identifier | FamilyMemberHistory.identifier | Goal.identifier | ImagingStudy.identifier | Immunization.identifier | List.identifier | MedicationAdministration.identifier | MedicationDispense.identifier | MedicationRequest.identifier | MedicationStatement.identifier | NutritionOrder.identifier | Observation.identifier | Procedure.identifier | RiskAssessment.identifier | ServiceRequest.identifier | SupplyDelivery.identifier | SupplyRequest.identifier | VisionPrescription.identifier')")
	if err == nil && identifierResult.String() != "undefined" {
		s.Identifier = []byte(identifierResult.String())
	}
	// extracting Language
	languageResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'language')")
	if err == nil && languageResult.String() != "undefined" {
		s.Language = []byte(languageResult.String())
	}
	// extracting Length
	lengthResult, err := vm.RunString("extractCatchallSearchParameters(fhirResource, 'Encounter.length')")
	if err == nil && lengthResult.String() != "undefined" {
		s.Length = []byte(lengthResult.String())
	}
	// extracting Location
	locationResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Encounter.location.location')")
	if err == nil && locationResult.String() != "undefined" {
		s.Location = []byte(locationResult.String())
	}
	// extracting LocationPeriod
	locationPeriodResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'Encounter.location.period')")
	if err == nil && locationPeriodResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, locationPeriodResult.String())
		if err == nil {
			s.LocationPeriod = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", locationPeriodResult.String())
			if err == nil {
				s.LocationPeriod = &d
			}
		}
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
	partOfResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Encounter.partOf')")
	if err == nil && partOfResult.String() != "undefined" {
		s.PartOf = []byte(partOfResult.String())
	}
	// extracting Participant
	participantResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Encounter.participant.individual')")
	if err == nil && participantResult.String() != "undefined" {
		s.Participant = []byte(participantResult.String())
	}
	// extracting ParticipantType
	participantTypeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Encounter.participant.type')")
	if err == nil && participantTypeResult.String() != "undefined" {
		s.ParticipantType = []byte(participantTypeResult.String())
	}
	// extracting Practitioner
	practitionerResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Encounter.participant.individual.where(resolve() is Practitioner)')")
	if err == nil && practitionerResult.String() != "undefined" {
		s.Practitioner = []byte(practitionerResult.String())
	}
	// extracting ReasonCode
	reasonCodeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Encounter.reasonCode')")
	if err == nil && reasonCodeResult.String() != "undefined" {
		s.ReasonCode = []byte(reasonCodeResult.String())
	}
	// extracting ReasonReference
	reasonReferenceResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Encounter.reasonReference')")
	if err == nil && reasonReferenceResult.String() != "undefined" {
		s.ReasonReference = []byte(reasonReferenceResult.String())
	}
	// extracting ServiceProvider
	serviceProviderResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Encounter.serviceProvider')")
	if err == nil && serviceProviderResult.String() != "undefined" {
		s.ServiceProvider = []byte(serviceProviderResult.String())
	}
	// extracting SpecialArrangement
	specialArrangementResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Encounter.hospitalization.specialArrangement')")
	if err == nil && specialArrangementResult.String() != "undefined" {
		s.SpecialArrangement = []byte(specialArrangementResult.String())
	}
	// extracting Status
	statusResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'Encounter.status')")
	if err == nil && statusResult.String() != "undefined" {
		s.Status = []byte(statusResult.String())
	}
	// extracting Subject
	subjectResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'Encounter.subject')")
	if err == nil && subjectResult.String() != "undefined" {
		s.Subject = []byte(subjectResult.String())
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirEncounter) TableName() string {
	return "fhir_encounter"
}
