package models

// VirtualServiceDetail Type: The set of values required to describe a virtual service's connection details, including some limitations of the service.
type VirtualServiceDetail struct {
	Id                           *string                `json:"id,omitempty"`                           // Unique id for inter-element referencing
	ChannelType                  *Coding                `json:"channelType,omitempty"`                  // Channel Type
	AddressUrl                   *string                `json:"addressUrl,omitempty"`                   // Contact address/number
	AddressString                *string                `json:"addressString,omitempty"`                // Contact address/number
	AddressContactPoint          *ContactPoint          `json:"addressContactPoint,omitempty"`          // Contact address/number
	AddressExtendedContactDetail *ExtendedContactDetail `json:"addressExtendedContactDetail,omitempty"` // Contact address/number
	AdditionalInfo               []string               `json:"additionalInfo,omitempty"`               // Web address to see alternative connection details
	MaxParticipants              *int                   `json:"maxParticipants,omitempty"`              // Maximum number of participants supported by the virtual service
	SessionKey                   *string                `json:"sessionKey,omitempty"`                   // Session Key required by the virtual service
}
