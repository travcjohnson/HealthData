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

type FhirSpecimen struct {
	models.OriginBase
	// The accession number associated with the specimen
	// https://hl7.org/fhir/r4/search.html#token
	Accession datatypes.JSON `gorm:"column:accession;type:text;serializer:json" json:"accession,omitempty"`
	// The code for the body site from where the specimen originated
	// https://hl7.org/fhir/r4/search.html#token
	Bodysite datatypes.JSON `gorm:"column:bodysite;type:text;serializer:json" json:"bodysite,omitempty"`
	// The date the specimen was collected
	// https://hl7.org/fhir/r4/search.html#date
	Collected time.Time `gorm:"column:collected;type:datetime" json:"collected,omitempty"`
	// Who collected the specimen
	// https://hl7.org/fhir/r4/search.html#reference
	Collector datatypes.JSON `gorm:"column:collector;type:text;serializer:json" json:"collector,omitempty"`
	// The kind of specimen container
	// https://hl7.org/fhir/r4/search.html#token
	Container datatypes.JSON `gorm:"column:container;type:text;serializer:json" json:"container,omitempty"`
	// The unique identifier associated with the specimen container
	// https://hl7.org/fhir/r4/search.html#token
	ContainerId datatypes.JSON `gorm:"column:containerId;type:text;serializer:json" json:"containerId,omitempty"`
	// The unique identifier associated with the specimen
	// https://hl7.org/fhir/r4/search.html#token
	Identifier datatypes.JSON `gorm:"column:identifier;type:text;serializer:json" json:"identifier,omitempty"`
	// Language of the resource content
	// https://hl7.org/fhir/r4/search.html#token
	Language datatypes.JSON `gorm:"column:language;type:text;serializer:json" json:"language,omitempty"`
	// When the resource version last changed
	// https://hl7.org/fhir/r4/search.html#date
	LastUpdated time.Time `gorm:"column:lastUpdated;type:datetime" json:"lastUpdated,omitempty"`
	// The parent of the specimen
	// https://hl7.org/fhir/r4/search.html#reference
	Parent datatypes.JSON `gorm:"column:parent;type:text;serializer:json" json:"parent,omitempty"`
	// Profiles this resource claims to conform to
	// https://hl7.org/fhir/r4/search.html#reference
	Profile datatypes.JSON `gorm:"column:profile;type:text;serializer:json" json:"profile,omitempty"`
	// The raw resource content in JSON format
	// https://hl7.org/fhir/r4/search.html#special
	RawResource datatypes.JSON `gorm:"column:rawResource;type:text;serializer:json" json:"rawResource,omitempty"`
	// Identifies where the resource comes from
	// https://hl7.org/fhir/r4/search.html#uri
	SourceUri string `gorm:"column:sourceUri;type:text" json:"sourceUri,omitempty"`
	// available | unavailable | unsatisfactory | entered-in-error
	// https://hl7.org/fhir/r4/search.html#token
	Status datatypes.JSON `gorm:"column:status;type:text;serializer:json" json:"status,omitempty"`
	// The subject of the specimen
	// https://hl7.org/fhir/r4/search.html#reference
	Subject datatypes.JSON `gorm:"column:subject;type:text;serializer:json" json:"subject,omitempty"`
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

func (s *FhirSpecimen) SetOriginBase(originBase models.OriginBase) {
	s.OriginBase = originBase
}
func (s *FhirSpecimen) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"accession":   "token",
		"bodysite":    "token",
		"collected":   "date",
		"collector":   "reference",
		"container":   "token",
		"containerId": "token",
		"identifier":  "token",
		"language":    "token",
		"lastUpdated": "date",
		"parent":      "reference",
		"profile":     "reference",
		"rawResource": "special",
		"sourceUri":   "uri",
		"status":      "token",
		"subject":     "reference",
		"tag":         "token",
		"text":        "string",
		"type":        "special",
	}
	return searchParameters
}
func (s *FhirSpecimen) PopulateAndExtractSearchParameters(rawResource json.RawMessage) error {
	s.RawResource = datatypes.JSON(rawResource)
	// unmarshal the raw resource (bytes) into a map
	var rawResourceMap map[string]interface{}
	err := json.Unmarshal(rawResource, &rawResourceMap)
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
	vm.Set("fhirResource", rawResourceMap)
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
	// extracting Container
	containerResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Specimen.container.type'))")
	if err == nil && containerResult.String() != "undefined" {
		s.Container = []byte(containerResult.String())
	}
	// extracting Status
	statusResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Specimen.status'))")
	if err == nil && statusResult.String() != "undefined" {
		s.Status = []byte(statusResult.String())
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
	// extracting Collected
	collectedResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Specimen.collection.collected')[0]")
	if err == nil && collectedResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, collectedResult.String())
		if err == nil {
			s.Collected = t
		}
	}
	// extracting Parent
	parentResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Specimen.parent'))")
	if err == nil && parentResult.String() != "undefined" {
		s.Parent = []byte(parentResult.String())
	}
	// extracting LastUpdated
	lastUpdatedResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Resource.meta.lastUpdated')[0]")
	if err == nil && lastUpdatedResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, lastUpdatedResult.String())
		if err == nil {
			s.LastUpdated = t
		}
	}
	// extracting Subject
	subjectResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Specimen.subject'))")
	if err == nil && subjectResult.String() != "undefined" {
		s.Subject = []byte(subjectResult.String())
	}
	// extracting Language
	languageResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Resource.language'))")
	if err == nil && languageResult.String() != "undefined" {
		s.Language = []byte(languageResult.String())
	}
	// extracting Tag
	tagResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Resource.meta.tag'))")
	if err == nil && tagResult.String() != "undefined" {
		s.Tag = []byte(tagResult.String())
	}
	// extracting Bodysite
	bodysiteResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Specimen.collection.bodySite'))")
	if err == nil && bodysiteResult.String() != "undefined" {
		s.Bodysite = []byte(bodysiteResult.String())
	}
	// extracting Collector
	collectorResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Specimen.collection.collector'))")
	if err == nil && collectorResult.String() != "undefined" {
		s.Collector = []byte(collectorResult.String())
	}
	// extracting Identifier
	identifierResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Specimen.identifier'))")
	if err == nil && identifierResult.String() != "undefined" {
		s.Identifier = []byte(identifierResult.String())
	}
	// extracting Accession
	accessionResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Specimen.accessionIdentifier'))")
	if err == nil && accessionResult.String() != "undefined" {
		s.Accession = []byte(accessionResult.String())
	}
	// extracting ContainerId
	containerIdResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Specimen.container.identifier'))")
	if err == nil && containerIdResult.String() != "undefined" {
		s.ContainerId = []byte(containerIdResult.String())
	}
	return nil
}
