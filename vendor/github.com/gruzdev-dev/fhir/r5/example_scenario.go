package models

import (
	"encoding/json"
	"fmt"
)

// A walkthrough of a workflow showing the interaction between systems and the instances shared, possibly including the evolution of instances over time.
type ExampleScenario struct {
	Id                     *string                   `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta                     `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string                   `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string                   `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative                `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage         `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Url                    *string                   `json:"url,omitempty" bson:"url,omitempty"`                                         // Canonical identifier for this example scenario, represented as a URI (globally unique)
	Identifier             []Identifier              `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // Additional identifier for the example scenario
	Version                *string                   `json:"version,omitempty" bson:"version,omitempty"`                                 // Business version of the example scenario
	VersionAlgorithmString *string                   `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"` // How to compare versions
	VersionAlgorithmCoding *Coding                   `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"` // How to compare versions
	Name                   *string                   `json:"name,omitempty" bson:"name,omitempty"`                                       // Name for this example scenario (computer friendly)
	Title                  *string                   `json:"title,omitempty" bson:"title,omitempty"`                                     // Name for this example scenario (human friendly)
	Status                 string                    `json:"status" bson:"status"`                                                       // draft | active | retired | unknown
	Experimental           bool                      `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // For testing only - never for real usage
	Date                   *string                   `json:"date,omitempty" bson:"date,omitempty"`                                       // Date last changed
	Publisher              *string                   `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // Name of the publisher/steward (organization or individual)
	Contact                []ContactDetail           `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Description            *string                   `json:"description,omitempty" bson:"description,omitempty"`                         // Natural language description of the ExampleScenario
	UseContext             []UsageContext            `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // The context that the content is intended to support
	Jurisdiction           []CodeableConcept         `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                       // Jurisdiction of the authority that maintains the example scenario (if applicable)
	Purpose                *string                   `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // The purpose of the example, e.g. to illustrate a scenario
	Copyright              *string                   `json:"copyright,omitempty" bson:"copyright,omitempty"`                             // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel         *string                   `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                  // Copyright holder and year(s)
	Actor                  []ExampleScenarioActor    `json:"actor,omitempty" bson:"actor,omitempty"`                                     // Individual involved in exchange
	Instance               []ExampleScenarioInstance `json:"instance,omitempty" bson:"instance,omitempty"`                               // Data used in the scenario
	Process                []ExampleScenarioProcess  `json:"process,omitempty" bson:"process,omitempty"`                                 // Major process within scenario
}

func (r *ExampleScenario) Validate() error {
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
	for i, item := range r.Actor {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Actor[%d]: %w", i, err)
		}
	}
	for i, item := range r.Instance {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Instance[%d]: %w", i, err)
		}
	}
	for i, item := range r.Process {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Process[%d]: %w", i, err)
		}
	}
	return nil
}

