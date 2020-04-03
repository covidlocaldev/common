package resources

import (
	"github.com/covidlocaldev/common/emr/codes"
	"time"
)

type DiagnosticReport struct {
	// http://hl7.org/implement/standards/emr/STU3/diagnosticreport.html
}

type Observation struct {
	// http://hl7.org/implement/standards/emr/STU3/observation.html
	BaseResource
	BasedOn           []Reference             `firestore:"basedOn,omitempty" json:"basedOn,omitempty"`
	Status            codes.ObservationStatus `firestore:"status,omitempty" json:"status"`
	Category          []CodeableConcept       `firestore:"category,omitempty" json:"category,omitempty"`
	Code              CodeableConcept         `firestore:"code,omitempty" json:"code"`
	Subject           *Reference              `firestore:"subject,omitempty" json:"subject,omitempty"`
	Context           *Reference              `firestore:"context,omitempty" json:"context,omitempty"`
	EffectiveDateTime *time.Time              `firestore:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod   *Period                 `firestore:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Issued            *time.Time              `firestore:"issued,omitempty" json:"issued,omitempty"`
	Performer         []Reference             `firestore:"performer,omitempty" json:"performer,omitempty"`
	ValueQuantity     *Quantity               `firestore:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueString       string                  `firestore:"valueString,omitempty" json:"valueString,omitempty"`
	ValueDateTime     *time.Time              `firestore:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	Comment           string                  `firestore:"comment,omitempty" json:"comment,omitempty"`
	BodySite          *CodeableConcept        `firestore:"bodySite,omitempty" json:"bodySite,omitempty"`
	Method            *CodeableConcept        `firestore:"method,omitempty" json:"method,omitempty"`
	Specimen          *Reference              `firestore:"specimen,omitempty" json:"specimen,omitempty"`
}

type BodySite struct {
	Identifier  []Identifier                          `firestore:"identifier,omitempty" json:"identifier"`
	Active      bool                                  `firestore:"active,omitempty" json:"active"`
	Code        codes.BodyStructure                   `firestore:"code,omitempty" json:"code"`
	Qualifier   []codes.BodySiteLocationQualifierCode `firestore:"qualifier,omitempty" json:"qualifier"`
	Description string                                `firestore:"description,omitempty" json:"description"`
	Image       []Attachment                          `firestore:"image,omitempty" json:"image"`
	Patient     *Patient                              `firestore:"patient,omitempty" json:"patient"`
}
