package models

import (
	"encoding/json"
	"fmt"
)

// The findings and interpretation of diagnostic tests performed on patients, groups of patients, products, substances, devices, and locations, and/or specimens derived from these. The report includes clinical context such as requesting provider information, and some mix of atomic results, images, textual and coded interpretations, and formatted representation of diagnostic reports. The report also includes non-clinical context such as batch analysis and stability reporting of products and substances.
type DiagnosticReport struct {
	Id                 *string                          `json:"id,omitempty" bson:"id,omitempty"`                                  // Logical id of this artifact
	Meta               *Meta                            `json:"meta,omitempty" bson:"meta,omitempty"`                              // Metadata about the resource
	ImplicitRules      *string                          `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`           // A set of rules under which this content was created
	Language           *string                          `json:"language,omitempty" bson:"language,omitempty"`                      // Language of the resource content
	Text               *Narrative                       `json:"text,omitempty" bson:"text,omitempty"`                              // Text summary of the resource, for human interpretation
	Contained          []json.RawMessage                `json:"contained,omitempty" bson:"contained,omitempty"`                    // Contained, inline Resources
	Identifier         []Identifier                     `json:"identifier,omitempty" bson:"identifier,omitempty"`                  // Business identifier for report
	BasedOn            []Reference                      `json:"basedOn,omitempty" bson:"based_on,omitempty"`                       // What was requested
	Status             string                           `json:"status" bson:"status"`                                              // registered | partial | preliminary | modified | final | amended | corrected | appended | cancelled | entered-in-error | unknown
	Category           []CodeableConcept                `json:"category,omitempty" bson:"category,omitempty"`                      // Service category
	Code               *CodeableConcept                 `json:"code" bson:"code"`                                                  // Name/Code for this diagnostic report
	Subject            *Reference                       `json:"subject,omitempty" bson:"subject,omitempty"`                        // The subject of the report - usually, but not always, the patient
	RelatesTo          []RelatedArtifact                `json:"relatesTo,omitempty" bson:"relates_to,omitempty"`                   // Related DiagnosticReports
	Encounter          *Reference                       `json:"encounter,omitempty" bson:"encounter,omitempty"`                    // Encounter associated with the DiagnosticReport
	EffectiveDateTime  *string                          `json:"effectiveDateTime,omitempty" bson:"effective_date_time,omitempty"`  // Clinically relevant time/time-period for the results that are included in the report
	EffectivePeriod    *Period                          `json:"effectivePeriod,omitempty" bson:"effective_period,omitempty"`       // Clinically relevant time/time-period for the results that are included in the report
	Issued             *string                          `json:"issued,omitempty" bson:"issued,omitempty"`                          // DateTime this version was made
	Procedure          []Reference                      `json:"procedure,omitempty" bson:"procedure,omitempty"`                    // The performed procedure(s) from which the report was produced
	Performer          []Reference                      `json:"performer,omitempty" bson:"performer,omitempty"`                    // Responsible Diagnostic Service
	ResultsInterpreter []Reference                      `json:"resultsInterpreter,omitempty" bson:"results_interpreter,omitempty"` // Who analyzed and reported the conclusions and interpretations
	Specimen           []Reference                      `json:"specimen,omitempty" bson:"specimen,omitempty"`                      // Specimens this report is based on
	Result             []Reference                      `json:"result,omitempty" bson:"result,omitempty"`                          // Observations
	Note               []Annotation                     `json:"note,omitempty" bson:"note,omitempty"`                              // Comments about the diagnostic report
	Study              []Reference                      `json:"study,omitempty" bson:"study,omitempty"`                            // Reference to full details of an analysis associated with the diagnostic report
	SupportingInfo     []DiagnosticReportSupportingInfo `json:"supportingInfo,omitempty" bson:"supporting_info,omitempty"`         // Additional information supporting the diagnostic report
	Media              []DiagnosticReportMedia          `json:"media,omitempty" bson:"media,omitempty"`                            // Key images or data associated with this report
	Composition        *Reference                       `json:"composition,omitempty" bson:"composition,omitempty"`                // Reference to a Composition resource for the DiagnosticReport structure
	Conclusion         *string                          `json:"conclusion,omitempty" bson:"conclusion,omitempty"`                  // Clinical conclusion (interpretation) of test results
	ConclusionCode     []CodeableReference              `json:"conclusionCode,omitempty" bson:"conclusion_code,omitempty"`         // Codes and/or references for the clinical conclusion of test results
	Recomendation      []CodeableReference              `json:"recomendation,omitempty" bson:"recomendation,omitempty"`            // Recommendations based on findings and interpretations
	PresentedForm      []Attachment                     `json:"presentedForm,omitempty" bson:"presented_form,omitempty"`           // Entire report as issued
	Communication      []Reference                      `json:"communication,omitempty" bson:"communication,omitempty"`            // Communication initiated during the reporting process
	Comparison         *Reference                       `json:"comparison,omitempty" bson:"comparison,omitempty"`                  // Prior data and findings for comparison
}

