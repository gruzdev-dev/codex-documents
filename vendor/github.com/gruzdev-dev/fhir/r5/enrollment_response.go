package models

import (
	"encoding/json"
	"fmt"
)

// This resource provides enrollment and plan details from the processing of an EnrollmentRequest resource.
type EnrollmentResponse struct {
	Id              *string           `json:"id,omitempty" bson:"id,omitempty"`                            // Logical id of this artifact
	Meta            *Meta             `json:"meta,omitempty" bson:"meta,omitempty"`                        // Metadata about the resource
	ImplicitRules   *string           `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`     // A set of rules under which this content was created
	Language        *string           `json:"language,omitempty" bson:"language,omitempty"`                // Language of the resource content
	Text            *Narrative        `json:"text,omitempty" bson:"text,omitempty"`                        // Text summary of the resource, for human interpretation
	Contained       []json.RawMessage `json:"contained,omitempty" bson:"contained,omitempty"`              // Contained, inline Resources
	Identifier      []Identifier      `json:"identifier,omitempty" bson:"identifier,omitempty"`            // Business Identifier
	Status          *string           `json:"status,omitempty" bson:"status,omitempty"`                    // active | cancelled | draft | entered-in-error
	Request         *Reference        `json:"request,omitempty" bson:"request,omitempty"`                  // Claim reference
	Outcome         *string           `json:"outcome,omitempty" bson:"outcome,omitempty"`                  // queued | complete | error | partial
	Disposition     *string           `json:"disposition,omitempty" bson:"disposition,omitempty"`          // Disposition Message
	Created         *string           `json:"created,omitempty" bson:"created,omitempty"`                  // Creation date
	Organization    *Reference        `json:"organization,omitempty" bson:"organization,omitempty"`        // Insurer
	RequestProvider *Reference        `json:"requestProvider,omitempty" bson:"request_provider,omitempty"` // Responsible practitioner
	Candidate       *Reference        `json:"candidate,omitempty" bson:"candidate,omitempty"`              // The subject(s)to be enrolled
}

func (r *EnrollmentResponse) Validate() error {
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
	if r.Request != nil {
		if err := r.Request.Validate(); err != nil {
			return fmt.Errorf("Request: %w", err)
		}
	}
	if r.Organization != nil {
		if err := r.Organization.Validate(); err != nil {
			return fmt.Errorf("Organization: %w", err)
		}
	}
	if r.RequestProvider != nil {
		if err := r.RequestProvider.Validate(); err != nil {
			return fmt.Errorf("RequestProvider: %w", err)
		}
	}
	if r.Candidate != nil {
		if err := r.Candidate.Validate(); err != nil {
			return fmt.Errorf("Candidate: %w", err)
		}
	}
	return nil
}
