package models

import (
	"encoding/json"
	"fmt"
)

// A biological material originating from a biological entity intended to be transplanted or infused into another (possibly the same) biological entity.
type BiologicallyDerivedProduct struct {
	Id                      *string                               `json:"id,omitempty" bson:"id,omitempty"`                                             // Logical id of this artifact
	Meta                    *Meta                                 `json:"meta,omitempty" bson:"meta,omitempty"`                                         // Metadata about the resource
	ImplicitRules           *string                               `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                      // A set of rules under which this content was created
	Language                *string                               `json:"language,omitempty" bson:"language,omitempty"`                                 // Language of the resource content
	Text                    *Narrative                            `json:"text,omitempty" bson:"text,omitempty"`                                         // Text summary of the resource, for human interpretation
	Contained               []json.RawMessage                     `json:"contained,omitempty" bson:"contained,omitempty"`                               // Contained, inline Resources
	ProductCategory         []CodeableConcept                     `json:"productCategory,omitempty" bson:"product_category,omitempty"`                  // A category or classification of the product
	ProductCode             *CodeableConcept                      `json:"productCode,omitempty" bson:"product_code,omitempty"`                          // A code that identifies the kind of this biologically derived product
	Parent                  []Reference                           `json:"parent,omitempty" bson:"parent,omitempty"`                                     // The parent biologically-derived product
	Request                 []Reference                           `json:"request,omitempty" bson:"request,omitempty"`                                   // Request to obtain and/or infuse this product
	Identifier              []Identifier                          `json:"identifier,omitempty" bson:"identifier,omitempty"`                             // Instance identifier
	BiologicalSourceEvent   *Identifier                           `json:"biologicalSourceEvent,omitempty" bson:"biological_source_event,omitempty"`     // An identifier that supports traceability to the event during which material in this product from one or more biological entities was obtained or pooled
	ProcessingFacility      []Reference                           `json:"processingFacility,omitempty" bson:"processing_facility,omitempty"`            // Processing facilities responsible for the labeling and distribution of this biologically derived product
	Division                *string                               `json:"division,omitempty" bson:"division,omitempty"`                                 // A unique identifier for an aliquot of a product
	ProductStatus           *Coding                               `json:"productStatus,omitempty" bson:"product_status,omitempty"`                      // available | unavailable | processed | applied | discarded
	ExpirationDate          *string                               `json:"expirationDate,omitempty" bson:"expiration_date,omitempty"`                    // Date, and where relevant time, of expiration
	Collection              *BiologicallyDerivedProductCollection `json:"collection,omitempty" bson:"collection,omitempty"`                             // How this product was collected
	StorageTempRequirements *Range                                `json:"storageTempRequirements,omitempty" bson:"storage_temp_requirements,omitempty"` // Product storage temperature requirements
	Property                []BiologicallyDerivedProductProperty  `json:"property,omitempty" bson:"property,omitempty"`                                 // A property that is specific to this BiologicallyDerviedProduct instance
}

func (r *BiologicallyDerivedProduct) Validate() error {
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
	for i, item := range r.ProductCategory {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ProductCategory[%d]: %w", i, err)
		}
	}
	if r.ProductCode != nil {
		if err := r.ProductCode.Validate(); err != nil {
			return fmt.Errorf("ProductCode: %w", err)
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
	for i, item := range r.Identifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Identifier[%d]: %w", i, err)
		}
	}
	if r.BiologicalSourceEvent != nil {
		if err := r.BiologicalSourceEvent.Validate(); err != nil {
			return fmt.Errorf("BiologicalSourceEvent: %w", err)
		}
	}
	for i, item := range r.ProcessingFacility {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ProcessingFacility[%d]: %w", i, err)
		}
	}
	if r.ProductStatus != nil {
		if err := r.ProductStatus.Validate(); err != nil {
			return fmt.Errorf("ProductStatus: %w", err)
		}
	}
	if r.Collection != nil {
		if err := r.Collection.Validate(); err != nil {
			return fmt.Errorf("Collection: %w", err)
		}
	}
	if r.StorageTempRequirements != nil {
		if err := r.StorageTempRequirements.Validate(); err != nil {
			return fmt.Errorf("StorageTempRequirements: %w", err)
		}
	}
	for i, item := range r.Property {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Property[%d]: %w", i, err)
		}
	}
	return nil
}

type BiologicallyDerivedProductCollection struct {
	Id                 *string    `json:"id,omitempty" bson:"id,omitempty"`                                  // Unique id for inter-element referencing
	Collector          *Reference `json:"collector,omitempty" bson:"collector,omitempty"`                    // Individual performing the collection
	SourcePatient      *Reference `json:"sourcePatient,omitempty" bson:"source_patient,omitempty"`           // The patient who underwent the medical procedure to collect the product
	SourceOrganization *Reference `json:"sourceOrganization,omitempty" bson:"source_organization,omitempty"` // The organization that facilitated the collection
	CollectedDateTime  *string    `json:"collectedDateTime,omitempty" bson:"collected_date_time,omitempty"`  // Time of product collection
	CollectedPeriod    *Period    `json:"collectedPeriod,omitempty" bson:"collected_period,omitempty"`       // Time of product collection
	Procedure          *Reference `json:"procedure,omitempty" bson:"procedure,omitempty"`                    // The procedure involved in the collection
}

func (r *BiologicallyDerivedProductCollection) Validate() error {
	if r.Collector != nil {
		if err := r.Collector.Validate(); err != nil {
			return fmt.Errorf("Collector: %w", err)
		}
	}
	if r.SourcePatient != nil {
		if err := r.SourcePatient.Validate(); err != nil {
			return fmt.Errorf("SourcePatient: %w", err)
		}
	}
	if r.SourceOrganization != nil {
		if err := r.SourceOrganization.Validate(); err != nil {
			return fmt.Errorf("SourceOrganization: %w", err)
		}
	}
	if r.CollectedPeriod != nil {
		if err := r.CollectedPeriod.Validate(); err != nil {
			return fmt.Errorf("CollectedPeriod: %w", err)
		}
	}
	if r.Procedure != nil {
		if err := r.Procedure.Validate(); err != nil {
			return fmt.Errorf("Procedure: %w", err)
		}
	}
	return nil
}

type BiologicallyDerivedProductProperty struct {
	Id                   *string          `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Type                 *CodeableConcept `json:"type" bson:"type"`                                   // Code that specifies the property
	ValueBoolean         *bool            `json:"valueBoolean" bson:"value_boolean"`                  // Property values
	ValueInteger         *int             `json:"valueInteger" bson:"value_integer"`                  // Property values
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept" bson:"value_codeable_concept"` // Property values
	ValuePeriod          *Period          `json:"valuePeriod" bson:"value_period"`                    // Property values
	ValueQuantity        *Quantity        `json:"valueQuantity" bson:"value_quantity"`                // Property values
	ValueRange           *Range           `json:"valueRange" bson:"value_range"`                      // Property values
	ValueRatio           *Ratio           `json:"valueRatio" bson:"value_ratio"`                      // Property values
	ValueString          *string          `json:"valueString" bson:"value_string"`                    // Property values
	ValueAttachment      *Attachment      `json:"valueAttachment" bson:"value_attachment"`            // Property values
}

