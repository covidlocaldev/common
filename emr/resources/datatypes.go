package resources

import (
	"fmt"
	"github.com/covidlocaldev/common/fhir/codes"
	"time"
)

type Reference struct {
	Reference  string      `firestore:"reference,omitempty" json:"reference"`
	Identifier *Identifier `firestore:"identifier,omitempty" json:"identifier,omitempty"`
	Display    string      `firestore:"display,omitempty" json:"display"`
}

func ResourceReference(resourceType, resourceID, display string) *Reference {
	return &Reference{
		Reference: fmt.Sprintf("%s/%s", resourceType, resourceID),
		Display:   display,
	}
}

type Date string

type Identifier struct {
	Use      codes.IdentifierUse `firestore:"use,omitempty" json:"use"`
	Type     CodeableConcept     `firestore:"type,omitempty" json:"type"`
	System   string              `firestore:"system,omitempty" json:"system"` // The namespace for the identifier value
	Value    string              `firestore:"value,omitempty" json:"value"`
	Assigner *Reference          `firestore:"assigner,omitempty" json:"assigner"` // Reference Types: Organization
}

type HumanName struct {
	Use    codes.NameUse `firestore:"use,omitempty" json:"use"`
	Text   string        `firestore:"text,omitempty" json:"text,omitempty"`
	Family string        `firestore:"family,omitempty" json:"family,omitempty"`
	Given  []string      `firestore:"given,omitempty" json:"given,omitempty"`
	Prefix string        `firestore:"prefix,omitempty" json:"prefix,omitempty"`
	Suffix string        `firestore:"suffix,omitempty" json:"suffix,omitempty"`
}

type ContactPoint struct {
	System codes.ContactPointSystem `firestore:"system,omitempty" json:"system"`
	Value  string                   `firestore:"value,omitempty" json:"value,omitempty"`
	Use    codes.ContactPointUse    `firestore:"use,omitempty" json:"use"`
	Rank   int                      `firestore:"rank,omitempty" json:"rank,omitempty"`
}

type Address struct {
	Use        codes.AddressUse  `firestore:"use,omitempty" json:"use"`
	Type       codes.AddressType `firestore:"type,omitempty" json:"type"`
	Text       string            `firestore:"text,omitempty" json:"text,omitempty"`
	Line       []string          `firestore:"line,omitempty" json:"line"`
	City       string            `firestore:"city,omitempty" json:"city"`
	District   string            `firestore:"district,omitempty" json:"district,omitempty"`
	State      string            `firestore:"state,omitempty" json:"state"`
	PostalCode string            `firestore:"postalCode,omitempty" json:"postalCode"`
	Country    string            `firestore:"country,omitempty" json:"country"`
}

type Attachment struct {
	ID          string         `firestore:"id,omitempty" json:"id"`
	ContentType codes.MimeType `firestore:"contentType,omitempty" json:"contentType"`
	Language    codes.Language `firestore:"language,omitempty" json:"language"`
	URL         string         `firestore:"url,omitempty" json:"url"`
	Size        uint64         `firestore:"size,omitempty" json:"size"`
	Title       string         `firestore:"title,omitempty" json:"title"`
	Creation    time.Time      `firestore:"creation,omitempty" json:"creation"`
}

type Contact struct {
	Relationship []codes.ContactRole         `firestore:"relationship,omitempty" json:"relationship"`
	Name         *HumanName                  `firestore:"name,omitempty" json:"name"`
	Telecom      []ContactPoint              `firestore:"telecom,omitempty" json:"telecom"`
	Address      *Address                    `firestore:"address,omitempty" json:"address"`
	Gender       *codes.AdministrativeGender `firestore:"gender,omitempty" json:"gender"`
}

type Communication struct {
	Language  CodeableConcept `firestore:"language,omitempty" json:"language"`
	Preferred bool            `firestore:"preferred,omitempty" json:"preferred"`
}

type PatientLink struct {
	Other Reference      `firestore:"other,omitempty" json:"other"` // Reference Types: Patient, RelatedPerson
	Type  codes.LinkType `firestore:"type,omitempty" json:"type"`
}

type Qualification struct {
	Identifier Identifier      `firestore:"identifier,omitempty" json:"identifier"`
	Code       CodeableConcept `firestore:"code,omitempty" json:"code"`
	Issuer     *Reference      `firestore:"issuer,omitempty" json:"issuer"` // Organization that regulates and issues the qualification
}

