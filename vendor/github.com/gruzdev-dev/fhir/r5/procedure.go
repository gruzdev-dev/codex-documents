package models

import (
	"encoding/json"
	"fmt"
)

// An action that is or was performed on or for a patient, practitioner, device, organization, or location. For example, this can be a physical intervention on a patient like an operation, or less invasive like long term services, counseling, or hypnotherapy.  This can be a quality or safety inspection for a location, organization, or device.  This can be an accreditation procedure on a practitioner for licensing.
type Procedure struct {
	Id                 *string                `json:"id,omitempty" bson:"id,omitempty"`                                   // Logical id of this artifact
	Meta               *Meta                  `json:"meta,omitempty" bson:"meta,omitempty"`                               // Metadata about the resource
	ImplicitRules      *string                `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`            // A set of rules under which this content was created
	Language           *string                `json:"language,omitempty" bson:"language,omitempty"`                       // Language of the resource content
	Text               *Narrative             `json:"text,omitempty" bson:"text,omitempty"`                               // Text summary of the resource, for human interpretation
	Contained          []json.RawMessage      `json:"contained,omitempty" bson:"contained,omitempty"`                     // Contained, inline Resources
	Identifier         []Identifier           `json:"identifier,omitempty" bson:"identifier,omitempty"`                   // External Identifiers for this procedure
	BasedOn            []Reference            `json:"basedOn,omitempty" bson:"based_on,omitempty"`                        // A request for this procedure
	PartOf             []Reference            `json:"partOf,omitempty" bson:"part_of,omitempty"`                          // Part of referenced event
	Status             string                 `json:"status" bson:"status"`                                               // preparation | in-progress | not-done | on-hold | stopped | completed | entered-in-error | unknown
	StatusReason       *CodeableConcept       `json:"statusReason,omitempty" bson:"status_reason,omitempty"`              // Reason for current status
	Category           []CodeableConcept      `json:"category,omitempty" bson:"category,omitempty"`                       // Classification of the procedure
	Code               *CodeableConcept       `json:"code,omitempty" bson:"code,omitempty"`                               // Identification of the procedure
	Subject            *Reference             `json:"subject" bson:"subject"`                                             // Individual or entity the procedure was performed on
	Focus              *Reference             `json:"focus,omitempty" bson:"focus,omitempty"`                             // Who is the target of the procedure when it is not the subject of record only
	Encounter          *Reference             `json:"encounter,omitempty" bson:"encounter,omitempty"`                     // The Encounter during which this Procedure was created
	OccurrenceDateTime *string                `json:"occurrenceDateTime,omitempty" bson:"occurrence_date_time,omitempty"` // When the procedure occurred or is occurring
	OccurrencePeriod   *Period                `json:"occurrencePeriod,omitempty" bson:"occurrence_period,omitempty"`      // When the procedure occurred or is occurring
	OccurrenceString   *string                `json:"occurrenceString,omitempty" bson:"occurrence_string,omitempty"`      // When the procedure occurred or is occurring
	OccurrenceAge      *Age                   `json:"occurrenceAge,omitempty" bson:"occurrence_age,omitempty"`            // When the procedure occurred or is occurring
	OccurrenceRange    *Range                 `json:"occurrenceRange,omitempty" bson:"occurrence_range,omitempty"`        // When the procedure occurred or is occurring
	OccurrenceTiming   *Timing                `json:"occurrenceTiming,omitempty" bson:"occurrence_timing,omitempty"`      // When the procedure occurred or is occurring
	Recorded           *string                `json:"recorded,omitempty" bson:"recorded,omitempty"`                       // When the procedure was first captured in the subject's record
	Recorder           *Reference             `json:"recorder,omitempty" bson:"recorder,omitempty"`                       // Who recorded the procedure
	ReportedBoolean    *bool                  `json:"reportedBoolean,omitempty" bson:"reported_boolean,omitempty"`        // Reported rather than primary record
	ReportedReference  *Reference             `json:"reportedReference,omitempty" bson:"reported_reference,omitempty"`    // Reported rather than primary record
	Performer          []ProcedurePerformer   `json:"performer,omitempty" bson:"performer,omitempty"`                     // Who performed the procedure and what they did
	Location           *Reference             `json:"location,omitempty" bson:"location,omitempty"`                       // Where the procedure happened
	Reason             []CodeableReference    `json:"reason,omitempty" bson:"reason,omitempty"`                           // The justification that the procedure was performed
	BodySite           []CodeableConcept      `json:"bodySite,omitempty" bson:"body_site,omitempty"`                      // Target body sites
	BodyStructure      *Reference             `json:"bodyStructure,omitempty" bson:"body_structure,omitempty"`            // Target body structure
	Outcome            []CodeableReference    `json:"outcome,omitempty" bson:"outcome,omitempty"`                         // The result of procedure
	Report             []Reference            `json:"report,omitempty" bson:"report,omitempty"`                           // Any report resulting from the procedure
	Complication       []CodeableReference    `json:"complication,omitempty" bson:"complication,omitempty"`               // Complication following the procedure
	FollowUp           []CodeableReference    `json:"followUp,omitempty" bson:"follow_up,omitempty"`                      // Instructions for follow up
	Note               []Annotation           `json:"note,omitempty" bson:"note,omitempty"`                               // Additional information about the procedure
	FocalDevice        []ProcedureFocalDevice `json:"focalDevice,omitempty" bson:"focal_device,omitempty"`                // Manipulated, implanted, or removed device
	Used               []CodeableReference    `json:"used,omitempty" bson:"used,omitempty"`                               // Items used during procedure
	SupportingInfo     []Reference            `json:"supportingInfo,omitempty" bson:"supporting_info,omitempty"`          // Extra information relevant to the procedure
}

