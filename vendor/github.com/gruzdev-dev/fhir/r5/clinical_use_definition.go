package models

import (
	"encoding/json"
	"fmt"
)

// A single issue - either an indication, contraindication, interaction, undesirable effect or warning for a medicinal product, medication, device or procedure.
type ClinicalUseDefinition struct {
	ResourceType      string                                  `json:"resourceType" bson:"resource_type"`                               // Type of resource
	Id                *string                                 `json:"id,omitempty" bson:"id,omitempty"`                                // Logical id of this artifact
	Meta              *Meta                                   `json:"meta,omitempty" bson:"meta,omitempty"`                            // Metadata about the resource
	ImplicitRules     *string                                 `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`         // A set of rules under which this content was created
	Language          *string                                 `json:"language,omitempty" bson:"language,omitempty"`                    // Language of the resource content
	Text              *Narrative                              `json:"text,omitempty" bson:"text,omitempty"`                            // Text summary of the resource, for human interpretation
	Contained         []json.RawMessage                       `json:"contained,omitempty" bson:"contained,omitempty"`                  // Contained, inline Resources
	Identifier        []Identifier                            `json:"identifier,omitempty" bson:"identifier,omitempty"`                // Business identifier for this issue
	Type              string                                  `json:"type" bson:"type"`                                                // indication | contraindication | interaction | undesirable-effect | warning
	Category          []CodeableConcept                       `json:"category,omitempty" bson:"category,omitempty"`                    // A categorisation of the issue, primarily for dividing warnings into subject heading areas such as "Pregnancy", "Overdose"
	Subject           []CodeableReference                     `json:"subject" bson:"subject"`                                          // The medication, product, substance, device, procedure etc. for which this is an indication, contraindication, interaction, undesirable effect, or warning
	Status            *CodeableConcept                        `json:"status,omitempty" bson:"status,omitempty"`                        // Whether this is a current issue or one that has been retired etc
	UndesirableEffect *ClinicalUseDefinitionUndesirableEffect `json:"undesirableEffect,omitempty" bson:"undesirable_effect,omitempty"` // A possible negative outcome from the use of this treatment
	Indication        *ClinicalUseDefinitionIndication        `json:"indication,omitempty" bson:"indication,omitempty"`                // Specifics for when this is an indication
	Contraindication  *ClinicalUseDefinitionContraindication  `json:"contraindication,omitempty" bson:"contraindication,omitempty"`    // Specifics for when this is a contraindication
	Interaction       *ClinicalUseDefinitionInteraction       `json:"interaction,omitempty" bson:"interaction,omitempty"`              // Specifics for when this is an interaction
	Population        []Reference                             `json:"population,omitempty" bson:"population,omitempty"`                // The population group to which this applies
	Library           []string                                `json:"library,omitempty" bson:"library,omitempty"`                      // Logic used by the clinical use definition
	Warning           *ClinicalUseDefinitionWarning           `json:"warning,omitempty" bson:"warning,omitempty"`                      // Critical environmental, health or physical risks or hazards. For example 'Do not operate heavy machinery', 'May cause drowsiness'
}

func (r *ClinicalUseDefinition) Validate() error {
	if r.ResourceType != "ClinicalUseDefinition" {
		return fmt.Errorf("invalid resourceType: expected 'ClinicalUseDefinition', got '%s'", r.ResourceType)
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
	var emptyString string
	if r.Type == emptyString {
		return fmt.Errorf("field 'Type' is required")
	}
	for i, item := range r.Category {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Category[%d]: %w", i, err)
		}
	}
	if len(r.Subject) < 1 {
		return fmt.Errorf("field 'Subject' must have at least 1 elements")
	}
	for i, item := range r.Subject {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Subject[%d]: %w", i, err)
		}
	}
	if r.Status != nil {
		if err := r.Status.Validate(); err != nil {
			return fmt.Errorf("Status: %w", err)
		}
	}
	if r.UndesirableEffect != nil {
		if err := r.UndesirableEffect.Validate(); err != nil {
			return fmt.Errorf("UndesirableEffect: %w", err)
		}
	}
	if r.Indication != nil {
		if err := r.Indication.Validate(); err != nil {
			return fmt.Errorf("Indication: %w", err)
		}
	}
	if r.Contraindication != nil {
		if err := r.Contraindication.Validate(); err != nil {
			return fmt.Errorf("Contraindication: %w", err)
		}
	}
	if r.Interaction != nil {
		if err := r.Interaction.Validate(); err != nil {
			return fmt.Errorf("Interaction: %w", err)
		}
	}
	for i, item := range r.Population {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Population[%d]: %w", i, err)
		}
	}
	if r.Warning != nil {
		if err := r.Warning.Validate(); err != nil {
			return fmt.Errorf("Warning: %w", err)
		}
	}
	return nil
}

