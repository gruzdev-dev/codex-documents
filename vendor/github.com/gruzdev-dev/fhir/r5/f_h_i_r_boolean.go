package models

// boolean Type: Value of "true" or "false"
type FHIRBoolean struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"`       // xml:id (or equivalent in JSON)
	Value *bool   `json:"value,omitempty" bson:"value,omitempty"` // Primitive value for boolean
}

func (r *FHIRBoolean) Validate() error {
	return nil
}
