package models

import (
	"encoding/json"
	"fmt"
)

// This Resource provides one or more comments, classifiers or ratings about a Resource and supports attribution and rights management metadata for the added content.
type ArtifactAssessment struct {
	Id                *string                       `json:"id,omitempty" bson:"id,omitempty"`                           // Logical id of this artifact
	Meta              *Meta                         `json:"meta,omitempty" bson:"meta,omitempty"`                       // Metadata about the resource
	ImplicitRules     *string                       `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`    // A set of rules under which this content was created
	Language          *string                       `json:"language,omitempty" bson:"language,omitempty"`               // Language of the resource content
	Text              *Narrative                    `json:"text,omitempty" bson:"text,omitempty"`                       // Text summary of the resource, for human interpretation
	Contained         []json.RawMessage             `json:"contained,omitempty" bson:"contained,omitempty"`             // Contained, inline Resources
	Identifier        []Identifier                  `json:"identifier,omitempty" bson:"identifier,omitempty"`           // Additional identifier for the artifact assessment
	Title             *string                       `json:"title,omitempty" bson:"title,omitempty"`                     // A label for use in displaying and selecting the artifact assessment
	CiteAs            *string                       `json:"citeAs,omitempty" bson:"cite_as,omitempty"`                  // How to cite the comment or rating
	ArtifactReference *Reference                    `json:"artifactReference" bson:"artifact_reference"`                // The artifact assessed, commented upon or rated
	ArtifactCanonical *string                       `json:"artifactCanonical" bson:"artifact_canonical"`                // The artifact assessed, commented upon or rated
	ArtifactUri       *string                       `json:"artifactUri" bson:"artifact_uri"`                            // The artifact assessed, commented upon or rated
	RelatesTo         []ArtifactAssessmentRelatesTo `json:"relatesTo,omitempty" bson:"relates_to,omitempty"`            // Relationship to other Resources
	Date              *string                       `json:"date,omitempty" bson:"date,omitempty"`                       // Date last changed
	Copyright         *string                       `json:"copyright,omitempty" bson:"copyright,omitempty"`             // Notice about intellectual property ownership, can include restrictions on use
	ApprovalDate      *string                       `json:"approvalDate,omitempty" bson:"approval_date,omitempty"`      // When the artifact assessment was approved by publisher
	LastReviewDate    *string                       `json:"lastReviewDate,omitempty" bson:"last_review_date,omitempty"` // When the artifact assessment was last reviewed by the publisher
	Content           []ArtifactAssessmentContent   `json:"content,omitempty" bson:"content,omitempty"`                 // Comment, classifier, or rating content
	WorkflowStatus    *string                       `json:"workflowStatus,omitempty" bson:"workflow_status,omitempty"`  // submitted | triaged | waiting-for-input | resolved-no-change | resolved-change-required | deferred | duplicate | applied | published | entered-in-error
	Disposition       *string                       `json:"disposition,omitempty" bson:"disposition,omitempty"`         // unresolved | not-persuasive | persuasive | persuasive-with-modification | not-persuasive-with-modification
}

func (r *ArtifactAssessment) Validate() error {
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
	if r.ArtifactReference == nil {
		return fmt.Errorf("field 'ArtifactReference' is required")
	}
	if r.ArtifactReference != nil {
		if err := r.ArtifactReference.Validate(); err != nil {
			return fmt.Errorf("ArtifactReference: %w", err)
		}
	}
	if r.ArtifactCanonical == nil {
		return fmt.Errorf("field 'ArtifactCanonical' is required")
	}
	if r.ArtifactUri == nil {
		return fmt.Errorf("field 'ArtifactUri' is required")
	}
	for i, item := range r.RelatesTo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RelatesTo[%d]: %w", i, err)
		}
	}
	for i, item := range r.Content {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Content[%d]: %w", i, err)
		}
	}
	return nil
}

