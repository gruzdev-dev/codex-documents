package models

import (
	"fmt"
)

// RelatedArtifact Type: Related artifacts such as dependencies, components, additional documentation, justification, or bibliographic references.
type RelatedArtifact struct {
	Id                 *string     `json:"id,omitempty" bson:"id,omitempty"`                                  // Unique id for inter-element referencing
	Type               string      `json:"type" bson:"type"`                                                  // documentation | justification | citation | predecessor | successor | derived-from | depends-on | composed-of | part-of
	Label              *string     `json:"label,omitempty" bson:"label,omitempty"`                            // Short label
	Display            *string     `json:"display,omitempty" bson:"display,omitempty"`                        // Brief description of the related artifact
	Citation           *string     `json:"citation,omitempty" bson:"citation,omitempty"`                      // Bibliographic citation for the artifact
	Document           *Attachment `json:"document,omitempty" bson:"document,omitempty"`                      // What document is being referenced
	Resource           *string     `json:"resource,omitempty" bson:"resource,omitempty"`                      // What artifact is being referenced
	ResourceReference  *Reference  `json:"resourceReference,omitempty" bson:"resource_reference,omitempty"`   // What artifact, if not a conformance resource
	ArtifactMarkdown   *string     `json:"artifactMarkdown,omitempty" bson:"artifact_markdown,omitempty"`     // What document, citation, artifact, or resource is being referenced
	ArtifactAttachment *Attachment `json:"artifactAttachment,omitempty" bson:"artifact_attachment,omitempty"` // What document, citation, artifact, or resource is being referenced
	ArtifactCanonical  *string     `json:"artifactCanonical,omitempty" bson:"artifact_canonical,omitempty"`   // What document, citation, artifact, or resource is being referenced
	ArtifactReference  *Reference  `json:"artifactReference,omitempty" bson:"artifact_reference,omitempty"`   // What document, citation, artifact, or resource is being referenced
}

func (r *RelatedArtifact) Validate() error {
	var emptyString string
	if r.Type == emptyString {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Document != nil {
		if err := r.Document.Validate(); err != nil {
			return fmt.Errorf("Document: %w", err)
		}
	}
	if r.ResourceReference != nil {
		if err := r.ResourceReference.Validate(); err != nil {
			return fmt.Errorf("ResourceReference: %w", err)
		}
	}
	if r.ArtifactAttachment != nil {
		if err := r.ArtifactAttachment.Validate(); err != nil {
			return fmt.Errorf("ArtifactAttachment: %w", err)
		}
	}
	if r.ArtifactReference != nil {
		if err := r.ArtifactReference.Validate(); err != nil {
			return fmt.Errorf("ArtifactReference: %w", err)
		}
	}
	return nil
}
