package fhir

// Signature Type: A signature along with supporting context. The signature may be a digital signature that is cryptographic in nature, or some other signature acceptable to the domain. This other signature may be as simple as a graphical image representing a hand-written signature, or a signature ceremony Different signature approaches have different utilities.
type Signature struct {
	Id           *string    `json:"id,omitempty"`           // Unique id for inter-element referencing
	Type         []Coding   `json:"type,omitempty"`         // Indication of the reason the entity signed the object(s)
	When         *string    `json:"when,omitempty"`         // When the signature was created
	Who          *Reference `json:"who,omitempty"`          // Who signed
	OnBehalfOf   *Reference `json:"onBehalfOf,omitempty"`   // The party represented
	TargetFormat *string    `json:"targetFormat,omitempty"` // The technical format of the signed resources
	SigFormat    *string    `json:"sigFormat,omitempty"`    // The technical format of the signature
	Data         *string    `json:"data,omitempty"`         // The actual signature content (XML Signature, JSON Jose, picture, etc.)
}
