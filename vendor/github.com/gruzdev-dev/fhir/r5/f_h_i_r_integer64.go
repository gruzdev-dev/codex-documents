package models

// integer64 Type: A very large whole number
type FHIRInteger64 struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"`       // xml:id (or equivalent in JSON)
	Value *int64  `json:"value,omitempty" bson:"value,omitempty"` // Primitive value for integer64
}

func (r *FHIRInteger64) Validate() error {
	return nil
}
