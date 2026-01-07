package models

import (
	"encoding/json"
	"fmt"
)

// The Evidence Resource provides a machine-interpretable expression of an evidence concept including the evidence variables (e.g., population, exposures/interventions, comparators, outcomes, measured variables, confounding variables), the statistics, and the certainty of this evidence.
type Evidence struct {
	Id                     *string                      `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta                        `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string                      `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string                      `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative                   `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage            `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Url                    *string                      `json:"url,omitempty" bson:"url,omitempty"`                                         // Canonical identifier for this evidence, represented as a globally unique URI
	Identifier             []Identifier                 `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // Additional identifier for the summary
	Version                *string                      `json:"version,omitempty" bson:"version,omitempty"`                                 // Business version of this summary
	VersionAlgorithmString *string                      `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"` // How to compare versions
	VersionAlgorithmCoding *Coding                      `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"` // How to compare versions
	Name                   *string                      `json:"name,omitempty" bson:"name,omitempty"`                                       // Name for this summary (machine friendly)
	Title                  *string                      `json:"title,omitempty" bson:"title,omitempty"`                                     // Name for this summary (human friendly)
	CiteAs                 *string                      `json:"citeAs,omitempty" bson:"cite_as,omitempty"`                                  // Display of how to cite this Evidence
	Status                 string                       `json:"status" bson:"status"`                                                       // draft | active | retired | unknown
	Experimental           bool                         `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // For testing only - never for real usage
	Date                   *string                      `json:"date,omitempty" bson:"date,omitempty"`                                       // Date last changed
	ApprovalDate           *string                      `json:"approvalDate,omitempty" bson:"approval_date,omitempty"`                      // When the summary was approved by publisher
	LastReviewDate         *string                      `json:"lastReviewDate,omitempty" bson:"last_review_date,omitempty"`                 // When the summary was last reviewed by the publisher
	Author                 []ContactDetail              `json:"author,omitempty" bson:"author,omitempty"`                                   // Who authored the content
	Publisher              *string                      `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // Name of the publisher/steward (organization or individual)
	Contact                []ContactDetail              `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Recorder               []ContactDetail              `json:"recorder,omitempty" bson:"recorder,omitempty"`                               // Who entered the data for the evidence
	Editor                 []ContactDetail              `json:"editor,omitempty" bson:"editor,omitempty"`                                   // Who edited the content
	Reviewer               []ContactDetail              `json:"reviewer,omitempty" bson:"reviewer,omitempty"`                               // Who reviewed the content
	Endorser               []ContactDetail              `json:"endorser,omitempty" bson:"endorser,omitempty"`                               // Who endorsed the content
	UseContext             []UsageContext               `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // The context that the content is intended to support
	Purpose                *string                      `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // Why this Evidence is defined
	Copyright              *string                      `json:"copyright,omitempty" bson:"copyright,omitempty"`                             // Intellectual property ownership, may include restrictions on use
	CopyrightLabel         *string                      `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                  // Copyright holder and year(s)
	RelatesTo              []EvidenceRelatesTo          `json:"relatesTo,omitempty" bson:"relates_to,omitempty"`                            // Relationships to other Resources
	Description            *string                      `json:"description,omitempty" bson:"description,omitempty"`                         // Description of the particular summary
	Assertion              *string                      `json:"assertion,omitempty" bson:"assertion,omitempty"`                             // Declarative description of the Evidence
	Note                   []Annotation                 `json:"note,omitempty" bson:"note,omitempty"`                                       // Footnotes and/or explanatory notes
	VariableDefinition     []EvidenceVariableDefinition `json:"variableDefinition,omitempty" bson:"variable_definition,omitempty"`          // Description, classification, and definition of a single variable
	SynthesisType          []CodeableConcept            `json:"synthesisType,omitempty" bson:"synthesis_type,omitempty"`                    // The design of the synthesis (combination of studies) that produced this evidence
	StudyDesign            []CodeableConcept            `json:"studyDesign,omitempty" bson:"study_design,omitempty"`                        // The design of the study that produced this evidence
	Statistic              []EvidenceStatistic          `json:"statistic,omitempty" bson:"statistic,omitempty"`                             // Values and parameters for a single statistic
	Certainty              []EvidenceCertainty          `json:"certainty,omitempty" bson:"certainty,omitempty"`                             // Certainty or quality of the evidence
}

