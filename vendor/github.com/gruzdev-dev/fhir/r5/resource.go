package models

import (
	"fmt"
)

// This is the base resource type for everything.
type Resource struct {
	ResourceType  string  `json:"resourceType" bson:"resource_type"`                       // Type of resource
	Id            *string `json:"id,omitempty" bson:"id,omitempty"`                        // Logical id of this artifact
	Meta          *Meta   `json:"meta,omitempty" bson:"meta,omitempty"`                    // Metadata about the resource
	ImplicitRules *string `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"` // A set of rules under which this content was created
	Language      *string `json:"language,omitempty" bson:"language,omitempty"`            // Language of the resource content
}

func (r *Resource) Validate() error {
	if r.ResourceType != "Resource" {
		return fmt.Errorf("invalid resourceType: expected 'Resource', got '%s'", r.ResourceType)
	}
	if r.Meta != nil {
		if err := r.Meta.Validate(); err != nil {
			return fmt.Errorf("Meta: %w", err)
		}
	}
	return nil
}
