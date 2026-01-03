package fhir

// Narrative Type: A human-readable summary of the resource conveying the essential clinical and business information for the resource.
type Narrative struct {
	Id     *string `json:"id,omitempty"`     // Unique id for inter-element referencing
	Status string  `json:"status,omitempty"` // generated | extensions | additional | empty
	Div    string  `json:"div,omitempty"`    // Limited xhtml content
}
