package models

import (
	"encoding/json"
	"fmt"
)

// A sample to be used for analysis.
type Specimen struct {
	ResourceType  string               `json:"resourceType" bson:"resource_type"`                       // Type of resource
	Id            *string              `json:"id,omitempty" bson:"id,omitempty"`                        // Logical id of this artifact
	Meta          *Meta                `json:"meta,omitempty" bson:"meta,omitempty"`                    // Metadata about the resource
	ImplicitRules *string              `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"` // A set of rules under which this content was created
	Language      *string              `json:"language,omitempty" bson:"language,omitempty"`            // Language of the resource content
	Text          *Narrative           `json:"text,omitempty" bson:"text,omitempty"`                    // Text summary of the resource, for human interpretation
	Contained     []json.RawMessage    `json:"contained,omitempty" bson:"contained,omitempty"`          // Contained, inline Resources
	Identifier    []Identifier         `json:"identifier,omitempty" bson:"identifier,omitempty"`        // Primary specimen identifier
	Status        *string              `json:"status,omitempty" bson:"status,omitempty"`                // available | unavailable | unsatisfactory | entered-in-error
	Type          *CodeableConcept     `json:"type,omitempty" bson:"type,omitempty"`                    // Kind of material that forms the specimen
	Subject       *Reference           `json:"subject,omitempty" bson:"subject,omitempty"`              // Where the specimen came from. This may be from patient(s), from a location (e.g., the source of an environmental sample), or a sampling of a substance, a biologically-derived product, or a device
	ReceivedTime  *string              `json:"receivedTime,omitempty" bson:"received_time,omitempty"`   // The time when specimen is received by the testing laboratory
	Parent        []Reference          `json:"parent,omitempty" bson:"parent,omitempty"`                // Specimen from which this specimen originated
	Request       []Reference          `json:"request,omitempty" bson:"request,omitempty"`              // Why the specimen was collected
	Combined      *string              `json:"combined,omitempty" bson:"combined,omitempty"`            // grouped | pooled
	Role          []CodeableConcept    `json:"role,omitempty" bson:"role,omitempty"`                    // The role the specimen serves
	Feature       []SpecimenFeature    `json:"feature,omitempty" bson:"feature,omitempty"`              // The physical feature of a specimen
	Collection    *SpecimenCollection  `json:"collection,omitempty" bson:"collection,omitempty"`        // Collection details
	Processing    []SpecimenProcessing `json:"processing,omitempty" bson:"processing,omitempty"`        // Processing and processing step details
	Container     []SpecimenContainer  `json:"container,omitempty" bson:"container,omitempty"`          // Direct container of specimen (tube/slide, etc.)
	Condition     []CodeableConcept    `json:"condition,omitempty" bson:"condition,omitempty"`          // State of the specimen
	Note          []Annotation         `json:"note,omitempty" bson:"note,omitempty"`                    // Comments
}

