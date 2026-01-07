package models

import (
	"encoding/json"
	"fmt"
)

// Describes a stream of resource state changes identified by trigger criteria and annotated with labels useful to filter projections from this topic.
type SubscriptionTopic struct {
	Id                     *string                    `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta                      `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string                    `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string                    `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative                 `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage          `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Url                    string                     `json:"url" bson:"url"`                                                             // Canonical identifier for this subscription topic, represented as an absolute URI (globally unique)
	Identifier             []Identifier               `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // Business identifier for subscription topic
	Version                *string                    `json:"version,omitempty" bson:"version,omitempty"`                                 // Business version of the subscription topic
	VersionAlgorithmString *string                    `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"` // How to compare versions
	VersionAlgorithmCoding *Coding                    `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"` // How to compare versions
	Name                   *string                    `json:"name,omitempty" bson:"name,omitempty"`                                       // Name for this subscription topic (computer friendly)
	Title                  *string                    `json:"title,omitempty" bson:"title,omitempty"`                                     // Name for this subscription topic (human friendly)
	DerivedFrom            []string                   `json:"derivedFrom,omitempty" bson:"derived_from,omitempty"`                        // Based on FHIR protocol or definition
	Status                 string                     `json:"status" bson:"status"`                                                       // draft | active | retired | unknown
	Experimental           bool                       `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // If For testing only - never for real usage
	Date                   *string                    `json:"date,omitempty" bson:"date,omitempty"`                                       // Date status first applied
	Publisher              *string                    `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // The name of the individual or organization that published the SubscriptionTopic
	Contact                []ContactDetail            `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Description            *string                    `json:"description,omitempty" bson:"description,omitempty"`                         // Natural language description of the SubscriptionTopic
	UseContext             []UsageContext             `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // Content intends to support these contexts
	Jurisdiction           []CodeableConcept          `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                       // Intended jurisdiction of the SubscriptionTopic (if applicable)
	Purpose                *string                    `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // Why this SubscriptionTopic is defined
	Copyright              *string                    `json:"copyright,omitempty" bson:"copyright,omitempty"`                             // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel         *string                    `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                  // Copyright holder and year(s)
	ApprovalDate           *string                    `json:"approvalDate,omitempty" bson:"approval_date,omitempty"`                      // When SubscriptionTopic is/was approved by publisher
	LastReviewDate         *string                    `json:"lastReviewDate,omitempty" bson:"last_review_date,omitempty"`                 // Date the Subscription Topic was last reviewed by the publisher
	EffectivePeriod        *Period                    `json:"effectivePeriod,omitempty" bson:"effective_period,omitempty"`                // The effective date range for the SubscriptionTopic
	Trigger                []SubscriptionTopicTrigger `json:"trigger,omitempty" bson:"trigger,omitempty"`                                 // Definition of a trigger for the subscription topic
}

func (r *SubscriptionTopic) Validate() error {
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
	if r.Url == emptyString {
		return fmt.Errorf("field 'Url' is required")
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
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
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
	if r.EffectivePeriod != nil {
		if err := r.EffectivePeriod.Validate(); err != nil {
			return fmt.Errorf("EffectivePeriod: %w", err)
		}
	}
	for i, item := range r.Trigger {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Trigger[%d]: %w", i, err)
		}
	}
	return nil
}

type SubscriptionTopicTriggerCanFilterBy struct {
	Id               *string  `json:"id,omitempty" bson:"id,omitempty"`                              // Unique id for inter-element referencing
	Description      *string  `json:"description,omitempty" bson:"description,omitempty"`            // Description of this filter parameter
	Resource         *string  `json:"resource,omitempty" bson:"resource,omitempty"`                  // URL of the triggering Resource that this filter applies to
	FilterParameter  string   `json:"filterParameter" bson:"filter_parameter"`                       // Human-readable and computation-friendly name for a filter parameter usable by subscriptions on this topic, via Subscription.filterBy.filterParameter
	FilterDefinition *string  `json:"filterDefinition,omitempty" bson:"filter_definition,omitempty"` // Canonical URL for a filterParameter definition
	Comparator       []string `json:"comparator,omitempty" bson:"comparator,omitempty"`              // eq | ne | gt | lt | ge | le | sa | eb | ap
	Modifier         []string `json:"modifier,omitempty" bson:"modifier,omitempty"`                  // missing | exact | contains | not | text | in | not-in | below | above | type | identifier | of-type | code-text | text-advanced | iterate
}

func (r *SubscriptionTopicTriggerCanFilterBy) Validate() error {
	var emptyString string
	if r.FilterParameter == emptyString {
		return fmt.Errorf("field 'FilterParameter' is required")
	}
	return nil
}

