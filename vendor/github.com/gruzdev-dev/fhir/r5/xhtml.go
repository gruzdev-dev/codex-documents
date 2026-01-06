package models

import (
	"fmt"
)

// xhtml Type definition
type xhtml struct {
}

func (r *xhtml) Validate() error {
	return nil
}

type Xhtml struct {
	Id    *string `json:"id,omitempty" bson:"id,omitempty"` // xml:id (or equivalent in JSON)
	Value string  `json:"value" bson:"value"`               // Actual xhtml
}

func (r *Xhtml) Validate() error {
	var emptyString string
	if r.Value == emptyString {
		return fmt.Errorf("field 'Value' is required")
	}
	return nil
}