type ArtifactAssessmentContent struct {
	Id          *string                       `json:"id,omitempty" bson:"id,omitempty"`                     // Unique id for inter-element referencing
	Summary     *string                       `json:"summary,omitempty" bson:"summary,omitempty"`           // Brief summary of the content
	Type        *CodeableConcept              `json:"type,omitempty" bson:"type,omitempty"`                 // What type of content
	Classifier  []CodeableConcept             `json:"classifier,omitempty" bson:"classifier,omitempty"`     // Rating, classifier, or assessment
	Quantity    *Quantity                     `json:"quantity,omitempty" bson:"quantity,omitempty"`         // Quantitative rating
	Author      []Reference                   `json:"author,omitempty" bson:"author,omitempty"`             // Who authored the content
	Path        []string                      `json:"path,omitempty" bson:"path,omitempty"`                 // What the comment is directed to
	RelatesTo   []ArtifactAssessmentRelatesTo `json:"relatesTo,omitempty" bson:"relates_to,omitempty"`      // Relationship to other Resources
	FreeToShare bool                          `json:"freeToShare,omitempty" bson:"free_to_share,omitempty"` // Acceptable to publicly share the content
	Component   []ArtifactAssessmentContent   `json:"component,omitempty" bson:"component,omitempty"`       // Comment, classifier, or rating content
}

func (r *ArtifactAssessmentContent) Validate() error {
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	for i, item := range r.Classifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Classifier[%d]: %w", i, err)
		}
	}
	if r.Quantity != nil {
		if err := r.Quantity.Validate(); err != nil {
			return fmt.Errorf("Quantity: %w", err)
		}
	}
	for i, item := range r.Author {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Author[%d]: %w", i, err)
		}
	}
	for i, item := range r.RelatesTo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RelatesTo[%d]: %w", i, err)
		}
	}
	for i, item := range r.Component {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Component[%d]: %w", i, err)
		}
	}
	return nil
}

type ArtifactAssessmentRelatesTo struct {
	Id               *string          `json:"id,omitempty" bson:"id,omitempty"`          // Unique id for inter-element referencing
	Type             *CodeableConcept `json:"type" bson:"type"`                          // documentation | justification | citation | predecessor | successor | derived-from | depends-on | composed-of | part-of | amends | amended-with | appends | appended-with | cites | cited-by | comments-on | comment-in | contains | contained-in | corrects | correction-in | replaces | replaced-with | retracts | retracted-by | signs | similar-to | supports | supported-with | transforms | transformed-into | transformed-with | documents | specification-of | created-with | cite-as | reprint | reprint-of | summarizes
	TargetUri        *string          `json:"targetUri" bson:"target_uri"`               // The artifact that is related to this ArtifactAssessment
	TargetAttachment *Attachment      `json:"targetAttachment" bson:"target_attachment"` // The artifact that is related to this ArtifactAssessment
	TargetCanonical  *string          `json:"targetCanonical" bson:"target_canonical"`   // The artifact that is related to this ArtifactAssessment
	TargetReference  *Reference       `json:"targetReference" bson:"target_reference"`   // The artifact that is related to this ArtifactAssessment
	TargetMarkdown   *string          `json:"targetMarkdown" bson:"target_markdown"`     // The artifact that is related to this ArtifactAssessment
}

func (r *ArtifactAssessmentRelatesTo) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.TargetUri == nil {
		return fmt.Errorf("field 'TargetUri' is required")
	}
	if r.TargetAttachment == nil {
		return fmt.Errorf("field 'TargetAttachment' is required")
	}
	if r.TargetAttachment != nil {
		if err := r.TargetAttachment.Validate(); err != nil {
			return fmt.Errorf("TargetAttachment: %w", err)
		}
	}
	if r.TargetCanonical == nil {
		return fmt.Errorf("field 'TargetCanonical' is required")
	}
	if r.TargetReference == nil {
		return fmt.Errorf("field 'TargetReference' is required")
	}
	if r.TargetReference != nil {
		if err := r.TargetReference.Validate(); err != nil {
			return fmt.Errorf("TargetReference: %w", err)
		}
	}
	if r.TargetMarkdown == nil {
		return fmt.Errorf("field 'TargetMarkdown' is required")
	}
	return nil
}