type SubscriptionTopicTriggerNotificationShape struct {
	Id           *string                                                 `json:"id,omitempty" bson:"id,omitempty"`                      // Unique id for inter-element referencing
	Resource     string                                                  `json:"resource" bson:"resource"`                              // URL of the key definition that is the focus in a notification shape
	Include      []string                                                `json:"include,omitempty" bson:"include,omitempty"`            // Include directives, rooted in the resource for this shape
	RevInclude   []string                                                `json:"revInclude,omitempty" bson:"rev_include,omitempty"`     // Reverse include directives, rooted in the resource for this shape
	RelatedQuery []SubscriptionTopicTriggerNotificationShapeRelatedQuery `json:"relatedQuery,omitempty" bson:"related_query,omitempty"` // Query describing data relevant to this notification
}

func (r *SubscriptionTopicTriggerNotificationShape) Validate() error {
	var emptyString string
	if r.Resource == emptyString {
		return fmt.Errorf("field 'Resource' is required")
	}
	for i, item := range r.RelatedQuery {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RelatedQuery[%d]: %w", i, err)
		}
	}
	return nil
}

type SubscriptionTopicTriggerNotificationShapeRelatedQuery struct {
	Id        *string `json:"id,omitempty" bson:"id,omitempty"`                // Unique id for inter-element referencing
	QueryType *Coding `json:"queryType,omitempty" bson:"query_type,omitempty"` // Coded information describing the type of data this query provides
	Query     string  `json:"query" bson:"query"`                              // Query to perform
}

func (r *SubscriptionTopicTriggerNotificationShapeRelatedQuery) Validate() error {
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

type SubscriptionTopicTrigger struct {
	Id                   *string                                     `json:"id,omitempty" bson:"id,omitempty"`                                      // Unique id for inter-element referencing
	Description          *string                                     `json:"description,omitempty" bson:"description,omitempty"`                    // Text representation of the resource trigger
	Resource             string                                      `json:"resource" bson:"resource"`                                              // Key Data Type, Resource (reference to definition), or relevant definition for this trigger
	SupportedInteraction []string                                    `json:"supportedInteraction,omitempty" bson:"supported_interaction,omitempty"` // create | update | delete
	QueryCriteria        *SubscriptionTopicTriggerQueryCriteria      `json:"queryCriteria,omitempty" bson:"query_criteria,omitempty"`               // Query based trigger rule
	FhirPathCriteria     *string                                     `json:"fhirPathCriteria,omitempty" bson:"fhir_path_criteria,omitempty"`        // FHIRPath based trigger rule
	Event                *CodeableConcept                            `json:"event,omitempty" bson:"event,omitempty"`                                // Event which can trigger a notification from the SubscriptionTopic
	CanFilterBy          []SubscriptionTopicTriggerCanFilterBy       `json:"canFilterBy,omitempty" bson:"can_filter_by,omitempty"`                  // Properties by which a Subscription can filter notifications based on this trigger
	NotificationShape    []SubscriptionTopicTriggerNotificationShape `json:"notificationShape,omitempty" bson:"notification_shape,omitempty"`       // Properties for describing the shape of notifications generated by this trigger
}

func (r *SubscriptionTopicTrigger) Validate() error {
	var emptyString string
	if r.Resource == emptyString {
		return fmt.Errorf("field 'Resource' is required")
	}
	if r.QueryCriteria != nil {
		if err := r.QueryCriteria.Validate(); err != nil {
			return fmt.Errorf("QueryCriteria: %w", err)
		}
	}
	if r.Event != nil {
		if err := r.Event.Validate(); err != nil {
			return fmt.Errorf("Event: %w", err)
		}
	}
	for i, item := range r.CanFilterBy {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("CanFilterBy[%d]: %w", i, err)
		}
	}
	for i, item := range r.NotificationShape {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("NotificationShape[%d]: %w", i, err)
		}
	}
	return nil
}

type SubscriptionTopicTriggerQueryCriteria struct {
	Id              *string `json:"id,omitempty" bson:"id,omitempty"`                             // Unique id for inter-element referencing
	Previous        *string `json:"previous,omitempty" bson:"previous,omitempty"`                 // Rule applied to previous resource state
	ResultForCreate *string `json:"resultForCreate,omitempty" bson:"result_for_create,omitempty"` // test-passes | test-fails
	Current         *string `json:"current,omitempty" bson:"current,omitempty"`                   // Rule applied to current resource state
	ResultForDelete *string `json:"resultForDelete,omitempty" bson:"result_for_delete,omitempty"` // test-passes | test-fails
	RequireBoth     bool    `json:"requireBoth,omitempty" bson:"require_both,omitempty"`          // Both must be true flag
}

func (r *SubscriptionTopicTriggerQueryCriteria) Validate() error {
	return nil
}
