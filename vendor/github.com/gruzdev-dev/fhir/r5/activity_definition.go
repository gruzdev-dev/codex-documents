package models

import (
	"encoding/json"
	"fmt"
)

// This resource allows for the definition of some activity to be performed, independent of a particular patient, practitioner, or other performance context.
type ActivityDefinition struct {
	ResourceType                 string                           `json:"resourceType" bson:"resource_type"`                                                      // Type of resource
	Id                           *string                          `json:"id,omitempty" bson:"id,omitempty"`                                                       // Logical id of this artifact
	Meta                         *Meta                            `json:"meta,omitempty" bson:"meta,omitempty"`                                                   // Metadata about the resource
	ImplicitRules                *string                          `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                                // A set of rules under which this content was created
	Language                     *string                          `json:"language,omitempty" bson:"language,omitempty"`                                           // Language of the resource content
	Text                         *Narrative                       `json:"text,omitempty" bson:"text,omitempty"`                                                   // Text summary of the resource, for human interpretation
	Contained                    []json.RawMessage                `json:"contained,omitempty" bson:"contained,omitempty"`                                         // Contained, inline Resources
	Url                          *string                          `json:"url,omitempty" bson:"url,omitempty"`                                                     // Canonical identifier for this activity definition, represented as a URI (globally unique)
	Identifier                   []Identifier                     `json:"identifier,omitempty" bson:"identifier,omitempty"`                                       // Additional identifier for the activity definition
	Version                      *string                          `json:"version,omitempty" bson:"version,omitempty"`                                             // Business version of the activity definition
	VersionAlgorithmString       *string                          `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"`             // How to compare versions
	VersionAlgorithmCoding       *Coding                          `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"`             // How to compare versions
	Name                         *string                          `json:"name,omitempty" bson:"name,omitempty"`                                                   // Name for this activity definition (computer friendly)
	Title                        *string                          `json:"title,omitempty" bson:"title,omitempty"`                                                 // Name for this activity definition (human friendly)
	Subtitle                     *string                          `json:"subtitle,omitempty" bson:"subtitle,omitempty"`                                           // Subordinate title of the activity definition
	Status                       string                           `json:"status" bson:"status"`                                                                   // draft | active | retired | unknown
	Experimental                 bool                             `json:"experimental,omitempty" bson:"experimental,omitempty"`                                   // For testing only - never for real usage
	SubjectCodeableConcept       *CodeableConcept                 `json:"subjectCodeableConcept,omitempty" bson:"subject_codeable_concept,omitempty"`             // Type of individual the activity definition is intended for
	SubjectReference             *Reference                       `json:"subjectReference,omitempty" bson:"subject_reference,omitempty"`                          // Type of individual the activity definition is intended for
	SubjectCanonical             *string                          `json:"subjectCanonical,omitempty" bson:"subject_canonical,omitempty"`                          // Type of individual the activity definition is intended for
	Date                         *string                          `json:"date,omitempty" bson:"date,omitempty"`                                                   // Date last changed
	Publisher                    *string                          `json:"publisher,omitempty" bson:"publisher,omitempty"`                                         // Name of the publisher/steward (organization or individual)
	Contact                      []ContactDetail                  `json:"contact,omitempty" bson:"contact,omitempty"`                                             // Contact details for the publisher
	Description                  *string                          `json:"description,omitempty" bson:"description,omitempty"`                                     // Natural language description of the activity definition
	UseContext                   []UsageContext                   `json:"useContext,omitempty" bson:"use_context,omitempty"`                                      // The context that the content is intended to support
	Jurisdiction                 []CodeableConcept                `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                                   // Jurisdiction of the authority that maintains the activity definition (if applicable)
	Purpose                      *string                          `json:"purpose,omitempty" bson:"purpose,omitempty"`                                             // Why this activity definition is defined
	Usage                        *string                          `json:"usage,omitempty" bson:"usage,omitempty"`                                                 // Describes the clinical usage of the activity definition
	Copyright                    *string                          `json:"copyright,omitempty" bson:"copyright,omitempty"`                                         // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel               *string                          `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                              // Copyright holder and year(s)
	ApprovalDate                 *string                          `json:"approvalDate,omitempty" bson:"approval_date,omitempty"`                                  // When the activity definition was approved by publisher
	LastReviewDate               *string                          `json:"lastReviewDate,omitempty" bson:"last_review_date,omitempty"`                             // When the activity definition was last reviewed by the publisher
	EffectivePeriod              *Period                          `json:"effectivePeriod,omitempty" bson:"effective_period,omitempty"`                            // When the activity definition is expected to be used
	Topic                        []CodeableConcept                `json:"topic,omitempty" bson:"topic,omitempty"`                                                 // E.g. Education, Treatment, Assessment, etc
	Author                       []ContactDetail                  `json:"author,omitempty" bson:"author,omitempty"`                                               // Who authored the content
	Editor                       []ContactDetail                  `json:"editor,omitempty" bson:"editor,omitempty"`                                               // Who edited the content
	Reviewer                     []ContactDetail                  `json:"reviewer,omitempty" bson:"reviewer,omitempty"`                                           // Who reviewed the content
	Endorser                     []ContactDetail                  `json:"endorser,omitempty" bson:"endorser,omitempty"`                                           // Who endorsed the content
	RelatedArtifact              []RelatedArtifact                `json:"relatedArtifact,omitempty" bson:"related_artifact,omitempty"`                            // Additional documentation, citations, etc
	Library                      []string                         `json:"library,omitempty" bson:"library,omitempty"`                                             // Logic used by the activity definition
	Kind                         *string                          `json:"kind,omitempty" bson:"kind,omitempty"`                                                   // Kind of resource
	Profile                      *string                          `json:"profile,omitempty" bson:"profile,omitempty"`                                             // What profile the resource needs to conform to
	Code                         *CodeableConcept                 `json:"code,omitempty" bson:"code,omitempty"`                                                   // Detail type of activity
	Intent                       *string                          `json:"intent,omitempty" bson:"intent,omitempty"`                                               // proposal | solicit-offer | offer-response | plan | directive | order | original-order | reflex-order | filler-order | instance-order | option
	Priority                     *string                          `json:"priority,omitempty" bson:"priority,omitempty"`                                           // routine | urgent | asap | stat
	DoNotPerform                 bool                             `json:"doNotPerform,omitempty" bson:"do_not_perform,omitempty"`                                 // True if the activity should not be performed
	TimingTiming                 *Timing                          `json:"timingTiming,omitempty" bson:"timing_timing,omitempty"`                                  // When activity is to occur
	TimingAge                    *Age                             `json:"timingAge,omitempty" bson:"timing_age,omitempty"`                                        // When activity is to occur
	TimingRange                  *Range                           `json:"timingRange,omitempty" bson:"timing_range,omitempty"`                                    // When activity is to occur
	TimingDuration               *Duration                        `json:"timingDuration,omitempty" bson:"timing_duration,omitempty"`                              // When activity is to occur
	TimingRelativeTime           *RelativeTime                    `json:"timingRelativeTime,omitempty" bson:"timing_relative_time,omitempty"`                     // When activity is to occur
	AsNeededBoolean              *bool                            `json:"asNeededBoolean,omitempty" bson:"as_needed_boolean,omitempty"`                           // Preconditions for service
	AsNeededCodeableConcept      *CodeableConcept                 `json:"asNeededCodeableConcept,omitempty" bson:"as_needed_codeable_concept,omitempty"`          // Preconditions for service
	Location                     *CodeableReference               `json:"location,omitempty" bson:"location,omitempty"`                                           // Where it should happen
	Participant                  []ActivityDefinitionParticipant  `json:"participant,omitempty" bson:"participant,omitempty"`                                     // Who should participate in the action
	ProductReference             *Reference                       `json:"productReference,omitempty" bson:"product_reference,omitempty"`                          // What's administered/supplied
	ProductCodeableConcept       *CodeableConcept                 `json:"productCodeableConcept,omitempty" bson:"product_codeable_concept,omitempty"`             // What's administered/supplied
	Quantity                     *Quantity                        `json:"quantity,omitempty" bson:"quantity,omitempty"`                                           // How much is administered/consumed/supplied
	Dosage                       []Dosage                         `json:"dosage,omitempty" bson:"dosage,omitempty"`                                               // Detailed dosage instructions
	BodySite                     []CodeableConcept                `json:"bodySite,omitempty" bson:"body_site,omitempty"`                                          // What part of body to perform on
	SpecimenRequirement          []string                         `json:"specimenRequirement,omitempty" bson:"specimen_requirement,omitempty"`                    // What specimens are required to perform this action
	ObservationRequirement       []string                         `json:"observationRequirement,omitempty" bson:"observation_requirement,omitempty"`              // What observations are required to perform this action
	ObservationResultRequirement []string                         `json:"observationResultRequirement,omitempty" bson:"observation_result_requirement,omitempty"` // What observations must be produced by this action
	Transform                    *string                          `json:"transform,omitempty" bson:"transform,omitempty"`                                         // Transform to apply the template
	DynamicValue                 []ActivityDefinitionDynamicValue `json:"dynamicValue,omitempty" bson:"dynamic_value,omitempty"`                                  // Dynamic aspects of the definition
}

