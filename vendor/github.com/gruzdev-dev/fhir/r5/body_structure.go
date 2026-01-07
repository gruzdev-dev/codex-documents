package models

import (
	"encoding/json"
	"fmt"
)

// Record details about an anatomical structure.  This resource may be used when a coded concept does not provide the necessary detail needed for the use case.
type BodyStructure struct {
	ResourceType      string                           `json:"resourceType" bson:"resource_type"`                               // Type of resource
	Id                *string                          `json:"id,omitempty" bson:"id,omitempty"`                                // Logical id of this artifact
	Meta              *Meta                            `json:"meta,omitempty" bson:"meta,omitempty"`                            // Metadata about the resource
	ImplicitRules     *string                          `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`         // A set of rules under which this content was created
	Language          *string                          `json:"language,omitempty" bson:"language,omitempty"`                    // Language of the resource content
	Text              *Narrative                       `json:"text,omitempty" bson:"text,omitempty"`                            // Text summary of the resource, for human interpretation
	Contained         []json.RawMessage                `json:"contained,omitempty" bson:"contained,omitempty"`                  // Contained, inline Resources
	Identifier        []Identifier                     `json:"identifier,omitempty" bson:"identifier,omitempty"`                // Bodystructure identifier
	Active            *bool                            `json:"active,omitempty" bson:"active,omitempty"`                        // Whether this record is in active use
	IncludedStructure []BodyStructureIncludedStructure `json:"includedStructure" bson:"included_structure"`                     // Included anatomic location(s)
	ExcludedStructure []BodyStructureIncludedStructure `json:"excludedStructure,omitempty" bson:"excluded_structure,omitempty"` // Excluded anatomic locations(s)
	Description       *string                          `json:"description,omitempty" bson:"description,omitempty"`              // Text description
	Image             []Attachment                     `json:"image,omitempty" bson:"image,omitempty"`                          // Attached images
	Patient           *Reference                       `json:"patient" bson:"patient"`                                          // Who this is about
}

func (r *BodyStructure) Validate() error {
	if r.ResourceType != "BodyStructure" {
		return fmt.Errorf("invalid resourceType: expected 'BodyStructure', got '%s'", r.ResourceType)
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
	if len(r.IncludedStructure) < 1 {
		return fmt.Errorf("field 'IncludedStructure' must have at least 1 elements")
	}
	for i, item := range r.IncludedStructure {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("IncludedStructure[%d]: %w", i, err)
		}
	}
	for i, item := range r.ExcludedStructure {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ExcludedStructure[%d]: %w", i, err)
		}
	}
	for i, item := range r.Image {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Image[%d]: %w", i, err)
		}
	}
	if r.Patient == nil {
		return fmt.Errorf("field 'Patient' is required")
	}
	if r.Patient != nil {
		if err := r.Patient.Validate(); err != nil {
			return fmt.Errorf("Patient: %w", err)
		}
	}
	return nil
}

type BodyStructureIncludedStructure struct {
	Id                      *string                                                 `json:"id,omitempty" bson:"id,omitempty"`                                             // Unique id for inter-element referencing
	Structure               *CodeableConcept                                        `json:"structure" bson:"structure"`                                                   // Code that represents the included structure
	Laterality              *CodeableConcept                                        `json:"laterality,omitempty" bson:"laterality,omitempty"`                             // Code that represents the included structure laterality
	BodyLandmarkOrientation []BodyStructureIncludedStructureBodyLandmarkOrientation `json:"bodyLandmarkOrientation,omitempty" bson:"body_landmark_orientation,omitempty"` // Landmark relative location
	SpatialReference        []Reference                                             `json:"spatialReference,omitempty" bson:"spatial_reference,omitempty"`                // Cartesian reference for structure
	Image                   []Attachment                                            `json:"image,omitempty" bson:"image,omitempty"`                                       // Image(s) of structural aspects
	Qualifier               []CodeableConcept                                       `json:"qualifier,omitempty" bson:"qualifier,omitempty"`                               // Code that represents the included structure qualifier
	Morphology              *CodeableConcept                                        `json:"morphology,omitempty" bson:"morphology,omitempty"`                             // Kind of Structure
}

func (r *BodyStructureIncludedStructure) Validate() error {
	if r.Structure == nil {
		return fmt.Errorf("field 'Structure' is required")
	}
	if r.Structure != nil {
		if err := r.Structure.Validate(); err != nil {
			return fmt.Errorf("Structure: %w", err)
		}
	}
	if r.Laterality != nil {
		if err := r.Laterality.Validate(); err != nil {
			return fmt.Errorf("Laterality: %w", err)
		}
	}
	for i, item := range r.BodyLandmarkOrientation {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("BodyLandmarkOrientation[%d]: %w", i, err)
		}
	}
	for i, item := range r.SpatialReference {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SpatialReference[%d]: %w", i, err)
		}
	}
	for i, item := range r.Image {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Image[%d]: %w", i, err)
		}
	}
	for i, item := range r.Qualifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Qualifier[%d]: %w", i, err)
		}
	}
	if r.Morphology != nil {
		if err := r.Morphology.Validate(); err != nil {
			return fmt.Errorf("Morphology: %w", err)
		}
	}
	return nil
}

type BodyStructureIncludedStructureBodyLandmarkOrientation struct {
	Id                   *string                                                                     `json:"id,omitempty" bson:"id,omitempty"`                                       // Unique id for inter-element referencing
	LandmarkDescription  []CodeableConcept                                                           `json:"landmarkDescription,omitempty" bson:"landmark_description,omitempty"`    // Explanation of landmark
	ClockFacePosition    []CodeableConcept                                                           `json:"clockFacePosition,omitempty" bson:"clock_face_position,omitempty"`       // Clockface orientation
	DistanceFromLandmark []BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark `json:"distanceFromLandmark,omitempty" bson:"distance_from_landmark,omitempty"` // Landmark relative location
	SurfaceOrientation   []CodeableConcept                                                           `json:"surfaceOrientation,omitempty" bson:"surface_orientation,omitempty"`      // Relative landmark surface orientation
}

func (r *BodyStructureIncludedStructureBodyLandmarkOrientation) Validate() error {
	for i, item := range r.LandmarkDescription {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("LandmarkDescription[%d]: %w", i, err)
		}
	}
	for i, item := range r.ClockFacePosition {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ClockFacePosition[%d]: %w", i, err)
		}
	}
	for i, item := range r.DistanceFromLandmark {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("DistanceFromLandmark[%d]: %w", i, err)
		}
	}
	for i, item := range r.SurfaceOrientation {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SurfaceOrientation[%d]: %w", i, err)
		}
	}
	return nil
}

type BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark struct {
	Id     *string             `json:"id,omitempty" bson:"id,omitempty"`         // Unique id for inter-element referencing
	Device []CodeableReference `json:"device,omitempty" bson:"device,omitempty"` // Measurement device
	Value  []Quantity          `json:"value,omitempty" bson:"value,omitempty"`   // Measured distance from body landmark
}

func (r *BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark) Validate() error {
	for i, item := range r.Device {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Device[%d]: %w", i, err)
		}
	}
	for i, item := range r.Value {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Value[%d]: %w", i, err)
		}
	}
	return nil
}
