package models

// integer Type: A whole number
type FHIRInteger struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"`       // xml:id (or equivalent in JSON)
	Value *int    `json:"value,omitempty" bson:"value,omitempty"` // Primitive value for integer
}

func (r *FHIRInteger) Validate() error {
	return nil
}
