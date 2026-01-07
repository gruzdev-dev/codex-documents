package models

import (
	"encoding/json"
	"fmt"
)

// The MeasureReport resource contains the results of the calculation of a measure; and optionally a reference to the resources involved in that calculation.
type MeasureReport struct {
	ResourceType      string               `json:"resourceType" bson:"resource_type"`                               // Type of resource
	Id                *string              `json:"id,omitempty" bson:"id,omitempty"`                                // Logical id of this artifact
	Meta              *Meta                `json:"meta,omitempty" bson:"meta,omitempty"`                            // Metadata about the resource
	ImplicitRules     *string              `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`         // A set of rules under which this content was created
	Language          *string              `json:"language,omitempty" bson:"language,omitempty"`                    // Language of the resource content
	Text              *Narrative           `json:"text,omitempty" bson:"text,omitempty"`                            // Text summary of the resource, for human interpretation
	Contained         []json.RawMessage    `json:"contained,omitempty" bson:"contained,omitempty"`                  // Contained, inline Resources
	Identifier        []Identifier         `json:"identifier,omitempty" bson:"identifier,omitempty"`                // Additional identifier for the MeasureReport
	Category          *CodeableConcept     `json:"category,omitempty" bson:"category,omitempty"`                    // The category of measure report instance this is (example codes include deqm, ra, vbp)
	Messages          *Reference           `json:"messages,omitempty" bson:"messages,omitempty"`                    // Evaluation messages
	Status            string               `json:"status" bson:"status"`                                            // complete | pending | error
	Type              string               `json:"type" bson:"type"`                                                // individual | subject-list | summary | data-exchange
	DataUpdateType    *string              `json:"dataUpdateType,omitempty" bson:"data_update_type,omitempty"`      // incremental | snapshot
	Measure           *string              `json:"measure,omitempty" bson:"measure,omitempty"`                      // What measure was calculated
	Subject           *Reference           `json:"subject,omitempty" bson:"subject,omitempty"`                      // What individual(s) the report is for
	Date              *string              `json:"date,omitempty" bson:"date,omitempty"`                            // When the measure report was generated
	Reporter          *Reference           `json:"reporter,omitempty" bson:"reporter,omitempty"`                    // Who is reporting the data
	ReportingVendor   *Reference           `json:"reportingVendor,omitempty" bson:"reporting_vendor,omitempty"`     // What vendor prepared the data
	Location          []Reference          `json:"location,omitempty" bson:"location,omitempty"`                    // Where the reported data is from
	Period            *Period              `json:"period" bson:"period"`                                            // What period the report covers
	InputParameters   *Reference           `json:"inputParameters,omitempty" bson:"input_parameters,omitempty"`     // What parameters were provided to the report
	Group             []MeasureReportGroup `json:"group,omitempty" bson:"group,omitempty"`                          // Measure results for each group
	SupplementalData  []Reference          `json:"supplementalData,omitempty" bson:"supplemental_data,omitempty"`   // Additional information collected for the report
	EvaluatedResource []Reference          `json:"evaluatedResource,omitempty" bson:"evaluated_resource,omitempty"` // What data was used to calculate the measure score
}

func (r *MeasureReport) Validate() error {
	if r.ResourceType != "MeasureReport" {
		return fmt.Errorf("invalid resourceType: expected 'MeasureReport', got '%s'", r.ResourceType)
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
	if r.Category != nil {
		if err := r.Category.Validate(); err != nil {
			return fmt.Errorf("Category: %w", err)
		}
	}
	if r.Messages != nil {
		if err := r.Messages.Validate(); err != nil {
			return fmt.Errorf("Messages: %w", err)
		}
	}
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.Type == emptyString {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Subject != nil {
		if err := r.Subject.Validate(); err != nil {
			return fmt.Errorf("Subject: %w", err)
		}
	}
	if r.Reporter != nil {
		if err := r.Reporter.Validate(); err != nil {
			return fmt.Errorf("Reporter: %w", err)
		}
	}
	if r.ReportingVendor != nil {
		if err := r.ReportingVendor.Validate(); err != nil {
			return fmt.Errorf("ReportingVendor: %w", err)
		}
	}
	for i, item := range r.Location {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Location[%d]: %w", i, err)
		}
	}
	if r.Period == nil {
		return fmt.Errorf("field 'Period' is required")
	}
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	if r.InputParameters != nil {
		if err := r.InputParameters.Validate(); err != nil {
			return fmt.Errorf("InputParameters: %w", err)
		}
	}
	for i, item := range r.Group {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Group[%d]: %w", i, err)
		}
	}
	for i, item := range r.SupplementalData {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SupplementalData[%d]: %w", i, err)
		}
	}
	for i, item := range r.EvaluatedResource {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("EvaluatedResource[%d]: %w", i, err)
		}
	}
	return nil
}

