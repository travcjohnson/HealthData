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

type FhirAllergyIntolerance struct {
	models.ResourceBase
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
	Date *time.Time `gorm:"column:date;type:datetime" json:"date,omitempty"`
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
	LastDate *time.Time `gorm:"column:lastDate;type:datetime" json:"lastDate,omitempty"`
	// Clinical symptoms/signs associated with the Event
	// https://hl7.org/fhir/r4/search.html#token
	Manifestation datatypes.JSON `gorm:"column:manifestation;type:text;serializer:json" json:"manifestation,omitempty"`
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
	// Date(/time) when manifestations showed
	// https://hl7.org/fhir/r4/search.html#date
	Onset *time.Time `gorm:"column:onset;type:datetime" json:"onset,omitempty"`
	// Who recorded the sensitivity
	// https://hl7.org/fhir/r4/search.html#reference
	Recorder datatypes.JSON `gorm:"column:recorder;type:text;serializer:json" json:"recorder,omitempty"`
	// How the subject was exposed to the substance
	// https://hl7.org/fhir/r4/search.html#token
	Route datatypes.JSON `gorm:"column:route;type:text;serializer:json" json:"route,omitempty"`
	// mild | moderate | severe (of event as a whole)
	// https://hl7.org/fhir/r4/search.html#token
	Severity datatypes.JSON `gorm:"column:severity;type:text;serializer:json" json:"severity,omitempty"`
	// Text search against the narrative
	// https://hl7.org/fhir/r4/search.html#string
	Text datatypes.JSON `gorm:"column:text;type:text;serializer:json" json:"text,omitempty"`
	/*
	   Multiple Resources:

	   * [AllergyIntolerance](allergyintolerance.html): allergy | intolerance - Underlying mechanism (if known)
	   * [Composition](composition.html): Kind of composition (LOINC if possible)
	   * [DocumentManifest](documentmanifest.html): Kind of document set
	   * [DocumentReference](documentreference.html): Kind of document (LOINC if possible)
	   * [Encounter](encounter.html): Specific type of encounter
	   * [EpisodeOfCare](episodeofcare.html): Type/class  - e.g. specialist referral, disease management
	*/
	// https://hl7.org/fhir/r4/search.html#token
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
	// unconfirmed | confirmed | refuted | entered-in-error
	// https://hl7.org/fhir/r4/search.html#token
	VerificationStatus datatypes.JSON `gorm:"column:verificationStatus;type:text;serializer:json" json:"verificationStatus,omitempty"`
}

