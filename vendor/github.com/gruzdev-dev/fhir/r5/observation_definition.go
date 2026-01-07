package models

import (
	"encoding/json"
	"fmt"
)

// Set of definitional characteristics for a kind of observation or measurement produced or consumed by an orderable health care service.
type ObservationDefinition struct {
	ResourceType           string                                `json:"resourceType" bson:"resource_type"`                                          // Type of resource
	Id                     *string                               `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta                                 `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string                               `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string                               `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative                            `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage                     `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Url                    *string                               `json:"url,omitempty" bson:"url,omitempty"`                                         // Logical canonical URL to reference this ObservationDefinition (globally unique)
	Identifier             *Identifier                           `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // Business identifier of the ObservationDefinition
	Version                *string                               `json:"version,omitempty" bson:"version,omitempty"`                                 // Business version of the ObservationDefinition
	VersionAlgorithmString *string                               `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"` // How to compare versions
	VersionAlgorithmCoding *Coding                               `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"` // How to compare versions
	Name                   *string                               `json:"name,omitempty" bson:"name,omitempty"`                                       // Name for this ObservationDefinition (computer friendly)
	Title                  *string                               `json:"title,omitempty" bson:"title,omitempty"`                                     // Name for this ObservationDefinition (human friendly)
	Status                 string                                `json:"status" bson:"status"`                                                       // draft | active | retired | unknown
	Experimental           bool                                  `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // If For testing only - never for real usage
	Date                   *string                               `json:"date,omitempty" bson:"date,omitempty"`                                       // Date last changed
	Publisher              *string                               `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // The name of the individual or organization that published the ObservationDefinition
	Contact                []ContactDetail                       `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Description            *string                               `json:"description,omitempty" bson:"description,omitempty"`                         // Natural language description of the ObservationDefinition
	UseContext             []UsageContext                        `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // Content intends to support these contexts
	Jurisdiction           []CodeableConcept                     `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                       // Jurisdiction of the authority that maintains the this ObservationDefinition (if applicable)
	Purpose                *string                               `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // Why this ObservationDefinition is defined
	Copyright              *string                               `json:"copyright,omitempty" bson:"copyright,omitempty"`                             // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel         *string                               `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                  // Copyright holder and year(s)
	ApprovalDate           *string                               `json:"approvalDate,omitempty" bson:"approval_date,omitempty"`                      // When ObservationDefinition was approved by publisher
	LastReviewDate         *string                               `json:"lastReviewDate,omitempty" bson:"last_review_date,omitempty"`                 // Date on which the asset content was last reviewed by the publisher
	EffectivePeriod        *Period                               `json:"effectivePeriod,omitempty" bson:"effective_period,omitempty"`                // The effective date range for the ObservationDefinition
	DerivedFromCanonical   []string                              `json:"derivedFromCanonical,omitempty" bson:"derived_from_canonical,omitempty"`     // Based on FHIR definition of another observation
	DerivedFromUri         []string                              `json:"derivedFromUri,omitempty" bson:"derived_from_uri,omitempty"`                 // Based on external definition
	Subject                []CodeableConcept                     `json:"subject,omitempty" bson:"subject,omitempty"`                                 // Type of subject for the defined observation
	PerformerType          *CodeableConcept                      `json:"performerType,omitempty" bson:"performer_type,omitempty"`                    // Desired kind of performer for such kind of observation
	Category               []CodeableConcept                     `json:"category,omitempty" bson:"category,omitempty"`                               // General type of observation
	Code                   *CodeableConcept                      `json:"code" bson:"code"`                                                           // Type of observation
	PermittedDataType      []string                              `json:"permittedDataType,omitempty" bson:"permitted_data_type,omitempty"`           // Quantity | CodeableConcept | string | boolean | integer | Range | Ratio | SampledData | time | dateTime | Period
	MultipleResultsAllowed bool                                  `json:"multipleResultsAllowed,omitempty" bson:"multiple_results_allowed,omitempty"` // Multiple results allowed for conforming observations
	BodyStructure          *CodeableReference                    `json:"bodyStructure,omitempty" bson:"body_structure,omitempty"`                    // Body structure to be observed
	Method                 *CodeableConcept                      `json:"method,omitempty" bson:"method,omitempty"`                                   // Method used to produce the observation
	Specimen               []Reference                           `json:"specimen,omitempty" bson:"specimen,omitempty"`                               // Kind of specimen used by this type of observation
	DeviceReference        *Reference                            `json:"deviceReference,omitempty" bson:"device_reference,omitempty"`                // Measurement device or model of device
	DeviceCanonical        *string                               `json:"deviceCanonical,omitempty" bson:"device_canonical,omitempty"`                // Measurement device or model of device
	PreferredReportName    *string                               `json:"preferredReportName,omitempty" bson:"preferred_report_name,omitempty"`       // The preferred name to be used when reporting the observation results
	PermittedUnit          []Coding                              `json:"permittedUnit,omitempty" bson:"permitted_unit,omitempty"`                    // Unit for quantitative results
	QualifiedValue         []ObservationDefinitionQualifiedValue `json:"qualifiedValue,omitempty" bson:"qualified_value,omitempty"`                  // Set of qualified values for observation results
	HasMember              []Reference                           `json:"hasMember,omitempty" bson:"has_member,omitempty"`                            // Definitions of related resources belonging to this kind of observation group
	Component              []ObservationDefinitionComponent      `json:"component,omitempty" bson:"component,omitempty"`                             // Component results
}