type ExampleScenarioProcessStep struct {
	Id          *string                                 `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Number      *string                                 `json:"number,omitempty" bson:"number,omitempty"`           // Sequential number of the step
	Process     *ExampleScenarioProcess                 `json:"process,omitempty" bson:"process,omitempty"`         // Step is nested process
	Workflow    *string                                 `json:"workflow,omitempty" bson:"workflow,omitempty"`       // Step is nested workflow
	Operation   *ExampleScenarioProcessStepOperation    `json:"operation,omitempty" bson:"operation,omitempty"`     // Step is simple action
	Alternative []ExampleScenarioProcessStepAlternative `json:"alternative,omitempty" bson:"alternative,omitempty"` // Alternate non-typical step action
	Pause       bool                                    `json:"pause,omitempty" bson:"pause,omitempty"`             // Pause in the flow?
}

func (r *ExampleScenarioProcessStep) Validate() error {
	if r.Process != nil {
		if err := r.Process.Validate(); err != nil {
			return fmt.Errorf("Process: %w", err)
		}
	}
	if r.Operation != nil {
		if err := r.Operation.Validate(); err != nil {
			return fmt.Errorf("Operation: %w", err)
		}
	}
	for i, item := range r.Alternative {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Alternative[%d]: %w", i, err)
		}
	}
	return nil
}

type ExampleScenarioProcessStepOperation struct {
	Id              *string                                   `json:"id,omitempty" bson:"id,omitempty"`                            // Unique id for inter-element referencing
	Type            *Coding                                   `json:"type,omitempty" bson:"type,omitempty"`                        // Kind of action
	Title           string                                    `json:"title" bson:"title"`                                          // Label for step
	Initiator       *string                                   `json:"initiator,omitempty" bson:"initiator,omitempty"`              // Who starts the operation
	Receiver        *string                                   `json:"receiver,omitempty" bson:"receiver,omitempty"`                // Who receives the operation
	Description     *string                                   `json:"description,omitempty" bson:"description,omitempty"`          // Human-friendly description of the operation
	InitiatorActive bool                                      `json:"initiatorActive,omitempty" bson:"initiator_active,omitempty"` // Initiator stays active?
	ReceiverActive  bool                                      `json:"receiverActive,omitempty" bson:"receiver_active,omitempty"`   // Receiver stays active?
	Request         *ExampleScenarioInstanceContainedInstance `json:"request,omitempty" bson:"request,omitempty"`                  // Instance transmitted on invocation
	Response        *ExampleScenarioInstanceContainedInstance `json:"response,omitempty" bson:"response,omitempty"`                // Instance transmitted on invocation response
}

func (r *ExampleScenarioProcessStepOperation) Validate() error {
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	var emptyString string
	if r.Title == emptyString {
		return fmt.Errorf("field 'Title' is required")
	}
	if r.Request != nil {
		if err := r.Request.Validate(); err != nil {
			return fmt.Errorf("Request: %w", err)
		}
	}
	if r.Response != nil {
		if err := r.Response.Validate(); err != nil {
			return fmt.Errorf("Response: %w", err)
		}
	}
	return nil
}

type ExampleScenarioProcessStepAlternative struct {
	Id          *string                      `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Title       string                       `json:"title" bson:"title"`                                 // Label for alternative
	Description *string                      `json:"description,omitempty" bson:"description,omitempty"` // Human-readable description of option
	Step        []ExampleScenarioProcessStep `json:"step,omitempty" bson:"step,omitempty"`               // Alternative action(s)
}

func (r *ExampleScenarioProcessStepAlternative) Validate() error {
	var emptyString string
	if r.Title == emptyString {
		return fmt.Errorf("field 'Title' is required")
	}
	for i, item := range r.Step {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Step[%d]: %w", i, err)
		}
	}
	return nil
}