func (r *BiologicallyDerivedProductProperty) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.ValueBoolean == nil {
		return fmt.Errorf("field 'ValueBoolean' is required")
	}
	if r.ValueInteger == nil {
		return fmt.Errorf("field 'ValueInteger' is required")
	}
	if r.ValueCodeableConcept == nil {
		return fmt.Errorf("field 'ValueCodeableConcept' is required")
	}
	if r.ValueCodeableConcept != nil {
		if err := r.ValueCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("ValueCodeableConcept: %w", err)
		}
	}
	if r.ValuePeriod == nil {
		return fmt.Errorf("field 'ValuePeriod' is required")
	}
	if r.ValuePeriod != nil {
		if err := r.ValuePeriod.Validate(); err != nil {
			return fmt.Errorf("ValuePeriod: %w", err)
		}
	}
	if r.ValueQuantity == nil {
		return fmt.Errorf("field 'ValueQuantity' is required")
	}
	if r.ValueQuantity != nil {
		if err := r.ValueQuantity.Validate(); err != nil {
			return fmt.Errorf("ValueQuantity: %w", err)
		}
	}
	if r.ValueRange == nil {
		return fmt.Errorf("field 'ValueRange' is required")
	}
	if r.ValueRange != nil {
		if err := r.ValueRange.Validate(); err != nil {
			return fmt.Errorf("ValueRange: %w", err)
		}
	}
	if r.ValueRatio == nil {
		return fmt.Errorf("field 'ValueRatio' is required")
	}
	if r.ValueRatio != nil {
		if err := r.ValueRatio.Validate(); err != nil {
			return fmt.Errorf("ValueRatio: %w", err)
		}
	}
	if r.ValueString == nil {
		return fmt.Errorf("field 'ValueString' is required")
	}
	if r.ValueAttachment == nil {
		return fmt.Errorf("field 'ValueAttachment' is required")
	}
	if r.ValueAttachment != nil {
		if err := r.ValueAttachment.Validate(); err != nil {
			return fmt.Errorf("ValueAttachment: %w", err)
		}
	}
	return nil
}
