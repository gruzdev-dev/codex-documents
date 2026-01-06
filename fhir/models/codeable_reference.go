package models

// CodeableReference Type: A reference to a resource (by instance), or instead, a reference to a concept defined in a terminology or ontology (by class).
type CodeableReference struct {
	Id        *string          `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Concept   *CodeableConcept `json:"concept,omitempty" bson:"concept,omitempty"`     // Reference to a concept (by class)
	Reference *Reference       `json:"reference,omitempty" bson:"reference,omitempty"` // Reference to a resource (by instance)
}