type MeasureReportGroup struct {
	Id                          *string                        `json:"id,omitempty" bson:"id,omitempty"`                                                      // Unique id for inter-element referencing
	LinkId                      *string                        `json:"linkId,omitempty" bson:"link_id,omitempty"`                                             // Pointer to specific group from Measure
	Title                       *string                        `json:"title,omitempty" bson:"title,omitempty"`                                                // Title of a group. Note- this value is copied from this element in Measure
	CalculatedDate              *string                        `json:"calculatedDate,omitempty" bson:"calculated_date,omitempty"`                             // The date the Measure Report was calculated
	Code                        *CodeableConcept               `json:"code,omitempty" bson:"code,omitempty"`                                                  // Meaning of the group
	Description                 *string                        `json:"description,omitempty" bson:"description,omitempty"`                                    // Summary description
	Type                        *CodeableConcept               `json:"type,omitempty" bson:"type,omitempty"`                                                  // process | outcome | structure | patient-reported-outcome
	Subject                     *Reference                     `json:"subject,omitempty" bson:"subject,omitempty"`                                            // What individual(s) the report is for
	Scoring                     *CodeableConcept               `json:"scoring,omitempty" bson:"scoring,omitempty"`                                            // What scoring method (e.g. proportion, ratio, continuous-variable)
	ImprovementNotation         *CodeableConcept               `json:"improvementNotation,omitempty" bson:"improvement_notation,omitempty"`                   // increase | decrease
	ImprovementNotationGuidance *string                        `json:"improvementNotationGuidance,omitempty" bson:"improvement_notation_guidance,omitempty"`  // Explanation of improvement notation
	Population                  []MeasureReportGroupPopulation `json:"population,omitempty" bson:"population,omitempty"`                                      // The populations in the group
	MeasureScoreQuantity        *Quantity                      `json:"measureScoreQuantity,omitempty" bson:"measure_score_quantity,omitempty"`                // What score this group achieved
	MeasureScoreDateTime        *string                        `json:"measureScoreDateTime,omitempty" bson:"measure_score_date_time,omitempty"`               // What score this group achieved
	MeasureScoreCodeableConcept *CodeableConcept               `json:"measureScoreCodeableConcept,omitempty" bson:"measure_score_codeable_concept,omitempty"` // What score this group achieved
	MeasureScorePeriod          *Period                        `json:"measureScorePeriod,omitempty" bson:"measure_score_period,omitempty"`                    // What score this group achieved
	MeasureScoreRange           *Range                         `json:"measureScoreRange,omitempty" bson:"measure_score_range,omitempty"`                      // What score this group achieved
	MeasureScoreDuration        *Duration                      `json:"measureScoreDuration,omitempty" bson:"measure_score_duration,omitempty"`                // What score this group achieved
	MeasureScoreBoolean         *bool                          `json:"measureScoreBoolean,omitempty" bson:"measure_score_boolean,omitempty"`                  // What score this group achieved
	MeasureScoreRatio           *Ratio                         `json:"measureScoreRatio,omitempty" bson:"measure_score_ratio,omitempty"`                      // What score this group achieved
	Stratifier                  []MeasureReportGroupStratifier `json:"stratifier,omitempty" bson:"stratifier,omitempty"`                                      // Stratification results
}