func (r *ObservationDefinition) Validate() error {
	if r.ResourceType != "ObservationDefinition" {
		return fmt.Errorf("invalid resourceType: expected 'ObservationDefinition', got '%s'", r.ResourceType)
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
	if r.Identifier != nil {
		if err := r.Identifier.Validate(); err != nil {
			return fmt.Errorf("Identifier: %w", err)
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
	for i, item := range r.Subject {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Subject[%d]: %w", i, err)
		}
	}
	if r.PerformerType != nil {
		if err := r.PerformerType.Validate(); err != nil {
			return fmt.Errorf("PerformerType: %w", err)
		}
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
	if r.BodyStructure != nil {
		if err := r.BodyStructure.Validate(); err != nil {
			return fmt.Errorf("BodyStructure: %w", err)
		}
	}
	if r.Method != nil {
		if err := r.Method.Validate(); err != nil {
			return fmt.Errorf("Method: %w", err)
		}
	}
	for i, item := range r.Specimen {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Specimen[%d]: %w", i, err)
		}
	}
	if r.DeviceReference != nil {
		if err := r.DeviceReference.Validate(); err != nil {
			return fmt.Errorf("DeviceReference: %w", err)
		}
	}
	for i, item := range r.PermittedUnit {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("PermittedUnit[%d]: %w", i, err)
		}
	}
	for i, item := range r.QualifiedValue {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("QualifiedValue[%d]: %w", i, err)
		}
	}
	for i, item := range r.HasMember {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("HasMember[%d]: %w", i, err)
		}
	}
	for i, item := range r.Component {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Component[%d]: %w", i, err)
		}
	}
	return nil
}

type ObservationDefinitionQualifiedValue struct {
	Id                         *string           `json:"id,omitempty" bson:"id,omitempty"`                                                     // Unique id for inter-element referencing
	Context                    *CodeableConcept  `json:"context,omitempty" bson:"context,omitempty"`                                           // Context qualifier for the set of qualified values
	AppliesTo                  []CodeableConcept `json:"appliesTo,omitempty" bson:"applies_to,omitempty"`                                      // Targetted population for the set of qualified values
	SexParameterForClinicalUse *string           `json:"sexParameterForClinicalUse,omitempty" bson:"sex_parameter_for_clinical_use,omitempty"` // male | female | other | unknown
	Age                        *Range            `json:"age,omitempty" bson:"age,omitempty"`                                                   // Applicable age range for the set of qualified values
	GestationalAge             *Range            `json:"gestationalAge,omitempty" bson:"gestational_age,omitempty"`                            // Applicable gestational age range for the set of qualified values
	Condition                  *string           `json:"condition,omitempty" bson:"condition,omitempty"`                                       // Validity criterion for the qualified value
	RangeCategory              *string           `json:"rangeCategory,omitempty" bson:"range_category,omitempty"`                              // reference | critical | absolute
	Range                      *Range            `json:"range,omitempty" bson:"range,omitempty"`                                               // The range for continuous or ordinal observations
	ValidCodedValueSet         *string           `json:"validCodedValueSet,omitempty" bson:"valid_coded_value_set,omitempty"`                  // Value set of valid coded values as part of this set of qualified values
	NormalCodedValueSet        *string           `json:"normalCodedValueSet,omitempty" bson:"normal_coded_value_set,omitempty"`                // Value set of normal coded values as part of this set of qualified values
	AbnormalCodedValueSet      *string           `json:"abnormalCodedValueSet,omitempty" bson:"abnormal_coded_value_set,omitempty"`            // Value set of abnormal coded values as part of this set of qualified values
	CriticalCodedValueSet      *string           `json:"criticalCodedValueSet,omitempty" bson:"critical_coded_value_set,omitempty"`            // Value set of critical coded values as part of this set of qualified values
	Interpretation             []CodeableConcept `json:"interpretation,omitempty" bson:"interpretation,omitempty"`                             // Expected coded interpretation values
}

func (r *ObservationDefinitionQualifiedValue) Validate() error {
	if r.Context != nil {
		if err := r.Context.Validate(); err != nil {
			return fmt.Errorf("Context: %w", err)
		}
	}
	for i, item := range r.AppliesTo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AppliesTo[%d]: %w", i, err)
		}
	}
	if r.Age != nil {
		if err := r.Age.Validate(); err != nil {
			return fmt.Errorf("Age: %w", err)
		}
	}
	if r.GestationalAge != nil {
		if err := r.GestationalAge.Validate(); err != nil {
			return fmt.Errorf("GestationalAge: %w", err)
		}
	}
	if r.Range != nil {
		if err := r.Range.Validate(); err != nil {
			return fmt.Errorf("Range: %w", err)
		}
	}
	for i, item := range r.Interpretation {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Interpretation[%d]: %w", i, err)
		}
	}
	return nil
}

type ObservationDefinitionComponent struct {
	Id                *string                               `json:"id,omitempty" bson:"id,omitempty"`                                 // Unique id for inter-element referencing
	Code              *CodeableConcept                      `json:"code" bson:"code"`                                                 // Type of observation
	PermittedDataType []string                              `json:"permittedDataType,omitempty" bson:"permitted_data_type,omitempty"` // Quantity | CodeableConcept | string | boolean | integer | Range | Ratio | SampledData | time | dateTime | Period
	PermittedUnit     []Coding                              `json:"permittedUnit,omitempty" bson:"permitted_unit,omitempty"`          // Unit for quantitative results
	QualifiedValue    []ObservationDefinitionQualifiedValue `json:"qualifiedValue,omitempty" bson:"qualified_value,omitempty"`        // Set of qualified values for observation results
}

func (r *ObservationDefinitionComponent) Validate() error {
	if r.Code == nil {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	for i, item := range r.PermittedUnit {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("PermittedUnit[%d]: %w", i, err)
		}
	}
	for i, item := range r.QualifiedValue {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("QualifiedValue[%d]: %w", i, err)
		}
	}
	return nil
}
