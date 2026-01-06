package models

// Element Type: Base definition for all elements in a resource.
type Element struct {
	Id *string `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
}

func (r *Element) Validate() error {
	return nil
}
