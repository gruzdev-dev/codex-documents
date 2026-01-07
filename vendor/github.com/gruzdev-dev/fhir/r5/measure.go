package models

import (
	"encoding/json"
	"fmt"
)

// The Measure resource provides the definition of a quality measure.
type Measure struct {
	ResourceType                    string                    `json:"resourceType" bson:"resource_type"`                                                            // Type of resource
	Id                              *string                   `json:"id,omitempty" bson:"id,omitempty"`                                                             // Logical id of this artifact
	Meta                            *Meta                     `json:"meta,omitempty" bson:"meta,omitempty"`                                                         // Metadata about the resource
	ImplicitRules                   *string                   `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                                      // A set of rules under which this content was created
	Language                        *string                   `json:"language,omitempty" bson:"language,omitempty"`                                                 // Language of the resource content
	Text                            *Narrative                `json:"text,omitempty" bson:"text,omitempty"`                                                         // Text summary of the resource, for human interpretation
	Contained                       []json.RawMessage         `json:"contained,omitempty" bson:"contained,omitempty"`                                               // Contained, inline Resources
	Url                             *string                   `json:"url,omitempty" bson:"url,omitempty"`                                                           // Canonical identifier for this measure, represented as a URI (globally unique)
	Identifier                      []Identifier              `json:"identifier,omitempty" bson:"identifier,omitempty"`                                             // Additional identifier for the measure
	Version                         *string                   `json:"version,omitempty" bson:"version,omitempty"`                                                   // Business version of the measure
	VersionAlgorithmString          *string                   `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"`                   // How to compare versions
	VersionAlgorithmCoding          *Coding                   `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"`                   // How to compare versions
	Name                            *string                   `json:"name,omitempty" bson:"name,omitempty"`                                                         // Name for this measure (computer friendly)
	Title                           *string                   `json:"title,omitempty" bson:"title,omitempty"`                                                       // Name for this measure (human friendly)
	Subtitle                        *string                   `json:"subtitle,omitempty" bson:"subtitle,omitempty"`                                                 // Subordinate title of the measure
	Status                          string                    `json:"status" bson:"status"`                                                                         // draft | active | retired | unknown
	Experimental                    bool                      `json:"experimental,omitempty" bson:"experimental,omitempty"`                                         // For testing only - never for real usage
	SubjectCodeableConcept          *CodeableConcept          `json:"subjectCodeableConcept,omitempty" bson:"subject_codeable_concept,omitempty"`                   // E.g. Patient, Practitioner, RelatedPerson, Organization, Location, Device
	SubjectReference                *Reference                `json:"subjectReference,omitempty" bson:"subject_reference,omitempty"`                                // E.g. Patient, Practitioner, RelatedPerson, Organization, Location, Device
	Date                            *string                   `json:"date,omitempty" bson:"date,omitempty"`                                                         // Date last changed
	Publisher                       *string                   `json:"publisher,omitempty" bson:"publisher,omitempty"`                                               // Name of the publisher/steward (organization or individual)
	Contact                         []ContactDetail           `json:"contact,omitempty" bson:"contact,omitempty"`                                                   // Contact details for the publisher
	Description                     *string                   `json:"description,omitempty" bson:"description,omitempty"`                                           // Natural language description of the measure
	UseContext                      []UsageContext            `json:"useContext,omitempty" bson:"use_context,omitempty"`                                            // The context that the content is intended to support
	Jurisdiction                    []CodeableConcept         `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                                         // Jurisdiction of the authority that maintains the measure (if applicable)
	Purpose                         *string                   `json:"purpose,omitempty" bson:"purpose,omitempty"`                                                   // Why this measure is defined
	Usage                           *string                   `json:"usage,omitempty" bson:"usage,omitempty"`                                                       // Describes the clinical usage of the measure
	Copyright                       *string                   `json:"copyright,omitempty" bson:"copyright,omitempty"`                                               // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel                  *string                   `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                                    // Copyright holder and year(s)
	ApprovalDate                    *string                   `json:"approvalDate,omitempty" bson:"approval_date,omitempty"`                                        // When the measure was approved by publisher
	LastReviewDate                  *string                   `json:"lastReviewDate,omitempty" bson:"last_review_date,omitempty"`                                   // When the measure was last reviewed by the publisher
	EffectivePeriod                 *Period                   `json:"effectivePeriod,omitempty" bson:"effective_period,omitempty"`                                  // When the measure is expected to be used
	ReportingFrequency              *Quantity                 `json:"reportingFrequency,omitempty" bson:"reporting_frequency,omitempty"`                            // The frequency in which this measure should be reported (e.g. 1 '/a' - yearly, 4 '/a' - quarterly)
	Topic                           []CodeableConcept         `json:"topic,omitempty" bson:"topic,omitempty"`                                                       // The category of the measure, such as Education, Treatment, Assessment, etc
	Author                          []ContactDetail           `json:"author,omitempty" bson:"author,omitempty"`                                                     // Who authored the content
	Editor                          []ContactDetail           `json:"editor,omitempty" bson:"editor,omitempty"`                                                     // Who edited the content
	Reviewer                        []ContactDetail           `json:"reviewer,omitempty" bson:"reviewer,omitempty"`                                                 // Who reviewed the content
	Endorser                        []ContactDetail           `json:"endorser,omitempty" bson:"endorser,omitempty"`                                                 // Who endorsed the content
	RelatedArtifact                 []RelatedArtifact         `json:"relatedArtifact,omitempty" bson:"related_artifact,omitempty"`                                  // Additional documentation, citations, etc
	Library                         []string                  `json:"library,omitempty" bson:"library,omitempty"`                                                   // Logic used by the measure
	Disclaimer                      *string                   `json:"disclaimer,omitempty" bson:"disclaimer,omitempty"`                                             // Disclaimer for use of the measure or its referenced content
	RiskAdjustment                  *string                   `json:"riskAdjustment,omitempty" bson:"risk_adjustment,omitempty"`                                    // How risk adjustment is applied for this measure
	RateAggregation                 *string                   `json:"rateAggregation,omitempty" bson:"rate_aggregation,omitempty"`                                  // How is rate aggregation performed for this measure
	Rationale                       *string                   `json:"rationale,omitempty" bson:"rationale,omitempty"`                                               // Justification for the measure in terms of impact, gap in care, and evidence
	ClinicalRecommendationStatement *string                   `json:"clinicalRecommendationStatement,omitempty" bson:"clinical_recommendation_statement,omitempty"` // Summary of clinical guidelines
	Term                            []MeasureTerm             `json:"term,omitempty" bson:"term,omitempty"`                                                         // Defined terms used in the measure documentation
	Guidance                        *string                   `json:"guidance,omitempty" bson:"guidance,omitempty"`                                                 // Additional guidance for implementers (deprecated)
	Group                           []MeasureGroup            `json:"group,omitempty" bson:"group,omitempty"`                                                       // Population criteria group
	SupplementalData                []MeasureSupplementalData `json:"supplementalData,omitempty" bson:"supplemental_data,omitempty"`                                // What other data should be reported with the measure
}

