package models

// BackboneType Type: Base definition for the few data types that are allowed to carry modifier extensions.
type BackboneType struct {
	Id *string `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
}

func (r *BackboneType) Validate() error {
	return nil
}
