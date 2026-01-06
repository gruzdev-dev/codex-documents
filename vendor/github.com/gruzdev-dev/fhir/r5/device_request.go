package models

import (
	"encoding/json"
	"fmt"
)

// Represents a request a device to be provided to a specific patient. The device may be an implantable device to be subsequently implanted, or an external assistive device, such as a walker, to be delivered and subsequently be used.
type DeviceRequest struct {
	Id                     *string                  `json:"id,omitempty" bson:"id,omitempty"`                                   // Logical id of this artifact
	Meta                   *Meta                    `json:"meta,omitempty" bson:"meta,omitempty"`                               // Metadata about the resource
	ImplicitRules          *string                  `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`            // A set of rules under which this content was created
	Language               *string                  `json:"language,omitempty" bson:"language,omitempty"`                       // Language of the resource content
	Text                   *Narrative               `json:"text,omitempty" bson:"text,omitempty"`                               // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage        `json:"contained,omitempty" bson:"contained,omitempty"`                     // Contained, inline Resources
	Identifier             []Identifier             `json:"identifier,omitempty" bson:"identifier,omitempty"`                   // External Request identifier
	BasedOn                []Reference              `json:"basedOn,omitempty" bson:"based_on,omitempty"`                        // What request fulfills
	Replaces               []Reference              `json:"replaces,omitempty" bson:"replaces,omitempty"`                       // What request replaces
	GroupIdentifier        *Identifier              `json:"groupIdentifier,omitempty" bson:"group_identifier,omitempty"`        // Identifier of composite request
	Status                 *string                  `json:"status,omitempty" bson:"status,omitempty"`                           // draft | active | on-hold | entered-in-error | ended | completed | revoked | unknown
	Intent                 string                   `json:"intent" bson:"intent"`                                               // proposal | solicit-offer | offer-response | plan | directive | order | original-order | reflex-order | filler-order | instance-order | option
	Priority               *string                  `json:"priority,omitempty" bson:"priority,omitempty"`                       // routine | urgent | asap | stat
	DoNotPerform           bool                     `json:"doNotPerform,omitempty" bson:"do_not_perform,omitempty"`             // True if the request is to stop or not to start using the device
	ProductCodeableConcept *CodeableConcept         `json:"productCodeableConcept" bson:"product_codeable_concept"`             // Device requested
	ProductReference       *Reference               `json:"productReference" bson:"product_reference"`                          // Device requested
	ProductCanonical       *string                  `json:"productCanonical" bson:"product_canonical"`                          // Device requested
	Quantity               *int                     `json:"quantity,omitempty" bson:"quantity,omitempty"`                       // Quantity of devices to supply
	Parameter              []DeviceRequestParameter `json:"parameter,omitempty" bson:"parameter,omitempty"`                     // Device details
	Subject                *Reference               `json:"subject" bson:"subject"`                                             // Focus of request
	Encounter              *Reference               `json:"encounter,omitempty" bson:"encounter,omitempty"`                     // Encounter motivating request
	OccurrenceDateTime     *string                  `json:"occurrenceDateTime,omitempty" bson:"occurrence_date_time,omitempty"` // Desired time or schedule for use
	OccurrencePeriod       *Period                  `json:"occurrencePeriod,omitempty" bson:"occurrence_period,omitempty"`      // Desired time or schedule for use
	OccurrenceTiming       *Timing                  `json:"occurrenceTiming,omitempty" bson:"occurrence_timing,omitempty"`      // Desired time or schedule for use
	AuthoredOn             *string                  `json:"authoredOn,omitempty" bson:"authored_on,omitempty"`                  // When recorded
	Requester              *Reference               `json:"requester,omitempty" bson:"requester,omitempty"`                     // Who/what submitted the device request
	Performer              *CodeableReference       `json:"performer,omitempty" bson:"performer,omitempty"`                     // Requested Filler
	Location               []CodeableReference      `json:"location,omitempty" bson:"location,omitempty"`                       // Requested location
	Reason                 []CodeableReference      `json:"reason,omitempty" bson:"reason,omitempty"`                           // Coded/Linked Reason for request
	AsNeeded               bool                     `json:"asNeeded,omitempty" bson:"as_needed,omitempty"`                      // PRN status of request
	AsNeededFor            *CodeableConcept         `json:"asNeededFor,omitempty" bson:"as_needed_for,omitempty"`               // Device usage reason
	Insurance              []Reference              `json:"insurance,omitempty" bson:"insurance,omitempty"`                     // Associated insurance coverage
	SupportingInfo         []Reference              `json:"supportingInfo,omitempty" bson:"supporting_info,omitempty"`          // Additional clinical information
	Note                   []Annotation             `json:"note,omitempty" bson:"note,omitempty"`                               // Notes or comments
	RelevantHistory        []Reference              `json:"relevantHistory,omitempty" bson:"relevant_history,omitempty"`        // Request provenance
}

