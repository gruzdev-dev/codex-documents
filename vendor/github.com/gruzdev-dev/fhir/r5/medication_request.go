package models

import (
	"encoding/json"
	"fmt"
)

// An order or request for both supply of the medication and the instructions for administration of the medication to a patient. The resource is called "MedicationRequest" rather than "MedicationPrescription" or "MedicationOrder" to generalize the use across inpatient and outpatient settings, including care plans, etc., and to harmonize with workflow patterns.
type MedicationRequest struct {
	ResourceType            string                            `json:"resourceType" bson:"resource_type"`                                            // Type of resource
	Id                      *string                           `json:"id,omitempty" bson:"id,omitempty"`                                             // Logical id of this artifact
	Meta                    *Meta                             `json:"meta,omitempty" bson:"meta,omitempty"`                                         // Metadata about the resource
	ImplicitRules           *string                           `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                      // A set of rules under which this content was created
	Language                *string                           `json:"language,omitempty" bson:"language,omitempty"`                                 // Language of the resource content
	Text                    *Narrative                        `json:"text,omitempty" bson:"text,omitempty"`                                         // Text summary of the resource, for human interpretation
	Contained               []json.RawMessage                 `json:"contained,omitempty" bson:"contained,omitempty"`                               // Contained, inline Resources
	Identifier              []Identifier                      `json:"identifier,omitempty" bson:"identifier,omitempty"`                             // External ids for this request
	BasedOn                 []Reference                       `json:"basedOn,omitempty" bson:"based_on,omitempty"`                                  // A plan or request that is fulfilled in whole or in part by this medication request
	PriorPrescription       *Reference                        `json:"priorPrescription,omitempty" bson:"prior_prescription,omitempty"`              // Reference to an order/prescription that is being replaced by this MedicationRequest
	GroupIdentifier         *Identifier                       `json:"groupIdentifier,omitempty" bson:"group_identifier,omitempty"`                  // Composite request this is part of
	Status                  string                            `json:"status" bson:"status"`                                                         // active | on-hold | ended | stopped | completed | cancelled | entered-in-error | draft | unknown
	StatusReason            *CodeableConcept                  `json:"statusReason,omitempty" bson:"status_reason,omitempty"`                        // Reason for current status
	StatusChanged           *string                           `json:"statusChanged,omitempty" bson:"status_changed,omitempty"`                      // When the status was changed
	Intent                  string                            `json:"intent" bson:"intent"`                                                         // proposal | plan | order | original-order | reflex-order | filler-order | instance-order | option (immutable)
	Category                []CodeableConcept                 `json:"category,omitempty" bson:"category,omitempty"`                                 // Grouping or category of medication request
	Priority                *string                           `json:"priority,omitempty" bson:"priority,omitempty"`                                 // routine | urgent | asap | stat
	DoNotPerform            *bool                             `json:"doNotPerform,omitempty" bson:"do_not_perform,omitempty"`                       // If true, indicates the provider is ordering a patient should not take the specified medication
	Medication              *CodeableReference                `json:"medication" bson:"medication"`                                                 // Medication to be taken
	Subject                 *Reference                        `json:"subject" bson:"subject"`                                                       // Individual or group for whom the medication has been requested
	InformationSource       []Reference                       `json:"informationSource,omitempty" bson:"information_source,omitempty"`              // The person or organization who provided the information about this request, if the source is someone other than the requestor
	Encounter               *Reference                        `json:"encounter,omitempty" bson:"encounter,omitempty"`                               // Encounter created as part of encounter/admission/stay
	SupportingInformation   []Reference                       `json:"supportingInformation,omitempty" bson:"supporting_information,omitempty"`      // Information to support fulfilling of the medication
	AuthoredOn              *string                           `json:"authoredOn,omitempty" bson:"authored_on,omitempty"`                            // When request was initially authored
	Requester               *Reference                        `json:"requester,omitempty" bson:"requester,omitempty"`                               // Who/What requested the Request
	IsRecordOfRequest       *bool                             `json:"isRecordOfRequest,omitempty" bson:"is_record_of_request,omitempty"`            // Whether this is record of a Medication Request or the actual request itself
	PerformerType           *CodeableConcept                  `json:"performerType,omitempty" bson:"performer_type,omitempty"`                      // Desired kind of performer of the medication administration
	Performer               []Reference                       `json:"performer,omitempty" bson:"performer,omitempty"`                               // Intended performer of administration
	Device                  []CodeableReference               `json:"device,omitempty" bson:"device,omitempty"`                                     // Intended type of device for the administration
	Recorder                *Reference                        `json:"recorder,omitempty" bson:"recorder,omitempty"`                                 // Person who entered the request
	Reason                  []CodeableReference               `json:"reason,omitempty" bson:"reason,omitempty"`                                     // Reason or indication for ordering or not ordering the medication
	CourseOfTherapyType     *CodeableConcept                  `json:"courseOfTherapyType,omitempty" bson:"course_of_therapy_type,omitempty"`        // Overall pattern of medication administration
	Insurance               []Reference                       `json:"insurance,omitempty" bson:"insurance,omitempty"`                               // Associated insurance coverage
	Note                    []Annotation                      `json:"note,omitempty" bson:"note,omitempty"`                                         // Information about the prescription
	EffectiveTimingDuration *Duration                         `json:"effectiveTimingDuration,omitempty" bson:"effective_timing_duration,omitempty"` // Period over which the medication is to be taken, can be specified as a duration or a range
	EffectiveTimingRange    *Range                            `json:"effectiveTimingRange,omitempty" bson:"effective_timing_range,omitempty"`       // Period over which the medication is to be taken, can be specified as a duration or a range
	EffectiveTimingPeriod   *Period                           `json:"effectiveTimingPeriod,omitempty" bson:"effective_timing_period,omitempty"`     // Period over which the medication is to be taken, can be specified as a duration or a range
	DosageInstruction       *DosageDetails                    `json:"dosageInstruction,omitempty" bson:"dosage_instruction,omitempty"`              // Specific instructions for how the medication should be taken
	DispenseRequest         *MedicationRequestDispenseRequest `json:"dispenseRequest,omitempty" bson:"dispense_request,omitempty"`                  // Medication supply authorization
	Substitution            *MedicationRequestSubstitution    `json:"substitution,omitempty" bson:"substitution,omitempty"`                         // Any restrictions on medication substitution
	EventHistory            []Reference                       `json:"eventHistory,omitempty" bson:"event_history,omitempty"`                        // A list of events of interest in the lifecycle
}

func (r *MedicationRequest) Validate() error {
	if r.ResourceType != "MedicationRequest" {
		return fmt.Errorf("invalid resourceType: expected 'MedicationRequest', got '%s'", r.ResourceType)
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
	if r.PriorPrescription != nil {
		if err := r.PriorPrescription.Validate(); err != nil {
			return fmt.Errorf("PriorPrescription: %w", err)
		}
	}
	if r.GroupIdentifier != nil {
		if err := r.GroupIdentifier.Validate(); err != nil {
			return fmt.Errorf("GroupIdentifier: %w", err)
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
	if r.Intent == emptyString {
		return fmt.Errorf("field 'Intent' is required")
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
	for i, item := range r.InformationSource {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("InformationSource[%d]: %w", i, err)
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
	if r.Requester != nil {
		if err := r.Requester.Validate(); err != nil {
			return fmt.Errorf("Requester: %w", err)
		}
	}
	if r.PerformerType != nil {
		if err := r.PerformerType.Validate(); err != nil {
			return fmt.Errorf("PerformerType: %w", err)
		}
	}
	for i, item := range r.Performer {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Performer[%d]: %w", i, err)
		}
	}
	for i, item := range r.Device {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Device[%d]: %w", i, err)
		}
	}
	if r.Recorder != nil {
		if err := r.Recorder.Validate(); err != nil {
			return fmt.Errorf("Recorder: %w", err)
		}
	}
	for i, item := range r.Reason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reason[%d]: %w", i, err)
		}
	}
	if r.CourseOfTherapyType != nil {
		if err := r.CourseOfTherapyType.Validate(); err != nil {
			return fmt.Errorf("CourseOfTherapyType: %w", err)
		}
	}
	for i, item := range r.Insurance {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Insurance[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	if r.EffectiveTimingDuration != nil {
		if err := r.EffectiveTimingDuration.Validate(); err != nil {
			return fmt.Errorf("EffectiveTimingDuration: %w", err)
		}
	}
	if r.EffectiveTimingRange != nil {
		if err := r.EffectiveTimingRange.Validate(); err != nil {
			return fmt.Errorf("EffectiveTimingRange: %w", err)
		}
	}
	if r.EffectiveTimingPeriod != nil {
		if err := r.EffectiveTimingPeriod.Validate(); err != nil {
			return fmt.Errorf("EffectiveTimingPeriod: %w", err)
		}
	}
	if r.DosageInstruction != nil {
		if err := r.DosageInstruction.Validate(); err != nil {
			return fmt.Errorf("DosageInstruction: %w", err)
		}
	}
	if r.DispenseRequest != nil {
		if err := r.DispenseRequest.Validate(); err != nil {
			return fmt.Errorf("DispenseRequest: %w", err)
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

type MedicationRequestSubstitution struct {
	Id                     *string          `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	AllowedBoolean         *bool            `json:"allowedBoolean" bson:"allowed_boolean"`                  // Whether substitution is allowed or not
	AllowedCodeableConcept *CodeableConcept `json:"allowedCodeableConcept" bson:"allowed_codeable_concept"` // Whether substitution is allowed or not
	Reason                 *CodeableConcept `json:"reason,omitempty" bson:"reason,omitempty"`               // Why should (not) substitution be made
}

