package models

import (
	"encoding/json"
	"fmt"
)

// This resource describes the properties (regulated, has real time clock, etc.), administrative (manufacturer name, model number, serial number, firmware, etc.), and type (knee replacement, blood pressure cuff, MRI, etc.) of a physical unit (these values do not change much within a given module, for example the serial number, manufacturer name, and model number). An actual unit may consist of several modules in a distinct hierarchy and these are represented by multiple Device resources and bound through the 'parent' element.
type Device struct {
	Id                    *string               `json:"id,omitempty" bson:"id,omitempty"`                                         // Logical id of this artifact
	Meta                  *Meta                 `json:"meta,omitempty" bson:"meta,omitempty"`                                     // Metadata about the resource
	ImplicitRules         *string               `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                  // A set of rules under which this content was created
	Language              *string               `json:"language,omitempty" bson:"language,omitempty"`                             // Language of the resource content
	Text                  *Narrative            `json:"text,omitempty" bson:"text,omitempty"`                                     // Text summary of the resource, for human interpretation
	Contained             []json.RawMessage     `json:"contained,omitempty" bson:"contained,omitempty"`                           // Contained, inline Resources
	Identifier            []Identifier          `json:"identifier,omitempty" bson:"identifier,omitempty"`                         // Instance identifier
	Definition            *string               `json:"definition,omitempty" bson:"definition,omitempty"`                         // The reference to the definition for the device
	UdiCarrier            []DeviceUdiCarrier    `json:"udiCarrier,omitempty" bson:"udi_carrier,omitempty"`                        // Unique Device Identifier (UDI) value
	Status                *string               `json:"status,omitempty" bson:"status,omitempty"`                                 // active | inactive | entered-in-error | unknown
	AvailabilityStatus    *CodeableConcept      `json:"availabilityStatus,omitempty" bson:"availability_status,omitempty"`        // lost | damaged | destroyed | available
	BiologicalSourceEvent *Identifier           `json:"biologicalSourceEvent,omitempty" bson:"biological_source_event,omitempty"` // A production identifier of the donation, collection, or pooling event from which biological material in this device was derived
	Manufacturer          *string               `json:"manufacturer,omitempty" bson:"manufacturer,omitempty"`                     // Name of device manufacturer
	ManufactureDate       *string               `json:"manufactureDate,omitempty" bson:"manufacture_date,omitempty"`              // A production identifier that indicates the date when the device was made
	ExpirationDate        *string               `json:"expirationDate,omitempty" bson:"expiration_date,omitempty"`                // A production identifier that indicates the date and time of expiry of this device (if applicable)
	LotNumber             *string               `json:"lotNumber,omitempty" bson:"lot_number,omitempty"`                          // A production identifier that indicates the Lot number of manufacture
	SerialNumber          *string               `json:"serialNumber,omitempty" bson:"serial_number,omitempty"`                    // A production identifier that indicates the Serial number assigned by the manufacturer
	Name                  []DeviceName          `json:"name,omitempty" bson:"name,omitempty"`                                     // The name or names of the device as known to the manufacturer and/or patient
	ModelNumber           *string               `json:"modelNumber,omitempty" bson:"model_number,omitempty"`                      // The manufacturer's model number for the device
	PartNumber            *string               `json:"partNumber,omitempty" bson:"part_number,omitempty"`                        // The part number or catalog number of the device
	Category              []CodeableConcept     `json:"category,omitempty" bson:"category,omitempty"`                             // Indicates a high-level grouping of the device
	Type                  []CodeableConcept     `json:"type,omitempty" bson:"type,omitempty"`                                     // The kind or type of device
	DeviceVersion         []DeviceDeviceVersion `json:"deviceVersion,omitempty" bson:"device_version,omitempty"`                  // The actual design of the device or software version running on the device
	ConformsTo            []DeviceConformsTo    `json:"conformsTo,omitempty" bson:"conforms_to,omitempty"`                        // Identifies the standards, specifications, or formal guidances for the capabilities supported by the device
	Property              []DeviceProperty      `json:"property,omitempty" bson:"property,omitempty"`                             // Inherent, essentially fixed, characteristics of the device.  e.g., time properties, size, material, etc.
	Additive              []DeviceAdditive      `json:"additive,omitempty" bson:"additive,omitempty"`                             // Material added to a container device
	Contact               []ContactPoint        `json:"contact,omitempty" bson:"contact,omitempty"`                               // Details for human/organization for support
	Location              *Reference            `json:"location,omitempty" bson:"location,omitempty"`                             // Where the device is found
	Note                  []Annotation          `json:"note,omitempty" bson:"note,omitempty"`                                     // Device notes and comments
	Safety                []CodeableConcept     `json:"safety,omitempty" bson:"safety,omitempty"`                                 // Safety Characteristics of Device
	Parent                *Reference            `json:"parent,omitempty" bson:"parent,omitempty"`                                 // The higher level or encompassing device that this device is a logical part of
}

func (r *Device) Validate() error {
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
	for i, item := range r.UdiCarrier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("UdiCarrier[%d]: %w", i, err)
		}
	}
	if r.AvailabilityStatus != nil {
		if err := r.AvailabilityStatus.Validate(); err != nil {
			return fmt.Errorf("AvailabilityStatus: %w", err)
		}
	}
	if r.BiologicalSourceEvent != nil {
		if err := r.BiologicalSourceEvent.Validate(); err != nil {
			return fmt.Errorf("BiologicalSourceEvent: %w", err)
		}
	}
	for i, item := range r.Name {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Name[%d]: %w", i, err)
		}
	}
	for i, item := range r.Category {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Category[%d]: %w", i, err)
		}
	}
	for i, item := range r.Type {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Type[%d]: %w", i, err)
		}
	}
	for i, item := range r.DeviceVersion {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("DeviceVersion[%d]: %w", i, err)
		}
	}
	for i, item := range r.ConformsTo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ConformsTo[%d]: %w", i, err)
		}
	}
	for i, item := range r.Property {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Property[%d]: %w", i, err)
		}
	}
	for i, item := range r.Additive {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Additive[%d]: %w", i, err)
		}
	}
	for i, item := range r.Contact {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Contact[%d]: %w", i, err)
		}
	}
	if r.Location != nil {
		if err := r.Location.Validate(); err != nil {
			return fmt.Errorf("Location: %w", err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	for i, item := range r.Safety {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Safety[%d]: %w", i, err)
		}
	}
	if r.Parent != nil {
		if err := r.Parent.Validate(); err != nil {
			return fmt.Errorf("Parent: %w", err)
		}
	}
	return nil
}

type DeviceDeviceVersion struct {
	Id          *string          `json:"id,omitempty" bson:"id,omitempty"`                    // Unique id for inter-element referencing
	Type        *CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`                // The type of the device version, e.g. manufacturer, approved, internal
	Component   *Identifier      `json:"component,omitempty" bson:"component,omitempty"`      // The hardware or software module of the device to which the version applies
	InstallDate *string          `json:"installDate,omitempty" bson:"install_date,omitempty"` // The date the version was installed on the device
	Value       string           `json:"value" bson:"value"`                                  // The version text
}