func (r *Evidence) Validate() error {
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
	for i, item := range r.Author {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Author[%d]: %w", i, err)
		}
	}
	for i, item := range r.Contact {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Contact[%d]: %w", i, err)
		}
	}
	for i, item := range r.Recorder {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Recorder[%d]: %w", i, err)
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
	for i, item := range r.UseContext {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("UseContext[%d]: %w", i, err)
		}
	}
	for i, item := range r.RelatesTo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RelatesTo[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	for i, item := range r.VariableDefinition {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("VariableDefinition[%d]: %w", i, err)
		}
	}
	for i, item := range r.SynthesisType {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SynthesisType[%d]: %w", i, err)
		}
	}
	for i, item := range r.StudyDesign {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("StudyDesign[%d]: %w", i, err)
		}
	}
	for i, item := range r.Statistic {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Statistic[%d]: %w", i, err)
		}
	}
	for i, item := range r.Certainty {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Certainty[%d]: %w", i, err)
		}
	}
	return nil
}

type EvidenceCertainty struct {
	Id           *string             `json:"id,omitempty" bson:"id,omitempty"`                     // Unique id for inter-element referencing
	Description  *string             `json:"description,omitempty" bson:"description,omitempty"`   // Textual description of certainty
	Note         []Annotation        `json:"note,omitempty" bson:"note,omitempty"`                 // Footnotes and/or explanatory notes
	Type         *CodeableConcept    `json:"type,omitempty" bson:"type,omitempty"`                 // Aspect of certainty being rated
	Rating       *CodeableConcept    `json:"rating,omitempty" bson:"rating,omitempty"`             // Assessment or judgement of the aspect
	Rater        []string            `json:"rater,omitempty" bson:"rater,omitempty"`               // Individual or group who did the rating
	Subcomponent []EvidenceCertainty `json:"subcomponent,omitempty" bson:"subcomponent,omitempty"` // A domain or subdomain of certainty
}

func (r *EvidenceCertainty) Validate() error {
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Rating != nil {
		if err := r.Rating.Validate(); err != nil {
			return fmt.Errorf("Rating: %w", err)
		}
	}
	for i, item := range r.Subcomponent {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Subcomponent[%d]: %w", i, err)
		}
	}
	return nil
}

type EvidenceStatistic struct {
	Id                  *string                                `json:"id,omitempty" bson:"id,omitempty"`                                    // Unique id for inter-element referencing
	Description         *string                                `json:"description,omitempty" bson:"description,omitempty"`                  // A natural language summary of the statistic
	Note                []Annotation                           `json:"note,omitempty" bson:"note,omitempty"`                                // Footnotes and/or explanatory notes
	StatisticType       *CodeableConcept                       `json:"statisticType,omitempty" bson:"statistic_type,omitempty"`             // Type of statistic, e.g., relative risk
	Category            *CodeableConcept                       `json:"category,omitempty" bson:"category,omitempty"`                        // Associated category for categorical variable
	Quantity            *Quantity                              `json:"quantity,omitempty" bson:"quantity,omitempty"`                        // Statistic value
	NumberOfEvents      *int                                   `json:"numberOfEvents,omitempty" bson:"number_of_events,omitempty"`          // The number of events associated with the statistic
	NumberAffected      *int                                   `json:"numberAffected,omitempty" bson:"number_affected,omitempty"`           // The number of participants affected
	SampleSize          *EvidenceStatisticSampleSize           `json:"sampleSize,omitempty" bson:"sample_size,omitempty"`                   // Count of participants in the study sample
	AttributeEstimate   []EvidenceStatisticAttributeEstimate   `json:"attributeEstimate,omitempty" bson:"attribute_estimate,omitempty"`     // An attribute of the Statistic
	ModelCharacteristic []EvidenceStatisticModelCharacteristic `json:"modelCharacteristic,omitempty" bson:"model_characteristic,omitempty"` // An aspect of the statistical model
}