func (r *MedicationRequestSubstitution) Validate() error {
	if r.AllowedBoolean == nil {
		return fmt.Errorf("field 'AllowedBoolean' is required")
	}
	if r.AllowedCodeableConcept == nil {
		return fmt.Errorf("field 'AllowedCodeableConcept' is required")
	}
	if r.AllowedCodeableConcept != nil {
		if err := r.AllowedCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("AllowedCodeableConcept: %w", err)
		}
	}
	if r.Reason != nil {
		if err := r.Reason.Validate(); err != nil {
			return fmt.Errorf("Reason: %w", err)
		}
	}
	return nil
}

type MedicationRequestDispenseRequest struct {
	Id                     *string                                      `json:"id,omitempty" bson:"id,omitempty"`                                            // Unique id for inter-element referencing
	InitialFill            *MedicationRequestDispenseRequestInitialFill `json:"initialFill,omitempty" bson:"initial_fill,omitempty"`                         // First fill details
	DispenseInterval       *Duration                                    `json:"dispenseInterval,omitempty" bson:"dispense_interval,omitempty"`               // Minimum period of time between dispenses
	ValidityPeriod         *Period                                      `json:"validityPeriod,omitempty" bson:"validity_period,omitempty"`                   // Time period supply is authorized for
	NumberOfRepeatsAllowed *int                                         `json:"numberOfRepeatsAllowed,omitempty" bson:"number_of_repeats_allowed,omitempty"` // Number of refills authorized
	Quantity               *Quantity                                    `json:"quantity,omitempty" bson:"quantity,omitempty"`                                // Amount of medication to supply per dispense
	ExpectedSupplyDuration *Duration                                    `json:"expectedSupplyDuration,omitempty" bson:"expected_supply_duration,omitempty"`  // Number of days supply per dispense
	Dispenser              *Reference                                   `json:"dispenser,omitempty" bson:"dispenser,omitempty"`                              // Intended performer of dispense
	DispenserInstruction   []CodeableConcept                            `json:"dispenserInstruction,omitempty" bson:"dispenser_instruction,omitempty"`       // Additional information for the dispenser
	DoseAdministrationAid  *CodeableConcept                             `json:"doseAdministrationAid,omitempty" bson:"dose_administration_aid,omitempty"`    // Type of adherence packaging to use for the dispense
	Destination            *Reference                                   `json:"destination,omitempty" bson:"destination,omitempty"`                          // Where the medication is expected to be delivered
}