func (r *DeviceDeviceVersion) Validate() error {
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Component != nil {
		if err := r.Component.Validate(); err != nil {
			return fmt.Errorf("Component: %w", err)
		}
	}
	var emptyString string
	if r.Value == emptyString {
		return fmt.Errorf("field 'Value' is required")
	}
	return nil
}

type DeviceConformsTo struct {
	Id            *string          `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Category      *CodeableConcept `json:"category,omitempty" bson:"category,omitempty"` // Describes the common type of the standard, specification, or formal guidance.  communication | performance | measurement
	Specification *CodeableConcept `json:"specification" bson:"specification"`           // Identifies the standard, specification, or formal guidance that the device adheres to
	Version       *string          `json:"version,omitempty" bson:"version,omitempty"`   // Specific form or variant of the standard
}

func (r *DeviceConformsTo) Validate() error {
	if r.Category != nil {
		if err := r.Category.Validate(); err != nil {
			return fmt.Errorf("Category: %w", err)
		}
	}
	if r.Specification == nil {
		return fmt.Errorf("field 'Specification' is required")
	}
	if r.Specification != nil {
		if err := r.Specification.Validate(); err != nil {
			return fmt.Errorf("Specification: %w", err)
		}
	}
	return nil
}

type DeviceProperty struct {
	Id                   *string          `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Type                 *CodeableConcept `json:"type" bson:"type"`                                   // Code that specifies the property being represented
	ValueQuantity        *Quantity        `json:"valueQuantity" bson:"value_quantity"`                // Value of the property
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept" bson:"value_codeable_concept"` // Value of the property
	ValueString          *string          `json:"valueString" bson:"value_string"`                    // Value of the property
	ValueBoolean         *bool            `json:"valueBoolean" bson:"value_boolean"`                  // Value of the property
	ValueInteger         *int             `json:"valueInteger" bson:"value_integer"`                  // Value of the property
	ValueRange           *Range           `json:"valueRange" bson:"value_range"`                      // Value of the property
	ValueAttachment      *Attachment      `json:"valueAttachment" bson:"value_attachment"`            // Value of the property
}