func (r *Measure) Validate() error {
	if r.ResourceType != "Measure" {
		return fmt.Errorf("invalid resourceType: expected 'Measure', got '%s'", r.ResourceType)
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
	if r.VersionAlgorithmCoding != nil {
		if err := r.VersionAlgorithmCoding.Validate(); err != nil {
			return fmt.Errorf("VersionAlgorithmCoding: %w", err)
		}
	}
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
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
	for i, item := range r.Contact {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Contact[%d]: %w", i, err)
		}
	}
	for i, item := range r.UseContext {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("UseContext[%d]: %w", i, err)
		}
	}
	for i, item := range r.Jurisdiction {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Jurisdiction[%d]: %w", i, err)
		}
	}
	if r.EffectivePeriod != nil {
		if err := r.EffectivePeriod.Validate(); err != nil {
			return fmt.Errorf("EffectivePeriod: %w", err)
		}
	}
	if r.ReportingFrequency != nil {
		if err := r.ReportingFrequency.Validate(); err != nil {
			return fmt.Errorf("ReportingFrequency: %w", err)
		}
	}
	for i, item := range r.Topic {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Topic[%d]: %w", i, err)
		}
	}
	for i, item := range r.Author {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Author[%d]: %w", i, err)
		}
	}
	for i, item := range r.Editor {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Editor[%d]: %w", i, err)
		}
	}
	for i, item := range r.Reviewer {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reviewer[%d]: %w", i, err)
		}
	}
	for i, item := range r.Endorser {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Endorser[%d]: %w", i, err)
		}
	}
	for i, item := range r.RelatedArtifact {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RelatedArtifact[%d]: %w", i, err)
		}
	}
	for i, item := range r.Term {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Term[%d]: %w", i, err)
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
	return nil
}