func (r *ActivityDefinition) Validate() error {
	if r.ResourceType != "ActivityDefinition" {
		return fmt.Errorf("invalid resourceType: expected 'ActivityDefinition', got '%s'", r.ResourceType)
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
	if r.VersionAlgorithmCoding != nil {
		if err := r.VersionAlgorithmCoding.Validate(); err != nil {
			return fmt.Errorf("VersionAlgorithmCoding: %w", err)
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
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.TimingTiming != nil {
		if err := r.TimingTiming.Validate(); err != nil {
			return fmt.Errorf("TimingTiming: %w", err)
		}
	}
	if r.TimingAge != nil {
		if err := r.TimingAge.Validate(); err != nil {
			return fmt.Errorf("TimingAge: %w", err)
		}
	}
	if r.TimingRange != nil {
		if err := r.TimingRange.Validate(); err != nil {
			return fmt.Errorf("TimingRange: %w", err)
		}
	}
	if r.TimingDuration != nil {
		if err := r.TimingDuration.Validate(); err != nil {
			return fmt.Errorf("TimingDuration: %w", err)
		}
	}
	if r.TimingRelativeTime != nil {
		if err := r.TimingRelativeTime.Validate(); err != nil {
			return fmt.Errorf("TimingRelativeTime: %w", err)
		}
	}
	if r.AsNeededCodeableConcept != nil {
		if err := r.AsNeededCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("AsNeededCodeableConcept: %w", err)
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
	if r.ProductReference != nil {
		if err := r.ProductReference.Validate(); err != nil {
			return fmt.Errorf("ProductReference: %w", err)
		}
	}
	if r.ProductCodeableConcept != nil {
		if err := r.ProductCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("ProductCodeableConcept: %w", err)
		}
	}
	if r.Quantity != nil {
		if err := r.Quantity.Validate(); err != nil {
			return fmt.Errorf("Quantity: %w", err)
		}
	}
	for i, item := range r.Dosage {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Dosage[%d]: %w", i, err)
		}
	}
	for i, item := range r.BodySite {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("BodySite[%d]: %w", i, err)
		}
	}
	for i, item := range r.DynamicValue {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("DynamicValue[%d]: %w", i, err)
		}
	}
	return nil
}