func (r *DiagnosticReport) Validate() error {
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
	for i, item := range r.BasedOn {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("BasedOn[%d]: %w", i, err)
		}
	}
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	for i, item := range r.Category {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Category[%d]: %w", i, err)
		}
	}
	if r.Code == nil {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.Subject != nil {
		if err := r.Subject.Validate(); err != nil {
			return fmt.Errorf("Subject: %w", err)
		}
	}
	for i, item := range r.RelatesTo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RelatesTo[%d]: %w", i, err)
		}
	}
	if r.Encounter != nil {
		if err := r.Encounter.Validate(); err != nil {
			return fmt.Errorf("Encounter: %w", err)
		}
	}
	if r.EffectivePeriod != nil {
		if err := r.EffectivePeriod.Validate(); err != nil {
			return fmt.Errorf("EffectivePeriod: %w", err)
		}
	}
	for i, item := range r.Procedure {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Procedure[%d]: %w", i, err)
		}
	}
	for i, item := range r.Performer {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Performer[%d]: %w", i, err)
		}
	}
	for i, item := range r.ResultsInterpreter {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ResultsInterpreter[%d]: %w", i, err)
		}
	}
	for i, item := range r.Specimen {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Specimen[%d]: %w", i, err)
		}
	}
	for i, item := range r.Result {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Result[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	for i, item := range r.Study {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Study[%d]: %w", i, err)
		}
	}
	for i, item := range r.SupportingInfo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SupportingInfo[%d]: %w", i, err)
		}
	}
	for i, item := range r.Media {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Media[%d]: %w", i, err)
		}
	}
	if r.Composition != nil {
		if err := r.Composition.Validate(); err != nil {
			return fmt.Errorf("Composition: %w", err)
		}
	}
	for i, item := range r.ConclusionCode {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ConclusionCode[%d]: %w", i, err)
		}
	}
	for i, item := range r.Recomendation {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Recomendation[%d]: %w", i, err)
		}
	}
	for i, item := range r.PresentedForm {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("PresentedForm[%d]: %w", i, err)
		}
	}
	for i, item := range r.Communication {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Communication[%d]: %w", i, err)
		}
	}
	if r.Comparison != nil {
		if err := r.Comparison.Validate(); err != nil {
			return fmt.Errorf("Comparison: %w", err)
		}
	}
	return nil
}

type DiagnosticReportSupportingInfo struct {
	Id        *string          `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Type      *CodeableConcept `json:"type" bson:"type"`                 // Supporting information role code
	Reference *Reference       `json:"reference" bson:"reference"`       // Supporting information reference
}

func (r *DiagnosticReportSupportingInfo) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Reference == nil {
		return fmt.Errorf("field 'Reference' is required")
	}
	if r.Reference != nil {
		if err := r.Reference.Validate(); err != nil {
			return fmt.Errorf("Reference: %w", err)
		}
	}
	return nil
}

type DiagnosticReportMedia struct {
	Id      *string    `json:"id,omitempty" bson:"id,omitempty"`           // Unique id for inter-element referencing
	Comment *string    `json:"comment,omitempty" bson:"comment,omitempty"` // Comment about the image or data (e.g. explanation)
	Link    *Reference `json:"link" bson:"link"`                           // Reference to the image or data source
}

func (r *DiagnosticReportMedia) Validate() error {
	if r.Link == nil {
		return fmt.Errorf("field 'Link' is required")
	}
	if r.Link != nil {
		if err := r.Link.Validate(); err != nil {
			return fmt.Errorf("Link: %w", err)
		}
	}
	return nil
}