func (r *EvidenceStatistic) Validate() error {
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	if r.StatisticType != nil {
		if err := r.StatisticType.Validate(); err != nil {
			return fmt.Errorf("StatisticType: %w", err)
		}
	}
	if r.Category != nil {
		if err := r.Category.Validate(); err != nil {
			return fmt.Errorf("Category: %w", err)
		}
	}
	if r.Quantity != nil {
		if err := r.Quantity.Validate(); err != nil {
			return fmt.Errorf("Quantity: %w", err)
		}
	}
	if r.SampleSize != nil {
		if err := r.SampleSize.Validate(); err != nil {
			return fmt.Errorf("SampleSize: %w", err)
		}
	}
	for i, item := range r.AttributeEstimate {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AttributeEstimate[%d]: %w", i, err)
		}
	}
	for i, item := range r.ModelCharacteristic {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ModelCharacteristic[%d]: %w", i, err)
		}
	}
	return nil
}

type EvidenceStatisticSampleSize struct {
	Id                   *string      `json:"id,omitempty" bson:"id,omitempty"`                                       // Unique id for inter-element referencing
	Description          *string      `json:"description,omitempty" bson:"description,omitempty"`                     // Textual description of sample size for statistic
	Note                 []Annotation `json:"note,omitempty" bson:"note,omitempty"`                                   // Footnote or explanatory note about the sample size
	NumberOfStudies      *int         `json:"numberOfStudies,omitempty" bson:"number_of_studies,omitempty"`           // Number of contributing studies
	NumberOfParticipants *int         `json:"numberOfParticipants,omitempty" bson:"number_of_participants,omitempty"` // Total number of participants
	KnownDataCount       *int         `json:"knownDataCount,omitempty" bson:"known_data_count,omitempty"`             // Number of participants with known results for measured variables
	NumberAnalyzed       *int         `json:"numberAnalyzed,omitempty" bson:"number_analyzed,omitempty"`              // Total number of participants who were analayzed
}

func (r *EvidenceStatisticSampleSize) Validate() error {
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	return nil
}

type EvidenceStatisticModelCharacteristic struct {
	Id                   *string                                        `json:"id,omitempty" bson:"id,omitempty"`                                       // Unique id for inter-element referencing
	Code                 *CodeableConcept                               `json:"code" bson:"code"`                                                       // Model specification
	ValueQuantity        *Quantity                                      `json:"valueQuantity,omitempty" bson:"value_quantity,omitempty"`                // The specific value (when paired with code)
	ValueRange           *Range                                         `json:"valueRange,omitempty" bson:"value_range,omitempty"`                      // The specific value (when paired with code)
	ValueCodeableConcept *CodeableConcept                               `json:"valueCodeableConcept,omitempty" bson:"value_codeable_concept,omitempty"` // The specific value (when paired with code)
	Intended             bool                                           `json:"intended,omitempty" bson:"intended,omitempty"`                           // The plan for analysis
	Applied              bool                                           `json:"applied,omitempty" bson:"applied,omitempty"`                             // This model characteristic is part of the analysis that was applied, whether or not the analysis followed the plan
	Variable             []EvidenceStatisticModelCharacteristicVariable `json:"variable,omitempty" bson:"variable,omitempty"`                           // A variable adjusted for in the adjusted analysis
	Attribute            []EvidenceStatisticAttributeEstimate           `json:"attribute,omitempty" bson:"attribute,omitempty"`                         // An attribute of the model characteristic
}

func (r *EvidenceStatisticModelCharacteristic) Validate() error {
	if r.Code == nil {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
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
	if r.ValueCodeableConcept != nil {
		if err := r.ValueCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("ValueCodeableConcept: %w", err)
		}
	}
	for i, item := range r.Variable {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Variable[%d]: %w", i, err)
		}
	}
	for i, item := range r.Attribute {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Attribute[%d]: %w", i, err)
		}
	}
	return nil
}

