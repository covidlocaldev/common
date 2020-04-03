package resources

import (
	"github.com/covidlocaldev/common/fhir/codes"
)

type Organization struct {
	BaseResource
	Active  bool              `firestore:"active,omitempty" json:"active"`
	Type    []CodeableConcept `firestore:"type,omitempty" json:"type"`
	Name    string            `firestore:"name,omitempty" json:"name"`
	Alias   []string          `firestore:"alias,omitempty" json:"alias"`
	Telecom []ContactPoint    `firestore:"telecom,omitempty" json:"telecom"`
	Address []Address         `firestore:"address,omitempty" json:"address"`
	PartOf  *Reference         `firestore:"partOf,omitempty" json:"partOf,omitempty"`
}

func (o Organization) GetResourceType() string { return "Organization" }
func (o Organization) IsProcedurePerformer()        {}
func (o Organization) IsObservationPerformer()      {}
func (o Organization) IsCarePlanAuthor()            {}
func (o Organization) IsGoalSubject()               {}
func (o Organization) IsCarePlanActivityPerformer() {}

type Location struct {
	// http://hl7.org/implement/standards/fhir/STU3/location.html
	Identifier   []Identifier                          `firestore:"identifier,omitempty" json:"identifier"`
	Status       codes.LocationStatus                  `firestore:"status,omitempty" json:"status"`
	Name         string                                `firestore:"name,omitempty" json:"name"`
	Mode         codes.LocationMode                    `firestore:"mode,omitempty" json:"mode"`
	Type         codes.ServiceDeliveryLocationRoleType `firestore:"type,omitempty" json:"type"`
	Telecom      []ContactPoint                        `firestore:"telecom,omitempty" json:"telecom"`
	Address      Address                               `firestore:"address,omitempty" json:"address"`
	PhysicalType codes.LocationType                    `firestore:"physicalType,omitempty" json:"physicalType"`
}

func (l Location) IsObservationSubject() {}

type HealthcareService struct {
	Identifier Identifier            `firestore:"identifier,omitempty" json:"identifier"`
	Active     bool                  `firestore:"active,omitempty" json:"active"`
	Category   codes.ServiceCategory `firestore:"category,omitempty" json:"category"`
	Type       []codes.ServiceType   `firestore:"type,omitempty" json:"type"`
	Specialty  []codes.Specialty     `firestore:"specialty,omitempty" json:"specialty"`
	Location   []Location            `firestore:"location,omitempty" json:"location"`
	Name       string                `firestore:"name,omitempty" json:"name"`
	Comment    string                `firestore:"comment,omitempty" json:"comment"`
	Photo      Attachment            `firestore:"photo,omitempty" json:"photo"`
	Telecom    []ContactPoint        `firestore:"telecom,omitempty" json:"telecom"`
}

func (hs HealthcareService) IsProcedureDefinition() {}

type Substance struct{}

func (s Substance) IsProductReference() {}
