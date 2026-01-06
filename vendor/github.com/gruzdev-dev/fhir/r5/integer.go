package models

// integer Type: A whole number
type integer struct {
}

func (r *integer) Validate() error {
	return nil
}

type Integer struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"`       // xml:id (or equivalent in JSON)
	Value *int    `json:"value,omitempty" bson:"value,omitempty"` // Primitive value for integer
}

func (r *Integer) Validate() error {
	return nil
}
