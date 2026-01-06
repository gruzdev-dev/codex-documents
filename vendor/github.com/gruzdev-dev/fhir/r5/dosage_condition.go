package models

import (
	"fmt"
)

// DosageCondition Type: DosageCondition expresses a time or time period as relative to the time of an event defined in data types other than dateTime.
type DosageCondition struct {
	Id                         *string                `json:"id,omitempty" bson:"id,omitempty"`                                                    // Unique id for inter-element referencing
	Code                       *CodeableConcept       `json:"code" bson:"code"`                                                                    // The specific event occurrence or resource context used as a base point (reference point) in time
	Details                    *CodeableConcept       `json:"details,omitempty" bson:"details,omitempty"`                                          // Additional details about the event - depends on the code
	Operation                  *string                `json:"operation,omitempty" bson:"operation,omitempty"`                                      // eq | ne | in | nin | gt | lt | ge | le | sa | eb | ap
	ValueBase64Binary          *string                `json:"valueBase64Binary,omitempty" bson:"value_base64_binary,omitempty"`                    // The value for this critera
	ValueBoolean               *bool                  `json:"valueBoolean,omitempty" bson:"value_boolean,omitempty"`                               // The value for this critera
	ValueCanonical             *string                `json:"valueCanonical,omitempty" bson:"value_canonical,omitempty"`                           // The value for this critera
	ValueCode                  *string                `json:"valueCode,omitempty" bson:"value_code,omitempty"`                                     // The value for this critera
	ValueDate                  *string                `json:"valueDate,omitempty" bson:"value_date,omitempty"`                                     // The value for this critera
	ValueDateTime              *string                `json:"valueDateTime,omitempty" bson:"value_date_time,omitempty"`                            // The value for this critera
	ValueDecimal               *float64               `json:"valueDecimal,omitempty" bson:"value_decimal,omitempty"`                               // The value for this critera
	ValueId                    *string                `json:"valueId,omitempty" bson:"value_id,omitempty"`                                         // The value for this critera
	ValueInstant               *string                `json:"valueInstant,omitempty" bson:"value_instant,omitempty"`                               // The value for this critera
	ValueInteger               *int                   `json:"valueInteger,omitempty" bson:"value_integer,omitempty"`                               // The value for this critera
	ValueInteger64             *int64                 `json:"valueInteger64,omitempty" bson:"value_integer64,omitempty"`                           // The value for this critera
	ValueMarkdown              *string                `json:"valueMarkdown,omitempty" bson:"value_markdown,omitempty"`                             // The value for this critera
	ValueOid                   *string                `json:"valueOid,omitempty" bson:"value_oid,omitempty"`                                       // The value for this critera
	ValuePositiveInt           *int                   `json:"valuePositiveInt,omitempty" bson:"value_positive_int,omitempty"`                      // The value for this critera
	ValueString                *string                `json:"valueString,omitempty" bson:"value_string,omitempty"`                                 // The value for this critera
	ValueTime                  *string                `json:"valueTime,omitempty" bson:"value_time,omitempty"`                                     // The value for this critera
	ValueUnsignedInt           *int                   `json:"valueUnsignedInt,omitempty" bson:"value_unsigned_int,omitempty"`                      // The value for this critera
	ValueUri                   *string                `json:"valueUri,omitempty" bson:"value_uri,omitempty"`                                       // The value for this critera
	ValueUrl                   *string                `json:"valueUrl,omitempty" bson:"value_url,omitempty"`                                       // The value for this critera
	ValueUuid                  *uuid                  `json:"valueUuid,omitempty" bson:"value_uuid,omitempty"`                                     // The value for this critera
	ValueAddress               *Address               `json:"valueAddress,omitempty" bson:"value_address,omitempty"`                               // The value for this critera
	ValueAge                   *Age                   `json:"valueAge,omitempty" bson:"value_age,omitempty"`                                       // The value for this critera
	ValueAnnotation            *Annotation            `json:"valueAnnotation,omitempty" bson:"value_annotation,omitempty"`                         // The value for this critera
	ValueAttachment            *Attachment            `json:"valueAttachment,omitempty" bson:"value_attachment,omitempty"`                         // The value for this critera
	ValueCodeableConcept       *CodeableConcept       `json:"valueCodeableConcept,omitempty" bson:"value_codeable_concept,omitempty"`              // The value for this critera
	ValueCodeableReference     *CodeableReference     `json:"valueCodeableReference,omitempty" bson:"value_codeable_reference,omitempty"`          // The value for this critera
	ValueCoding                *Coding                `json:"valueCoding,omitempty" bson:"value_coding,omitempty"`                                 // The value for this critera
	ValueContactPoint          *ContactPoint          `json:"valueContactPoint,omitempty" bson:"value_contact_point,omitempty"`                    // The value for this critera
	ValueCount                 *Count                 `json:"valueCount,omitempty" bson:"value_count,omitempty"`                                   // The value for this critera
	ValueDistance              *Distance              `json:"valueDistance,omitempty" bson:"value_distance,omitempty"`                             // The value for this critera
	ValueDuration              *Duration              `json:"valueDuration,omitempty" bson:"value_duration,omitempty"`                             // The value for this critera
	ValueHumanName             *HumanName             `json:"valueHumanName,omitempty" bson:"value_human_name,omitempty"`                          // The value for this critera
	ValueIdentifier            *Identifier            `json:"valueIdentifier,omitempty" bson:"value_identifier,omitempty"`                         // The value for this critera
	ValueMoney                 *Money                 `json:"valueMoney,omitempty" bson:"value_money,omitempty"`                                   // The value for this critera
	ValuePeriod                *Period                `json:"valuePeriod,omitempty" bson:"value_period,omitempty"`                                 // The value for this critera
	ValueQuantity              *Quantity              `json:"valueQuantity,omitempty" bson:"value_quantity,omitempty"`                             // The value for this critera
	ValueRange                 *Range                 `json:"valueRange,omitempty" bson:"value_range,omitempty"`                                   // The value for this critera
	ValueRatio                 *Ratio                 `json:"valueRatio,omitempty" bson:"value_ratio,omitempty"`                                   // The value for this critera
	ValueRatioRange            *RatioRange            `json:"valueRatioRange,omitempty" bson:"value_ratio_range,omitempty"`                        // The value for this critera
	ValueReference             *Reference             `json:"valueReference,omitempty" bson:"value_reference,omitempty"`                           // The value for this critera
	ValueSampledData           *SampledData           `json:"valueSampledData,omitempty" bson:"value_sampled_data,omitempty"`                      // The value for this critera
	ValueSignature             *Signature             `json:"valueSignature,omitempty" bson:"value_signature,omitempty"`                           // The value for this critera
	ValueTiming                *Timing                `json:"valueTiming,omitempty" bson:"value_timing,omitempty"`                                 // The value for this critera
	ValueContactDetail         *ContactDetail         `json:"valueContactDetail,omitempty" bson:"value_contact_detail,omitempty"`                  // The value for this critera
	ValueDataRequirement       *DataRequirement       `json:"valueDataRequirement,omitempty" bson:"value_data_requirement,omitempty"`              // The value for this critera
	ValueExpression            *Expression            `json:"valueExpression,omitempty" bson:"value_expression,omitempty"`                         // The value for this critera
	ValueParameterDefinition   *ParameterDefinition   `json:"valueParameterDefinition,omitempty" bson:"value_parameter_definition,omitempty"`      // The value for this critera
	ValueRelatedArtifact       *RelatedArtifact       `json:"valueRelatedArtifact,omitempty" bson:"value_related_artifact,omitempty"`              // The value for this critera
	ValueTriggerDefinition     *TriggerDefinition     `json:"valueTriggerDefinition,omitempty" bson:"value_trigger_definition,omitempty"`          // The value for this critera
	ValueUsageContext          *UsageContext          `json:"valueUsageContext,omitempty" bson:"value_usage_context,omitempty"`                    // The value for this critera
	ValueAvailability          *Availability          `json:"valueAvailability,omitempty" bson:"value_availability,omitempty"`                     // The value for this critera
	ValueExtendedContactDetail *ExtendedContactDetail `json:"valueExtendedContactDetail,omitempty" bson:"value_extended_contact_detail,omitempty"` // The value for this critera
	ValueVirtualServiceDetail  *VirtualServiceDetail  `json:"valueVirtualServiceDetail,omitempty" bson:"value_virtual_service_detail,omitempty"`   // The value for this critera
	ValueDosage                *Dosage                `json:"valueDosage,omitempty" bson:"value_dosage,omitempty"`                                 // The value for this critera
	ValueMeta                  *Meta                  `json:"valueMeta,omitempty" bson:"value_meta,omitempty"`                                     // The value for this critera
	Text                       *string                `json:"text,omitempty" bson:"text,omitempty"`                                                // Free-text description
}

func (r *DosageCondition) Validate() error {
	if r.Code == nil {
		return fmt.Errorf("field 'Code' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.Details != nil {
		if err := r.Details.Validate(); err != nil {
			return fmt.Errorf("Details: %w", err)
		}
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
