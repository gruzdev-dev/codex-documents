package models

import (
	"encoding/json"
	"fmt"
)

// A clinical or business level record of information being transmitted or shared; e.g. an alert that was sent to a responsible provider, a public health agency communication to a provider/reporter in response to a case report for a reportable condition.
type Communication struct {
	Id            *string                `json:"id,omitempty" bson:"id,omitempty"`                        // Logical id of this artifact
	Meta          *Meta                  `json:"meta,omitempty" bson:"meta,omitempty"`                    // Metadata about the resource
	ImplicitRules *string                `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"` // A set of rules under which this content was created
	Language      *string                `json:"language,omitempty" bson:"language,omitempty"`            // Language of the resource content
	Text          *Narrative             `json:"text,omitempty" bson:"text,omitempty"`                    // Text summary of the resource, for human interpretation
	Contained     []json.RawMessage      `json:"contained,omitempty" bson:"contained,omitempty"`          // Contained, inline Resources
	Identifier    []Identifier           `json:"identifier,omitempty" bson:"identifier,omitempty"`        // Unique identifier
	BasedOn       []Reference            `json:"basedOn,omitempty" bson:"based_on,omitempty"`             // Request fulfilled by this communication
	PartOf        []Reference            `json:"partOf,omitempty" bson:"part_of,omitempty"`               // Part of referenced event (e.g. Communication, Procedure)
	InResponseTo  []Reference            `json:"inResponseTo,omitempty" bson:"in_response_to,omitempty"`  // Reply to
	Status        string                 `json:"status" bson:"status"`                                    // preparation | in-progress | not-done | on-hold | stopped | completed | entered-in-error | unknown
	StatusReason  *CodeableConcept       `json:"statusReason,omitempty" bson:"status_reason,omitempty"`   // Reason for current status
	Category      []CodeableConcept      `json:"category,omitempty" bson:"category,omitempty"`            // Message category
	Priority      *string                `json:"priority,omitempty" bson:"priority,omitempty"`            // routine | urgent | asap | stat
	Medium        []CodeableConcept      `json:"medium,omitempty" bson:"medium,omitempty"`                // A channel of communication
	Subject       *Reference             `json:"subject,omitempty" bson:"subject,omitempty"`              // Focus of message
	Topic         *CodeableConcept       `json:"topic,omitempty" bson:"topic,omitempty"`                  // Description of the purpose/content
	About         []Reference            `json:"about,omitempty" bson:"about,omitempty"`                  // Resources that pertain to this communication
	Encounter     *Reference             `json:"encounter,omitempty" bson:"encounter,omitempty"`          // The Encounter during which this Communication was created
	Sent          *string                `json:"sent,omitempty" bson:"sent,omitempty"`                    // When sent
	Received      *string                `json:"received,omitempty" bson:"received,omitempty"`            // When received
	Recipient     []Reference            `json:"recipient,omitempty" bson:"recipient,omitempty"`          // Who the information is shared with
	Sender        *Reference             `json:"sender,omitempty" bson:"sender,omitempty"`                // Who shares the information
	Reason        []CodeableReference    `json:"reason,omitempty" bson:"reason,omitempty"`                // Indication for message
	Payload       []CommunicationPayload `json:"payload,omitempty" bson:"payload,omitempty"`              // Message payload
	Note          []Annotation           `json:"note,omitempty" bson:"note,omitempty"`                    // Comments made about the communication
}

func (r *Communication) Validate() error {
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
	for i, item := range r.InResponseTo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("InResponseTo[%d]: %w", i, err)
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
	if r.Topic != nil {
		if err := r.Topic.Validate(); err != nil {
			return fmt.Errorf("Topic: %w", err)
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
	for i, item := range r.Recipient {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Recipient[%d]: %w", i, err)
		}
	}
	if r.Sender != nil {
		if err := r.Sender.Validate(); err != nil {
			return fmt.Errorf("Sender: %w", err)
		}
	}
	for i, item := range r.Reason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reason[%d]: %w", i, err)
		}
	}
	for i, item := range r.Payload {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Payload[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	return nil
}

type CommunicationPayload struct {
	Id                     *string          `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	ContentAttachment      *Attachment      `json:"contentAttachment" bson:"content_attachment"`            // Message part content
	ContentReference       *Reference       `json:"contentReference" bson:"content_reference"`              // Message part content
	ContentCodeableConcept *CodeableConcept `json:"contentCodeableConcept" bson:"content_codeable_concept"` // Message part content
}

func (r *CommunicationPayload) Validate() error {
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
