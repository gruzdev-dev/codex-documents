package models

import (
	"fmt"
)

// MarketingStatus Type: The marketing status describes the date when an item is actually put on the market or the date as of which it is no longer available.
type MarketingStatus struct {
	Id           *string          `json:"id,omitempty" bson:"id,omitempty"`                     // Unique id for inter-element referencing
	Country      *CodeableConcept `json:"country,omitempty" bson:"country,omitempty"`           // The country in which the marketing status applies
	Jurisdiction *CodeableConcept `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"` // The jurisdiction in which the marketing status applies
	Status       *CodeableConcept `json:"status" bson:"status"`                                 // This attribute provides information on the status of the marketing of the item
	DateRange    *Period          `json:"dateRange,omitempty" bson:"date_range,omitempty"`      // The dates that the item is made available on the market by the owner (or where applicable, the manufacturer/distributor) in a country and/or jurisdiction. Note that “on the market” refers to the release of the item into the distribution chain
	RestoreDate  *string          `json:"restoreDate,omitempty" bson:"restore_date,omitempty"`  // The date when the item is due to be placed back on the market by the owner, manufacturer or distributor, after a suspension
}

func (r *MarketingStatus) Validate() error {
	if r.Country != nil {
		if err := r.Country.Validate(); err != nil {
			return fmt.Errorf("Country: %w", err)
		}
	}
	if r.Jurisdiction != nil {
		if err := r.Jurisdiction.Validate(); err != nil {
			return fmt.Errorf("Jurisdiction: %w", err)
		}
	}
	if r.Status == nil {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.Status != nil {
		if err := r.Status.Validate(); err != nil {
			return fmt.Errorf("Status: %w", err)
		}
	}
	if r.DateRange != nil {
		if err := r.DateRange.Validate(); err != nil {
			return fmt.Errorf("DateRange: %w", err)
		}
	}
	return nil
}