func (r *Procedure) Validate() error {
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
	for i, item := range r.BasedOn {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("BasedOn[%d]: %w", i, err)
		}
	}
	for i, item := range r.PartOf {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("PartOf[%d]: %w", i, err)
		}
	}
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.StatusReason != nil {
		if err := r.StatusReason.Validate(); err != nil {
			return fmt.Errorf("StatusReason: %w", err)
		}
	}
	for i, item := range r.Category {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Category[%d]: %w", i, err)
		}
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.Subject == nil {
		return fmt.Errorf("field 'Subject' is required")
	}
	if r.Subject != nil {
		if err := r.Subject.Validate(); err != nil {
			return fmt.Errorf("Subject: %w", err)
		}
	}
	if r.Focus != nil {
		if err := r.Focus.Validate(); err != nil {
			return fmt.Errorf("Focus: %w", err)
		}
	}
	if r.Encounter != nil {
		if err := r.Encounter.Validate(); err != nil {
			return fmt.Errorf("Encounter: %w", err)
		}
	}
	if r.OccurrencePeriod != nil {
		if err := r.OccurrencePeriod.Validate(); err != nil {
			return fmt.Errorf("OccurrencePeriod: %w", err)
		}
	}
	if r.OccurrenceAge != nil {
		if err := r.OccurrenceAge.Validate(); err != nil {
			return fmt.Errorf("OccurrenceAge: %w", err)
		}
	}
	if r.OccurrenceRange != nil {
		if err := r.OccurrenceRange.Validate(); err != nil {
			return fmt.Errorf("OccurrenceRange: %w", err)
		}
	}
	if r.OccurrenceTiming != nil {
		if err := r.OccurrenceTiming.Validate(); err != nil {
			return fmt.Errorf("OccurrenceTiming: %w", err)
		}
	}
	if r.Recorder != nil {
		if err := r.Recorder.Validate(); err != nil {
			return fmt.Errorf("Recorder: %w", err)
		}
	}
	if r.ReportedReference != nil {
		if err := r.ReportedReference.Validate(); err != nil {
			return fmt.Errorf("ReportedReference: %w", err)
		}
	}
	for i, item := range r.Performer {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Performer[%d]: %w", i, err)
		}
	}
	if r.Location != nil {
		if err := r.Location.Validate(); err != nil {
			return fmt.Errorf("Location: %w", err)
		}
	}
	for i, item := range r.Reason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reason[%d]: %w", i, err)
		}
	}
	for i, item := range r.BodySite {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("BodySite[%d]: %w", i, err)
		}
	}
	if r.BodyStructure != nil {
		if err := r.BodyStructure.Validate(); err != nil {
			return fmt.Errorf("BodyStructure: %w", err)
		}
	}
	for i, item := range r.Outcome {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Outcome[%d]: %w", i, err)
		}
	}
	for i, item := range r.Report {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Report[%d]: %w", i, err)
		}
	}
	for i, item := range r.Complication {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Complication[%d]: %w", i, err)
		}
	}
	for i, item := range r.FollowUp {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("FollowUp[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	for i, item := range r.FocalDevice {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("FocalDevice[%d]: %w", i, err)
		}
	}
	for i, item := range r.Used {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Used[%d]: %w", i, err)
		}
	}
	for i, item := range r.SupportingInfo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SupportingInfo[%d]: %w", i, err)
		}
	}
	return nil
}

type ProcedurePerformer struct {
	Id         *string          `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Function   *CodeableConcept `json:"function,omitempty" bson:"function,omitempty"`       // Type of performance
	Actor      *Reference       `json:"actor" bson:"actor"`                                 // Who performed the procedure
	OnBehalfOf *Reference       `json:"onBehalfOf,omitempty" bson:"on_behalf_of,omitempty"` // Organization the device or practitioner was acting for
	Period     *Period          `json:"period,omitempty" bson:"period,omitempty"`           // When the performer performed the procedure
}

func (r *ProcedurePerformer) Validate() error {
	if r.Function != nil {
		if err := r.Function.Validate(); err != nil {
			return fmt.Errorf("Function: %w", err)
		}
	}
	if r.Actor == nil {
		return fmt.Errorf("field 'Actor' is required")
	}
	if r.Actor != nil {
		if err := r.Actor.Validate(); err != nil {
			return fmt.Errorf("Actor: %w", err)
		}
	}
	if r.OnBehalfOf != nil {
		if err := r.OnBehalfOf.Validate(); err != nil {
			return fmt.Errorf("OnBehalfOf: %w", err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	return nil
}

type ProcedureFocalDevice struct {
	Id          *string          `json:"id,omitempty" bson:"id,omitempty"`         // Unique id for inter-element referencing
	Action      *CodeableConcept `json:"action,omitempty" bson:"action,omitempty"` // Kind of change to device
	Manipulated *Reference       `json:"manipulated" bson:"manipulated"`           // Device that was changed
}

func (r *ProcedureFocalDevice) Validate() error {
	if r.Action != nil {
		if err := r.Action.Validate(); err != nil {
			return fmt.Errorf("Action: %w", err)
		}
	}
	if r.Manipulated == nil {
		return fmt.Errorf("field 'Manipulated' is required")
	}
	if r.Manipulated != nil {
		if err := r.Manipulated.Validate(); err != nil {
			return fmt.Errorf("Manipulated: %w", err)
		}
	}
	return nil
}
