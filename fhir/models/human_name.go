package models

import (
	"fmt"
)

// HumanName Type: A name, normally of a human, that can be used for other living entities (e.g. animals but not organizations) that have been assigned names by a human and may need the use of name parts or the need for usage information.
type HumanName struct {
	Id     *string  `json:"id,omitempty" bson:"id,omitempty"`         // Unique id for inter-element referencing
	Use    *string  `json:"use,omitempty" bson:"use,omitempty"`       // usual | official | temp | nickname | anonymous | old | maiden
	Text   *string  `json:"text,omitempty" bson:"text,omitempty"`     // Text representation of the full name
	Family *string  `json:"family,omitempty" bson:"family,omitempty"` // Family name (often called 'Surname')
	Given  []string `json:"given,omitempty" bson:"given,omitempty"`   // Given names (not always 'first'). Includes middle names
	Prefix []string `json:"prefix,omitempty" bson:"prefix,omitempty"` // Parts that come before the name
	Suffix []string `json:"suffix,omitempty" bson:"suffix,omitempty"` // Parts that come after the name
	Period *Period  `json:"period,omitempty" bson:"period,omitempty"` // Time period when name was/is in use
}

func (r *HumanName) Validate() error {
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	return nil
}
