package models

// boolean Type: Value of "true" or "false"
type boolean struct {
}

func (r *boolean) Validate() error {
	return nil
}

type Boolean struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"`       // xml:id (or equivalent in JSON)
	Value bool    `json:"value,omitempty" bson:"value,omitempty"` // Primitive value for boolean
}

func (r *Boolean) Validate() error {
	return nil
}
