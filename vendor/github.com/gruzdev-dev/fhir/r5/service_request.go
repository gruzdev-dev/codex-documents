package models

import (
	"encoding/json"
	"fmt"
)

// A record of a request for service such as diagnostic investigations, treatments, or operations to be performed.
type ServiceRequest struct {
	Id                 *string                            `json:"id,omitempty" bson:"id,omitempty"`                                   // Logical id of this artifact
	Meta               *Meta                              `json:"meta,omitempty" bson:"meta,omitempty"`                               // Metadata about the resource
	ImplicitRules      *string                            `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`            // A set of rules under which this content was created
	Language           *string                            `json:"language,omitempty" bson:"language,omitempty"`                       // Language of the resource content
	Text               *Narrative                         `json:"text,omitempty" bson:"text,omitempty"`                               // Text summary of the resource, for human interpretation
	Contained          []json.RawMessage                  `json:"contained,omitempty" bson:"contained,omitempty"`                     // Contained, inline Resources
	Identifier         []Identifier                       `json:"identifier,omitempty" bson:"identifier,omitempty"`                   // Identifiers assigned to this order
	BasedOn            []Reference                        `json:"basedOn,omitempty" bson:"based_on,omitempty"`                        // What request fulfills
	Replaces           []Reference                        `json:"replaces,omitempty" bson:"replaces,omitempty"`                       // What request replaces
	Requisition        *Identifier                        `json:"requisition,omitempty" bson:"requisition,omitempty"`                 // Composite Request ID
	Status             string                             `json:"status" bson:"status"`                                               // draft | active | on-hold | entered-in-error | ended | completed | revoked | unknown
	StatusReason       []CodeableConcept                  `json:"statusReason,omitempty" bson:"status_reason,omitempty"`              // Reason for current status
	Intent             string                             `json:"intent" bson:"intent"`                                               // proposal | solicit-offer | offer-response | plan | directive | order | original-order | reflex-order | filler-order | instance-order | option
	Category           []CodeableConcept                  `json:"category,omitempty" bson:"category,omitempty"`                       // Classification of service
	Priority           *string                            `json:"priority,omitempty" bson:"priority,omitempty"`                       // routine | urgent | asap | stat
	DoNotPerform       bool                               `json:"doNotPerform,omitempty" bson:"do_not_perform,omitempty"`             // True if service/procedure should not be performed
	Code               *CodeableReference                 `json:"code,omitempty" bson:"code,omitempty"`                               // What is being requested/ordered
	OrderDetail        []ServiceRequestOrderDetail        `json:"orderDetail,omitempty" bson:"order_detail,omitempty"`                // Additional information about the request
	QuantityQuantity   *Quantity                          `json:"quantityQuantity,omitempty" bson:"quantity_quantity,omitempty"`      // Service amount
	QuantityRatio      *Ratio                             `json:"quantityRatio,omitempty" bson:"quantity_ratio,omitempty"`            // Service amount
	QuantityRange      *Range                             `json:"quantityRange,omitempty" bson:"quantity_range,omitempty"`            // Service amount
	Subject            *Reference                         `json:"subject" bson:"subject"`                                             // Individual or Entity the service is ordered for
	Focus              []Reference                        `json:"focus,omitempty" bson:"focus,omitempty"`                             // What the service request is about, when it is not about the subject of record
	Encounter          *Reference                         `json:"encounter,omitempty" bson:"encounter,omitempty"`                     // Encounter in which the request was created
	OccurrenceDateTime *string                            `json:"occurrenceDateTime,omitempty" bson:"occurrence_date_time,omitempty"` // When service should occur
	OccurrencePeriod   *Period                            `json:"occurrencePeriod,omitempty" bson:"occurrence_period,omitempty"`      // When service should occur
	OccurrenceTiming   *Timing                            `json:"occurrenceTiming,omitempty" bson:"occurrence_timing,omitempty"`      // When service should occur
	AsNeeded           bool                               `json:"asNeeded,omitempty" bson:"as_needed,omitempty"`                      // Perform the service "as needed"
	AsNeededFor        []CodeableConcept                  `json:"asNeededFor,omitempty" bson:"as_needed_for,omitempty"`               // Specified criteria for the service
	AuthoredOn         *string                            `json:"authoredOn,omitempty" bson:"authored_on,omitempty"`                  // Date request signed
	Requester          *Reference                         `json:"requester,omitempty" bson:"requester,omitempty"`                     // Who/what is requesting service
	PerformerType      *CodeableConcept                   `json:"performerType,omitempty" bson:"performer_type,omitempty"`            // Performer role
	Performer          []Reference                        `json:"performer,omitempty" bson:"performer,omitempty"`                     // Requested performer
	Location           []CodeableReference                `json:"location,omitempty" bson:"location,omitempty"`                       // Requested location
	Reason             []CodeableReference                `json:"reason,omitempty" bson:"reason,omitempty"`                           // Reason or indication for requesting the service
	Insurance          []Reference                        `json:"insurance,omitempty" bson:"insurance,omitempty"`                     // Associated insurance coverage
	SupportingInfo     []CodeableReference                `json:"supportingInfo,omitempty" bson:"supporting_info,omitempty"`          // Additional clinical information
	Specimen           []Reference                        `json:"specimen,omitempty" bson:"specimen,omitempty"`                       // Procedure Samples
	BodyStructure      *CodeableReference                 `json:"bodyStructure,omitempty" bson:"body_structure,omitempty"`            // BodyStructure-based location on the body
	Note               []Annotation                       `json:"note,omitempty" bson:"note,omitempty"`                               // Comments
	PatientInstruction []ServiceRequestPatientInstruction `json:"patientInstruction,omitempty" bson:"patient_instruction,omitempty"`  // Patient or consumer-oriented instructions
	RelevantHistory    []Reference                        `json:"relevantHistory,omitempty" bson:"relevant_history,omitempty"`        // Request provenance
}

