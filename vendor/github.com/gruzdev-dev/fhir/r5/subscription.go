package models

import (
	"encoding/json"
	"fmt"
)

// The subscription resource describes a particular client's request to be notified about a SubscriptionTopic.
type Subscription struct {
	Id              *string                 `json:"id,omitempty" bson:"id,omitempty"`                            // Logical id of this artifact
	Meta            *Meta                   `json:"meta,omitempty" bson:"meta,omitempty"`                        // Metadata about the resource
	ImplicitRules   *string                 `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`     // A set of rules under which this content was created
	Language        *string                 `json:"language,omitempty" bson:"language,omitempty"`                // Language of the resource content
	Text            *Narrative              `json:"text,omitempty" bson:"text,omitempty"`                        // Text summary of the resource, for human interpretation
	Contained       []json.RawMessage       `json:"contained,omitempty" bson:"contained,omitempty"`              // Contained, inline Resources
	Identifier      []Identifier            `json:"identifier,omitempty" bson:"identifier,omitempty"`            // Additional identifiers (business identifier)
	Name            *string                 `json:"name,omitempty" bson:"name,omitempty"`                        // Human readable name for this subscription
	Status          string                  `json:"status" bson:"status"`                                        // requested | active | error | off | entered-in-error
	Topic           string                  `json:"topic" bson:"topic"`                                          // Reference to the subscription topic being subscribed to
	Contact         []ContactPoint          `json:"contact,omitempty" bson:"contact,omitempty"`                  // Contact details for source (e.g. troubleshooting)
	End             *string                 `json:"end,omitempty" bson:"end,omitempty"`                          // When to automatically delete the subscription
	ManagingEntity  *Reference              `json:"managingEntity,omitempty" bson:"managing_entity,omitempty"`   // Entity responsible for Subscription changes
	Reason          *string                 `json:"reason,omitempty" bson:"reason,omitempty"`                    // Description of why this subscription was created
	FilterBy        []SubscriptionFilterBy  `json:"filterBy,omitempty" bson:"filter_by,omitempty"`               // Criteria for narrowing the subscription topic stream
	ChannelType     *Coding                 `json:"channelType" bson:"channel_type"`                             // Channel type for notifications
	Endpoint        *string                 `json:"endpoint,omitempty" bson:"endpoint,omitempty"`                // URL where the channel sends notifications
	Parameter       []SubscriptionParameter `json:"parameter,omitempty" bson:"parameter,omitempty"`              // Channel type dependent information
	HeartbeatPeriod *int                    `json:"heartbeatPeriod,omitempty" bson:"heartbeat_period,omitempty"` // Interval in seconds to send 'heartbeat' notification
	Timeout         *int                    `json:"timeout,omitempty" bson:"timeout,omitempty"`                  // Timeout in seconds to attempt notification delivery
	ContentType     *string                 `json:"contentType,omitempty" bson:"content_type,omitempty"`         // MIME type to send, or omit for no payload
	Content         *string                 `json:"content,omitempty" bson:"content,omitempty"`                  // empty | id-only | full-resource
	MaxCount        *int                    `json:"maxCount,omitempty" bson:"max_count,omitempty"`               // Maximum number of events that can be combined in a single notification
}

func (r *Subscription) Validate() error {
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
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.Topic == emptyString {
		return fmt.Errorf("field 'Topic' is required")
	}
	for i, item := range r.Contact {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Contact[%d]: %w", i, err)
		}
	}
	if r.ManagingEntity != nil {
		if err := r.ManagingEntity.Validate(); err != nil {
			return fmt.Errorf("ManagingEntity: %w", err)
		}
	}
	for i, item := range r.FilterBy {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("FilterBy[%d]: %w", i, err)
		}
	}
	if r.ChannelType == nil {
		return fmt.Errorf("field 'ChannelType' is required")
	}
	if r.ChannelType != nil {
		if err := r.ChannelType.Validate(); err != nil {
			return fmt.Errorf("ChannelType: %w", err)
		}
	}
	for i, item := range r.Parameter {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Parameter[%d]: %w", i, err)
		}
	}
	return nil
}

type SubscriptionFilterBy struct {
	Id              *string           `json:"id,omitempty" bson:"id,omitempty"`                 // Unique id for inter-element referencing
	Resource        *string           `json:"resource,omitempty" bson:"resource,omitempty"`     // Allowed Resource (reference to definition) for this Subscription filter
	FilterParameter string            `json:"filterParameter" bson:"filter_parameter"`          // Filter label defined in SubscriptionTopic
	Comparator      *string           `json:"comparator,omitempty" bson:"comparator,omitempty"` // eq | ne | gt | lt | ge | le | sa | eb | ap
	Modifier        *string           `json:"modifier,omitempty" bson:"modifier,omitempty"`     // missing | exact | contains | not | text | in | not-in | below | above | type | identifier | of-type | code-text | text-advanced | iterate
	Value           string            `json:"value" bson:"value"`                               // Literal value or resource path
	Event           []CodeableConcept `json:"event,omitempty" bson:"event,omitempty"`           // Event to filter by
}

func (r *SubscriptionFilterBy) Validate() error {
	var emptyString string
	if r.FilterParameter == emptyString {
		return fmt.Errorf("field 'FilterParameter' is required")
	}
	if r.Value == emptyString {
		return fmt.Errorf("field 'Value' is required")
	}
	for i, item := range r.Event {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Event[%d]: %w", i, err)
		}
	}
	return nil
}

type SubscriptionParameter struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Name  string  `json:"name" bson:"name"`                 // Name (key) of the parameter
	Value string  `json:"value" bson:"value"`               // Value of the parameter to use or pass through
}

func (r *SubscriptionParameter) Validate() error {
	var emptyString string
	if r.Name == emptyString {
		return fmt.Errorf("field 'Name' is required")
	}
	if r.Value == emptyString {
		return fmt.Errorf("field 'Value' is required")
	}
	return nil
}
