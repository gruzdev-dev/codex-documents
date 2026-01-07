package models

import (
	"encoding/json"
	"fmt"
)

// Describes the event of a patient consuming or otherwise being administered a medication.  This may be as simple as swallowing a tablet or it may be a long running infusion. Related resources tie this event to the authorizing prescription, and the specific encounter between patient and health care practitioner. This event can also be used to record waste using a status of not-done and the appropriate statusReason.
type MedicationAdministration struct {
	ResourceType          string                              `json:"resourceType" bson:"resource_type"`                                       // Type of resource
	Id                    *string                             `json:"id,omitempty" bson:"id,omitempty"`                                        // Logical id of this artifact
	Meta                  *Meta                               `json:"meta,omitempty" bson:"meta,omitempty"`                                    // Metadata about the resource
	ImplicitRules         *string                             `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                 // A set of rules under which this content was created
	Language              *string                             `json:"language,omitempty" bson:"language,omitempty"`                            // Language of the resource content
	Text                  *Narrative                          `json:"text,omitempty" bson:"text,omitempty"`                                    // Text summary of the resource, for human interpretation
	Contained             []json.RawMessage                   `json:"contained,omitempty" bson:"contained,omitempty"`                          // Contained, inline Resources
	Identifier            []Identifier                        `json:"identifier,omitempty" bson:"identifier,omitempty"`                        // External identifier
	BasedOn               []Reference                         `json:"basedOn,omitempty" bson:"based_on,omitempty"`                             // Plan this is fulfilled by this administration
	PartOf                []Reference                         `json:"partOf,omitempty" bson:"part_of,omitempty"`                               // Part of referenced event
	Status                string                              `json:"status" bson:"status"`                                                    // in-progress | not-done | on-hold | completed | entered-in-error | stopped | unknown
	StatusReason          []CodeableConcept                   `json:"statusReason,omitempty" bson:"status_reason,omitempty"`                   // Reason status of the administration changed
	Category              []CodeableConcept                   `json:"category,omitempty" bson:"category,omitempty"`                            // Type of medication administration
	Medication            *CodeableReference                  `json:"medication" bson:"medication"`                                            // What was administered
	Subject               *Reference                          `json:"subject" bson:"subject"`                                                  // Who received medication
	Encounter             *Reference                          `json:"encounter,omitempty" bson:"encounter,omitempty"`                          // Encounter administered as part of
	SupportingInformation []Reference                         `json:"supportingInformation,omitempty" bson:"supporting_information,omitempty"` // Additional information to support administration
	OccurrenceDateTime    *string                             `json:"occurrenceDateTime" bson:"occurrence_date_time"`                          // Specific date/time or interval of time during which the administration took place (or did not take place)
	OccurrencePeriod      *Period                             `json:"occurrencePeriod" bson:"occurrence_period"`                               // Specific date/time or interval of time during which the administration took place (or did not take place)
	OccurrenceTiming      *Timing                             `json:"occurrenceTiming" bson:"occurrence_timing"`                               // Specific date/time or interval of time during which the administration took place (or did not take place)
	Recorded              *string                             `json:"recorded,omitempty" bson:"recorded,omitempty"`                            // When the MedicationAdministration was first captured in the subject's record
	IsSubPotent           *bool                               `json:"isSubPotent,omitempty" bson:"is_sub_potent,omitempty"`                    // An indication that the full ordered dose was not administered
	SubPotentReason       []CodeableConcept                   `json:"subPotentReason,omitempty" bson:"sub_potent_reason,omitempty"`            // Reason full dose was not administered
	Performer             []MedicationAdministrationPerformer `json:"performer,omitempty" bson:"performer,omitempty"`                          // Who or what performed the medication administration and what type of performance they did
	Reason                []CodeableReference                 `json:"reason,omitempty" bson:"reason,omitempty"`                                // Reason that supports why the medication was administered
	Request               *Reference                          `json:"request,omitempty" bson:"request,omitempty"`                              // Request administration performed against
	Device                []CodeableReference                 `json:"device,omitempty" bson:"device,omitempty"`                                // Device used to administer
	Note                  []Annotation                        `json:"note,omitempty" bson:"note,omitempty"`                                    // Information about the administration
	Dosage                *MedicationAdministrationDosage     `json:"dosage,omitempty" bson:"dosage,omitempty"`                                // Details of how medication was taken
	EventHistory          []Reference                         `json:"eventHistory,omitempty" bson:"event_history,omitempty"`                   // A list of events of interest in the lifecycle
}

func (r *MedicationAdministration) Validate() error {
	if r.ResourceType != "MedicationAdministration" {
		return fmt.Errorf("invalid resourceType: expected 'MedicationAdministration', got '%s'", r.ResourceType)
	}
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
	for i, item := range r.StatusReason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("StatusReason[%d]: %w", i, err)
		}
	}
	for i, item := range r.Category {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Category[%d]: %w", i, err)
		}
	}
	if r.Medication == nil {
		return fmt.Errorf("field 'Medication' is required")
	}
	if r.Medication != nil {
		if err := r.Medication.Validate(); err != nil {
			return fmt.Errorf("Medication: %w", err)
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
	if r.Encounter != nil {
		if err := r.Encounter.Validate(); err != nil {
			return fmt.Errorf("Encounter: %w", err)
		}
	}
	for i, item := range r.SupportingInformation {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SupportingInformation[%d]: %w", i, err)
		}
	}
	if r.OccurrenceDateTime == nil {
		return fmt.Errorf("field 'OccurrenceDateTime' is required")
	}
	if r.OccurrencePeriod == nil {
		return fmt.Errorf("field 'OccurrencePeriod' is required")
	}
	if r.OccurrencePeriod != nil {
		if err := r.OccurrencePeriod.Validate(); err != nil {
			return fmt.Errorf("OccurrencePeriod: %w", err)
		}
	}
	if r.OccurrenceTiming == nil {
		return fmt.Errorf("field 'OccurrenceTiming' is required")
	}
	if r.OccurrenceTiming != nil {
		if err := r.OccurrenceTiming.Validate(); err != nil {
			return fmt.Errorf("OccurrenceTiming: %w", err)
		}
	}
	for i, item := range r.SubPotentReason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SubPotentReason[%d]: %w", i, err)
		}
	}
	for i, item := range r.Performer {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Performer[%d]: %w", i, err)
		}
	}
	for i, item := range r.Reason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reason[%d]: %w", i, err)
		}
	}
	if r.Request != nil {
		if err := r.Request.Validate(); err != nil {
			return fmt.Errorf("Request: %w", err)
		}
	}
	for i, item := range r.Device {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Device[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	if r.Dosage != nil {
		if err := r.Dosage.Validate(); err != nil {
			return fmt.Errorf("Dosage: %w", err)
		}
	}
	for i, item := range r.EventHistory {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("EventHistory[%d]: %w", i, err)
		}
	}
	return nil
}

type MedicationAdministrationPerformer struct {
	Id       *string            `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Function *CodeableConcept   `json:"function,omitempty" bson:"function,omitempty"` // Type of performance
	Actor    *CodeableReference `json:"actor" bson:"actor"`                           // Who or what performed the medication administration
}