func (r *ServiceRequest) Validate() error {
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
	for i, item := range r.Replaces {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Replaces[%d]: %w", i, err)
		}
	}
	if r.Requisition != nil {
		if err := r.Requisition.Validate(); err != nil {
			return fmt.Errorf("Requisition: %w", err)
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
	if r.Intent == emptyString {
		return fmt.Errorf("field 'Intent' is required")
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
	for i, item := range r.OrderDetail {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("OrderDetail[%d]: %w", i, err)
		}
	}
	if r.QuantityQuantity != nil {
		if err := r.QuantityQuantity.Validate(); err != nil {
			return fmt.Errorf("QuantityQuantity: %w", err)
		}
	}
	if r.QuantityRatio != nil {
		if err := r.QuantityRatio.Validate(); err != nil {
			return fmt.Errorf("QuantityRatio: %w", err)
		}
	}
	if r.QuantityRange != nil {
		if err := r.QuantityRange.Validate(); err != nil {
			return fmt.Errorf("QuantityRange: %w", err)
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
	for i, item := range r.Focus {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Focus[%d]: %w", i, err)
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
	if r.OccurrenceTiming != nil {
		if err := r.OccurrenceTiming.Validate(); err != nil {
			return fmt.Errorf("OccurrenceTiming: %w", err)
		}
	}
	for i, item := range r.AsNeededFor {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AsNeededFor[%d]: %w", i, err)
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
	for i, item := range r.Location {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Location[%d]: %w", i, err)
		}
	}
	for i, item := range r.Reason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reason[%d]: %w", i, err)
		}
	}
	for i, item := range r.Insurance {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Insurance[%d]: %w", i, err)
		}
	}
	for i, item := range r.SupportingInfo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SupportingInfo[%d]: %w", i, err)
		}
	}
	for i, item := range r.Specimen {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Specimen[%d]: %w", i, err)
		}
	}
	if r.BodyStructure != nil {
		if err := r.BodyStructure.Validate(); err != nil {
			return fmt.Errorf("BodyStructure: %w", err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	for i, item := range r.PatientInstruction {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("PatientInstruction[%d]: %w", i, err)
		}
	}
	for i, item := range r.RelevantHistory {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RelevantHistory[%d]: %w", i, err)
		}
	}
	return nil
}

type ServiceRequestOrderDetail struct {
	Id                            *string                              `json:"id,omitempty" bson:"id,omitempty"`                                                          // Unique id for inter-element referencing
	ParameterFocusCodeableConcept *CodeableConcept                     `json:"parameterFocusCodeableConcept,omitempty" bson:"parameter_focus_codeable_concept,omitempty"` // The context of the order details by reference
	ParameterFocusReference       *Reference                           `json:"parameterFocusReference,omitempty" bson:"parameter_focus_reference,omitempty"`              // The context of the order details by reference
	ParameterFocusCanonical       *string                              `json:"parameterFocusCanonical,omitempty" bson:"parameter_focus_canonical,omitempty"`              // The context of the order details by reference
	Parameter                     []ServiceRequestOrderDetailParameter `json:"parameter" bson:"parameter"`                                                                // The parameter details for the service being requested
}

