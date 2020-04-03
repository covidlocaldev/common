package codes

type IdentifierUse = string

const (
	IdentifierUseUsual     IdentifierUse = "usual"
	IdentifierUseOfficial  IdentifierUse = "official"
	IdentifierUseSecondary IdentifierUse = "secondary"
)

type IdentifierType = string // http://hl7.org/implement/standards/fhir/STU3/v2/0203/index.html#v2-0203-PRN

const (
	IdentifierTypeDriversLicenseNumber             IdentifierType = "DL"
	IdentifierTypeMedicalLicenseNumber             IdentifierType = "MD"
	IdentifierTypeNationalInsurancePayorIdentifier IdentifierType = "NIIP"
	IdentifierTypePassportNumber                   IdentifierType = "PPN"
	IdentifierTypePersonNumber                     IdentifierType = "PN"
	IdentifierTypeProviderNumber                   IdentifierType = "PRN"
	IdentifierTypeResourceNumber                   IdentifierType = "RN"
)

type NameUse = string

const (
	NameUseUsual     NameUse = "usual"
	NameUseOfficial  NameUse = "official"
	NameUseAnonymous NameUse = "anonymous"
)

type ContactPointSystem = string

const (
	ContactPointSystemPhone ContactPointSystem = "phone"
	ContactPointSystemFax   ContactPointSystem = "fax"
	ContactPointSystemEmail ContactPointSystem = "email"
	ContactPointSystemPager ContactPointSystem = "pager"
	ContactPointSystemURL   ContactPointSystem = "url"
	ContactPointSystemSMS   ContactPointSystem = "sms"
	ContactPointSystemOther ContactPointSystem = "other"
)

type ContactPointUse = string

const (
	ContactPointUseHome   ContactPointUse = "home"
	ContactPointUseWork   ContactPointUse = "work"
	ContactPointUseMobile ContactPointUse = "mobile"
)

type AdministrativeGender = string

const (
	AdministrativeGenderMale    AdministrativeGender = "male"
	AdministrativeGenderFemale  AdministrativeGender = "female"
	AdministrativeGenderOther   AdministrativeGender = "other"
	AdministrativeGenderUnknown AdministrativeGender = "unknown"
)

type AddressUse = string

const (
	AddressUseHome AddressUse = "home"
	AddressUseWork AddressUse = "work"
)

type AddressType = string

const (
	AddressTypePostal   AddressType = "postal"
	AddressTypePhysical AddressType = "physical"
	AddressTypeBoth     AddressType = "both"
)

type MaritalStatus = string

const (
	MartialStatusNeverMarried MaritalStatus = "S"
	MaritalStatusMarried      MaritalStatus = "M"
	MaritalStatusDivorced     MaritalStatus = "D"
)

type MimeType = string

const (
	MimeTypeBMP  MimeType = "image/bmp"
	MimeTypeJPEG MimeType = "image/jpeg"
	MimeTypePNG  MimeType = "image/png"
	MimeTypeWAV  MimeType = "audio/wav"
)

type Language = string

const (
	LanguageEnglish Language = "en"
	LanguageSpanish Language = "es"
)

type ContactRole = string

const (
	ContactRoleEmergencyContact ContactRole = "C"
	ContactRoleNextOfKin        ContactRole = "N"
	ContactRoleOther            ContactRole = "O"
	ContactRoleUnknown          ContactRole = "U"
)

type LinkType = string

const (
	LinkTypeReplacedBy LinkType = "replaced-by"
	LinkTypeReplaces   LinkType = "replaces"
	LinkTypeRefer      LinkType = "refer"
	LinkTypeSeeAlso    LinkType = "seealso"
)

type PatientRelationshipType = string

const (
	PatientRelationshipTypeEmergencyContact       PatientRelationshipType = "C"
	PatientRelationshipTypeContactPerson          PatientRelationshipType = "CP"
	PatientRelationshipTypeEmergencyContactPerson PatientRelationshipType = "EP"
	PatientRelationshipTypeNextOfKin              PatientRelationshipType = "N"
	PatientRelationshipTypeOther                  PatientRelationshipType = "O"
	PatientRelationshipTypeUnknown                PatientRelationshipType = "U"
)