type MeasureGroup struct {
	Id                          *string                  `json:"id,omitempty" bson:"id,omitempty"`                                                     // Unique id for inter-element referencing
	LinkId                      *string                  `json:"linkId,omitempty" bson:"link_id,omitempty"`                                            // Unique id for group in measure
	Title                       *string                  `json:"title,omitempty" bson:"title,omitempty"`                                               // Title of the group. This title is expected in the corresponding MeasureReport.group.title
	Code                        *CodeableConcept         `json:"code,omitempty" bson:"code,omitempty"`                                                 // Meaning of the group
	Description                 *string                  `json:"description,omitempty" bson:"description,omitempty"`                                   // Summary description
	Type                        []CodeableConcept        `json:"type,omitempty" bson:"type,omitempty"`                                                 // process | outcome | structure | patient-reported-outcome
	SubjectCodeableConcept      *CodeableConcept         `json:"subjectCodeableConcept,omitempty" bson:"subject_codeable_concept,omitempty"`           // E.g. Patient, Practitioner, RelatedPerson, Organization, Location, Device
	SubjectReference            *Reference               `json:"subjectReference,omitempty" bson:"subject_reference,omitempty"`                        // E.g. Patient, Practitioner, RelatedPerson, Organization, Location, Device
	Basis                       *string                  `json:"basis,omitempty" bson:"basis,omitempty"`                                               // Population basis
	BasisRequirement            []DataRequirement        `json:"basisRequirement,omitempty" bson:"basis_requirement,omitempty"`                        // Population basis requirements
	Scoring                     *CodeableConcept         `json:"scoring,omitempty" bson:"scoring,omitempty"`                                           // proportion | ratio | continuous-variable | cohort | composite
	ScoringUnit                 *CodeableConcept         `json:"scoringUnit,omitempty" bson:"scoring_unit,omitempty"`                                  // What units?
	ScoringPrecision            *int                     `json:"scoringPrecision,omitempty" bson:"scoring_precision,omitempty"`                        // How many decimals (The number of decimal places to include in the score when the score is a decimal-valued result)
	CompositeScoring            *CodeableConcept         `json:"compositeScoring,omitempty" bson:"composite_scoring,omitempty"`                        // opportunity | all-or-nothing | linear | weighted
	Component                   []MeasureGroupComponent  `json:"component,omitempty" bson:"component,omitempty"`                                       // A component of a composite measure
	RateAggregation             *string                  `json:"rateAggregation,omitempty" bson:"rate_aggregation,omitempty"`                          // How is rate aggregation performed for this measure
	ImprovementNotation         *CodeableConcept         `json:"improvementNotation,omitempty" bson:"improvement_notation,omitempty"`                  // increase | decrease
	ImprovementNotationGuidance *string                  `json:"improvementNotationGuidance,omitempty" bson:"improvement_notation_guidance,omitempty"` // Explanation of improvement notation
	Library                     []string                 `json:"library,omitempty" bson:"library,omitempty"`                                           // Logic used by the measure group
	Population                  []MeasureGroupPopulation `json:"population,omitempty" bson:"population,omitempty"`                                     // Population criteria
	Stratifier                  []MeasureGroupStratifier `json:"stratifier,omitempty" bson:"stratifier,omitempty"`                                     // Stratifier criteria for the measure
}

