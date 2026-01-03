package models

import "encoding/json"

// An interaction between healthcare provider(s), and/or patient(s) for the purpose of providing healthcare service(s) or assessing the health status of patient(s).
type Encounter struct {
	Id                 *string                   `json:"id,omitempty"`                 // Logical id of this artifact
	Meta               *Meta                     `json:"meta,omitempty"`               // Metadata about the resource
	ImplicitRules      *string                   `json:"implicitRules,omitempty"`      // A set of rules under which this content was created
	Language           *string                   `json:"language,omitempty"`           // Language of the resource content
	Text               *Narrative                `json:"text,omitempty"`               // Text summary of the resource, for human interpretation
	Contained          []json.RawMessage         `json:"contained,omitempty"`          // Contained, inline Resources
	Identifier         []Identifier              `json:"identifier,omitempty"`         // Identifier(s) by which this encounter is known
	Status             string                    `json:"status,omitempty"`             // planned | in-progress | on-hold | discharged | completed | cancelled | discontinued | entered-in-error | unknown
	BusinessStatus     []EncounterBusinessStatus `json:"businessStatus,omitempty"`     // A granular, workflows specific set of statuses that apply to the encounter
	Class              []CodeableConcept         `json:"class,omitempty"`              // Classification of patient encounter context - e.g. Inpatient, outpatient
	Priority           *CodeableConcept          `json:"priority,omitempty"`           // Indicates the urgency of the encounter
	Type               []CodeableConcept         `json:"type,omitempty"`               // Specific type of encounter (e.g. e-mail consultation, surgical day-care, ...)
	ServiceType        []CodeableReference       `json:"serviceType,omitempty"`        // Specific type of service
	Subject            *Reference                `json:"subject,omitempty"`            // The patient or group related to this encounter
	SubjectStatus      *CodeableConcept          `json:"subjectStatus,omitempty"`      // The current status of the subject in relation to the Encounter
	EpisodeOfCare      []Reference               `json:"episodeOfCare,omitempty"`      // Episode(s) of care that this encounter should be recorded against
	BasedOn            []Reference               `json:"basedOn,omitempty"`            // The request that initiated this encounter
	CareTeam           []Reference               `json:"careTeam,omitempty"`           // The group(s) that are allocated to participate in this encounter
	PartOf             *Reference                `json:"partOf,omitempty"`             // Another Encounter this encounter is part of
	ServiceProvider    *Reference                `json:"serviceProvider,omitempty"`    // The organization (facility) responsible for this encounter
	Participant        []EncounterParticipant    `json:"participant,omitempty"`        // List of participants involved in the encounter
	Appointment        []Reference               `json:"appointment,omitempty"`        // The appointment that scheduled this encounter
	VirtualService     []VirtualServiceDetail    `json:"virtualService,omitempty"`     // Connection details of a virtual service (e.g. conference call)
	ActualPeriod       *Period                   `json:"actualPeriod,omitempty"`       // The actual start and end time of the encounter
	PlannedStartDate   *string                   `json:"plannedStartDate,omitempty"`   // The planned start date/time (or admission date) of the encounter
	PlannedEndDate     *string                   `json:"plannedEndDate,omitempty"`     // The planned end date/time (or discharge date) of the encounter
	Length             *Duration                 `json:"length,omitempty"`             // Actual quantity of time the encounter lasted (less time absent)
	Reason             []EncounterReason         `json:"reason,omitempty"`             // The list of medical reasons that are expected to be addressed during the episode of care
	Diagnosis          []EncounterDiagnosis      `json:"diagnosis,omitempty"`          // The list of diagnosis relevant to this encounter
	Account            []Reference               `json:"account,omitempty"`            // The set of accounts that may be used for billing for this Encounter
	DietPreference     []CodeableConcept         `json:"dietPreference,omitempty"`     // Diet preferences reported by the patient
	SpecialArrangement []CodeableConcept         `json:"specialArrangement,omitempty"` // Wheelchair, translator, stretcher, etc
	SpecialCourtesy    []CodeableConcept         `json:"specialCourtesy,omitempty"`    // Special courtesies (VIP, board member)
	Admission          *EncounterAdmission       `json:"admission,omitempty"`          // Details about the admission to a healthcare service
	Location           []EncounterLocation       `json:"location,omitempty"`           // List of locations where the patient has been
}

type EncounterAdmission struct {
	Id                     *string          `json:"id,omitempty"`                     // Unique id for inter-element referencing
	PreAdmissionIdentifier *Identifier      `json:"preAdmissionIdentifier,omitempty"` // Pre-admission identifier
	Origin                 *Reference       `json:"origin,omitempty"`                 // The location/organization from which the patient came before admission
	AdmitSource            *CodeableConcept `json:"admitSource,omitempty"`            // From where patient was admitted (physician referral, transfer)
	ReAdmission            *CodeableConcept `json:"reAdmission,omitempty"`            // Indicates that the patient is being re-admitted
	Destination            *Reference       `json:"destination,omitempty"`            // Location/organization to which the patient is discharged
	DischargeDisposition   *CodeableConcept `json:"dischargeDisposition,omitempty"`   // Category or kind of location after discharge
}

type EncounterLocation struct {
	Id       *string          `json:"id,omitempty"`       // Unique id for inter-element referencing
	Location Reference        `json:"location,omitempty"` // Location the encounter takes place
	Status   *string          `json:"status,omitempty"`   // planned | active | reserved | completed
	Form     *CodeableConcept `json:"form,omitempty"`     // The physical type of the location (usually the level in the location hierarchy - bed, room, ward, virtual etc.)
	Period   *Period          `json:"period,omitempty"`   // Time period during which the patient was present at the location
}

type EncounterBusinessStatus struct {
	Id            *string         `json:"id,omitempty"`            // Unique id for inter-element referencing
	Code          CodeableConcept `json:"code,omitempty"`          // The current business status
	Type          *Coding         `json:"type,omitempty"`          // The kind of workflow the status is tracking
	EffectiveDate *string         `json:"effectiveDate,omitempty"` // When the encounter entered this business status
}

type EncounterParticipant struct {
	Id     *string           `json:"id,omitempty"`     // Unique id for inter-element referencing
	Type   []CodeableConcept `json:"type,omitempty"`   // Role of participant in encounter
	Period *Period           `json:"period,omitempty"` // Period of time during the encounter that the participant participated
	Actor  *Reference        `json:"actor,omitempty"`  // The individual, device, or service participating in the encounter
}

type EncounterReason struct {
	Id    *string             `json:"id,omitempty"`    // Unique id for inter-element referencing
	Use   []CodeableConcept   `json:"use,omitempty"`   // What the reason value should be used for/as
	Value []CodeableReference `json:"value,omitempty"` // Reason the encounter takes place (core or reference)
}

type EncounterDiagnosis struct {
	Id        *string             `json:"id,omitempty"`        // Unique id for inter-element referencing
	Condition []CodeableReference `json:"condition,omitempty"` // The diagnosis relevant to the encounter
	Use       []CodeableConcept   `json:"use,omitempty"`       // Role that this diagnosis has within the encounter (e.g. admission, billing, discharge …)
}
