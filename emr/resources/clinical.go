package resources

import (
	"github.com/covidlocaldev/common/emr/codes"
	"time"
)

type AllergyIntolerance struct {
	Identifier         []Identifier                               `firestore:"identifier,omitempty" json:"identifier"`
	ClinicalStatus     codes.AllergyIntoleranceClinicalStatus     `firestore:"clinicalStatus,omitempty" json:"clinicalStatus"`
	VerificationStatus codes.AllergyIntoleranceVerificationStatus `firestore:"verificationStatus,omitempty" json:"verificationStatus"`
	Type               codes.AllergyIntoleranceType               `firestore:"type,omitempty" json:"type"`
	Category           []codes.AllergyIntoleranceCategory         `firestore:"category,omitempty" json:"category"`
	Criticality        codes.AllergyIntoleranceCriticality        `firestore:"criticality,omitempty" json:"criticality"`
	Code               codes.AllergyIntoleranceCode               `firestore:"code,omitempty" json:"code"`
	Patient            Patient                                    `firestore:"patient,omitempty" json:"patient"`
	Note               []string                                   `firestore:"note,omitempty" json:"note"`
	Reaction           []AllergyIntoleranceReaction               `firestore:"reaction,omitempty" json:"reaction"`
}

type Condition struct {
	Identifier         []Identifier                      `firestore:"identifier,omitempty" json:"identifier"`
	ClinicalStatus     codes.ConditionClinicalStatus     `firestore:"clinicalStatus,omitempty" json:"clinicalStatus,omitempty"`
	VerificationStatus codes.ConditionVerificationStatus `firestore:"verificationStatus,omitempty" json:"verificationStatus,omitempty"`
	Category           []CodeableConcept                 `firestore:"category,omitempty" json:"category,omitempty"`
	Severity           CodeableConcept                   `firestore:"severity,omitempty" json:"severity,omitempty"`
	Code               CodeableConcept                   `firestore:"code,omitempty" json:"code,omitempty"`
	BodySite           []CodeableConcept                 `firestore:"bodySite,omitempty" json:"bodySite,omitempty"`
	Subject            Reference                         `firestore:"subject,omitempty" json:"subject"`
	Context            *Reference                        `firestore:"context,omitempty" json:"context,omitempty"`
	OnsetDateTime      *time.Time                        `firestore:"onsetDateTime,omitempty" json:"onsetDateTime,omitempty"`
	OnsetPeriod        *Period                           `firestore:"onsetPeriod,omitempty" json:"onsetPeriod,omitempty"`
	OnsetString        string                            `firestore:"onsetString,omitempty" json:"onsetString,omitempty"`
	AbatementDateTime  *time.Time                        `firestore:"abatement,omitempty" json:"abatement,omitempty"`
	AbatementBoolean   bool                              `firestore:"abatementBoolean,omitempty" json:"abatementBoolean,omitempty"`
	AbatementPeriod    *Period                           `firestore:"abatementPeriod,omitempty" json:"abatementPeriod,omitempty"`
	AbatementString    string                            `firestore:"abatementString,omitempty" json:"abatementString,omitempty"`
	AssertedDate       *time.Time                        `firestore:"assertedDate,omitempty" json:"assertedDate,omitempty"`
	Asserter           *Reference                        `firestore:"asserter,omitempty" json:"asserter,omitempty"`
	Stage              *ConditionStage                   `firestore:"stage,omitempty" json:"stage,omitempty"`
	Evidence           []ConditionEvidence               `firestore:"evidence,omitempty" json:"evidence,omitempty"`
	Note               []Annotation                      `firestore:"note,omitempty" json:"note,omitempty"`
}

type Procedure struct {
	// http://hl7.org/implement/standards/emr/STU3/procedure.html
	BaseResource
	Definition         []Reference          `firestore:"definition,omitempty" json:"definition,omitempty"`
	BasedOn            []Reference          `firestore:"basedOn,omitempty" json:"basedOn,omitempty"`
	PartOf             []Reference          `firestore:"partOf,omitempty" json:"partOf,omitempty"`
	Status             codes.EventStatus    `firestore:"status,omitempty" json:"status"`
	NotDone            bool                 `firestore:"notDone,omitempty" json:"notDone,omitempty"`
	NotDoneReason      *CodeableConcept     `firestore:"notDoneReason,omitempty" json:"notDoneReason,omitempty"`
	Category           *CodeableConcept     `firestore:"category,omitempty" json:"category,omitempty"`
	Subject            Reference            `firestore:"subject,omitempty" json:"subject"`
	Context            *Reference           `firestore:"context,omitempty" json:"context,omitempty"`
	PerformedDateTime  *time.Time           `firestore:"performedDateTime,omitempty" json:"performedDateTime,omitempty"`
	PerformedPeriod    *Period              `firestore:"performedPeriod,omitempty" json:"performedPeriod,omitempty"`
	Performer          []ProcedurePerformer `firestore:"performer,omitempty" json:"performer,omitempty"`
	Location           *Reference           `firestore:"location,omitempty" json:"location,omitempty"`
	ReasonCode         []CodeableConcept    `firestore:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference    []Reference          `firestore:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	BodySite           []CodeableConcept    `firestore:"bodySite,omitempty" json:"bodySite,omitempty"`
	Outcome            *CodeableConcept     `firestore:"outcome,omitempty" json:"outcome,omitempty"`
	Report             []Reference          `firestore:"report,omitempty" json:"report,omitempty"`
	Complication       []CodeableConcept    `firestore:"complication,omitempty" json:"complication,omitempty"`
	ComplicationDetail []Reference          `firestore:"complicationDetail,omitempty" json:"complicationDetail,omitempty"`
	FollowUp           []CodeableConcept    `firestore:"followUp,omitempty" json:"followUp,omitempty"`
	Note               []Annotation         `firestore:"note,omitempty" json:"note,omitempty"`
}

