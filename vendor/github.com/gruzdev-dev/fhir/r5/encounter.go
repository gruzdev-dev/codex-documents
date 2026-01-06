package models

import (
	"encoding/json"
	"fmt"
)

// An interaction between healthcare provider(s), and/or patient(s) for the purpose of providing healthcare service(s) or assessing the health status of patient(s).
type Encounter struct {
	Id                 *string                   `json:"id,omitempty" bson:"id,omitempty"`                                  // Logical id of this artifact
	Meta               *Meta                     `json:"meta,omitempty" bson:"meta,omitempty"`                              // Metadata about the resource
	ImplicitRules      *string                   `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`           // A set of rules under which this content was created
	Language           *string                   `json:"language,omitempty" bson:"language,omitempty"`                      // Language of the resource content
	Text               *Narrative                `json:"text,omitempty" bson:"text,omitempty"`                              // Text summary of the resource, for human interpretation
	Contained          []json.RawMessage         `json:"contained,omitempty" bson:"contained,omitempty"`                    // Contained, inline Resources
	Identifier         []Identifier              `json:"identifier,omitempty" bson:"identifier,omitempty"`                  // Identifier(s) by which this encounter is known
	Status             string                    `json:"status" bson:"status"`                                              // planned | in-progress | on-hold | discharged | completed | cancelled | discontinued | entered-in-error | unknown
	BusinessStatus     []EncounterBusinessStatus `json:"businessStatus,omitempty" bson:"business_status,omitempty"`         // A granular, workflows specific set of statuses that apply to the encounter
	Class              []CodeableConcept         `json:"class,omitempty" bson:"class,omitempty"`                            // Classification of patient encounter context - e.g. Inpatient, outpatient
	Priority           *CodeableConcept          `json:"priority,omitempty" bson:"priority,omitempty"`                      // Indicates the urgency of the encounter
	Type               []CodeableConcept         `json:"type,omitempty" bson:"type,omitempty"`                              // Specific type of encounter (e.g. e-mail consultation, surgical day-care, ...)
	ServiceType        []CodeableReference       `json:"serviceType,omitempty" bson:"service_type,omitempty"`               // Specific type of service
	Subject            *Reference                `json:"subject,omitempty" bson:"subject,omitempty"`                        // The patient or group related to this encounter
	SubjectStatus      *CodeableConcept          `json:"subjectStatus,omitempty" bson:"subject_status,omitempty"`           // The current status of the subject in relation to the Encounter
	EpisodeOfCare      []Reference               `json:"episodeOfCare,omitempty" bson:"episode_of_care,omitempty"`          // Episode(s) of care that this encounter should be recorded against
	BasedOn            []Reference               `json:"basedOn,omitempty" bson:"based_on,omitempty"`                       // The request that initiated this encounter
	CareTeam           []Reference               `json:"careTeam,omitempty" bson:"care_team,omitempty"`                     // The group(s) that are allocated to participate in this encounter
	PartOf             *Reference                `json:"partOf,omitempty" bson:"part_of,omitempty"`                         // Another Encounter this encounter is part of
	ServiceProvider    *Reference                `json:"serviceProvider,omitempty" bson:"service_provider,omitempty"`       // The organization (facility) responsible for this encounter
	Participant        []EncounterParticipant    `json:"participant,omitempty" bson:"participant,omitempty"`                // List of participants involved in the encounter
	Appointment        []Reference               `json:"appointment,omitempty" bson:"appointment,omitempty"`                // The appointment that scheduled this encounter
	VirtualService     []VirtualServiceDetail    `json:"virtualService,omitempty" bson:"virtual_service,omitempty"`         // Connection details of a virtual service (e.g. conference call)
	ActualPeriod       *Period                   `json:"actualPeriod,omitempty" bson:"actual_period,omitempty"`             // The actual start and end time of the encounter
	PlannedStartDate   *string                   `json:"plannedStartDate,omitempty" bson:"planned_start_date,omitempty"`    // The planned start date/time (or admission date) of the encounter
	PlannedEndDate     *string                   `json:"plannedEndDate,omitempty" bson:"planned_end_date,omitempty"`        // The planned end date/time (or discharge date) of the encounter
	Length             *Duration                 `json:"length,omitempty" bson:"length,omitempty"`                          // Actual quantity of time the encounter lasted (less time absent)
	Reason             []EncounterReason         `json:"reason,omitempty" bson:"reason,omitempty"`                          // The list of medical reasons that are expected to be addressed during the episode of care
	Diagnosis          []EncounterDiagnosis      `json:"diagnosis,omitempty" bson:"diagnosis,omitempty"`                    // The list of diagnosis relevant to this encounter
	Account            []Reference               `json:"account,omitempty" bson:"account,omitempty"`                        // The set of accounts that may be used for billing for this Encounter
	DietPreference     []CodeableConcept         `json:"dietPreference,omitempty" bson:"diet_preference,omitempty"`         // Diet preferences reported by the patient
	SpecialArrangement []CodeableConcept         `json:"specialArrangement,omitempty" bson:"special_arrangement,omitempty"` // Wheelchair, translator, stretcher, etc
	SpecialCourtesy    []CodeableConcept         `json:"specialCourtesy,omitempty" bson:"special_courtesy,omitempty"`       // Special courtesies (VIP, board member)
	Admission          *EncounterAdmission       `json:"admission,omitempty" bson:"admission,omitempty"`                    // Details about the admission to a healthcare service
	Location           []EncounterLocation       `json:"location,omitempty" bson:"location,omitempty"`                      // List of locations where the patient has been
}