func (r *MeasureReportGroup) Validate() error {
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Subject != nil {
		if err := r.Subject.Validate(); err != nil {
			return fmt.Errorf("Subject: %w", err)
		}
	}
	if r.Scoring != nil {
		if err := r.Scoring.Validate(); err != nil {
			return fmt.Errorf("Scoring: %w", err)
		}
	}
	if r.ImprovementNotation != nil {
		if err := r.ImprovementNotation.Validate(); err != nil {
			return fmt.Errorf("ImprovementNotation: %w", err)
		}
	}
	for i, item := range r.Population {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Population[%d]: %w", i, err)
		}
	}
	if r.MeasureScoreQuantity != nil {
		if err := r.MeasureScoreQuantity.Validate(); err != nil {
			return fmt.Errorf("MeasureScoreQuantity: %w", err)
		}
	}
	if r.MeasureScoreCodeableConcept != nil {
		if err := r.MeasureScoreCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("MeasureScoreCodeableConcept: %w", err)
		}
	}
	if r.MeasureScorePeriod != nil {
		if err := r.MeasureScorePeriod.Validate(); err != nil {
			return fmt.Errorf("MeasureScorePeriod: %w", err)
		}
	}
	if r.MeasureScoreRange != nil {
		if err := r.MeasureScoreRange.Validate(); err != nil {
			return fmt.Errorf("MeasureScoreRange: %w", err)
		}
	}
	if r.MeasureScoreDuration != nil {
		if err := r.MeasureScoreDuration.Validate(); err != nil {
			return fmt.Errorf("MeasureScoreDuration: %w", err)
		}
	}
	if r.MeasureScoreRatio != nil {
		if err := r.MeasureScoreRatio.Validate(); err != nil {
			return fmt.Errorf("MeasureScoreRatio: %w", err)
		}
	}
	for i, item := range r.Stratifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Stratifier[%d]: %w", i, err)
		}
	}
	return nil
}

type MeasureReportGroupPopulation struct {
	Id             *string          `json:"id,omitempty" bson:"id,omitempty"`                          // Unique id for inter-element referencing
	LinkId         *string          `json:"linkId,omitempty" bson:"link_id,omitempty"`                 // Pointer to specific population from Measure
	Title          *string          `json:"title,omitempty" bson:"title,omitempty"`                    // Title of a group. Note- this value is copied from this element in Measure
	Code           *CodeableConcept `json:"code,omitempty" bson:"code,omitempty"`                      // initial-population | numerator | numerator-exclusion | denominator | denominator-exclusion | denominator-exception | measure-population | measure-population-exclusion | measure-observation
	Description    *string          `json:"description,omitempty" bson:"description,omitempty"`        // The human readable description of this population criteria
	Count          *int             `json:"count,omitempty" bson:"count,omitempty"`                    // Size of the population
	CountQuantity  *Quantity        `json:"countQuantity,omitempty" bson:"count_quantity,omitempty"`   // Size of the population as a quantity
	SubjectResults *Reference       `json:"subjectResults,omitempty" bson:"subject_results,omitempty"` // For subject-list reports, the subject results in this population
	SubjectReport  []Reference      `json:"subjectReport,omitempty" bson:"subject_report,omitempty"`   // For subject-list reports, references to the individual reports for subjects in this population
	Subjects       *Reference       `json:"subjects,omitempty" bson:"subjects,omitempty"`              // What individual(s) in the population
}

func (r *MeasureReportGroupPopulation) Validate() error {
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.CountQuantity != nil {
		if err := r.CountQuantity.Validate(); err != nil {
			return fmt.Errorf("CountQuantity: %w", err)
		}
	}
	if r.SubjectResults != nil {
		if err := r.SubjectResults.Validate(); err != nil {
			return fmt.Errorf("SubjectResults: %w", err)
		}
	}
	for i, item := range r.SubjectReport {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SubjectReport[%d]: %w", i, err)
		}
	}
	if r.Subjects != nil {
		if err := r.Subjects.Validate(); err != nil {
			return fmt.Errorf("Subjects: %w", err)
		}
	}
	return nil
}

