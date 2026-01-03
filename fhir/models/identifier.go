package fhir

// Identifier Type: An identifier - identifies some entity uniquely and unambiguously. Typically this is used for business identifiers.
type Identifier struct {
	Id       *string          `json:"id,omitempty"`       // Unique id for inter-element referencing
	Use      *string          `json:"use,omitempty"`      // usual | official | temp | secondary | old (If known)
	Type     *CodeableConcept `json:"type,omitempty"`     // Description of identifier
	System   *string          `json:"system,omitempty"`   // The namespace for the identifier value
	Value    *string          `json:"value,omitempty"`    // The value that is unique
	Period   *Period          `json:"period,omitempty"`   // Time period when id is/was valid for use
	Assigner *Reference       `json:"assigner,omitempty"` // Organization that issued id (may be just text)
}
