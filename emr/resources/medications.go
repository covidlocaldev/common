package resources

import "time"

type MedicationStatus = string

const (
	MedicationStatusActive         MedicationStatus = "active"
	MedicationStatusInactive       MedicationStatus = "inactive"
	MedicationStatusEnteredInError MedicationStatus = "entered-in-error"
)

type Medication struct {
	// http://hl7.org/implement/standards/emr/STU3/medication.html
	BaseResource
	Code             *CodeableConcept `firestore:"code,omitempty" json:"code"`
	Status           MedicationStatus `firestore:"status,omitempty" json:"status"`
	IsBrand          bool             `firestore:"isBrand,omitempty" json:"isBrand"`
	IsOverTheCounter bool             `firestore:"isOverTheCounter,omitempty" json:"isOverTheCounter"`
	Manufacturer     *Reference       `firestore:"manufacturer,omitempty" json:"manufacturer"`
	Form             *CodeableConcept `firestore:"form,omitempty" json:"form"`
}

func (m Medication) GetResourceType() string { return "Medication" }

type MedicationRequestStatus = string

const (
	MedicationRequestStatusActive         MedicationRequestStatus = "active"
	MedicationRequestStatusOnHold         MedicationRequestStatus = "on-hold"
	MedicationRequestStatusCancelled      MedicationRequestStatus = "cancelled"
	MedicationRequestStatusCompleted      MedicationRequestStatus = "completed"
	MedicationRequestStatusEnteredInError MedicationRequestStatus = "entered-in-error"
	MedicationRequestStatusStopped        MedicationRequestStatus = "stopped"
	MedicationRequestStatusDraft          MedicationRequestStatus = "draft"
	MedicationRequestStatusUnknown        MedicationRequestStatus = "unknown"
)

type MedicationRequestIntent = string

const (
	MedicationRequestIntentProposal      MedicationRequestIntent = "proposal"
	MedicationRequestIntentPlan          MedicationRequestIntent = "plan"
	MedicationRequestIntentOrder         MedicationRequestIntent = "order"
	MedicationRequestIntentInstanceOrder MedicationRequestIntent = "instance-order"
)

type MedicationRequestPriority = string

const (
	MedicationRequestPriorityRoutine MedicationRequestPriority = "routine"
	MedicationRequestPriorityUrgent  MedicationRequestPriority = "urgent"
	MedicationRequestPriorityStat    MedicationRequestPriority = "stat"
	MedicationRequestPriorityASAP    MedicationRequestPriority = "asap"
)

type MedicationRequestRequester struct {
	Agent      Reference  `firestore:"agent,omitempty" json:"agent"`
	OnBehalfOf *Reference `firestore:"onBehalfOf,omitempty" json:"onBehalfOf"`
}

type Dosage struct {
	BaseResource
	Sequence                 int64             `firestore:"sequence,omitempty" json:"sequence"`
	Text                     string            `firestore:"text,omitempty" json:"text"`
	AdditionalInstruction    []CodeableConcept `firestore:"additionalInstruction,omitempty" json:"additionalInstruction"`
	PatientInstruction       string            `firestore:"patientInstruction,omitempty" json:"patientInstruction"`
	Timing                   *Timing           `firestore:"timing,omitempty" json:"timing"`
	AsNeededBoolean          bool              `firestore:"asNeededBoolean,omitempty" json:"asNeededBoolean"`
	Site                     *CodeableConcept  `firestore:"site,omitempty" json:"site"`
	Route                    *CodeableConcept  `firestore:"route,omitempty" json:"route"`
	Method                   *CodeableConcept  `firestore:"method,omitempty" json:"method"`
	DoseQuantity             *Quantity         `firestore:"doseQuantity,omitempty" json:"doseQuantity"`
	MaxDosePerPeriod         *Ratio            `firestore:"maxDosePeriod,omitempty" json:"maxDosePeriod"`
	MaxDosePerAdministration *Quantity         `firestore:"maxDosePerAdministration,omitempty" json:"maxDosePerAdministration"`
	MaxDosePerLifetime       *Quantity         `firestore:"maxDosePerLifetime,omitempty" json:"maxDosePerLifetime"`
	RateQuantity             *Quantity         `firestore:"rateQuantity,omitempty" json:"rateQuantity"`
}