func (r *MeasureGroup) Validate() error {
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	for i, item := range r.Type {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Type[%d]: %w", i, err)
		}
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
	for i, item := range r.BasisRequirement {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("BasisRequirement[%d]: %w", i, err)
		}
	}
	if r.Scoring != nil {
		if err := r.Scoring.Validate(); err != nil {
			return fmt.Errorf("Scoring: %w", err)
		}
	}
	if r.ScoringUnit != nil {
		if err := r.ScoringUnit.Validate(); err != nil {
			return fmt.Errorf("ScoringUnit: %w", err)
		}
	}
	if r.CompositeScoring != nil {
		if err := r.CompositeScoring.Validate(); err != nil {
			return fmt.Errorf("CompositeScoring: %w", err)
		}
	}
	for i, item := range r.Component {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Component[%d]: %w", i, err)
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
	for i, item := range r.Stratifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Stratifier[%d]: %w", i, err)
		}
	}
	return nil
}

type MeasureGroupComponent struct {
	Id      *string  `json:"id,omitempty" bson:"id,omitempty"`            // Unique id for inter-element referencing
	Measure *string  `json:"measure,omitempty" bson:"measure,omitempty"`  // What measure?
	GroupId *string  `json:"groupId,omitempty" bson:"group_id,omitempty"` // What group?
	Weight  *float64 `json:"weight,omitempty" bson:"weight,omitempty"`    // What weight?
}

func (r *MeasureGroupComponent) Validate() error {
	return nil
}

type MeasureGroupStratifier struct {
	Id              *string                           `json:"id,omitempty" bson:"id,omitempty"`                            // Unique id for inter-element referencing
	LinkId          *string                           `json:"linkId,omitempty" bson:"link_id,omitempty"`                   // Unique id for stratifier in measure
	Title           *string                           `json:"title,omitempty" bson:"title,omitempty"`                      // Title of a group's stratifier. This title is expected in the corresponding MeasureReport.group.title
	Code            *CodeableConcept                  `json:"code,omitempty" bson:"code,omitempty"`                        // Meaning of the stratifier
	Description     *string                           `json:"description,omitempty" bson:"description,omitempty"`          // The human readable description of this stratifier
	Criteria        *Expression                       `json:"criteria,omitempty" bson:"criteria,omitempty"`                // How the measure should be stratified
	GroupDefinition *Reference                        `json:"groupDefinition,omitempty" bson:"group_definition,omitempty"` // A group resource that defines this population
	Component       []MeasureGroupStratifierComponent `json:"component,omitempty" bson:"component,omitempty"`              // Stratifier criteria component for the measure
}

func (r *MeasureGroupStratifier) Validate() error {
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.Criteria != nil {
		if err := r.Criteria.Validate(); err != nil {
			return fmt.Errorf("Criteria: %w", err)
		}
	}
	if r.GroupDefinition != nil {
		if err := r.GroupDefinition.Validate(); err != nil {
			return fmt.Errorf("GroupDefinition: %w", err)
		}
	}
	for i, item := range r.Component {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Component[%d]: %w", i, err)
		}
	}
	return nil
}

type MeasureTerm struct {
	Id         *string          `json:"id,omitempty" bson:"id,omitempty"`                 // Unique id for inter-element referencing
	Code       *CodeableConcept `json:"code,omitempty" bson:"code,omitempty"`             // What term?
	Definition *string          `json:"definition,omitempty" bson:"definition,omitempty"` // Meaning of the term
}

func (r *MeasureTerm) Validate() error {
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	return nil
}

