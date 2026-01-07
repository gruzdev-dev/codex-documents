package models

import (
	"fmt"
)

// A resource that represents the data of a single raw artifact as digital content accessible in its native format.  A Binary resource can contain any content, whether text, image, pdf, zip archive, etc.
type Binary struct {
	ResourceType    string     `json:"resourceType" bson:"resource_type"`                           // Type of resource
	Id              *string    `json:"id,omitempty" bson:"id,omitempty"`                            // Logical id of this artifact
	Meta            *Meta      `json:"meta,omitempty" bson:"meta,omitempty"`                        // Metadata about the resource
	ImplicitRules   *string    `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`     // A set of rules under which this content was created
	Language        *string    `json:"language,omitempty" bson:"language,omitempty"`                // Language of the resource content
	ContentType     string     `json:"contentType" bson:"content_type"`                             // MimeType of the binary content
	SecurityContext *Reference `json:"securityContext,omitempty" bson:"security_context,omitempty"` // Identifies another resource to use as proxy when enforcing access control
	Data            *string    `json:"data,omitempty" bson:"data,omitempty"`                        // The actual content
}

func (r *Binary) Validate() error {
	if r.ResourceType != "Binary" {
		return fmt.Errorf("invalid resourceType: expected 'Binary', got '%s'", r.ResourceType)
	}
	if r.Meta != nil {
		if err := r.Meta.Validate(); err != nil {
			return fmt.Errorf("Meta: %w", err)
		}
	}
	var emptyString string
	if r.ContentType == emptyString {
		return fmt.Errorf("field 'ContentType' is required")
	}
	if r.SecurityContext != nil {
		if err := r.SecurityContext.Validate(); err != nil {
			return fmt.Errorf("SecurityContext: %w", err)
		}
	}
	return nil
}
