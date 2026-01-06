package models

import (
	"fmt"
)

// Address Type: An address expressed using postal conventions (as opposed to GPS or other location definition formats).  This data type may be used to convey addresses for use in delivering mail as well as for visiting locations which might not be valid for mail delivery.  There are a variety of postal address formats defined around the world. The ISO21090-codedString may be used to provide a coded representation of the contents of strings in an Address.
type Address struct {
	Id         *string  `json:"id,omitempty" bson:"id,omitempty"`                  // Unique id for inter-element referencing
	Use        *string  `json:"use,omitempty" bson:"use,omitempty"`                // home | work | temp | old | billing - purpose of this address
	Type       *string  `json:"type,omitempty" bson:"type,omitempty"`              // postal | physical | both
	Text       *string  `json:"text,omitempty" bson:"text,omitempty"`              // Text representation of the address
	Line       []string `json:"line,omitempty" bson:"line,omitempty"`              // Street name, number, direction & P.O. Box etc.
	City       *string  `json:"city,omitempty" bson:"city,omitempty"`              // Name of city, town etc.
	District   *string  `json:"district,omitempty" bson:"district,omitempty"`      // District name (aka county)
	State      *string  `json:"state,omitempty" bson:"state,omitempty"`            // Sub-unit of country (abbreviations ok)
	PostalCode *string  `json:"postalCode,omitempty" bson:"postal_code,omitempty"` // Postal code for area
	Country    *string  `json:"country,omitempty" bson:"country,omitempty"`        // Country (e.g. may be ISO 3166 2 or 3 letter code)
	Period     *Period  `json:"period,omitempty" bson:"period,omitempty"`          // Time period when address was/is in use
}

func (r *Address) Validate() error {
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	return nil
}