type MeasureGroupPopulation struct {
	Id                *string          `json:"id,omitempty" bson:"id,omitempty"`                                 // Unique id for inter-element referencing
	LinkId            *string          `json:"linkId,omitempty" bson:"link_id,omitempty"`                        // Unique id for population in measure
	Title             *string          `json:"title,omitempty" bson:"title,omitempty"`                           // Title of the group's population. This title is expected in the corresponding MeasureReport.group.population.title
	Code              *CodeableConcept `json:"code,omitempty" bson:"code,omitempty"`                             // initial-population | numerator | numerator-exclusion | denominator | denominator-exclusion | denominator-exception | measure-population | measure-population-exclusion | measure-observation
	Description       *string          `json:"description,omitempty" bson:"description,omitempty"`               // The human readable description of this population criteria
	Criteria          *Expression      `json:"criteria,omitempty" bson:"criteria,omitempty"`                     // The criteria that defines this population
	GroupDefinition   *Reference       `json:"groupDefinition,omitempty" bson:"group_definition,omitempty"`      // A group resource that defines this population
	InputPopulationId *string          `json:"inputPopulationId,omitempty" bson:"input_population_id,omitempty"` // Which population
	AggregateMethod   *CodeableConcept `json:"aggregateMethod,omitempty" bson:"aggregate_method,omitempty"`      // Aggregation method for a measure score (e.g. sum, average, median, minimum, maximum, count)
}

func (r *MeasureGroupPopulation) Validate() error {
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.Criteria != nil {
		if err := r.Criteria.Validate(); err != nil {
			return fmt.Errorf("Criteria: %w", err)
		}
	}
	if r.GroupDefinition != nil {
		if err := r.GroupDefinition.Validate(); err != nil {
			return fmt.Errorf("GroupDefinition: %w", err)
		}
	}
	if r.AggregateMethod != nil {
		if err := r.AggregateMethod.Validate(); err != nil {
			return fmt.Errorf("AggregateMethod: %w", err)
		}
	}
	return nil
}

type MeasureGroupStratifierComponent struct {
	Id              *string          `json:"id,omitempty" bson:"id,omitempty"`                            // Unique id for inter-element referencing
	LinkId          *string          `json:"linkId,omitempty" bson:"link_id,omitempty"`                   // Unique id for stratifier component in measure
	Code            *CodeableConcept `json:"code,omitempty" bson:"code,omitempty"`                        // Meaning of the stratifier component
	Description     *string          `json:"description,omitempty" bson:"description,omitempty"`          // The human readable description of this stratifier component
	Criteria        *Expression      `json:"criteria,omitempty" bson:"criteria,omitempty"`                // Component of how the measure should be stratified
	GroupDefinition *Reference       `json:"groupDefinition,omitempty" bson:"group_definition,omitempty"` // A group resource that defines this population
	ValueSet        *string          `json:"valueSet,omitempty" bson:"value_set,omitempty"`               // What stratum values?
	Unit            *string          `json:"unit,omitempty" bson:"unit,omitempty"`                        // What units?
}

func (r *MeasureGroupStratifierComponent) Validate() error {
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.Criteria != nil {
		if err := r.Criteria.Validate(); err != nil {
			return fmt.Errorf("Criteria: %w", err)
		}
	}
	if r.GroupDefinition != nil {
		if err := r.GroupDefinition.Validate(); err != nil {
			return fmt.Errorf("GroupDefinition: %w", err)
		}
	}
	return nil
}

type MeasureSupplementalData struct {
	Id          *string           `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	LinkId      *string           `json:"linkId,omitempty" bson:"link_id,omitempty"`          // Unique id for supplementalData in measure
	Code        *CodeableConcept  `json:"code,omitempty" bson:"code,omitempty"`               // Meaning of the supplemental data
	Usage       []CodeableConcept `json:"usage,omitempty" bson:"usage,omitempty"`             // supplemental-data | risk-adjustment-factor
	Description *string           `json:"description,omitempty" bson:"description,omitempty"` // The human readable description of this supplemental data
	Criteria    *Expression       `json:"criteria" bson:"criteria"`                           // Expression describing additional data to be reported
	ValueSet    *string           `json:"valueSet,omitempty" bson:"value_set,omitempty"`      // What supplemental data values?
	Unit        *string           `json:"unit,omitempty" bson:"unit,omitempty"`               // What units?
}

func (r *MeasureSupplementalData) Validate() error {
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	for i, item := range r.Usage {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Usage[%d]: %w", i, err)
		}
	}
	if r.Criteria == nil {
		return fmt.Errorf("field 'Criteria' is required")
	}
	if r.Criteria != nil {
		if err := r.Criteria.Validate(); err != nil {
			return fmt.Errorf("Criteria: %w", err)
		}
	}
	return nil
}
