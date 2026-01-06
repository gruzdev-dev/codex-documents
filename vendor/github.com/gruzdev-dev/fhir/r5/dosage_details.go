package models

import (
	"fmt"
)

// DosageDetails Type: Indicates how the medication is/was taken or should be taken by the patient.
type DosageDetails struct {
	Id                  *string             `json:"id,omitempty" bson:"id,omitempty"`                                    // Unique id for inter-element referencing
	RenderedInstruction *string             `json:"renderedInstruction,omitempty" bson:"rendered_instruction,omitempty"` // Full representation of the dosage instructions
	Simple              *Dosage             `json:"simple,omitempty" bson:"simple,omitempty"`                            // Dosage details if it is a simple dose
	Step                []DosageDetailsStep `json:"step,omitempty" bson:"step,omitempty"`                                // One step in a sequence of steps that comprise the dosage course
	Safety              *DosageSafety       `json:"safety,omitempty" bson:"safety,omitempty"`                            // Safety Information about the combined dose course
}

func (r *DosageDetails) Validate() error {
	if r.Simple != nil {
		if err := r.Simple.Validate(); err != nil {
			return fmt.Errorf("Simple: %w", err)
		}
	}
	for i, item := range r.Step {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Step[%d]: %w", i, err)
		}
	}
	if r.Safety != nil {
		if err := r.Safety.Validate(); err != nil {
			return fmt.Errorf("Safety: %w", err)
		}
	}
	return nil
}

type DosageDetailsStep struct {
	Id        *string       `json:"id,omitempty" bson:"id,omitempty"`         // Unique id for inter-element referencing
	Start     *RelativeTime `json:"start,omitempty" bson:"start,omitempty"`   // When the step starts
	End       *RelativeTime `json:"end,omitempty" bson:"end,omitempty"`       // When the step ends
	Count     *int          `json:"count,omitempty" bson:"count,omitempty"`   // How many times to do this step (if not 1)
	Component []Dosage      `json:"component" bson:"component"`               // A dosage details that is part of this step
	Safety    *DosageSafety `json:"safety,omitempty" bson:"safety,omitempty"` // Safety Information about this step of the dose course
}

func (r *DosageDetailsStep) Validate() error {
	if r.Start != nil {
		if err := r.Start.Validate(); err != nil {
			return fmt.Errorf("Start: %w", err)
		}
	}
	if r.End != nil {
		if err := r.End.Validate(); err != nil {
			return fmt.Errorf("End: %w", err)
		}
	}
	if len(r.Component) < 1 {
		return fmt.Errorf("field 'Component' must have at least 1 elements")
	}
	for i, item := range r.Component {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Component[%d]: %w", i, err)
		}
	}
	if r.Safety != nil {
		if err := r.Safety.Validate(); err != nil {
			return fmt.Errorf("Safety: %w", err)
		}
	}
	return nil
}
