package models

import (
	"encoding/json"
	"fmt"
)

// A request to supply a diet, formula feeding (enteral) or oral nutritional supplement to an individual or group.
type NutritionOrder struct {
	Id                     *string                       `json:"id,omitempty" bson:"id,omitempty"`                                           // Logical id of this artifact
	Meta                   *Meta                         `json:"meta,omitempty" bson:"meta,omitempty"`                                       // Metadata about the resource
	ImplicitRules          *string                       `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                    // A set of rules under which this content was created
	Language               *string                       `json:"language,omitempty" bson:"language,omitempty"`                               // Language of the resource content
	Text                   *Narrative                    `json:"text,omitempty" bson:"text,omitempty"`                                       // Text summary of the resource, for human interpretation
	Contained              []json.RawMessage             `json:"contained,omitempty" bson:"contained,omitempty"`                             // Contained, inline Resources
	Identifier             []Identifier                  `json:"identifier,omitempty" bson:"identifier,omitempty"`                           // Identifiers assigned to this order
	BasedOn                []Reference                   `json:"basedOn,omitempty" bson:"based_on,omitempty"`                                // What this order fulfills
	GroupIdentifier        *Identifier                   `json:"groupIdentifier,omitempty" bson:"group_identifier,omitempty"`                // Composite Request ID
	Status                 string                        `json:"status" bson:"status"`                                                       // draft | active | on-hold | entered-in-error | ended | completed | revoked | unknown
	Intent                 string                        `json:"intent" bson:"intent"`                                                       // proposal | solicit-offer | offer-response | plan | directive | order | original-order | reflex-order | filler-order | instance-order | option
	Priority               *string                       `json:"priority,omitempty" bson:"priority,omitempty"`                               // routine | urgent | asap | stat
	Subject                *Reference                    `json:"subject" bson:"subject"`                                                     // Who requires the diet, formula or nutritional supplement
	Encounter              *Reference                    `json:"encounter,omitempty" bson:"encounter,omitempty"`                             // The encounter associated with this nutrition order
	SupportingInformation  []Reference                   `json:"supportingInformation,omitempty" bson:"supporting_information,omitempty"`    // Information to support fulfilling of the nutrition order
	DateTime               string                        `json:"dateTime" bson:"date_time"`                                                  // Date and time the nutrition order was requested
	Requester              *Reference                    `json:"requester,omitempty" bson:"requester,omitempty"`                             // Who ordered the diet, formula or nutritional supplement
	Performer              []CodeableReference           `json:"performer,omitempty" bson:"performer,omitempty"`                             // Who is intended to perform the administration of the nutrition order
	AllergyIntolerance     []Reference                   `json:"allergyIntolerance,omitempty" bson:"allergy_intolerance,omitempty"`          // List of the patient's food and nutrition-related allergies and intolerances
	FoodPreferenceModifier []CodeableConcept             `json:"foodPreferenceModifier,omitempty" bson:"food_preference_modifier,omitempty"` // Order-specific modifier about the type of food that should be given
	ExcludeFoodModifier    []CodeableConcept             `json:"excludeFoodModifier,omitempty" bson:"exclude_food_modifier,omitempty"`       // Food that should not be given
	OutsideFoodAllowed     bool                          `json:"outsideFoodAllowed,omitempty" bson:"outside_food_allowed,omitempty"`         // Capture if patient is permitted to consume food from outside of current setting brought by the patient, family, and/or caregiver
	OralDiet               *NutritionOrderOralDiet       `json:"oralDiet,omitempty" bson:"oral_diet,omitempty"`                              // Oral diet components
	Supplement             []NutritionOrderSupplement    `json:"supplement,omitempty" bson:"supplement,omitempty"`                           // Supplement components
	EnteralFormula         *NutritionOrderEnteralFormula `json:"enteralFormula,omitempty" bson:"enteral_formula,omitempty"`                  // Enteral formula product
	Additive               []NutritionOrderAdditive      `json:"additive,omitempty" bson:"additive,omitempty"`                               // Modular additive to add to the oral diet, supplement, and/or enteral feeding
	Note                   []Annotation                  `json:"note,omitempty" bson:"note,omitempty"`                                       // Comments
}

func (r *NutritionOrder) Validate() error {
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
	if r.GroupIdentifier != nil {
		if err := r.GroupIdentifier.Validate(); err != nil {
			return fmt.Errorf("GroupIdentifier: %w", err)
		}
	}
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.Intent == emptyString {
		return fmt.Errorf("field 'Intent' is required")
	}
	if r.Subject == nil {
		return fmt.Errorf("field 'Subject' is required")
	}
	if r.Subject != nil {
		if err := r.Subject.Validate(); err != nil {
			return fmt.Errorf("Subject: %w", err)
		}
	}
	if r.Encounter != nil {
		if err := r.Encounter.Validate(); err != nil {
			return fmt.Errorf("Encounter: %w", err)
		}
	}
	for i, item := range r.SupportingInformation {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SupportingInformation[%d]: %w", i, err)
		}
	}
	if r.DateTime == emptyString {
		return fmt.Errorf("field 'DateTime' is required")
	}
	if r.Requester != nil {
		if err := r.Requester.Validate(); err != nil {
			return fmt.Errorf("Requester: %w", err)
		}
	}
	for i, item := range r.Performer {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Performer[%d]: %w", i, err)
		}
	}
	for i, item := range r.AllergyIntolerance {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AllergyIntolerance[%d]: %w", i, err)
		}
	}
	for i, item := range r.FoodPreferenceModifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("FoodPreferenceModifier[%d]: %w", i, err)
		}
	}
	for i, item := range r.ExcludeFoodModifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ExcludeFoodModifier[%d]: %w", i, err)
		}
	}
	if r.OralDiet != nil {
		if err := r.OralDiet.Validate(); err != nil {
			return fmt.Errorf("OralDiet: %w", err)
		}
	}
	for i, item := range r.Supplement {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Supplement[%d]: %w", i, err)
		}
	}
	if r.EnteralFormula != nil {
		if err := r.EnteralFormula.Validate(); err != nil {
			return fmt.Errorf("EnteralFormula: %w", err)
		}
	}
	for i, item := range r.Additive {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Additive[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	return nil
}

type NutritionOrderEnteralFormulaAdministration struct {
	Id           *string                                             `json:"id,omitempty" bson:"id,omitempty"`                      // Unique id for inter-element referencing
	Schedule     *NutritionOrderEnteralFormulaAdministrationSchedule `json:"schedule,omitempty" bson:"schedule,omitempty"`          // Scheduling information for enteral feeding products
	Quantity     *Quantity                                           `json:"quantity,omitempty" bson:"quantity,omitempty"`          // The volume of formula feeding to provide
	RateQuantity *Quantity                                           `json:"rateQuantity,omitempty" bson:"rate_quantity,omitempty"` // Speed with which the formula feeding is provided per period of time
	RateRatio    *Ratio                                              `json:"rateRatio,omitempty" bson:"rate_ratio,omitempty"`       // Speed with which the formula feeding is provided per period of time
}

func (r *NutritionOrderEnteralFormulaAdministration) Validate() error {
	if r.Schedule != nil {
		if err := r.Schedule.Validate(); err != nil {
			return fmt.Errorf("Schedule: %w", err)
		}
	}
	if r.Quantity != nil {
		if err := r.Quantity.Validate(); err != nil {
			return fmt.Errorf("Quantity: %w", err)
		}
	}
	if r.RateQuantity != nil {
		if err := r.RateQuantity.Validate(); err != nil {
			return fmt.Errorf("RateQuantity: %w", err)
		}
	}
	if r.RateRatio != nil {
		if err := r.RateRatio.Validate(); err != nil {
			return fmt.Errorf("RateRatio: %w", err)
		}
	}
	return nil
}

type NutritionOrderEnteralFormulaAdministrationSchedule struct {
	Id          *string          `json:"id,omitempty" bson:"id,omitempty"`                     // Unique id for inter-element referencing
	Timing      []Timing         `json:"timing,omitempty" bson:"timing,omitempty"`             // Scheduled frequency of enteral feeding
	AsNeeded    bool             `json:"asNeeded,omitempty" bson:"as_needed,omitempty"`        // Take 'as needed'
	AsNeededFor *CodeableConcept `json:"asNeededFor,omitempty" bson:"as_needed_for,omitempty"` // Take 'as needed' for x
}

func (r *NutritionOrderEnteralFormulaAdministrationSchedule) Validate() error {
	for i, item := range r.Timing {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Timing[%d]: %w", i, err)
		}
	}
	if r.AsNeededFor != nil {
		if err := r.AsNeededFor.Validate(); err != nil {
			return fmt.Errorf("AsNeededFor: %w", err)
		}
	}
	return nil
}

type NutritionOrderAdditive struct {
	Id                    *string            `json:"id,omitempty" bson:"id,omitempty"`                                         // Unique id for inter-element referencing
	ModularType           *CodeableReference `json:"modularType,omitempty" bson:"modular_type,omitempty"`                      // Type of modular component to add to the oral diet, supplement, and/or enteral feeding
	ProductName           *string            `json:"productName,omitempty" bson:"product_name,omitempty"`                      // Product or brand name of the modular additive
	Quantity              *Quantity          `json:"quantity,omitempty" bson:"quantity,omitempty"`                             // Amount of additive to be given or mixed in with the oral diet, supplement, and/or enteral feeding
	RouteOfAdministration []CodeableConcept  `json:"routeOfAdministration,omitempty" bson:"route_of_administration,omitempty"` // How the additive should enter the patient's gastrointestinal tract
}

func (r *NutritionOrderAdditive) Validate() error {
	if r.ModularType != nil {
		if err := r.ModularType.Validate(); err != nil {
			return fmt.Errorf("ModularType: %w", err)
		}
	}
	if r.Quantity != nil {
		if err := r.Quantity.Validate(); err != nil {
			return fmt.Errorf("Quantity: %w", err)
		}
	}
	for i, item := range r.RouteOfAdministration {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RouteOfAdministration[%d]: %w", i, err)
		}
	}
	return nil
}

type NutritionOrderOralDiet struct {
	Id             *string                          `json:"id,omitempty" bson:"id,omitempty"`                          // Unique id for inter-element referencing
	Type           []CodeableConcept                `json:"type,omitempty" bson:"type,omitempty"`                      // Type of oral diet or diet restrictions that can be consumed orally
	Schedule       *NutritionOrderOralDietSchedule  `json:"schedule,omitempty" bson:"schedule,omitempty"`              // Scheduling information for oral diets
	Nutrient       []NutritionOrderOralDietNutrient `json:"nutrient,omitempty" bson:"nutrient,omitempty"`              // The nutrient that is modified and the quantity in the diet
	Texture        []NutritionOrderOralDietTexture  `json:"texture,omitempty" bson:"texture,omitempty"`                // Texture modifications in addition to the oral diet type
	Instruction    *string                          `json:"instruction,omitempty" bson:"instruction,omitempty"`        // Instructions or additional information about the oral diet
	CaloricDensity *Quantity                        `json:"caloricDensity,omitempty" bson:"caloric_density,omitempty"` // Amount of energy per specified volume of oral diet
}

func (r *NutritionOrderOralDiet) Validate() error {
	for i, item := range r.Type {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Type[%d]: %w", i, err)
		}
	}
	if r.Schedule != nil {
		if err := r.Schedule.Validate(); err != nil {
			return fmt.Errorf("Schedule: %w", err)
		}
	}
	for i, item := range r.Nutrient {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Nutrient[%d]: %w", i, err)
		}
	}
	for i, item := range r.Texture {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Texture[%d]: %w", i, err)
		}
	}
	if r.CaloricDensity != nil {
		if err := r.CaloricDensity.Validate(); err != nil {
			return fmt.Errorf("CaloricDensity: %w", err)
		}
	}
	return nil
}

type NutritionOrderOralDietSchedule struct {
	Id          *string          `json:"id,omitempty" bson:"id,omitempty"`                     // Unique id for inter-element referencing
	Timing      []Timing         `json:"timing,omitempty" bson:"timing,omitempty"`             // Scheduled frequency of diet
	AsNeeded    bool             `json:"asNeeded,omitempty" bson:"as_needed,omitempty"`        // Take 'as needed'
	AsNeededFor *CodeableConcept `json:"asNeededFor,omitempty" bson:"as_needed_for,omitempty"` // Take 'as needed' for x
}

func (r *NutritionOrderOralDietSchedule) Validate() error {
	for i, item := range r.Timing {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Timing[%d]: %w", i, err)
		}
	}
	if r.AsNeededFor != nil {
		if err := r.AsNeededFor.Validate(); err != nil {
			return fmt.Errorf("AsNeededFor: %w", err)
		}
	}
	return nil
}

type NutritionOrderOralDietTexture struct {
	Id       *string          `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Modifier *CodeableConcept `json:"modifier,omitempty" bson:"modifier,omitempty"` // Food (i.e. solid and/or liquid) texture modifications in addition to those in the oral diet type
	Type     *CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`         // Food (i.e. solid and/or liquid) types that undergo texture alteration
}

func (r *NutritionOrderOralDietTexture) Validate() error {
	if r.Modifier != nil {
		if err := r.Modifier.Validate(); err != nil {
			return fmt.Errorf("Modifier: %w", err)
		}
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	return nil
}

type NutritionOrderSupplement struct {
	Id             *string                           `json:"id,omitempty" bson:"id,omitempty"`                          // Unique id for inter-element referencing
	Type           *CodeableReference                `json:"type,omitempty" bson:"type,omitempty"`                      // Type of supplement product requested
	ProductName    *string                           `json:"productName,omitempty" bson:"product_name,omitempty"`       // Product or brand name of the nutritional supplement
	Schedule       *NutritionOrderSupplementSchedule `json:"schedule,omitempty" bson:"schedule,omitempty"`              // Scheduling information for supplements
	Quantity       *Quantity                         `json:"quantity,omitempty" bson:"quantity,omitempty"`              // Amount of the nutritional supplement
	Instruction    *string                           `json:"instruction,omitempty" bson:"instruction,omitempty"`        // Instructions or additional information about the oral supplement
	CaloricDensity *Quantity                         `json:"caloricDensity,omitempty" bson:"caloric_density,omitempty"` // Amount of energy per specified volume of supplement that is required
}

func (r *NutritionOrderSupplement) Validate() error {
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Schedule != nil {
		if err := r.Schedule.Validate(); err != nil {
			return fmt.Errorf("Schedule: %w", err)
		}
	}
	if r.Quantity != nil {
		if err := r.Quantity.Validate(); err != nil {
			return fmt.Errorf("Quantity: %w", err)
		}
	}
	if r.CaloricDensity != nil {
		if err := r.CaloricDensity.Validate(); err != nil {
			return fmt.Errorf("CaloricDensity: %w", err)
		}
	}
	return nil
}

type NutritionOrderOralDietNutrient struct {
	Id       *string          `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Modifier *CodeableConcept `json:"modifier,omitempty" bson:"modifier,omitempty"` // Nutrient modified in the oral diet type
	Amount   *Quantity        `json:"amount,omitempty" bson:"amount,omitempty"`     // Quantity of the specified nutrient
}