func (r *Encounter) Validate() error {
	if r.Meta != nil {
		if err := r.Meta.Validate(); err != nil {
			return fmt.Errorf("Meta: %w", err)
		}
	}
	if r.Text != nil {
		if err := r.Text.Validate(); err != nil {
			return fmt.Errorf("Text: %w", err)
		}
	}
	for i, item := range r.Identifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Identifier[%d]: %w", i, err)
		}
	}
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	for i, item := range r.BusinessStatus {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("BusinessStatus[%d]: %w", i, err)
		}
	}
	for i, item := range r.Class {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Class[%d]: %w", i, err)
		}
	}
	if r.Priority != nil {
		if err := r.Priority.Validate(); err != nil {
			return fmt.Errorf("Priority: %w", err)
		}
	}
	for i, item := range r.Type {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Type[%d]: %w", i, err)
		}
	}
	for i, item := range r.ServiceType {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ServiceType[%d]: %w", i, err)
		}
	}
	if r.Subject != nil {
		if err := r.Subject.Validate(); err != nil {
			return fmt.Errorf("Subject: %w", err)
		}
	}
	if r.SubjectStatus != nil {
		if err := r.SubjectStatus.Validate(); err != nil {
			return fmt.Errorf("SubjectStatus: %w", err)
		}
	}
	for i, item := range r.EpisodeOfCare {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("EpisodeOfCare[%d]: %w", i, err)
		}
	}
	for i, item := range r.BasedOn {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("BasedOn[%d]: %w", i, err)
		}
	}
	for i, item := range r.CareTeam {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("CareTeam[%d]: %w", i, err)
		}
	}
	if r.PartOf != nil {
		if err := r.PartOf.Validate(); err != nil {
			return fmt.Errorf("PartOf: %w", err)
		}
	}
	if r.ServiceProvider != nil {
		if err := r.ServiceProvider.Validate(); err != nil {
			return fmt.Errorf("ServiceProvider: %w", err)
		}
	}
	for i, item := range r.Participant {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Participant[%d]: %w", i, err)
		}
	}
	for i, item := range r.Appointment {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Appointment[%d]: %w", i, err)
		}
	}
	for i, item := range r.VirtualService {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("VirtualService[%d]: %w", i, err)
		}
	}
	if r.ActualPeriod != nil {
		if err := r.ActualPeriod.Validate(); err != nil {
			return fmt.Errorf("ActualPeriod: %w", err)
		}
	}
	if r.Length != nil {
		if err := r.Length.Validate(); err != nil {
			return fmt.Errorf("Length: %w", err)
		}
	}
	for i, item := range r.Reason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reason[%d]: %w", i, err)
		}
	}
	for i, item := range r.Diagnosis {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Diagnosis[%d]: %w", i, err)
		}
	}
	for i, item := range r.Account {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Account[%d]: %w", i, err)
		}
	}
	for i, item := range r.DietPreference {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("DietPreference[%d]: %w", i, err)
		}
	}
	for i, item := range r.SpecialArrangement {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SpecialArrangement[%d]: %w", i, err)
		}
	}
	for i, item := range r.SpecialCourtesy {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SpecialCourtesy[%d]: %w", i, err)
		}
	}
	if r.Admission != nil {
		if err := r.Admission.Validate(); err != nil {
			return fmt.Errorf("Admission: %w", err)
		}
	}
	for i, item := range r.Location {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Location[%d]: %w", i, err)
		}
	}
	return nil
}

type EncounterAdmission struct {
	Id                     *string          `json:"id,omitempty" bson:"id,omitempty"`                                           // Unique id for inter-element referencing
	PreAdmissionIdentifier *Identifier      `json:"preAdmissionIdentifier,omitempty" bson:"pre_admission_identifier,omitempty"` // Pre-admission identifier
	Origin                 *Reference       `json:"origin,omitempty" bson:"origin,omitempty"`                                   // The location/organization from which the patient came before admission
	AdmitSource            *CodeableConcept `json:"admitSource,omitempty" bson:"admit_source,omitempty"`                        // From where patient was admitted (physician referral, transfer)
	ReAdmission            *CodeableConcept `json:"reAdmission,omitempty" bson:"re_admission,omitempty"`                        // Indicates that the patient is being re-admitted
	Destination            *Reference       `json:"destination,omitempty" bson:"destination,omitempty"`                         // Location/organization to which the patient is discharged
	DischargeDisposition   *CodeableConcept `json:"dischargeDisposition,omitempty" bson:"discharge_disposition,omitempty"`      // Category or kind of location after discharge
}

