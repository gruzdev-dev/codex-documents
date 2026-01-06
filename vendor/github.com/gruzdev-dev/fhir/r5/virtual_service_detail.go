package models

import (
	"fmt"
)

// VirtualServiceDetail Type: The set of values required to describe a virtual service's connection details, including some limitations of the service.
type VirtualServiceDetail struct {
	Id                           *string                `json:"id,omitempty" bson:"id,omitempty"`                                                        // Unique id for inter-element referencing
	ChannelType                  *Coding                `json:"channelType,omitempty" bson:"channel_type,omitempty"`                                     // Channel Type
	AddressUrl                   *string                `json:"addressUrl,omitempty" bson:"address_url,omitempty"`                                       // Contact address/number
	AddressString                *string                `json:"addressString,omitempty" bson:"address_string,omitempty"`                                 // Contact address/number
	AddressContactPoint          *ContactPoint          `json:"addressContactPoint,omitempty" bson:"address_contact_point,omitempty"`                    // Contact address/number
	AddressExtendedContactDetail *ExtendedContactDetail `json:"addressExtendedContactDetail,omitempty" bson:"address_extended_contact_detail,omitempty"` // Contact address/number
	AdditionalInfo               []string               `json:"additionalInfo,omitempty" bson:"additional_info,omitempty"`                               // Web address to see alternative connection details
	MaxParticipants              *int                   `json:"maxParticipants,omitempty" bson:"max_participants,omitempty"`                             // Maximum number of participants supported by the virtual service
	SessionKey                   *string                `json:"sessionKey,omitempty" bson:"session_key,omitempty"`                                       // Session Key required by the virtual service
}

func (r *VirtualServiceDetail) Validate() error {
	if r.ChannelType != nil {
		if err := r.ChannelType.Validate(); err != nil {
			return fmt.Errorf("ChannelType: %w", err)
		}
	}
	if r.AddressContactPoint != nil {
		if err := r.AddressContactPoint.Validate(); err != nil {
			return fmt.Errorf("AddressContactPoint: %w", err)
		}
	}
	if r.AddressExtendedContactDetail != nil {
		if err := r.AddressExtendedContactDetail.Validate(); err != nil {
			return fmt.Errorf("AddressExtendedContactDetail: %w", err)
		}
	}
	return nil
}
