package models

import (
	"encoding/json"
	"fmt"
)

// This resource allows for the definition of various types of plans as a sharable, consumable, and executable artifact. The resource is general enough to support the description of a broad range of clinical and non-clinical artifacts such as clinical decision support rules, order sets, protocols, drug quality specifications, and drug manufacturing process.
type PlanDefinition struct {
	Id                      *string                `json:"id,omitempty" bson:"id,omitempty"`                                              // Logical id of this artifact
	Meta                    *Meta                  `json:"meta,omitempty" bson:"meta,omitempty"`                                          // Metadata about the resource
	ImplicitRules           *string                `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                       // A set of rules under which this content was created
	Language                *string                `json:"language,omitempty" bson:"language,omitempty"`                                  // Language of the resource content
	Text                    *Narrative             `json:"text,omitempty" bson:"text,omitempty"`                                          // Text summary of the resource, for human interpretation
	Contained               []json.RawMessage      `json:"contained,omitempty" bson:"contained,omitempty"`                                // Contained, inline Resources
	Url                     *string                `json:"url,omitempty" bson:"url,omitempty"`                                            // Canonical identifier for this plan definition, represented as a URI (globally unique)
	Identifier              []Identifier           `json:"identifier,omitempty" bson:"identifier,omitempty"`                              // Additional identifier for the plan definition
	Version                 *string                `json:"version,omitempty" bson:"version,omitempty"`                                    // Business version of the plan definition
	VersionAlgorithmString  *string                `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"`    // How to compare versions
	VersionAlgorithmCoding  *Coding                `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"`    // How to compare versions
	Name                    *string                `json:"name,omitempty" bson:"name,omitempty"`                                          // Name for this plan definition (computer friendly)
	Title                   *string                `json:"title,omitempty" bson:"title,omitempty"`                                        // Name for this plan definition (human friendly)
	Subtitle                *string                `json:"subtitle,omitempty" bson:"subtitle,omitempty"`                                  // Subordinate title of the plan definition
	Type                    *CodeableConcept       `json:"type,omitempty" bson:"type,omitempty"`                                          // order-set | protocol | eca-rule | workflow-definition | etc.
	Status                  string                 `json:"status" bson:"status"`                                                          // draft | active | retired | unknown
	Experimental            bool                   `json:"experimental,omitempty" bson:"experimental,omitempty"`                          // For testing only - never for real usage
	SubjectCodeableConcept  *CodeableConcept       `json:"subjectCodeableConcept,omitempty" bson:"subject_codeable_concept,omitempty"`    // Type of individual the plan definition is focused on
	SubjectReference        *Reference             `json:"subjectReference,omitempty" bson:"subject_reference,omitempty"`                 // Type of individual the plan definition is focused on
	SubjectCanonical        *string                `json:"subjectCanonical,omitempty" bson:"subject_canonical,omitempty"`                 // Type of individual the plan definition is focused on
	Date                    *string                `json:"date,omitempty" bson:"date,omitempty"`                                          // Date last changed
	Publisher               *string                `json:"publisher,omitempty" bson:"publisher,omitempty"`                                // Name of the publisher/steward (organization or individual)
	Contact                 []ContactDetail        `json:"contact,omitempty" bson:"contact,omitempty"`                                    // Contact details for the publisher
	Description             *string                `json:"description,omitempty" bson:"description,omitempty"`                            // Natural language description of the plan definition
	UseContext              []UsageContext         `json:"useContext,omitempty" bson:"use_context,omitempty"`                             // The context that the content is intended to support
	Jurisdiction            []CodeableConcept      `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                          // Jurisdiction of the authority that maintains the plan definition (if applicable)
	Purpose                 *string                `json:"purpose,omitempty" bson:"purpose,omitempty"`                                    // Why this plan definition is defined
	Usage                   *string                `json:"usage,omitempty" bson:"usage,omitempty"`                                        // Describes the clinical usage of the plan
	Copyright               *string                `json:"copyright,omitempty" bson:"copyright,omitempty"`                                // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel          *string                `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                     // Copyright holder and year(s)
	ApprovalDate            *string                `json:"approvalDate,omitempty" bson:"approval_date,omitempty"`                         // When the plan definition was approved by publisher
	LastReviewDate          *string                `json:"lastReviewDate,omitempty" bson:"last_review_date,omitempty"`                    // When the plan definition was last reviewed by the publisher
	EffectivePeriod         *Period                `json:"effectivePeriod,omitempty" bson:"effective_period,omitempty"`                   // When the plan definition is expected to be used
	Topic                   []CodeableConcept      `json:"topic,omitempty" bson:"topic,omitempty"`                                        // E.g. Education, Treatment, Assessment
	Author                  []ContactDetail        `json:"author,omitempty" bson:"author,omitempty"`                                      // Who authored the content
	Editor                  []ContactDetail        `json:"editor,omitempty" bson:"editor,omitempty"`                                      // Who edited the content
	Reviewer                []ContactDetail        `json:"reviewer,omitempty" bson:"reviewer,omitempty"`                                  // Who reviewed the content
	Endorser                []ContactDetail        `json:"endorser,omitempty" bson:"endorser,omitempty"`                                  // Who endorsed the content
	RelatedArtifact         []RelatedArtifact      `json:"relatedArtifact,omitempty" bson:"related_artifact,omitempty"`                   // Additional documentation, citations
	Library                 []string               `json:"library,omitempty" bson:"library,omitempty"`                                    // Logic used by the plan definition
	Goal                    []PlanDefinitionGoal   `json:"goal,omitempty" bson:"goal,omitempty"`                                          // What the plan is trying to accomplish
	Actor                   []PlanDefinitionActor  `json:"actor,omitempty" bson:"actor,omitempty"`                                        // Actors within the plan
	Action                  []PlanDefinitionAction `json:"action,omitempty" bson:"action,omitempty"`                                      // Action defined by the plan
	AsNeededBoolean         *bool                  `json:"asNeededBoolean,omitempty" bson:"as_needed_boolean,omitempty"`                  // Preconditions for service
	AsNeededCodeableConcept *CodeableConcept       `json:"asNeededCodeableConcept,omitempty" bson:"as_needed_codeable_concept,omitempty"` // Preconditions for service
}

