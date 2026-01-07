package models

import (
	"encoding/json"
	"fmt"
)

// A collection of error, warning, or information messages that result from a system action.
type OperationOutcome struct {
	ResourceType  string                  `json:"resourceType" bson:"resource_type"`                       // Type of resource
	Id            *string                 `json:"id,omitempty" bson:"id,omitempty"`                        // Logical id of this artifact
	Meta          *Meta                   `json:"meta,omitempty" bson:"meta,omitempty"`                    // Metadata about the resource
	ImplicitRules *string                 `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"` // A set of rules under which this content was created
	Language      *string                 `json:"language,omitempty" bson:"language,omitempty"`            // Language of the resource content
	Text          *Narrative              `json:"text,omitempty" bson:"text,omitempty"`                    // Text summary of the resource, for human interpretation
	Contained     []json.RawMessage       `json:"contained,omitempty" bson:"contained,omitempty"`          // Contained, inline Resources
	Issue         []OperationOutcomeIssue `json:"issue" bson:"issue"`                                      // A single issue associated with the action
}

func (r *OperationOutcome) Validate() error {
	if r.ResourceType != "OperationOutcome" {
		return fmt.Errorf("invalid resourceType: expected 'OperationOutcome', got '%s'", r.ResourceType)
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
	if len(r.Issue) < 1 {
		return fmt.Errorf("field 'Issue' must have at least 1 elements")
	}
	for i, item := range r.Issue {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Issue[%d]: %w", i, err)
		}
	}
	return nil
}

type OperationOutcomeIssue struct {
	Id          *string          `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Severity    string           `json:"severity" bson:"severity"`                           // fatal | error | warning | information | success
	Code        string           `json:"code" bson:"code"`                                   // Error or warning code
	Details     *CodeableConcept `json:"details,omitempty" bson:"details,omitempty"`         // Additional details about the error
	Diagnostics *string          `json:"diagnostics,omitempty" bson:"diagnostics,omitempty"` // Additional diagnostic information about the issue
	Location    []string         `json:"location,omitempty" bson:"location,omitempty"`       // Deprecated: Path of element(s) related to issue
	Expression  []string         `json:"expression,omitempty" bson:"expression,omitempty"`   // FHIRPath of element(s) related to issue
}

func (r *OperationOutcomeIssue) Validate() error {
	var emptyString string
	if r.Severity == emptyString {
		return fmt.Errorf("field 'Severity' is required")
	}
	if r.Code == emptyString {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Details != nil {
		if err := r.Details.Validate(); err != nil {
			return fmt.Errorf("Details: %w", err)
		}
	}
	return nil
}