type MeasureReportGroupStratifier struct {
	Id          *string                               `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	LinkId      *string                               `json:"linkId,omitempty" bson:"link_id,omitempty"`          // Pointer to specific stratifier from Measure
	Title       *string                               `json:"title,omitempty" bson:"title,omitempty"`             // Title of a group's stratifier. Note- this value is copied from this element in Measure
	Code        *CodeableConcept                      `json:"code,omitempty" bson:"code,omitempty"`               // What stratifier of the group
	Description *string                               `json:"description,omitempty" bson:"description,omitempty"` // The human readable description of this stratifier
	Stratum     []MeasureReportGroupStratifierStratum `json:"stratum,omitempty" bson:"stratum,omitempty"`         // Stratum results, one for each unique value, or set of values, in the stratifier, or stratifier components
}

func (r *MeasureReportGroupStratifier) Validate() error {
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	for i, item := range r.Stratum {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Stratum[%d]: %w", i, err)
		}
	}
	return nil
}

type MeasureReportGroupStratifierStratum struct {
	Id                          *string                                         `json:"id,omitempty" bson:"id,omitempty"`                                                      // Unique id for inter-element referencing
	ValueCodeableConcept        *CodeableConcept                                `json:"valueCodeableConcept,omitempty" bson:"value_codeable_concept,omitempty"`                // The stratum value, e.g. male
	ValueBoolean                *bool                                           `json:"valueBoolean,omitempty" bson:"value_boolean,omitempty"`                                 // The stratum value, e.g. male
	ValueQuantity               *Quantity                                       `json:"valueQuantity,omitempty" bson:"value_quantity,omitempty"`                               // The stratum value, e.g. male
	ValueRange                  *Range                                          `json:"valueRange,omitempty" bson:"value_range,omitempty"`                                     // The stratum value, e.g. male
	ValueReference              *Reference                                      `json:"valueReference,omitempty" bson:"value_reference,omitempty"`                             // The stratum value, e.g. male
	Component                   []MeasureReportGroupStratifierStratumComponent  `json:"component,omitempty" bson:"component,omitempty"`                                        // Stratifier component values
	Population                  []MeasureReportGroupStratifierStratumPopulation `json:"population,omitempty" bson:"population,omitempty"`                                      // Population results in this stratum
	MeasureScoreQuantity        *Quantity                                       `json:"measureScoreQuantity,omitempty" bson:"measure_score_quantity,omitempty"`                // What score this stratum achieved
	MeasureScoreDateTime        *string                                         `json:"measureScoreDateTime,omitempty" bson:"measure_score_date_time,omitempty"`               // What score this stratum achieved
	MeasureScoreCodeableConcept *CodeableConcept                                `json:"measureScoreCodeableConcept,omitempty" bson:"measure_score_codeable_concept,omitempty"` // What score this stratum achieved
	MeasureScorePeriod          *Period                                         `json:"measureScorePeriod,omitempty" bson:"measure_score_period,omitempty"`                    // What score this stratum achieved
	MeasureScoreRange           *Range                                          `json:"measureScoreRange,omitempty" bson:"measure_score_range,omitempty"`                      // What score this stratum achieved
	MeasureScoreDuration        *Duration                                       `json:"measureScoreDuration,omitempty" bson:"measure_score_duration,omitempty"`                // What score this stratum achieved
	MeasureScoreBoolean         *bool                                           `json:"measureScoreBoolean,omitempty" bson:"measure_score_boolean,omitempty"`                  // What score this stratum achieved
	MeasureScoreRatio           *Ratio                                          `json:"measureScoreRatio,omitempty" bson:"measure_score_ratio,omitempty"`                      // What score this stratum achieved
}

func (r *MeasureReportGroupStratifierStratum) Validate() error {
	if r.ValueCodeableConcept != nil {
		if err := r.ValueCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("ValueCodeableConcept: %w", err)
		}
	}
	if r.ValueQuantity != nil {
		if err := r.ValueQuantity.Validate(); err != nil {
			return fmt.Errorf("ValueQuantity: %w", err)
		}
	}
	if r.ValueRange != nil {
		if err := r.ValueRange.Validate(); err != nil {
			return fmt.Errorf("ValueRange: %w", err)
		}
	}
	if r.ValueReference != nil {
		if err := r.ValueReference.Validate(); err != nil {
			return fmt.Errorf("ValueReference: %w", err)
		}
	}
	for i, item := range r.Component {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Component[%d]: %w", i, err)
		}
	}
	for i, item := range r.Population {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Population[%d]: %w", i, err)
		}
	}
	if r.MeasureScoreQuantity != nil {
		if err := r.MeasureScoreQuantity.Validate(); err != nil {
			return fmt.Errorf("MeasureScoreQuantity: %w", err)
		}
	}
	if r.MeasureScoreCodeableConcept != nil {
		if err := r.MeasureScoreCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("MeasureScoreCodeableConcept: %w", err)
		}
	}
	if r.MeasureScorePeriod != nil {
		if err := r.MeasureScorePeriod.Validate(); err != nil {
			return fmt.Errorf("MeasureScorePeriod: %w", err)
		}
	}
	if r.MeasureScoreRange != nil {
		if err := r.MeasureScoreRange.Validate(); err != nil {
			return fmt.Errorf("MeasureScoreRange: %w", err)
		}
	}
	if r.MeasureScoreDuration != nil {
		if err := r.MeasureScoreDuration.Validate(); err != nil {
			return fmt.Errorf("MeasureScoreDuration: %w", err)
		}
	}
	if r.MeasureScoreRatio != nil {
		if err := r.MeasureScoreRatio.Validate(); err != nil {
			return fmt.Errorf("MeasureScoreRatio: %w", err)
		}
	}
	return nil
}

type MeasureReportGroupStratifierStratumComponent struct {
	Id                   *string          `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	LinkId               *string          `json:"linkId,omitempty" bson:"link_id,omitempty"`          // Pointer to specific stratifier component from Measure
	Code                 *CodeableConcept `json:"code" bson:"code"`                                   // What stratifier component of the group
	Description          *string          `json:"description,omitempty" bson:"description,omitempty"` // The human readable description of this stratifier component
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept" bson:"value_codeable_concept"` // The stratum component value, e.g. male
	ValueBoolean         *bool            `json:"valueBoolean" bson:"value_boolean"`                  // The stratum component value, e.g. male
	ValueQuantity        *Quantity        `json:"valueQuantity" bson:"value_quantity"`                // The stratum component value, e.g. male
	ValueRange           *Range           `json:"valueRange" bson:"value_range"`                      // The stratum component value, e.g. male
	ValueReference       *Reference       `json:"valueReference" bson:"value_reference"`              // The stratum component value, e.g. male
}