type QualificationCode = string // http://hl7.org/implement/standards/fhir/STU3/v2/0360/2.7/index.html

type IdentityAssuranceLevel = string

const (
	IdentityAssuranceLevel1 IdentityAssuranceLevel = "level1"
	IdentityAssuranceLevel2 IdentityAssuranceLevel = "level2"
	IdentityAssuranceLevel3 IdentityAssuranceLevel = "level3"
	IdentityAssuranceLevel4 IdentityAssuranceLevel = "level4"
)

type PractitionerRoleCode = string

const (
	PractitionerRoleCodeDoctor     PractitionerRoleCode = "doctor"
	PractitionerRoleCodeNurse      PractitionerRoleCode = "nurse"
	PractitionerRoleCodePharmacist PractitionerRoleCode = "pharmacist"
	PractitionerRoleCodeResearcher PractitionerRoleCode = "researcher"
	PractitionerRoleCodeTeacher    PractitionerRoleCode = "teacher"
)

type Specialty = string // http://hl7.org/implement/standards/fhir/STU3/valueset-c80-practice-codes.html

type LocationStatus = string

const (
	LocationStatusActive    LocationStatus = "active"
	LocationStatusSuspended LocationStatus = "suspended"
	LocationStatusInactive  LocationStatus = "inactive"
)

type LocationMode = string

const (
	LocationModeInstance LocationMode = "instance"
	LocationModeKind     LocationMode = "kind"
)

type ServiceDeliveryLocationRoleType = string // http://hl7.org/implement/standards/fhir/STU3/v3/ServiceDeliveryLocationRoleType/vs.html
const (
	ServiceDeliveryLocationRoleTypeHospital           ServiceDeliveryLocationRoleType = "HOSP"
	ServiceDeliveryLocationRoleTypeOutpatientFacility ServiceDeliveryLocationRoleType = "OF"
)

type LocationType = string

const (
	LocationTypeSite     LocationType = "si"
	LocationTypeBuilding LocationType = "bu"
)

type ServiceCategory = string // http://hl7.org/implement/standards/fhir/STU3/valueset-service-category.html

type ServiceType = string // http://hl7.org/implement/standards/fhir/STU3/valueset-service-type.html

type AllergyIntoleranceClinicalStatus = string

const (
	AllergyIntoleranceClinicalStatusActive   AllergyIntoleranceClinicalStatus = "active"
	AllergyIntoleranceClinicalStatusInactive AllergyIntoleranceClinicalStatus = "inactive"
	AllergyIntoleranceClinicalStatusResolved AllergyIntoleranceClinicalStatus = "resolved"
)

type AllergyIntoleranceVerificationStatus = string

const (
	AllergyIntoleranceVerificationStatusUnconfirmed    AllergyIntoleranceVerificationStatus = "unconfirmed"
	AllergyIntoleranceVerificationStatusConfirmed      AllergyIntoleranceVerificationStatus = "confirmed"
	AllergyIntoleranceVerificationStatusRefuted        AllergyIntoleranceVerificationStatus = "refuted"
	AllergyIntoleranceVerificationStatusEnteredInError AllergyIntoleranceVerificationStatus = "entered-in-error"
)

type AllergyIntoleranceType = string

const (
	AllergyIntoleranceTypeAllergy            AllergyIntoleranceType = "allergy"
	AllergyIntoleranceTypeAllergyIntolerance AllergyIntoleranceType = "intolerance"
)

type AllergyIntoleranceCategory = string

const (
	AllergyIntoleranceCategoryFood        AllergyIntoleranceCategory = "food"
	AllergyIntoleranceCategoryMedication  AllergyIntoleranceCategory = "medication"
	AllergyIntoleranceCategoryEnvironment AllergyIntoleranceCategory = "environment"
	AllergyIntoleranceCategoryBiologic    AllergyIntoleranceCategory = "biologic"
)