type EvidenceRelatesTo struct {
	Id               *string          `json:"id,omitempty" bson:"id,omitempty"`          // Unique id for inter-element referencing
	Type             *CodeableConcept `json:"type" bson:"type"`                          // documentation | justification | citation | predecessor | successor | derived-from | depends-on | composed-of | part-of | amends | amended-with | appends | appended-with | cites | cited-by | comments-on | comment-in | contains | contained-in | corrects | correction-in | replaces | replaced-with | retracts | retracted-by | signs | similar-to | supports | supported-with | transforms | transformed-into | transformed-with | documents | specification-of | created-with | cite-as | reprint | reprint-of | summarizes
	TargetUri        *string          `json:"targetUri" bson:"target_uri"`               // The artifact that is related to this Evidence
	TargetAttachment *Attachment      `json:"targetAttachment" bson:"target_attachment"` // The artifact that is related to this Evidence
	TargetCanonical  *string          `json:"targetCanonical" bson:"target_canonical"`   // The artifact that is related to this Evidence
	TargetReference  *Reference       `json:"targetReference" bson:"target_reference"`   // The artifact that is related to this Evidence
	TargetMarkdown   *string          `json:"targetMarkdown" bson:"target_markdown"`     // The artifact that is related to this Evidence
}

func (r *EvidenceRelatesTo) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.TargetUri == nil {
		return fmt.Errorf("field 'TargetUri' is required")
	}
	if r.TargetAttachment == nil {
		return fmt.Errorf("field 'TargetAttachment' is required")
	}
	if r.TargetAttachment != nil {
		if err := r.TargetAttachment.Validate(); err != nil {
			return fmt.Errorf("TargetAttachment: %w", err)
		}
	}
	if r.TargetCanonical == nil {
		return fmt.Errorf("field 'TargetCanonical' is required")
	}
	if r.TargetReference == nil {
		return fmt.Errorf("field 'TargetReference' is required")
	}
	if r.TargetReference != nil {
		if err := r.TargetReference.Validate(); err != nil {
			return fmt.Errorf("TargetReference: %w", err)
		}
	}
	if r.TargetMarkdown == nil {
		return fmt.Errorf("field 'TargetMarkdown' is required")
	}
	return nil
}

type EvidenceVariableDefinition struct {
	Id                 *string          `json:"id,omitempty" bson:"id,omitempty"`                                  // Unique id for inter-element referencing
	Description        *string          `json:"description,omitempty" bson:"description,omitempty"`                // A text description or summary of the variable
	Note               []Annotation     `json:"note,omitempty" bson:"note,omitempty"`                              // Footnotes and/or explanatory notes
	VariableRole       string           `json:"variableRole" bson:"variable_role"`                                 // population | exposure | outcome | covariate
	RoleSubtype        *CodeableConcept `json:"roleSubtype,omitempty" bson:"role_subtype,omitempty"`               // Sub-classification of the role of the variable
	ComparatorCategory *string          `json:"comparatorCategory,omitempty" bson:"comparator_category,omitempty"` // The reference value used for comparison
	Observed           *Reference       `json:"observed,omitempty" bson:"observed,omitempty"`                      // Definition of the actual variable related to the statistic(s)
	Intended           *Reference       `json:"intended,omitempty" bson:"intended,omitempty"`                      // Definition of the intended variable related to the Evidence
	DirectnessMatch    *CodeableConcept `json:"directnessMatch,omitempty" bson:"directness_match,omitempty"`       // low | moderate | high | exact
}

func (r *EvidenceVariableDefinition) Validate() error {
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	var emptyString string
	if r.VariableRole == emptyString {
		return fmt.Errorf("field 'VariableRole' is required")
	}
	if r.RoleSubtype != nil {
		if err := r.RoleSubtype.Validate(); err != nil {
			return fmt.Errorf("RoleSubtype: %w", err)
		}
	}
	if r.Observed != nil {
		if err := r.Observed.Validate(); err != nil {
			return fmt.Errorf("Observed: %w", err)
		}
	}
	if r.Intended != nil {
		if err := r.Intended.Validate(); err != nil {
			return fmt.Errorf("Intended: %w", err)
		}
	}
	if r.DirectnessMatch != nil {
		if err := r.DirectnessMatch.Validate(); err != nil {
			return fmt.Errorf("DirectnessMatch: %w", err)
		}
	}
	return nil
}

