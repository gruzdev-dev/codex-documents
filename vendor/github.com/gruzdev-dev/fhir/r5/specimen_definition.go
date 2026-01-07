package models

import (
	"encoding/json"
	"fmt"
)

// A kind of specimen with associated set of requirements.
type SpecimenDefinition struct {
	Id                     *string                        `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta                          `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string                        `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string                        `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative                     `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage              `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Url                    *string                        `json:"url,omitempty" bson:"url,omitempty"`                                         // Logical canonical URL to reference this SpecimenDefinition (globally unique)
	Identifier             *Identifier                    `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // Business identifier
	Version                *string                        `json:"version,omitempty" bson:"version,omitempty"`                                 // Business version of the SpecimenDefinition
	VersionAlgorithmString *string                        `json:"versionAlgorithmString,omitempty" bson:"version_algorithm_string,omitempty"` // How to compare versions
	VersionAlgorithmCoding *Coding                        `json:"versionAlgorithmCoding,omitempty" bson:"version_algorithm_coding,omitempty"` // How to compare versions
	Name                   *string                        `json:"name,omitempty" bson:"name,omitempty"`                                       // Name for this {{title}} (computer friendly)
	Title                  *string                        `json:"title,omitempty" bson:"title,omitempty"`                                     // Name for this SpecimenDefinition (Human friendly)
	DerivedFromCanonical   []string                       `json:"derivedFromCanonical,omitempty" bson:"derived_from_canonical,omitempty"`     // Based on FHIR definition of another SpecimenDefinition
	DerivedFromUri         []string                       `json:"derivedFromUri,omitempty" bson:"derived_from_uri,omitempty"`                 // Based on external definition
	Status                 string                         `json:"status" bson:"status"`                                                       // draft | active | retired | unknown
	Experimental           bool                           `json:"experimental,omitempty" bson:"experimental,omitempty"`                       // If this SpecimenDefinition is not for real usage
	SubjectCodeableConcept *CodeableConcept               `json:"subjectCodeableConcept,omitempty" bson:"subject_codeable_concept,omitempty"` // Type of subject for specimen collection
	SubjectReference       *Reference                     `json:"subjectReference,omitempty" bson:"subject_reference,omitempty"`              // Type of subject for specimen collection
	Date                   *string                        `json:"date,omitempty" bson:"date,omitempty"`                                       // Date status first applied
	Publisher              *string                        `json:"publisher,omitempty" bson:"publisher,omitempty"`                             // The name of the individual or organization that published the SpecimenDefinition
	Contact                []ContactDetail                `json:"contact,omitempty" bson:"contact,omitempty"`                                 // Contact details for the publisher
	Description            *string                        `json:"description,omitempty" bson:"description,omitempty"`                         // Natural language description of the SpecimenDefinition
	UseContext             []UsageContext                 `json:"useContext,omitempty" bson:"use_context,omitempty"`                          // Content intends to support these contexts
	Jurisdiction           []CodeableConcept              `json:"jurisdiction,omitempty" bson:"jurisdiction,omitempty"`                       // Jurisdiction of the authority that maintains the this SpecimenDefinition (if applicable)
	Purpose                *string                        `json:"purpose,omitempty" bson:"purpose,omitempty"`                                 // Why this SpecimenDefinition is defined
	Copyright              *string                        `json:"copyright,omitempty" bson:"copyright,omitempty"`                             // Notice about intellectual property ownership, can include restrictions on use
	CopyrightLabel         *string                        `json:"copyrightLabel,omitempty" bson:"copyright_label,omitempty"`                  // Copyright holder and year(s)
	ApprovalDate           *string                        `json:"approvalDate,omitempty" bson:"approval_date,omitempty"`                      // When SpecimenDefinition was approved by publisher
	LastReviewDate         *string                        `json:"lastReviewDate,omitempty" bson:"last_review_date,omitempty"`                 // The date on which the asset content was last reviewed by the publisher
	EffectivePeriod        *Period                        `json:"effectivePeriod,omitempty" bson:"effective_period,omitempty"`                // The effective date range for the SpecimenDefinition
	TypeCollected          *CodeableConcept               `json:"typeCollected,omitempty" bson:"type_collected,omitempty"`                    // Kind of material to collect
	PatientPreparation     []CodeableConcept              `json:"patientPreparation,omitempty" bson:"patient_preparation,omitempty"`          // Patient preparation for collection
	TimeAspect             *string                        `json:"timeAspect,omitempty" bson:"time_aspect,omitempty"`                          // Time aspect for collection
	Collection             []CodeableConcept              `json:"collection,omitempty" bson:"collection,omitempty"`                           // Specimen collection procedure
	TypeTested             []SpecimenDefinitionTypeTested `json:"typeTested,omitempty" bson:"type_tested,omitempty"`                          // Specimen in container intended for testing by lab
}

func (r *SpecimenDefinition) Validate() error {
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
	if r.TypeCollected != nil {
		if err := r.TypeCollected.Validate(); err != nil {
			return fmt.Errorf("TypeCollected: %w", err)
		}
	}
	for i, item := range r.PatientPreparation {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("PatientPreparation[%d]: %w", i, err)
		}
	}
	for i, item := range r.Collection {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Collection[%d]: %w", i, err)
		}
	}
	for i, item := range r.TypeTested {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("TypeTested[%d]: %w", i, err)
		}
	}
	return nil
}

type SpecimenDefinitionTypeTestedContainer struct {
	Id                    *string                                         `json:"id,omitempty" bson:"id,omitempty"`                                         // Unique id for inter-element referencing
	Material              *CodeableConcept                                `json:"material,omitempty" bson:"material,omitempty"`                             // The material type used for the container
	Type                  *CodeableConcept                                `json:"type,omitempty" bson:"type,omitempty"`                                     // Kind of container associated with the kind of specimen
	Cap                   *CodeableConcept                                `json:"cap,omitempty" bson:"cap,omitempty"`                                       // Color of container cap
	Description           *string                                         `json:"description,omitempty" bson:"description,omitempty"`                       // The description of the kind of container
	Capacity              *Quantity                                       `json:"capacity,omitempty" bson:"capacity,omitempty"`                             // The capacity of this kind of container
	MinimumVolumeQuantity *Quantity                                       `json:"minimumVolumeQuantity,omitempty" bson:"minimum_volume_quantity,omitempty"` // Minimum volume
	MinimumVolumeString   *string                                         `json:"minimumVolumeString,omitempty" bson:"minimum_volume_string,omitempty"`     // Minimum volume
	Additive              []SpecimenDefinitionTypeTestedContainerAdditive `json:"additive,omitempty" bson:"additive,omitempty"`                             // Additive associated with container
	Preparation           *string                                         `json:"preparation,omitempty" bson:"preparation,omitempty"`                       // Special processing applied to the container for this specimen type
}

func (r *SpecimenDefinitionTypeTestedContainer) Validate() error {
	if r.Material != nil {
		if err := r.Material.Validate(); err != nil {
			return fmt.Errorf("Material: %w", err)
		}
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Cap != nil {
		if err := r.Cap.Validate(); err != nil {
			return fmt.Errorf("Cap: %w", err)
		}
	}
	if r.Capacity != nil {
		if err := r.Capacity.Validate(); err != nil {
			return fmt.Errorf("Capacity: %w", err)
		}
	}
	if r.MinimumVolumeQuantity != nil {
		if err := r.MinimumVolumeQuantity.Validate(); err != nil {
			return fmt.Errorf("MinimumVolumeQuantity: %w", err)
		}
	}
	for i, item := range r.Additive {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Additive[%d]: %w", i, err)
		}
	}
	return nil
}

type SpecimenDefinitionTypeTestedContainerAdditive struct {
	Id                      *string          `json:"id,omitempty" bson:"id,omitempty"`                         // Unique id for inter-element referencing
	AdditiveCodeableConcept *CodeableConcept `json:"additiveCodeableConcept" bson:"additive_codeable_concept"` // Additive associated with container
	AdditiveReference       *Reference       `json:"additiveReference" bson:"additive_reference"`              // Additive associated with container
}

func (r *SpecimenDefinitionTypeTestedContainerAdditive) Validate() error {
	if r.AdditiveCodeableConcept == nil {
		return fmt.Errorf("field 'AdditiveCodeableConcept' is required")
	}
	if r.AdditiveCodeableConcept != nil {
		if err := r.AdditiveCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("AdditiveCodeableConcept: %w", err)
		}
	}
	if r.AdditiveReference == nil {
		return fmt.Errorf("field 'AdditiveReference' is required")
	}
	if r.AdditiveReference != nil {
		if err := r.AdditiveReference.Validate(); err != nil {
			return fmt.Errorf("AdditiveReference: %w", err)
		}
	}
	return nil
}

type SpecimenDefinitionTypeTestedHandling struct {
	Id                   *string          `json:"id,omitempty" bson:"id,omitempty"`                                      // Unique id for inter-element referencing
	TemperatureQualifier *CodeableConcept `json:"temperatureQualifier,omitempty" bson:"temperature_qualifier,omitempty"` // Qualifies the interval of temperature
	TemperatureRange     *Range           `json:"temperatureRange,omitempty" bson:"temperature_range,omitempty"`         // Temperature range for these handling instructions
	MaxDuration          *Duration        `json:"maxDuration,omitempty" bson:"max_duration,omitempty"`                   // Maximum preservation time
	Instruction          *string          `json:"instruction,omitempty" bson:"instruction,omitempty"`                    // Preservation instruction
}

func (r *SpecimenDefinitionTypeTestedHandling) Validate() error {
	if r.TemperatureQualifier != nil {
		if err := r.TemperatureQualifier.Validate(); err != nil {
			return fmt.Errorf("TemperatureQualifier: %w", err)
		}
	}
	if r.TemperatureRange != nil {
		if err := r.TemperatureRange.Validate(); err != nil {
			return fmt.Errorf("TemperatureRange: %w", err)
		}
	}
	if r.MaxDuration != nil {
		if err := r.MaxDuration.Validate(); err != nil {
			return fmt.Errorf("MaxDuration: %w", err)
		}
	}
	return nil
}

type SpecimenDefinitionTypeTested struct {
	Id                 *string                                `json:"id,omitempty" bson:"id,omitempty"`                                  // Unique id for inter-element referencing
	IsDerived          bool                                   `json:"isDerived,omitempty" bson:"is_derived,omitempty"`                   // Primary or secondary specimen
	Type               *CodeableConcept                       `json:"type,omitempty" bson:"type,omitempty"`                              // Type of intended specimen
	Preference         string                                 `json:"preference" bson:"preference"`                                      // preferred | alternate
	Container          *SpecimenDefinitionTypeTestedContainer `json:"container,omitempty" bson:"container,omitempty"`                    // The specimen's container
	Requirement        *string                                `json:"requirement,omitempty" bson:"requirement,omitempty"`                // Requirements for specimen delivery and special handling
	RetentionTime      *Duration                              `json:"retentionTime,omitempty" bson:"retention_time,omitempty"`           // The usual time for retaining this kind of specimen
	SingleUse          bool                                   `json:"singleUse,omitempty" bson:"single_use,omitempty"`                   // Specimen for single use only
	RejectionCriterion []CodeableConcept                      `json:"rejectionCriterion,omitempty" bson:"rejection_criterion,omitempty"` // Criterion specified for specimen rejection
	Handling           []SpecimenDefinitionTypeTestedHandling `json:"handling,omitempty" bson:"handling,omitempty"`                      // Specimen handling before testing
	TestingDestination []CodeableConcept                      `json:"testingDestination,omitempty" bson:"testing_destination,omitempty"` // Where the specimen will be tested
}

func (r *SpecimenDefinitionTypeTested) Validate() error {
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	var emptyString string
	if r.Preference == emptyString {
		return fmt.Errorf("field 'Preference' is required")
	}
	if r.Container != nil {
		if err := r.Container.Validate(); err != nil {
			return fmt.Errorf("Container: %w", err)
		}
	}
	if r.RetentionTime != nil {
		if err := r.RetentionTime.Validate(); err != nil {
			return fmt.Errorf("RetentionTime: %w", err)
		}
	}
	for i, item := range r.RejectionCriterion {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RejectionCriterion[%d]: %w", i, err)
		}
	}
	for i, item := range r.Handling {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Handling[%d]: %w", i, err)
		}
	}
	for i, item := range r.TestingDestination {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("TestingDestination[%d]: %w", i, err)
		}
	}
	return nil
}
