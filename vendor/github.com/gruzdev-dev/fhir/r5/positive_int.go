package models

// positiveInt type: An integer with a value that is positive (e.g. >0)
type positiveInt struct {
}

func (r *positiveInt) Validate() error {
	return nil
}

type PositiveInt struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"`       // xml:id (or equivalent in JSON)
	Value *string `json:"value,omitempty" bson:"value,omitempty"` // Primitive value for positiveInt
}

func (r *PositiveInt) Validate() error {
	return nil
}