type ClinicalUseDefinitionUndesirableEffect struct {
	Id                     *string            `json:"id,omitempty" bson:"id,omitempty"`                                           // Unique id for inter-element referencing
	SymptomConditionEffect *CodeableReference `json:"symptomConditionEffect,omitempty" bson:"symptom_condition_effect,omitempty"` // The situation in which the undesirable effect may manifest
	Classification         *CodeableConcept   `json:"classification,omitempty" bson:"classification,omitempty"`                   // High level classification of the effect
	FrequencyOfOccurrence  *CodeableConcept   `json:"frequencyOfOccurrence,omitempty" bson:"frequency_of_occurrence,omitempty"`   // How often the effect is seen
	Management             []CodeableConcept  `json:"management,omitempty" bson:"management,omitempty"`                           // Actions for managing the undesirable effect
}

func (r *ClinicalUseDefinitionUndesirableEffect) Validate() error {
	if r.SymptomConditionEffect != nil {
		if err := r.SymptomConditionEffect.Validate(); err != nil {
			return fmt.Errorf("SymptomConditionEffect: %w", err)
		}
	}
	if r.Classification != nil {
		if err := r.Classification.Validate(); err != nil {
			return fmt.Errorf("Classification: %w", err)
		}
	}
	if r.FrequencyOfOccurrence != nil {
		if err := r.FrequencyOfOccurrence.Validate(); err != nil {
			return fmt.Errorf("FrequencyOfOccurrence: %w", err)
		}
	}
	for i, item := range r.Management {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Management[%d]: %w", i, err)
		}
	}
	return nil
}

type ClinicalUseDefinitionIndication struct {
	Id                      *string                                       `json:"id,omitempty" bson:"id,omitempty"`                                             // Unique id for inter-element referencing
	DiseaseSymptomProcedure *CodeableReference                            `json:"diseaseSymptomProcedure,omitempty" bson:"disease_symptom_procedure,omitempty"` // The situation that is being documented as an indication for this item
	DiseaseStatus           *CodeableReference                            `json:"diseaseStatus,omitempty" bson:"disease_status,omitempty"`                      // The status of the disease or symptom for the indication
	Comorbidity             []CodeableReference                           `json:"comorbidity,omitempty" bson:"comorbidity,omitempty"`                           // A comorbidity or coinfection as part of the indication
	IntendedEffect          []CodeableReference                           `json:"intendedEffect,omitempty" bson:"intended_effect,omitempty"`                    // The intended effect, aim or strategy to be achieved
	DurationRange           *Range                                        `json:"durationRange,omitempty" bson:"duration_range,omitempty"`                      // Timing or duration information
	DurationString          *string                                       `json:"durationString,omitempty" bson:"duration_string,omitempty"`                    // Timing or duration information
	UndesirableEffect       []ClinicalUseDefinitionUndesirableEffect      `json:"undesirableEffect,omitempty" bson:"undesirable_effect,omitempty"`              // An unwanted side effect or negative outcome of the subject of this resource when being used for this indication
	Applicability           *Expression                                   `json:"applicability,omitempty" bson:"applicability,omitempty"`                       // An expression that returns true or false, indicating whether the indication is applicable or not, after having applied its other elements
	OtherTherapy            []ClinicalUseDefinitionIndicationOtherTherapy `json:"otherTherapy,omitempty" bson:"other_therapy,omitempty"`                        // Information about use of the product in relation to other therapies described as part of the contraindication
}

