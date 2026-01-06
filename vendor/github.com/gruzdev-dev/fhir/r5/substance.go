package models

import (
	"encoding/json"
	"fmt"
)

// A homogeneous material with a definite composition.
type Substance struct {
	Id            *string            `json:"id,omitempty" bson:"id,omitempty"`                        // Logical id of this artifact
	Meta          *Meta              `json:"meta,omitempty" bson:"meta,omitempty"`                    // Metadata about the resource
	ImplicitRules *string            `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"` // A set of rules under which this content was created
	Language      *string            `json:"language,omitempty" bson:"language,omitempty"`            // Language of the resource content
	Text          *Narrative         `json:"text,omitempty" bson:"text,omitempty"`                    // Text summary of the resource, for human interpretation
	Contained     []json.RawMessage  `json:"contained,omitempty" bson:"contained,omitempty"`          // Contained, inline Resources
	Identifier    []Identifier       `json:"identifier,omitempty" bson:"identifier,omitempty"`        // Unique identifier
	Status        *string            `json:"status,omitempty" bson:"status,omitempty"`                // active | inactive | entered-in-error
	Category      []CodeableConcept  `json:"category,omitempty" bson:"category,omitempty"`            // What class/type of substance this is
	Code          *CodeableReference `json:"code" bson:"code"`                                        // What substance this is
	Description   *string            `json:"description,omitempty" bson:"description,omitempty"`      // Textual description of the substance, comments
	Expiry        *string            `json:"expiry,omitempty" bson:"expiry,omitempty"`                // When no longer valid to use
	Quantity      *Quantity          `json:"quantity,omitempty" bson:"quantity,omitempty"`            // Amount of substance in the package
}

func (r *Substance) Validate() error {
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
	for i, item := range r.Category {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Category[%d]: %w", i, err)
		}
	}
	if r.Code == nil {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.Quantity != nil {
		if err := r.Quantity.Validate(); err != nil {
			return fmt.Errorf("Quantity: %w", err)
		}
	}
	return nil
}