func (r *PlanDefinition) Validate() error {
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
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.SubjectCodeableConcept != nil {
		if err := r.SubjectCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("SubjectCodeableConcept: %w", err)
		}
	}
	if r.SubjectReference != nil {
		if err := r.SubjectReference.Validate(); err != nil {
			return fmt.Errorf("SubjectReference: %w", err)
		}
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
	for i, item := range r.Topic {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Topic[%d]: %w", i, err)
		}
	}
	for i, item := range r.Author {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Author[%d]: %w", i, err)
		}
	}
	for i, item := range r.Editor {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Editor[%d]: %w", i, err)
		}
	}
	for i, item := range r.Reviewer {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reviewer[%d]: %w", i, err)
		}
	}
	for i, item := range r.Endorser {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Endorser[%d]: %w", i, err)
		}
	}
	for i, item := range r.RelatedArtifact {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RelatedArtifact[%d]: %w", i, err)
		}
	}
	for i, item := range r.Goal {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Goal[%d]: %w", i, err)
		}
	}
	for i, item := range r.Actor {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Actor[%d]: %w", i, err)
		}
	}
	for i, item := range r.Action {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Action[%d]: %w", i, err)
		}
	}
	if r.AsNeededCodeableConcept != nil {
		if err := r.AsNeededCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("AsNeededCodeableConcept: %w", err)
		}
	}
	return nil
}

type PlanDefinitionActionOutput struct {
	Id          *string          `json:"id,omitempty" bson:"id,omitempty"`                    // Unique id for inter-element referencing
	Title       *string          `json:"title,omitempty" bson:"title,omitempty"`              // User-visible title
	Requirement *DataRequirement `json:"requirement,omitempty" bson:"requirement,omitempty"`  // What data is provided
	RelatedData *string          `json:"relatedData,omitempty" bson:"related_data,omitempty"` // What data is provided
}

