package models

import (
	"fmt"
)

// string Type: A sequence of Unicode characters
type FHIRString struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"`       // xml:id (or equivalent in JSON)
	Value *string `json:"value,omitempty" bson:"value,omitempty"` // Primitive value for string
}

func (r *FHIRString) Validate() error {
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
