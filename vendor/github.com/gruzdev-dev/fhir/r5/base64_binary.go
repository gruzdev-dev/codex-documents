package models

// base64Binary Type: A stream of bytes
type base64Binary struct {
}

func (r *base64Binary) Validate() error {
	return nil
}

type Base64Binary struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"`       // xml:id (or equivalent in JSON)
	Value *string `json:"value,omitempty" bson:"value,omitempty"` // Primitive value for base64Binary
}

func (r *Base64Binary) Validate() error {
	return nil
}