func (r *PlanDefinitionActionOutput) Validate() error {
	if r.Requirement != nil {
		if err := r.Requirement.Validate(); err != nil {
			return fmt.Errorf("Requirement: %w", err)
		}
	}
	return nil
}

type PlanDefinitionActionParticipant struct {
	Id            *string          `json:"id,omitempty" bson:"id,omitempty"`                        // Unique id for inter-element referencing
	ActorId       *string          `json:"actorId,omitempty" bson:"actor_id,omitempty"`             // What actor
	Type          *string          `json:"type,omitempty" bson:"type,omitempty"`                    // careteam | device | group | healthcareservice | location | organization | patient | practitioner | practitionerrole | relatedperson
	TypeCanonical *string          `json:"typeCanonical,omitempty" bson:"type_canonical,omitempty"` // Who or what can participate
	TypeReference *Reference       `json:"typeReference,omitempty" bson:"type_reference,omitempty"` // Who or what can participate
	Role          *CodeableConcept `json:"role,omitempty" bson:"role,omitempty"`                    // E.g. Nurse, Surgeon, Parent
	Function      *CodeableConcept `json:"function,omitempty" bson:"function,omitempty"`            // E.g. Author, Reviewer, Witness, etc
}

func (r *PlanDefinitionActionParticipant) Validate() error {
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
	return nil
}

type PlanDefinitionActionDynamicValue struct {
	Id         *string     `json:"id,omitempty" bson:"id,omitempty"`                 // Unique id for inter-element referencing
	Path       *string     `json:"path,omitempty" bson:"path,omitempty"`             // The path to the element to be set dynamically
	Expression *Expression `json:"expression,omitempty" bson:"expression,omitempty"` // An expression that provides the dynamic value for the customization
}

func (r *PlanDefinitionActionDynamicValue) Validate() error {
	if r.Expression != nil {
		if err := r.Expression.Validate(); err != nil {
			return fmt.Errorf("Expression: %w", err)
		}
	}
	return nil
}

type PlanDefinitionGoal struct {
	Id            *string                    `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	Category      *CodeableConcept           `json:"category,omitempty" bson:"category,omitempty"`           // E.g. Treatment, dietary, behavioral
	Description   *CodeableConcept           `json:"description" bson:"description"`                         // Code or text describing the goal
	Priority      *CodeableConcept           `json:"priority,omitempty" bson:"priority,omitempty"`           // high-priority | medium-priority | low-priority
	Start         *CodeableConcept           `json:"start,omitempty" bson:"start,omitempty"`                 // When goal pursuit begins
	Addresses     []CodeableConcept          `json:"addresses,omitempty" bson:"addresses,omitempty"`         // What does the goal address
	Documentation []RelatedArtifact          `json:"documentation,omitempty" bson:"documentation,omitempty"` // Supporting documentation for the goal
	Target        []PlanDefinitionGoalTarget `json:"target,omitempty" bson:"target,omitempty"`               // Target outcome for the goal
}

func (r *PlanDefinitionGoal) Validate() error {
	if r.Category != nil {
		if err := r.Category.Validate(); err != nil {
			return fmt.Errorf("Category: %w", err)
		}
	}
	if r.Description == nil {
		return fmt.Errorf("field 'Description' is required")
	}
	if r.Description != nil {
		if err := r.Description.Validate(); err != nil {
			return fmt.Errorf("Description: %w", err)
		}
	}
	if r.Priority != nil {
		if err := r.Priority.Validate(); err != nil {
			return fmt.Errorf("Priority: %w", err)
		}
	}
	if r.Start != nil {
		if err := r.Start.Validate(); err != nil {
			return fmt.Errorf("Start: %w", err)
		}
	}
	for i, item := range r.Addresses {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Addresses[%d]: %w", i, err)
		}
	}
	for i, item := range r.Documentation {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Documentation[%d]: %w", i, err)
		}
	}
	for i, item := range r.Target {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Target[%d]: %w", i, err)
		}
	}
	return nil
}

type PlanDefinitionActor struct {
	Id          *string                     `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Title       *string                     `json:"title,omitempty" bson:"title,omitempty"`             // User-visible title
	Description *string                     `json:"description,omitempty" bson:"description,omitempty"` // Describes the actor
	Option      []PlanDefinitionActorOption `json:"option" bson:"option"`                               // Who or what can be this actor
}

