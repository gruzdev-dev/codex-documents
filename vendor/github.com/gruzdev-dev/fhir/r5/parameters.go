package models

import (
	"encoding/json"
	"fmt"
)

// This resource is used to pass information into and back from an operation (whether invoked directly from REST or within a messaging environment).  It is not persisted or allowed to be referenced by other resources except as described in the definition of the Parameters resource.
type Parameters struct {
	Id            *string               `json:"id,omitempty" bson:"id,omitempty"`                        // Logical id of this artifact
	Meta          *Meta                 `json:"meta,omitempty" bson:"meta,omitempty"`                    // Metadata about the resource
	ImplicitRules *string               `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"` // A set of rules under which this content was created
	Language      *string               `json:"language,omitempty" bson:"language,omitempty"`            // Language of the resource content
	Parameter     []ParametersParameter `json:"parameter,omitempty" bson:"parameter,omitempty"`          // Operation Parameter
}

func (r *Parameters) Validate() error {
	if r.Meta != nil {
		if err := r.Meta.Validate(); err != nil {
			return fmt.Errorf("Meta: %w", err)
		}
	}
	for i, item := range r.Parameter {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Parameter[%d]: %w", i, err)
		}
	}
	return nil
}

type ParametersParameter struct {
	Id                         *string                `json:"id,omitempty" bson:"id,omitempty"`                                                    // Unique id for inter-element referencing
	Name                       string                 `json:"name" bson:"name"`                                                                    // Name from the definition
	ValueBase64Binary          *string                `json:"valueBase64Binary,omitempty" bson:"value_base64_binary,omitempty"`                    // If parameter is a data type
	ValueBoolean               *bool                  `json:"valueBoolean,omitempty" bson:"value_boolean,omitempty"`                               // If parameter is a data type
	ValueCanonical             *string                `json:"valueCanonical,omitempty" bson:"value_canonical,omitempty"`                           // If parameter is a data type
	ValueCode                  *string                `json:"valueCode,omitempty" bson:"value_code,omitempty"`                                     // If parameter is a data type
	ValueDate                  *string                `json:"valueDate,omitempty" bson:"value_date,omitempty"`                                     // If parameter is a data type
	ValueDateTime              *string                `json:"valueDateTime,omitempty" bson:"value_date_time,omitempty"`                            // If parameter is a data type
	ValueDecimal               *float64               `json:"valueDecimal,omitempty" bson:"value_decimal,omitempty"`                               // If parameter is a data type
	ValueId                    *string                `json:"valueId,omitempty" bson:"value_id,omitempty"`                                         // If parameter is a data type
	ValueInstant               *string                `json:"valueInstant,omitempty" bson:"value_instant,omitempty"`                               // If parameter is a data type
	ValueInteger               *int                   `json:"valueInteger,omitempty" bson:"value_integer,omitempty"`                               // If parameter is a data type
	ValueInteger64             *int64                 `json:"valueInteger64,omitempty" bson:"value_integer64,omitempty"`                           // If parameter is a data type
	ValueMarkdown              *string                `json:"valueMarkdown,omitempty" bson:"value_markdown,omitempty"`                             // If parameter is a data type
	ValueOid                   *string                `json:"valueOid,omitempty" bson:"value_oid,omitempty"`                                       // If parameter is a data type
	ValuePositiveInt           *int                   `json:"valuePositiveInt,omitempty" bson:"value_positive_int,omitempty"`                      // If parameter is a data type
	ValueString                *string                `json:"valueString,omitempty" bson:"value_string,omitempty"`                                 // If parameter is a data type
	ValueTime                  *string                `json:"valueTime,omitempty" bson:"value_time,omitempty"`                                     // If parameter is a data type
	ValueUnsignedInt           *int                   `json:"valueUnsignedInt,omitempty" bson:"value_unsigned_int,omitempty"`                      // If parameter is a data type
	ValueUri                   *string                `json:"valueUri,omitempty" bson:"value_uri,omitempty"`                                       // If parameter is a data type
	ValueUrl                   *string                `json:"valueUrl,omitempty" bson:"value_url,omitempty"`                                       // If parameter is a data type
	ValueUuid                  *uuid                  `json:"valueUuid,omitempty" bson:"value_uuid,omitempty"`                                     // If parameter is a data type
	ValueAddress               *Address               `json:"valueAddress,omitempty" bson:"value_address,omitempty"`                               // If parameter is a data type
	ValueAge                   *Age                   `json:"valueAge,omitempty" bson:"value_age,omitempty"`                                       // If parameter is a data type
	ValueAnnotation            *Annotation            `json:"valueAnnotation,omitempty" bson:"value_annotation,omitempty"`                         // If parameter is a data type
	ValueAttachment            *Attachment            `json:"valueAttachment,omitempty" bson:"value_attachment,omitempty"`                         // If parameter is a data type
	ValueCodeableConcept       *CodeableConcept       `json:"valueCodeableConcept,omitempty" bson:"value_codeable_concept,omitempty"`              // If parameter is a data type
	ValueCodeableReference     *CodeableReference     `json:"valueCodeableReference,omitempty" bson:"value_codeable_reference,omitempty"`          // If parameter is a data type
	ValueCoding                *Coding                `json:"valueCoding,omitempty" bson:"value_coding,omitempty"`                                 // If parameter is a data type
	ValueContactPoint          *ContactPoint          `json:"valueContactPoint,omitempty" bson:"value_contact_point,omitempty"`                    // If parameter is a data type
	ValueCount                 *Count                 `json:"valueCount,omitempty" bson:"value_count,omitempty"`                                   // If parameter is a data type
	ValueDistance              *Distance              `json:"valueDistance,omitempty" bson:"value_distance,omitempty"`                             // If parameter is a data type
	ValueDuration              *Duration              `json:"valueDuration,omitempty" bson:"value_duration,omitempty"`                             // If parameter is a data type
	ValueHumanName             *HumanName             `json:"valueHumanName,omitempty" bson:"value_human_name,omitempty"`                          // If parameter is a data type
	ValueIdentifier            *Identifier            `json:"valueIdentifier,omitempty" bson:"value_identifier,omitempty"`                         // If parameter is a data type
	ValueMoney                 *Money                 `json:"valueMoney,omitempty" bson:"value_money,omitempty"`                                   // If parameter is a data type
	ValuePeriod                *Period                `json:"valuePeriod,omitempty" bson:"value_period,omitempty"`                                 // If parameter is a data type
	ValueQuantity              *Quantity              `json:"valueQuantity,omitempty" bson:"value_quantity,omitempty"`                             // If parameter is a data type
	ValueRange                 *Range                 `json:"valueRange,omitempty" bson:"value_range,omitempty"`                                   // If parameter is a data type
	ValueRatio                 *Ratio                 `json:"valueRatio,omitempty" bson:"value_ratio,omitempty"`                                   // If parameter is a data type
	ValueRatioRange            *RatioRange            `json:"valueRatioRange,omitempty" bson:"value_ratio_range,omitempty"`                        // If parameter is a data type
	ValueReference             *Reference             `json:"valueReference,omitempty" bson:"value_reference,omitempty"`                           // If parameter is a data type
	ValueSampledData           *SampledData           `json:"valueSampledData,omitempty" bson:"value_sampled_data,omitempty"`                      // If parameter is a data type
	ValueSignature             *Signature             `json:"valueSignature,omitempty" bson:"value_signature,omitempty"`                           // If parameter is a data type
	ValueTiming                *Timing                `json:"valueTiming,omitempty" bson:"value_timing,omitempty"`                                 // If parameter is a data type
	ValueContactDetail         *ContactDetail         `json:"valueContactDetail,omitempty" bson:"value_contact_detail,omitempty"`                  // If parameter is a data type
	ValueDataRequirement       *DataRequirement       `json:"valueDataRequirement,omitempty" bson:"value_data_requirement,omitempty"`              // If parameter is a data type
	ValueExpression            *Expression            `json:"valueExpression,omitempty" bson:"value_expression,omitempty"`                         // If parameter is a data type
	ValueParameterDefinition   *ParameterDefinition   `json:"valueParameterDefinition,omitempty" bson:"value_parameter_definition,omitempty"`      // If parameter is a data type
	ValueRelatedArtifact       *RelatedArtifact       `json:"valueRelatedArtifact,omitempty" bson:"value_related_artifact,omitempty"`              // If parameter is a data type
	ValueTriggerDefinition     *TriggerDefinition     `json:"valueTriggerDefinition,omitempty" bson:"value_trigger_definition,omitempty"`          // If parameter is a data type
	ValueUsageContext          *UsageContext          `json:"valueUsageContext,omitempty" bson:"value_usage_context,omitempty"`                    // If parameter is a data type
	ValueAvailability          *Availability          `json:"valueAvailability,omitempty" bson:"value_availability,omitempty"`                     // If parameter is a data type
	ValueExtendedContactDetail *ExtendedContactDetail `json:"valueExtendedContactDetail,omitempty" bson:"value_extended_contact_detail,omitempty"` // If parameter is a data type
	ValueVirtualServiceDetail  *VirtualServiceDetail  `json:"valueVirtualServiceDetail,omitempty" bson:"value_virtual_service_detail,omitempty"`   // If parameter is a data type
	ValueDosage                *Dosage                `json:"valueDosage,omitempty" bson:"value_dosage,omitempty"`                                 // If parameter is a data type
	ValueMeta                  *Meta                  `json:"valueMeta,omitempty" bson:"value_meta,omitempty"`                                     // If parameter is a data type
	Resource                   json.RawMessage        `json:"resource,omitempty" bson:"resource,omitempty"`                                        // If parameter is a whole resource
	Part                       []ParametersParameter  `json:"part,omitempty" bson:"part,omitempty"`                                                // Named part of a multi-part parameter
}

