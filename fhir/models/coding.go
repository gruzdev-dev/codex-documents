package models

// Coding Type: A reference to a code defined by a terminology system.
type Coding struct {
	Id           *string `json:"id,omitempty"`           // Unique id for inter-element referencing
	System       *string `json:"system,omitempty"`       // Identity of the terminology system
	Version      *string `json:"version,omitempty"`      // Version of the system - if relevant
	Code         *string `json:"code,omitempty"`         // Symbol in syntax defined by the system
	Display      *string `json:"display,omitempty"`      // Representation defined by the system
	UserSelected bool    `json:"userSelected,omitempty"` // If this coding was chosen directly by the user
}