func (r *MedicationAdministrationPerformer) Validate() error {
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
	return nil
}

type MedicationAdministrationDosage struct {
	Id           *string          `json:"id,omitempty" bson:"id,omitempty"`                      // Unique id for inter-element referencing
	Text         *string          `json:"text,omitempty" bson:"text,omitempty"`                  // Free text dosage instructions e.g. SIG
	Site         *CodeableConcept `json:"site,omitempty" bson:"site,omitempty"`                  // Body site administered to
	Route        *CodeableConcept `json:"route,omitempty" bson:"route,omitempty"`                // Path of substance into body
	Method       *CodeableConcept `json:"method,omitempty" bson:"method,omitempty"`              // How drug was administered
	Dose         *Quantity        `json:"dose,omitempty" bson:"dose,omitempty"`                  // Amount of medication per dose
	RateRatio    *Ratio           `json:"rateRatio,omitempty" bson:"rate_ratio,omitempty"`       // Dose quantity per unit of time
	RateQuantity *Quantity        `json:"rateQuantity,omitempty" bson:"rate_quantity,omitempty"` // Dose quantity per unit of time
}

func (r *MedicationAdministrationDosage) Validate() error {
	if r.Site != nil {
		if err := r.Site.Validate(); err != nil {
			return fmt.Errorf("Site: %w", err)
		}
	}
	if r.Route != nil {
		if err := r.Route.Validate(); err != nil {
			return fmt.Errorf("Route: %w", err)
		}
	}
	if r.Method != nil {
		if err := r.Method.Validate(); err != nil {
			return fmt.Errorf("Method: %w", err)
		}
	}
	if r.Dose != nil {
		if err := r.Dose.Validate(); err != nil {
			return fmt.Errorf("Dose: %w", err)
		}
	}
	if r.RateRatio != nil {
		if err := r.RateRatio.Validate(); err != nil {
			return fmt.Errorf("RateRatio: %w", err)
		}
	}
	if r.RateQuantity != nil {
		if err := r.RateQuantity.Validate(); err != nil {
			return fmt.Errorf("RateQuantity: %w", err)
		}
	}
	return nil
}
