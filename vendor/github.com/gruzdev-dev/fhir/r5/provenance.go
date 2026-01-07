package models

import (
	"encoding/json"
	"fmt"
)

// Provenance of a resource is a record that describes entities and processes involved in producing and delivering or otherwise influencing that resource. Provenance provides a critical foundation for assessing authenticity, enabling trust, and allowing reproducibility. Provenance assertions are a form of contextual metadata and can themselves become important records with their own provenance. Provenance statement indicates clinical significance in terms of confidence in authenticity, reliability, and trustworthiness, integrity, and stage in lifecycle (e.g. Document Completion - has the artifact been legally authenticated), all of which MAY impact security, privacy, and trust policies.
type Provenance struct {
	ResourceType     string              `json:"resourceType" bson:"resource_type"`                              // Type of resource
	Id               *string             `json:"id,omitempty" bson:"id,omitempty"`                               // Logical id of this artifact
	Meta             *Meta               `json:"meta,omitempty" bson:"meta,omitempty"`                           // Metadata about the resource
	ImplicitRules    *string             `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`        // A set of rules under which this content was created
	Language         *string             `json:"language,omitempty" bson:"language,omitempty"`                   // Language of the resource content
	Text             *Narrative          `json:"text,omitempty" bson:"text,omitempty"`                           // Text summary of the resource, for human interpretation
	Contained        []json.RawMessage   `json:"contained,omitempty" bson:"contained,omitempty"`                 // Contained, inline Resources
	Target           []Reference         `json:"target" bson:"target"`                                           // Target Reference(s) (usually version specific)
	OccurredPeriod   *Period             `json:"occurredPeriod,omitempty" bson:"occurred_period,omitempty"`      // When the activity occurred
	OccurredDateTime *string             `json:"occurredDateTime,omitempty" bson:"occurred_date_time,omitempty"` // When the activity occurred
	Recorded         *string             `json:"recorded,omitempty" bson:"recorded,omitempty"`                   // When the activity was recorded / updated
	Policy           []string            `json:"policy,omitempty" bson:"policy,omitempty"`                       // Policy or plan the activity was defined by
	Location         *Reference          `json:"location,omitempty" bson:"location,omitempty"`                   // Where the activity occurred
	Authorization    []CodeableReference `json:"authorization,omitempty" bson:"authorization,omitempty"`         // Authorization (purposeOfUse) related to the event
	Why              *string             `json:"why,omitempty" bson:"why,omitempty"`                             // Why was the event performed?
	Activity         *CodeableConcept    `json:"activity,omitempty" bson:"activity,omitempty"`                   // Activity that occurred
	BasedOn          []Reference         `json:"basedOn,omitempty" bson:"based_on,omitempty"`                    // Workflow authorization within which this event occurred
	Patient          *Reference          `json:"patient,omitempty" bson:"patient,omitempty"`                     // The patient is the subject of the data created/updated (.target) by the activity
	Encounter        *Reference          `json:"encounter,omitempty" bson:"encounter,omitempty"`                 // Encounter within which this event occurred or which the event is tightly associated
	Agent            []ProvenanceAgent   `json:"agent" bson:"agent"`                                             // Actor involved
	Entity           []ProvenanceEntity  `json:"entity,omitempty" bson:"entity,omitempty"`                       // An entity used in this activity
	Signature        []Signature         `json:"signature,omitempty" bson:"signature,omitempty"`                 // Signature on target
}

func (r *Provenance) Validate() error {
	if r.ResourceType != "Provenance" {
		return fmt.Errorf("invalid resourceType: expected 'Provenance', got '%s'", r.ResourceType)
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
	if len(r.Target) < 1 {
		return fmt.Errorf("field 'Target' must have at least 1 elements")
	}
	for i, item := range r.Target {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Target[%d]: %w", i, err)
		}
	}
	if r.OccurredPeriod != nil {
		if err := r.OccurredPeriod.Validate(); err != nil {
			return fmt.Errorf("OccurredPeriod: %w", err)
		}
	}
	if r.Location != nil {
		if err := r.Location.Validate(); err != nil {
			return fmt.Errorf("Location: %w", err)
		}
	}
	for i, item := range r.Authorization {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Authorization[%d]: %w", i, err)
		}
	}
	if r.Activity != nil {
		if err := r.Activity.Validate(); err != nil {
			return fmt.Errorf("Activity: %w", err)
		}
	}
	for i, item := range r.BasedOn {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("BasedOn[%d]: %w", i, err)
		}
	}
	if r.Patient != nil {
		if err := r.Patient.Validate(); err != nil {
			return fmt.Errorf("Patient: %w", err)
		}
	}
	if r.Encounter != nil {
		if err := r.Encounter.Validate(); err != nil {
			return fmt.Errorf("Encounter: %w", err)
		}
	}
	if len(r.Agent) < 1 {
		return fmt.Errorf("field 'Agent' must have at least 1 elements")
	}
	for i, item := range r.Agent {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Agent[%d]: %w", i, err)
		}
	}
	for i, item := range r.Entity {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Entity[%d]: %w", i, err)
		}
	}
	for i, item := range r.Signature {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Signature[%d]: %w", i, err)
		}
	}
	return nil
}

type ProvenanceEntity struct {
	Id    *string           `json:"id,omitempty" bson:"id,omitempty"`       // Unique id for inter-element referencing
	Role  string            `json:"role" bson:"role"`                       // revision | quotation | source | instantiates | removal
	What  *Reference        `json:"what" bson:"what"`                       // Identity of entity
	Agent []ProvenanceAgent `json:"agent,omitempty" bson:"agent,omitempty"` // Entity is attributed to this agent
}

func (r *ProvenanceEntity) Validate() error {
	var emptyString string
	if r.Role == emptyString {
		return fmt.Errorf("field 'Role' is required")
	}
	if r.What == nil {
		return fmt.Errorf("field 'What' is required")
	}
	if r.What != nil {
		if err := r.What.Validate(); err != nil {
			return fmt.Errorf("What: %w", err)
		}
	}
	for i, item := range r.Agent {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Agent[%d]: %w", i, err)
		}
	}
	return nil
}

type ProvenanceAgent struct {
	Id         *string           `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Type       *CodeableConcept  `json:"type,omitempty" bson:"type,omitempty"`               // How the agent participated
	Role       []CodeableConcept `json:"role,omitempty" bson:"role,omitempty"`               // What the agents role was
	Who        *Reference        `json:"who" bson:"who"`                                     // The agent that participated in the event
	OnBehalfOf *Reference        `json:"onBehalfOf,omitempty" bson:"on_behalf_of,omitempty"` // The agent that delegated
}

func (r *ProvenanceAgent) Validate() error {
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	for i, item := range r.Role {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Role[%d]: %w", i, err)
		}
	}
	if r.Who == nil {
		return fmt.Errorf("field 'Who' is required")
	}
	if r.Who != nil {
		if err := r.Who.Validate(); err != nil {
			return fmt.Errorf("Who: %w", err)
		}
	}
	if r.OnBehalfOf != nil {
		if err := r.OnBehalfOf.Validate(); err != nil {
			return fmt.Errorf("OnBehalfOf: %w", err)
		}
	}
	return nil
}