func (r *EncounterAdmission) Validate() error {
	if r.PreAdmissionIdentifier != nil {
		if err := r.PreAdmissionIdentifier.Validate(); err != nil {
			return fmt.Errorf("PreAdmissionIdentifier: %w", err)
		}
	}
	if r.Origin != nil {
		if err := r.Origin.Validate(); err != nil {
			return fmt.Errorf("Origin: %w", err)
		}
	}
	if r.AdmitSource != nil {
		if err := r.AdmitSource.Validate(); err != nil {
			return fmt.Errorf("AdmitSource: %w", err)
		}
	}
	if r.ReAdmission != nil {
		if err := r.ReAdmission.Validate(); err != nil {
			return fmt.Errorf("ReAdmission: %w", err)
		}
	}
	if r.Destination != nil {
		if err := r.Destination.Validate(); err != nil {
			return fmt.Errorf("Destination: %w", err)
		}
	}
	if r.DischargeDisposition != nil {
		if err := r.DischargeDisposition.Validate(); err != nil {
			return fmt.Errorf("DischargeDisposition: %w", err)
		}
	}
	return nil
}

type EncounterLocation struct {
	Id       *string          `json:"id,omitempty" bson:"id,omitempty"`         // Unique id for inter-element referencing
	Location *Reference       `json:"location" bson:"location"`                 // Location the encounter takes place
	Status   *string          `json:"status,omitempty" bson:"status,omitempty"` // planned | active | reserved | completed
	Form     *CodeableConcept `json:"form,omitempty" bson:"form,omitempty"`     // The physical type of the location (usually the level in the location hierarchy - bed, room, ward, virtual etc.)
	Period   *Period          `json:"period,omitempty" bson:"period,omitempty"` // Time period during which the patient was present at the location
}

func (r *EncounterLocation) Validate() error {
	if r.Location == nil {
		return fmt.Errorf("field 'Location' is required")
	}
	if r.Location != nil {
		if err := r.Location.Validate(); err != nil {
			return fmt.Errorf("Location: %w", err)
		}
	}
	if r.Form != nil {
		if err := r.Form.Validate(); err != nil {
			return fmt.Errorf("Form: %w", err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	return nil
}

type EncounterBusinessStatus struct {
	Id            *string          `json:"id,omitempty" bson:"id,omitempty"`                        // Unique id for inter-element referencing
	Code          *CodeableConcept `json:"code" bson:"code"`                                        // The current business status
	Type          *Coding          `json:"type,omitempty" bson:"type,omitempty"`                    // The kind of workflow the status is tracking
	EffectiveDate *string          `json:"effectiveDate,omitempty" bson:"effective_date,omitempty"` // When the encounter entered this business status
}

func (r *EncounterBusinessStatus) Validate() error {
	if r.Code == nil {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	return nil
}

type EncounterParticipant struct {
	Id     *string           `json:"id,omitempty" bson:"id,omitempty"`         // Unique id for inter-element referencing
	Type   []CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`     // Role of participant in encounter
	Period *Period           `json:"period,omitempty" bson:"period,omitempty"` // Period of time during the encounter that the participant participated
	Actor  *Reference        `json:"actor,omitempty" bson:"actor,omitempty"`   // The individual, device, or service participating in the encounter
}

func (r *EncounterParticipant) Validate() error {
	for i, item := range r.Type {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Type[%d]: %w", i, err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	if r.Actor != nil {
		if err := r.Actor.Validate(); err != nil {
			return fmt.Errorf("Actor: %w", err)
		}
	}
	return nil
}

type EncounterReason struct {
	Id    *string             `json:"id,omitempty" bson:"id,omitempty"`       // Unique id for inter-element referencing
	Use   []CodeableConcept   `json:"use,omitempty" bson:"use,omitempty"`     // What the reason value should be used for/as
	Value []CodeableReference `json:"value,omitempty" bson:"value,omitempty"` // Reason the encounter takes place (core or reference)
}

func (r *EncounterReason) Validate() error {
	for i, item := range r.Use {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Use[%d]: %w", i, err)
		}
	}
	for i, item := range r.Value {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Value[%d]: %w", i, err)
		}
	}
	return nil
}

type EncounterDiagnosis struct {
	Id        *string             `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Condition []CodeableReference `json:"condition,omitempty" bson:"condition,omitempty"` // The diagnosis relevant to the encounter
	Use       []CodeableConcept   `json:"use,omitempty" bson:"use,omitempty"`             // Role that this diagnosis has within the encounter (e.g. admission, billing, discharge …)
}

func (r *EncounterDiagnosis) Validate() error {
	for i, item := range r.Condition {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Condition[%d]: %w", i, err)
		}
	}
	for i, item := range r.Use {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Use[%d]: %w", i, err)
		}
	}
	return nil
}
