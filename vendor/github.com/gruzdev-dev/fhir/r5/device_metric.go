package models

import (
	"encoding/json"
	"fmt"
)

// Describes a measurement, calculation or setting capability of a device. The DeviceMetric resource is derived from the ISO/IEEE 11073-10201 Domain Information Model standard, but is more widely applicable.
type DeviceMetric struct {
	ResourceType         string                    `json:"resourceType" bson:"resource_type"`                                     // Type of resource
	Id                   *string                   `json:"id,omitempty" bson:"id,omitempty"`                                      // Logical id of this artifact
	Meta                 *Meta                     `json:"meta,omitempty" bson:"meta,omitempty"`                                  // Metadata about the resource
	ImplicitRules        *string                   `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`               // A set of rules under which this content was created
	Language             *string                   `json:"language,omitempty" bson:"language,omitempty"`                          // Language of the resource content
	Text                 *Narrative                `json:"text,omitempty" bson:"text,omitempty"`                                  // Text summary of the resource, for human interpretation
	Contained            []json.RawMessage         `json:"contained,omitempty" bson:"contained,omitempty"`                        // Contained, inline Resources
	Identifier           []Identifier              `json:"identifier,omitempty" bson:"identifier,omitempty"`                      // Instance identifier
	Status               string                    `json:"status" bson:"status"`                                                  // active | inactive | entered-in-error | unknown
	OperationalStatus    *string                   `json:"operationalStatus,omitempty" bson:"operational_status,omitempty"`       // on | off | standby | unknown
	Category             *CodeableConcept          `json:"category" bson:"category"`                                              // The kind of metric represented
	Type                 *CodeableConcept          `json:"type" bson:"type"`                                                      // Identity of metric, for example Heart Rate or PEEP Setting
	Device               *Reference                `json:"device" bson:"device"`                                                  // The device to which this DeviceMetric applies
	Unit                 *CodeableConcept          `json:"unit,omitempty" bson:"unit,omitempty"`                                  // Unit of Measure for the Metric
	Color                *string                   `json:"color,omitempty" bson:"color,omitempty"`                                // Color name (from CSS4) or #RRGGBB code
	MeasurementFrequency *Quantity                 `json:"measurementFrequency,omitempty" bson:"measurement_frequency,omitempty"` // Indicates how often the metric is taken or recorded
	Availability         *CodeableConcept          `json:"availability,omitempty" bson:"availability,omitempty"`                  // The continuity of the metric (e.g., measurement)
	Calibration          []DeviceMetricCalibration `json:"calibration,omitempty" bson:"calibration,omitempty"`                    // Describes the calibrations that have been performed or that are required to be performed
}

func (r *DeviceMetric) Validate() error {
	if r.ResourceType != "DeviceMetric" {
		return fmt.Errorf("invalid resourceType: expected 'DeviceMetric', got '%s'", r.ResourceType)
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
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.Category == nil {
		return fmt.Errorf("field 'Category' is required")
	}
	if r.Category != nil {
		if err := r.Category.Validate(); err != nil {
			return fmt.Errorf("Category: %w", err)
		}
	}
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
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
	if r.Unit != nil {
		if err := r.Unit.Validate(); err != nil {
			return fmt.Errorf("Unit: %w", err)
		}
	}
	if r.MeasurementFrequency != nil {
		if err := r.MeasurementFrequency.Validate(); err != nil {
			return fmt.Errorf("MeasurementFrequency: %w", err)
		}
	}
	if r.Availability != nil {
		if err := r.Availability.Validate(); err != nil {
			return fmt.Errorf("Availability: %w", err)
		}
	}
	for i, item := range r.Calibration {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Calibration[%d]: %w", i, err)
		}
	}
	return nil
}

type DeviceMetricCalibration struct {
	Id    *string          `json:"id,omitempty" bson:"id,omitempty"`       // Unique id for inter-element referencing
	Type  *CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`   // The method of calibration
	State *string          `json:"state,omitempty" bson:"state,omitempty"` // not-calibrated | calibration-required | calibrated | unspecified
	Time  *string          `json:"time,omitempty" bson:"time,omitempty"`   // Describes the time last calibration has been performed
}

func (r *DeviceMetricCalibration) Validate() error {
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	return nil
}
