package models

// unsignedInt type: An integer with a value that is not negative (e.g. >= 0)
type FHIRInteger struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"`       // xml:id (or equivalent in JSON)
	Value *int    `json:"value,omitempty" bson:"value,omitempty"` // Primitive value for unsignedInt
}

func (r *FHIRInteger) Validate() error {
	return nil
}