func (r *ServiceRequestOrderDetail) Validate() error {
	if r.ParameterFocusCodeableConcept != nil {
		if err := r.ParameterFocusCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("ParameterFocusCodeableConcept: %w", err)
		}
	}
	if r.ParameterFocusReference != nil {
		if err := r.ParameterFocusReference.Validate(); err != nil {
			return fmt.Errorf("ParameterFocusReference: %w", err)
		}
	}
	if len(r.Parameter) < 1 {
		return fmt.Errorf("field 'Parameter' must have at least 1 elements")
	}
	for i, item := range r.Parameter {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Parameter[%d]: %w", i, err)
		}
	}
	return nil
}

type ServiceRequestOrderDetailParameter struct {
	Id                   *string          `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Code                 *CodeableConcept `json:"code" bson:"code"`                                   // The detail of the order being requested
	ValueQuantity        *Quantity        `json:"valueQuantity" bson:"value_quantity"`                // The value for the order detail
	ValueRatio           *Ratio           `json:"valueRatio" bson:"value_ratio"`                      // The value for the order detail
	ValueRange           *Range           `json:"valueRange" bson:"value_range"`                      // The value for the order detail
	ValueBoolean         *bool            `json:"valueBoolean" bson:"value_boolean"`                  // The value for the order detail
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept" bson:"value_codeable_concept"` // The value for the order detail
	ValueString          *string          `json:"valueString" bson:"value_string"`                    // The value for the order detail
	ValuePeriod          *Period          `json:"valuePeriod" bson:"value_period"`                    // The value for the order detail
}

func (r *ServiceRequestOrderDetailParameter) Validate() error {
	if r.Code == nil {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.ValueQuantity == nil {
		return fmt.Errorf("field 'ValueQuantity' is required")
	}
	if r.ValueQuantity != nil {
		if err := r.ValueQuantity.Validate(); err != nil {
			return fmt.Errorf("ValueQuantity: %w", err)
		}
	}
	if r.ValueRatio == nil {
		return fmt.Errorf("field 'ValueRatio' is required")
	}
	if r.ValueRatio != nil {
		if err := r.ValueRatio.Validate(); err != nil {
			return fmt.Errorf("ValueRatio: %w", err)
		}
	}
	if r.ValueRange == nil {
		return fmt.Errorf("field 'ValueRange' is required")
	}
	if r.ValueRange != nil {
		if err := r.ValueRange.Validate(); err != nil {
			return fmt.Errorf("ValueRange: %w", err)
		}
	}
	if r.ValueBoolean == nil {
		return fmt.Errorf("field 'ValueBoolean' is required")
	}
	if r.ValueCodeableConcept == nil {
		return fmt.Errorf("field 'ValueCodeableConcept' is required")
	}
	if r.ValueCodeableConcept != nil {
		if err := r.ValueCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("ValueCodeableConcept: %w", err)
		}
	}
	if r.ValueString == nil {
		return fmt.Errorf("field 'ValueString' is required")
	}
	if r.ValuePeriod == nil {
		return fmt.Errorf("field 'ValuePeriod' is required")
	}
	if r.ValuePeriod != nil {
		if err := r.ValuePeriod.Validate(); err != nil {
			return fmt.Errorf("ValuePeriod: %w", err)
		}
	}
	return nil
}

type ServiceRequestPatientInstruction struct {
	Id                   *string    `json:"id,omitempty" bson:"id,omitempty"`                                      // Unique id for inter-element referencing
	InstructionMarkdown  *string    `json:"instructionMarkdown,omitempty" bson:"instruction_markdown,omitempty"`   // Patient or consumer-oriented instructions
	InstructionReference *Reference `json:"instructionReference,omitempty" bson:"instruction_reference,omitempty"` // Patient or consumer-oriented instructions
}

func (r *ServiceRequestPatientInstruction) Validate() error {
	if r.InstructionReference != nil {
		if err := r.InstructionReference.Validate(); err != nil {
			return fmt.Errorf("InstructionReference: %w", err)
		}
	}
	return nil
}
