package models

import (
	"encoding/json"
	"fmt"
)

// A record of an event relevant for purposes such as operations, privacy, security, maintenance, and performance analysis.
type AuditEvent struct {
	ResourceType     string             `json:"resourceType" bson:"resource_type"`                              // Type of resource
	Id               *string            `json:"id,omitempty" bson:"id,omitempty"`                               // Logical id of this artifact
	Meta             *Meta              `json:"meta,omitempty" bson:"meta,omitempty"`                           // Metadata about the resource
	ImplicitRules    *string            `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`        // A set of rules under which this content was created
	Language         *string            `json:"language,omitempty" bson:"language,omitempty"`                   // Language of the resource content
	Text             *Narrative         `json:"text,omitempty" bson:"text,omitempty"`                           // Text summary of the resource, for human interpretation
	Contained        []json.RawMessage  `json:"contained,omitempty" bson:"contained,omitempty"`                 // Contained, inline Resources
	Type             *CodeableConcept   `json:"type" bson:"type"`                                               // High level categorization of audit event
	Subtype          []CodeableConcept  `json:"subtype,omitempty" bson:"subtype,omitempty"`                     // Specific type of event
	Action           *string            `json:"action,omitempty" bson:"action,omitempty"`                       // Type of action performed during the event
	Severity         *string            `json:"severity,omitempty" bson:"severity,omitempty"`                   // emergency | alert | critical | error | warning | notice | informational | debug
	OccurredPeriod   *Period            `json:"occurredPeriod,omitempty" bson:"occurred_period,omitempty"`      // When the activity occurred
	OccurredDateTime *string            `json:"occurredDateTime,omitempty" bson:"occurred_date_time,omitempty"` // When the activity occurred
	Recorded         string             `json:"recorded" bson:"recorded"`                                       // Time when the event was recorded
	Outcome          *AuditEventOutcome `json:"outcome,omitempty" bson:"outcome,omitempty"`                     // Whether the event succeeded or failed
	Authorization    []CodeableConcept  `json:"authorization,omitempty" bson:"authorization,omitempty"`         // Authorization related to the event
	BasedOn          []Reference        `json:"basedOn,omitempty" bson:"based_on,omitempty"`                    // Workflow authorization within which this event occurred
	Patient          *Reference         `json:"patient,omitempty" bson:"patient,omitempty"`                     // The patient is the subject of the data used/created/updated/deleted during the activity
	Encounter        *Reference         `json:"encounter,omitempty" bson:"encounter,omitempty"`                 // Encounter within which this event occurred or which the event is tightly associated
	Agent            []AuditEventAgent  `json:"agent" bson:"agent"`                                             // Actor involved in the event
	Source           *AuditEventSource  `json:"source" bson:"source"`                                           // Audit Event Reporter
	Entity           []AuditEventEntity `json:"entity,omitempty" bson:"entity,omitempty"`                       // Data or objects used
}

func (r *AuditEvent) Validate() error {
	if r.ResourceType != "AuditEvent" {
		return fmt.Errorf("invalid resourceType: expected 'AuditEvent', got '%s'", r.ResourceType)
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
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	for i, item := range r.Subtype {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Subtype[%d]: %w", i, err)
		}
	}
	if r.OccurredPeriod != nil {
		if err := r.OccurredPeriod.Validate(); err != nil {
			return fmt.Errorf("OccurredPeriod: %w", err)
		}
	}
	var emptyString string
	if r.Recorded == emptyString {
		return fmt.Errorf("field 'Recorded' is required")
	}
	if r.Outcome != nil {
		if err := r.Outcome.Validate(); err != nil {
			return fmt.Errorf("Outcome: %w", err)
		}
	}
	for i, item := range r.Authorization {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Authorization[%d]: %w", i, err)
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
	if r.Source == nil {
		return fmt.Errorf("field 'Source' is required")
	}
	if r.Source != nil {
		if err := r.Source.Validate(); err != nil {
			return fmt.Errorf("Source: %w", err)
		}
	}
	for i, item := range r.Entity {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Entity[%d]: %w", i, err)
		}
	}
	return nil
}