func (r *NutritionOrderOralDietNutrient) Validate() error {
	if r.Modifier != nil {
		if err := r.Modifier.Validate(); err != nil {
			return fmt.Errorf("Modifier: %w", err)
		}
	}
	if r.Amount != nil {
		if err := r.Amount.Validate(); err != nil {
			return fmt.Errorf("Amount: %w", err)
		}
	}
	return nil
}

type NutritionOrderSupplementSchedule struct {
	Id          *string          `json:"id,omitempty" bson:"id,omitempty"`                     // Unique id for inter-element referencing
	Timing      []Timing         `json:"timing,omitempty" bson:"timing,omitempty"`             // Scheduled frequency of supplement
	AsNeeded    bool             `json:"asNeeded,omitempty" bson:"as_needed,omitempty"`        // Take 'as needed'
	AsNeededFor *CodeableConcept `json:"asNeededFor,omitempty" bson:"as_needed_for,omitempty"` // Take 'as needed' for x
}

func (r *NutritionOrderSupplementSchedule) Validate() error {
	for i, item := range r.Timing {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Timing[%d]: %w", i, err)
		}
	}
	if r.AsNeededFor != nil {
		if err := r.AsNeededFor.Validate(); err != nil {
			return fmt.Errorf("AsNeededFor: %w", err)
		}
	}
	return nil
}

