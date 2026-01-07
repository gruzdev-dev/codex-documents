package models

// uuid type: A UUID, represented as a URI
type uuid struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"`       // xml:id (or equivalent in JSON)
	Value *string `json:"value,omitempty" bson:"value,omitempty"` // Primitive value for uuid
}

func (r *uuid) Validate() error {
	return nil
}