func (r *PlanDefinitionActor) Validate() error {
	if len(r.Option) < 1 {
		return fmt.Errorf("field 'Option' must have at least 1 elements")
	}
	for i, item := range r.Option {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Option[%d]: %w", i, err)
		}
	}
	return nil
}

type PlanDefinitionActionCondition struct {
	Id         *string     `json:"id,omitempty" bson:"id,omitempty"`                 // Unique id for inter-element referencing
	Kind       string      `json:"kind" bson:"kind"`                                 // applicability | start | stop
	Expression *Expression `json:"expression,omitempty" bson:"expression,omitempty"` // Boolean-valued expression
}

func (r *PlanDefinitionActionCondition) Validate() error {
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

type PlanDefinitionActionInput struct {
	Id          *string          `json:"id,omitempty" bson:"id,omitempty"`                    // Unique id for inter-element referencing
	Title       *string          `json:"title,omitempty" bson:"title,omitempty"`              // User-visible title
	Requirement *DataRequirement `json:"requirement,omitempty" bson:"requirement,omitempty"`  // What data is provided
	RelatedData *string          `json:"relatedData,omitempty" bson:"related_data,omitempty"` // What data is provided
}

func (r *PlanDefinitionActionInput) Validate() error {
	if r.Requirement != nil {
		if err := r.Requirement.Validate(); err != nil {
			return fmt.Errorf("Requirement: %w", err)
		}
	}
	return nil
}

type PlanDefinitionActionRelatedAction struct {
	Id              *string   `json:"id,omitempty" bson:"id,omitempty"`                            // Unique id for inter-element referencing
	TargetId        string    `json:"targetId" bson:"target_id"`                                   // What action is this related to
	Relationship    string    `json:"relationship" bson:"relationship"`                            // before | before-start | before-end | concurrent | concurrent-with-start | concurrent-with-end | after | after-start | after-end
	EndRelationship *string   `json:"endRelationship,omitempty" bson:"end_relationship,omitempty"` // before | before-start | before-end | concurrent | concurrent-with-start | concurrent-with-end | after | after-start | after-end
	OffsetDuration  *Duration `json:"offsetDuration,omitempty" bson:"offset_duration,omitempty"`   // Time offset for the relationship
	OffsetRange     *Range    `json:"offsetRange,omitempty" bson:"offset_range,omitempty"`         // Time offset for the relationship
}

func (r *PlanDefinitionActionRelatedAction) Validate() error {
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

type PlanDefinitionGoalTarget struct {
	Id                    *string          `json:"id,omitempty" bson:"id,omitempty"`                                         // Unique id for inter-element referencing
	Measure               *CodeableConcept `json:"measure,omitempty" bson:"measure,omitempty"`                               // The parameter whose value is to be tracked
	DetailQuantity        *Quantity        `json:"detailQuantity,omitempty" bson:"detail_quantity,omitempty"`                // The target value to be achieved
	DetailRange           *Range           `json:"detailRange,omitempty" bson:"detail_range,omitempty"`                      // The target value to be achieved
	DetailCodeableConcept *CodeableConcept `json:"detailCodeableConcept,omitempty" bson:"detail_codeable_concept,omitempty"` // The target value to be achieved
	DetailString          *string          `json:"detailString,omitempty" bson:"detail_string,omitempty"`                    // The target value to be achieved
	DetailBoolean         *bool            `json:"detailBoolean,omitempty" bson:"detail_boolean,omitempty"`                  // The target value to be achieved
	DetailInteger         *int             `json:"detailInteger,omitempty" bson:"detail_integer,omitempty"`                  // The target value to be achieved
	DetailRatio           *Ratio           `json:"detailRatio,omitempty" bson:"detail_ratio,omitempty"`                      // The target value to be achieved
	Due                   *Duration        `json:"due,omitempty" bson:"due,omitempty"`                                       // Reach goal within
}

func (r *PlanDefinitionGoalTarget) Validate() error {
	if r.Measure != nil {
		if err := r.Measure.Validate(); err != nil {
			return fmt.Errorf("Measure: %w", err)
		}
	}
	if r.DetailQuantity != nil {
		if err := r.DetailQuantity.Validate(); err != nil {
			return fmt.Errorf("DetailQuantity: %w", err)
		}
	}
	if r.DetailRange != nil {
		if err := r.DetailRange.Validate(); err != nil {
			return fmt.Errorf("DetailRange: %w", err)
		}
	}
	if r.DetailCodeableConcept != nil {
		if err := r.DetailCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("DetailCodeableConcept: %w", err)
		}
	}
	if r.DetailRatio != nil {
		if err := r.DetailRatio.Validate(); err != nil {
			return fmt.Errorf("DetailRatio: %w", err)
		}
	}
	if r.Due != nil {
		if err := r.Due.Validate(); err != nil {
			return fmt.Errorf("Due: %w", err)
		}
	}
	return nil
}

type PlanDefinitionActorOption struct {
	Id            *string          `json:"id,omitempty" bson:"id,omitempty"`                        // Unique id for inter-element referencing
	Type          *string          `json:"type,omitempty" bson:"type,omitempty"`                    // careteam | device | group | healthcareservice | location | organization | patient | practitioner | practitionerrole | relatedperson
	TypeCanonical *string          `json:"typeCanonical,omitempty" bson:"type_canonical,omitempty"` // Who or what can participate
	TypeReference *Reference       `json:"typeReference,omitempty" bson:"type_reference,omitempty"` // Who or what can participate
	Role          *CodeableConcept `json:"role,omitempty" bson:"role,omitempty"`                    // E.g. Nurse, Surgeon, Parent
}

func (r *PlanDefinitionActorOption) Validate() error {
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
	return nil
}

type PlanDefinitionAction struct {
	Id                     *string                             `json:"id,omitempty" bson:"id,omitempty"`                                           // Unique id for inter-element referencing
	LinkId                 *string                             `json:"linkId,omitempty" bson:"link_id,omitempty"`                                  // Unique id for the action in the PlanDefinition
	Prefix                 *string                             `json:"prefix,omitempty" bson:"prefix,omitempty"`                                   // User-visible prefix for the action (e.g. 1. or A.)
	Title                  *string                             `json:"title,omitempty" bson:"title,omitempty"`                                     // User-visible title
	Description            *string                             `json:"description,omitempty" bson:"description,omitempty"`                         // Brief description of the action
	TextEquivalent         *string                             `json:"textEquivalent,omitempty" bson:"text_equivalent,omitempty"`                  // Static text equivalent of the action, used if the dynamic aspects cannot be interpreted by the receiving system
	Priority               *string                             `json:"priority,omitempty" bson:"priority,omitempty"`                               // routine | urgent | asap | stat
	Code                   *CodeableConcept                    `json:"code,omitempty" bson:"code,omitempty"`                                       // Code representing the meaning of the action or sub-actions
	Reason                 []CodeableConcept                   `json:"reason,omitempty" bson:"reason,omitempty"`                                   // Why the action should be performed
	Documentation          []RelatedArtifact                   `json:"documentation,omitempty" bson:"documentation,omitempty"`                     // Supporting documentation for the intended performer of the action
	GoalId                 []string                            `json:"goalId,omitempty" bson:"goal_id,omitempty"`                                  // What goals this action supports
	SubjectCodeableConcept *CodeableConcept                    `json:"subjectCodeableConcept,omitempty" bson:"subject_codeable_concept,omitempty"` // Type of individual the action is focused on
	SubjectReference       *Reference                          `json:"subjectReference,omitempty" bson:"subject_reference,omitempty"`              // Type of individual the action is focused on
	SubjectCanonical       *string                             `json:"subjectCanonical,omitempty" bson:"subject_canonical,omitempty"`              // Type of individual the action is focused on
	Trigger                []TriggerDefinition                 `json:"trigger,omitempty" bson:"trigger,omitempty"`                                 // When the action should be triggered
	Condition              []PlanDefinitionActionCondition     `json:"condition,omitempty" bson:"condition,omitempty"`                             // Whether or not the action is applicable
	Input                  []PlanDefinitionActionInput         `json:"input,omitempty" bson:"input,omitempty"`                                     // Input data requirements
	Output                 []PlanDefinitionActionOutput        `json:"output,omitempty" bson:"output,omitempty"`                                   // Output data definition
	RelatedAction          []PlanDefinitionActionRelatedAction `json:"relatedAction,omitempty" bson:"related_action,omitempty"`                    // Relationship to another action
	TimingAge              *Age                                `json:"timingAge,omitempty" bson:"timing_age,omitempty"`                            // When the action should take place
	TimingDuration         *Duration                           `json:"timingDuration,omitempty" bson:"timing_duration,omitempty"`                  // When the action should take place
	TimingRange            *Range                              `json:"timingRange,omitempty" bson:"timing_range,omitempty"`                        // When the action should take place
	TimingTiming           *Timing                             `json:"timingTiming,omitempty" bson:"timing_timing,omitempty"`                      // When the action should take place
	TimingRelativeTime     *RelativeTime                       `json:"timingRelativeTime,omitempty" bson:"timing_relative_time,omitempty"`         // When the action should take place
	Location               *CodeableReference                  `json:"location,omitempty" bson:"location,omitempty"`                               // Where it should happen
	Participant            []PlanDefinitionActionParticipant   `json:"participant,omitempty" bson:"participant,omitempty"`                         // Who should participate in the action
	Type                   *CodeableConcept                    `json:"type,omitempty" bson:"type,omitempty"`                                       // create | update | remove | fire-event | etc.
	ApplicabilityBehavior  *string                             `json:"applicabilityBehavior,omitempty" bson:"applicability_behavior,omitempty"`    // all | any
	GroupingBehavior       *string                             `json:"groupingBehavior,omitempty" bson:"grouping_behavior,omitempty"`              // visual-group | logical-group | sentence-group
	SelectionBehavior      *string                             `json:"selectionBehavior,omitempty" bson:"selection_behavior,omitempty"`            // any | all | all-or-none | exactly-one | at-most-one | one-or-more
	RequiredBehavior       *string                             `json:"requiredBehavior,omitempty" bson:"required_behavior,omitempty"`              // must | could | must-unless-documented
	PrecheckBehavior       *string                             `json:"precheckBehavior,omitempty" bson:"precheck_behavior,omitempty"`              // yes | no
	CardinalityBehavior    *string                             `json:"cardinalityBehavior,omitempty" bson:"cardinality_behavior,omitempty"`        // single | multiple
	DefinitionCanonical    *string                             `json:"definitionCanonical,omitempty" bson:"definition_canonical,omitempty"`        // Description of the activity to be performed
	DefinitionUri          *string                             `json:"definitionUri,omitempty" bson:"definition_uri,omitempty"`                    // Description of the activity to be performed
	Transform              *string                             `json:"transform,omitempty" bson:"transform,omitempty"`                             // Transform to apply the template
	DynamicValue           []PlanDefinitionActionDynamicValue  `json:"dynamicValue,omitempty" bson:"dynamic_value,omitempty"`                      // Dynamic aspects of the definition
	Action                 []PlanDefinitionAction              `json:"action,omitempty" bson:"action,omitempty"`                                   // A sub-action
}

func (r *PlanDefinitionAction) Validate() error {
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	for i, item := range r.Reason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reason[%d]: %w", i, err)
		}
	}
	for i, item := range r.Documentation {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Documentation[%d]: %w", i, err)
		}
	}
	if r.SubjectCodeableConcept != nil {
		if err := r.SubjectCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("SubjectCodeableConcept: %w", err)
		}
	}
	if r.SubjectReference != nil {
		if err := r.SubjectReference.Validate(); err != nil {
			return fmt.Errorf("SubjectReference: %w", err)
		}
	}
	for i, item := range r.Trigger {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Trigger[%d]: %w", i, err)
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
