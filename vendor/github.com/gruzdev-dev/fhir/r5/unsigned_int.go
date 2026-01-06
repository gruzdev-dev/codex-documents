package models

// unsignedInt type: An integer with a value that is not negative (e.g. >= 0)
type unsignedInt struct {
}

func (r *unsignedInt) Validate() error {
	return nil
}

type UnsignedInt struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"`       // xml:id (or equivalent in JSON)
	Value *string `json:"value,omitempty" bson:"value,omitempty"` // Primitive value for unsignedInt
}

func (r *UnsignedInt) Validate() error {
	return nil
}
