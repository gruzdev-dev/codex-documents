package models

import (
	"encoding/json"
	"fmt"
)

// Describes the event of a patient being administered a vaccine or a record of an immunization as reported by a patient, a clinician or another party.
type Immunization struct {
	Id                    *string                          `json:"id,omitempty" bson:"id,omitempty"`                                        // Logical id of this artifact
	Meta                  *Meta                            `json:"meta,omitempty" bson:"meta,omitempty"`                                    // Metadata about the resource
	ImplicitRules         *string                          `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                 // A set of rules under which this content was created
	Language              *string                          `json:"language,omitempty" bson:"language,omitempty"`                            // Language of the resource content
	Text                  *Narrative                       `json:"text,omitempty" bson:"text,omitempty"`                                    // Text summary of the resource, for human interpretation
	Contained             []json.RawMessage                `json:"contained,omitempty" bson:"contained,omitempty"`                          // Contained, inline Resources
	Identifier            []Identifier                     `json:"identifier,omitempty" bson:"identifier,omitempty"`                        // Business identifier
	BasedOn               []Reference                      `json:"basedOn,omitempty" bson:"based_on,omitempty"`                             // Authority that the immunization event is based on
	Status                string                           `json:"status" bson:"status"`                                                    // completed | entered-in-error | not-done
	StatusReason          *CodeableConcept                 `json:"statusReason,omitempty" bson:"status_reason,omitempty"`                   // Reason for current status
	VaccineCode           *CodeableConcept                 `json:"vaccineCode" bson:"vaccine_code"`                                         // Vaccine administered
	AdministeredProduct   *CodeableReference               `json:"administeredProduct,omitempty" bson:"administered_product,omitempty"`     // Product that was administered
	Manufacturer          *CodeableReference               `json:"manufacturer,omitempty" bson:"manufacturer,omitempty"`                    // Vaccine manufacturer
	LotNumber             *string                          `json:"lotNumber,omitempty" bson:"lot_number,omitempty"`                         // Vaccine lot number
	ExpirationDate        *string                          `json:"expirationDate,omitempty" bson:"expiration_date,omitempty"`               // Vaccine expiration date
	Patient               *Reference                       `json:"patient" bson:"patient"`                                                  // Who was immunized
	Encounter             *Reference                       `json:"encounter,omitempty" bson:"encounter,omitempty"`                          // Encounter immunization was part of
	SupportingInformation []Reference                      `json:"supportingInformation,omitempty" bson:"supporting_information,omitempty"` // Additional information in support of the immunization
	OccurrenceDateTime    *string                          `json:"occurrenceDateTime" bson:"occurrence_date_time"`                          // Vaccine administration date
	OccurrenceString      *string                          `json:"occurrenceString" bson:"occurrence_string"`                               // Vaccine administration date
	PrimarySource         bool                             `json:"primarySource,omitempty" bson:"primary_source,omitempty"`                 // Indicates context the data was captured in
	InformationSource     *CodeableReference               `json:"informationSource,omitempty" bson:"information_source,omitempty"`         // Indicates the source of a  reported record
	Location              *Reference                       `json:"location,omitempty" bson:"location,omitempty"`                            // The service delivery location
	Site                  *CodeableConcept                 `json:"site,omitempty" bson:"site,omitempty"`                                    // Body site vaccine  was administered
	Route                 *CodeableConcept                 `json:"route,omitempty" bson:"route,omitempty"`                                  // How vaccine entered body
	DoseQuantity          *Quantity                        `json:"doseQuantity,omitempty" bson:"dose_quantity,omitempty"`                   // Amount of vaccine administered
	Performer             []ImmunizationPerformer          `json:"performer,omitempty" bson:"performer,omitempty"`                          // Who performed event
	Note                  []Annotation                     `json:"note,omitempty" bson:"note,omitempty"`                                    // Additional immunization notes
	Reason                []CodeableReference              `json:"reason,omitempty" bson:"reason,omitempty"`                                // Why immunization occurred
	IsSubpotent           bool                             `json:"isSubpotent,omitempty" bson:"is_subpotent,omitempty"`                     // Dose potency
	SubpotentReason       []CodeableConcept                `json:"subpotentReason,omitempty" bson:"subpotent_reason,omitempty"`             // Reason for being subpotent
	ProgramEligibility    []ImmunizationProgramEligibility `json:"programEligibility,omitempty" bson:"program_eligibility,omitempty"`       // Patient eligibility for a specific vaccination program
	FundingSource         *CodeableConcept                 `json:"fundingSource,omitempty" bson:"funding_source,omitempty"`                 // Funding source for the vaccine
	Reaction              []ImmunizationReaction           `json:"reaction,omitempty" bson:"reaction,omitempty"`                            // Details of a reaction that followed the immunization
	ProtocolApplied       []ImmunizationProtocolApplied    `json:"protocolApplied,omitempty" bson:"protocol_applied,omitempty"`             // Protocol followed by the provider
}

func (r *Immunization) Validate() error {
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
	if r.StatusReason != nil {
		if err := r.StatusReason.Validate(); err != nil {
			return fmt.Errorf("StatusReason: %w", err)
		}
	}
	if r.VaccineCode == nil {
		return fmt.Errorf("field 'VaccineCode' is required")
	}
	if r.VaccineCode != nil {
		if err := r.VaccineCode.Validate(); err != nil {
			return fmt.Errorf("VaccineCode: %w", err)
		}
	}
	if r.AdministeredProduct != nil {
		if err := r.AdministeredProduct.Validate(); err != nil {
			return fmt.Errorf("AdministeredProduct: %w", err)
		}
	}
	if r.Manufacturer != nil {
		if err := r.Manufacturer.Validate(); err != nil {
			return fmt.Errorf("Manufacturer: %w", err)
		}
	}
	if r.Patient == nil {
		return fmt.Errorf("field 'Patient' is required")
	}
	if r.Patient != nil {
		if err := r.Patient.Validate(); err != nil {
			return fmt.Errorf("Patient: %w", err)
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
	if r.OccurrenceDateTime == nil {
		return fmt.Errorf("field 'OccurrenceDateTime' is required")
	}
	if r.OccurrenceString == nil {
		return fmt.Errorf("field 'OccurrenceString' is required")
	}
	if r.InformationSource != nil {
		if err := r.InformationSource.Validate(); err != nil {
			return fmt.Errorf("InformationSource: %w", err)
		}
	}
	if r.Location != nil {
		if err := r.Location.Validate(); err != nil {
			return fmt.Errorf("Location: %w", err)
		}
	}
	if r.Site != nil {
		if err := r.Site.Validate(); err != nil {
			return fmt.Errorf("Site: %w", err)
		}
	}
	if r.Route != nil {
		if err := r.Route.Validate(); err != nil {
			return fmt.Errorf("Route: %w", err)
		}
	}
	if r.DoseQuantity != nil {
		if err := r.DoseQuantity.Validate(); err != nil {
			return fmt.Errorf("DoseQuantity: %w", err)
		}
	}
	for i, item := range r.Performer {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Performer[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	for i, item := range r.Reason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reason[%d]: %w", i, err)
		}
	}
	for i, item := range r.SubpotentReason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SubpotentReason[%d]: %w", i, err)
		}
	}
	for i, item := range r.ProgramEligibility {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ProgramEligibility[%d]: %w", i, err)
		}
	}
	if r.FundingSource != nil {
		if err := r.FundingSource.Validate(); err != nil {
			return fmt.Errorf("FundingSource: %w", err)
		}
	}
	for i, item := range r.Reaction {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reaction[%d]: %w", i, err)
		}
	}
	for i, item := range r.ProtocolApplied {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ProtocolApplied[%d]: %w", i, err)
		}
	}
	return nil
}

type ImmunizationReaction struct {
	Id            *string            `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	Date          *string            `json:"date,omitempty" bson:"date,omitempty"`                   // When reaction started
	Manifestation *CodeableReference `json:"manifestation,omitempty" bson:"manifestation,omitempty"` // Additional information on reaction
	Reported      bool               `json:"reported,omitempty" bson:"reported,omitempty"`           // Indicates self-reported reaction
}

