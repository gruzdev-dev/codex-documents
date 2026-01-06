package models

// Narrative Type: A human-readable summary of the resource conveying the essential clinical and business information for the resource.
type Narrative struct {
	Id     *string `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Status string  `json:"status" bson:"status"`             // generated | extensions | additional | empty
	Div    string  `json:"div" bson:"div"`                   // Limited xhtml content
}