type AllergyIntoleranceCriticality = string

const (
	AllergyIntoleranceCriticalityLow            AllergyIntoleranceCriticality = "low"
	AllergyIntoleranceCriticalityHigh           AllergyIntoleranceCriticality = "high"
	AllergyIntoleranceCriticalityUnableToAssess AllergyIntoleranceCriticality = "unable-to-asses"
)

type AllergyIntoleranceCode = string // http://hl7.org/implement/standards/fhir/STU3/valueset-allergyintolerance-code.html

type SubstanceCode = string // http://hl7.org/implement/standards/fhir/STU3/valueset-substance-code.html

type ClinicalFindingCode = int // http://hl7.org/implement/standards/fhir/STU3/valueset-clinical-findings.html

type AllergyIntoleranceSeverity = string

const (
	AllergyIntoleranceSeverityMild     AllergyIntoleranceSeverity = "mild"
	AllergyIntoleranceSeverityModerate AllergyIntoleranceSeverity = "moderate"
	AllergyIntoleranceSeveritySevere   AllergyIntoleranceSeverity = "severe"
)

type RouteCode = int // http://hl7.org/implement/standards/fhir/STU3/valueset-route-codes.html

type ConditionClinicalStatus = string

const (
	ConditionClinicalStatusActive     ConditionClinicalStatus = "active"
	ConditionClinicalStatusRecurrence ConditionClinicalStatus = "recurrence"
	ConditionClinicalStatusInactive   ConditionClinicalStatus = "inactive"
	ConditionClinicalStatusRemissions ConditionClinicalStatus = "remission"
	ConditionClinicalStatusResolved   ConditionClinicalStatus = "resolved"
)

type ConditionVerificationStatus = string

const (
	ConditionVerificationStatusProvisional    ConditionVerificationStatus = "provisional"
	ConditionVerificationStatusDifferential   ConditionVerificationStatus = "differential"
	ConditionVerificationStatusConfirmed      ConditionVerificationStatus = "confirmed"
	ConditionVerificationStatusRefuted        ConditionVerificationStatus = "refuted"
	ConditionVerificationStatusEnteredInError ConditionVerificationStatus = "entered-in-error"
	ConditionVerificationStatusUnknown        ConditionVerificationStatus = "unknown"
)

type ConditionCategory = string

const (
	ConditionCategoryProblemListItem    ConditionCategory = "problem-list-item"
	ConditionCategoryEncounterDiagnosis ConditionCategory = "encounter-diagnosis"
)

type ConditionSeverity = int

const (
	ConditionSeveritySevere   ConditionSeverity = 24484000
	ConditionSeverityModerate ConditionSeverity = 6736007
	ConditionSeverityMild     ConditionSeverity = 255604002
)

type BodyStructure = int // http://hl7.org/implement/standards/fhir/STU3/valueset-body-site.html

type ConditionStageCode = int // http://hl7.org/implement/standards/fhir/STU3/valueset-condition-stage.html

type ManifestationAndSymptomCode = int // http://hl7.org/implement/standards/fhir/STU3/valueset-manifestation-or-symptom.html

type EventStatus = string

const (
	EventStatusPreparation    EventStatus = "preparation"
	EventStatusInProgress     EventStatus = "in-progress"
	EventStatusSuspended      EventStatus = "suspended"
	EventStatusAborted        EventStatus = "aborted"
	EventStatusCompleted      EventStatus = "completed"
	EventStatusEnteredInError             = "entered-in-error"
	EventStatusUnknown        EventStatus = "unknown"
)

type ProcedureNotPerformedReason = int // http://hl7.org/implement/standards/fhir/STU3/valueset-procedure-not-performed-reason.html

type ProcedureCategory = int // http://hl7.org/implement/standards/fhir/STU3/valueset-procedure-category.html