type ActivityDefinitionParticipant struct {
	Id            *string          `json:"id,omitempty" bson:"id,omitempty"`                        // Unique id for inter-element referencing
	Type          *string          `json:"type,omitempty" bson:"type,omitempty"`                    // careteam | device | group | healthcareservice | location | organization | patient | practitioner | practitionerrole | relatedperson
	TypeCanonical *string          `json:"typeCanonical,omitempty" bson:"type_canonical,omitempty"` // Who or what can participate
	TypeReference *Reference       `json:"typeReference,omitempty" bson:"type_reference,omitempty"` // Who or what can participate
	Role          *CodeableConcept `json:"role,omitempty" bson:"role,omitempty"`                    // E.g. Nurse, Surgeon, Parent, etc
	Function      *CodeableConcept `json:"function,omitempty" bson:"function,omitempty"`            // E.g. Author, Reviewer, Witness, etc
}

func (r *ActivityDefinitionParticipant) Validate() error {
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

type ActivityDefinitionDynamicValue struct {
	Id         *string     `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Path       string      `json:"path" bson:"path"`                 // The path to the element to be set dynamically
	Expression *Expression `json:"expression" bson:"expression"`     // An expression that provides the dynamic value for the customization
}

func (r *ActivityDefinitionDynamicValue) Validate() error {
	var emptyString string
	if r.Path == emptyString {
		return fmt.Errorf("field 'Path' is required")
	}
	if r.Expression == nil {
		return fmt.Errorf("field 'Expression' is required")
	}
	if r.Expression != nil {
		if err := r.Expression.Validate(); err != nil {
			return fmt.Errorf("Expression: %w", err)
		}
	}
	return nil
}