func (r *ParametersParameter) Validate() error {
	var emptyString string
	if r.Name == emptyString {
		return fmt.Errorf("field 'Name' is required")
	}
	if r.ValueUuid != nil {
		if err := r.ValueUuid.Validate(); err != nil {
			return fmt.Errorf("ValueUuid: %w", err)
		}
	}
	if r.ValueAddress != nil {
		if err := r.ValueAddress.Validate(); err != nil {
			return fmt.Errorf("ValueAddress: %w", err)
		}
	}
	if r.ValueAge != nil {
		if err := r.ValueAge.Validate(); err != nil {
			return fmt.Errorf("ValueAge: %w", err)
		}
	}
	if r.ValueAnnotation != nil {
		if err := r.ValueAnnotation.Validate(); err != nil {
			return fmt.Errorf("ValueAnnotation: %w", err)
		}
	}
	if r.ValueAttachment != nil {
		if err := r.ValueAttachment.Validate(); err != nil {
			return fmt.Errorf("ValueAttachment: %w", err)
		}
	}
	if r.ValueCodeableConcept != nil {
		if err := r.ValueCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("ValueCodeableConcept: %w", err)
		}
	}
	if r.ValueCodeableReference != nil {
		if err := r.ValueCodeableReference.Validate(); err != nil {
			return fmt.Errorf("ValueCodeableReference: %w", err)
		}
	}
	if r.ValueCoding != nil {
		if err := r.ValueCoding.Validate(); err != nil {
			return fmt.Errorf("ValueCoding: %w", err)
		}
	}
	if r.ValueContactPoint != nil {
		if err := r.ValueContactPoint.Validate(); err != nil {
			return fmt.Errorf("ValueContactPoint: %w", err)
		}
	}
	if r.ValueCount != nil {
		if err := r.ValueCount.Validate(); err != nil {
			return fmt.Errorf("ValueCount: %w", err)
		}
	}
	if r.ValueDistance != nil {
		if err := r.ValueDistance.Validate(); err != nil {
			return fmt.Errorf("ValueDistance: %w", err)
		}
	}
	if r.ValueDuration != nil {
		if err := r.ValueDuration.Validate(); err != nil {
			return fmt.Errorf("ValueDuration: %w", err)
		}
	}
	if r.ValueHumanName != nil {
		if err := r.ValueHumanName.Validate(); err != nil {
			return fmt.Errorf("ValueHumanName: %w", err)
		}
	}
	if r.ValueIdentifier != nil {
		if err := r.ValueIdentifier.Validate(); err != nil {
			return fmt.Errorf("ValueIdentifier: %w", err)
		}
	}
	if r.ValueMoney != nil {
		if err := r.ValueMoney.Validate(); err != nil {
			return fmt.Errorf("ValueMoney: %w", err)
		}
	}
	if r.ValuePeriod != nil {
		if err := r.ValuePeriod.Validate(); err != nil {
			return fmt.Errorf("ValuePeriod: %w", err)
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
	if r.ValueRatio != nil {
		if err := r.ValueRatio.Validate(); err != nil {
			return fmt.Errorf("ValueRatio: %w", err)
		}
	}
	if r.ValueRatioRange != nil {
		if err := r.ValueRatioRange.Validate(); err != nil {
			return fmt.Errorf("ValueRatioRange: %w", err)
		}
	}
	if r.ValueReference != nil {
		if err := r.ValueReference.Validate(); err != nil {
			return fmt.Errorf("ValueReference: %w", err)
		}
	}
	if r.ValueSampledData != nil {
		if err := r.ValueSampledData.Validate(); err != nil {
			return fmt.Errorf("ValueSampledData: %w", err)
		}
	}
	if r.ValueSignature != nil {
		if err := r.ValueSignature.Validate(); err != nil {
			return fmt.Errorf("ValueSignature: %w", err)
		}
	}
	if r.ValueTiming != nil {
		if err := r.ValueTiming.Validate(); err != nil {
			return fmt.Errorf("ValueTiming: %w", err)
		}
	}
	if r.ValueContactDetail != nil {
		if err := r.ValueContactDetail.Validate(); err != nil {
			return fmt.Errorf("ValueContactDetail: %w", err)
		}
	}
	if r.ValueDataRequirement != nil {
		if err := r.ValueDataRequirement.Validate(); err != nil {
			return fmt.Errorf("ValueDataRequirement: %w", err)
		}
	}
	if r.ValueExpression != nil {
		if err := r.ValueExpression.Validate(); err != nil {
			return fmt.Errorf("ValueExpression: %w", err)
		}
	}
	if r.ValueParameterDefinition != nil {
		if err := r.ValueParameterDefinition.Validate(); err != nil {
			return fmt.Errorf("ValueParameterDefinition: %w", err)
		}
	}
	if r.ValueRelatedArtifact != nil {
		if err := r.ValueRelatedArtifact.Validate(); err != nil {
			return fmt.Errorf("ValueRelatedArtifact: %w", err)
		}
	}
	if r.ValueTriggerDefinition != nil {
		if err := r.ValueTriggerDefinition.Validate(); err != nil {
			return fmt.Errorf("ValueTriggerDefinition: %w", err)
		}
	}
	if r.ValueUsageContext != nil {
		if err := r.ValueUsageContext.Validate(); err != nil {
			return fmt.Errorf("ValueUsageContext: %w", err)
		}
	}
	if r.ValueAvailability != nil {
		if err := r.ValueAvailability.Validate(); err != nil {
			return fmt.Errorf("ValueAvailability: %w", err)
		}
	}
	if r.ValueExtendedContactDetail != nil {
		if err := r.ValueExtendedContactDetail.Validate(); err != nil {
			return fmt.Errorf("ValueExtendedContactDetail: %w", err)
		}
	}
	if r.ValueVirtualServiceDetail != nil {
		if err := r.ValueVirtualServiceDetail.Validate(); err != nil {
			return fmt.Errorf("ValueVirtualServiceDetail: %w", err)
		}
	}
	if r.ValueDosage != nil {
		if err := r.ValueDosage.Validate(); err != nil {
			return fmt.Errorf("ValueDosage: %w", err)
		}
	}
	if r.ValueMeta != nil {
		if err := r.ValueMeta.Validate(); err != nil {
			return fmt.Errorf("ValueMeta: %w", err)
		}
	}
	for i, item := range r.Part {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Part[%d]: %w", i, err)
		}
	}
	return nil
}
