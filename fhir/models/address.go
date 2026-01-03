package fhir

// Address Type: An address expressed using postal conventions (as opposed to GPS or other location definition formats).  This data type may be used to convey addresses for use in delivering mail as well as for visiting locations which might not be valid for mail delivery.  There are a variety of postal address formats defined around the world. The ISO21090-codedString may be used to provide a coded representation of the contents of strings in an Address.
type Address struct {
	Id         *string  `json:"id,omitempty"`         // Unique id for inter-element referencing
	Use        *string  `json:"use,omitempty"`        // home | work | temp | old | billing - purpose of this address
	Type       *string  `json:"type,omitempty"`       // postal | physical | both
	Text       *string  `json:"text,omitempty"`       // Text representation of the address
	Line       []string `json:"line,omitempty"`       // Street name, number, direction & P.O. Box etc.
	City       *string  `json:"city,omitempty"`       // Name of city, town etc.
	District   *string  `json:"district,omitempty"`   // District name (aka county)
	State      *string  `json:"state,omitempty"`      // Sub-unit of country (abbreviations ok)
	PostalCode *string  `json:"postalCode,omitempty"` // Postal code for area
	Country    *string  `json:"country,omitempty"`    // Country (e.g. may be ISO 3166 2 or 3 letter code)
	Period     *Period  `json:"period,omitempty"`     // Time period when address was/is in use
}