type AuditEventOutcome struct {
	Id     *string           `json:"id,omitempty" bson:"id,omitempty"`         // Unique id for inter-element referencing
	Code   *Coding           `json:"code" bson:"code"`                         // Whether the event succeeded or failed
	Detail []CodeableConcept `json:"detail,omitempty" bson:"detail,omitempty"` // Additional outcome detail
}

func (r *AuditEventOutcome) Validate() error {
	if r.Code == nil {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	for i, item := range r.Detail {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Detail[%d]: %w", i, err)
		}
	}
	return nil
}

type AuditEventAgent struct {
	Id               *string           `json:"id,omitempty" bson:"id,omitempty"`                              // Unique id for inter-element referencing
	Type             *CodeableConcept  `json:"type,omitempty" bson:"type,omitempty"`                          // How agent participated
	Role             []CodeableConcept `json:"role,omitempty" bson:"role,omitempty"`                          // Agent role in the event
	Who              *Reference        `json:"who" bson:"who"`                                                // Identifier of who
	Requestor        bool              `json:"requestor,omitempty" bson:"requestor,omitempty"`                // Whether user is initiator
	Location         *Reference        `json:"location,omitempty" bson:"location,omitempty"`                  // The agent location when the event occurred
	Policy           []string          `json:"policy,omitempty" bson:"policy,omitempty"`                      // Policy that authorized the agent participation in the event
	NetworkReference *Reference        `json:"networkReference,omitempty" bson:"network_reference,omitempty"` // This agent network location for the activity
	NetworkUri       *string           `json:"networkUri,omitempty" bson:"network_uri,omitempty"`             // This agent network location for the activity
	NetworkString    *string           `json:"networkString,omitempty" bson:"network_string,omitempty"`       // This agent network location for the activity
	Authorization    []CodeableConcept `json:"authorization,omitempty" bson:"authorization,omitempty"`        // Allowable authorization for this agent
}

func (r *AuditEventAgent) Validate() error {
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
	if r.Location != nil {
		if err := r.Location.Validate(); err != nil {
			return fmt.Errorf("Location: %w", err)
		}
	}
	if r.NetworkReference != nil {
		if err := r.NetworkReference.Validate(); err != nil {
			return fmt.Errorf("NetworkReference: %w", err)
		}
	}
	for i, item := range r.Authorization {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Authorization[%d]: %w", i, err)
		}
	}
	return nil
}

type AuditEventSource struct {
	Id       *string           `json:"id,omitempty" bson:"id,omitempty"`     // Unique id for inter-element referencing
	Site     *Reference        `json:"site,omitempty" bson:"site,omitempty"` // Logical source location within the enterprise
	Observer *Reference        `json:"observer" bson:"observer"`             // The identity of source detecting the event
	Type     []CodeableConcept `json:"type,omitempty" bson:"type,omitempty"` // The type of source where event originated
}

func (r *AuditEventSource) Validate() error {
	if r.Site != nil {
		if err := r.Site.Validate(); err != nil {
			return fmt.Errorf("Site: %w", err)
		}
	}
	if r.Observer == nil {
		return fmt.Errorf("field 'Observer' is required")
	}
	if r.Observer != nil {
		if err := r.Observer.Validate(); err != nil {
			return fmt.Errorf("Observer: %w", err)
		}
	}
	for i, item := range r.Type {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Type[%d]: %w", i, err)
		}
	}
	return nil
}

type AuditEventEntity struct {
	Id            *string                  `json:"id,omitempty" bson:"id,omitempty"`                        // Unique id for inter-element referencing
	What          *Reference               `json:"what,omitempty" bson:"what,omitempty"`                    // Specific instance of resource
	Role          *CodeableConcept         `json:"role,omitempty" bson:"role,omitempty"`                    // What role the entity played
	SecurityLabel []CodeableConcept        `json:"securityLabel,omitempty" bson:"security_label,omitempty"` // Security labels on the entity
	Description   *string                  `json:"description,omitempty" bson:"description,omitempty"`      // Descriptive text
	Query         *string                  `json:"query,omitempty" bson:"query,omitempty"`                  // Query parameters
	Detail        []AuditEventEntityDetail `json:"detail,omitempty" bson:"detail,omitempty"`                // Additional Information about the entity
	Agent         []AuditEventAgent        `json:"agent,omitempty" bson:"agent,omitempty"`                  // Entity is attributed to this agent
}

