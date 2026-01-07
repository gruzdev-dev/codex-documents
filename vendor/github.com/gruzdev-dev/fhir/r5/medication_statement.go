package models

import (
	"encoding/json"
	"fmt"
)

// A record of a medication consumed by a patient.  A MedicationStatement may indicate that the patient may be taking the medication now or has taken the medication in the past or will be taking the medication in the future. The source of this information can be the patient, patient representative (e.g., spouse, significant other, family member, caregiver), or a clinician. A common scenario where this information is captured is during the history taking process during a patient encounter or stay. The medication information may come from sources such as the patient's memory, from a prescription bottle, or from a list of medications the patient, clinician or other party maintains. The primary difference between a MedicationStatement and a MedicationAdministration is that the medication administration has complete administration information and is based on actual administration information from the person who administered the medication. A MedicationStatement is often, if not always, less specific. There is no required date/time when the medication was administered, in fact we only know that a source has reported the patient is taking this medication, where details such as time, quantity, or rate or even medication product may be incomplete or missing or less precise. As stated earlier, the MedicationStatement information may come from the patient's memory, from a prescription bottle or from a list of medications the patient, clinician or other party. MedicationAdministration is more formal and is not missing detailed information.
type MedicationStatement struct {
	ResourceType               string                        `json:"resourceType" bson:"resource_type"`                                                  // Type of resource
	Id                         *string                       `json:"id,omitempty" bson:"id,omitempty"`                                                   // Logical id of this artifact
	Meta                       *Meta                         `json:"meta,omitempty" bson:"meta,omitempty"`                                               // Metadata about the resource
	ImplicitRules              *string                       `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                            // A set of rules under which this content was created
	Language                   *string                       `json:"language,omitempty" bson:"language,omitempty"`                                       // Language of the resource content
	Text                       *Narrative                    `json:"text,omitempty" bson:"text,omitempty"`                                               // Text summary of the resource, for human interpretation
	Contained                  []json.RawMessage             `json:"contained,omitempty" bson:"contained,omitempty"`                                     // Contained, inline Resources
	Identifier                 []Identifier                  `json:"identifier,omitempty" bson:"identifier,omitempty"`                                   // External identifier
	PartOf                     []Reference                   `json:"partOf,omitempty" bson:"part_of,omitempty"`                                          // Part of referenced event
	Status                     string                        `json:"status" bson:"status"`                                                               // recorded | entered-in-error | draft
	Category                   []CodeableConcept             `json:"category,omitempty" bson:"category,omitempty"`                                       // Type of medication statement
	Medication                 *CodeableReference            `json:"medication" bson:"medication"`                                                       // The medication that is/was taken (or not taken)
	Subject                    *Reference                    `json:"subject" bson:"subject"`                                                             // Who is/was taking  the medication
	Encounter                  *Reference                    `json:"encounter,omitempty" bson:"encounter,omitempty"`                                     // Encounter associated with MedicationStatement
	EffectiveDateTime          *string                       `json:"effectiveDateTime,omitempty" bson:"effective_date_time,omitempty"`                   // The date/time or interval when the medication is/was/will be taken
	EffectivePeriod            *Period                       `json:"effectivePeriod,omitempty" bson:"effective_period,omitempty"`                        // The date/time or interval when the medication is/was/will be taken
	EffectiveTiming            *Timing                       `json:"effectiveTiming,omitempty" bson:"effective_timing,omitempty"`                        // The date/time or interval when the medication is/was/will be taken
	DateAsserted               *string                       `json:"dateAsserted,omitempty" bson:"date_asserted,omitempty"`                              // When the usage was asserted?
	Author                     *Reference                    `json:"author,omitempty" bson:"author,omitempty"`                                           // Who/What authored the statement
	InformationSource          []Reference                   `json:"informationSource,omitempty" bson:"information_source,omitempty"`                    // Person or organization that provided the information about the taking of this medication
	DerivedFrom                []Reference                   `json:"derivedFrom,omitempty" bson:"derived_from,omitempty"`                                // Link to information used to derive the MedicationStatement
	Reason                     []CodeableReference           `json:"reason,omitempty" bson:"reason,omitempty"`                                           // Reason for why the medication is being/was taken
	Note                       []Annotation                  `json:"note,omitempty" bson:"note,omitempty"`                                               // Further information about the usage
	RelatedClinicalInformation []Reference                   `json:"relatedClinicalInformation,omitempty" bson:"related_clinical_information,omitempty"` // Link to information relevant to the usage of a medication
	Dosage                     *DosageDetails                `json:"dosage,omitempty" bson:"dosage,omitempty"`                                           // Details of how medication is/was taken or should be taken
	Adherence                  *MedicationStatementAdherence `json:"adherence,omitempty" bson:"adherence,omitempty"`                                     // Indicates whether the medication is or is not being consumed or administered
}

func (r *MedicationStatement) Validate() error {
	if r.ResourceType != "MedicationStatement" {
		return fmt.Errorf("invalid resourceType: expected 'MedicationStatement', got '%s'", r.ResourceType)
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
	for i, item := range r.PartOf {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("PartOf[%d]: %w", i, err)
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
	if r.Medication == nil {
		return fmt.Errorf("field 'Medication' is required")
	}
	if r.Medication != nil {
		if err := r.Medication.Validate(); err != nil {
			return fmt.Errorf("Medication: %w", err)
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
	if r.EffectivePeriod != nil {
		if err := r.EffectivePeriod.Validate(); err != nil {
			return fmt.Errorf("EffectivePeriod: %w", err)
		}
	}
	if r.EffectiveTiming != nil {
		if err := r.EffectiveTiming.Validate(); err != nil {
			return fmt.Errorf("EffectiveTiming: %w", err)
		}
	}
	if r.Author != nil {
		if err := r.Author.Validate(); err != nil {
			return fmt.Errorf("Author: %w", err)
		}
	}
	for i, item := range r.InformationSource {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("InformationSource[%d]: %w", i, err)
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
	for i, item := range r.RelatedClinicalInformation {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RelatedClinicalInformation[%d]: %w", i, err)
		}
	}
	if r.Dosage != nil {
		if err := r.Dosage.Validate(); err != nil {
			return fmt.Errorf("Dosage: %w", err)
		}
	}
	if r.Adherence != nil {
		if err := r.Adherence.Validate(); err != nil {
			return fmt.Errorf("Adherence: %w", err)
		}
	}
	return nil
}

type MedicationStatementAdherence struct {
	Id     *string          `json:"id,omitempty" bson:"id,omitempty"`         // Unique id for inter-element referencing
	Code   *CodeableConcept `json:"code" bson:"code"`                         // Type of adherence
	Reason *CodeableConcept `json:"reason,omitempty" bson:"reason,omitempty"` // Details of the reason for the current use of the medication
}

func (r *MedicationStatementAdherence) Validate() error {
	if r.Code == nil {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.Reason != nil {
		if err := r.Reason.Validate(); err != nil {
			return fmt.Errorf("Reason: %w", err)
		}
	}
	return nil
}
