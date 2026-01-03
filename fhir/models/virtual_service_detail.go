package fhir

// VirtualServiceDetail Type: The set of values required to describe a virtual service's connection details, including some limitations of the service.
type VirtualServiceDetail struct {
	Id              *string  `json:"id,omitempty"`              // Unique id for inter-element referencing
	ChannelType     *Coding  `json:"channelType,omitempty"`     // Channel Type
	Address         any      `json:"address[x],omitempty"`      // Contact address/number
	AdditionalInfo  []string `json:"additionalInfo,omitempty"`  // Web address to see alternative connection details
	MaxParticipants *int     `json:"maxParticipants,omitempty"` // Maximum number of participants supported by the virtual service
	SessionKey      *string  `json:"sessionKey,omitempty"`      // Session Key required by the virtual service
}
