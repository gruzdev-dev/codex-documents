package models

// code type: A string which has at least one character and no leading or trailing whitespace and where there is no whitespace other than single spaces in the contents
type code struct {
}

func (r *code) Validate() error {
	return nil
}

type Code struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"`       // xml:id (or equivalent in JSON)
	Value *string `json:"value,omitempty" bson:"value,omitempty"` // Primitive value for code
}

func (r *Code) Validate() error {
	return nil
}
