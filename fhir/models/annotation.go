package models

// Annotation Type: A  text note which also  contains information about who made the statement and when.
type Annotation struct {
	Id              *string    `json:"id,omitempty" bson:"id,omitempty"`                            // Unique id for inter-element referencing
	AuthorReference *Reference `json:"authorReference,omitempty" bson:"author_reference,omitempty"` // Individual responsible for the annotation
	AuthorString    *string    `json:"authorString,omitempty" bson:"author_string,omitempty"`       // Individual responsible for the annotation
	Time            *string    `json:"time,omitempty" bson:"time,omitempty"`                        // When the annotation was made
	Text            string     `json:"text" bson:"text"`                                            // The annotation  - text content (as markdown)
}
