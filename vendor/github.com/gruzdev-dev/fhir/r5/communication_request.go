package models

import (
	"encoding/json"
	"fmt"
)

// A request to convey information from a sender to a recipient.
type CommunicationRequest struct {
	ResourceType        string                        `json:"resourceType" bson:"resource_type"`                                   // Type of resource
	Id                  *string                       `json:"id,omitempty" bson:"id,omitempty"`                                    // Logical id of this artifact
	Meta                *Meta                         `json:"meta,omitempty" bson:"meta,omitempty"`                                // Metadata about the resource
	ImplicitRules       *string                       `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`             // A set of rules under which this content was created
	Language            *string                       `json:"language,omitempty" bson:"language,omitempty"`                        // Language of the resource content
	Text                *Narrative                    `json:"text,omitempty" bson:"text,omitempty"`                                // Text summary of the resource, for human interpretation
	Contained           []json.RawMessage             `json:"contained,omitempty" bson:"contained,omitempty"`                      // Contained, inline Resources
	Identifier          []Identifier                  `json:"identifier,omitempty" bson:"identifier,omitempty"`                    // Unique identifier
	BasedOn             []Reference                   `json:"basedOn,omitempty" bson:"based_on,omitempty"`                         // Fulfills plan or proposal
	Replaces            []Reference                   `json:"replaces,omitempty" bson:"replaces,omitempty"`                        // Request(s) replaced by this request
	GroupIdentifier     *Identifier                   `json:"groupIdentifier,omitempty" bson:"group_identifier,omitempty"`         // Composite request this is part of
	Status              string                        `json:"status" bson:"status"`                                                // draft | active | on-hold | entered-in-error | ended | completed | revoked | unknown
	StatusReason        *CodeableConcept              `json:"statusReason,omitempty" bson:"status_reason,omitempty"`               // Reason for current status
	Intent              string                        `json:"intent" bson:"intent"`                                                // proposal | solicit-offer | offer-response | plan | directive | order | original-order | reflex-order | filler-order | instance-order | option
	Category            []CodeableConcept             `json:"category,omitempty" bson:"category,omitempty"`                        // Message category
	Priority            *string                       `json:"priority,omitempty" bson:"priority,omitempty"`                        // routine | urgent | asap | stat
	DoNotPerform        *bool                         `json:"doNotPerform,omitempty" bson:"do_not_perform,omitempty"`              // True if request is prohibiting action
	Medium              []CodeableConcept             `json:"medium,omitempty" bson:"medium,omitempty"`                            // A channel of communication
	Subject             *Reference                    `json:"subject,omitempty" bson:"subject,omitempty"`                          // Focus of message
	About               []Reference                   `json:"about,omitempty" bson:"about,omitempty"`                              // Resources that pertain to this communication request
	Encounter           *Reference                    `json:"encounter,omitempty" bson:"encounter,omitempty"`                      // The Encounter during which this CommunicationRequest was created
	Payload             []CommunicationRequestPayload `json:"payload,omitempty" bson:"payload,omitempty"`                          // Message payload
	OccurrenceDateTime  *string                       `json:"occurrenceDateTime,omitempty" bson:"occurrence_date_time,omitempty"`  // When scheduled
	OccurrencePeriod    *Period                       `json:"occurrencePeriod,omitempty" bson:"occurrence_period,omitempty"`       // When scheduled
	AuthoredOn          *string                       `json:"authoredOn,omitempty" bson:"authored_on,omitempty"`                   // When request transitioned to being actionable
	Requester           *Reference                    `json:"requester,omitempty" bson:"requester,omitempty"`                      // Who asks for the information to be shared
	Recipient           []Reference                   `json:"recipient,omitempty" bson:"recipient,omitempty"`                      // Who to share the information with
	InformationProvider []Reference                   `json:"informationProvider,omitempty" bson:"information_provider,omitempty"` // Who should share the information
	Reason              []CodeableReference           `json:"reason,omitempty" bson:"reason,omitempty"`                            // Why is communication needed?
	Note                []Annotation                  `json:"note,omitempty" bson:"note,omitempty"`                                // Comments made about communication request
}

func (r *CommunicationRequest) Validate() error {
	if r.ResourceType != "CommunicationRequest" {
		return fmt.Errorf("invalid resourceType: expected 'CommunicationRequest', got '%s'", r.ResourceType)
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
	for i, item := range r.Medium {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Medium[%d]: %w", i, err)
		}
	}
	if r.Subject != nil {
		if err := r.Subject.Validate(); err != nil {
			return fmt.Errorf("Subject: %w", err)
		}
	}
	for i, item := range r.About {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("About[%d]: %w", i, err)
		}
	}
	if r.Encounter != nil {
		if err := r.Encounter.Validate(); err != nil {
			return fmt.Errorf("Encounter: %w", err)
		}
	}
	for i, item := range r.Payload {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Payload[%d]: %w", i, err)
		}
	}
	if r.OccurrencePeriod != nil {
		if err := r.OccurrencePeriod.Validate(); err != nil {
			return fmt.Errorf("OccurrencePeriod: %w", err)
		}
	}
	if r.Requester != nil {
		if err := r.Requester.Validate(); err != nil {
			return fmt.Errorf("Requester: %w", err)
		}
	}
	for i, item := range r.Recipient {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Recipient[%d]: %w", i, err)
		}
	}
	for i, item := range r.InformationProvider {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("InformationProvider[%d]: %w", i, err)
		}
	}
	for i, item := range r.Reason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reason[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	return nil
}

type CommunicationRequestPayload struct {
	Id                     *string          `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	ContentAttachment      *Attachment      `json:"contentAttachment" bson:"content_attachment"`            // Message part content
	ContentReference       *Reference       `json:"contentReference" bson:"content_reference"`              // Message part content
	ContentCodeableConcept *CodeableConcept `json:"contentCodeableConcept" bson:"content_codeable_concept"` // Message part content
}

func (r *CommunicationRequestPayload) Validate() error {
	if r.ContentAttachment == nil {
		return fmt.Errorf("field 'ContentAttachment' is required")
	}
	if r.ContentAttachment != nil {
		if err := r.ContentAttachment.Validate(); err != nil {
			return fmt.Errorf("ContentAttachment: %w", err)
		}
	}
	if r.ContentReference == nil {
		return fmt.Errorf("field 'ContentReference' is required")
	}
	if r.ContentReference != nil {
		if err := r.ContentReference.Validate(); err != nil {
			return fmt.Errorf("ContentReference: %w", err)
		}
	}
	if r.ContentCodeableConcept == nil {
		return fmt.Errorf("field 'ContentCodeableConcept' is required")
	}
	if r.ContentCodeableConcept != nil {
		if err := r.ContentCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("ContentCodeableConcept: %w", err)
		}
	}
	return nil
}