type ProcedurePerformer struct {
	Role       *CodeableConcept `firestore:"role,omitempty" json:"role,omitempty"`
	Actor      Reference        `firestore:"actor,omitempty" json:"actor"`
	OnBehalfOf *Reference       `firestore:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
}

type FamilyMemberHistory struct {
	// http://hl7.org/implement/standards/emr/STU3/familymemberhistory.html
	BaseResource
	Definition       []Reference                    `firestore:"definition,omitempty" json:"definition,omitempty"`
	Status           codes.FamilyHistoryStatusCode  `firestore:"status,omitempty" json:"status"`
	NotDone          bool                           `firestore:"notDone,omitempty" json:"notDone,omitempty"`
	NotDoneReason    *CodeableConcept               `firestore:"notDoneReason,omitempty" json:"notDoneReason,omitempty"`
	Patient          Reference                      `firestore:"patient,omitempty" json:"patient"`
	Date             *time.Time                     `firestore:"date,omitempty" json:"date,omitempty"`
	Name             string                         `firestore:"name,omitempty" json:"name,omitempty"`
	Relationship     CodeableConcept                `firestore:"relationship,omitempty" json:"relationship,omitempty"`
	Gender           *codes.AdministrativeGender    `firestore:"gender,omitempty" json:"gender,omitempty"`
	BornPeriod       *Period                        `firestore:"bornPeriod,omitempty" json:"bornPeriod,omitempty"`
	BornDate         *time.Time                     `firestore:"bornDate,omitempty" json:"bornDate,omitempty"`
	BornString       string                         `firestore:"bornString,omitempty" json:"bornString,omitempty"`
	AgeAge           int                            `firestore:"ageAge,omitempty" json:"ageAge,omitempty"`
	AgeString        string                         `firestore:"ageString,omitempty" json:"ageString,omitempty"`
	EstimatedAge     bool                           `firestore:"estimatedAge,omitempty" json:"estimatedAge,omitempty"`
	DeceasedBoolean  bool                           `firestore:"deceasedBoolean,omitempty" json:"deceasedBoolean,omitempty"`
	DeceasedDateTime *time.Time                     `firestore:"deceasedDateTime,omitempty" json:"deceasedDateTime,omitempty"`
	ReasonCode       []CodeableConcept              `firestore:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference  []Reference                    `firestore:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Note             []Annotation                   `firestore:"note,omitempty" json:"note,omitempty"`
	Condition        []FamilyMemberHistoryCondition `firestore:"condition,omitempty" json:"condition"`
}

type FamilyMemberHistoryCondition struct {
	Code        CodeableConcept  `firestore:"code,omitempty" json:"code"`
	Outcome     *CodeableConcept `firestore:"outcome,omitempty" json:"outcome,omitempty"`
	OnsetPeriod *Period          `firestore:"onsetPeriod,omitempty" json:"onsetPeriod,omitempty"`
	OnsetString string           `firestore:"onsetString,omitempty" json:"onsetString,omitempty"`
	Note        []Annotation     `firestore:"note,omitempty" json:"note"`
}

type ClinicalImpression struct {
	// http://hl7.org/implement/standards/emr/STU3/clinicalimpression.html
	BaseResource
	Status                   codes.ClinicalImpressionStatus    `firestore:"status,omitempty" json:"status"`
	Code                     *CodeableConcept                  `firestore:"code,omitempty" json:"code,omitempty"`
	Description              string                            `firestore:"description,omitempty" json:"description,omitempty"`
	Subject                  Reference                         `firestore:"subject,omitempty" json:"subject"`
	Context                  *Reference                        `firestore:"context,omitempty" json:"context,omitempty"`
	EffectiveDateTime        *time.Time                        `firestore:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"` // Time of assessment
	Date                     *time.Time                        `firestore:"date,omitempty" json:"date,omitempty"`              // When the assessment was documented
	Assessor                 *Reference                        `firestore:"assessor,omitempty" json:"assessor,omitempty"`
	Previous                 *Reference                        `firestore:"previous,omitempty" json:"previous,omitempty"`
	Problem                  []Reference                       `firestore:"problem,omitempty" json:"problem,omitempty"`
	Investigation            []ClinicalImpressionInvestigation `firestore:"investigation,omitempty" json:"investigation,omitempty"`
	Protocol                 []string                          `firestore:"protocol,omitempty" json:"protocol,omitempty"`
	Summary                  string                            `firestore:"summary,omitempty" json:"summary,omitempty"`
	Finding                  []ClinicalImpressionFinding       `firestore:"finding,omitempty" json:"finding,omitempty"`
	PrognosisCodeableConcept []CodeableConcept                 `firestore:"prognosisCodeableConcept,omitempty" json:"prognosisCodeableConcept,omitempty"`
	PrognosisReference       []Reference                       `firestore:"prognosisReference,omitempty" json:"prognosisReference,omitempty"`
	Action                   []Reference                       `firestore:"action,omitempty" json:"action,omitempty"`
	Note                     []Annotation                      `firestore:"note,omitempty" json:"note,omitempty"`
}

type ClinicalImpressionInvestigation struct {
	Code CodeableConcept `firestore:"code,omitempty" json:"code"`
	Item []Reference     `firestore:"item,omitempty" json:"item,omitempty"`
}

type ClinicalImpressionFinding struct {
	ItemCodeableConcept *CodeableConcept `firestore:"itemCodeableConcept,omitempty" json:"itemCodeableConcept,omitempty"`
	ItemReference       *Reference       `firestore:"itemReference,omitempty" json:"itemReference,omitempty"`
	Basis               string           `firestore:"basis,omitempty" json:"basis"`
}
