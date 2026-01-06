package models

import (
	"fmt"
)

// ProductShelfLife Type: The shelf-life and storage information for a medicinal product item or container can be described using this class.
type ProductShelfLife struct {
	Id                           *string           `json:"id,omitempty" bson:"id,omitempty"`                                                        // Unique id for inter-element referencing
	Type                         *CodeableConcept  `json:"type,omitempty" bson:"type,omitempty"`                                                    // This describes the shelf life, taking into account various scenarios such as shelf life of the packaged Medicinal Product itself, shelf life after transformation where necessary and shelf life after the first opening of a bottle, etc. The shelf life type shall be specified using an appropriate controlled vocabulary The controlled term and the controlled term identifier shall be specified
	PeriodDuration               *Duration         `json:"periodDuration,omitempty" bson:"period_duration,omitempty"`                               // The shelf life time period can be specified using a numerical value for the period of time and its unit of time measurement The unit of measurement shall be specified in accordance with ISO 11240 and the resulting terminology The symbol and the symbol identifier shall be used
	PeriodString                 *string           `json:"periodString,omitempty" bson:"period_string,omitempty"`                                   // The shelf life time period can be specified using a numerical value for the period of time and its unit of time measurement The unit of measurement shall be specified in accordance with ISO 11240 and the resulting terminology The symbol and the symbol identifier shall be used
	SpecialPrecautionsForStorage []CodeableConcept `json:"specialPrecautionsForStorage,omitempty" bson:"special_precautions_for_storage,omitempty"` // Special precautions for storage, if any, can be specified using an appropriate controlled vocabulary The controlled term and the controlled term identifier shall be specified
}

func (r *ProductShelfLife) Validate() error {
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.PeriodDuration != nil {
		if err := r.PeriodDuration.Validate(); err != nil {
			return fmt.Errorf("PeriodDuration: %w", err)
		}
	}
	for i, item := range r.SpecialPrecautionsForStorage {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SpecialPrecautionsForStorage[%d]: %w", i, err)
		}
	}
	return nil
}
