package models

// canonical type: A URI that is a reference to a canonical URL on a FHIR resource
type FHIRString struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"`       // xml:id (or equivalent in JSON)
	Value *string `json:"value,omitempty" bson:"value,omitempty"` // Primitive value for canonical
}

func (r *FHIRString) Validate() error {
	return nil
}
