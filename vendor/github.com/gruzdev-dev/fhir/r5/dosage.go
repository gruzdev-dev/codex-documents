package models

import (
	"fmt"
)

// Dosage Type: Indicates how the medication is/was taken or should be taken by the patient.
type Dosage struct {
	Id                    *string             `json:"id,omitempty" bson:"id,omitempty"`                                        // Unique id for inter-element referencing
	Condition             []DosageCondition   `json:"condition,omitempty" bson:"condition,omitempty"`                          // Whether the dosage applies (e.g. as a whole) (any can be true)
	Text                  *string             `json:"text,omitempty" bson:"text,omitempty"`                                    // Free text dosage instructions e.g. SIG
	AdditionalInstruction []CodeableConcept   `json:"additionalInstruction,omitempty" bson:"additional_instruction,omitempty"` // Supplemental instruction or warnings to the patient - e.g. "with meals", "may cause drowsiness"
	PatientInstruction    *string             `json:"patientInstruction,omitempty" bson:"patient_instruction,omitempty"`       // Patient or consumer oriented instructions
	Timing                *Timing             `json:"timing,omitempty" bson:"timing,omitempty"`                                // When medication should be administered
	AsNeeded              bool                `json:"asNeeded,omitempty" bson:"as_needed,omitempty"`                           // Take "as needed"
	AsNeededFor           []CodeableConcept   `json:"asNeededFor,omitempty" bson:"as_needed_for,omitempty"`                    // Take "as needed" (for x)
	Site                  *CodeableConcept    `json:"site,omitempty" bson:"site,omitempty"`                                    // Body site to administer to
	Route                 *CodeableConcept    `json:"route,omitempty" bson:"route,omitempty"`                                  // How drug should enter body
	Method                *CodeableConcept    `json:"method,omitempty" bson:"method,omitempty"`                                // Technique for administering medication
	DoseAndRate           []DosageDoseAndRate `json:"doseAndRate,omitempty" bson:"dose_and_rate,omitempty"`                    // Amount of medication administered, to be administered or typical amount to be administered
	Safety                *DosageSafety       `json:"safety,omitempty" bson:"safety,omitempty"`                                // Safety Information about the this dosage instructions
}

func (r *Dosage) Validate() error {
	for i, item := range r.Condition {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Condition[%d]: %w", i, err)
		}
	}
	for i, item := range r.AdditionalInstruction {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AdditionalInstruction[%d]: %w", i, err)
		}
	}
	if r.Timing != nil {
		if err := r.Timing.Validate(); err != nil {
			return fmt.Errorf("Timing: %w", err)
		}
	}
	for i, item := range r.AsNeededFor {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AsNeededFor[%d]: %w", i, err)
		}
	}
	if r.Site != nil {
		if err := r.Site.Validate(); err != nil {
			return fmt.Errorf("Site: %w", err)
		}
	}
	if r.Route != nil {
		if err := r.Route.Validate(); err != nil {
			return fmt.Errorf("Route: %w", err)
		}
	}
	if r.Method != nil {
		if err := r.Method.Validate(); err != nil {
			return fmt.Errorf("Method: %w", err)
		}
	}
	for i, item := range r.DoseAndRate {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("DoseAndRate[%d]: %w", i, err)
		}
	}
	if r.Safety != nil {
		if err := r.Safety.Validate(); err != nil {
			return fmt.Errorf("Safety: %w", err)
		}
	}
	return nil
}

type DosageDoseAndRate struct {
	Id             *string          `json:"id,omitempty" bson:"id,omitempty"`                          // Unique id for inter-element referencing
	Type           *CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`                      // The kind of dose or rate specified
	DoseRange      *Range           `json:"doseRange,omitempty" bson:"dose_range,omitempty"`           // Amount of medication per dose
	DoseQuantity   *Quantity        `json:"doseQuantity,omitempty" bson:"dose_quantity,omitempty"`     // Amount of medication per dose
	DoseExpression *Expression      `json:"doseExpression,omitempty" bson:"dose_expression,omitempty"` // Amount of medication per dose
	RateRatio      *Ratio           `json:"rateRatio,omitempty" bson:"rate_ratio,omitempty"`           // Amount of medication per unit of time
	RateRange      *Range           `json:"rateRange,omitempty" bson:"rate_range,omitempty"`           // Amount of medication per unit of time
	RateQuantity   *Quantity        `json:"rateQuantity,omitempty" bson:"rate_quantity,omitempty"`     // Amount of medication per unit of time
	RateExpression *Expression      `json:"rateExpression,omitempty" bson:"rate_expression,omitempty"` // Amount of medication per unit of time
}

func (r *DosageDoseAndRate) Validate() error {
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.DoseRange != nil {
		if err := r.DoseRange.Validate(); err != nil {
			return fmt.Errorf("DoseRange: %w", err)
		}
	}
	if r.DoseQuantity != nil {
		if err := r.DoseQuantity.Validate(); err != nil {
			return fmt.Errorf("DoseQuantity: %w", err)
		}
	}
	if r.DoseExpression != nil {
		if err := r.DoseExpression.Validate(); err != nil {
			return fmt.Errorf("DoseExpression: %w", err)
		}
	}
	if r.RateRatio != nil {
		if err := r.RateRatio.Validate(); err != nil {
			return fmt.Errorf("RateRatio: %w", err)
		}
	}
	if r.RateRange != nil {
		if err := r.RateRange.Validate(); err != nil {
			return fmt.Errorf("RateRange: %w", err)
		}
	}
	if r.RateQuantity != nil {
		if err := r.RateQuantity.Validate(); err != nil {
			return fmt.Errorf("RateQuantity: %w", err)
		}
	}
	if r.RateExpression != nil {
		if err := r.RateExpression.Validate(); err != nil {
			return fmt.Errorf("RateExpression: %w", err)
		}
	}
	return nil
}