type ExampleScenarioActor struct {
	Id          *string `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Key         string  `json:"key" bson:"key"`                                     // ID or acronym of the actor
	Type        *string `json:"type,omitempty" bson:"type,omitempty"`               // person | system | collective | other
	Title       string  `json:"title" bson:"title"`                                 // Label for actor when rendering
	Description *string `json:"description,omitempty" bson:"description,omitempty"` // Details about actor
	Definition  *string `json:"definition,omitempty" bson:"definition,omitempty"`   // Formal definition of actor
}

func (r *ExampleScenarioActor) Validate() error {
	var emptyString string
	if r.Key == emptyString {
		return fmt.Errorf("field 'Key' is required")
	}
	if r.Title == emptyString {
		return fmt.Errorf("field 'Title' is required")
	}
	return nil
}

type ExampleScenarioInstance struct {
	Id                        *string                                    `json:"id,omitempty" bson:"id,omitempty"`                                                 // Unique id for inter-element referencing
	Key                       string                                     `json:"key" bson:"key"`                                                                   // ID or acronym of the instance
	StructureType             *Coding                                    `json:"structureType" bson:"structure_type"`                                              // Data structure for example
	StructureVersion          *string                                    `json:"structureVersion,omitempty" bson:"structure_version,omitempty"`                    // E.g. 4.0.1
	StructureProfileCanonical *string                                    `json:"structureProfileCanonical,omitempty" bson:"structure_profile_canonical,omitempty"` // Rules instance adheres to
	StructureProfileUri       *string                                    `json:"structureProfileUri,omitempty" bson:"structure_profile_uri,omitempty"`             // Rules instance adheres to
	Title                     string                                     `json:"title" bson:"title"`                                                               // Label for instance
	Description               *string                                    `json:"description,omitempty" bson:"description,omitempty"`                               // Human-friendly description of the instance
	Content                   *Reference                                 `json:"content,omitempty" bson:"content,omitempty"`                                       // Example instance data
	Version                   []ExampleScenarioInstanceVersion           `json:"version,omitempty" bson:"version,omitempty"`                                       // Snapshot of instance that changes
	ContainedInstance         []ExampleScenarioInstanceContainedInstance `json:"containedInstance,omitempty" bson:"contained_instance,omitempty"`                  // Resources contained in the instance
}

func (r *ExampleScenarioInstance) Validate() error {
	var emptyString string
	if r.Key == emptyString {
		return fmt.Errorf("field 'Key' is required")
	}
	if r.StructureType == nil {
		return fmt.Errorf("field 'StructureType' is required")
	}
	if r.StructureType != nil {
		if err := r.StructureType.Validate(); err != nil {
			return fmt.Errorf("StructureType: %w", err)
		}
	}
	if r.Title == emptyString {
		return fmt.Errorf("field 'Title' is required")
	}
	if r.Content != nil {
		if err := r.Content.Validate(); err != nil {
			return fmt.Errorf("Content: %w", err)
		}
	}
	for i, item := range r.Version {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Version[%d]: %w", i, err)
		}
	}
	for i, item := range r.ContainedInstance {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ContainedInstance[%d]: %w", i, err)
		}
	}
	return nil
}

type ExampleScenarioInstanceVersion struct {
	Id          *string    `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Key         string     `json:"key" bson:"key"`                                     // ID or acronym of the version
	Title       string     `json:"title" bson:"title"`                                 // Label for instance version
	Description *string    `json:"description,omitempty" bson:"description,omitempty"` // Details about version
	Content     *Reference `json:"content,omitempty" bson:"content,omitempty"`         // Example instance version data
}

func (r *ExampleScenarioInstanceVersion) Validate() error {
	var emptyString string
	if r.Key == emptyString {
		return fmt.Errorf("field 'Key' is required")
	}
	if r.Title == emptyString {
		return fmt.Errorf("field 'Title' is required")
	}
	if r.Content != nil {
		if err := r.Content.Validate(); err != nil {
			return fmt.Errorf("Content: %w", err)
		}
	}
	return nil
}

type ExampleScenarioInstanceContainedInstance struct {
	Id                *string `json:"id,omitempty" bson:"id,omitempty"`                              // Unique id for inter-element referencing
	InstanceReference string  `json:"instanceReference" bson:"instance_reference"`                   // Key of contained instance
	VersionReference  *string `json:"versionReference,omitempty" bson:"version_reference,omitempty"` // Key of contained instance version
}

func (r *ExampleScenarioInstanceContainedInstance) Validate() error {
	var emptyString string
	if r.InstanceReference == emptyString {
		return fmt.Errorf("field 'InstanceReference' is required")
	}
	return nil
}

type ExampleScenarioProcess struct {
	Id             *string                      `json:"id,omitempty" bson:"id,omitempty"`                          // Unique id for inter-element referencing
	Title          string                       `json:"title" bson:"title"`                                        // Label for procss
	Description    *string                      `json:"description,omitempty" bson:"description,omitempty"`        // Human-friendly description of the process
	PreConditions  *string                      `json:"preConditions,omitempty" bson:"pre_conditions,omitempty"`   // Status before process starts
	PostConditions *string                      `json:"postConditions,omitempty" bson:"post_conditions,omitempty"` // Status after successful completion
	Step           []ExampleScenarioProcessStep `json:"step,omitempty" bson:"step,omitempty"`                      // Event within of the process
}

func (r *ExampleScenarioProcess) Validate() error {
	var emptyString string
	if r.Title == emptyString {
		return fmt.Errorf("field 'Title' is required")
	}
	for i, item := range r.Step {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Step[%d]: %w", i, err)
		}
	}
	return nil
}
