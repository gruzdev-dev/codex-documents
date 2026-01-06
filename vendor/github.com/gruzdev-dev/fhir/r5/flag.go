package models

import (
	"encoding/json"
	"fmt"
)

// Prospective warnings of potential issues when providing care to the patient.
type Flag struct {
	Id             *string           `json:"id,omitempty" bson:"id,omitempty"`                          // Logical id of this artifact
	Meta           *Meta             `json:"meta,omitempty" bson:"meta,omitempty"`                      // Metadata about the resource
	ImplicitRules  *string           `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`   // A set of rules under which this content was created
	Language       *string           `json:"language,omitempty" bson:"language,omitempty"`              // Language of the resource content
	Text           *Narrative        `json:"text,omitempty" bson:"text,omitempty"`                      // Text summary of the resource, for human interpretation
	Contained      []json.RawMessage `json:"contained,omitempty" bson:"contained,omitempty"`            // Contained, inline Resources
	Identifier     []Identifier      `json:"identifier,omitempty" bson:"identifier,omitempty"`          // Business identifier
	Status         *string           `json:"status,omitempty" bson:"status,omitempty"`                  // active | inactive | entered-in-error
	Category       []CodeableConcept `json:"category,omitempty" bson:"category,omitempty"`              // Clinical, administrative, etc
	Code           *CodeableConcept  `json:"code" bson:"code"`                                          // Coded or textual message to display to user
	Subject        *Reference        `json:"subject" bson:"subject"`                                    // Who/What is flag about?
	Period         *Period           `json:"period,omitempty" bson:"period,omitempty"`                  // Time period when flag is active
	Encounter      *Reference        `json:"encounter,omitempty" bson:"encounter,omitempty"`            // Alert relevant during encounter
	Author         *Reference        `json:"author,omitempty" bson:"author,omitempty"`                  // Flag creator
	SupportingInfo []Reference       `json:"supportingInfo,omitempty" bson:"supporting_info,omitempty"` // Extra information to use in context of the flag
}

func (r *Flag) Validate() error {
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
	if r.Subject == nil {
		return fmt.Errorf("field 'Subject' is required")
	}
	if r.Subject != nil {
		if err := r.Subject.Validate(); err != nil {
			return fmt.Errorf("Subject: %w", err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	if r.Encounter != nil {
		if err := r.Encounter.Validate(); err != nil {
			return fmt.Errorf("Encounter: %w", err)
		}
	}
	if r.Author != nil {
		if err := r.Author.Validate(); err != nil {
			return fmt.Errorf("Author: %w", err)
		}
	}
	for i, item := range r.SupportingInfo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SupportingInfo[%d]: %w", i, err)
		}
	}
	return nil
}
