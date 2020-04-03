package resources

import "github.com/covidlocaldev/common/emr/codes"

type CarePlan struct {
	// http://hl7.org/implement/standards/emr/STU3/careplan.html
	BaseResource
	Definition  []Reference          `firestore:"definition,omitempty" json:"definition,omitempty"`
	BasedOn     []Reference          `firestore:"basedOn,omitempty" json:"basedOn,omitempty"`
	Replaces    []Reference          `firestore:"replaces,omitempty" json:"replaces,omitempty"`
	PartOf      []Reference          `firestore:"partOf,omitempty" json:"partOf,omitempty"`
	Status      codes.CarePlanStatus `firestore:"status,omitempty" json:"status"`
	Intent      codes.CarePlanIntent `firestore:"intent,omitempty" json:"intent"`
	Category    []CodeableConcept    `firestore:"category,omitempty" json:"category,omitempty"`
	Title       string               `firestore:"title,omitempty" json:"title,omitempty"`
	Description string               `firestore:"description,omitempty" json:"description,omitempty"`
	Subject     Reference            `firestore:"subject,omitempty" json:"subject"`
	Context     *Reference           `firestore:"context,omitempty" json:"context,omitempty"`
	Period      *Period              `firestore:"period,omitempty" json:"period,omitempty"`
	Author      []Reference          `firestore:"author,omitempty" json:"author,omitempty"`
	CareTeam    []Reference          `firestore:"careTeam,omitempty" json:"careTeam,omitempty"`
	Addresses   []Reference          `firestore:"addresses,omitempty" json:"addresses,omitempty"`
	Goal        []Reference          `firestore:"goal,omitempty" json:"goal,omitempty"`
	Activity    []CarePlanActivity   `firestore:"activity,omitempty" json:"activity"`
}

type ProcedureRequest struct {
	// http://hl7.org/implement/standards/emr/STU3/procedurerequest.html
}

type ReferralRequest struct {
	// http://hl7.org/implement/standards/emr/STU3/referralrequest.html
}
