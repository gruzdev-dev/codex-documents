package models

// Reference Type: A reference from one resource to another.
type Reference struct {
	Id         *string     `json:"id,omitempty"`         // Unique id for inter-element referencing
	Reference  *string     `json:"reference,omitempty"`  // Literal reference, Relative, internal or absolute URL
	Type       *string     `json:"type,omitempty"`       // Type the reference refers to (e.g. "Patient") - must be a resource in resources
	Identifier *Identifier `json:"identifier,omitempty"` // Logical reference, when literal reference is not known
	Display    *string     `json:"display,omitempty"`    // Text alternative for the resource
}
