package models

import "encoding/json"

// Measurements and simple assertions made about a patient, device or other subject.
type Observation struct {
	Id                    *string                     `json:"id,omitempty"`                    // Logical id of this artifact
	Meta                  *Meta                       `json:"meta,omitempty"`                  // Metadata about the resource
	ImplicitRules         *string                     `json:"implicitRules,omitempty"`         // A set of rules under which this content was created
	Language              *string                     `json:"language,omitempty"`              // Language of the resource content
	Text                  *Narrative                  `json:"text,omitempty"`                  // Text summary of the resource, for human interpretation
	Contained             []json.RawMessage           `json:"contained,omitempty"`             // Contained, inline Resources
	Identifier            []Identifier                `json:"identifier,omitempty"`            // Business Identifier for observation
	BasedOn               []Reference                 `json:"basedOn,omitempty"`               // Fulfills plan, proposal or order
	TriggeredBy           []ObservationTriggeredBy    `json:"triggeredBy,omitempty"`           // Triggering observation(s)
	PartOf                []Reference                 `json:"partOf,omitempty"`                // Part of referenced event
	Status                string                      `json:"status,omitempty"`                // registered | specimen-in-process | preliminary | final | amended | corrected | appended | cancelled | entered-in-error | unknown | cannot-be-obtained
	Category              []CodeableConcept           `json:"category,omitempty"`              // Classification of  type of observation
	Code                  CodeableConcept             `json:"code,omitempty"`                  // Type of observation (code / type)
	Subject               *Reference                  `json:"subject,omitempty"`               // Who and/or what the observation is about
	Focus                 []Reference                 `json:"focus,omitempty"`                 // What the observation is about, when it is not about the subject of record
	Organizer             bool                        `json:"organizer,omitempty"`             // This observation organizes/groups a set of sub-observations
	Encounter             *Reference                  `json:"encounter,omitempty"`             // Healthcare event during which this observation is made. If you need to place the observation within one or more episodes of care, use the workflow-episodeOfCare extension
	EffectiveDateTime     *string                     `json:"effectiveDateTime,omitempty"`     // Clinically relevant time/time-period for observation
	EffectivePeriod       *Period                     `json:"effectivePeriod,omitempty"`       // Clinically relevant time/time-period for observation
	EffectiveTiming       *Timing                     `json:"effectiveTiming,omitempty"`       // Clinically relevant time/time-period for observation
	EffectiveInstant      *string                     `json:"effectiveInstant,omitempty"`      // Clinically relevant time/time-period for observation
	Issued                *string                     `json:"issued,omitempty"`                // Date/Time this version was made available
	Performer             []Reference                 `json:"performer,omitempty"`             // Who is responsible for the observation
	ValueQuantity         *Quantity                   `json:"valueQuantity,omitempty"`         // Actual result
	ValueCodeableConcept  *CodeableConcept            `json:"valueCodeableConcept,omitempty"`  // Actual result
	ValueString           *string                     `json:"valueString,omitempty"`           // Actual result
	ValueBoolean          *bool                       `json:"valueBoolean,omitempty"`          // Actual result
	ValueInteger          *int                        `json:"valueInteger,omitempty"`          // Actual result
	ValueRange            *Range                      `json:"valueRange,omitempty"`            // Actual result
	ValueRatio            *Ratio                      `json:"valueRatio,omitempty"`            // Actual result
	ValueSampledData      *SampledData                `json:"valueSampledData,omitempty"`      // Actual result
	ValueTime             *string                     `json:"valueTime,omitempty"`             // Actual result
	ValueDateTime         *string                     `json:"valueDateTime,omitempty"`         // Actual result
	ValuePeriod           *Period                     `json:"valuePeriod,omitempty"`           // Actual result
	ValueAttachment       *Attachment                 `json:"valueAttachment,omitempty"`       // Actual result
	DataAbsentReason      *CodeableConcept            `json:"dataAbsentReason,omitempty"`      // Why the result value is missing
	Interpretation        []CodeableConcept           `json:"interpretation,omitempty"`        // High, low, normal, etc
	InterpretationContext []CodeableReference         `json:"interpretationContext,omitempty"` // Context for understanding the observation
	Note                  []Annotation                `json:"note,omitempty"`                  // Comments about the observation
	BodySite              *CodeableConcept            `json:"bodySite,omitempty"`              // DEPRECATED: Observed body part
	BodyStructure         *CodeableReference          `json:"bodyStructure,omitempty"`         // Observed body structure
	Method                *CodeableConcept            `json:"method,omitempty"`                // How it was done
	Specimen              *Reference                  `json:"specimen,omitempty"`              // Specimen used for this observation
	Device                *Reference                  `json:"device,omitempty"`                // A reference to the device that generates the measurements or the device settings for the device
	ReferenceRange        []ObservationReferenceRange `json:"referenceRange,omitempty"`        // Provides guide for interpretation
	HasMember             []Reference                 `json:"hasMember,omitempty"`             // Related resource that belongs to the Observation group
	DerivedFrom           []Reference                 `json:"derivedFrom,omitempty"`           // Related resource from which the observation is made
	Component             []ObservationComponent      `json:"component,omitempty"`             // Component results
}

