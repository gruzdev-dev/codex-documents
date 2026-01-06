package models

import (
	"fmt"
)

// Annotation Type: A  text note which also  contains information about who made the statement and when.
type Annotation struct {
	Id              *string    `json:"id,omitempty" bson:"id,omitempty"`                            // Unique id for inter-element referencing
	AuthorReference *Reference `json:"authorReference,omitempty" bson:"author_reference,omitempty"` // Individual responsible for the annotation
	AuthorString    *string    `json:"authorString,omitempty" bson:"author_string,omitempty"`       // Individual responsible for the annotation
	Time            *string    `json:"time,omitempty" bson:"time,omitempty"`                        // When the annotation was made
	Text            string     `json:"text" bson:"text"`                                            // The annotation  - text content (as markdown)
}

func (r *Annotation) Validate() error {
	if r.AuthorReference != nil {
		if err := r.AuthorReference.Validate(); err != nil {
			return fmt.Errorf("AuthorReference: %w", err)
		}
	}
	var emptyString string
	if r.Text == emptyString {
		return fmt.Errorf("field 'Text' is required")
	}
	return nil
}
