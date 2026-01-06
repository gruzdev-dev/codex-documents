package models

import (
	"fmt"
)

// Meta Type: The metadata about a resource. This is content in the resource that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
type Meta struct {
	Id          *string  `json:"id,omitempty" bson:"id,omitempty"`                    // Unique id for inter-element referencing
	VersionId   *string  `json:"versionId,omitempty" bson:"version_id,omitempty"`     // Version specific identifier
	LastUpdated *string  `json:"lastUpdated,omitempty" bson:"last_updated,omitempty"` // When the resource version last changed
	Source      *string  `json:"source,omitempty" bson:"source,omitempty"`            // Identifies where the resource comes from
	Profile     []string `json:"profile,omitempty" bson:"profile,omitempty"`          // Profiles this resource claims to conform to
	Security    []Coding `json:"security,omitempty" bson:"security,omitempty"`        // Security Labels applied to this resource
	Tag         []Coding `json:"tag,omitempty" bson:"tag,omitempty"`                  // Tags applied to this resource
}

func (r *Meta) Validate() error {
	for i, item := range r.Security {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Security[%d]: %w", i, err)
		}
	}
	for i, item := range r.Tag {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Tag[%d]: %w", i, err)
		}
	}
	return nil
}
