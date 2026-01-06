package models

import (
	"encoding/json"
	"fmt"
)

// A set of related requests that can be used to capture intended activities that have inter-dependencies such as "give this medication after that one".
type RequestOrchestration struct {
	Id                    *string                      `json:"id,omitempty" bson:"id,omitempty"`                                        // Logical id of this artifact
	Meta                  *Meta                        `json:"meta,omitempty" bson:"meta,omitempty"`                                    // Metadata about the resource
	ImplicitRules         *string                      `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                 // A set of rules under which this content was created
	Language              *string                      `json:"language,omitempty" bson:"language,omitempty"`                            // Language of the resource content
	Text                  *Narrative                   `json:"text,omitempty" bson:"text,omitempty"`                                    // Text summary of the resource, for human interpretation
	Contained             []json.RawMessage            `json:"contained,omitempty" bson:"contained,omitempty"`                          // Contained, inline Resources
	Identifier            []Identifier                 `json:"identifier,omitempty" bson:"identifier,omitempty"`                        // Business identifier
	InstantiatesCanonical []string                     `json:"instantiatesCanonical,omitempty" bson:"instantiates_canonical,omitempty"` // Instantiates FHIR protocol or definition
	InstantiatesUri       []string                     `json:"instantiatesUri,omitempty" bson:"instantiates_uri,omitempty"`             // Instantiates external protocol or definition
	BasedOn               []Reference                  `json:"basedOn,omitempty" bson:"based_on,omitempty"`                             // Fulfills plan, proposal, or order
	Replaces              []Reference                  `json:"replaces,omitempty" bson:"replaces,omitempty"`                            // Request(s) replaced by this request
	GroupIdentifier       *Identifier                  `json:"groupIdentifier,omitempty" bson:"group_identifier,omitempty"`             // Composite request this is part of
	Status                string                       `json:"status" bson:"status"`                                                    // draft | active | on-hold | entered-in-error | ended | completed | revoked | unknown
	Intent                string                       `json:"intent" bson:"intent"`                                                    // proposal | solicit-offer | offer-response | plan | directive | order | original-order | reflex-order | filler-order | instance-order | option
	Priority              *string                      `json:"priority,omitempty" bson:"priority,omitempty"`                            // routine | urgent | asap | stat
	Code                  *CodeableConcept             `json:"code,omitempty" bson:"code,omitempty"`                                    // What's being requested/ordered
	Subject               *Reference                   `json:"subject,omitempty" bson:"subject,omitempty"`                              // Who the request orchestration is about
	Encounter             *Reference                   `json:"encounter,omitempty" bson:"encounter,omitempty"`                          // Created as part of
	AuthoredOn            *string                      `json:"authoredOn,omitempty" bson:"authored_on,omitempty"`                       // When the request orchestration was authored
	Author                *Reference                   `json:"author,omitempty" bson:"author,omitempty"`                                // Device or practitioner that authored the request orchestration
	Reason                []CodeableReference          `json:"reason,omitempty" bson:"reason,omitempty"`                                // Why the request orchestration is needed
	Goal                  []Reference                  `json:"goal,omitempty" bson:"goal,omitempty"`                                    // What goals
	Note                  []Annotation                 `json:"note,omitempty" bson:"note,omitempty"`                                    // Additional notes about the response
	Action                []RequestOrchestrationAction `json:"action,omitempty" bson:"action,omitempty"`                                // Proposed actions, if any
}

func (r *RequestOrchestration) Validate() error {
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
	if r.Intent == emptyString {
		return fmt.Errorf("field 'Intent' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
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
	if r.Author != nil {
		if err := r.Author.Validate(); err != nil {
			return fmt.Errorf("Author: %w", err)
		}
	}
	for i, item := range r.Reason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reason[%d]: %w", i, err)
		}
	}
	for i, item := range r.Goal {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Goal[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	for i, item := range r.Action {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Action[%d]: %w", i, err)
		}
	}
	return nil
}

type RequestOrchestrationActionOutput struct {
	Id          *string          `json:"id,omitempty" bson:"id,omitempty"`                    // Unique id for inter-element referencing
	Title       *string          `json:"title,omitempty" bson:"title,omitempty"`              // User-visible title
	Requirement *DataRequirement `json:"requirement,omitempty" bson:"requirement,omitempty"`  // What data is provided
	RelatedData *string          `json:"relatedData,omitempty" bson:"related_data,omitempty"` // What data is provided
}

func (r *RequestOrchestrationActionOutput) Validate() error {
	if r.Requirement != nil {
		if err := r.Requirement.Validate(); err != nil {
			return fmt.Errorf("Requirement: %w", err)
		}
	}
	return nil
}

type RequestOrchestrationActionRelatedAction struct {
	Id              *string   `json:"id,omitempty" bson:"id,omitempty"`                            // Unique id for inter-element referencing
	TargetId        string    `json:"targetId" bson:"target_id"`                                   // What action this is related to
	Relationship    string    `json:"relationship" bson:"relationship"`                            // before | before-start | before-end | concurrent | concurrent-with-start | concurrent-with-end | after | after-start | after-end
	EndRelationship *string   `json:"endRelationship,omitempty" bson:"end_relationship,omitempty"` // before | before-start | before-end | concurrent | concurrent-with-start | concurrent-with-end | after | after-start | after-end
	OffsetDuration  *Duration `json:"offsetDuration,omitempty" bson:"offset_duration,omitempty"`   // Time offset for the relationship
	OffsetRange     *Range    `json:"offsetRange,omitempty" bson:"offset_range,omitempty"`         // Time offset for the relationship
}

func (r *RequestOrchestrationActionRelatedAction) Validate() error {
	var emptyString string
	if r.TargetId == emptyString {
		return fmt.Errorf("field 'TargetId' is required")
	}
	if r.Relationship == emptyString {
		return fmt.Errorf("field 'Relationship' is required")
	}
	if r.OffsetDuration != nil {
		if err := r.OffsetDuration.Validate(); err != nil {
			return fmt.Errorf("OffsetDuration: %w", err)
		}
	}
	if r.OffsetRange != nil {
		if err := r.OffsetRange.Validate(); err != nil {
			return fmt.Errorf("OffsetRange: %w", err)
		}
	}
	return nil
}

type RequestOrchestrationActionParticipant struct {
	Id             *string          `json:"id,omitempty" bson:"id,omitempty"`                          // Unique id for inter-element referencing
	Type           *string          `json:"type,omitempty" bson:"type,omitempty"`                      // careteam | device | group | healthcareservice | location | organization | patient | practitioner | practitionerrole | relatedperson
	TypeCanonical  *string          `json:"typeCanonical,omitempty" bson:"type_canonical,omitempty"`   // Who or what can participate
	TypeReference  *Reference       `json:"typeReference,omitempty" bson:"type_reference,omitempty"`   // Who or what can participate
	Role           *CodeableConcept `json:"role,omitempty" bson:"role,omitempty"`                      // E.g. Nurse, Surgeon, Parent, etc
	Function       *CodeableConcept `json:"function,omitempty" bson:"function,omitempty"`              // E.g. Author, Reviewer, Witness, etc
	ActorCanonical *string          `json:"actorCanonical,omitempty" bson:"actor_canonical,omitempty"` // Who/what is participating?
	ActorReference *Reference       `json:"actorReference,omitempty" bson:"actor_reference,omitempty"` // Who/what is participating?
}

func (r *RequestOrchestrationActionParticipant) Validate() error {
	if r.TypeReference != nil {
		if err := r.TypeReference.Validate(); err != nil {
			return fmt.Errorf("TypeReference: %w", err)
		}
	}
	if r.Role != nil {
		if err := r.Role.Validate(); err != nil {
			return fmt.Errorf("Role: %w", err)
		}
	}
	if r.Function != nil {
		if err := r.Function.Validate(); err != nil {
			return fmt.Errorf("Function: %w", err)
		}
	}
	if r.ActorReference != nil {
		if err := r.ActorReference.Validate(); err != nil {
			return fmt.Errorf("ActorReference: %w", err)
		}
	}
	return nil
}

type RequestOrchestrationAction struct {
	Id                    *string                                   `json:"id,omitempty" bson:"id,omitempty"`                                        // Unique id for inter-element referencing
	LinkId                *string                                   `json:"linkId,omitempty" bson:"link_id,omitempty"`                               // Pointer to specific item from the PlanDefinition
	Prefix                *string                                   `json:"prefix,omitempty" bson:"prefix,omitempty"`                                // User-visible prefix for the action (e.g. 1. or A.)
	Title                 *string                                   `json:"title,omitempty" bson:"title,omitempty"`                                  // User-visible title
	Description           *string                                   `json:"description,omitempty" bson:"description,omitempty"`                      // Short description of the action
	TextEquivalent        *string                                   `json:"textEquivalent,omitempty" bson:"text_equivalent,omitempty"`               // Static text equivalent of the action, used if the dynamic aspects cannot be interpreted by the receiving system
	Priority              *string                                   `json:"priority,omitempty" bson:"priority,omitempty"`                            // routine | urgent | asap | stat
	Code                  []CodeableConcept                         `json:"code,omitempty" bson:"code,omitempty"`                                    // Code representing the meaning of the action or sub-actions
	Documentation         []RelatedArtifact                         `json:"documentation,omitempty" bson:"documentation,omitempty"`                  // Supporting documentation for the intended performer of the action
	Goal                  []Reference                               `json:"goal,omitempty" bson:"goal,omitempty"`                                    // What goals
	Condition             []RequestOrchestrationActionCondition     `json:"condition,omitempty" bson:"condition,omitempty"`                          // Whether or not the action is applicable
	Input                 []RequestOrchestrationActionInput         `json:"input,omitempty" bson:"input,omitempty"`                                  // Input data requirements
	Output                []RequestOrchestrationActionOutput        `json:"output,omitempty" bson:"output,omitempty"`                                // Output data definition
	RelatedAction         []RequestOrchestrationActionRelatedAction `json:"relatedAction,omitempty" bson:"related_action,omitempty"`                 // Relationship to another action
	TimingDateTime        *string                                   `json:"timingDateTime,omitempty" bson:"timing_date_time,omitempty"`              // When the action should take place
	TimingAge             *Age                                      `json:"timingAge,omitempty" bson:"timing_age,omitempty"`                         // When the action should take place
	TimingPeriod          *Period                                   `json:"timingPeriod,omitempty" bson:"timing_period,omitempty"`                   // When the action should take place
	TimingDuration        *Duration                                 `json:"timingDuration,omitempty" bson:"timing_duration,omitempty"`               // When the action should take place
	TimingRange           *Range                                    `json:"timingRange,omitempty" bson:"timing_range,omitempty"`                     // When the action should take place
	TimingTiming          *Timing                                   `json:"timingTiming,omitempty" bson:"timing_timing,omitempty"`                   // When the action should take place
	TimingRelativeTime    *RelativeTime                             `json:"timingRelativeTime,omitempty" bson:"timing_relative_time,omitempty"`      // When the action should take place
	Location              *CodeableReference                        `json:"location,omitempty" bson:"location,omitempty"`                            // Where it should happen
	Participant           []RequestOrchestrationActionParticipant   `json:"participant,omitempty" bson:"participant,omitempty"`                      // Who should perform the action
	Type                  *CodeableConcept                          `json:"type,omitempty" bson:"type,omitempty"`                                    // create | update | remove | fire-event
	ApplicabilityBehavior *string                                   `json:"applicabilityBehavior,omitempty" bson:"applicability_behavior,omitempty"` // all | any
	GroupingBehavior      *string                                   `json:"groupingBehavior,omitempty" bson:"grouping_behavior,omitempty"`           // visual-group | logical-group | sentence-group
	SelectionBehavior     *string                                   `json:"selectionBehavior,omitempty" bson:"selection_behavior,omitempty"`         // any | all | all-or-none | exactly-one | at-most-one | one-or-more
	RequiredBehavior      *string                                   `json:"requiredBehavior,omitempty" bson:"required_behavior,omitempty"`           // must | could | must-unless-documented
	PrecheckBehavior      *string                                   `json:"precheckBehavior,omitempty" bson:"precheck_behavior,omitempty"`           // yes | no
	CardinalityBehavior   *string                                   `json:"cardinalityBehavior,omitempty" bson:"cardinality_behavior,omitempty"`     // single | multiple
	Resource              *Reference                                `json:"resource,omitempty" bson:"resource,omitempty"`                            // The target of the action
	DefinitionCanonical   *string                                   `json:"definitionCanonical,omitempty" bson:"definition_canonical,omitempty"`     // Description of the activity to be performed
	DefinitionUri         *string                                   `json:"definitionUri,omitempty" bson:"definition_uri,omitempty"`                 // Description of the activity to be performed
	Transform             *string                                   `json:"transform,omitempty" bson:"transform,omitempty"`                          // Transform to apply the template
	DynamicValue          []RequestOrchestrationActionDynamicValue  `json:"dynamicValue,omitempty" bson:"dynamic_value,omitempty"`                   // Dynamic aspects of the definition
	Action                []RequestOrchestrationAction              `json:"action,omitempty" bson:"action,omitempty"`                                // Sub action
}

func (r *RequestOrchestrationAction) Validate() error {
	for i, item := range r.Code {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Code[%d]: %w", i, err)
		}
	}
	for i, item := range r.Documentation {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Documentation[%d]: %w", i, err)
		}
	}
	for i, item := range r.Goal {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Goal[%d]: %w", i, err)
		}
	}
	for i, item := range r.Condition {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Condition[%d]: %w", i, err)
		}
	}
	for i, item := range r.Input {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Input[%d]: %w", i, err)
		}
	}
	for i, item := range r.Output {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Output[%d]: %w", i, err)
		}
	}
	for i, item := range r.RelatedAction {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RelatedAction[%d]: %w", i, err)
		}
	}
	if r.TimingAge != nil {
		if err := r.TimingAge.Validate(); err != nil {
			return fmt.Errorf("TimingAge: %w", err)
		}
	}
	if r.TimingPeriod != nil {
		if err := r.TimingPeriod.Validate(); err != nil {
			return fmt.Errorf("TimingPeriod: %w", err)
		}
	}
	if r.TimingDuration != nil {
		if err := r.TimingDuration.Validate(); err != nil {
			return fmt.Errorf("TimingDuration: %w", err)
		}
	}
	if r.TimingRange != nil {
		if err := r.TimingRange.Validate(); err != nil {
			return fmt.Errorf("TimingRange: %w", err)
		}
	}
	if r.TimingTiming != nil {
		if err := r.TimingTiming.Validate(); err != nil {
			return fmt.Errorf("TimingTiming: %w", err)
		}
	}
	if r.TimingRelativeTime != nil {
		if err := r.TimingRelativeTime.Validate(); err != nil {
			return fmt.Errorf("TimingRelativeTime: %w", err)
		}
	}
	if r.Location != nil {
		if err := r.Location.Validate(); err != nil {
			return fmt.Errorf("Location: %w", err)
		}
	}
	for i, item := range r.Participant {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Participant[%d]: %w", i, err)
		}
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Resource != nil {
		if err := r.Resource.Validate(); err != nil {
			return fmt.Errorf("Resource: %w", err)
		}
	}
	for i, item := range r.DynamicValue {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("DynamicValue[%d]: %w", i, err)
		}
	}
	for i, item := range r.Action {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Action[%d]: %w", i, err)
		}
	}
	return nil
}

type RequestOrchestrationActionInput struct {
	Id          *string          `json:"id,omitempty" bson:"id,omitempty"`                    // Unique id for inter-element referencing
	Title       *string          `json:"title,omitempty" bson:"title,omitempty"`              // User-visible title
	Requirement *DataRequirement `json:"requirement,omitempty" bson:"requirement,omitempty"`  // What data is provided
	RelatedData *string          `json:"relatedData,omitempty" bson:"related_data,omitempty"` // What data is provided
}

func (r *RequestOrchestrationActionInput) Validate() error {
	if r.Requirement != nil {
		if err := r.Requirement.Validate(); err != nil {
			return fmt.Errorf("Requirement: %w", err)
		}
	}
	return nil
}

type RequestOrchestrationActionDynamicValue struct {
	Id         *string     `json:"id,omitempty" bson:"id,omitempty"`                 // Unique id for inter-element referencing
	Path       *string     `json:"path,omitempty" bson:"path,omitempty"`             // The path to the element to be set dynamically
	Expression *Expression `json:"expression,omitempty" bson:"expression,omitempty"` // An expression that provides the dynamic value for the customization
}

func (r *RequestOrchestrationActionDynamicValue) Validate() error {
	if r.Expression != nil {
		if err := r.Expression.Validate(); err != nil {
			return fmt.Errorf("Expression: %w", err)
		}
	}
	return nil
}

type RequestOrchestrationActionCondition struct {
	Id         *string     `json:"id,omitempty" bson:"id,omitempty"`                 // Unique id for inter-element referencing
	Kind       string      `json:"kind" bson:"kind"`                                 // applicability | start | stop
	Expression *Expression `json:"expression,omitempty" bson:"expression,omitempty"` // Boolean-valued expression
}

func (r *RequestOrchestrationActionCondition) Validate() error {
	var emptyString string
	if r.Kind == emptyString {
		return fmt.Errorf("field 'Kind' is required")
	}
	if r.Expression != nil {
		if err := r.Expression.Validate(); err != nil {
			return fmt.Errorf("Expression: %w", err)
		}
	}
	return nil
}
