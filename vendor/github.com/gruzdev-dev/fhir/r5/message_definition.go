package models

import (
	"encoding/json"
	"fmt"
)

// Defines the characteristics of a message that can be shared between systems, including the type of event that initiates the message, the content to be transmitted and what response(s), if any, are permitted.
type MessageDefinition struct {
	Id                     *string                            `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta                              `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string                            `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string                            `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative                         `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage                  `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Url                    *string                            `json:"url,omitempty" bson:"url,omitempty"`                                         // The cannonical URL for a given MessageDefinition
	Identifier             []Identifier                       `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // Business Identifier for a given MessageDefinition
	Version                *string                            `json:"version,omitempty" bson:"version,omitempty"`                                 // Business version of the message definition
	VersionAlgorithmString *string                            `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"` // How to compare versions
	VersionAlgorithmCoding *Coding                            `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"` // How to compare versions
	Name                   *string                            `json:"name,omitempty" bson:"name,omitempty"`                                       // Name for this message definition (computer friendly)
	Title                  *string                            `json:"title,omitempty" bson:"title,omitempty"`                                     // Name for this message definition (human friendly)
	Replaces               []string                           `json:"replaces,omitempty" bson:"replaces,omitempty"`                               // Takes the place of
	Status                 string                             `json:"status" bson:"status"`                                                       // draft | active | retired | unknown
	Experimental           bool                               `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // For testing only - never for real usage
	Date                   string                             `json:"date" bson:"date"`                                                           // Date last changed
	Publisher              *string                            `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // Name of the publisher/steward (organization or individual)
	Contact                []ContactDetail                    `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Description            *string                            `json:"description,omitempty" bson:"description,omitempty"`                         // Natural language description of the message definition
	UseContext             []UsageContext                     `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // The context that the content is intended to support
	Jurisdiction           []CodeableConcept                  `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                       // Jurisdiction of the authority that maintains the message definition (if applicable)
	Purpose                *string                            `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // Why this message definition is defined
	Copyright              *string                            `json:"copyright,omitempty" bson:"copyright,omitempty"`                             // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel         *string                            `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                  // Copyright holder and year(s)
	Base                   *string                            `json:"base,omitempty" bson:"base,omitempty"`                                       // Definition this one is based on
	Parent                 []string                           `json:"parent,omitempty" bson:"parent,omitempty"`                                   // Protocol/workflow this is part of
	EventCoding            *Coding                            `json:"eventCoding" bson:"event_coding"`                                            // Event code  or link to the EventDefinition
	EventUri               *string                            `json:"eventUri" bson:"event_uri"`                                                  // Event code  or link to the EventDefinition
	Category               *string                            `json:"category,omitempty" bson:"category,omitempty"`                               // consequence | currency | notification
	Focus                  []MessageDefinitionFocus           `json:"focus,omitempty" bson:"focus,omitempty"`                                     // Resource(s) that are the subject of the event
	ResponseRequired       *string                            `json:"responseRequired,omitempty" bson:"response_required,omitempty"`              // always | on-error | never | on-success
	AllowedResponse        []MessageDefinitionAllowedResponse `json:"allowedResponse,omitempty" bson:"allowed_response,omitempty"`                // Responses to this message
}

func (r *MessageDefinition) Validate() error {
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
	if r.VersionAlgorithmCoding != nil {
		if err := r.VersionAlgorithmCoding.Validate(); err != nil {
			return fmt.Errorf("VersionAlgorithmCoding: %w", err)
		}
	}
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.Date == emptyString {
		return fmt.Errorf("field 'Date' is required")
	}
	for i, item := range r.Contact {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Contact[%d]: %w", i, err)
		}
	}
	for i, item := range r.UseContext {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("UseContext[%d]: %w", i, err)
		}
	}
	for i, item := range r.Jurisdiction {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Jurisdiction[%d]: %w", i, err)
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
	for i, item := range r.Focus {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Focus[%d]: %w", i, err)
		}
	}
	for i, item := range r.AllowedResponse {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AllowedResponse[%d]: %w", i, err)
		}
	}
	return nil
}

type MessageDefinitionFocus struct {
	Id      *string `json:"id,omitempty" bson:"id,omitempty"`           // Unique id for inter-element referencing
	Code    string  `json:"code" bson:"code"`                           // Type of resource
	Profile *string `json:"profile,omitempty" bson:"profile,omitempty"` // Profile that must be adhered to by focus
	Min     int     `json:"min" bson:"min"`                             // Minimum number of focuses of this type
	Max     *string `json:"max,omitempty" bson:"max,omitempty"`         // Maximum number of focuses of this type
}

func (r *MessageDefinitionFocus) Validate() error {
	var emptyString string
	if r.Code == emptyString {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Min == 0 {
		return fmt.Errorf("field 'Min' is required")
	}
	return nil
}

type MessageDefinitionAllowedResponse struct {
	Id        *string `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Message   string  `json:"message" bson:"message"`                         // Reference to allowed message definition response
	Situation *string `json:"situation,omitempty" bson:"situation,omitempty"` // When should this response be used
}

func (r *MessageDefinitionAllowedResponse) Validate() error {
	var emptyString string
	if r.Message == emptyString {
		return fmt.Errorf("field 'Message' is required")
	}
	return nil
}
