package models

import (
	"encoding/json"
	"fmt"
)

// The header for a message exchange that is either requesting or responding to an action.  The reference(s) that are the subject of the action as well as other information related to the action are typically transmitted in a bundle in which the MessageHeader resource instance is the first resource in the bundle.
type MessageHeader struct {
	ResourceType   string                     `json:"resourceType" bson:"resource_type"`                       // Type of resource
	Id             *string                    `json:"id,omitempty" bson:"id,omitempty"`                        // Logical id of this artifact
	Meta           *Meta                      `json:"meta,omitempty" bson:"meta,omitempty"`                    // Metadata about the resource
	ImplicitRules  *string                    `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"` // A set of rules under which this content was created
	Language       *string                    `json:"language,omitempty" bson:"language,omitempty"`            // Language of the resource content
	Text           *Narrative                 `json:"text,omitempty" bson:"text,omitempty"`                    // Text summary of the resource, for human interpretation
	Contained      []json.RawMessage          `json:"contained,omitempty" bson:"contained,omitempty"`          // Contained, inline Resources
	EventCoding    *Coding                    `json:"eventCoding" bson:"event_coding"`                         // The real world event that triggered this messsage
	EventUri       *string                    `json:"eventUri" bson:"event_uri"`                               // The real world event that triggered this messsage
	EventCanonical *string                    `json:"eventCanonical" bson:"event_canonical"`                   // The real world event that triggered this messsage
	Destination    []MessageHeaderDestination `json:"destination,omitempty" bson:"destination,omitempty"`      // Message destination application(s)
	Source         *MessageHeaderSource       `json:"source" bson:"source"`                                    // Message source application
	Reason         *CodeableConcept           `json:"reason,omitempty" bson:"reason,omitempty"`                // Cause of event
	Response       *MessageHeaderResponse     `json:"response,omitempty" bson:"response,omitempty"`            // If this is a reply to prior message
	Focus          []Reference                `json:"focus,omitempty" bson:"focus,omitempty"`                  // The actual content of the message
	Definition     *string                    `json:"definition,omitempty" bson:"definition,omitempty"`        // Link to the definition for this message
}

func (r *MessageHeader) Validate() error {
	if r.ResourceType != "MessageHeader" {
		return fmt.Errorf("invalid resourceType: expected 'MessageHeader', got '%s'", r.ResourceType)
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
	if r.EventCoding == nil {
		return fmt.Errorf("field 'EventCoding' is required")
	}
	if r.EventCoding != nil {
		if err := r.EventCoding.Validate(); err != nil {
			return fmt.Errorf("EventCoding: %w", err)
		}
	}
	if r.EventUri == nil {
		return fmt.Errorf("field 'EventUri' is required")
	}
	if r.EventCanonical == nil {
		return fmt.Errorf("field 'EventCanonical' is required")
	}
	for i, item := range r.Destination {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Destination[%d]: %w", i, err)
		}
	}
	if r.Source == nil {
		return fmt.Errorf("field 'Source' is required")
	}
	if r.Source != nil {
		if err := r.Source.Validate(); err != nil {
			return fmt.Errorf("Source: %w", err)
		}
	}
	if r.Reason != nil {
		if err := r.Reason.Validate(); err != nil {
			return fmt.Errorf("Reason: %w", err)
		}
	}
	if r.Response != nil {
		if err := r.Response.Validate(); err != nil {
			return fmt.Errorf("Response: %w", err)
		}
	}
	for i, item := range r.Focus {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Focus[%d]: %w", i, err)
		}
	}
	return nil
}

type MessageHeaderDestination struct {
	Id                *string    `json:"id,omitempty" bson:"id,omitempty"`                                // Unique id for inter-element referencing
	EndpointUrl       *string    `json:"endpointUrl,omitempty" bson:"endpoint_url,omitempty"`             // Actual destination address or Endpoint resource
	EndpointReference *Reference `json:"endpointReference,omitempty" bson:"endpoint_reference,omitempty"` // Actual destination address or Endpoint resource
	Name              *string    `json:"name,omitempty" bson:"name,omitempty"`                            // Name of system
	Receiver          *Reference `json:"receiver,omitempty" bson:"receiver,omitempty"`                    // Intended "real-world" recipient for the data
}

func (r *MessageHeaderDestination) Validate() error {
	if r.EndpointReference != nil {
		if err := r.EndpointReference.Validate(); err != nil {
			return fmt.Errorf("EndpointReference: %w", err)
		}
	}
	if r.Receiver != nil {
		if err := r.Receiver.Validate(); err != nil {
			return fmt.Errorf("Receiver: %w", err)
		}
	}
	return nil
}

type MessageHeaderSource struct {
	Id                *string       `json:"id,omitempty" bson:"id,omitempty"`                                // Unique id for inter-element referencing
	EndpointUrl       *string       `json:"endpointUrl,omitempty" bson:"endpoint_url,omitempty"`             // Actual source address or Endpoint resource
	EndpointReference *Reference    `json:"endpointReference,omitempty" bson:"endpoint_reference,omitempty"` // Actual source address or Endpoint resource
	Name              *string       `json:"name,omitempty" bson:"name,omitempty"`                            // Name of system
	Software          *string       `json:"software,omitempty" bson:"software,omitempty"`                    // Name of software running the system
	Version           *string       `json:"version,omitempty" bson:"version,omitempty"`                      // Version of software running
	Contact           *ContactPoint `json:"contact,omitempty" bson:"contact,omitempty"`                      // Human contact for problems
	Sender            *Reference    `json:"sender,omitempty" bson:"sender,omitempty"`                        // Real world sender of the message
}

func (r *MessageHeaderSource) Validate() error {
	if r.EndpointReference != nil {
		if err := r.EndpointReference.Validate(); err != nil {
			return fmt.Errorf("EndpointReference: %w", err)
		}
	}
	if r.Contact != nil {
		if err := r.Contact.Validate(); err != nil {
			return fmt.Errorf("Contact: %w", err)
		}
	}
	if r.Sender != nil {
		if err := r.Sender.Validate(); err != nil {
			return fmt.Errorf("Sender: %w", err)
		}
	}
	return nil
}

type MessageHeaderResponse struct {
	Id         *string     `json:"id,omitempty" bson:"id,omitempty"`           // Unique id for inter-element referencing
	Identifier *Identifier `json:"identifier" bson:"identifier"`               // Bundle.identifier of original message
	Code       string      `json:"code" bson:"code"`                           // ok | transient-error | fatal-error
	Details    *Reference  `json:"details,omitempty" bson:"details,omitempty"` // Specific list of hints/warnings/errors
}

func (r *MessageHeaderResponse) Validate() error {
	if r.Identifier == nil {
		return fmt.Errorf("field 'Identifier' is required")
	}
	if r.Identifier != nil {
		if err := r.Identifier.Validate(); err != nil {
			return fmt.Errorf("Identifier: %w", err)
		}
	}
	var emptyString string
	if r.Code == emptyString {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Details != nil {
		if err := r.Details.Validate(); err != nil {
			return fmt.Errorf("Details: %w", err)
		}
	}
	return nil
}