func (r *AuditEventEntity) Validate() error {
	if r.What != nil {
		if err := r.What.Validate(); err != nil {
			return fmt.Errorf("What: %w", err)
		}
	}
	if r.Role != nil {
		if err := r.Role.Validate(); err != nil {
			return fmt.Errorf("Role: %w", err)
		}
	}
	for i, item := range r.SecurityLabel {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SecurityLabel[%d]: %w", i, err)
		}
	}
	for i, item := range r.Detail {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Detail[%d]: %w", i, err)
		}
	}
	for i, item := range r.Agent {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Agent[%d]: %w", i, err)
		}
	}
	return nil
}

type AuditEventEntityDetail struct {
	Id                   *string          `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Type                 *CodeableConcept `json:"type" bson:"type"`                                   // The name of the extra detail property
	ValueQuantity        *Quantity        `json:"valueQuantity" bson:"value_quantity"`                // Property value
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept" bson:"value_codeable_concept"` // Property value
	ValueString          *string          `json:"valueString" bson:"value_string"`                    // Property value
	ValueBoolean         *bool            `json:"valueBoolean" bson:"value_boolean"`                  // Property value
	ValueInteger         *int             `json:"valueInteger" bson:"value_integer"`                  // Property value
	ValueRange           *Range           `json:"valueRange" bson:"value_range"`                      // Property value
	ValueRatio           *Ratio           `json:"valueRatio" bson:"value_ratio"`                      // Property value
	ValueTime            *string          `json:"valueTime" bson:"value_time"`                        // Property value
	ValueDateTime        *string          `json:"valueDateTime" bson:"value_date_time"`               // Property value
	ValuePeriod          *Period          `json:"valuePeriod" bson:"value_period"`                    // Property value
	ValueBase64Binary    *string          `json:"valueBase64Binary" bson:"value_base64_binary"`       // Property value
}

func (r *AuditEventEntityDetail) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.ValueQuantity == nil {
		return fmt.Errorf("field 'ValueQuantity' is required")
	}
	if r.ValueQuantity != nil {
		if err := r.ValueQuantity.Validate(); err != nil {
			return fmt.Errorf("ValueQuantity: %w", err)
		}
	}
	if r.ValueCodeableConcept == nil {
		return fmt.Errorf("field 'ValueCodeableConcept' is required")
	}
	if r.ValueCodeableConcept != nil {
		if err := r.ValueCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("ValueCodeableConcept: %w", err)
		}
	}
	if r.ValueString == nil {
		return fmt.Errorf("field 'ValueString' is required")
	}
	if r.ValueBoolean == nil {
		return fmt.Errorf("field 'ValueBoolean' is required")
	}
	if r.ValueInteger == nil {
		return fmt.Errorf("field 'ValueInteger' is required")
	}
	if r.ValueRange == nil {
		return fmt.Errorf("field 'ValueRange' is required")
	}
	if r.ValueRange != nil {
		if err := r.ValueRange.Validate(); err != nil {
			return fmt.Errorf("ValueRange: %w", err)
		}
	}
	if r.ValueRatio == nil {
		return fmt.Errorf("field 'ValueRatio' is required")
	}
	if r.ValueRatio != nil {
		if err := r.ValueRatio.Validate(); err != nil {
			return fmt.Errorf("ValueRatio: %w", err)
		}
	}
	if r.ValueTime == nil {
		return fmt.Errorf("field 'ValueTime' is required")
	}
	if r.ValueDateTime == nil {
		return fmt.Errorf("field 'ValueDateTime' is required")
	}
	if r.ValuePeriod == nil {
		return fmt.Errorf("field 'ValuePeriod' is required")
	}
	if r.ValuePeriod != nil {
		if err := r.ValuePeriod.Validate(); err != nil {
			return fmt.Errorf("ValuePeriod: %w", err)
		}
	}
	if r.ValueBase64Binary == nil {
		return fmt.Errorf("field 'ValueBase64Binary' is required")
	}
	return nil
}
