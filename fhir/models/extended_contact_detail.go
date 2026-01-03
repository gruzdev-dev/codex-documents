package models

// ExtendedContactDetail Type: Specifies contact information for a specific purpose over a period of time, might be handled/monitored by a specific named person or organization.
type ExtendedContactDetail struct {
	Id           *string          `json:"id,omitempty"`           // Unique id for inter-element referencing
	Purpose      *CodeableConcept `json:"purpose,omitempty"`      // The type of contact
	Name         []HumanName      `json:"name,omitempty"`         // Name of an individual to contact
	Telecom      []ContactPoint   `json:"telecom,omitempty"`      // Contact details (e.g.phone/fax/url)
	Address      *Address         `json:"address,omitempty"`      // Address for the contact
	Organization *Reference       `json:"organization,omitempty"` // This contact detail is handled/monitored by a specific organization
	Period       *Period          `json:"period,omitempty"`       // Period that this contact was valid for usage
}
