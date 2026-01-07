package models

import (
	"encoding/json"
	"fmt"
)

// A medicinal product in the final form which is suitable for administering to a patient (after any mixing of multiple components, dissolution etc. has been performed).
type AdministrableProductDefinition struct {
	Id                    *string                                               `json:"id,omitempty" bson:"id,omitempty"`                                         // Logical id of this artifact
	Meta                  *Meta                                                 `json:"meta,omitempty" bson:"meta,omitempty"`                                     // Metadata about the resource
	ImplicitRules         *string                                               `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                  // A set of rules under which this content was created
	Language              *string                                               `json:"language,omitempty" bson:"language,omitempty"`                             // Language of the resource content
	Text                  *Narrative                                            `json:"text,omitempty" bson:"text,omitempty"`                                     // Text summary of the resource, for human interpretation
	Contained             []json.RawMessage                                     `json:"contained,omitempty" bson:"contained,omitempty"`                           // Contained, inline Resources
	Identifier            []Identifier                                          `json:"identifier,omitempty" bson:"identifier,omitempty"`                         // An identifier for the administrable product instance
	Status                string                                                `json:"status" bson:"status"`                                                     // draft | active | retired | unknown
	FormOf                []Reference                                           `json:"formOf,omitempty" bson:"form_of,omitempty"`                                // References a product from which one or more of the constituent parts of that product can be prepared and used as described by this administrable product
	AdministrableDoseForm *CodeableConcept                                      `json:"administrableDoseForm,omitempty" bson:"administrable_dose_form,omitempty"` // The dose form of the final product after necessary reconstitution or processing
	UnitOfPresentation    *CodeableConcept                                      `json:"unitOfPresentation,omitempty" bson:"unit_of_presentation,omitempty"`       // The presentation type in which this item is given to a patient. e.g. for a spray - 'puff'
	ProducedFrom          []Reference                                           `json:"producedFrom,omitempty" bson:"produced_from,omitempty"`                    // Indicates the specific manufactured items that are part of the 'formOf' product that are used in the preparation of this specific administrable form
	Ingredient            []CodeableConcept                                     `json:"ingredient,omitempty" bson:"ingredient,omitempty"`                         // The ingredients of this administrable medicinal product. This is only needed if the ingredients are not specified either using ManufacturedItemDefinition, or using incoming references from the Ingredient resource
	Device                *Reference                                            `json:"device,omitempty" bson:"device,omitempty"`                                 // A device that is integral to the medicinal product, in effect being considered as an "ingredient" of the medicinal product
	Description           *string                                               `json:"description,omitempty" bson:"description,omitempty"`                       // A general description of the product, when in its final form, suitable for administration e.g. effervescent blue liquid, to be swallowed
	Code                  []Coding                                              `json:"code,omitempty" bson:"code,omitempty"`                                     // A code that this product is known by, within some formal terminology. May be a PhPID
	Property              []AdministrableProductDefinitionProperty              `json:"property,omitempty" bson:"property,omitempty"`                             // Characteristics e.g. a product's onset of action
	RouteOfAdministration []AdministrableProductDefinitionRouteOfAdministration `json:"routeOfAdministration" bson:"route_of_administration"`                     // The path by which the product is taken into or makes contact with the body
}

func (r *AdministrableProductDefinition) Validate() error {
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
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	for i, item := range r.FormOf {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("FormOf[%d]: %w", i, err)
		}
	}
	if r.AdministrableDoseForm != nil {
		if err := r.AdministrableDoseForm.Validate(); err != nil {
			return fmt.Errorf("AdministrableDoseForm: %w", err)
		}
	}
	if r.UnitOfPresentation != nil {
		if err := r.UnitOfPresentation.Validate(); err != nil {
			return fmt.Errorf("UnitOfPresentation: %w", err)
		}
	}
	for i, item := range r.ProducedFrom {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ProducedFrom[%d]: %w", i, err)
		}
	}
	for i, item := range r.Ingredient {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Ingredient[%d]: %w", i, err)
		}
	}
	if r.Device != nil {
		if err := r.Device.Validate(); err != nil {
			return fmt.Errorf("Device: %w", err)
		}
	}
	for i, item := range r.Code {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Code[%d]: %w", i, err)
		}
	}
	for i, item := range r.Property {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Property[%d]: %w", i, err)
		}
	}
	if len(r.RouteOfAdministration) < 1 {
		return fmt.Errorf("field 'RouteOfAdministration' must have at least 1 elements")
	}
	for i, item := range r.RouteOfAdministration {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RouteOfAdministration[%d]: %w", i, err)
		}
	}
	return nil
}

type AdministrableProductDefinitionRouteOfAdministrationTargetSpecies struct {
	Id               *string                                                                            `json:"id,omitempty" bson:"id,omitempty"`                              // Unique id for inter-element referencing
	Code             *CodeableConcept                                                                   `json:"code" bson:"code"`                                              // Coded expression for the species
	WithdrawalPeriod []AdministrableProductDefinitionRouteOfAdministrationTargetSpeciesWithdrawalPeriod `json:"withdrawalPeriod,omitempty" bson:"withdrawal_period,omitempty"` // A species specific time during which consumption of animal product is not appropriate
}

func (r *AdministrableProductDefinitionRouteOfAdministrationTargetSpecies) Validate() error {
	if r.Code == nil {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	for i, item := range r.WithdrawalPeriod {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("WithdrawalPeriod[%d]: %w", i, err)
		}
	}
	return nil
}

type AdministrableProductDefinitionRouteOfAdministrationTargetSpeciesWithdrawalPeriod struct {
	Id                    *string          `json:"id,omitempty" bson:"id,omitempty"`                                        // Unique id for inter-element referencing
	Tissue                *CodeableConcept `json:"tissue" bson:"tissue"`                                                    // The type of tissue for which the withdrawal period applies, e.g. meat, milk
	Value                 *Quantity        `json:"value" bson:"value"`                                                      // A value for the time
	SupportingInformation *string          `json:"supportingInformation,omitempty" bson:"supporting_information,omitempty"` // Extra information about the withdrawal period
}

func (r *AdministrableProductDefinitionRouteOfAdministrationTargetSpeciesWithdrawalPeriod) Validate() error {
	if r.Tissue == nil {
		return fmt.Errorf("field 'Tissue' is required")
	}
	if r.Tissue != nil {
		if err := r.Tissue.Validate(); err != nil {
			return fmt.Errorf("Tissue: %w", err)
		}
	}
	if r.Value == nil {
		return fmt.Errorf("field 'Value' is required")
	}
	if r.Value != nil {
		if err := r.Value.Validate(); err != nil {
			return fmt.Errorf("Value: %w", err)
		}
	}
	return nil
}

type AdministrableProductDefinitionProperty struct {
	Id                   *string          `json:"id,omitempty" bson:"id,omitempty"`                                       // Unique id for inter-element referencing
	Type                 *CodeableConcept `json:"type" bson:"type"`                                                       // A code expressing the type of characteristic
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty" bson:"value_codeable_concept,omitempty"` // A value for the characteristic
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty" bson:"value_quantity,omitempty"`                // A value for the characteristic
	ValueRange           *Range           `json:"valueRange,omitempty" bson:"value_range,omitempty"`                      // A value for the characteristic
	ValueDate            *string          `json:"valueDate,omitempty" bson:"value_date,omitempty"`                        // A value for the characteristic
	ValueBoolean         *bool            `json:"valueBoolean,omitempty" bson:"value_boolean,omitempty"`                  // A value for the characteristic
	ValueMarkdown        *string          `json:"valueMarkdown,omitempty" bson:"value_markdown,omitempty"`                // A value for the characteristic
	ValueAttachment      *Attachment      `json:"valueAttachment,omitempty" bson:"value_attachment,omitempty"`            // A value for the characteristic
	ValueReference       *Reference       `json:"valueReference,omitempty" bson:"value_reference,omitempty"`              // A value for the characteristic
	Status               *CodeableConcept `json:"status,omitempty" bson:"status,omitempty"`                               // The status of characteristic e.g. assigned or pending
}

