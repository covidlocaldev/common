package resources

import (
	"errors"
	"fmt"
)

type EncounterStatus string

const (
	EncounterStatusPlanned    EncounterStatus = "planned"
	EncounterStatusArrived    EncounterStatus = "arrived"
	EncounterStatusTriaged    EncounterStatus = "triaged"
	EncounterStatusInProgress EncounterStatus = "in-progress"
	EncounterStatusOnLeave    EncounterStatus = "on-leave"
	EncounterStatusFinished   EncounterStatus = "finished"
)

func (e EncounterStatus) GetCodes() []Code {
	return []Code{EncounterStatusPlanned, EncounterStatusArrived, EncounterStatusTriaged, EncounterStatusInProgress, EncounterStatusOnLeave, EncounterStatusFinished}
}
func (e EncounterStatus) String() string { return string(e) }
func (e *EncounterStatus) UnmarshalJSON(b []byte) error {
	raw := string(b)[1 : len(b)-1]
	unmarshalled := UnmarshalCode(e, raw)

	if unmarshalled != nil {
		*e = unmarshalled.(EncounterStatus)
		return nil
	} else {
		return errors.New(fmt.Sprintf("%s is not a valid EncounterStatus", b))
	}
}

type EncounterStatusHistory struct {
	Status EncounterStatus `firestore:"status,omitempty" json:"status"`
	Period Period          `firestore:"period,omitempty" json:"period"`
}

type EncounterParticipant struct {
	Type       []CodeableConcept `firestore:"type,omitempty" json:"type"`
	Period     *Period           `firestore:"period,omitempty" json:"period,omitempty"`
	Individual *Reference        `firestore:"individual,omitempty" json:"individual,omitempty"`
}

type EncounterDiagnosis struct {
	Condition Reference        `firestore:"condition,omitempty" json:"condition,omitempty"`
	Role      *CodeableConcept `firestore:"role,omitempty" json:"role,omitempty"`
	Rank      float64          `firestore:"rank,omitempty" json:"rank,omitempty"`
}

type Encounter struct {
	BaseResource
	Status        EncounterStatus          `firestore:"status,omitempty" json:"status"`
	StatusHistory []EncounterStatusHistory `firestore:"statusHistory,omitempty" json:"statusHistory,omitempty"`
	Class         Coding                   `firestore:"class,omitempty" json:"class,omitempty"`
	Type          []CodeableConcept        `firestore:"type,omitempty" json:"type,omitempty"`
	Priority      *CodeableConcept         `firestore:"priority,omitempty" json:"priority,omitempty"`
	Subject       *Reference               `firestore:"subject,omitempty" json:"subject,omitempty"`
	EpisodeOfCare []Reference              `firestore:"episodeOfCare,omitempty" json:"episodeOfCare,omitempty"`
	Participant   []EncounterParticipant   `firestore:"participant,omitempty" json:"participant,omitempty"`
	Appointment   *Reference               `firestore:"appointment,omitempty" json:"appointment,omitempty"`
	Period        *Period                  `firestore:"period,omitempty" json:"period,omitempty"`
	Reason        []CodeableConcept        `firestore:"reason,omitempty" json:"reason,omitempty"`
	Diagnosis     []EncounterDiagnosis     `firestore:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	PartOf        *Reference               `firestore:"partOf,omitempty" json:"partOf,omitempty"`
}

func (e Encounter) GetResourceType() string      { return "Encounter" }

type EpisodeOfCare struct{}