func (r *ClinicalUseDefinitionIndication) Validate() error {
	if r.DiseaseSymptomProcedure != nil {
		if err := r.DiseaseSymptomProcedure.Validate(); err != nil {
			return fmt.Errorf("DiseaseSymptomProcedure: %w", err)
		}
	}
	if r.DiseaseStatus != nil {
		if err := r.DiseaseStatus.Validate(); err != nil {
			return fmt.Errorf("DiseaseStatus: %w", err)
		}
	}
	for i, item := range r.Comorbidity {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Comorbidity[%d]: %w", i, err)
		}
	}
	for i, item := range r.IntendedEffect {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("IntendedEffect[%d]: %w", i, err)
		}
	}
	if r.DurationRange != nil {
		if err := r.DurationRange.Validate(); err != nil {
			return fmt.Errorf("DurationRange: %w", err)
		}
	}
	for i, item := range r.UndesirableEffect {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("UndesirableEffect[%d]: %w", i, err)
		}
	}
	if r.Applicability != nil {
		if err := r.Applicability.Validate(); err != nil {
			return fmt.Errorf("Applicability: %w", err)
		}
	}
	for i, item := range r.OtherTherapy {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("OtherTherapy[%d]: %w", i, err)
		}
	}
	return nil
}

type ClinicalUseDefinitionContraindication struct {
	Id                      *string                                       `json:"id,omitempty" bson:"id,omitempty"`                                             // Unique id for inter-element referencing
	DiseaseSymptomProcedure *CodeableReference                            `json:"diseaseSymptomProcedure,omitempty" bson:"disease_symptom_procedure,omitempty"` // The situation that is being documented as contraindicating against this item
	DiseaseStatus           *CodeableReference                            `json:"diseaseStatus,omitempty" bson:"disease_status,omitempty"`                      // The status of the disease or symptom for the contraindication
	Comorbidity             []CodeableReference                           `json:"comorbidity,omitempty" bson:"comorbidity,omitempty"`                           // A comorbidity (concurrent condition) or coinfection
	Indication              []ClinicalUseDefinitionIndication             `json:"indication,omitempty" bson:"indication,omitempty"`                             // The indication which this is a contraindication for
	Applicability           *Expression                                   `json:"applicability,omitempty" bson:"applicability,omitempty"`                       // An expression that returns true or false, indicating whether the indication is applicable or not, after having applied its other elements
	Management              []CodeableConcept                             `json:"management,omitempty" bson:"management,omitempty"`                             // Actions for managing the contraindication
	OtherTherapy            []ClinicalUseDefinitionIndicationOtherTherapy `json:"otherTherapy,omitempty" bson:"other_therapy,omitempty"`                        // Information about use of the product in relation to other therapies described as part of the contraindication
}

func (r *ClinicalUseDefinitionContraindication) Validate() error {
	if r.DiseaseSymptomProcedure != nil {
		if err := r.DiseaseSymptomProcedure.Validate(); err != nil {
			return fmt.Errorf("DiseaseSymptomProcedure: %w", err)
		}
	}
	if r.DiseaseStatus != nil {
		if err := r.DiseaseStatus.Validate(); err != nil {
			return fmt.Errorf("DiseaseStatus: %w", err)
		}
	}
	for i, item := range r.Comorbidity {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Comorbidity[%d]: %w", i, err)
		}
	}
	for i, item := range r.Indication {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Indication[%d]: %w", i, err)
		}
	}
	if r.Applicability != nil {
		if err := r.Applicability.Validate(); err != nil {
			return fmt.Errorf("Applicability: %w", err)
		}
	}
	for i, item := range r.Management {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Management[%d]: %w", i, err)
		}
	}
	for i, item := range r.OtherTherapy {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("OtherTherapy[%d]: %w", i, err)
		}
	}
	return nil
}