func (r *MeasureReportGroupStratifierStratumComponent) Validate() error {
	if r.Code == nil {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
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
	if r.ValueBoolean == nil {
		return fmt.Errorf("field 'ValueBoolean' is required")
	}
	if r.ValueQuantity == nil {
		return fmt.Errorf("field 'ValueQuantity' is required")
	}
	if r.ValueQuantity != nil {
		if err := r.ValueQuantity.Validate(); err != nil {
			return fmt.Errorf("ValueQuantity: %w", err)
		}
	}
	if r.ValueRange == nil {
		return fmt.Errorf("field 'ValueRange' is required")
	}
	if r.ValueRange != nil {
		if err := r.ValueRange.Validate(); err != nil {
			return fmt.Errorf("ValueRange: %w", err)
		}
	}
	if r.ValueReference == nil {
		return fmt.Errorf("field 'ValueReference' is required")
	}
	if r.ValueReference != nil {
		if err := r.ValueReference.Validate(); err != nil {
			return fmt.Errorf("ValueReference: %w", err)
		}
	}
	return nil
}

type MeasureReportGroupStratifierStratumPopulation struct {
	Id             *string          `json:"id,omitempty" bson:"id,omitempty"`                          // Unique id for inter-element referencing
	LinkId         *string          `json:"linkId,omitempty" bson:"link_id,omitempty"`                 // Pointer to specific population from Measure
	Code           *CodeableConcept `json:"code,omitempty" bson:"code,omitempty"`                      // initial-population | numerator | numerator-exclusion | denominator | denominator-exclusion | denominator-exception | measure-population | measure-population-exclusion | measure-observation
	Count          *int             `json:"count,omitempty" bson:"count,omitempty"`                    // Size of the population
	CountQuantity  *Quantity        `json:"countQuantity,omitempty" bson:"count_quantity,omitempty"`   // Size of the population as a quantity
	SubjectResults *Reference       `json:"subjectResults,omitempty" bson:"subject_results,omitempty"` // For subject-list reports, the subject results in this population
	SubjectReport  []Reference      `json:"subjectReport,omitempty" bson:"subject_report,omitempty"`   // For subject-list reports, a subject result in this population
	Subjects       *Reference       `json:"subjects,omitempty" bson:"subjects,omitempty"`              // What individual(s) in the population
}

func (r *MeasureReportGroupStratifierStratumPopulation) Validate() error {
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.CountQuantity != nil {
		if err := r.CountQuantity.Validate(); err != nil {
			return fmt.Errorf("CountQuantity: %w", err)
		}
	}
	if r.SubjectResults != nil {
		if err := r.SubjectResults.Validate(); err != nil {
			return fmt.Errorf("SubjectResults: %w", err)
		}
	}
	for i, item := range r.SubjectReport {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SubjectReport[%d]: %w", i, err)
		}
	}
	if r.Subjects != nil {
		if err := r.Subjects.Validate(); err != nil {
			return fmt.Errorf("Subjects: %w", err)
		}
	}
	return nil
}