type ProcedurePerformerRoleCode = int // http://hl7.org/implement/standards/fhir/STU3/valueset-performer-role.html

type ProcedureReasonCode = int // http://hl7.org/implement/standards/fhir/STU3/valueset-procedure-reason.html

type ProcedureOutcomeCode = int // http://hl7.org/implement/standards/fhir/STU3/valueset-procedure-outcome.html
const (
	ProcedureOutcomeCodeSuccessful          ProcedureOutcomeCode = 385669000
	ProcedureOutcomeCodeUnsuccessful        ProcedureOutcomeCode = 385671000
	ProcedureOutcomeCodePartiallySuccessful ProcedureOutcomeCode = 385670004
)

type ConditionCode = int // http://hl7.org/implement/standards/fhir/STU3/valueset-condition-code.html

type ProcedureFollowUpCode = int // http://hl7.org/implement/standards/fhir/STU3/valueset-procedure-followup.html

type FamilyHistoryStatusCode = string

const (
	FamilyHistoryStatusCodePartial        FamilyHistoryStatusCode = "partial"
	FamilyHistoryStatusCodeCompleted      FamilyHistoryStatusCode = "completed"
	FamilyHistoryStatusCodeEnteredInError FamilyHistoryStatusCode = "entered-in-error"
	FamilyHistoryStatusCodeHealthUnknown  FamilyHistoryStatusCode = "health-unknown"
)

type FamilyHistoryNotDoneReason = string

const (
	FamilyHistoryNotDoneReasonSubjectUnknown FamilyHistoryNotDoneReason = "subject-unknown"
	FamilyHistoryNotDoneReasonWithheld       FamilyHistoryNotDoneReason = "withheld"
	FamilyHistoryNotDoneReasonUnableToObtain FamilyHistoryNotDoneReason = "unable-to-obtain"
	FamilyHistoryNotDoneReasonDeferred       FamilyHistoryNotDoneReason = "deferred"
)

type FamilyMemberCode = string // http://hl7.org/implement/standards/fhir/STU3/v3/FamilyMember/vs.html

type ConditionOutcomeCode = int // http://hl7.org/implement/standards/fhir/STU3/valueset-condition-outcome.html

type ClinicalImpressionStatus = string

const (
	ClinicalImpressionStatusDraft          ClinicalImpressionStatus = "draft"
	ClinicalImpressionStatusCompleted      ClinicalImpressionStatus = "completed"
	ClinicalImpressionStatusEnteredInError ClinicalImpressionStatus = "entered-in-error"
)

type InvestigationTypeCode = int // http://hl7.org/implement/standards/fhir/STU3/valueset-investigation-sets.html
const (
	InvestigationTypeExaminationOrSigns = 271336007
	InvestigationTypeHistoryOrSymptoms  = 160237006
)

type ClinicalImpressionPrognosis = int // http://hl7.org/implement/standards/fhir/STU3/valueset-clinicalimpression-prognosis.html

type BodySiteLocationQualifierCode = int // http://hl7.org/implement/standards/fhir/STU3/valueset-bodysite-relative-location.html

type ObservationStatus = string

const (
	ObservationStatusRegistered  ObservationStatus = "registered"
	ObservationStatusPreliminary ObservationStatus = "preliminary"
	ObservationStatusFinal       ObservationStatus = "final"
	ObservationStatusAmended     ObservationStatus = "amended"
)

type ObservationCategoryCode = string // http://hl7.org/implement/standards/fhir/STU3/valueset-observation-category.html

type ObservationCode = string // http://hl7.org/implement/standards/fhir/STU3/valueset-observation-codes.html

type ObservationMethodCode = string // http://hl7.org/implement/standards/fhir/STU3/valueset-observation-methods.html

type CarePlanStatus = string

