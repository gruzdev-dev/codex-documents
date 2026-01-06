package models

import (
	"fmt"
)

// string Type: A sequence of Unicode characters
type string struct {
}

func (r *string) Validate() error {
	return nil
}

type String struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"`       // xml:id (or equivalent in JSON)
	Value *string `json:"value,omitempty" bson:"value,omitempty"` // Primitive value for string
}

func (r *String) Validate() error {
	var emptyString string
	if r.Value != nil {
		valStr := ""
		if *r.Value != emptyString {
			valStr = fmt.Sprint(*r.Value)
		}
		if len(valStr) > 1048576 {
			return fmt.Errorf("field 'Value' exceeds maxLength 1048576")
		}
	}
	return nil
}