func (r *AdministrableProductDefinitionProperty) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
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
	if r.ValueAttachment != nil {
		if err := r.ValueAttachment.Validate(); err != nil {
			return fmt.Errorf("ValueAttachment: %w", err)
		}
	}
	if r.ValueReference != nil {
		if err := r.ValueReference.Validate(); err != nil {
			return fmt.Errorf("ValueReference: %w", err)
		}
	}
	if r.Status != nil {
		if err := r.Status.Validate(); err != nil {
			return fmt.Errorf("Status: %w", err)
		}
	}
	return nil
}

type AdministrableProductDefinitionRouteOfAdministration struct {
	Id                        *string                                                            `json:"id,omitempty" bson:"id,omitempty"`                                                   // Unique id for inter-element referencing
	Code                      *CodeableConcept                                                   `json:"code" bson:"code"`                                                                   // Coded expression for the route
	FirstDose                 *Quantity                                                          `json:"firstDose,omitempty" bson:"first_dose,omitempty"`                                    // The first dose (dose quantity) administered can be specified for the product
	MaxSingleDose             *Quantity                                                          `json:"maxSingleDose,omitempty" bson:"max_single_dose,omitempty"`                           // The maximum single dose that can be administered
	MaxDosePerDay             *Quantity                                                          `json:"maxDosePerDay,omitempty" bson:"max_dose_per_day,omitempty"`                          // The maximum dose quantity to be administered in any one 24-h period
	MaxDosePerTreatmentPeriod *Ratio                                                             `json:"maxDosePerTreatmentPeriod,omitempty" bson:"max_dose_per_treatment_period,omitempty"` // The maximum dose per treatment period that can be administered
	MaxTreatmentPeriod        *Duration                                                          `json:"maxTreatmentPeriod,omitempty" bson:"max_treatment_period,omitempty"`                 // The maximum treatment period during which the product can be administered
	TargetSpecies             []AdministrableProductDefinitionRouteOfAdministrationTargetSpecies `json:"targetSpecies,omitempty" bson:"target_species,omitempty"`                            // A species for which this route applies
}

func (r *AdministrableProductDefinitionRouteOfAdministration) Validate() error {
	if r.Code == nil {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.FirstDose != nil {
		if err := r.FirstDose.Validate(); err != nil {
			return fmt.Errorf("FirstDose: %w", err)
		}
	}
	if r.MaxSingleDose != nil {
		if err := r.MaxSingleDose.Validate(); err != nil {
			return fmt.Errorf("MaxSingleDose: %w", err)
		}
	}
	if r.MaxDosePerDay != nil {
		if err := r.MaxDosePerDay.Validate(); err != nil {
			return fmt.Errorf("MaxDosePerDay: %w", err)
		}
	}
	if r.MaxDosePerTreatmentPeriod != nil {
		if err := r.MaxDosePerTreatmentPeriod.Validate(); err != nil {
			return fmt.Errorf("MaxDosePerTreatmentPeriod: %w", err)
		}
	}
	if r.MaxTreatmentPeriod != nil {
		if err := r.MaxTreatmentPeriod.Validate(); err != nil {
			return fmt.Errorf("MaxTreatmentPeriod: %w", err)
		}
	}
	for i, item := range r.TargetSpecies {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("TargetSpecies[%d]: %w", i, err)
		}
	}
	return nil
}
