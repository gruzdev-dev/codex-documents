package models

import (
	"fmt"
)

// Extension Type: Optional Extension Element - found in all resources.
type Extension struct {
	Id                         *string                `json:"id,omitempty" bson:"id,omitempty"`                                                    // Unique id for inter-element referencing
	Url                        string                 `json:"url" bson:"url"`                                                                      // identifies the meaning of the extension
	ValueBase64Binary          *string                `json:"valueBase64Binary,omitempty" bson:"value_base64_binary,omitempty"`                    // Value of extension
	ValueBoolean               *bool                  `json:"valueBoolean,omitempty" bson:"value_boolean,omitempty"`                               // Value of extension
	ValueCanonical             *string                `json:"valueCanonical,omitempty" bson:"value_canonical,omitempty"`                           // Value of extension
	ValueCode                  *string                `json:"valueCode,omitempty" bson:"value_code,omitempty"`                                     // Value of extension
	ValueDate                  *string                `json:"valueDate,omitempty" bson:"value_date,omitempty"`                                     // Value of extension
	ValueDateTime              *string                `json:"valueDateTime,omitempty" bson:"value_date_time,omitempty"`                            // Value of extension
	ValueDecimal               *float64               `json:"valueDecimal,omitempty" bson:"value_decimal,omitempty"`                               // Value of extension
	ValueId                    *string                `json:"valueId,omitempty" bson:"value_id,omitempty"`                                         // Value of extension
	ValueInstant               *string                `json:"valueInstant,omitempty" bson:"value_instant,omitempty"`                               // Value of extension
	ValueInteger               *int                   `json:"valueInteger,omitempty" bson:"value_integer,omitempty"`                               // Value of extension
	ValueInteger64             *int64                 `json:"valueInteger64,omitempty" bson:"value_integer64,omitempty"`                           // Value of extension
	ValueMarkdown              *string                `json:"valueMarkdown,omitempty" bson:"value_markdown,omitempty"`                             // Value of extension
	ValueOid                   *string                `json:"valueOid,omitempty" bson:"value_oid,omitempty"`                                       // Value of extension
	ValuePositiveInt           *int                   `json:"valuePositiveInt,omitempty" bson:"value_positive_int,omitempty"`                      // Value of extension
	ValueString                *string                `json:"valueString,omitempty" bson:"value_string,omitempty"`                                 // Value of extension
	ValueTime                  *string                `json:"valueTime,omitempty" bson:"value_time,omitempty"`                                     // Value of extension
	ValueUnsignedInt           *int                   `json:"valueUnsignedInt,omitempty" bson:"value_unsigned_int,omitempty"`                      // Value of extension
	ValueUri                   *string                `json:"valueUri,omitempty" bson:"value_uri,omitempty"`                                       // Value of extension
	ValueUrl                   *string                `json:"valueUrl,omitempty" bson:"value_url,omitempty"`                                       // Value of extension
	ValueUuid                  *uuid                  `json:"valueUuid,omitempty" bson:"value_uuid,omitempty"`                                     // Value of extension
	ValueAddress               *Address               `json:"valueAddress,omitempty" bson:"value_address,omitempty"`                               // Value of extension
	ValueAge                   *Age                   `json:"valueAge,omitempty" bson:"value_age,omitempty"`                                       // Value of extension
	ValueAnnotation            *Annotation            `json:"valueAnnotation,omitempty" bson:"value_annotation,omitempty"`                         // Value of extension
	ValueAttachment            *Attachment            `json:"valueAttachment,omitempty" bson:"value_attachment,omitempty"`                         // Value of extension
	ValueCodeableConcept       *CodeableConcept       `json:"valueCodeableConcept,omitempty" bson:"value_codeable_concept,omitempty"`              // Value of extension
	ValueCodeableReference     *CodeableReference     `json:"valueCodeableReference,omitempty" bson:"value_codeable_reference,omitempty"`          // Value of extension
	ValueCoding                *Coding                `json:"valueCoding,omitempty" bson:"value_coding,omitempty"`                                 // Value of extension
	ValueContactPoint          *ContactPoint          `json:"valueContactPoint,omitempty" bson:"value_contact_point,omitempty"`                    // Value of extension
	ValueCount                 *Count                 `json:"valueCount,omitempty" bson:"value_count,omitempty"`                                   // Value of extension
	ValueDistance              *Distance              `json:"valueDistance,omitempty" bson:"value_distance,omitempty"`                             // Value of extension
	ValueDuration              *Duration              `json:"valueDuration,omitempty" bson:"value_duration,omitempty"`                             // Value of extension
	ValueHumanName             *HumanName             `json:"valueHumanName,omitempty" bson:"value_human_name,omitempty"`                          // Value of extension
	ValueIdentifier            *Identifier            `json:"valueIdentifier,omitempty" bson:"value_identifier,omitempty"`                         // Value of extension
	ValueMoney                 *Money                 `json:"valueMoney,omitempty" bson:"value_money,omitempty"`                                   // Value of extension
	ValuePeriod                *Period                `json:"valuePeriod,omitempty" bson:"value_period,omitempty"`                                 // Value of extension
	ValueQuantity              *Quantity              `json:"valueQuantity,omitempty" bson:"value_quantity,omitempty"`                             // Value of extension
	ValueRange                 *Range                 `json:"valueRange,omitempty" bson:"value_range,omitempty"`                                   // Value of extension
	ValueRatio                 *Ratio                 `json:"valueRatio,omitempty" bson:"value_ratio,omitempty"`                                   // Value of extension
	ValueRatioRange            *RatioRange            `json:"valueRatioRange,omitempty" bson:"value_ratio_range,omitempty"`                        // Value of extension
	ValueReference             *Reference             `json:"valueReference,omitempty" bson:"value_reference,omitempty"`                           // Value of extension
	ValueSampledData           *SampledData           `json:"valueSampledData,omitempty" bson:"value_sampled_data,omitempty"`                      // Value of extension
	ValueSignature             *Signature             `json:"valueSignature,omitempty" bson:"value_signature,omitempty"`                           // Value of extension
	ValueTiming                *Timing                `json:"valueTiming,omitempty" bson:"value_timing,omitempty"`                                 // Value of extension
	ValueContactDetail         *ContactDetail         `json:"valueContactDetail,omitempty" bson:"value_contact_detail,omitempty"`                  // Value of extension
	ValueDataRequirement       *DataRequirement       `json:"valueDataRequirement,omitempty" bson:"value_data_requirement,omitempty"`              // Value of extension
	ValueExpression            *Expression            `json:"valueExpression,omitempty" bson:"value_expression,omitempty"`                         // Value of extension
	ValueParameterDefinition   *ParameterDefinition   `json:"valueParameterDefinition,omitempty" bson:"value_parameter_definition,omitempty"`      // Value of extension
	ValueRelatedArtifact       *RelatedArtifact       `json:"valueRelatedArtifact,omitempty" bson:"value_related_artifact,omitempty"`              // Value of extension
	ValueTriggerDefinition     *TriggerDefinition     `json:"valueTriggerDefinition,omitempty" bson:"value_trigger_definition,omitempty"`          // Value of extension
	ValueUsageContext          *UsageContext          `json:"valueUsageContext,omitempty" bson:"value_usage_context,omitempty"`                    // Value of extension
	ValueAvailability          *Availability          `json:"valueAvailability,omitempty" bson:"value_availability,omitempty"`                     // Value of extension
	ValueExtendedContactDetail *ExtendedContactDetail `json:"valueExtendedContactDetail,omitempty" bson:"value_extended_contact_detail,omitempty"` // Value of extension
	ValueVirtualServiceDetail  *VirtualServiceDetail  `json:"valueVirtualServiceDetail,omitempty" bson:"value_virtual_service_detail,omitempty"`   // Value of extension
	ValueDosage                *Dosage                `json:"valueDosage,omitempty" bson:"value_dosage,omitempty"`                                 // Value of extension
	ValueMeta                  *Meta                  `json:"valueMeta,omitempty" bson:"value_meta,omitempty"`                                     // Value of extension
}

func (r *Extension) Validate() error {
	var emptyString string
	if r.Url == emptyString {
		return fmt.Errorf("field 'Url' is required")
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
	return nil
}
