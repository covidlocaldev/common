package resources

import (
	"github.com/covidlocaldev/common/emr/codes"
	"time"
)

type Person struct {
	// http://hl7.org/implement/standards/emr/STU3/person.html
	BaseResource
	Identifier []Identifier               `firestore:"identifier,omitempty" json:"identifier"`
	Active     bool                       `firestore:"active,omitempty" json:"active,omitempty"`
	Name       []HumanName                `firestore:"name,omitempty" json:"name,omitempty"`
	Telecom    []ContactPoint             `firestore:"telecom,omitempty" json:"telecom,omitempty"`
	Gender     codes.AdministrativeGender `firestore:"gender,omitempty" json:"gender,omitempty"`
	BirthDate  *Date                      `firestore:"birthDate,omitempty" json:"birthDate,omitempty"`
	Address    []Address                  `firestore:"address,omitempty" json:"address,omitempty"`
	Photo      []Attachment               `firestore:"photo,omitempty" json:"photo,omitempty"`
	Link       []PersonLink               `firestore:"link,omitempty" json:"link,omitempty"`
}

func (p Person) GetResourceType() string { return "Person" }

type Patient struct {
	// http://hl7.org/implement/standards/emr/STU3/patient.html
	Person
	DeceasedBoolean      bool             `firestore:"deceasedBoolean,omitempty" json:"deceasedBoolean,omitempty"`
	DeceasedDateTime     *time.Time       `firestore:"deceasedDateTime,omitempty" json:"deceasedDateTime,omitempty"`
	MaritalStatus        *CodeableConcept `firestore:"maritalStatus,omitempty" json:"maritalStatus"`
	Contact              []Contact        `firestore:"contact,omitempty" json:"contact"`
	Communication        []Communication  `firestore:"communication,omitempty" json:"communication"`
	GeneralPractitioner  []Reference      `firestore:"generalPractitioner,omitempty" json:"generalPractitioner"`  // Reference Types: Organization, Practitioner
	ManagingOrganization *Reference       `firestore:"managingOrganization,omitempty" json:"managingOrganization"` // Reference Types: Organization
	Link                 []PatientLink    `firestore:"link,omitempty" json:"link"`
}

func (p Patient) GetResourceType() string { return "Patient" }

type RelatedPerson struct {
	// http://hl7.org/implement/standards/emr/STU3/relatedperson.html
	Person
	Relationship CodeableConcept `firestore:"relationship,omitempty" json:"relationship"`
}

type Practitioner struct {
	// http://hl7.org/implement/standards/emr/STU3/practitioner.html
	Person
	Qualification []Qualification   `firestore:"qualification,omitempty" json:"qualification"`
	Communication []CodeableConcept `firestore:"communication,omitempty" json:"communication"`
}

func (p Practitioner) GetResourceType() string      { return "Practitioner" }

type PractitionerRole struct {
	// http://hl7.org/implement/standards/emr/STU3/practitionerrole.html
	Identifier        []Identifier                 `firestore:"identifier,omitempty" json:"identifier"`
	Active            bool                         `firestore:"active,omitempty" json:"active"`
	Organization      Organization                 `firestore:"organization,omitempty" json:"organization"` // Organization where the roles are available
	Code              []codes.PractitionerRoleCode `firestore:"code,omitempty" json:"code"`
	Speciality        []codes.Specialty            `firestore:"specialty,omitempty" json:"specialty"`
	Location          []Location                   `firestore:"location,omitempty" json:"location"`
	HealthcareService []HealthcareService          `firestore:"healthcareService,omitempty" json:"healthcareService"`
	Telecom           []ContactPoint               `firestore:"telecom,omitempty" json:"telecom"`
}
