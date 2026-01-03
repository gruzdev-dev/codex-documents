package fhir

// ContactPoint Type: Details for all kinds of technology mediated contact points for a person or organization, including telephone, email, etc.
type ContactPoint struct {
	Id     *string `json:"id,omitempty"`     // Unique id for inter-element referencing
	System *string `json:"system,omitempty"` // phone | fax | email | pager | url | sms | other
	Value  *string `json:"value,omitempty"`  // The actual contact point details
	Use    *string `json:"use,omitempty"`    // home | work | temp | old | mobile - purpose of this contact point
	Rank   *int    `json:"rank,omitempty"`   // Specify preferred order of use (1 = highest)
	Period *Period `json:"period,omitempty"` // Time period when the contact point was/is in use
}
