package models

import (
	"encoding/json"
	"fmt"
)

// The SubscriptionStatus resource describes the state of a Subscription during notifications. It is not persisted.
type SubscriptionStatus struct {
	ResourceType                 string                                `json:"resourceType" bson:"resource_type"`                                                       // Type of resource
	Id                           *string                               `json:"id,omitempty" bson:"id,omitempty"`                                                        // Logical id of this artifact
	Meta                         *Meta                                 `json:"meta,omitempty" bson:"meta,omitempty"`                                                    // Metadata about the resource
	ImplicitRules                *string                               `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                                 // A set of rules under which this content was created
	Language                     *string                               `json:"language,omitempty" bson:"language,omitempty"`                                            // Language of the resource content
	Text                         *Narrative                            `json:"text,omitempty" bson:"text,omitempty"`                                                    // Text summary of the resource, for human interpretation
	Contained                    []json.RawMessage                     `json:"contained,omitempty" bson:"contained,omitempty"`                                          // Contained, inline Resources
	Status                       *string                               `json:"status,omitempty" bson:"status,omitempty"`                                                // requested | active | error | off | entered-in-error
	Type                         string                                `json:"type" bson:"type"`                                                                        // handshake | heartbeat | event-notification | query-status | query-event
	EventsSinceSubscriptionStart *int64                                `json:"eventsSinceSubscriptionStart,omitempty" bson:"events_since_subscription_start,omitempty"` // Events since the Subscription was created
	NotificationEvent            []SubscriptionStatusNotificationEvent `json:"notificationEvent,omitempty" bson:"notification_event,omitempty"`                         // Detailed information about any events relevant to this notification
	Subscription                 *Reference                            `json:"subscription" bson:"subscription"`                                                        // Reference to the Subscription responsible for this notification
	Topic                        *string                               `json:"topic,omitempty" bson:"topic,omitempty"`                                                  // Reference to the SubscriptionTopic this notification relates to
	Error                        []CodeableConcept                     `json:"error,omitempty" bson:"error,omitempty"`                                                  // List of errors on the subscription
}

func (r *SubscriptionStatus) Validate() error {
	if r.ResourceType != "SubscriptionStatus" {
		return fmt.Errorf("invalid resourceType: expected 'SubscriptionStatus', got '%s'", r.ResourceType)
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
	var emptyString string
	if r.Type == emptyString {
		return fmt.Errorf("field 'Type' is required")
	}
	for i, item := range r.NotificationEvent {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("NotificationEvent[%d]: %w", i, err)
		}
	}
	if r.Subscription == nil {
		return fmt.Errorf("field 'Subscription' is required")
	}
	if r.Subscription != nil {
		if err := r.Subscription.Validate(); err != nil {
			return fmt.Errorf("Subscription: %w", err)
		}
	}
	for i, item := range r.Error {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Error[%d]: %w", i, err)
		}
	}
	return nil
}

type SubscriptionStatusNotificationEvent struct {
	Id                *string                                                `json:"id,omitempty" bson:"id,omitempty"`                                // Unique id for inter-element referencing
	EventNumber       int64                                                  `json:"eventNumber" bson:"event_number"`                                 // Sequencing index of this event
	TriggerEvent      []CodeableConcept                                      `json:"triggerEvent,omitempty" bson:"trigger_event,omitempty"`           // Event that triggered this notification
	Timestamp         *string                                                `json:"timestamp,omitempty" bson:"timestamp,omitempty"`                  // The instant this event occurred
	Focus             *Reference                                             `json:"focus,omitempty" bson:"focus,omitempty"`                          // Reference to the primary resource or information of this event
	AdditionalContext []Reference                                            `json:"additionalContext,omitempty" bson:"additional_context,omitempty"` // References related to the focus resource and/or context of this event
	RelatedQuery      []SubscriptionStatusNotificationEventRelatedQuery      `json:"relatedQuery,omitempty" bson:"related_query,omitempty"`           // Query describing data relevant to this notification
	AuthorizationHint []SubscriptionStatusNotificationEventAuthorizationHint `json:"authorizationHint,omitempty" bson:"authorization_hint,omitempty"` // Authorization information relevant to a notification
}

func (r *SubscriptionStatusNotificationEvent) Validate() error {
	if r.EventNumber == 0 {
		return fmt.Errorf("field 'EventNumber' is required")
	}
	for i, item := range r.TriggerEvent {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("TriggerEvent[%d]: %w", i, err)
		}
	}
	if r.Focus != nil {
		if err := r.Focus.Validate(); err != nil {
			return fmt.Errorf("Focus: %w", err)
		}
	}
	for i, item := range r.AdditionalContext {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AdditionalContext[%d]: %w", i, err)
		}
	}
	for i, item := range r.RelatedQuery {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RelatedQuery[%d]: %w", i, err)
		}
	}
	for i, item := range r.AuthorizationHint {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AuthorizationHint[%d]: %w", i, err)
		}
	}
	return nil
}

type SubscriptionStatusNotificationEventRelatedQuery struct {
	Id        *string `json:"id,omitempty" bson:"id,omitempty"`                // Unique id for inter-element referencing
	QueryType *Coding `json:"queryType,omitempty" bson:"query_type,omitempty"` // Coded information describing the type of data this query provides
	Query     string  `json:"query" bson:"query"`                              // Query to perform
}

func (r *SubscriptionStatusNotificationEventRelatedQuery) Validate() error {
	if r.QueryType != nil {
		if err := r.QueryType.Validate(); err != nil {
			return fmt.Errorf("QueryType: %w", err)
		}
	}
	var emptyString string
	if r.Query == emptyString {
		return fmt.Errorf("field 'Query' is required")
	}
	return nil
}

type SubscriptionStatusNotificationEventAuthorizationHint struct {
	Id                *string `json:"id,omitempty" bson:"id,omitempty"`            // Unique id for inter-element referencing
	AuthorizationType *Coding `json:"authorizationType" bson:"authorization_type"` // Classification of the authorization hint, e.g., 'oAuthChallengeToken'
	Value             *string `json:"value,omitempty" bson:"value,omitempty"`      // Authorization value, as defined by the 'authorizationType'
}

func (r *SubscriptionStatusNotificationEventAuthorizationHint) Validate() error {
	if r.AuthorizationType == nil {
		return fmt.Errorf("field 'AuthorizationType' is required")
	}
	if r.AuthorizationType != nil {
		if err := r.AuthorizationType.Validate(); err != nil {
			return fmt.Errorf("AuthorizationType: %w", err)
		}
	}
	return nil
}