func (s *FhirAllergyIntolerance) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"asserter":             "reference",
		"category":             "token",
		"clinicalStatus":       "token",
		"code":                 "token",
		"criticality":          "token",
		"date":                 "date",
		"id":                   "keyword",
		"identifier":           "token",
		"language":             "token",
		"lastDate":             "date",
		"manifestation":        "token",
		"metaLastUpdated":      "date",
		"metaProfile":          "reference",
		"metaTag":              "token",
		"metaVersionId":        "keyword",
		"onset":                "date",
		"recorder":             "reference",
		"route":                "token",
		"severity":             "token",
		"sort_date":            "date",
		"source_id":            "keyword",
		"source_resource_id":   "keyword",
		"source_resource_type": "keyword",
		"source_uri":           "keyword",
		"text":                 "string",
		"type":                 "token",
		"verificationStatus":   "token",
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
	// extracting Asserter
	asserterResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'AllergyIntolerance.asserter')")
	if err == nil && asserterResult.String() != "undefined" {
		s.Asserter = []byte(asserterResult.String())
	}
	// extracting Category
	categoryResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'AllergyIntolerance.category')")
	if err == nil && categoryResult.String() != "undefined" {
		s.Category = []byte(categoryResult.String())
	}
	// extracting ClinicalStatus
	clinicalStatusResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'AllergyIntolerance.clinicalStatus')")
	if err == nil && clinicalStatusResult.String() != "undefined" {
		s.ClinicalStatus = []byte(clinicalStatusResult.String())
	}
	// extracting Code
	codeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'AllergyIntolerance.code | AllergyIntolerance.reaction.substance | Condition.code | (DeviceRequest.codeCodeableConcept) | DiagnosticReport.code | FamilyMemberHistory.condition.code | List.code | Medication.code | (MedicationAdministration.medicationCodeableConcept) | (MedicationDispense.medicationCodeableConcept) | (MedicationRequest.medicationCodeableConcept) | (MedicationStatement.medicationCodeableConcept) | Observation.code | Procedure.code | ServiceRequest.code')")
	if err == nil && codeResult.String() != "undefined" {
		s.Code = []byte(codeResult.String())
	}
	// extracting Criticality
	criticalityResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'AllergyIntolerance.criticality')")
	if err == nil && criticalityResult.String() != "undefined" {
		s.Criticality = []byte(criticalityResult.String())
	}
	// extracting Date
	dateResult, err := vm.RunString("extractDateSearchParameters(fhirResource, 'AllergyIntolerance.recordedDate | CarePlan.period | CareTeam.period | ClinicalImpression.date | Composition.date | Consent.dateTime | DiagnosticReport.effectiveDateTime | DiagnosticReport.effectivePeriod | Encounter.period | EpisodeOfCare.period | FamilyMemberHistory.date | Flag.period | (Immunization.occurrenceDateTime) | List.date | Observation.effectiveDateTime | Observation.effectivePeriod | Observation.effectiveTiming | Observation.effectiveInstant | Procedure.performedDateTime | Procedure.performedPeriod | Procedure.performedString | Procedure.performedAge | Procedure.performedRange | (RiskAssessment.occurrenceDateTime) | SupplyRequest.authoredOn')")
	if err == nil && dateResult.String() != "undefined" {
		if t, err := time.Parse(time.RFC3339, dateResult.String()); err == nil {
			s.Date = &t
		} else if t, err = time.Parse("2006-01-02", dateResult.String()); err == nil {
			s.Date = &t
		} else if t, err = time.Parse("2006-01", dateResult.String()); err == nil {
			s.Date = &t
		} else if t, err = time.Parse("2006", dateResult.String()); err == nil {
			s.Date = &t
		}
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
	// extracting LastDate
	lastDateResult, err := vm.RunString("extractDateSearchParameters(fhirResource, 'AllergyIntolerance.lastOccurrence')")
	if err == nil && lastDateResult.String() != "undefined" {
		if t, err := time.Parse(time.RFC3339, lastDateResult.String()); err == nil {
			s.LastDate = &t
		} else if t, err = time.Parse("2006-01-02", lastDateResult.String()); err == nil {
			s.LastDate = &t
		} else if t, err = time.Parse("2006-01", lastDateResult.String()); err == nil {
			s.LastDate = &t
		} else if t, err = time.Parse("2006", lastDateResult.String()); err == nil {
			s.LastDate = &t
		}
	}
	// extracting Manifestation
	manifestationResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'AllergyIntolerance.reaction.manifestation')")
	if err == nil && manifestationResult.String() != "undefined" {
		s.Manifestation = []byte(manifestationResult.String())
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
	// extracting Onset
	onsetResult, err := vm.RunString("extractDateSearchParameters(fhirResource, 'AllergyIntolerance.reaction.onset')")
	if err == nil && onsetResult.String() != "undefined" {
		if t, err := time.Parse(time.RFC3339, onsetResult.String()); err == nil {
			s.Onset = &t
		} else if t, err = time.Parse("2006-01-02", onsetResult.String()); err == nil {
			s.Onset = &t
		} else if t, err = time.Parse("2006-01", onsetResult.String()); err == nil {
			s.Onset = &t
		} else if t, err = time.Parse("2006", onsetResult.String()); err == nil {
			s.Onset = &t
		}
	}
	// extracting Recorder
	recorderResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'AllergyIntolerance.recorder')")
	if err == nil && recorderResult.String() != "undefined" {
		s.Recorder = []byte(recorderResult.String())
	}
	// extracting Route
	routeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'AllergyIntolerance.reaction.exposureRoute')")
	if err == nil && routeResult.String() != "undefined" {
		s.Route = []byte(routeResult.String())
	}
	// extracting Severity
	severityResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'AllergyIntolerance.reaction.severity')")
	if err == nil && severityResult.String() != "undefined" {
		s.Severity = []byte(severityResult.String())
	}
	// extracting Text
	textResult, err := vm.RunString("extractStringSearchParameters(fhirResource, 'text')")
	if err == nil && textResult.String() != "undefined" {
		s.Text = []byte(textResult.String())
	}
	// extracting Type
	typeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'AllergyIntolerance.type | Composition.type | DocumentManifest.type | DocumentReference.type | Encounter.type | EpisodeOfCare.type')")
	if err == nil && typeResult.String() != "undefined" {
		s.Type = []byte(typeResult.String())
	}
	// extracting VerificationStatus
	verificationStatusResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'AllergyIntolerance.verificationStatus')")
	if err == nil && verificationStatusResult.String() != "undefined" {
		s.VerificationStatus = []byte(verificationStatusResult.String())
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirAllergyIntolerance) TableName() string {
	return "fhir_allergy_intolerance"
}
