package models

import (
	"encoding/json"
	"fmt"
)

// A record of intake by a patient.  A NutritionIntake may indicate that the patient may be consuming the food (i.e., solid and/or liquid), breastmilk, infant formula, supplements, enteral formula now or has consumed it in the past.  The source of this information can be the patient, significant other (such as a family member or spouse), or a clinician.  A common scenario where this information is captured is during the history taking process during a patient visit or stay or through an app that tracks food (i.e., solid and/or liquid), breastmilk, infant formula, supplements, enteral formula consumed.   The consumption information may come from sources such as the patient's memory, from a nutrition label, or from a clinician documenting observed intake.
type NutritionIntake struct {
	Id                 *string                        `json:"id,omitempty" bson:"id,omitempty"`                                   // Logical id of this artifact
	Meta               *Meta                          `json:"meta,omitempty" bson:"meta,omitempty"`                               // Metadata about the resource
	ImplicitRules      *string                        `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`            // A set of rules under which this content was created
	Language           *string                        `json:"language,omitempty" bson:"language,omitempty"`                       // Language of the resource content
	Text               *Narrative                     `json:"text,omitempty" bson:"text,omitempty"`                               // Text summary of the resource, for human interpretation
	Contained          []json.RawMessage              `json:"contained,omitempty" bson:"contained,omitempty"`                     // Contained, inline Resources
	Identifier         []Identifier                   `json:"identifier,omitempty" bson:"identifier,omitempty"`                   // External identifier
	BasedOn            []Reference                    `json:"basedOn,omitempty" bson:"based_on,omitempty"`                        // Fulfils plan, proposal or order
	PartOf             []Reference                    `json:"partOf,omitempty" bson:"part_of,omitempty"`                          // Part of referenced event
	Status             string                         `json:"status" bson:"status"`                                               // preparation | in-progress | not-done | on-hold | stopped | completed | entered-in-error | unknown
	StatusReason       []CodeableConcept              `json:"statusReason,omitempty" bson:"status_reason,omitempty"`              // Reason for current status
	Code               *CodeableConcept               `json:"code,omitempty" bson:"code,omitempty"`                               // Code representing an overall type of nutrition intake
	Subject            *Reference                     `json:"subject" bson:"subject"`                                             // Who is/was consuming the food (i.e. solid and/or liquid)
	Encounter          *Reference                     `json:"encounter,omitempty" bson:"encounter,omitempty"`                     // Encounter associated with NutritionIntake
	OccurrenceDateTime *string                        `json:"occurrenceDateTime,omitempty" bson:"occurrence_date_time,omitempty"` // The date/time or interval when the food (i.e. solid and/or liquid) is/was consumed
	OccurrencePeriod   *Period                        `json:"occurrencePeriod,omitempty" bson:"occurrence_period,omitempty"`      // The date/time or interval when the food (i.e. solid and/or liquid) is/was consumed
	Recorded           *string                        `json:"recorded,omitempty" bson:"recorded,omitempty"`                       // When the intake was recorded
	ReportedBoolean    *bool                          `json:"reportedBoolean,omitempty" bson:"reported_boolean,omitempty"`        // Indicates if this is a reported rather than a primary record.  Can also indicate the source that provided the information about the consumption
	ReportedReference  *Reference                     `json:"reportedReference,omitempty" bson:"reported_reference,omitempty"`    // Indicates if this is a reported rather than a primary record.  Can also indicate the source that provided the information about the consumption
	NutritionItem      []NutritionIntakeNutritionItem `json:"nutritionItem,omitempty" bson:"nutrition_item,omitempty"`            // The nutrition product intended for consumption and/or administration
	Performer          []NutritionIntakePerformer     `json:"performer,omitempty" bson:"performer,omitempty"`                     // Who or what performed the intake and how they were involved
	Location           *Reference                     `json:"location,omitempty" bson:"location,omitempty"`                       // Where the intake occurred
	DerivedFrom        []Reference                    `json:"derivedFrom,omitempty" bson:"derived_from,omitempty"`                // Additional supporting information
	Reason             []CodeableReference            `json:"reason,omitempty" bson:"reason,omitempty"`                           // Reason for why the food (i.e. solid and/or liquid) is /was consumed
	Note               []Annotation                   `json:"note,omitempty" bson:"note,omitempty"`                               // Further information about the consumption
}

func (r *NutritionIntake) Validate() error {
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
	for i, item := range r.PartOf {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("PartOf[%d]: %w", i, err)
		}
	}
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	for i, item := range r.StatusReason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("StatusReason[%d]: %w", i, err)
		}
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
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
	if r.OccurrencePeriod != nil {
		if err := r.OccurrencePeriod.Validate(); err != nil {
			return fmt.Errorf("OccurrencePeriod: %w", err)
		}
	}
	if r.ReportedReference != nil {
		if err := r.ReportedReference.Validate(); err != nil {
			return fmt.Errorf("ReportedReference: %w", err)
		}
	}
	for i, item := range r.NutritionItem {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("NutritionItem[%d]: %w", i, err)
		}
	}
	for i, item := range r.Performer {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Performer[%d]: %w", i, err)
		}
	}
	if r.Location != nil {
		if err := r.Location.Validate(); err != nil {
			return fmt.Errorf("Location: %w", err)
		}
	}
	for i, item := range r.DerivedFrom {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("DerivedFrom[%d]: %w", i, err)
		}
	}
	for i, item := range r.Reason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reason[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	return nil
}