type PersonLink struct {
	Target    Reference                    `firestore:"target,omitempty" json:"target"`
	Assurance codes.IdentityAssuranceLevel `firestore:"assurance,omitempty" json:"assurance"`
}

type AllergyIntoleranceReaction struct {
	Substance     codes.SubstanceCode              `firestore:"substance,omitempty" json:"substance"`
	Manifestation []codes.ClinicalFindingCode      `firestore:"manifestation,omitempty" json:"manifestation"`
	Description   string                           `firestore:"description,omitempty" json:"description"`
	Onset         *time.Time                       `firestore:"onset,omitempty" json:"onset"`
	Severity      codes.AllergyIntoleranceSeverity `firestore:"severity,omitempty" json:"severity"`
	ExposureRoute codes.RouteCode                  `firestore:"exposureRoute,omitempty" json:"exposureRoute"`
	Note          []string                         `firestore:"note,omitempty" json:"note"`
}

type ConditionStage struct {
	Summary    CodeableConcept `firestore:"summary,omitempty" json:"summary,omitempty"`
	Assessment []Reference     `firestore:"assessment,omitempty" json:"assessment,omitempty"`
}

type ConditionEvidence struct {
	Code   []CodeableConcept `firestore:"code,omitempty" json:"code,omitempty"`
	Detail []Reference       `firestore:"detail,omitempty" json:"detail,omitempty"`
}

type Period struct {
	Start *time.Time `firestore:"start,omitempty" json:"start"`
	End   *time.Time `firestore:"end,omitempty" json:"end"`
}

type Coding struct {
	System       string `firestore:"system,omitempty" json:"system"`
	Version      string `firestore:"version,omitempty" json:"version"`
	Code         string `firestore:"code,omitempty" json:"code"`
	Display      string `firestore:"display,omitempty" json:"display"`
	UserSelected bool   `firestore:"userSelected,omitempty" json:"userSelected"`
}

type CodeableConcept struct {
	Coding []Coding `firestore:"coding,omitempty" json:"coding"`
	Text   string   `firestore:"text,omitempty" json:"text"`
}

type QuantityComparatorCode = string

const (
	QuantityComparatorCodeLessThan           QuantityComparatorCode = "<"
	QuantityComparatorCodeLessThanOrEqual    QuantityComparatorCode = "<="
	QuantityComparatorCodeGreaterThan        QuantityComparatorCode = ">"
	QuantityComparatorCodeGreaterThanOrEqual QuantityComparatorCode = ">="
)

type Quantity struct {
	Value      float64                `firestore:"value,omitempty" json:"value"`
	Comparator QuantityComparatorCode `firestore:"comparator,omitempty" json:"comparator"`
	Unit       string                 `firestore:"unit,omitempty" json:"unit"`
	System     string                 `firestore:"system,omitempty" json:"system"`
	Code       string                 `firestore:"code,omitempty" json:"code"`
}

type Range struct {
	Low  *Quantity `firestore:"low,omitempty" json:"low"`
	High *Quantity `firestore:"high,omitempty" json:"high"`
}

type Ratio struct {
	Numerator   *Quantity `firestore:"numerator,omitempty" json:"numerator"`
	Denominator *Quantity `firestore:"denominator,omitempty" json:"denominator"`
}

// http://hl7.org/implement/standards/fhir/STU3/datatypes.html#Timing
type Timing struct {
	Event  *time.Time       `firestore:"event,omitempty" json:"event"`
	Repeat *TimingRepeat    `firestore:"repeat,omitempty" json:"repeat"`
	Code   *CodeableConcept `firestore:"code,omitempty" json:"code"`
}

type TimingRepeat struct {
	Count        int64      `firestore:"count,omitempty" json:"count"`
	CountMax     int64      `firestore:"countMax,omitempty" json:"countMax"`
	Duration     float64    `firestore:"duration,omitempty" json:"duration"`
	DurationMax  float64    `firestore:"durationMax,omitempty" json:"durationMax"`
	DurationUnit UnitOfTime `firestore:"durationUnit,omitempty" json:"durationUnit"`
	Frequency    int64      `firestore:"frequency,omitempty" json:"frequency"`
	FrequencyMax int64      `firestore:"frequencyMax,omitempty" json:"frequencyMax"`
	Period       float64    `firestore:"period,omitempty" json:"period"`
	PeriodMax    float64    `firestore:"periodMax,omitempty" json:"periodMax"`
	PeriodUnit   UnitOfTime `firestore:"periodUnit,omitempty" json:"periodUnit"`
	DayOfWeek    *DayOfWeek `firestore:"dayOfWeek,omitempty" json:"dayOfWeek"`
	TimeOfDay    *time.Time `firestore:"timeOfDay,omitempty" json:"timeOfDay"`
}

