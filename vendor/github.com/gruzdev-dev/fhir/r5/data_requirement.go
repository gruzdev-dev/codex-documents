package models

import (
	"fmt"
)

// DataRequirement Type: Describes a required data item for evaluation in terms of the type of data, and optional code or date-based filters of the data.
type DataRequirement struct {
	Id                     *string                      `json:"id,omitempty" bson:"id,omitempty"`                                           // Unique id for inter-element referencing
	Type                   string                       `json:"type" bson:"type"`                                                           // The type of the required data
	Profile                []string                     `json:"profile,omitempty" bson:"profile,omitempty"`                                 // The profile of the required data
	SubjectCodeableConcept *CodeableConcept             `json:"subjectCodeableConcept,omitempty" bson:"subject_codeable_concept,omitempty"` // E.g. Patient, Practitioner, RelatedPerson, Organization, Location, Device
	SubjectReference       *Reference                   `json:"subjectReference,omitempty" bson:"subject_reference,omitempty"`              // E.g. Patient, Practitioner, RelatedPerson, Organization, Location, Device
	MustSupport            []string                     `json:"mustSupport,omitempty" bson:"must_support,omitempty"`                        // Indicates specific structure elements that are referenced by the knowledge module
	CodeFilter             []DataRequirementCodeFilter  `json:"codeFilter,omitempty" bson:"code_filter,omitempty"`                          // What codes are expected
	DateFilter             []DataRequirementDateFilter  `json:"dateFilter,omitempty" bson:"date_filter,omitempty"`                          // What dates/date ranges are expected
	ValueFilter            []DataRequirementValueFilter `json:"valueFilter,omitempty" bson:"value_filter,omitempty"`                        // What values are expected
	Limit                  *int                         `json:"limit,omitempty" bson:"limit,omitempty"`                                     // Number of results
	Sort                   []DataRequirementSort        `json:"sort,omitempty" bson:"sort,omitempty"`                                       // Order of the results
}

func (r *DataRequirement) Validate() error {
	var emptyString string
	if r.Type == emptyString {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.SubjectCodeableConcept != nil {
		if err := r.SubjectCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("SubjectCodeableConcept: %w", err)
		}
	}
	if r.SubjectReference != nil {
		if err := r.SubjectReference.Validate(); err != nil {
			return fmt.Errorf("SubjectReference: %w", err)
		}
	}
	for i, item := range r.CodeFilter {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("CodeFilter[%d]: %w", i, err)
		}
	}
	for i, item := range r.DateFilter {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("DateFilter[%d]: %w", i, err)
		}
	}
	for i, item := range r.ValueFilter {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ValueFilter[%d]: %w", i, err)
		}
	}
	for i, item := range r.Sort {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Sort[%d]: %w", i, err)
		}
	}
	return nil
}

type DataRequirementDateFilter struct {
	Id            *string   `json:"id,omitempty" bson:"id,omitempty"`                         // Unique id for inter-element referencing
	Path          *string   `json:"path,omitempty" bson:"path,omitempty"`                     // A date-valued attribute to filter on
	SearchParam   *string   `json:"searchParam,omitempty" bson:"search_param,omitempty"`      // A date valued parameter to search on
	ValueDateTime *string   `json:"valueDateTime,omitempty" bson:"value_date_time,omitempty"` // The value of the filter, as a Period, DateTime, or Duration value
	ValuePeriod   *Period   `json:"valuePeriod,omitempty" bson:"value_period,omitempty"`      // The value of the filter, as a Period, DateTime, or Duration value
	ValueDuration *Duration `json:"valueDuration,omitempty" bson:"value_duration,omitempty"`  // The value of the filter, as a Period, DateTime, or Duration value
}

func (r *DataRequirementDateFilter) Validate() error {
	if r.ValuePeriod != nil {
		if err := r.ValuePeriod.Validate(); err != nil {
			return fmt.Errorf("ValuePeriod: %w", err)
		}
	}
	if r.ValueDuration != nil {
		if err := r.ValueDuration.Validate(); err != nil {
			return fmt.Errorf("ValueDuration: %w", err)
		}
	}
	return nil
}

type DataRequirementValueFilter struct {
	Id            *string   `json:"id,omitempty" bson:"id,omitempty"`                         // Unique id for inter-element referencing
	Path          *string   `json:"path,omitempty" bson:"path,omitempty"`                     // An attribute to filter on
	SearchParam   *string   `json:"searchParam,omitempty" bson:"search_param,omitempty"`      // A parameter to search on
	Comparator    *string   `json:"comparator,omitempty" bson:"comparator,omitempty"`         // eq | gt | lt | ge | le | sa | eb
	ValueDateTime *string   `json:"valueDateTime,omitempty" bson:"value_date_time,omitempty"` // The value of the filter, as a Period, DateTime, or Duration value
	ValuePeriod   *Period   `json:"valuePeriod,omitempty" bson:"value_period,omitempty"`      // The value of the filter, as a Period, DateTime, or Duration value
	ValueDuration *Duration `json:"valueDuration,omitempty" bson:"value_duration,omitempty"`  // The value of the filter, as a Period, DateTime, or Duration value
}

func (r *DataRequirementValueFilter) Validate() error {
	if r.ValuePeriod != nil {
		if err := r.ValuePeriod.Validate(); err != nil {
			return fmt.Errorf("ValuePeriod: %w", err)
		}
	}
	if r.ValueDuration != nil {
		if err := r.ValueDuration.Validate(); err != nil {
			return fmt.Errorf("ValueDuration: %w", err)
		}
	}
	return nil
}

type DataRequirementSort struct {
	Id        *string `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Path      string  `json:"path" bson:"path"`                 // The name of the attribute to perform the sort
	Direction string  `json:"direction" bson:"direction"`       // ascending | descending
}

func (r *DataRequirementSort) Validate() error {
	var emptyString string
	if r.Path == emptyString {
		return fmt.Errorf("field 'Path' is required")
	}
	if r.Direction == emptyString {
		return fmt.Errorf("field 'Direction' is required")
	}
	return nil
}

type DataRequirementCodeFilter struct {
	Id          *string  `json:"id,omitempty" bson:"id,omitempty"`                    // Unique id for inter-element referencing
	Path        *string  `json:"path,omitempty" bson:"path,omitempty"`                // A code-valued attribute to filter on
	SearchParam *string  `json:"searchParam,omitempty" bson:"search_param,omitempty"` // A coded (token) parameter to search on
	ValueSet    *string  `json:"valueSet,omitempty" bson:"value_set,omitempty"`       // ValueSet for the filter
	Code        []Coding `json:"code,omitempty" bson:"code,omitempty"`                // What code is expected
}

func (r *DataRequirementCodeFilter) Validate() error {
	for i, item := range r.Code {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Code[%d]: %w", i, err)
		}
	}
	return nil
}