func (r *MedicationRequestDispenseRequest) Validate() error {
	if r.InitialFill != nil {
		if err := r.InitialFill.Validate(); err != nil {
			return fmt.Errorf("InitialFill: %w", err)
		}
	}
	if r.DispenseInterval != nil {
		if err := r.DispenseInterval.Validate(); err != nil {
			return fmt.Errorf("DispenseInterval: %w", err)
		}
	}
	if r.ValidityPeriod != nil {
		if err := r.ValidityPeriod.Validate(); err != nil {
			return fmt.Errorf("ValidityPeriod: %w", err)
		}
	}
	if r.Quantity != nil {
		if err := r.Quantity.Validate(); err != nil {
			return fmt.Errorf("Quantity: %w", err)
		}
	}
	if r.ExpectedSupplyDuration != nil {
		if err := r.ExpectedSupplyDuration.Validate(); err != nil {
			return fmt.Errorf("ExpectedSupplyDuration: %w", err)
		}
	}
	if r.Dispenser != nil {
		if err := r.Dispenser.Validate(); err != nil {
			return fmt.Errorf("Dispenser: %w", err)
		}
	}
	for i, item := range r.DispenserInstruction {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("DispenserInstruction[%d]: %w", i, err)
		}
	}
	if r.DoseAdministrationAid != nil {
		if err := r.DoseAdministrationAid.Validate(); err != nil {
			return fmt.Errorf("DoseAdministrationAid: %w", err)
		}
	}
	if r.Destination != nil {
		if err := r.Destination.Validate(); err != nil {
			return fmt.Errorf("Destination: %w", err)
		}
	}
	return nil
}

type MedicationRequestDispenseRequestInitialFill struct {
	Id       *string   `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Quantity *Quantity `json:"quantity,omitempty" bson:"quantity,omitempty"` // First fill quantity
	Duration *Duration `json:"duration,omitempty" bson:"duration,omitempty"` // First fill duration
}

func (r *MedicationRequestDispenseRequestInitialFill) Validate() error {
	if r.Quantity != nil {
		if err := r.Quantity.Validate(); err != nil {
			return fmt.Errorf("Quantity: %w", err)
		}
	}
	if r.Duration != nil {
		if err := r.Duration.Validate(); err != nil {
			return fmt.Errorf("Duration: %w", err)
		}
	}
	return nil
}
