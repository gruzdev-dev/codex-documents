package models

import (
	"encoding/json"
	"fmt"
)

// This resource provides the insurance enrollment details to the insurer regarding a specified coverage.
type EnrollmentRequest struct {
	Id            *string           `json:"id,omitempty" bson:"id,omitempty"`                        // Logical id of this artifact
	Meta          *Meta             `json:"meta,omitempty" bson:"meta,omitempty"`                    // Metadata about the resource
	ImplicitRules *string           `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"` // A set of rules under which this content was created
	Language      *string           `json:"language,omitempty" bson:"language,omitempty"`            // Language of the resource content
	Text          *Narrative        `json:"text,omitempty" bson:"text,omitempty"`                    // Text summary of the resource, for human interpretation
	Contained     []json.RawMessage `json:"contained,omitempty" bson:"contained,omitempty"`          // Contained, inline Resources
	Identifier    []Identifier      `json:"identifier,omitempty" bson:"identifier,omitempty"`        // Business Identifier
	Status        *string           `json:"status,omitempty" bson:"status,omitempty"`                // active | cancelled | draft | entered-in-error
	Created       *string           `json:"created,omitempty" bson:"created,omitempty"`              // Creation date
	Insurer       *Reference        `json:"insurer,omitempty" bson:"insurer,omitempty"`              // Target
	Provider      *Reference        `json:"provider,omitempty" bson:"provider,omitempty"`            // Responsible practitioner
	Candidate     *Reference        `json:"candidate,omitempty" bson:"candidate,omitempty"`          // The subject(s)to be enrolled
	Coverage      *Reference        `json:"coverage,omitempty" bson:"coverage,omitempty"`            // Insurance information
}

func (r *EnrollmentRequest) Validate() error {
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
	if r.Insurer != nil {
		if err := r.Insurer.Validate(); err != nil {
			return fmt.Errorf("Insurer: %w", err)
		}
	}
	if r.Provider != nil {
		if err := r.Provider.Validate(); err != nil {
			return fmt.Errorf("Provider: %w", err)
		}
	}
	if r.Candidate != nil {
		if err := r.Candidate.Validate(); err != nil {
			return fmt.Errorf("Candidate: %w", err)
		}
	}
	if r.Coverage != nil {
		if err := r.Coverage.Validate(); err != nil {
			return fmt.Errorf("Coverage: %w", err)
		}
	}
	return nil
}
