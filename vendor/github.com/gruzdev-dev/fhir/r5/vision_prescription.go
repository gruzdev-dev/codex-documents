package models

import (
	"encoding/json"
	"fmt"
)

// An authorization for the provision of glasses and/or contact lenses to a patient.
type VisionPrescription struct {
	ResourceType      string                                `json:"resourceType" bson:"resource_type"`                           // Type of resource
	Id                *string                               `json:"id,omitempty" bson:"id,omitempty"`                            // Logical id of this artifact
	Meta              *Meta                                 `json:"meta,omitempty" bson:"meta,omitempty"`                        // Metadata about the resource
	ImplicitRules     *string                               `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`     // A set of rules under which this content was created
	Language          *string                               `json:"language,omitempty" bson:"language,omitempty"`                // Language of the resource content
	Text              *Narrative                            `json:"text,omitempty" bson:"text,omitempty"`                        // Text summary of the resource, for human interpretation
	Contained         []json.RawMessage                     `json:"contained,omitempty" bson:"contained,omitempty"`              // Contained, inline Resources
	Identifier        []Identifier                          `json:"identifier,omitempty" bson:"identifier,omitempty"`            // Business Identifier for vision prescription
	BasedOn           []Reference                           `json:"basedOn,omitempty" bson:"based_on,omitempty"`                 // What prescription fulfills
	GroupIdentifier   *Identifier                           `json:"groupIdentifier,omitempty" bson:"group_identifier,omitempty"` // Composite request identifier
	Status            string                                `json:"status" bson:"status"`                                        // active | cancelled | draft | entered-in-error
	Priority          *string                               `json:"priority,omitempty" bson:"priority,omitempty"`                // routine | urgent | asap | stat
	Created           string                                `json:"created" bson:"created"`                                      // Response creation date
	Patient           *Reference                            `json:"patient" bson:"patient"`                                      // Who prescription is for
	Encounter         *Reference                            `json:"encounter,omitempty" bson:"encounter,omitempty"`              // Created during encounter / admission / stay
	DateWritten       string                                `json:"dateWritten" bson:"date_written"`                             // When prescription was authorized
	Prescriber        *Reference                            `json:"prescriber" bson:"prescriber"`                                // Who authorized the vision prescription
	LensSpecification []VisionPrescriptionLensSpecification `json:"lensSpecification" bson:"lens_specification"`                 // Vision lens authorization
}

func (r *VisionPrescription) Validate() error {
	if r.ResourceType != "VisionPrescription" {
		return fmt.Errorf("invalid resourceType: expected 'VisionPrescription', got '%s'", r.ResourceType)
	}
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
	for i, item := range r.BasedOn {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("BasedOn[%d]: %w", i, err)
		}
	}
	if r.GroupIdentifier != nil {
		if err := r.GroupIdentifier.Validate(); err != nil {
			return fmt.Errorf("GroupIdentifier: %w", err)
		}
	}
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.Created == emptyString {
		return fmt.Errorf("field 'Created' is required")
	}
	if r.Patient == nil {
		return fmt.Errorf("field 'Patient' is required")
	}
	if r.Patient != nil {
		if err := r.Patient.Validate(); err != nil {
			return fmt.Errorf("Patient: %w", err)
		}
	}
	if r.Encounter != nil {
		if err := r.Encounter.Validate(); err != nil {
			return fmt.Errorf("Encounter: %w", err)
		}
	}
	if r.DateWritten == emptyString {
		return fmt.Errorf("field 'DateWritten' is required")
	}
	if r.Prescriber == nil {
		return fmt.Errorf("field 'Prescriber' is required")
	}
	if r.Prescriber != nil {
		if err := r.Prescriber.Validate(); err != nil {
			return fmt.Errorf("Prescriber: %w", err)
		}
	}
	if len(r.LensSpecification) < 1 {
		return fmt.Errorf("field 'LensSpecification' must have at least 1 elements")
	}
	for i, item := range r.LensSpecification {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("LensSpecification[%d]: %w", i, err)
		}
	}
	return nil
}

type VisionPrescriptionLensSpecification struct {
	Id        *string                                    `json:"id,omitempty" bson:"id,omitempty"`                // Unique id for inter-element referencing
	Product   *CodeableConcept                           `json:"product" bson:"product"`                          // Product to be supplied
	Eye       string                                     `json:"eye" bson:"eye"`                                  // right | left
	Sphere    *float64                                   `json:"sphere,omitempty" bson:"sphere,omitempty"`        // Power of the lens
	Cylinder  *float64                                   `json:"cylinder,omitempty" bson:"cylinder,omitempty"`    // Lens power for astigmatism
	Axis      *int                                       `json:"axis,omitempty" bson:"axis,omitempty"`            // Lens meridian which contain no power for astigmatism
	Prism     []VisionPrescriptionLensSpecificationPrism `json:"prism,omitempty" bson:"prism,omitempty"`          // Eye alignment compensation
	Add       *float64                                   `json:"add,omitempty" bson:"add,omitempty"`              // Added power for multifocal levels
	Power     *float64                                   `json:"power,omitempty" bson:"power,omitempty"`          // Contact lens power
	BackCurve *float64                                   `json:"backCurve,omitempty" bson:"back_curve,omitempty"` // Contact lens back curvature
	Diameter  *float64                                   `json:"diameter,omitempty" bson:"diameter,omitempty"`    // Contact lens diameter
	Duration  *Quantity                                  `json:"duration,omitempty" bson:"duration,omitempty"`    // Lens wear duration
	Color     *string                                    `json:"color,omitempty" bson:"color,omitempty"`          // Color required
	Brand     *string                                    `json:"brand,omitempty" bson:"brand,omitempty"`          // Brand required
	Note      []Annotation                               `json:"note,omitempty" bson:"note,omitempty"`            // Notes for coatings
}

func (r *VisionPrescriptionLensSpecification) Validate() error {
	if r.Product == nil {
		return fmt.Errorf("field 'Product' is required")
	}
	if r.Product != nil {
		if err := r.Product.Validate(); err != nil {
			return fmt.Errorf("Product: %w", err)
		}
	}
	var emptyString string
	if r.Eye == emptyString {
		return fmt.Errorf("field 'Eye' is required")
	}
	for i, item := range r.Prism {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Prism[%d]: %w", i, err)
		}
	}
	if r.Duration != nil {
		if err := r.Duration.Validate(); err != nil {
			return fmt.Errorf("Duration: %w", err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	return nil
}

type VisionPrescriptionLensSpecificationPrism struct {
	Id     *string `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Amount float64 `json:"amount" bson:"amount"`             // Amount of adjustment
	Base   string  `json:"base" bson:"base"`                 // up | down | in | out
}

func (r *VisionPrescriptionLensSpecificationPrism) Validate() error {
	if r.Amount == 0 {
		return fmt.Errorf("field 'Amount' is required")
	}
	var emptyString string
	if r.Base == emptyString {
		return fmt.Errorf("field 'Base' is required")
	}
	return nil
}
