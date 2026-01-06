package models

import (
	"encoding/json"
	"fmt"
)

// A resource that includes narrative, extensions, and contained resources.
type DomainResource struct {
	Id            *string           `json:"id,omitempty" bson:"id,omitempty"`                        // Logical id of this artifact
	Meta          *Meta             `json:"meta,omitempty" bson:"meta,omitempty"`                    // Metadata about the resource
	ImplicitRules *string           `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"` // A set of rules under which this content was created
	Language      *string           `json:"language,omitempty" bson:"language,omitempty"`            // Language of the resource content
	Text          *Narrative        `json:"text,omitempty" bson:"text,omitempty"`                    // Text summary of the resource, for human interpretation
	Contained     []json.RawMessage `json:"contained,omitempty" bson:"contained,omitempty"`          // Contained, inline Resources
}

func (r *DomainResource) Validate() error {
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
	return nil
}