type NutritionOrderEnteralFormula struct {
	Id                            *string                                      `json:"id,omitempty" bson:"id,omitempty"`                                                          // Unique id for inter-element referencing
	Type                          *CodeableReference                           `json:"type,omitempty" bson:"type,omitempty"`                                                      // Type of patient enteral feeding
	ProductName                   *string                                      `json:"productName,omitempty" bson:"product_name,omitempty"`                                       // Product or brand name of the enteral feeding
	DeliveryDeviceCodeableConcept *CodeableConcept                             `json:"deliveryDeviceCodeableConcept,omitempty" bson:"delivery_device_codeable_concept,omitempty"` // Intended type of device for the enteral feeding administration
	DeliveryDeviceCanonical       *string                                      `json:"deliveryDeviceCanonical,omitempty" bson:"delivery_device_canonical,omitempty"`              // Intended type of device for the enteral feeding administration
	CaloricDensity                *Quantity                                    `json:"caloricDensity,omitempty" bson:"caloric_density,omitempty"`                                 // Amount of energy per specified volume of feeding that is required
	RouteOfAdministration         []CodeableConcept                            `json:"routeOfAdministration,omitempty" bson:"route_of_administration,omitempty"`                  // How the enteral feeding should enter the patient's gastrointestinal tract
	Administration                []NutritionOrderEnteralFormulaAdministration `json:"administration,omitempty" bson:"administration,omitempty"`                                  // Formula feeding instruction as structured data
	MaxVolumeToAdminister         *Quantity                                    `json:"maxVolumeToAdminister,omitempty" bson:"max_volume_to_administer,omitempty"`                 // Upper limit on formula feeding volume per unit of time
	AdministrationInstruction     *string                                      `json:"administrationInstruction,omitempty" bson:"administration_instruction,omitempty"`           // Formula feeding instructions expressed as text
}

func (r *NutritionOrderEnteralFormula) Validate() error {
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.DeliveryDeviceCodeableConcept != nil {
		if err := r.DeliveryDeviceCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("DeliveryDeviceCodeableConcept: %w", err)
		}
	}
	if r.CaloricDensity != nil {
		if err := r.CaloricDensity.Validate(); err != nil {
			return fmt.Errorf("CaloricDensity: %w", err)
		}
	}
	for i, item := range r.RouteOfAdministration {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RouteOfAdministration[%d]: %w", i, err)
		}
	}
	for i, item := range r.Administration {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Administration[%d]: %w", i, err)
		}
	}
	if r.MaxVolumeToAdminister != nil {
		if err := r.MaxVolumeToAdminister.Validate(); err != nil {
			return fmt.Errorf("MaxVolumeToAdminister: %w", err)
		}
	}
	return nil
}
