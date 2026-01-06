package models

import "encoding/json"

// Measurements and simple assertions made about a patient, device or other subject.
type Observation struct {
	Id                    *string                     `json:"id,omitempty" bson:"id,omitempty"`                                        // Logical id of this artifact
	Meta                  *Meta                       `json:"meta,omitempty" bson:"meta,omitempty"`                                    // Metadata about the resource
	ImplicitRules         *string                     `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                 // A set of rules under which this content was created
	Language              *string                     `json:"language,omitempty" bson:"language,omitempty"`                            // Language of the resource content
	Text                  *Narrative                  `json:"text,omitempty" bson:"text,omitempty"`                                    // Text summary of the resource, for human interpretation
	Contained             []json.RawMessage           `json:"contained,omitempty" bson:"contained,omitempty"`                          // Contained, inline Resources
	Identifier            []Identifier                `json:"identifier,omitempty" bson:"identifier,omitempty"`                        // Business Identifier for observation
	BasedOn               []Reference                 `json:"basedOn,omitempty" bson:"based_on,omitempty"`                             // Fulfills plan, proposal or order
	TriggeredBy           []ObservationTriggeredBy    `json:"triggeredBy,omitempty" bson:"triggered_by,omitempty"`                     // Triggering observation(s)
	PartOf                []Reference                 `json:"partOf,omitempty" bson:"part_of,omitempty"`                               // Part of referenced event
	Status                string                      `json:"status" bson:"status"`                                                    // registered | specimen-in-process | preliminary | final | amended | corrected | appended | cancelled | entered-in-error | unknown | cannot-be-obtained
	Category              []CodeableConcept           `json:"category,omitempty" bson:"category,omitempty"`                            // Classification of  type of observation
	Code                  *CodeableConcept            `json:"code" bson:"code"`                                                        // Type of observation (code / type)
	Subject               *Reference                  `json:"subject,omitempty" bson:"subject,omitempty"`                              // Who and/or what the observation is about
	Focus                 []Reference                 `json:"focus,omitempty" bson:"focus,omitempty"`                                  // What the observation is about, when it is not about the subject of record
	Organizer             bool                        `json:"organizer,omitempty" bson:"organizer,omitempty"`                          // This observation organizes/groups a set of sub-observations
	Encounter             *Reference                  `json:"encounter,omitempty" bson:"encounter,omitempty"`                          // Healthcare event during which this observation is made. If you need to place the observation within one or more episodes of care, use the workflow-episodeOfCare extension
	EffectiveDateTime     *string                     `json:"effectiveDateTime,omitempty" bson:"effective_date_time,omitempty"`        // Clinically relevant time/time-period for observation
	EffectivePeriod       *Period                     `json:"effectivePeriod,omitempty" bson:"effective_period,omitempty"`             // Clinically relevant time/time-period for observation
	EffectiveTiming       *Timing                     `json:"effectiveTiming,omitempty" bson:"effective_timing,omitempty"`             // Clinically relevant time/time-period for observation
	EffectiveInstant      *string                     `json:"effectiveInstant,omitempty" bson:"effective_instant,omitempty"`           // Clinically relevant time/time-period for observation
	Issued                *string                     `json:"issued,omitempty" bson:"issued,omitempty"`                                // Date/Time this version was made available
	Performer             []Reference                 `json:"performer,omitempty" bson:"performer,omitempty"`                          // Who is responsible for the observation
	ValueQuantity         *Quantity                   `json:"valueQuantity,omitempty" bson:"value_quantity,omitempty"`                 // Actual result
	ValueCodeableConcept  *CodeableConcept            `json:"valueCodeableConcept,omitempty" bson:"value_codeable_concept,omitempty"`  // Actual result
	ValueString           *string                     `json:"valueString,omitempty" bson:"value_string,omitempty"`                     // Actual result
	ValueBoolean          *bool                       `json:"valueBoolean,omitempty" bson:"value_boolean,omitempty"`                   // Actual result
	ValueInteger          *int                        `json:"valueInteger,omitempty" bson:"value_integer,omitempty"`                   // Actual result
	ValueRange            *Range                      `json:"valueRange,omitempty" bson:"value_range,omitempty"`                       // Actual result
	ValueRatio            *Ratio                      `json:"valueRatio,omitempty" bson:"value_ratio,omitempty"`                       // Actual result
	ValueSampledData      *SampledData                `json:"valueSampledData,omitempty" bson:"value_sampled_data,omitempty"`          // Actual result
	ValueTime             *string                     `json:"valueTime,omitempty" bson:"value_time,omitempty"`                         // Actual result
	ValueDateTime         *string                     `json:"valueDateTime,omitempty" bson:"value_date_time,omitempty"`                // Actual result
	ValuePeriod           *Period                     `json:"valuePeriod,omitempty" bson:"value_period,omitempty"`                     // Actual result
	ValueAttachment       *Attachment                 `json:"valueAttachment,omitempty" bson:"value_attachment,omitempty"`             // Actual result
	DataAbsentReason      *CodeableConcept            `json:"dataAbsentReason,omitempty" bson:"data_absent_reason,omitempty"`          // Why the result value is missing
	Interpretation        []CodeableConcept           `json:"interpretation,omitempty" bson:"interpretation,omitempty"`                // High, low, normal, etc
	InterpretationContext []CodeableReference         `json:"interpretationContext,omitempty" bson:"interpretation_context,omitempty"` // Context for understanding the observation
	Note                  []Annotation                `json:"note,omitempty" bson:"note,omitempty"`                                    // Comments about the observation
	BodySite              *CodeableConcept            `json:"bodySite,omitempty" bson:"body_site,omitempty"`                           // DEPRECATED: Observed body part
	BodyStructure         *CodeableReference          `json:"bodyStructure,omitempty" bson:"body_structure,omitempty"`                 // Observed body structure
	Method                *CodeableConcept            `json:"method,omitempty" bson:"method,omitempty"`                                // How it was done
	Specimen              *Reference                  `json:"specimen,omitempty" bson:"specimen,omitempty"`                            // Specimen used for this observation
	Device                *Reference                  `json:"device,omitempty" bson:"device,omitempty"`                                // A reference to the device that generates the measurements or the device settings for the device
	ReferenceRange        []ObservationReferenceRange `json:"referenceRange,omitempty" bson:"reference_range,omitempty"`               // Provides guide for interpretation
	HasMember             []Reference                 `json:"hasMember,omitempty" bson:"has_member,omitempty"`                         // Related resource that belongs to the Observation group
	DerivedFrom           []Reference                 `json:"derivedFrom,omitempty" bson:"derived_from,omitempty"`                     // Related resource from which the observation is made
	Component             []ObservationComponent      `json:"component,omitempty" bson:"component,omitempty"`                          // Component results
}

type ObservationTriggeredBy struct {
	Id          *string    `json:"id,omitempty" bson:"id,omitempty"`         // Unique id for inter-element referencing
	Observation *Reference `json:"observation" bson:"observation"`           // Triggering observation
	Type        string     `json:"type" bson:"type"`                         // reflex | repeat | re-run
	Reason      *string    `json:"reason,omitempty" bson:"reason,omitempty"` // Reason that the observation was triggered
}

type ObservationReferenceRange struct {
	Id          *string           `json:"id,omitempty" bson:"id,omitempty"`                    // Unique id for inter-element referencing
	Low         *Quantity         `json:"low,omitempty" bson:"low,omitempty"`                  // Low Range, if relevant
	High        *Quantity         `json:"high,omitempty" bson:"high,omitempty"`                // High Range, if relevant
	NormalValue *CodeableConcept  `json:"normalValue,omitempty" bson:"normal_value,omitempty"` // Normal value, if relevant
	Type        *CodeableConcept  `json:"type,omitempty" bson:"type,omitempty"`                // Reference range qualifier
	AppliesTo   []CodeableConcept `json:"appliesTo,omitempty" bson:"applies_to,omitempty"`     // Reference range population
	Age         *Range            `json:"age,omitempty" bson:"age,omitempty"`                  // Applicable age range, if relevant
	Text        *string           `json:"text,omitempty" bson:"text,omitempty"`                // Text based reference range in an observation
}

type ObservationComponent struct {
	Id                   *string                     `json:"id,omitempty" bson:"id,omitempty"`                                       // Unique id for inter-element referencing
	Code                 *CodeableConcept            `json:"code" bson:"code"`                                                       // Type of component observation (code / type)
	ValueQuantity        *Quantity                   `json:"valueQuantity,omitempty" bson:"value_quantity,omitempty"`                // Actual component result
	ValueCodeableConcept *CodeableConcept            `json:"valueCodeableConcept,omitempty" bson:"value_codeable_concept,omitempty"` // Actual component result
	ValueString          *string                     `json:"valueString,omitempty" bson:"value_string,omitempty"`                    // Actual component result
	ValueBoolean         *bool                       `json:"valueBoolean,omitempty" bson:"value_boolean,omitempty"`                  // Actual component result
	ValueInteger         *int                        `json:"valueInteger,omitempty" bson:"value_integer,omitempty"`                  // Actual component result
	ValueRange           *Range                      `json:"valueRange,omitempty" bson:"value_range,omitempty"`                      // Actual component result
	ValueRatio           *Ratio                      `json:"valueRatio,omitempty" bson:"value_ratio,omitempty"`                      // Actual component result
	ValueSampledData     *SampledData                `json:"valueSampledData,omitempty" bson:"value_sampled_data,omitempty"`         // Actual component result
	ValueTime            *string                     `json:"valueTime,omitempty" bson:"value_time,omitempty"`                        // Actual component result
	ValueDateTime        *string                     `json:"valueDateTime,omitempty" bson:"value_date_time,omitempty"`               // Actual component result
	ValuePeriod          *Period                     `json:"valuePeriod,omitempty" bson:"value_period,omitempty"`                    // Actual component result
	ValueAttachment      *Attachment                 `json:"valueAttachment,omitempty" bson:"value_attachment,omitempty"`            // Actual component result
	DataAbsentReason     *CodeableConcept            `json:"dataAbsentReason,omitempty" bson:"data_absent_reason,omitempty"`         // Why the component result value is missing
	Interpretation       []CodeableConcept           `json:"interpretation,omitempty" bson:"interpretation,omitempty"`               // High, low, normal, etc
	ReferenceRange       []ObservationReferenceRange `json:"referenceRange,omitempty" bson:"reference_range,omitempty"`              // Provides guide for interpretation of component result value
}
