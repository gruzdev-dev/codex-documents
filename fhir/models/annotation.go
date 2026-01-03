package models

// Annotation Type: A  text note which also  contains information about who made the statement and when.
type Annotation struct {
	Id              *string    `json:"id,omitempty"`              // Unique id for inter-element referencing
	AuthorReference *Reference `json:"authorReference,omitempty"` // Individual responsible for the annotation
	AuthorString    *string    `json:"authorString,omitempty"`    // Individual responsible for the annotation
	Time            *string    `json:"time,omitempty"`            // When the annotation was made
	Text            string     `json:"text,omitempty"`            // The annotation  - text content (as markdown)
}