type DispenseRequest struct {
	ValidityPeriod         *Period    `firestore:"validityPeriod,omitempty" json:"validityPeriod"`
	NumberOfRepeatsAllowed int64      `firestore:"numberOfRepeatsAllowed,omitempty" json:"numberOfRepeatsAllowed"`
	Quantity               *Quantity  `firestore:"quantity,omitempty" json:"quantity"`
	ExpectedSupplyDuration *Quantity  `firestore:"expectedSupplyDuration,omitempty" json:"expectedSupplyDuration"`
	Performer              *Reference `firestore:"performer,omitempty" json:"performer"`
}

type MedicationRequestSubstitution struct {
	Allowed bool             `firestore:"allowed,omitempty" json:"allowed"`
	Reason  *CodeableConcept `firestore:"reason,omitempty" json:"reason"`
}

type MedicationRequest struct {
	// http://hl7.org/implement/standards/emr/STU3/medicationrequest.html
	BaseResource
	Definition                []Reference                    `firestore:"definition,omitempty" json:"definition"`
	BasedOn                   []Reference                    `firestore:"basedOn,omitempty" json:"basedOn"`
	GroupIdentifier           *Identifier                    `firestore:"groupIdentifier,omitempty" json:"groupIdentifier"`
	Status                    MedicationRequestStatus        `firestore:"status,omitempty" json:"status"`
	Intent                    MedicationRequestIntent        `firestore:"intent,omitempty" json:"intent"`
	Category                  *CodeableConcept               `firestore:"category,omitempty" json:"category"`
	Priority                  MedicationRequestPriority      `firestore:"priority,omitempty" json:"priority"`
	MedicationCodeableConcept *CodeableConcept               `firestore:"medicationCodeableConcept,omitempty" json:"medicationCodeableConcept"`
	MedicationReference       *Reference                     `firestore:"medicationReference,omitempty" json:"medicationReference"`
	Subject                   Reference                      `firestore:"subject,omitempty" json:"subject"`
	Context                   *Reference                     `firestore:"context,omitempty" json:"context"`
	SupportingInformation     []Reference                    `firestore:"supportingInformation,omitempty" json:"supportingInformation"`
	AuthoredOn                *time.Time                     `firestore:"authoredOn,omitempty" json:"authoredOn"`
	Requester                 *MedicationRequestRequester    `firestore:"requester,omitempty" json:"requester"`
	Recorder                  *Reference                     `firestore:"recorder,omitempty" json:"recorder"`
	ReasonCode                []CodeableConcept              `firestore:"reasonCode,omitempty" json:"reasonCode"`
	ReasonReference           *Reference                     `firestore:"reasonReference,omitempty" json:"reasonReference"`
	Note                      []Annotation                   `firestore:"note,omitempty" json:"note"`
	DosageInstruction         []Dosage                       `firestore:"dosageInstruction,omitempty" json:"dosageInstruction"`
	DispenseRequest           *DispenseRequest               `firestore:"dispenseRequest,omitempty" json:"dispenseRequest"`
	Substitution              *MedicationRequestSubstitution `firestore:"substitution,omitempty" json:"substitution"`
	PriorPrescription         *Reference                     `firestore:"priorPrescription,omitempty" json:"priorPrescription"`
	DetectedIssue             *Reference                     `firestore:"detectedIssue,omitempty" json:"detectedIssue"`
	EventHistory              []Reference                    `firestore:"eventHistory,omitempty" json:"eventHistory"`
}