type UnitOfTime = string

const (
	UnitOfTimeSecond UnitOfTime = "s"
	UnitOfTimeMinute UnitOfTime = "min"
	UnitOfTimeHour   UnitOfTime = "h"
	UnitOfTimeDay    UnitOfTime = "d"
	UnitOfTimeWeek   UnitOfTime = "wk"
	UnitOfTimeMonth  UnitOfTime = "mo"
)

type DayOfWeek = string

const (
	DayOfWeekMonday    DayOfWeek = "mon"
	DayOfWeekTuesday   DayOfWeek = "tue"
	DayOfWeekWednesday DayOfWeek = "wed"
	DayOfWeekThursday  DayOfWeek = "thu"
	DayOfWeekFriday    DayOfWeek = "fri"
	DayOfWeekSaturday  DayOfWeek = "saturday"
	DayOfWeekSunday    DayOfWeek = "sunday"
)

type Annotation struct {
	AuthorReference *Reference `firestore:"authorReference,omitempty" json:"authorReference,omitempty"`
	AuthorString    string     `firestore:"authorString,omitempty" json:"authorString,omitempty"`
	Time            *time.Time `firestore:"time,omitempty" json:"time,omitempty"`
	Text            string     `firestore:"text,omitempty" json:"text"`
}

type CarePlanActivityDetail struct {
	Category               *CodeableConcept             `firestore:"category,omitempty" json:"category,omitempty"`
	Definition             *Reference                   `firestore:"definition,omitempty" json:"definition,omitempty"`
	Code                   *CodeableConcept             `firestore:"code,omitempty" json:"code,omitempty"`
	ReasonCode             []CodeableConcept            `firestore:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference        []Reference                  `firestore:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Goal                   []Reference                  `firestore:"goal,omitempty" json:"goal,omitempty"`
	Status                 codes.CarePlanActivityStatus `firestore:"status,omitempty" json:"status"`
	StatusReason           string                       `firestore:"statusReason,omitempty" json:"statusReason,omitempty"`
	Prohibited             bool                         `firestore:"prohibited,omitempty" json:"prohibited,omitempty"`
	ScheduledPeriod        *Period                      `firestore:"scheduledPeriod,omitempty" json:"scheduledPeriod,omitempty"`
	ScheduledString        string                       `firestore:"scheduledString,omitempty" json:"scheduledString,omitempty"`
	Location               *Reference                   `firestore:"location,omitempty" json:"location,omitempty"`
	Performer              []Reference                  `firestore:"performer,omitempty" json:"performer,omitempty"`
	ProductCodeableConcept *CodeableConcept             `firestore:"productCodeableConcept,omitempty" json:"productCodeableConcept,omitempty"`
	ProductReference       *Reference                   `firestore:"productReference,omitempty" json:"productReference,omitempty"`
	DailyAmount            *Quantity                    `firestore:"dailyAmount,omitempty" json:"dailyAmount,omitempty"`
	Quantity               *Quantity                    `firestore:"quantity,omitempty" json:"quantity,omitempty"`
	Description            string                       `firestore:"description,omitempty" json:"description,omitempty"`
}

type CarePlanActivity struct {
	OutcomeCodeableConcept []CodeableConcept       `firestore:"outcomeCodeableConcept,omitempty" json:"outcomeCodeableConcept,omitempty"`
	OutcomeReference       []Reference             `firestore:"outcomeReference,omitempty" json:"outcomeReference,omitempty"`
	Progress               []Annotation            `firestore:"progress,omitempty" json:"progress,omitempty"`
	Reference              *Reference              `firestore:"reference,omitempty" json:"reference,omitempty"`
	Detail                 *CarePlanActivityDetail `firestore:"detail,omitempty" json:"detail,omitempty"`
	Note                   []Annotation            `firestore:"note,omitempty" json:"note,omitempty"`
}