func (r *Specimen) Validate() error {
	if r.ResourceType != "Specimen" {
		return fmt.Errorf("invalid resourceType: expected 'Specimen', got '%s'", r.ResourceType)
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
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Subject != nil {
		if err := r.Subject.Validate(); err != nil {
			return fmt.Errorf("Subject: %w", err)
		}
	}
	for i, item := range r.Parent {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Parent[%d]: %w", i, err)
		}
	}
	for i, item := range r.Request {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Request[%d]: %w", i, err)
		}
	}
	for i, item := range r.Role {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Role[%d]: %w", i, err)
		}
	}
	for i, item := range r.Feature {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Feature[%d]: %w", i, err)
		}
	}
	if r.Collection != nil {
		if err := r.Collection.Validate(); err != nil {
			return fmt.Errorf("Collection: %w", err)
		}
	}
	for i, item := range r.Processing {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Processing[%d]: %w", i, err)
		}
	}
	for i, item := range r.Container {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Container[%d]: %w", i, err)
		}
	}
	for i, item := range r.Condition {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Condition[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	return nil
}

type SpecimenFeature struct {
	Id          *string          `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Type        *CodeableConcept `json:"type" bson:"type"`                 // Highlighted feature
	Description string           `json:"description" bson:"description"`   // Information about the feature
}

func (r *SpecimenFeature) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	var emptyString string
	if r.Description == emptyString {
		return fmt.Errorf("field 'Description' is required")
	}
	return nil
}

type SpecimenCollection struct {
	Id                           *string            `json:"id,omitempty" bson:"id,omitempty"`                                                        // Unique id for inter-element referencing
	Collector                    *Reference         `json:"collector,omitempty" bson:"collector,omitempty"`                                          // Who collected the specimen
	CollectedDateTime            *string            `json:"collectedDateTime,omitempty" bson:"collected_date_time,omitempty"`                        // Collection time
	CollectedPeriod              *Period            `json:"collectedPeriod,omitempty" bson:"collected_period,omitempty"`                             // Collection time
	Duration                     *Duration          `json:"duration,omitempty" bson:"duration,omitempty"`                                            // How long it took to collect specimen
	Quantity                     *Quantity          `json:"quantity,omitempty" bson:"quantity,omitempty"`                                            // The quantity of specimen collected
	Method                       *CodeableConcept   `json:"method,omitempty" bson:"method,omitempty"`                                                // Technique used to perform collection
	DeviceCodeableConcept        *CodeableConcept   `json:"deviceCodeableConcept,omitempty" bson:"device_codeable_concept,omitempty"`                // Device used to perform collection
	DeviceReference              *Reference         `json:"deviceReference,omitempty" bson:"device_reference,omitempty"`                             // Device used to perform collection
	DeviceCanonical              *string            `json:"deviceCanonical,omitempty" bson:"device_canonical,omitempty"`                             // Device used to perform collection
	Procedure                    *Reference         `json:"procedure,omitempty" bson:"procedure,omitempty"`                                          // The procedure that collects the specimen
	BodySite                     *CodeableReference `json:"bodySite,omitempty" bson:"body_site,omitempty"`                                           // Anatomical collection site
	FastingStatusCodeableConcept *CodeableConcept   `json:"fastingStatusCodeableConcept,omitempty" bson:"fasting_status_codeable_concept,omitempty"` // Whether or how long patient abstained from food and/or drink
	FastingStatusDuration        *Duration          `json:"fastingStatusDuration,omitempty" bson:"fasting_status_duration,omitempty"`                // Whether or how long patient abstained from food and/or drink
}

func (r *SpecimenCollection) Validate() error {
	if r.Collector != nil {
		if err := r.Collector.Validate(); err != nil {
			return fmt.Errorf("Collector: %w", err)
		}
	}
	if r.CollectedPeriod != nil {
		if err := r.CollectedPeriod.Validate(); err != nil {
			return fmt.Errorf("CollectedPeriod: %w", err)
		}
	}
	if r.Duration != nil {
		if err := r.Duration.Validate(); err != nil {
			return fmt.Errorf("Duration: %w", err)
		}
	}
	if r.Quantity != nil {
		if err := r.Quantity.Validate(); err != nil {
			return fmt.Errorf("Quantity: %w", err)
		}
	}
	if r.Method != nil {
		if err := r.Method.Validate(); err != nil {
			return fmt.Errorf("Method: %w", err)
		}
	}
	if r.DeviceCodeableConcept != nil {
		if err := r.DeviceCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("DeviceCodeableConcept: %w", err)
		}
	}
	if r.DeviceReference != nil {
		if err := r.DeviceReference.Validate(); err != nil {
			return fmt.Errorf("DeviceReference: %w", err)
		}
	}
	if r.Procedure != nil {
		if err := r.Procedure.Validate(); err != nil {
			return fmt.Errorf("Procedure: %w", err)
		}
	}
	if r.BodySite != nil {
		if err := r.BodySite.Validate(); err != nil {
			return fmt.Errorf("BodySite: %w", err)
		}
	}
	if r.FastingStatusCodeableConcept != nil {
		if err := r.FastingStatusCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("FastingStatusCodeableConcept: %w", err)
		}
	}
	if r.FastingStatusDuration != nil {
		if err := r.FastingStatusDuration.Validate(); err != nil {
			return fmt.Errorf("FastingStatusDuration: %w", err)
		}
	}
	return nil
}

type SpecimenProcessing struct {
	Id                    *string             `json:"id,omitempty" bson:"id,omitempty"`                                         // Unique id for inter-element referencing
	Description           *string             `json:"description,omitempty" bson:"description,omitempty"`                       // Textual description of procedure
	Method                *CodeableConcept    `json:"method,omitempty" bson:"method,omitempty"`                                 // Indicates the treatment step  applied to the specimen
	Performer             *Reference          `json:"performer,omitempty" bson:"performer,omitempty"`                           // Entity processing specimen
	DeviceCodeableConcept *CodeableConcept    `json:"deviceCodeableConcept,omitempty" bson:"device_codeable_concept,omitempty"` // Device used to process the specimen
	DeviceReference       *Reference          `json:"deviceReference,omitempty" bson:"device_reference,omitempty"`              // Device used to process the specimen
	DeviceCanonical       *string             `json:"deviceCanonical,omitempty" bson:"device_canonical,omitempty"`              // Device used to process the specimen
	Additive              []CodeableReference `json:"additive,omitempty" bson:"additive,omitempty"`                             // Material used in the processing step
	TimeDateTime          *string             `json:"timeDateTime,omitempty" bson:"time_date_time,omitempty"`                   // Date and time of specimen processing
	TimePeriod            *Period             `json:"timePeriod,omitempty" bson:"time_period,omitempty"`                        // Date and time of specimen processing
	TimeDuration          *Duration           `json:"timeDuration,omitempty" bson:"time_duration,omitempty"`                    // Date and time of specimen processing
}

func (r *SpecimenProcessing) Validate() error {
	if r.Method != nil {
		if err := r.Method.Validate(); err != nil {
			return fmt.Errorf("Method: %w", err)
		}
	}
	if r.Performer != nil {
		if err := r.Performer.Validate(); err != nil {
			return fmt.Errorf("Performer: %w", err)
		}
	}
	if r.DeviceCodeableConcept != nil {
		if err := r.DeviceCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("DeviceCodeableConcept: %w", err)
		}
	}
	if r.DeviceReference != nil {
		if err := r.DeviceReference.Validate(); err != nil {
			return fmt.Errorf("DeviceReference: %w", err)
		}
	}
	for i, item := range r.Additive {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Additive[%d]: %w", i, err)
		}
	}
	if r.TimePeriod != nil {
		if err := r.TimePeriod.Validate(); err != nil {
			return fmt.Errorf("TimePeriod: %w", err)
		}
	}
	if r.TimeDuration != nil {
		if err := r.TimeDuration.Validate(); err != nil {
			return fmt.Errorf("TimeDuration: %w", err)
		}
	}
	return nil
}

type SpecimenContainer struct {
	Id                    *string          `json:"id,omitempty" bson:"id,omitempty"`                              // Unique id for inter-element referencing
	DeviceCodeableConcept *CodeableConcept `json:"deviceCodeableConcept" bson:"device_codeable_concept"`          // Device resource for the container
	DeviceReference       *Reference       `json:"deviceReference" bson:"device_reference"`                       // Device resource for the container
	DeviceCanonical       *string          `json:"deviceCanonical" bson:"device_canonical"`                       // Device resource for the container
	SpecimenQuantity      *Quantity        `json:"specimenQuantity,omitempty" bson:"specimen_quantity,omitempty"` // Quantity of specimen within container
}

func (r *SpecimenContainer) Validate() error {
	if r.DeviceCodeableConcept == nil {
		return fmt.Errorf("field 'DeviceCodeableConcept' is required")
	}
	if r.DeviceCodeableConcept != nil {
		if err := r.DeviceCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("DeviceCodeableConcept: %w", err)
		}
	}
	if r.DeviceReference == nil {
		return fmt.Errorf("field 'DeviceReference' is required")
	}
	if r.DeviceReference != nil {
		if err := r.DeviceReference.Validate(); err != nil {
			return fmt.Errorf("DeviceReference: %w", err)
		}
	}
	if r.DeviceCanonical == nil {
		return fmt.Errorf("field 'DeviceCanonical' is required")
	}
	if r.SpecimenQuantity != nil {
		if err := r.SpecimenQuantity.Validate(); err != nil {
			return fmt.Errorf("SpecimenQuantity: %w", err)
		}
	}
	return nil
}
