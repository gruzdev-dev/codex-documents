package fhir

// Annotation Type: A  text note which also  contains information about who made the statement and when.
type Annotation struct {
	Id     *string `json:"id,omitempty"`        // Unique id for inter-element referencing
	Author any     `json:"author[x],omitempty"` // Individual responsible for the annotation
	Time   *string `json:"time,omitempty"`      // When the annotation was made
	Text   string  `json:"text,omitempty"`      // The annotation  - text content (as markdown)
}
