package models

import (
	"encoding/json"
	"fmt"
)

// Indicates that a medication product is to be or has been dispensed for a named person/patient.  This includes a description of the medication product (supply) provided and the instructions for administering the medication.  The medication dispense is the result of a pharmacy system responding to a medication order.
type MedicationDispense struct {
	ResourceType            string                          `json:"resourceType" bson:"resource_type"`                                           // Type of resource
	Id                      *string                         `json:"id,omitempty" bson:"id,omitempty"`                                            // Logical id of this artifact
	Meta                    *Meta                           `json:"meta,omitempty" bson:"meta,omitempty"`                                        // Metadata about the resource
	ImplicitRules           *string                         `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                     // A set of rules under which this content was created
	Language                *string                         `json:"language,omitempty" bson:"language,omitempty"`                                // Language of the resource content
	Text                    *Narrative                      `json:"text,omitempty" bson:"text,omitempty"`                                        // Text summary of the resource, for human interpretation
	Contained               []json.RawMessage               `json:"contained,omitempty" bson:"contained,omitempty"`                              // Contained, inline Resources
	Identifier              []Identifier                    `json:"identifier,omitempty" bson:"identifier,omitempty"`                            // External identifier
	BasedOn                 []Reference                     `json:"basedOn,omitempty" bson:"based_on,omitempty"`                                 // Plan that is fulfilled by this dispense
	PartOf                  []Reference                     `json:"partOf,omitempty" bson:"part_of,omitempty"`                                   // Event that dispense is part of
	Status                  string                          `json:"status" bson:"status"`                                                        // preparation | in-progress | cancelled | on-hold | completed | entered-in-error | unfulfilled | declined | unknown
	NotPerformedReason      *CodeableReference              `json:"notPerformedReason,omitempty" bson:"not_performed_reason,omitempty"`          // Why a dispense was not performed
	StatusChanged           *string                         `json:"statusChanged,omitempty" bson:"status_changed,omitempty"`                     // When the status changed
	Category                []CodeableConcept               `json:"category,omitempty" bson:"category,omitempty"`                                // Type of medication dispense
	Medication              *CodeableReference              `json:"medication" bson:"medication"`                                                // What medication was (or was intended to be) supplied
	Subject                 *Reference                      `json:"subject" bson:"subject"`                                                      // Who the dispense is for
	Encounter               *Reference                      `json:"encounter,omitempty" bson:"encounter,omitempty"`                              // Encounter associated with event
	SupportingInformation   []Reference                     `json:"supportingInformation,omitempty" bson:"supporting_information,omitempty"`     // Information that supports the dispensing of the medication
	Performer               []MedicationDispensePerformer   `json:"performer,omitempty" bson:"performer,omitempty"`                              // Who performed event
	Location                *Reference                      `json:"location,omitempty" bson:"location,omitempty"`                                // Where the dispense occurred
	AuthorizingPrescription []Reference                     `json:"authorizingPrescription,omitempty" bson:"authorizing_prescription,omitempty"` // Medication order that authorizes the dispense
	Type                    *CodeableConcept                `json:"type,omitempty" bson:"type,omitempty"`                                        // Trial fill, partial fill, emergency fill, etc
	Quantity                *Quantity                       `json:"quantity,omitempty" bson:"quantity,omitempty"`                                // Amount of medication
	DaysSupply              *Quantity                       `json:"daysSupply,omitempty" bson:"days_supply,omitempty"`                           // Amount of medication expressed as a timing amount
	FillNumber              *int                            `json:"fillNumber,omitempty" bson:"fill_number,omitempty"`                           // A number that represents the known fill this dispense represents
	Recorded                *string                         `json:"recorded,omitempty" bson:"recorded,omitempty"`                                // When the recording of the dispense started
	WhenPrepared            *string                         `json:"whenPrepared,omitempty" bson:"when_prepared,omitempty"`                       // When product was packaged and reviewed
	WhenHandedOver          *string                         `json:"whenHandedOver,omitempty" bson:"when_handed_over,omitempty"`                  // When product was given out
	Destination             *Reference                      `json:"destination,omitempty" bson:"destination,omitempty"`                          // Where the medication was/will be sent
	Receiver                []Reference                     `json:"receiver,omitempty" bson:"receiver,omitempty"`                                // Who collected the medication or where the medication was delivered
	Note                    []Annotation                    `json:"note,omitempty" bson:"note,omitempty"`                                        // Information about the dispense
	DosageInstruction       *DosageDetails                  `json:"dosageInstruction,omitempty" bson:"dosage_instruction,omitempty"`             // How the medication is to be used by the patient or administered by the caregiver
	DoseAdministrationAid   *CodeableConcept                `json:"doseAdministrationAid,omitempty" bson:"dose_administration_aid,omitempty"`    // Type of adherence packaging to use for the dispense
	Substitution            *MedicationDispenseSubstitution `json:"substitution,omitempty" bson:"substitution,omitempty"`                        // Whether a substitution was performed on the dispense
	EventHistory            []Reference                     `json:"eventHistory,omitempty" bson:"event_history,omitempty"`                       // A list of relevant lifecycle events
}

func (r *MedicationDispense) Validate() error {
	if r.ResourceType != "MedicationDispense" {
		return fmt.Errorf("invalid resourceType: expected 'MedicationDispense', got '%s'", r.ResourceType)
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
	if r.NotPerformedReason != nil {
		if err := r.NotPerformedReason.Validate(); err != nil {
			return fmt.Errorf("NotPerformedReason: %w", err)
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
	for i, item := range r.AuthorizingPrescription {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AuthorizingPrescription[%d]: %w", i, err)
		}
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Quantity != nil {
		if err := r.Quantity.Validate(); err != nil {
			return fmt.Errorf("Quantity: %w", err)
		}
	}
	if r.DaysSupply != nil {
		if err := r.DaysSupply.Validate(); err != nil {
			return fmt.Errorf("DaysSupply: %w", err)
		}
	}
	if r.Destination != nil {
		if err := r.Destination.Validate(); err != nil {
			return fmt.Errorf("Destination: %w", err)
		}
	}
	for i, item := range r.Receiver {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Receiver[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	if r.DosageInstruction != nil {
		if err := r.DosageInstruction.Validate(); err != nil {
			return fmt.Errorf("DosageInstruction: %w", err)
		}
	}
	if r.DoseAdministrationAid != nil {
		if err := r.DoseAdministrationAid.Validate(); err != nil {
			return fmt.Errorf("DoseAdministrationAid: %w", err)
		}
	}
	if r.Substitution != nil {
		if err := r.Substitution.Validate(); err != nil {
			return fmt.Errorf("Substitution: %w", err)
		}
	}
	for i, item := range r.EventHistory {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("EventHistory[%d]: %w", i, err)
		}
	}
	return nil
}

type MedicationDispensePerformer struct {
	Id       *string          `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Function *CodeableConcept `json:"function,omitempty" bson:"function,omitempty"` // Who performed the dispense and what they did
	Actor    *Reference       `json:"actor" bson:"actor"`                           // Individual who was performing
}

func (r *MedicationDispensePerformer) Validate() error {
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

type MedicationDispenseSubstitution struct {
	Id               *string           `json:"id,omitempty" bson:"id,omitempty"`                              // Unique id for inter-element referencing
	WasSubstituted   bool              `json:"wasSubstituted" bson:"was_substituted"`                         // Whether a substitution was or was not performed on the dispense
	Type             *CodeableConcept  `json:"type,omitempty" bson:"type,omitempty"`                          // Code signifying whether a different drug was dispensed from what was prescribed
	Reason           []CodeableConcept `json:"reason,omitempty" bson:"reason,omitempty"`                      // Why was substitution made
	ResponsibleParty *Reference        `json:"responsibleParty,omitempty" bson:"responsible_party,omitempty"` // Who is responsible for the substitution
}

func (r *MedicationDispenseSubstitution) Validate() error {
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	for i, item := range r.Reason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reason[%d]: %w", i, err)
		}
	}
	if r.ResponsibleParty != nil {
		if err := r.ResponsibleParty.Validate(); err != nil {
			return fmt.Errorf("ResponsibleParty: %w", err)
		}
	}
	return nil
}