type EvidenceStatisticAttributeEstimate struct {
	Id                *string                              `json:"id,omitempty" bson:"id,omitempty"`                                // Unique id for inter-element referencing
	Description       *string                              `json:"description,omitempty" bson:"description,omitempty"`              // Textual description of the attribute estimate
	Note              []Annotation                         `json:"note,omitempty" bson:"note,omitempty"`                            // Footnote or explanatory note about the estimate
	Type              *CodeableConcept                     `json:"type,omitempty" bson:"type,omitempty"`                            // The type of attribute estimate, e.g., confidence interval or p value
	Quantity          *Quantity                            `json:"quantity,omitempty" bson:"quantity,omitempty"`                    // The singular quantity of the attribute estimate, for attribute estimates represented as single values, which may include a unit of measure
	Level             *float64                             `json:"level,omitempty" bson:"level,omitempty"`                          // Level of confidence interval, e.g., 0.95 for 95% confidence interval
	Range             *Range                               `json:"range,omitempty" bson:"range,omitempty"`                          // Lower and upper bound values of the attribute estimate
	AttributeEstimate []EvidenceStatisticAttributeEstimate `json:"attributeEstimate,omitempty" bson:"attribute_estimate,omitempty"` // A nested attribute estimate; which is the attribute estimate of an attribute estimate
}

func (r *EvidenceStatisticAttributeEstimate) Validate() error {
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Quantity != nil {
		if err := r.Quantity.Validate(); err != nil {
			return fmt.Errorf("Quantity: %w", err)
		}
	}
	if r.Range != nil {
		if err := r.Range.Validate(); err != nil {
			return fmt.Errorf("Range: %w", err)
		}
	}
	for i, item := range r.AttributeEstimate {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AttributeEstimate[%d]: %w", i, err)
		}
	}
	return nil
}

type EvidenceStatisticModelCharacteristicVariable struct {
	Id                 *string           `json:"id,omitempty" bson:"id,omitempty"`                        // Unique id for inter-element referencing
	VariableDefinition *Reference        `json:"variableDefinition" bson:"variable_definition"`           // Description and definition of the variable
	Handling           *CodeableConcept  `json:"handling,omitempty" bson:"handling,omitempty"`            // boolean | continuous | dichotomous | ordinal | polychotomous | time-to-event | not-specified
	ValueCategory      []CodeableConcept `json:"valueCategory,omitempty" bson:"value_category,omitempty"` // Qualitative label used for grouping values of a dichotomous, ordinal, or polychotomous variable
	ValueQuantity      []Quantity        `json:"valueQuantity,omitempty" bson:"value_quantity,omitempty"` // Quantitative label used for grouping values of a dichotomous, ordinal, or polychotomous variable
	ValueRange         []Range           `json:"valueRange,omitempty" bson:"value_range,omitempty"`       // Range of quantitative labels used for grouping values of a dichotomous, ordinal, or polychotomous variable
}

func (r *EvidenceStatisticModelCharacteristicVariable) Validate() error {
	if r.VariableDefinition == nil {
		return fmt.Errorf("field 'VariableDefinition' is required")
	}
	if r.VariableDefinition != nil {
		if err := r.VariableDefinition.Validate(); err != nil {
			return fmt.Errorf("VariableDefinition: %w", err)
		}
	}
	if r.Handling != nil {
		if err := r.Handling.Validate(); err != nil {
			return fmt.Errorf("Handling: %w", err)
		}
	}
	for i, item := range r.ValueCategory {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ValueCategory[%d]: %w", i, err)
		}
	}
	for i, item := range r.ValueQuantity {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ValueQuantity[%d]: %w", i, err)
		}
	}
	for i, item := range r.ValueRange {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ValueRange[%d]: %w", i, err)
		}
	}
	return nil
}