type ClinicalUseDefinitionInteraction struct {
	Id          *string                                       `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Interactant []ClinicalUseDefinitionInteractionInteractant `json:"interactant,omitempty" bson:"interactant,omitempty"` // The specific medication, product, food etc. or laboratory test that interacts
	Type        *CodeableConcept                              `json:"type,omitempty" bson:"type,omitempty"`               // The type of the interaction e.g. drug-drug interaction, drug-lab test interaction
	Effect      *CodeableReference                            `json:"effect,omitempty" bson:"effect,omitempty"`           // The effect of the interaction, for example "reduced gastric absorption of primary medication"
	Incidence   *CodeableConcept                              `json:"incidence,omitempty" bson:"incidence,omitempty"`     // The incidence of the interaction, e.g. theoretical, observed
	Management  []CodeableConcept                             `json:"management,omitempty" bson:"management,omitempty"`   // Actions for managing the interaction
	Severity    *CodeableConcept                              `json:"severity,omitempty" bson:"severity,omitempty"`       // The severity of the interaction
}

func (r *ClinicalUseDefinitionInteraction) Validate() error {
	for i, item := range r.Interactant {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Interactant[%d]: %w", i, err)
		}
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Effect != nil {
		if err := r.Effect.Validate(); err != nil {
			return fmt.Errorf("Effect: %w", err)
		}
	}
	if r.Incidence != nil {
		if err := r.Incidence.Validate(); err != nil {
			return fmt.Errorf("Incidence: %w", err)
		}
	}
	for i, item := range r.Management {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Management[%d]: %w", i, err)
		}
	}
	if r.Severity != nil {
		if err := r.Severity.Validate(); err != nil {
			return fmt.Errorf("Severity: %w", err)
		}
	}
	return nil
}

type ClinicalUseDefinitionInteractionInteractant struct {
	Id                  *string          `json:"id,omitempty" bson:"id,omitempty"`                 // Unique id for inter-element referencing
	ItemReference       *Reference       `json:"itemReference" bson:"item_reference"`              // The specific medication, product, food etc. or laboratory test that interacts
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept" bson:"item_codeable_concept"` // The specific medication, product, food etc. or laboratory test that interacts
	Route               *CodeableConcept `json:"route,omitempty" bson:"route,omitempty"`           // The route by which the item is administered to cause the interaction
}

func (r *ClinicalUseDefinitionInteractionInteractant) Validate() error {
	if r.ItemReference == nil {
		return fmt.Errorf("field 'ItemReference' is required")
	}
	if r.ItemReference != nil {
		if err := r.ItemReference.Validate(); err != nil {
			return fmt.Errorf("ItemReference: %w", err)
		}
	}
	if r.ItemCodeableConcept == nil {
		return fmt.Errorf("field 'ItemCodeableConcept' is required")
	}
	if r.ItemCodeableConcept != nil {
		if err := r.ItemCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("ItemCodeableConcept: %w", err)
		}
	}
	if r.Route != nil {
		if err := r.Route.Validate(); err != nil {
			return fmt.Errorf("Route: %w", err)
		}
	}
	return nil
}

type ClinicalUseDefinitionIndicationOtherTherapy struct {
	Id               *string            `json:"id,omitempty" bson:"id,omitempty"`          // Unique id for inter-element referencing
	RelationshipType *CodeableConcept   `json:"relationshipType" bson:"relationship_type"` // The type of relationship between the product indication/contraindication and another therapy
	Treatment        *CodeableReference `json:"treatment" bson:"treatment"`                // Reference to a specific medication, substance etc. as part of an indication or contraindication
}

func (r *ClinicalUseDefinitionIndicationOtherTherapy) Validate() error {
	if r.RelationshipType == nil {
		return fmt.Errorf("field 'RelationshipType' is required")
	}
	if r.RelationshipType != nil {
		if err := r.RelationshipType.Validate(); err != nil {
			return fmt.Errorf("RelationshipType: %w", err)
		}
	}
	if r.Treatment == nil {
		return fmt.Errorf("field 'Treatment' is required")
	}
	if r.Treatment != nil {
		if err := r.Treatment.Validate(); err != nil {
			return fmt.Errorf("Treatment: %w", err)
		}
	}
	return nil
}

type ClinicalUseDefinitionWarning struct {
	Id          *string          `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Description *string          `json:"description,omitempty" bson:"description,omitempty"` // A textual definition of this warning, with formatting
	Code        *CodeableConcept `json:"code,omitempty" bson:"code,omitempty"`               // A coded or unformatted textual definition of this warning
}

func (r *ClinicalUseDefinitionWarning) Validate() error {
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	return nil
}
