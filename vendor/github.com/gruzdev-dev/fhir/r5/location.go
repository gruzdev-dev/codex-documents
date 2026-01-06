package models

import (
	"encoding/json"
	"fmt"
)

// Details and position information for a place where services are provided and resources and participants may be stored, found, contained, or accommodated.
type Location struct {
	Id                   *string                 `json:"id,omitempty" bson:"id,omitempty"`                                      // Logical id of this artifact
	Meta                 *Meta                   `json:"meta,omitempty" bson:"meta,omitempty"`                                  // Metadata about the resource
	ImplicitRules        *string                 `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`               // A set of rules under which this content was created
	Language             *string                 `json:"language,omitempty" bson:"language,omitempty"`                          // Language of the resource content
	Text                 *Narrative              `json:"text,omitempty" bson:"text,omitempty"`                                  // Text summary of the resource, for human interpretation
	Contained            []json.RawMessage       `json:"contained,omitempty" bson:"contained,omitempty"`                        // Contained, inline Resources
	Identifier           []Identifier            `json:"identifier,omitempty" bson:"identifier,omitempty"`                      // Unique code or number identifying the location to its users
	Status               *string                 `json:"status,omitempty" bson:"status,omitempty"`                              // active | suspended | inactive
	OperationalStatus    *Coding                 `json:"operationalStatus,omitempty" bson:"operational_status,omitempty"`       // The operational status of the location (typically only for a bed/room)
	Code                 []CodeableConcept       `json:"code,omitempty" bson:"code,omitempty"`                                  // Codes that identify this location
	Name                 *string                 `json:"name,omitempty" bson:"name,omitempty"`                                  // Name of the location as used by humans
	Alias                []string                `json:"alias,omitempty" bson:"alias,omitempty"`                                // A list of alternate names that the location is known as, or was known as, in the past
	Description          *string                 `json:"description,omitempty" bson:"description,omitempty"`                    // Additional details about the location that could be displayed as further information to identify the location beyond its name
	Mode                 *string                 `json:"mode,omitempty" bson:"mode,omitempty"`                                  // instance | kind
	Type                 []CodeableConcept       `json:"type,omitempty" bson:"type,omitempty"`                                  // Types of services available at this location
	Contact              []ExtendedContactDetail `json:"contact,omitempty" bson:"contact,omitempty"`                            // Official contact details for the location
	Address              *Address                `json:"address,omitempty" bson:"address,omitempty"`                            // Physical location
	Form                 *CodeableConcept        `json:"form,omitempty" bson:"form,omitempty"`                                  // Physical form of the location
	Position             *LocationPosition       `json:"position,omitempty" bson:"position,omitempty"`                          // The absolute geographic location
	ManagingOrganization *Reference              `json:"managingOrganization,omitempty" bson:"managing_organization,omitempty"` // Organization responsible for provisioning and upkeep
	PartOf               *Reference              `json:"partOf,omitempty" bson:"part_of,omitempty"`                             // Another Location this one is physically a part of
	Characteristic       []CodeableConcept       `json:"characteristic,omitempty" bson:"characteristic,omitempty"`              // Collection of characteristics (attributes)
	HoursOfOperation     *Availability           `json:"hoursOfOperation,omitempty" bson:"hours_of_operation,omitempty"`        // What days/times during a week is this location usually open (including exceptions)
	VirtualService       []VirtualServiceDetail  `json:"virtualService,omitempty" bson:"virtual_service,omitempty"`             // Connection details of a virtual service (e.g. conference call)
	Endpoint             []Reference             `json:"endpoint,omitempty" bson:"endpoint,omitempty"`                          // Technical endpoints providing access to services operated for the location
}

func (r *Location) Validate() error {
	if r.Meta != nil {
		if err := r.Meta.Validate(); err != nil {
			return fmt.Errorf("Meta: %w", err)
		}
	}
	if r.Text != nil {
		if err := r.Text.Validate(); err != nil {
			return fmt.Errorf("Text: %w", err)
		}
	}
	for i, item := range r.Identifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Identifier[%d]: %w", i, err)
		}
	}
	if r.OperationalStatus != nil {
		if err := r.OperationalStatus.Validate(); err != nil {
			return fmt.Errorf("OperationalStatus: %w", err)
		}
	}
	for i, item := range r.Code {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Code[%d]: %w", i, err)
		}
	}
	for i, item := range r.Type {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Type[%d]: %w", i, err)
		}
	}
	for i, item := range r.Contact {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Contact[%d]: %w", i, err)
		}
	}
	if r.Address != nil {
		if err := r.Address.Validate(); err != nil {
			return fmt.Errorf("Address: %w", err)
		}
	}
	if r.Form != nil {
		if err := r.Form.Validate(); err != nil {
			return fmt.Errorf("Form: %w", err)
		}
	}
	if r.Position != nil {
		if err := r.Position.Validate(); err != nil {
			return fmt.Errorf("Position: %w", err)
		}
	}
	if r.ManagingOrganization != nil {
		if err := r.ManagingOrganization.Validate(); err != nil {
			return fmt.Errorf("ManagingOrganization: %w", err)
		}
	}
	if r.PartOf != nil {
		if err := r.PartOf.Validate(); err != nil {
			return fmt.Errorf("PartOf: %w", err)
		}
	}
	for i, item := range r.Characteristic {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Characteristic[%d]: %w", i, err)
		}
	}
	if r.HoursOfOperation != nil {
		if err := r.HoursOfOperation.Validate(); err != nil {
			return fmt.Errorf("HoursOfOperation: %w", err)
		}
	}
	for i, item := range r.VirtualService {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("VirtualService[%d]: %w", i, err)
		}
	}
	for i, item := range r.Endpoint {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Endpoint[%d]: %w", i, err)
		}
	}
	return nil
}

type LocationPosition struct {
	Id        *string  `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Longitude float64  `json:"longitude" bson:"longitude"`                   // Longitude with WGS84 datum
	Latitude  float64  `json:"latitude" bson:"latitude"`                     // Latitude with WGS84 datum
	Altitude  *float64 `json:"altitude,omitempty" bson:"altitude,omitempty"` // Altitude with WGS84 datum
}

func (r *LocationPosition) Validate() error {
	if r.Longitude == 0 {
		return fmt.Errorf("field 'Longitude' is required")
	}
	if r.Latitude == 0 {
		return fmt.Errorf("field 'Latitude' is required")
	}
	return nil
}