func (r *DeviceRequest) Validate() error {
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
	if r.GroupIdentifier != nil {
		if err := r.GroupIdentifier.Validate(); err != nil {
			return fmt.Errorf("GroupIdentifier: %w", err)
		}
	}
	var emptyString string
	if r.Intent == emptyString {
		return fmt.Errorf("field 'Intent' is required")
	}
	if r.ProductCodeableConcept == nil {
		return fmt.Errorf("field 'ProductCodeableConcept' is required")
	}
	if r.ProductCodeableConcept != nil {
		if err := r.ProductCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("ProductCodeableConcept: %w", err)
		}
	}
	if r.ProductReference == nil {
		return fmt.Errorf("field 'ProductReference' is required")
	}
	if r.ProductReference != nil {
		if err := r.ProductReference.Validate(); err != nil {
			return fmt.Errorf("ProductReference: %w", err)
		}
	}
	if r.ProductCanonical == nil {
		return fmt.Errorf("field 'ProductCanonical' is required")
	}
	for i, item := range r.Parameter {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Parameter[%d]: %w", i, err)
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
	if r.Requester != nil {
		if err := r.Requester.Validate(); err != nil {
			return fmt.Errorf("Requester: %w", err)
		}
	}
	if r.Performer != nil {
		if err := r.Performer.Validate(); err != nil {
			return fmt.Errorf("Performer: %w", err)
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
	if r.AsNeededFor != nil {
		if err := r.AsNeededFor.Validate(); err != nil {
			return fmt.Errorf("AsNeededFor: %w", err)
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
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	for i, item := range r.RelevantHistory {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RelevantHistory[%d]: %w", i, err)
		}
	}
	return nil
}

type DeviceRequestParameter struct {
	Id                   *string          `json:"id,omitempty" bson:"id,omitempty"`                                       // Unique id for inter-element referencing
	Code                 *CodeableConcept `json:"code,omitempty" bson:"code,omitempty"`                                   // Device detail
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty" bson:"value_codeable_concept,omitempty"` // Value of detail
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty" bson:"value_quantity,omitempty"`                // Value of detail
	ValueRange           *Range           `json:"valueRange,omitempty" bson:"value_range,omitempty"`                      // Value of detail
	ValueBoolean         *bool            `json:"valueBoolean,omitempty" bson:"value_boolean,omitempty"`                  // Value of detail
}

func (r *DeviceRequestParameter) Validate() error {
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.ValueCodeableConcept != nil {
		if err := r.ValueCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("ValueCodeableConcept: %w", err)
		}
	}
	if r.ValueQuantity != nil {
		if err := r.ValueQuantity.Validate(); err != nil {
			return fmt.Errorf("ValueQuantity: %w", err)
		}
	}
	if r.ValueRange != nil {
		if err := r.ValueRange.Validate(); err != nil {
			return fmt.Errorf("ValueRange: %w", err)
		}
	}
	return nil
}