func (r *ImmunizationReaction) Validate() error {
	if r.Manifestation != nil {
		if err := r.Manifestation.Validate(); err != nil {
			return fmt.Errorf("Manifestation: %w", err)
		}
	}
	return nil
}

type ImmunizationProtocolApplied struct {
	Id            *string           `json:"id,omitempty" bson:"id,omitempty"`                        // Unique id for inter-element referencing
	Series        *string           `json:"series,omitempty" bson:"series,omitempty"`                // Name of vaccine series
	Authority     *Reference        `json:"authority,omitempty" bson:"authority,omitempty"`          // Who is responsible for publishing the recommendations
	TargetDisease []CodeableConcept `json:"targetDisease,omitempty" bson:"target_disease,omitempty"` // Vaccine preventable disease being targeted
	DoseNumber    *CodeableConcept  `json:"doseNumber,omitempty" bson:"dose_number,omitempty"`       // Dose number within series
	SeriesDoses   *CodeableConcept  `json:"seriesDoses,omitempty" bson:"series_doses,omitempty"`     // Recommended number of doses for immunity
}

func (r *ImmunizationProtocolApplied) Validate() error {
	if r.Authority != nil {
		if err := r.Authority.Validate(); err != nil {
			return fmt.Errorf("Authority: %w", err)
		}
	}
	for i, item := range r.TargetDisease {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("TargetDisease[%d]: %w", i, err)
		}
	}
	if r.DoseNumber != nil {
		if err := r.DoseNumber.Validate(); err != nil {
			return fmt.Errorf("DoseNumber: %w", err)
		}
	}
	if r.SeriesDoses != nil {
		if err := r.SeriesDoses.Validate(); err != nil {
			return fmt.Errorf("SeriesDoses: %w", err)
		}
	}
	return nil
}

type ImmunizationPerformer struct {
	Id       *string          `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Function *CodeableConcept `json:"function,omitempty" bson:"function,omitempty"` // Type of performance
	Actor    *Reference       `json:"actor" bson:"actor"`                           // Individual or organization who performed the event
}

func (r *ImmunizationPerformer) Validate() error {
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

type ImmunizationProgramEligibility struct {
	Id            *string          `json:"id,omitempty" bson:"id,omitempty"`    // Unique id for inter-element referencing
	Program       *CodeableConcept `json:"program" bson:"program"`              // The program that eligibility is declared for
	ProgramStatus *CodeableConcept `json:"programStatus" bson:"program_status"` // The patient's eligibility status for the program
}

func (r *ImmunizationProgramEligibility) Validate() error {
	if r.Program == nil {
		return fmt.Errorf("field 'Program' is required")
	}
	if r.Program != nil {
		if err := r.Program.Validate(); err != nil {
			return fmt.Errorf("Program: %w", err)
		}
	}
	if r.ProgramStatus == nil {
		return fmt.Errorf("field 'ProgramStatus' is required")
	}
	if r.ProgramStatus != nil {
		if err := r.ProgramStatus.Validate(); err != nil {
			return fmt.Errorf("ProgramStatus: %w", err)
		}
	}
	return nil
}