type NutritionIntakeNutritionItem struct {
	Id               *string                                       `json:"id,omitempty" bson:"id,omitempty"`                              // Unique id for inter-element referencing
	Type             *CodeableConcept                              `json:"type,omitempty" bson:"type,omitempty"`                          // The type of food (i.e. solid or liquid) product
	NutritionProduct *CodeableReference                            `json:"nutritionProduct,omitempty" bson:"nutrition_product,omitempty"` // A product used for nutritional purposes (e.g. food or supplement)
	ConsumedItem     []NutritionIntakeNutritionItemConsumedItem    `json:"consumedItem,omitempty" bson:"consumed_item,omitempty"`         // What nutrition item was consumed
	NotConsumedItem  []NutritionIntakeNutritionItemNotConsumedItem `json:"notConsumedItem,omitempty" bson:"not_consumed_item,omitempty"`  // What nutrition item was not consumed
}

func (r *NutritionIntakeNutritionItem) Validate() error {
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.NutritionProduct != nil {
		if err := r.NutritionProduct.Validate(); err != nil {
			return fmt.Errorf("NutritionProduct: %w", err)
		}
	}
	for i, item := range r.ConsumedItem {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ConsumedItem[%d]: %w", i, err)
		}
	}
	for i, item := range r.NotConsumedItem {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("NotConsumedItem[%d]: %w", i, err)
		}
	}
	return nil
}

type NutritionIntakeNutritionItemConsumedItem struct {
	Id           *string                                               `json:"id,omitempty" bson:"id,omitempty"`                      // Unique id for inter-element referencing
	Schedule     *Timing                                               `json:"schedule,omitempty" bson:"schedule,omitempty"`          // Scheduled frequency of consumption
	Amount       *Quantity                                             `json:"amount,omitempty" bson:"amount,omitempty"`              // Quantity of the specified food (i.e. solid and/or liquid)
	RateQuantity *Quantity                                             `json:"rateQuantity,omitempty" bson:"rate_quantity,omitempty"` // Rate of enteral feeding administration
	RateRatio    *Ratio                                                `json:"rateRatio,omitempty" bson:"rate_ratio,omitempty"`       // Rate of enteral feeding administration
	TotalIntake  []NutritionIntakeNutritionItemConsumedItemTotalIntake `json:"totalIntake,omitempty" bson:"total_intake,omitempty"`   // Nutrients and/or energy contained in the intake
}

func (r *NutritionIntakeNutritionItemConsumedItem) Validate() error {
	if r.Schedule != nil {
		if err := r.Schedule.Validate(); err != nil {
			return fmt.Errorf("Schedule: %w", err)
		}
	}
	if r.Amount != nil {
		if err := r.Amount.Validate(); err != nil {
			return fmt.Errorf("Amount: %w", err)
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
	for i, item := range r.TotalIntake {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("TotalIntake[%d]: %w", i, err)
		}
	}
	return nil
}

type NutritionIntakeNutritionItemConsumedItemTotalIntake struct {
	Id       *string            `json:"id,omitempty" bson:"id,omitempty"`         // Unique id for inter-element referencing
	Nutrient *CodeableReference `json:"nutrient" bson:"nutrient"`                 // Type of nutrient consumed in the intake
	Amount   *Quantity          `json:"amount" bson:"amount"`                     // Total amount of nutrient consumed
	Energy   *Quantity          `json:"energy,omitempty" bson:"energy,omitempty"` // Total energy consumed in kilocalories or kilojoules
}

func (r *NutritionIntakeNutritionItemConsumedItemTotalIntake) Validate() error {
	if r.Nutrient == nil {
		return fmt.Errorf("field 'Nutrient' is required")
	}
	if r.Nutrient != nil {
		if err := r.Nutrient.Validate(); err != nil {
			return fmt.Errorf("Nutrient: %w", err)
		}
	}
	if r.Amount == nil {
		return fmt.Errorf("field 'Amount' is required")
	}
	if r.Amount != nil {
		if err := r.Amount.Validate(); err != nil {
			return fmt.Errorf("Amount: %w", err)
		}
	}
	if r.Energy != nil {
		if err := r.Energy.Validate(); err != nil {
			return fmt.Errorf("Energy: %w", err)
		}
	}
	return nil
}

type NutritionIntakeNutritionItemNotConsumedItem struct {
	Id       *string          `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Reason   *CodeableConcept `json:"reason,omitempty" bson:"reason,omitempty"`     // Reason the nutrition item was not consumed
	Schedule *Timing          `json:"schedule,omitempty" bson:"schedule,omitempty"` // The intended frequency of consumption that was not followed
	Amount   *Quantity        `json:"amount,omitempty" bson:"amount,omitempty"`     // Quantity of the specified food (i.e. solid and/or liquid) that was not consumed
}

func (r *NutritionIntakeNutritionItemNotConsumedItem) Validate() error {
	if r.Reason != nil {
		if err := r.Reason.Validate(); err != nil {
			return fmt.Errorf("Reason: %w", err)
		}
	}
	if r.Schedule != nil {
		if err := r.Schedule.Validate(); err != nil {
			return fmt.Errorf("Schedule: %w", err)
		}
	}
	if r.Amount != nil {
		if err := r.Amount.Validate(); err != nil {
			return fmt.Errorf("Amount: %w", err)
		}
	}
	return nil
}

type NutritionIntakePerformer struct {
	Id       *string          `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Function *CodeableConcept `json:"function,omitempty" bson:"function,omitempty"` // Type of performer
	Actor    *Reference       `json:"actor" bson:"actor"`                           // Who or what performed the intake
}

func (r *NutritionIntakePerformer) Validate() error {
	if r.Function != nil {
		if err := r.Function.Validate(); err != nil {
			return fmt.Errorf("Function: %w", err)
		}
	}
	if r.Actor == nil {
		return fmt.Errorf("field 'Actor' is required")
	}
	if r.Actor != nil {
		if err := r.Actor.Validate(); err != nil {
			return fmt.Errorf("Actor: %w", err)
		}
	}
	return nil
}
