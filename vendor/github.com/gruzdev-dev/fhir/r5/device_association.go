package models

import (
	"encoding/json"
	"fmt"
)

// A record of association of a device.
type DeviceAssociation struct {
	Id                *string           `json:"id,omitempty" bson:"id,omitempty"`                                // Logical id of this artifact
	Meta              *Meta             `json:"meta,omitempty" bson:"meta,omitempty"`                            // Metadata about the resource
	ImplicitRules     *string           `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`         // A set of rules under which this content was created
	Language          *string           `json:"language,omitempty" bson:"language,omitempty"`                    // Language of the resource content
	Text              *Narrative        `json:"text,omitempty" bson:"text,omitempty"`                            // Text summary of the resource, for human interpretation
	Contained         []json.RawMessage `json:"contained,omitempty" bson:"contained,omitempty"`                  // Contained, inline Resources
	Identifier        []Identifier      `json:"identifier,omitempty" bson:"identifier,omitempty"`                // Instance identifier
	Device            *Reference        `json:"device" bson:"device"`                                            // Reference to the device that is being associated
	Relationship      []CodeableConcept `json:"relationship,omitempty" bson:"relationship,omitempty"`            // Describes the relationship between the device and subject
	Status            string            `json:"status" bson:"status"`                                            // active | inactive | entered-in-error | unknown
	StatusReason      []CodeableConcept `json:"statusReason,omitempty" bson:"status_reason,omitempty"`           // The reasons given for the current association status
	AssociationStatus *CodeableConcept  `json:"associationStatus,omitempty" bson:"association_status,omitempty"` // State of the device’s association
	Subject           *Reference        `json:"subject,omitempty" bson:"subject,omitempty"`                      // The entity(ies) that the device is on or associated with
	Focus             *Reference        `json:"focus,omitempty" bson:"focus,omitempty"`                          // The target of the association
	BodyStructure     *Reference        `json:"bodyStructure,omitempty" bson:"body_structure,omitempty"`         // Current anatomical location of the device in/on subject
	Period            *Period           `json:"period,omitempty" bson:"period,omitempty"`                        // Begin and end dates and times for the device association
}

func (r *DeviceAssociation) Validate() error {
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
	if r.Device == nil {
		return fmt.Errorf("field 'Device' is required")
	}
	if r.Device != nil {
		if err := r.Device.Validate(); err != nil {
			return fmt.Errorf("Device: %w", err)
		}
	}
	for i, item := range r.Relationship {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Relationship[%d]: %w", i, err)
		}
	}
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	for i, item := range r.StatusReason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("StatusReason[%d]: %w", i, err)
		}
	}
	if r.AssociationStatus != nil {
		if err := r.AssociationStatus.Validate(); err != nil {
			return fmt.Errorf("AssociationStatus: %w", err)
		}
	}
	if r.Subject != nil {
		if err := r.Subject.Validate(); err != nil {
			return fmt.Errorf("Subject: %w", err)
		}
	}
	if r.Focus != nil {
		if err := r.Focus.Validate(); err != nil {
			return fmt.Errorf("Focus: %w", err)
		}
	}
	if r.BodyStructure != nil {
		if err := r.BodyStructure.Validate(); err != nil {
			return fmt.Errorf("BodyStructure: %w", err)
		}
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	return nil
}
