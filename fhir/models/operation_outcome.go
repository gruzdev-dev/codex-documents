package fhir

import "encoding/json"

// A collection of error, warning, or information messages that result from a system action.
type OperationOutcome struct {
	Id            *string                 `json:"id,omitempty"`            // Logical id of this artifact
	Meta          *Meta                   `json:"meta,omitempty"`          // Metadata about the resource
	ImplicitRules *string                 `json:"implicitRules,omitempty"` // A set of rules under which this content was created
	Language      *string                 `json:"language,omitempty"`      // Language of the resource content
	Text          *Narrative              `json:"text,omitempty"`          // Text summary of the resource, for human interpretation
	Contained     []json.RawMessage       `json:"contained,omitempty"`     // Contained, inline Resources
	Issue         []OperationOutcomeIssue `json:"issue,omitempty"`         // A single issue associated with the action
}

type OperationOutcomeIssue struct {
	Id          *string          `json:"id,omitempty"`          // Unique id for inter-element referencing
	Severity    string           `json:"severity,omitempty"`    // fatal | error | warning | information | success
	Code        string           `json:"code,omitempty"`        // Error or warning code
	Details     *CodeableConcept `json:"details,omitempty"`     // Additional details about the error
	Diagnostics *string          `json:"diagnostics,omitempty"` // Additional diagnostic information about the issue
	Location    []string         `json:"location,omitempty"`    // Deprecated: Path of element(s) related to issue
	Expression  []string         `json:"expression,omitempty"`  // FHIRPath of element(s) related to issue
}