const (
	CarePlanStatusDraft          CarePlanStatus = "draft"
	CarePlanStatusActive         CarePlanStatus = "active"
	CarePlanStatusSuspended      CarePlanStatus = "suspended"
	CarePlanStatusCompleted      CarePlanStatus = "completed"
	CarePlanStatusEnteredInError CarePlanStatus = "entered-in-error"
	CarePlanStatusCancelled      CarePlanStatus = "cancelled"
	CarePlanStatusUnknown        CarePlanStatus = "unkonown"
)

type CarePlanIntent = string

const (
	CarePlanIntentProposal CarePlanIntent = "proposal"
	CarePlanIntentPlan     CarePlanIntent = "plan"
	CarePlanIntentOrder    CarePlanIntent = "order"
	CarePlanIntentOption   CarePlanIntent = "option"
)

type CarePlanCategoryCode = int // http://hl7.org/implement/standards/fhir/STU3/valueset-care-plan-category.html

type GoalStatus = string

const (
	GoalStatusProposed       GoalStatus = "proposed"
	GoalStatusAccepted       GoalStatus = "accepted"
	GoalStatusPlanned        GoalStatus = "planned"
	GoalStatusInProgress     GoalStatus = "in-progress"
	GoalStatusOnTarget       GoalStatus = "on-target"
	GoalStatusAheadOfTarget  GoalStatus = "ahead-of-target"
	GoalStatusBehindTarget   GoalStatus = "behind-target"
	GoalStatusSustaining     GoalStatus = "sustaining"
	GoalStatusAchieved       GoalStatus = "achieved"
	GoalStatusOnHold         GoalStatus = "on-hold"
	GoalStatusCancelled      GoalStatus = "cancelled"
	GoalStatusEnteredInError GoalStatus = "entered-in-error"
	GoalStatusRejected       GoalStatus = "rejected"
)

type GoalCategory = string

const (
	GoalCategoryDietary       GoalCategory = "dietary"
	GoalCategorySafety        GoalCategory = "safety"
	GoalCategoryBehavioral    GoalCategory = "behavioral"
	GoalCategoryNursing       GoalCategory = "nursing"
	GoalCategoryPhysiotherapy GoalCategory = "physiotherapy"
)

type GoalPriority = string

const (
	GoalPriorityHigh   GoalPriority = "high-priority"
	GoalPriorityMedium GoalPriority = "medium-priority"
	GoalPriorityLow    GoalPriority = "low-priority"
)

type CarePlanActivityDetailCategory = string

const (
	CarePlanActivityDetailCategoryDiet        CarePlanActivityDetailCategory = "diet"
	CarePlanActivityDetailCategoryDrug        CarePlanActivityDetailCategory = "drug"
	CarePlanActivityDetailCategoryEncounter   CarePlanActivityDetailCategory = "encounter"
	CarePlanActivityDetailCategoryObservation CarePlanActivityDetailCategory = "observation"
	CarePlanActivityDetailCategoryProcedure   CarePlanActivityDetailCategory = "procedure"
	CarePlanActivityDetailCategorySupply      CarePlanActivityDetailCategory = "supply"
	CarePlanActivityDetailCategoryOther       CarePlanActivityDetailCategory = "other"
)

type CarePlanActivityCode = int // http://hl7.org/implement/standards/fhir/STU3/valueset-care-plan-activity.html

type ActivityReasonCode = int // http://hl7.org/implement/standards/fhir/STU3/valueset-activity-reason.html

type CarePlanActivityStatus = string

const (
	CarePlanActivityStatusNotStarted CarePlanActivityStatus = "not-started"
	CarePlanActivityStatusScheduled  CarePlanActivityStatus = "scheduled"
	CarePlanActivityStatusInProgress CarePlanActivityStatus = "in-progress"
	CarePlanActivityStatusOnHold     CarePlanActivityStatus = "on-hold"
	CarePlanActivityStatusCompleted  CarePlanActivityStatus = "completed"
	CarePlanActivityStatusCancelled  CarePlanActivityStatus = "cancelled"
	CarePlanActivityStatusUnknown    CarePlanActivityStatus = "unknown"
)