type ObservationTriggeredBy struct {
	Id          *string   `json:"id,omitempty"`          // Unique id for inter-element referencing
	Observation Reference `json:"observation,omitempty"` // Triggering observation
	Type        string    `json:"type,omitempty"`        // reflex | repeat | re-run
	Reason      *string   `json:"reason,omitempty"`      // Reason that the observation was triggered
}

type ObservationReferenceRange struct {
	Id          *string           `json:"id,omitempty"`          // Unique id for inter-element referencing
	Low         *Quantity         `json:"low,omitempty"`         // Low Range, if relevant
	High        *Quantity         `json:"high,omitempty"`        // High Range, if relevant
	NormalValue *CodeableConcept  `json:"normalValue,omitempty"` // Normal value, if relevant
	Type        *CodeableConcept  `json:"type,omitempty"`        // Reference range qualifier
	AppliesTo   []CodeableConcept `json:"appliesTo,omitempty"`   // Reference range population
	Age         *Range            `json:"age,omitempty"`         // Applicable age range, if relevant
	Text        *string           `json:"text,omitempty"`        // Text based reference range in an observation
}

type ObservationComponent struct {
	Id                   *string                              `json:"id,omitempty"`                   // Unique id for inter-element referencing
	Code                 CodeableConcept                      `json:"code,omitempty"`                 // Type of component observation (code / type)
	ValueQuantity        *Quantity                            `json:"valueQuantity,omitempty"`        // Actual component result
	ValueCodeableConcept *CodeableConcept                     `json:"valueCodeableConcept,omitempty"` // Actual component result
	ValueString          *string                              `json:"valueString,omitempty"`          // Actual component result
	ValueBoolean         *bool                                `json:"valueBoolean,omitempty"`         // Actual component result
	ValueInteger         *int                                 `json:"valueInteger,omitempty"`         // Actual component result
	ValueRange           *Range                               `json:"valueRange,omitempty"`           // Actual component result
	ValueRatio           *Ratio                               `json:"valueRatio,omitempty"`           // Actual component result
	ValueSampledData     *SampledData                         `json:"valueSampledData,omitempty"`     // Actual component result
	ValueTime            *string                              `json:"valueTime,omitempty"`            // Actual component result
	ValueDateTime        *string                              `json:"valueDateTime,omitempty"`        // Actual component result
	ValuePeriod          *Period                              `json:"valuePeriod,omitempty"`          // Actual component result
	ValueAttachment      *Attachment                          `json:"valueAttachment,omitempty"`      // Actual component result
	DataAbsentReason     *CodeableConcept                     `json:"dataAbsentReason,omitempty"`     // Why the component result value is missing
	Interpretation       []CodeableConcept                    `json:"interpretation,omitempty"`       // High, low, normal, etc
	ReferenceRange       []ObservationComponentReferenceRange `json:"referenceRange,omitempty"`       // Provides guide for interpretation of component result value
}

type ObservationComponentReferenceRange struct {
}
