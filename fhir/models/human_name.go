package models

// HumanName Type: A name, normally of a human, that can be used for other living entities (e.g. animals but not organizations) that have been assigned names by a human and may need the use of name parts or the need for usage information.
type HumanName struct {
	Id     *string  `json:"id,omitempty"`     // Unique id for inter-element referencing
	Use    *string  `json:"use,omitempty"`    // usual | official | temp | nickname | anonymous | old | maiden
	Text   *string  `json:"text,omitempty"`   // Text representation of the full name
	Family *string  `json:"family,omitempty"` // Family name (often called 'Surname')
	Given  []string `json:"given,omitempty"`  // Given names (not always 'first'). Includes middle names
	Prefix []string `json:"prefix,omitempty"` // Parts that come before the name
	Suffix []string `json:"suffix,omitempty"` // Parts that come after the name
	Period *Period  `json:"period,omitempty"` // Time period when name was/is in use
}
