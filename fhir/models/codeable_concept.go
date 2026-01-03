package models

// CodeableConcept Type: A concept that may be defined by a formal reference to a terminology or ontology or may be provided by text.
type CodeableConcept struct {
	Id     *string  `json:"id,omitempty"`     // Unique id for inter-element referencing
	Coding []Coding `json:"coding,omitempty"` // Code defined by a terminology system
	Text   *string  `json:"text,omitempty"`   // Plain text representation of the concept
}