func (r *DeviceProperty) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.ValueQuantity == nil {
		return fmt.Errorf("field 'ValueQuantity' is required")
	}
	if r.ValueQuantity != nil {
		if err := r.ValueQuantity.Validate(); err != nil {
			return fmt.Errorf("ValueQuantity: %w", err)
		}
	}
	if r.ValueCodeableConcept == nil {
		return fmt.Errorf("field 'ValueCodeableConcept' is required")
	}
	if r.ValueCodeableConcept != nil {
		if err := r.ValueCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("ValueCodeableConcept: %w", err)
		}
	}
	if r.ValueString == nil {
		return fmt.Errorf("field 'ValueString' is required")
	}
	if r.ValueBoolean == nil {
		return fmt.Errorf("field 'ValueBoolean' is required")
	}
	if r.ValueInteger == nil {
		return fmt.Errorf("field 'ValueInteger' is required")
	}
	if r.ValueRange == nil {
		return fmt.Errorf("field 'ValueRange' is required")
	}
	if r.ValueRange != nil {
		if err := r.ValueRange.Validate(); err != nil {
			return fmt.Errorf("ValueRange: %w", err)
		}
	}
	if r.ValueAttachment == nil {
		return fmt.Errorf("field 'ValueAttachment' is required")
	}
	if r.ValueAttachment != nil {
		if err := r.ValueAttachment.Validate(); err != nil {
			return fmt.Errorf("ValueAttachment: %w", err)
		}
	}
	return nil
}

type DeviceAdditive struct {
	Id       *string            `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Type     *CodeableReference `json:"type" bson:"type"`                             // The additive substance
	Quantity *Quantity          `json:"quantity,omitempty" bson:"quantity,omitempty"` // Quantity of additive substance within container
}

func (r *DeviceAdditive) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Quantity != nil {
		if err := r.Quantity.Validate(); err != nil {
			return fmt.Errorf("Quantity: %w", err)
		}
	}
	return nil
}

type DeviceUdiCarrier struct {
	Id                     *string `json:"id,omitempty" bson:"id,omitempty"`                                           // Unique id for inter-element referencing
	DeviceIdentifier       string  `json:"deviceIdentifier" bson:"device_identifier"`                                  // Mandatory fixed portion of UDI
	DeviceIdentifierSystem *string `json:"deviceIdentifierSystem,omitempty" bson:"device_identifier_system,omitempty"` // The namespace for the device identifier value
	Issuer                 string  `json:"issuer" bson:"issuer"`                                                       // UDI Issuing Organization
	Jurisdiction           *string `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                       // Regional UDI authority
	CarrierAIDC            *string `json:"carrierAIDC,omitempty" bson:"carrier_a_i_d_c,omitempty"`                     // UDI Machine Readable value
	CarrierHRF             *string `json:"carrierHRF,omitempty" bson:"carrier_h_r_f,omitempty"`                        // UDI Human Readable value
	EntryType              *string `json:"entryType,omitempty" bson:"entry_type,omitempty"`                            // barcode | rfid | manual | card | self-reported | electronic-transmission | unknown
}

func (r *DeviceUdiCarrier) Validate() error {
	var emptyString string
	if r.DeviceIdentifier == emptyString {
		return fmt.Errorf("field 'DeviceIdentifier' is required")
	}
	if r.Issuer == emptyString {
		return fmt.Errorf("field 'Issuer' is required")
	}
	return nil
}

type DeviceName struct {
	Id      *string          `json:"id,omitempty" bson:"id,omitempty"`           // Unique id for inter-element referencing
	Value   string           `json:"value" bson:"value"`                         // The term that names the device
	Type    *CodeableConcept `json:"type" bson:"type"`                           // registered-name | user-friendly-name | patient-reported-name
	Display bool             `json:"display,omitempty" bson:"display,omitempty"` // The preferred device name
}

func (r *DeviceName) Validate() error {
	var emptyString string
	if r.Value == emptyString {
		return fmt.Errorf("field 'Value' is required")
	}
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	return nil
}
