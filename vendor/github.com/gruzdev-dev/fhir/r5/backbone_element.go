package models

// BackboneElement Type: Base definition for all elements that are defined inside a resource - but not those in a data type.
type BackboneElement struct {
	Id *string `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
}

func (r *BackboneElement) Validate() error {
	return nil
}
