package models

// PrimitiveType Type: The base type for all re-useable types defined that have a simple property.
type PrimitiveType struct {
	Id *string `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
}

func (r *PrimitiveType) Validate() error {
	return nil
}
