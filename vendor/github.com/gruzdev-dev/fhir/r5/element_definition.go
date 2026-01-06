package models

import (
	"fmt"
)

// ElementDefinition Type: Captures constraints on each element within the resource, profile, or extension.
type ElementDefinition struct {
	Id                                *string                       `json:"id,omitempty" bson:"id,omitempty"`                                                                   // Unique id for inter-element referencing
	Path                              string                        `json:"path" bson:"path"`                                                                                   // Path of the element in the hierarchy of elements
	Representation                    []string                      `json:"representation,omitempty" bson:"representation,omitempty"`                                           // xmlAttr | xmlText | typeAttr | cdaText | xhtml
	SliceName                         *string                       `json:"sliceName,omitempty" bson:"slice_name,omitempty"`                                                    // Name for this particular element (in a set of slices)
	SliceIsConstraining               bool                          `json:"sliceIsConstraining,omitempty" bson:"slice_is_constraining,omitempty"`                               // If this slice definition constrains an inherited slice definition (or not)
	Label                             *string                       `json:"label,omitempty" bson:"label,omitempty"`                                                             // String to display with or prompt for element
	Code                              []Coding                      `json:"code,omitempty" bson:"code,omitempty"`                                                               // Corresponding codes in terminologies
	Slicing                           *ElementDefinitionSlicing     `json:"slicing,omitempty" bson:"slicing,omitempty"`                                                         // This element is sliced - slices follow
	Short                             *string                       `json:"short,omitempty" bson:"short,omitempty"`                                                             // Concise definition for space-constrained presentation
	Definition                        *string                       `json:"definition,omitempty" bson:"definition,omitempty"`                                                   // Full formal definition as narrative text
	Comment                           *string                       `json:"comment,omitempty" bson:"comment,omitempty"`                                                         // Comments about the use of this element
	Requirements                      *string                       `json:"requirements,omitempty" bson:"requirements,omitempty"`                                               // Requirements satisfied by this element/structure and its constraints
	Alias                             []string                      `json:"alias,omitempty" bson:"alias,omitempty"`                                                             // Other names
	Min                               *int                          `json:"min,omitempty" bson:"min,omitempty"`                                                                 // Minimum Cardinality
	Max                               *string                       `json:"max,omitempty" bson:"max,omitempty"`                                                                 // Maximum Cardinality (a number or *)
	Base                              *ElementDefinitionBase        `json:"base,omitempty" bson:"base,omitempty"`                                                               // Base definition information for tools
	ContentReference                  *string                       `json:"contentReference,omitempty" bson:"content_reference,omitempty"`                                      // Reference to definition of content for the element
	Type                              []ElementDefinitionType       `json:"type,omitempty" bson:"type,omitempty"`                                                               // Data type and Profile for this element
	DefaultValueBase64Binary          *string                       `json:"defaultValueBase64Binary,omitempty" bson:"default_value_base64_binary,omitempty"`                    // Specified value if missing from instance
	DefaultValueBoolean               *bool                         `json:"defaultValueBoolean,omitempty" bson:"default_value_boolean,omitempty"`                               // Specified value if missing from instance
	DefaultValueCanonical             *string                       `json:"defaultValueCanonical,omitempty" bson:"default_value_canonical,omitempty"`                           // Specified value if missing from instance
	DefaultValueCode                  *string                       `json:"defaultValueCode,omitempty" bson:"default_value_code,omitempty"`                                     // Specified value if missing from instance
	DefaultValueDate                  *string                       `json:"defaultValueDate,omitempty" bson:"default_value_date,omitempty"`                                     // Specified value if missing from instance
	DefaultValueDateTime              *string                       `json:"defaultValueDateTime,omitempty" bson:"default_value_date_time,omitempty"`                            // Specified value if missing from instance
	DefaultValueDecimal               *float64                      `json:"defaultValueDecimal,omitempty" bson:"default_value_decimal,omitempty"`                               // Specified value if missing from instance
	DefaultValueId                    *string                       `json:"defaultValueId,omitempty" bson:"default_value_id,omitempty"`                                         // Specified value if missing from instance
	DefaultValueInstant               *string                       `json:"defaultValueInstant,omitempty" bson:"default_value_instant,omitempty"`                               // Specified value if missing from instance
	DefaultValueInteger               *int                          `json:"defaultValueInteger,omitempty" bson:"default_value_integer,omitempty"`                               // Specified value if missing from instance
	DefaultValueInteger64             *int64                        `json:"defaultValueInteger64,omitempty" bson:"default_value_integer64,omitempty"`                           // Specified value if missing from instance
	DefaultValueMarkdown              *string                       `json:"defaultValueMarkdown,omitempty" bson:"default_value_markdown,omitempty"`                             // Specified value if missing from instance
	DefaultValueOid                   *string                       `json:"defaultValueOid,omitempty" bson:"default_value_oid,omitempty"`                                       // Specified value if missing from instance
	DefaultValuePositiveInt           *int                          `json:"defaultValuePositiveInt,omitempty" bson:"default_value_positive_int,omitempty"`                      // Specified value if missing from instance
	DefaultValueString                *string                       `json:"defaultValueString,omitempty" bson:"default_value_string,omitempty"`                                 // Specified value if missing from instance
	DefaultValueTime                  *string                       `json:"defaultValueTime,omitempty" bson:"default_value_time,omitempty"`                                     // Specified value if missing from instance
	DefaultValueUnsignedInt           *int                          `json:"defaultValueUnsignedInt,omitempty" bson:"default_value_unsigned_int,omitempty"`                      // Specified value if missing from instance
	DefaultValueUri                   *string                       `json:"defaultValueUri,omitempty" bson:"default_value_uri,omitempty"`                                       // Specified value if missing from instance
	DefaultValueUrl                   *string                       `json:"defaultValueUrl,omitempty" bson:"default_value_url,omitempty"`                                       // Specified value if missing from instance
	DefaultValueUuid                  *uuid                         `json:"defaultValueUuid,omitempty" bson:"default_value_uuid,omitempty"`                                     // Specified value if missing from instance
	DefaultValueAddress               *Address                      `json:"defaultValueAddress,omitempty" bson:"default_value_address,omitempty"`                               // Specified value if missing from instance
	DefaultValueAge                   *Age                          `json:"defaultValueAge,omitempty" bson:"default_value_age,omitempty"`                                       // Specified value if missing from instance
	DefaultValueAnnotation            *Annotation                   `json:"defaultValueAnnotation,omitempty" bson:"default_value_annotation,omitempty"`                         // Specified value if missing from instance
	DefaultValueAttachment            *Attachment                   `json:"defaultValueAttachment,omitempty" bson:"default_value_attachment,omitempty"`                         // Specified value if missing from instance
	DefaultValueCodeableConcept       *CodeableConcept              `json:"defaultValueCodeableConcept,omitempty" bson:"default_value_codeable_concept,omitempty"`              // Specified value if missing from instance
	DefaultValueCodeableReference     *CodeableReference            `json:"defaultValueCodeableReference,omitempty" bson:"default_value_codeable_reference,omitempty"`          // Specified value if missing from instance
	DefaultValueCoding                *Coding                       `json:"defaultValueCoding,omitempty" bson:"default_value_coding,omitempty"`                                 // Specified value if missing from instance
	DefaultValueContactPoint          *ContactPoint                 `json:"defaultValueContactPoint,omitempty" bson:"default_value_contact_point,omitempty"`                    // Specified value if missing from instance
	DefaultValueCount                 *Count                        `json:"defaultValueCount,omitempty" bson:"default_value_count,omitempty"`                                   // Specified value if missing from instance
	DefaultValueDistance              *Distance                     `json:"defaultValueDistance,omitempty" bson:"default_value_distance,omitempty"`                             // Specified value if missing from instance
	DefaultValueDuration              *Duration                     `json:"defaultValueDuration,omitempty" bson:"default_value_duration,omitempty"`                             // Specified value if missing from instance
	DefaultValueHumanName             *HumanName                    `json:"defaultValueHumanName,omitempty" bson:"default_value_human_name,omitempty"`                          // Specified value if missing from instance
	DefaultValueIdentifier            *Identifier                   `json:"defaultValueIdentifier,omitempty" bson:"default_value_identifier,omitempty"`                         // Specified value if missing from instance
	DefaultValueMoney                 *Money                        `json:"defaultValueMoney,omitempty" bson:"default_value_money,omitempty"`                                   // Specified value if missing from instance
	DefaultValuePeriod                *Period                       `json:"defaultValuePeriod,omitempty" bson:"default_value_period,omitempty"`                                 // Specified value if missing from instance
	DefaultValueQuantity              *Quantity                     `json:"defaultValueQuantity,omitempty" bson:"default_value_quantity,omitempty"`                             // Specified value if missing from instance
	DefaultValueRange                 *Range                        `json:"defaultValueRange,omitempty" bson:"default_value_range,omitempty"`                                   // Specified value if missing from instance
	DefaultValueRatio                 *Ratio                        `json:"defaultValueRatio,omitempty" bson:"default_value_ratio,omitempty"`                                   // Specified value if missing from instance
	DefaultValueRatioRange            *RatioRange                   `json:"defaultValueRatioRange,omitempty" bson:"default_value_ratio_range,omitempty"`                        // Specified value if missing from instance
	DefaultValueReference             *Reference                    `json:"defaultValueReference,omitempty" bson:"default_value_reference,omitempty"`                           // Specified value if missing from instance
	DefaultValueSampledData           *SampledData                  `json:"defaultValueSampledData,omitempty" bson:"default_value_sampled_data,omitempty"`                      // Specified value if missing from instance
	DefaultValueSignature             *Signature                    `json:"defaultValueSignature,omitempty" bson:"default_value_signature,omitempty"`                           // Specified value if missing from instance
	DefaultValueTiming                *Timing                       `json:"defaultValueTiming,omitempty" bson:"default_value_timing,omitempty"`                                 // Specified value if missing from instance
	DefaultValueContactDetail         *ContactDetail                `json:"defaultValueContactDetail,omitempty" bson:"default_value_contact_detail,omitempty"`                  // Specified value if missing from instance
	DefaultValueDataRequirement       *DataRequirement              `json:"defaultValueDataRequirement,omitempty" bson:"default_value_data_requirement,omitempty"`              // Specified value if missing from instance
	DefaultValueExpression            *Expression                   `json:"defaultValueExpression,omitempty" bson:"default_value_expression,omitempty"`                         // Specified value if missing from instance
	DefaultValueParameterDefinition   *ParameterDefinition          `json:"defaultValueParameterDefinition,omitempty" bson:"default_value_parameter_definition,omitempty"`      // Specified value if missing from instance
	DefaultValueRelatedArtifact       *RelatedArtifact              `json:"defaultValueRelatedArtifact,omitempty" bson:"default_value_related_artifact,omitempty"`              // Specified value if missing from instance
	DefaultValueTriggerDefinition     *TriggerDefinition            `json:"defaultValueTriggerDefinition,omitempty" bson:"default_value_trigger_definition,omitempty"`          // Specified value if missing from instance
	DefaultValueUsageContext          *UsageContext                 `json:"defaultValueUsageContext,omitempty" bson:"default_value_usage_context,omitempty"`                    // Specified value if missing from instance
	DefaultValueAvailability          *Availability                 `json:"defaultValueAvailability,omitempty" bson:"default_value_availability,omitempty"`                     // Specified value if missing from instance
	DefaultValueExtendedContactDetail *ExtendedContactDetail        `json:"defaultValueExtendedContactDetail,omitempty" bson:"default_value_extended_contact_detail,omitempty"` // Specified value if missing from instance
	DefaultValueVirtualServiceDetail  *VirtualServiceDetail         `json:"defaultValueVirtualServiceDetail,omitempty" bson:"default_value_virtual_service_detail,omitempty"`   // Specified value if missing from instance
	DefaultValueDosage                *Dosage                       `json:"defaultValueDosage,omitempty" bson:"default_value_dosage,omitempty"`                                 // Specified value if missing from instance
	DefaultValueMeta                  *Meta                         `json:"defaultValueMeta,omitempty" bson:"default_value_meta,omitempty"`                                     // Specified value if missing from instance
	MeaningWhenMissing                *string                       `json:"meaningWhenMissing,omitempty" bson:"meaning_when_missing,omitempty"`                                 // Implicit meaning when this element is missing
	OrderMeaning                      *string                       `json:"orderMeaning,omitempty" bson:"order_meaning,omitempty"`                                              // What the order of the elements means
	FixedBase64Binary                 *string                       `json:"fixedBase64Binary,omitempty" bson:"fixed_base64_binary,omitempty"`                                   // Value must be exactly this
	FixedBoolean                      *bool                         `json:"fixedBoolean,omitempty" bson:"fixed_boolean,omitempty"`                                              // Value must be exactly this
	FixedCanonical                    *string                       `json:"fixedCanonical,omitempty" bson:"fixed_canonical,omitempty"`                                          // Value must be exactly this
	FixedCode                         *string                       `json:"fixedCode,omitempty" bson:"fixed_code,omitempty"`                                                    // Value must be exactly this
	FixedDate                         *string                       `json:"fixedDate,omitempty" bson:"fixed_date,omitempty"`                                                    // Value must be exactly this
	FixedDateTime                     *string                       `json:"fixedDateTime,omitempty" bson:"fixed_date_time,omitempty"`                                           // Value must be exactly this
	FixedDecimal                      *float64                      `json:"fixedDecimal,omitempty" bson:"fixed_decimal,omitempty"`                                              // Value must be exactly this
	FixedId                           *string                       `json:"fixedId,omitempty" bson:"fixed_id,omitempty"`                                                        // Value must be exactly this
	FixedInstant                      *string                       `json:"fixedInstant,omitempty" bson:"fixed_instant,omitempty"`                                              // Value must be exactly this
	FixedInteger                      *int                          `json:"fixedInteger,omitempty" bson:"fixed_integer,omitempty"`                                              // Value must be exactly this
	FixedInteger64                    *int64                        `json:"fixedInteger64,omitempty" bson:"fixed_integer64,omitempty"`                                          // Value must be exactly this
	FixedMarkdown                     *string                       `json:"fixedMarkdown,omitempty" bson:"fixed_markdown,omitempty"`                                            // Value must be exactly this
	FixedOid                          *string                       `json:"fixedOid,omitempty" bson:"fixed_oid,omitempty"`                                                      // Value must be exactly this
	FixedPositiveInt                  *int                          `json:"fixedPositiveInt,omitempty" bson:"fixed_positive_int,omitempty"`                                     // Value must be exactly this
	FixedString                       *string                       `json:"fixedString,omitempty" bson:"fixed_string,omitempty"`                                                // Value must be exactly this
	FixedTime                         *string                       `json:"fixedTime,omitempty" bson:"fixed_time,omitempty"`                                                    // Value must be exactly this
	FixedUnsignedInt                  *int                          `json:"fixedUnsignedInt,omitempty" bson:"fixed_unsigned_int,omitempty"`                                     // Value must be exactly this
	FixedUri                          *string                       `json:"fixedUri,omitempty" bson:"fixed_uri,omitempty"`                                                      // Value must be exactly this
	FixedUrl                          *string                       `json:"fixedUrl,omitempty" bson:"fixed_url,omitempty"`                                                      // Value must be exactly this
	FixedUuid                         *uuid                         `json:"fixedUuid,omitempty" bson:"fixed_uuid,omitempty"`                                                    // Value must be exactly this
	FixedAddress                      *Address                      `json:"fixedAddress,omitempty" bson:"fixed_address,omitempty"`                                              // Value must be exactly this
	FixedAge                          *Age                          `json:"fixedAge,omitempty" bson:"fixed_age,omitempty"`                                                      // Value must be exactly this
	FixedAnnotation                   *Annotation                   `json:"fixedAnnotation,omitempty" bson:"fixed_annotation,omitempty"`                                        // Value must be exactly this
	FixedAttachment                   *Attachment                   `json:"fixedAttachment,omitempty" bson:"fixed_attachment,omitempty"`                                        // Value must be exactly this
	FixedCodeableConcept              *CodeableConcept              `json:"fixedCodeableConcept,omitempty" bson:"fixed_codeable_concept,omitempty"`                             // Value must be exactly this
	FixedCodeableReference            *CodeableReference            `json:"fixedCodeableReference,omitempty" bson:"fixed_codeable_reference,omitempty"`                         // Value must be exactly this
	FixedCoding                       *Coding                       `json:"fixedCoding,omitempty" bson:"fixed_coding,omitempty"`                                                // Value must be exactly this
	FixedContactPoint                 *ContactPoint                 `json:"fixedContactPoint,omitempty" bson:"fixed_contact_point,omitempty"`                                   // Value must be exactly this
	FixedCount                        *Count                        `json:"fixedCount,omitempty" bson:"fixed_count,omitempty"`                                                  // Value must be exactly this
	FixedDistance                     *Distance                     `json:"fixedDistance,omitempty" bson:"fixed_distance,omitempty"`                                            // Value must be exactly this
	FixedDuration                     *Duration                     `json:"fixedDuration,omitempty" bson:"fixed_duration,omitempty"`                                            // Value must be exactly this
	FixedHumanName                    *HumanName                    `json:"fixedHumanName,omitempty" bson:"fixed_human_name,omitempty"`                                         // Value must be exactly this
	FixedIdentifier                   *Identifier                   `json:"fixedIdentifier,omitempty" bson:"fixed_identifier,omitempty"`                                        // Value must be exactly this
	FixedMoney                        *Money                        `json:"fixedMoney,omitempty" bson:"fixed_money,omitempty"`                                                  // Value must be exactly this
	FixedPeriod                       *Period                       `json:"fixedPeriod,omitempty" bson:"fixed_period,omitempty"`                                                // Value must be exactly this
	FixedQuantity                     *Quantity                     `json:"fixedQuantity,omitempty" bson:"fixed_quantity,omitempty"`                                            // Value must be exactly this
	FixedRange                        *Range                        `json:"fixedRange,omitempty" bson:"fixed_range,omitempty"`                                                  // Value must be exactly this
	FixedRatio                        *Ratio                        `json:"fixedRatio,omitempty" bson:"fixed_ratio,omitempty"`                                                  // Value must be exactly this
	FixedRatioRange                   *RatioRange                   `json:"fixedRatioRange,omitempty" bson:"fixed_ratio_range,omitempty"`                                       // Value must be exactly this
	FixedReference                    *Reference                    `json:"fixedReference,omitempty" bson:"fixed_reference,omitempty"`                                          // Value must be exactly this
	FixedSampledData                  *SampledData                  `json:"fixedSampledData,omitempty" bson:"fixed_sampled_data,omitempty"`                                     // Value must be exactly this
	FixedSignature                    *Signature                    `json:"fixedSignature,omitempty" bson:"fixed_signature,omitempty"`                                          // Value must be exactly this
	FixedTiming                       *Timing                       `json:"fixedTiming,omitempty" bson:"fixed_timing,omitempty"`                                                // Value must be exactly this
	FixedContactDetail                *ContactDetail                `json:"fixedContactDetail,omitempty" bson:"fixed_contact_detail,omitempty"`                                 // Value must be exactly this
	FixedDataRequirement              *DataRequirement              `json:"fixedDataRequirement,omitempty" bson:"fixed_data_requirement,omitempty"`                             // Value must be exactly this
	FixedExpression                   *Expression                   `json:"fixedExpression,omitempty" bson:"fixed_expression,omitempty"`                                        // Value must be exactly this
	FixedParameterDefinition          *ParameterDefinition          `json:"fixedParameterDefinition,omitempty" bson:"fixed_parameter_definition,omitempty"`                     // Value must be exactly this
	FixedRelatedArtifact              *RelatedArtifact              `json:"fixedRelatedArtifact,omitempty" bson:"fixed_related_artifact,omitempty"`                             // Value must be exactly this
	FixedTriggerDefinition            *TriggerDefinition            `json:"fixedTriggerDefinition,omitempty" bson:"fixed_trigger_definition,omitempty"`                         // Value must be exactly this
	FixedUsageContext                 *UsageContext                 `json:"fixedUsageContext,omitempty" bson:"fixed_usage_context,omitempty"`                                   // Value must be exactly this
	FixedAvailability                 *Availability                 `json:"fixedAvailability,omitempty" bson:"fixed_availability,omitempty"`                                    // Value must be exactly this
	FixedExtendedContactDetail        *ExtendedContactDetail        `json:"fixedExtendedContactDetail,omitempty" bson:"fixed_extended_contact_detail,omitempty"`                // Value must be exactly this
	FixedVirtualServiceDetail         *VirtualServiceDetail         `json:"fixedVirtualServiceDetail,omitempty" bson:"fixed_virtual_service_detail,omitempty"`                  // Value must be exactly this
	FixedDosage                       *Dosage                       `json:"fixedDosage,omitempty" bson:"fixed_dosage,omitempty"`                                                // Value must be exactly this
	FixedMeta                         *Meta                         `json:"fixedMeta,omitempty" bson:"fixed_meta,omitempty"`                                                    // Value must be exactly this
	PatternBase64Binary               *string                       `json:"patternBase64Binary,omitempty" bson:"pattern_base64_binary,omitempty"`                               // Value must have at least these property values
	PatternBoolean                    *bool                         `json:"patternBoolean,omitempty" bson:"pattern_boolean,omitempty"`                                          // Value must have at least these property values
	PatternCanonical                  *string                       `json:"patternCanonical,omitempty" bson:"pattern_canonical,omitempty"`                                      // Value must have at least these property values
	PatternCode                       *string                       `json:"patternCode,omitempty" bson:"pattern_code,omitempty"`                                                // Value must have at least these property values
	PatternDate                       *string                       `json:"patternDate,omitempty" bson:"pattern_date,omitempty"`                                                // Value must have at least these property values
	PatternDateTime                   *string                       `json:"patternDateTime,omitempty" bson:"pattern_date_time,omitempty"`                                       // Value must have at least these property values
	PatternDecimal                    *float64                      `json:"patternDecimal,omitempty" bson:"pattern_decimal,omitempty"`                                          // Value must have at least these property values
	PatternId                         *string                       `json:"patternId,omitempty" bson:"pattern_id,omitempty"`                                                    // Value must have at least these property values
	PatternInstant                    *string                       `json:"patternInstant,omitempty" bson:"pattern_instant,omitempty"`                                          // Value must have at least these property values
	PatternInteger                    *int                          `json:"patternInteger,omitempty" bson:"pattern_integer,omitempty"`                                          // Value must have at least these property values
	PatternInteger64                  *int64                        `json:"patternInteger64,omitempty" bson:"pattern_integer64,omitempty"`                                      // Value must have at least these property values
	PatternMarkdown                   *string                       `json:"patternMarkdown,omitempty" bson:"pattern_markdown,omitempty"`                                        // Value must have at least these property values
	PatternOid                        *string                       `json:"patternOid,omitempty" bson:"pattern_oid,omitempty"`                                                  // Value must have at least these property values
	PatternPositiveInt                *int                          `json:"patternPositiveInt,omitempty" bson:"pattern_positive_int,omitempty"`                                 // Value must have at least these property values
	PatternString                     *string                       `json:"patternString,omitempty" bson:"pattern_string,omitempty"`                                            // Value must have at least these property values
	PatternTime                       *string                       `json:"patternTime,omitempty" bson:"pattern_time,omitempty"`                                                // Value must have at least these property values
	PatternUnsignedInt                *int                          `json:"patternUnsignedInt,omitempty" bson:"pattern_unsigned_int,omitempty"`                                 // Value must have at least these property values
	PatternUri                        *string                       `json:"patternUri,omitempty" bson:"pattern_uri,omitempty"`                                                  // Value must have at least these property values
	PatternUrl                        *string                       `json:"patternUrl,omitempty" bson:"pattern_url,omitempty"`                                                  // Value must have at least these property values
	PatternUuid                       *uuid                         `json:"patternUuid,omitempty" bson:"pattern_uuid,omitempty"`                                                // Value must have at least these property values
	PatternAddress                    *Address                      `json:"patternAddress,omitempty" bson:"pattern_address,omitempty"`                                          // Value must have at least these property values
	PatternAge                        *Age                          `json:"patternAge,omitempty" bson:"pattern_age,omitempty"`                                                  // Value must have at least these property values
	PatternAnnotation                 *Annotation                   `json:"patternAnnotation,omitempty" bson:"pattern_annotation,omitempty"`                                    // Value must have at least these property values
	PatternAttachment                 *Attachment                   `json:"patternAttachment,omitempty" bson:"pattern_attachment,omitempty"`                                    // Value must have at least these property values
	PatternCodeableConcept            *CodeableConcept              `json:"patternCodeableConcept,omitempty" bson:"pattern_codeable_concept,omitempty"`                         // Value must have at least these property values
	PatternCodeableReference          *CodeableReference            `json:"patternCodeableReference,omitempty" bson:"pattern_codeable_reference,omitempty"`                     // Value must have at least these property values
	PatternCoding                     *Coding                       `json:"patternCoding,omitempty" bson:"pattern_coding,omitempty"`                                            // Value must have at least these property values
	PatternContactPoint               *ContactPoint                 `json:"patternContactPoint,omitempty" bson:"pattern_contact_point,omitempty"`                               // Value must have at least these property values
	PatternCount                      *Count                        `json:"patternCount,omitempty" bson:"pattern_count,omitempty"`                                              // Value must have at least these property values
	PatternDistance                   *Distance                     `json:"patternDistance,omitempty" bson:"pattern_distance,omitempty"`                                        // Value must have at least these property values
	PatternDuration                   *Duration                     `json:"patternDuration,omitempty" bson:"pattern_duration,omitempty"`                                        // Value must have at least these property values
	PatternHumanName                  *HumanName                    `json:"patternHumanName,omitempty" bson:"pattern_human_name,omitempty"`                                     // Value must have at least these property values
	PatternIdentifier                 *Identifier                   `json:"patternIdentifier,omitempty" bson:"pattern_identifier,omitempty"`                                    // Value must have at least these property values
	PatternMoney                      *Money                        `json:"patternMoney,omitempty" bson:"pattern_money,omitempty"`                                              // Value must have at least these property values
	PatternPeriod                     *Period                       `json:"patternPeriod,omitempty" bson:"pattern_period,omitempty"`                                            // Value must have at least these property values
	PatternQuantity                   *Quantity                     `json:"patternQuantity,omitempty" bson:"pattern_quantity,omitempty"`                                        // Value must have at least these property values
	PatternRange                      *Range                        `json:"patternRange,omitempty" bson:"pattern_range,omitempty"`                                              // Value must have at least these property values
	PatternRatio                      *Ratio                        `json:"patternRatio,omitempty" bson:"pattern_ratio,omitempty"`                                              // Value must have at least these property values
	PatternRatioRange                 *RatioRange                   `json:"patternRatioRange,omitempty" bson:"pattern_ratio_range,omitempty"`                                   // Value must have at least these property values
	PatternReference                  *Reference                    `json:"patternReference,omitempty" bson:"pattern_reference,omitempty"`                                      // Value must have at least these property values
	PatternSampledData                *SampledData                  `json:"patternSampledData,omitempty" bson:"pattern_sampled_data,omitempty"`                                 // Value must have at least these property values
	PatternSignature                  *Signature                    `json:"patternSignature,omitempty" bson:"pattern_signature,omitempty"`                                      // Value must have at least these property values
	PatternTiming                     *Timing                       `json:"patternTiming,omitempty" bson:"pattern_timing,omitempty"`                                            // Value must have at least these property values
	PatternContactDetail              *ContactDetail                `json:"patternContactDetail,omitempty" bson:"pattern_contact_detail,omitempty"`                             // Value must have at least these property values
	PatternDataRequirement            *DataRequirement              `json:"patternDataRequirement,omitempty" bson:"pattern_data_requirement,omitempty"`                         // Value must have at least these property values
	PatternExpression                 *Expression                   `json:"patternExpression,omitempty" bson:"pattern_expression,omitempty"`                                    // Value must have at least these property values
	PatternParameterDefinition        *ParameterDefinition          `json:"patternParameterDefinition,omitempty" bson:"pattern_parameter_definition,omitempty"`                 // Value must have at least these property values
	PatternRelatedArtifact            *RelatedArtifact              `json:"patternRelatedArtifact,omitempty" bson:"pattern_related_artifact,omitempty"`                         // Value must have at least these property values
	PatternTriggerDefinition          *TriggerDefinition            `json:"patternTriggerDefinition,omitempty" bson:"pattern_trigger_definition,omitempty"`                     // Value must have at least these property values
	PatternUsageContext               *UsageContext                 `json:"patternUsageContext,omitempty" bson:"pattern_usage_context,omitempty"`                               // Value must have at least these property values
	PatternAvailability               *Availability                 `json:"patternAvailability,omitempty" bson:"pattern_availability,omitempty"`                                // Value must have at least these property values
	PatternExtendedContactDetail      *ExtendedContactDetail        `json:"patternExtendedContactDetail,omitempty" bson:"pattern_extended_contact_detail,omitempty"`            // Value must have at least these property values
	PatternVirtualServiceDetail       *VirtualServiceDetail         `json:"patternVirtualServiceDetail,omitempty" bson:"pattern_virtual_service_detail,omitempty"`              // Value must have at least these property values
	PatternDosage                     *Dosage                       `json:"patternDosage,omitempty" bson:"pattern_dosage,omitempty"`                                            // Value must have at least these property values
	PatternMeta                       *Meta                         `json:"patternMeta,omitempty" bson:"pattern_meta,omitempty"`                                                // Value must have at least these property values
	Example                           []ElementDefinitionExample    `json:"example,omitempty" bson:"example,omitempty"`                                                         // Example value (as defined for type)
	MinValueDate                      *string                       `json:"minValueDate,omitempty" bson:"min_value_date,omitempty"`                                             // Minimum Allowed Value (for some types)
	MinValueDateTime                  *string                       `json:"minValueDateTime,omitempty" bson:"min_value_date_time,omitempty"`                                    // Minimum Allowed Value (for some types)
	MinValueInstant                   *string                       `json:"minValueInstant,omitempty" bson:"min_value_instant,omitempty"`                                       // Minimum Allowed Value (for some types)
	MinValueTime                      *string                       `json:"minValueTime,omitempty" bson:"min_value_time,omitempty"`                                             // Minimum Allowed Value (for some types)
	MinValueDecimal                   *float64                      `json:"minValueDecimal,omitempty" bson:"min_value_decimal,omitempty"`                                       // Minimum Allowed Value (for some types)
	MinValueInteger                   *int                          `json:"minValueInteger,omitempty" bson:"min_value_integer,omitempty"`                                       // Minimum Allowed Value (for some types)
	MinValueInteger64                 *int64                        `json:"minValueInteger64,omitempty" bson:"min_value_integer64,omitempty"`                                   // Minimum Allowed Value (for some types)
	MinValuePositiveInt               *int                          `json:"minValuePositiveInt,omitempty" bson:"min_value_positive_int,omitempty"`                              // Minimum Allowed Value (for some types)
	MinValueUnsignedInt               *int                          `json:"minValueUnsignedInt,omitempty" bson:"min_value_unsigned_int,omitempty"`                              // Minimum Allowed Value (for some types)
	MinValueQuantity                  *Quantity                     `json:"minValueQuantity,omitempty" bson:"min_value_quantity,omitempty"`                                     // Minimum Allowed Value (for some types)
	MaxValueDate                      *string                       `json:"maxValueDate,omitempty" bson:"max_value_date,omitempty"`                                             // Maximum Allowed Value (for some types)
	MaxValueDateTime                  *string                       `json:"maxValueDateTime,omitempty" bson:"max_value_date_time,omitempty"`                                    // Maximum Allowed Value (for some types)
	MaxValueInstant                   *string                       `json:"maxValueInstant,omitempty" bson:"max_value_instant,omitempty"`                                       // Maximum Allowed Value (for some types)
	MaxValueTime                      *string                       `json:"maxValueTime,omitempty" bson:"max_value_time,omitempty"`                                             // Maximum Allowed Value (for some types)
	MaxValueDecimal                   *float64                      `json:"maxValueDecimal,omitempty" bson:"max_value_decimal,omitempty"`                                       // Maximum Allowed Value (for some types)
	MaxValueInteger                   *int                          `json:"maxValueInteger,omitempty" bson:"max_value_integer,omitempty"`                                       // Maximum Allowed Value (for some types)
	MaxValueInteger64                 *int64                        `json:"maxValueInteger64,omitempty" bson:"max_value_integer64,omitempty"`                                   // Maximum Allowed Value (for some types)
	MaxValuePositiveInt               *int                          `json:"maxValuePositiveInt,omitempty" bson:"max_value_positive_int,omitempty"`                              // Maximum Allowed Value (for some types)
	MaxValueUnsignedInt               *int                          `json:"maxValueUnsignedInt,omitempty" bson:"max_value_unsigned_int,omitempty"`                              // Maximum Allowed Value (for some types)
	MaxValueQuantity                  *Quantity                     `json:"maxValueQuantity,omitempty" bson:"max_value_quantity,omitempty"`                                     // Maximum Allowed Value (for some types)
	MaxLength                         *int                          `json:"maxLength,omitempty" bson:"max_length,omitempty"`                                                    // Max length for string type data
	Condition                         []string                      `json:"condition,omitempty" bson:"condition,omitempty"`                                                     // Reference to invariant about presence
	Constraint                        []ElementDefinitionConstraint `json:"constraint,omitempty" bson:"constraint,omitempty"`                                                   // Condition that must evaluate to true
	MustHaveValue                     bool                          `json:"mustHaveValue,omitempty" bson:"must_have_value,omitempty"`                                           // For primitives, that a value must be present - not replaced by an extension
	ValueAlternatives                 []string                      `json:"valueAlternatives,omitempty" bson:"value_alternatives,omitempty"`                                    // Extensions that are allowed to replace a primitive value
	MustSupport                       bool                          `json:"mustSupport,omitempty" bson:"must_support,omitempty"`                                                // If the element must be supported (discouraged - see obligations)
	IsModifier                        bool                          `json:"isModifier,omitempty" bson:"is_modifier,omitempty"`                                                  // If this modifies the meaning of other elements
	IsModifierReason                  *string                       `json:"isModifierReason,omitempty" bson:"is_modifier_reason,omitempty"`                                     // Reason that this element is marked as a modifier
	IsSummary                         bool                          `json:"isSummary,omitempty" bson:"is_summary,omitempty"`                                                    // Include when _summary = true?
	Binding                           *ElementDefinitionBinding     `json:"binding,omitempty" bson:"binding,omitempty"`                                                         // ValueSet details if this is coded
	Mapping                           []ElementDefinitionMapping    `json:"mapping,omitempty" bson:"mapping,omitempty"`                                                         // Map element to another set of definitions
}

func (r *ElementDefinition) Validate() error {
	var emptyString string
	if r.Path == emptyString {
		return fmt.Errorf("field 'Path' is required")
	}
	for i, item := range r.Code {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Code[%d]: %w", i, err)
		}
	}
	if r.Slicing != nil {
		if err := r.Slicing.Validate(); err != nil {
			return fmt.Errorf("Slicing: %w", err)
		}
	}
	if r.Base != nil {
		if err := r.Base.Validate(); err != nil {
			return fmt.Errorf("Base: %w", err)
		}
	}
	for i, item := range r.Type {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Type[%d]: %w", i, err)
		}
	}
	if r.DefaultValueUuid != nil {
		if err := r.DefaultValueUuid.Validate(); err != nil {
			return fmt.Errorf("DefaultValueUuid: %w", err)
		}
	}
	if r.DefaultValueAddress != nil {
		if err := r.DefaultValueAddress.Validate(); err != nil {
			return fmt.Errorf("DefaultValueAddress: %w", err)
		}
	}
	if r.DefaultValueAge != nil {
		if err := r.DefaultValueAge.Validate(); err != nil {
			return fmt.Errorf("DefaultValueAge: %w", err)
		}
	}
	if r.DefaultValueAnnotation != nil {
		if err := r.DefaultValueAnnotation.Validate(); err != nil {
			return fmt.Errorf("DefaultValueAnnotation: %w", err)
		}
	}
	if r.DefaultValueAttachment != nil {
		if err := r.DefaultValueAttachment.Validate(); err != nil {
			return fmt.Errorf("DefaultValueAttachment: %w", err)
		}
	}
	if r.DefaultValueCodeableConcept != nil {
		if err := r.DefaultValueCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("DefaultValueCodeableConcept: %w", err)
		}
	}
	if r.DefaultValueCodeableReference != nil {
		if err := r.DefaultValueCodeableReference.Validate(); err != nil {
			return fmt.Errorf("DefaultValueCodeableReference: %w", err)
		}
	}
	if r.DefaultValueCoding != nil {
		if err := r.DefaultValueCoding.Validate(); err != nil {
			return fmt.Errorf("DefaultValueCoding: %w", err)
		}
	}
	if r.DefaultValueContactPoint != nil {
		if err := r.DefaultValueContactPoint.Validate(); err != nil {
			return fmt.Errorf("DefaultValueContactPoint: %w", err)
		}
	}
	if r.DefaultValueCount != nil {
		if err := r.DefaultValueCount.Validate(); err != nil {
			return fmt.Errorf("DefaultValueCount: %w", err)
		}
	}
	if r.DefaultValueDistance != nil {
		if err := r.DefaultValueDistance.Validate(); err != nil {
			return fmt.Errorf("DefaultValueDistance: %w", err)
		}
	}
	if r.DefaultValueDuration != nil {
		if err := r.DefaultValueDuration.Validate(); err != nil {
			return fmt.Errorf("DefaultValueDuration: %w", err)
		}
	}
	if r.DefaultValueHumanName != nil {
		if err := r.DefaultValueHumanName.Validate(); err != nil {
			return fmt.Errorf("DefaultValueHumanName: %w", err)
		}
	}
	if r.DefaultValueIdentifier != nil {
		if err := r.DefaultValueIdentifier.Validate(); err != nil {
			return fmt.Errorf("DefaultValueIdentifier: %w", err)
		}
	}
	if r.DefaultValueMoney != nil {
		if err := r.DefaultValueMoney.Validate(); err != nil {
			return fmt.Errorf("DefaultValueMoney: %w", err)
		}
	}
	if r.DefaultValuePeriod != nil {
		if err := r.DefaultValuePeriod.Validate(); err != nil {
			return fmt.Errorf("DefaultValuePeriod: %w", err)
		}
	}
	if r.DefaultValueQuantity != nil {
		if err := r.DefaultValueQuantity.Validate(); err != nil {
			return fmt.Errorf("DefaultValueQuantity: %w", err)
		}
	}
	if r.DefaultValueRange != nil {
		if err := r.DefaultValueRange.Validate(); err != nil {
			return fmt.Errorf("DefaultValueRange: %w", err)
		}
	}
	if r.DefaultValueRatio != nil {
		if err := r.DefaultValueRatio.Validate(); err != nil {
			return fmt.Errorf("DefaultValueRatio: %w", err)
		}
	}
	if r.DefaultValueRatioRange != nil {
		if err := r.DefaultValueRatioRange.Validate(); err != nil {
			return fmt.Errorf("DefaultValueRatioRange: %w", err)
		}
	}
	if r.DefaultValueReference != nil {
		if err := r.DefaultValueReference.Validate(); err != nil {
			return fmt.Errorf("DefaultValueReference: %w", err)
		}
	}
	if r.DefaultValueSampledData != nil {
		if err := r.DefaultValueSampledData.Validate(); err != nil {
			return fmt.Errorf("DefaultValueSampledData: %w", err)
		}
	}
	if r.DefaultValueSignature != nil {
		if err := r.DefaultValueSignature.Validate(); err != nil {
			return fmt.Errorf("DefaultValueSignature: %w", err)
		}
	}
	if r.DefaultValueTiming != nil {
		if err := r.DefaultValueTiming.Validate(); err != nil {
			return fmt.Errorf("DefaultValueTiming: %w", err)
		}
	}
	if r.DefaultValueContactDetail != nil {
		if err := r.DefaultValueContactDetail.Validate(); err != nil {
			return fmt.Errorf("DefaultValueContactDetail: %w", err)
		}
	}
	if r.DefaultValueDataRequirement != nil {
		if err := r.DefaultValueDataRequirement.Validate(); err != nil {
			return fmt.Errorf("DefaultValueDataRequirement: %w", err)
		}
	}
	if r.DefaultValueExpression != nil {
		if err := r.DefaultValueExpression.Validate(); err != nil {
			return fmt.Errorf("DefaultValueExpression: %w", err)
		}
	}
	if r.DefaultValueParameterDefinition != nil {
		if err := r.DefaultValueParameterDefinition.Validate(); err != nil {
			return fmt.Errorf("DefaultValueParameterDefinition: %w", err)
		}
	}
	if r.DefaultValueRelatedArtifact != nil {
		if err := r.DefaultValueRelatedArtifact.Validate(); err != nil {
			return fmt.Errorf("DefaultValueRelatedArtifact: %w", err)
		}
	}
	if r.DefaultValueTriggerDefinition != nil {
		if err := r.DefaultValueTriggerDefinition.Validate(); err != nil {
			return fmt.Errorf("DefaultValueTriggerDefinition: %w", err)
		}
	}
	if r.DefaultValueUsageContext != nil {
		if err := r.DefaultValueUsageContext.Validate(); err != nil {
			return fmt.Errorf("DefaultValueUsageContext: %w", err)
		}
	}
	if r.DefaultValueAvailability != nil {
		if err := r.DefaultValueAvailability.Validate(); err != nil {
			return fmt.Errorf("DefaultValueAvailability: %w", err)
		}
	}
	if r.DefaultValueExtendedContactDetail != nil {
		if err := r.DefaultValueExtendedContactDetail.Validate(); err != nil {
			return fmt.Errorf("DefaultValueExtendedContactDetail: %w", err)
		}
	}
	if r.DefaultValueVirtualServiceDetail != nil {
		if err := r.DefaultValueVirtualServiceDetail.Validate(); err != nil {
			return fmt.Errorf("DefaultValueVirtualServiceDetail: %w", err)
		}
	}
	if r.DefaultValueDosage != nil {
		if err := r.DefaultValueDosage.Validate(); err != nil {
			return fmt.Errorf("DefaultValueDosage: %w", err)
		}
	}
	if r.DefaultValueMeta != nil {
		if err := r.DefaultValueMeta.Validate(); err != nil {
			return fmt.Errorf("DefaultValueMeta: %w", err)
		}
	}
	if r.FixedUuid != nil {
		if err := r.FixedUuid.Validate(); err != nil {
			return fmt.Errorf("FixedUuid: %w", err)
		}
	}
	if r.FixedAddress != nil {
		if err := r.FixedAddress.Validate(); err != nil {
			return fmt.Errorf("FixedAddress: %w", err)
		}
	}
	if r.FixedAge != nil {
		if err := r.FixedAge.Validate(); err != nil {
			return fmt.Errorf("FixedAge: %w", err)
		}
	}
	if r.FixedAnnotation != nil {
		if err := r.FixedAnnotation.Validate(); err != nil {
			return fmt.Errorf("FixedAnnotation: %w", err)
		}
	}
	if r.FixedAttachment != nil {
		if err := r.FixedAttachment.Validate(); err != nil {
			return fmt.Errorf("FixedAttachment: %w", err)
		}
	}
	if r.FixedCodeableConcept != nil {
		if err := r.FixedCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("FixedCodeableConcept: %w", err)
		}
	}
	if r.FixedCodeableReference != nil {
		if err := r.FixedCodeableReference.Validate(); err != nil {
			return fmt.Errorf("FixedCodeableReference: %w", err)
		}
	}
	if r.FixedCoding != nil {
		if err := r.FixedCoding.Validate(); err != nil {
			return fmt.Errorf("FixedCoding: %w", err)
		}
	}
	if r.FixedContactPoint != nil {
		if err := r.FixedContactPoint.Validate(); err != nil {
			return fmt.Errorf("FixedContactPoint: %w", err)
		}
	}
	if r.FixedCount != nil {
		if err := r.FixedCount.Validate(); err != nil {
			return fmt.Errorf("FixedCount: %w", err)
		}
	}
	if r.FixedDistance != nil {
		if err := r.FixedDistance.Validate(); err != nil {
			return fmt.Errorf("FixedDistance: %w", err)
		}
	}
	if r.FixedDuration != nil {
		if err := r.FixedDuration.Validate(); err != nil {
			return fmt.Errorf("FixedDuration: %w", err)
		}
	}
	if r.FixedHumanName != nil {
		if err := r.FixedHumanName.Validate(); err != nil {
			return fmt.Errorf("FixedHumanName: %w", err)
		}
	}
	if r.FixedIdentifier != nil {
		if err := r.FixedIdentifier.Validate(); err != nil {
			return fmt.Errorf("FixedIdentifier: %w", err)
		}
	}
	if r.FixedMoney != nil {
		if err := r.FixedMoney.Validate(); err != nil {
			return fmt.Errorf("FixedMoney: %w", err)
		}
	}
	if r.FixedPeriod != nil {
		if err := r.FixedPeriod.Validate(); err != nil {
			return fmt.Errorf("FixedPeriod: %w", err)
		}
	}
	if r.FixedQuantity != nil {
		if err := r.FixedQuantity.Validate(); err != nil {
			return fmt.Errorf("FixedQuantity: %w", err)
		}
	}
	if r.FixedRange != nil {
		if err := r.FixedRange.Validate(); err != nil {
			return fmt.Errorf("FixedRange: %w", err)
		}
	}
	if r.FixedRatio != nil {
		if err := r.FixedRatio.Validate(); err != nil {
			return fmt.Errorf("FixedRatio: %w", err)
		}
	}
	if r.FixedRatioRange != nil {
		if err := r.FixedRatioRange.Validate(); err != nil {
			return fmt.Errorf("FixedRatioRange: %w", err)
		}
	}
	if r.FixedReference != nil {
		if err := r.FixedReference.Validate(); err != nil {
			return fmt.Errorf("FixedReference: %w", err)
		}
	}
	if r.FixedSampledData != nil {
		if err := r.FixedSampledData.Validate(); err != nil {
			return fmt.Errorf("FixedSampledData: %w", err)
		}
	}
	if r.FixedSignature != nil {
		if err := r.FixedSignature.Validate(); err != nil {
			return fmt.Errorf("FixedSignature: %w", err)
		}
	}
	if r.FixedTiming != nil {
		if err := r.FixedTiming.Validate(); err != nil {
			return fmt.Errorf("FixedTiming: %w", err)
		}
	}
	if r.FixedContactDetail != nil {
		if err := r.FixedContactDetail.Validate(); err != nil {
			return fmt.Errorf("FixedContactDetail: %w", err)
		}
	}
	if r.FixedDataRequirement != nil {
		if err := r.FixedDataRequirement.Validate(); err != nil {
			return fmt.Errorf("FixedDataRequirement: %w", err)
		}
	}
	if r.FixedExpression != nil {
		if err := r.FixedExpression.Validate(); err != nil {
			return fmt.Errorf("FixedExpression: %w", err)
		}
	}
	if r.FixedParameterDefinition != nil {
		if err := r.FixedParameterDefinition.Validate(); err != nil {
			return fmt.Errorf("FixedParameterDefinition: %w", err)
		}
	}
	if r.FixedRelatedArtifact != nil {
		if err := r.FixedRelatedArtifact.Validate(); err != nil {
			return fmt.Errorf("FixedRelatedArtifact: %w", err)
		}
	}
	if r.FixedTriggerDefinition != nil {
		if err := r.FixedTriggerDefinition.Validate(); err != nil {
			return fmt.Errorf("FixedTriggerDefinition: %w", err)
		}
	}
	if r.FixedUsageContext != nil {
		if err := r.FixedUsageContext.Validate(); err != nil {
			return fmt.Errorf("FixedUsageContext: %w", err)
		}
	}
	if r.FixedAvailability != nil {
		if err := r.FixedAvailability.Validate(); err != nil {
			return fmt.Errorf("FixedAvailability: %w", err)
		}
	}
	if r.FixedExtendedContactDetail != nil {
		if err := r.FixedExtendedContactDetail.Validate(); err != nil {
			return fmt.Errorf("FixedExtendedContactDetail: %w", err)
		}
	}
	if r.FixedVirtualServiceDetail != nil {
		if err := r.FixedVirtualServiceDetail.Validate(); err != nil {
			return fmt.Errorf("FixedVirtualServiceDetail: %w", err)
		}
	}
	if r.FixedDosage != nil {
		if err := r.FixedDosage.Validate(); err != nil {
			return fmt.Errorf("FixedDosage: %w", err)
		}
	}
	if r.FixedMeta != nil {
		if err := r.FixedMeta.Validate(); err != nil {
			return fmt.Errorf("FixedMeta: %w", err)
		}
	}
	if r.PatternUuid != nil {
		if err := r.PatternUuid.Validate(); err != nil {
			return fmt.Errorf("PatternUuid: %w", err)
		}
	}
	if r.PatternAddress != nil {
		if err := r.PatternAddress.Validate(); err != nil {
			return fmt.Errorf("PatternAddress: %w", err)
		}
	}
	if r.PatternAge != nil {
		if err := r.PatternAge.Validate(); err != nil {
			return fmt.Errorf("PatternAge: %w", err)
		}
	}
	if r.PatternAnnotation != nil {
		if err := r.PatternAnnotation.Validate(); err != nil {
			return fmt.Errorf("PatternAnnotation: %w", err)
		}
	}
	if r.PatternAttachment != nil {
		if err := r.PatternAttachment.Validate(); err != nil {
			return fmt.Errorf("PatternAttachment: %w", err)
		}
	}
	if r.PatternCodeableConcept != nil {
		if err := r.PatternCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("PatternCodeableConcept: %w", err)
		}
	}
	if r.PatternCodeableReference != nil {
		if err := r.PatternCodeableReference.Validate(); err != nil {
			return fmt.Errorf("PatternCodeableReference: %w", err)
		}
	}
	if r.PatternCoding != nil {
		if err := r.PatternCoding.Validate(); err != nil {
			return fmt.Errorf("PatternCoding: %w", err)
		}
	}
	if r.PatternContactPoint != nil {
		if err := r.PatternContactPoint.Validate(); err != nil {
			return fmt.Errorf("PatternContactPoint: %w", err)
		}
	}
	if r.PatternCount != nil {
		if err := r.PatternCount.Validate(); err != nil {
			return fmt.Errorf("PatternCount: %w", err)
		}
	}
	if r.PatternDistance != nil {
		if err := r.PatternDistance.Validate(); err != nil {
			return fmt.Errorf("PatternDistance: %w", err)
		}
	}
	if r.PatternDuration != nil {
		if err := r.PatternDuration.Validate(); err != nil {
			return fmt.Errorf("PatternDuration: %w", err)
		}
	}
	if r.PatternHumanName != nil {
		if err := r.PatternHumanName.Validate(); err != nil {
			return fmt.Errorf("PatternHumanName: %w", err)
		}
	}
	if r.PatternIdentifier != nil {
		if err := r.PatternIdentifier.Validate(); err != nil {
			return fmt.Errorf("PatternIdentifier: %w", err)
		}
	}
	if r.PatternMoney != nil {
		if err := r.PatternMoney.Validate(); err != nil {
			return fmt.Errorf("PatternMoney: %w", err)
		}
	}
	if r.PatternPeriod != nil {
		if err := r.PatternPeriod.Validate(); err != nil {
			return fmt.Errorf("PatternPeriod: %w", err)
		}
	}
	if r.PatternQuantity != nil {
		if err := r.PatternQuantity.Validate(); err != nil {
			return fmt.Errorf("PatternQuantity: %w", err)
		}
	}
	if r.PatternRange != nil {
		if err := r.PatternRange.Validate(); err != nil {
			return fmt.Errorf("PatternRange: %w", err)
		}
	}
	if r.PatternRatio != nil {
		if err := r.PatternRatio.Validate(); err != nil {
			return fmt.Errorf("PatternRatio: %w", err)
		}
	}
	if r.PatternRatioRange != nil {
		if err := r.PatternRatioRange.Validate(); err != nil {
			return fmt.Errorf("PatternRatioRange: %w", err)
		}
	}
	if r.PatternReference != nil {
		if err := r.PatternReference.Validate(); err != nil {
			return fmt.Errorf("PatternReference: %w", err)
		}
	}
	if r.PatternSampledData != nil {
		if err := r.PatternSampledData.Validate(); err != nil {
			return fmt.Errorf("PatternSampledData: %w", err)
		}
	}
	if r.PatternSignature != nil {
		if err := r.PatternSignature.Validate(); err != nil {
			return fmt.Errorf("PatternSignature: %w", err)
		}
	}
	if r.PatternTiming != nil {
		if err := r.PatternTiming.Validate(); err != nil {
			return fmt.Errorf("PatternTiming: %w", err)
		}
	}
	if r.PatternContactDetail != nil {
		if err := r.PatternContactDetail.Validate(); err != nil {
			return fmt.Errorf("PatternContactDetail: %w", err)
		}
	}
	if r.PatternDataRequirement != nil {
		if err := r.PatternDataRequirement.Validate(); err != nil {
			return fmt.Errorf("PatternDataRequirement: %w", err)
		}
	}
	if r.PatternExpression != nil {
		if err := r.PatternExpression.Validate(); err != nil {
			return fmt.Errorf("PatternExpression: %w", err)
		}
	}
	if r.PatternParameterDefinition != nil {
		if err := r.PatternParameterDefinition.Validate(); err != nil {
			return fmt.Errorf("PatternParameterDefinition: %w", err)
		}
	}
	if r.PatternRelatedArtifact != nil {
		if err := r.PatternRelatedArtifact.Validate(); err != nil {
			return fmt.Errorf("PatternRelatedArtifact: %w", err)
		}
	}
	if r.PatternTriggerDefinition != nil {
		if err := r.PatternTriggerDefinition.Validate(); err != nil {
			return fmt.Errorf("PatternTriggerDefinition: %w", err)
		}
	}
	if r.PatternUsageContext != nil {
		if err := r.PatternUsageContext.Validate(); err != nil {
			return fmt.Errorf("PatternUsageContext: %w", err)
		}
	}
	if r.PatternAvailability != nil {
		if err := r.PatternAvailability.Validate(); err != nil {
			return fmt.Errorf("PatternAvailability: %w", err)
		}
	}
	if r.PatternExtendedContactDetail != nil {
		if err := r.PatternExtendedContactDetail.Validate(); err != nil {
			return fmt.Errorf("PatternExtendedContactDetail: %w", err)
		}
	}
	if r.PatternVirtualServiceDetail != nil {
		if err := r.PatternVirtualServiceDetail.Validate(); err != nil {
			return fmt.Errorf("PatternVirtualServiceDetail: %w", err)
		}
	}
	if r.PatternDosage != nil {
		if err := r.PatternDosage.Validate(); err != nil {
			return fmt.Errorf("PatternDosage: %w", err)
		}
	}
	if r.PatternMeta != nil {
		if err := r.PatternMeta.Validate(); err != nil {
			return fmt.Errorf("PatternMeta: %w", err)
		}
	}
	for i, item := range r.Example {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Example[%d]: %w", i, err)
		}
	}
	if r.MinValueQuantity != nil {
		if err := r.MinValueQuantity.Validate(); err != nil {
			return fmt.Errorf("MinValueQuantity: %w", err)
		}
	}
	if r.MaxValueQuantity != nil {
		if err := r.MaxValueQuantity.Validate(); err != nil {
			return fmt.Errorf("MaxValueQuantity: %w", err)
		}
	}
	for i, item := range r.Constraint {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Constraint[%d]: %w", i, err)
		}
	}
	if r.Binding != nil {
		if err := r.Binding.Validate(); err != nil {
			return fmt.Errorf("Binding: %w", err)
		}
	}
	for i, item := range r.Mapping {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Mapping[%d]: %w", i, err)
		}
	}
	return nil
}

type ElementDefinitionSlicing struct {
	Id            *string                                 `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	Discriminator []ElementDefinitionSlicingDiscriminator `json:"discriminator,omitempty" bson:"discriminator,omitempty"` // Element values that are used to distinguish the slices
	Description   *string                                 `json:"description,omitempty" bson:"description,omitempty"`     // Text description of how slicing works (or not)
	Ordered       bool                                    `json:"ordered,omitempty" bson:"ordered,omitempty"`             // If elements must be in same order as slices
	Rules         string                                  `json:"rules" bson:"rules"`                                     // closed | open | openAtEnd
}

func (r *ElementDefinitionSlicing) Validate() error {
	for i, item := range r.Discriminator {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Discriminator[%d]: %w", i, err)
		}
	}
	var emptyString string
	if r.Rules == emptyString {
		return fmt.Errorf("field 'Rules' is required")
	}
	return nil
}

type ElementDefinitionSlicingDiscriminator struct {
	Id   *string `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Type string  `json:"type" bson:"type"`                 // value | exists | type | profile | position
	Path string  `json:"path" bson:"path"`                 // Path to element value
}

func (r *ElementDefinitionSlicingDiscriminator) Validate() error {
	var emptyString string
	if r.Type == emptyString {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Path == emptyString {
		return fmt.Errorf("field 'Path' is required")
	}
	return nil
}

type ElementDefinitionBase struct {
	Id   *string `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Path string  `json:"path" bson:"path"`                 // Path that identifies the base element
	Min  int     `json:"min" bson:"min"`                   // Min cardinality of the base element
	Max  string  `json:"max" bson:"max"`                   // Max cardinality of the base element
}

func (r *ElementDefinitionBase) Validate() error {
	var emptyString string
	if r.Path == emptyString {
		return fmt.Errorf("field 'Path' is required")
	}
	if r.Min == 0 {
		return fmt.Errorf("field 'Min' is required")
	}
	if r.Max == emptyString {
		return fmt.Errorf("field 'Max' is required")
	}
	return nil
}

type ElementDefinitionType struct {
	Id            *string  `json:"id,omitempty" bson:"id,omitempty"`                        // Unique id for inter-element referencing
	Code          string   `json:"code" bson:"code"`                                        // Data type or Resource (reference to definition)
	Profile       []string `json:"profile,omitempty" bson:"profile,omitempty"`              // Profiles (StructureDefinition or IG) - one must apply
	TargetProfile []string `json:"targetProfile,omitempty" bson:"target_profile,omitempty"` // Profile (StructureDefinition or IG) on the Reference/canonical target - one must apply
	Aggregation   []string `json:"aggregation,omitempty" bson:"aggregation,omitempty"`      // contained | referenced | bundled - how aggregated
	Versioning    *string  `json:"versioning,omitempty" bson:"versioning,omitempty"`        // either | independent | specific
}

func (r *ElementDefinitionType) Validate() error {
	var emptyString string
	if r.Code == emptyString {
		return fmt.Errorf("field 'Code' is required")
	}
	return nil
}

type ElementDefinitionExample struct {
	Id                         *string                `json:"id,omitempty" bson:"id,omitempty"`                                // Unique id for inter-element referencing
	Label                      string                 `json:"label" bson:"label"`                                              // Describes the purpose of this example
	ValueBase64Binary          *string                `json:"valueBase64Binary" bson:"value_base64_binary"`                    // Value of Example (one of allowed types)
	ValueBoolean               *bool                  `json:"valueBoolean" bson:"value_boolean"`                               // Value of Example (one of allowed types)
	ValueCanonical             *string                `json:"valueCanonical" bson:"value_canonical"`                           // Value of Example (one of allowed types)
	ValueCode                  *string                `json:"valueCode" bson:"value_code"`                                     // Value of Example (one of allowed types)
	ValueDate                  *string                `json:"valueDate" bson:"value_date"`                                     // Value of Example (one of allowed types)
	ValueDateTime              *string                `json:"valueDateTime" bson:"value_date_time"`                            // Value of Example (one of allowed types)
	ValueDecimal               *float64               `json:"valueDecimal" bson:"value_decimal"`                               // Value of Example (one of allowed types)
	ValueId                    *string                `json:"valueId" bson:"value_id"`                                         // Value of Example (one of allowed types)
	ValueInstant               *string                `json:"valueInstant" bson:"value_instant"`                               // Value of Example (one of allowed types)
	ValueInteger               *int                   `json:"valueInteger" bson:"value_integer"`                               // Value of Example (one of allowed types)
	ValueInteger64             *int64                 `json:"valueInteger64" bson:"value_integer64"`                           // Value of Example (one of allowed types)
	ValueMarkdown              *string                `json:"valueMarkdown" bson:"value_markdown"`                             // Value of Example (one of allowed types)
	ValueOid                   *string                `json:"valueOid" bson:"value_oid"`                                       // Value of Example (one of allowed types)
	ValuePositiveInt           *int                   `json:"valuePositiveInt" bson:"value_positive_int"`                      // Value of Example (one of allowed types)
	ValueString                *string                `json:"valueString" bson:"value_string"`                                 // Value of Example (one of allowed types)
	ValueTime                  *string                `json:"valueTime" bson:"value_time"`                                     // Value of Example (one of allowed types)
	ValueUnsignedInt           *int                   `json:"valueUnsignedInt" bson:"value_unsigned_int"`                      // Value of Example (one of allowed types)
	ValueUri                   *string                `json:"valueUri" bson:"value_uri"`                                       // Value of Example (one of allowed types)
	ValueUrl                   *string                `json:"valueUrl" bson:"value_url"`                                       // Value of Example (one of allowed types)
	ValueUuid                  *uuid                  `json:"valueUuid" bson:"value_uuid"`                                     // Value of Example (one of allowed types)
	ValueAddress               *Address               `json:"valueAddress" bson:"value_address"`                               // Value of Example (one of allowed types)
	ValueAge                   *Age                   `json:"valueAge" bson:"value_age"`                                       // Value of Example (one of allowed types)
	ValueAnnotation            *Annotation            `json:"valueAnnotation" bson:"value_annotation"`                         // Value of Example (one of allowed types)
	ValueAttachment            *Attachment            `json:"valueAttachment" bson:"value_attachment"`                         // Value of Example (one of allowed types)
	ValueCodeableConcept       *CodeableConcept       `json:"valueCodeableConcept" bson:"value_codeable_concept"`              // Value of Example (one of allowed types)
	ValueCodeableReference     *CodeableReference     `json:"valueCodeableReference" bson:"value_codeable_reference"`          // Value of Example (one of allowed types)
	ValueCoding                *Coding                `json:"valueCoding" bson:"value_coding"`                                 // Value of Example (one of allowed types)
	ValueContactPoint          *ContactPoint          `json:"valueContactPoint" bson:"value_contact_point"`                    // Value of Example (one of allowed types)
	ValueCount                 *Count                 `json:"valueCount" bson:"value_count"`                                   // Value of Example (one of allowed types)
	ValueDistance              *Distance              `json:"valueDistance" bson:"value_distance"`                             // Value of Example (one of allowed types)
	ValueDuration              *Duration              `json:"valueDuration" bson:"value_duration"`                             // Value of Example (one of allowed types)
	ValueHumanName             *HumanName             `json:"valueHumanName" bson:"value_human_name"`                          // Value of Example (one of allowed types)
	ValueIdentifier            *Identifier            `json:"valueIdentifier" bson:"value_identifier"`                         // Value of Example (one of allowed types)
	ValueMoney                 *Money                 `json:"valueMoney" bson:"value_money"`                                   // Value of Example (one of allowed types)
	ValuePeriod                *Period                `json:"valuePeriod" bson:"value_period"`                                 // Value of Example (one of allowed types)
	ValueQuantity              *Quantity              `json:"valueQuantity" bson:"value_quantity"`                             // Value of Example (one of allowed types)
	ValueRange                 *Range                 `json:"valueRange" bson:"value_range"`                                   // Value of Example (one of allowed types)
	ValueRatio                 *Ratio                 `json:"valueRatio" bson:"value_ratio"`                                   // Value of Example (one of allowed types)
	ValueRatioRange            *RatioRange            `json:"valueRatioRange" bson:"value_ratio_range"`                        // Value of Example (one of allowed types)
	ValueReference             *Reference             `json:"valueReference" bson:"value_reference"`                           // Value of Example (one of allowed types)
	ValueSampledData           *SampledData           `json:"valueSampledData" bson:"value_sampled_data"`                      // Value of Example (one of allowed types)
	ValueSignature             *Signature             `json:"valueSignature" bson:"value_signature"`                           // Value of Example (one of allowed types)
	ValueTiming                *Timing                `json:"valueTiming" bson:"value_timing"`                                 // Value of Example (one of allowed types)
	ValueContactDetail         *ContactDetail         `json:"valueContactDetail" bson:"value_contact_detail"`                  // Value of Example (one of allowed types)
	ValueDataRequirement       *DataRequirement       `json:"valueDataRequirement" bson:"value_data_requirement"`              // Value of Example (one of allowed types)
	ValueExpression            *Expression            `json:"valueExpression" bson:"value_expression"`                         // Value of Example (one of allowed types)
	ValueParameterDefinition   *ParameterDefinition   `json:"valueParameterDefinition" bson:"value_parameter_definition"`      // Value of Example (one of allowed types)
	ValueRelatedArtifact       *RelatedArtifact       `json:"valueRelatedArtifact" bson:"value_related_artifact"`              // Value of Example (one of allowed types)
	ValueTriggerDefinition     *TriggerDefinition     `json:"valueTriggerDefinition" bson:"value_trigger_definition"`          // Value of Example (one of allowed types)
	ValueUsageContext          *UsageContext          `json:"valueUsageContext" bson:"value_usage_context"`                    // Value of Example (one of allowed types)
	ValueAvailability          *Availability          `json:"valueAvailability" bson:"value_availability"`                     // Value of Example (one of allowed types)
	ValueExtendedContactDetail *ExtendedContactDetail `json:"valueExtendedContactDetail" bson:"value_extended_contact_detail"` // Value of Example (one of allowed types)
	ValueVirtualServiceDetail  *VirtualServiceDetail  `json:"valueVirtualServiceDetail" bson:"value_virtual_service_detail"`   // Value of Example (one of allowed types)
	ValueDosage                *Dosage                `json:"valueDosage" bson:"value_dosage"`                                 // Value of Example (one of allowed types)
	ValueMeta                  *Meta                  `json:"valueMeta" bson:"value_meta"`                                     // Value of Example (one of allowed types)
}

func (r *ElementDefinitionExample) Validate() error {
	var emptyString string
	if r.Label == emptyString {
		return fmt.Errorf("field 'Label' is required")
	}
	if r.ValueBase64Binary == nil {
		return fmt.Errorf("field 'ValueBase64Binary' is required")
	}
	if r.ValueBoolean == nil {
		return fmt.Errorf("field 'ValueBoolean' is required")
	}
	if r.ValueCanonical == nil {
		return fmt.Errorf("field 'ValueCanonical' is required")
	}
	if r.ValueCode == nil {
		return fmt.Errorf("field 'ValueCode' is required")
	}
	if r.ValueDate == nil {
		return fmt.Errorf("field 'ValueDate' is required")
	}
	if r.ValueDateTime == nil {
		return fmt.Errorf("field 'ValueDateTime' is required")
	}
	if r.ValueDecimal == nil {
		return fmt.Errorf("field 'ValueDecimal' is required")
	}
	if r.ValueId == nil {
		return fmt.Errorf("field 'ValueId' is required")
	}
	if r.ValueInstant == nil {
		return fmt.Errorf("field 'ValueInstant' is required")
	}
	if r.ValueInteger == nil {
		return fmt.Errorf("field 'ValueInteger' is required")
	}
	if r.ValueInteger64 == nil {
		return fmt.Errorf("field 'ValueInteger64' is required")
	}
	if r.ValueMarkdown == nil {
		return fmt.Errorf("field 'ValueMarkdown' is required")
	}
	if r.ValueOid == nil {
		return fmt.Errorf("field 'ValueOid' is required")
	}
	if r.ValuePositiveInt == nil {
		return fmt.Errorf("field 'ValuePositiveInt' is required")
	}
	if r.ValueString == nil {
		return fmt.Errorf("field 'ValueString' is required")
	}
	if r.ValueTime == nil {
		return fmt.Errorf("field 'ValueTime' is required")
	}
	if r.ValueUnsignedInt == nil {
		return fmt.Errorf("field 'ValueUnsignedInt' is required")
	}
	if r.ValueUri == nil {
		return fmt.Errorf("field 'ValueUri' is required")
	}
	if r.ValueUrl == nil {
		return fmt.Errorf("field 'ValueUrl' is required")
	}
	if r.ValueUuid == nil {
		return fmt.Errorf("field 'ValueUuid' is required")
	}
	if r.ValueUuid != nil {
		if err := r.ValueUuid.Validate(); err != nil {
			return fmt.Errorf("ValueUuid: %w", err)
		}
	}
	if r.ValueAddress == nil {
		return fmt.Errorf("field 'ValueAddress' is required")
	}
	if r.ValueAddress != nil {
		if err := r.ValueAddress.Validate(); err != nil {
			return fmt.Errorf("ValueAddress: %w", err)
		}
	}
	if r.ValueAge == nil {
		return fmt.Errorf("field 'ValueAge' is required")
	}
	if r.ValueAge != nil {
		if err := r.ValueAge.Validate(); err != nil {
			return fmt.Errorf("ValueAge: %w", err)
		}
	}
	if r.ValueAnnotation == nil {
		return fmt.Errorf("field 'ValueAnnotation' is required")
	}
	if r.ValueAnnotation != nil {
		if err := r.ValueAnnotation.Validate(); err != nil {
			return fmt.Errorf("ValueAnnotation: %w", err)
		}
	}
	if r.ValueAttachment == nil {
		return fmt.Errorf("field 'ValueAttachment' is required")
	}
	if r.ValueAttachment != nil {
		if err := r.ValueAttachment.Validate(); err != nil {
			return fmt.Errorf("ValueAttachment: %w", err)
		}
	}
	if r.ValueCodeableConcept == nil {
		return fmt.Errorf("field 'ValueCodeableConcept' is required")
	}
	if r.ValueCodeableConcept != nil {
		if err := r.ValueCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("ValueCodeableConcept: %w", err)
		}
	}
	if r.ValueCodeableReference == nil {
		return fmt.Errorf("field 'ValueCodeableReference' is required")
	}
	if r.ValueCodeableReference != nil {
		if err := r.ValueCodeableReference.Validate(); err != nil {
			return fmt.Errorf("ValueCodeableReference: %w", err)
		}
	}
	if r.ValueCoding == nil {
		return fmt.Errorf("field 'ValueCoding' is required")
	}
	if r.ValueCoding != nil {
		if err := r.ValueCoding.Validate(); err != nil {
			return fmt.Errorf("ValueCoding: %w", err)
		}
	}
	if r.ValueContactPoint == nil {
		return fmt.Errorf("field 'ValueContactPoint' is required")
	}
	if r.ValueContactPoint != nil {
		if err := r.ValueContactPoint.Validate(); err != nil {
			return fmt.Errorf("ValueContactPoint: %w", err)
		}
	}
	if r.ValueCount == nil {
		return fmt.Errorf("field 'ValueCount' is required")
	}
	if r.ValueCount != nil {
		if err := r.ValueCount.Validate(); err != nil {
			return fmt.Errorf("ValueCount: %w", err)
		}
	}
	if r.ValueDistance == nil {
		return fmt.Errorf("field 'ValueDistance' is required")
	}
	if r.ValueDistance != nil {
		if err := r.ValueDistance.Validate(); err != nil {
			return fmt.Errorf("ValueDistance: %w", err)
		}
	}
	if r.ValueDuration == nil {
		return fmt.Errorf("field 'ValueDuration' is required")
	}
	if r.ValueDuration != nil {
		if err := r.ValueDuration.Validate(); err != nil {
			return fmt.Errorf("ValueDuration: %w", err)
		}
	}
	if r.ValueHumanName == nil {
		return fmt.Errorf("field 'ValueHumanName' is required")
	}
	if r.ValueHumanName != nil {
		if err := r.ValueHumanName.Validate(); err != nil {
			return fmt.Errorf("ValueHumanName: %w", err)
		}
	}
	if r.ValueIdentifier == nil {
		return fmt.Errorf("field 'ValueIdentifier' is required")
	}
	if r.ValueIdentifier != nil {
		if err := r.ValueIdentifier.Validate(); err != nil {
			return fmt.Errorf("ValueIdentifier: %w", err)
		}
	}
	if r.ValueMoney == nil {
		return fmt.Errorf("field 'ValueMoney' is required")
	}
	if r.ValueMoney != nil {
		if err := r.ValueMoney.Validate(); err != nil {
			return fmt.Errorf("ValueMoney: %w", err)
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
	if r.ValueRatioRange == nil {
		return fmt.Errorf("field 'ValueRatioRange' is required")
	}
	if r.ValueRatioRange != nil {
		if err := r.ValueRatioRange.Validate(); err != nil {
			return fmt.Errorf("ValueRatioRange: %w", err)
		}
	}
	if r.ValueReference == nil {
		return fmt.Errorf("field 'ValueReference' is required")
	}
	if r.ValueReference != nil {
		if err := r.ValueReference.Validate(); err != nil {
			return fmt.Errorf("ValueReference: %w", err)
		}
	}
	if r.ValueSampledData == nil {
		return fmt.Errorf("field 'ValueSampledData' is required")
	}
	if r.ValueSampledData != nil {
		if err := r.ValueSampledData.Validate(); err != nil {
			return fmt.Errorf("ValueSampledData: %w", err)
		}
	}
	if r.ValueSignature == nil {
		return fmt.Errorf("field 'ValueSignature' is required")
	}
	if r.ValueSignature != nil {
		if err := r.ValueSignature.Validate(); err != nil {
			return fmt.Errorf("ValueSignature: %w", err)
		}
	}
	if r.ValueTiming == nil {
		return fmt.Errorf("field 'ValueTiming' is required")
	}
	if r.ValueTiming != nil {
		if err := r.ValueTiming.Validate(); err != nil {
			return fmt.Errorf("ValueTiming: %w", err)
		}
	}
	if r.ValueContactDetail == nil {
		return fmt.Errorf("field 'ValueContactDetail' is required")
	}
	if r.ValueContactDetail != nil {
		if err := r.ValueContactDetail.Validate(); err != nil {
			return fmt.Errorf("ValueContactDetail: %w", err)
		}
	}
	if r.ValueDataRequirement == nil {
		return fmt.Errorf("field 'ValueDataRequirement' is required")
	}
	if r.ValueDataRequirement != nil {
		if err := r.ValueDataRequirement.Validate(); err != nil {
			return fmt.Errorf("ValueDataRequirement: %w", err)
		}
	}
	if r.ValueExpression == nil {
		return fmt.Errorf("field 'ValueExpression' is required")
	}
	if r.ValueExpression != nil {
		if err := r.ValueExpression.Validate(); err != nil {
			return fmt.Errorf("ValueExpression: %w", err)
		}
	}
	if r.ValueParameterDefinition == nil {
		return fmt.Errorf("field 'ValueParameterDefinition' is required")
	}
	if r.ValueParameterDefinition != nil {
		if err := r.ValueParameterDefinition.Validate(); err != nil {
			return fmt.Errorf("ValueParameterDefinition: %w", err)
		}
	}
	if r.ValueRelatedArtifact == nil {
		return fmt.Errorf("field 'ValueRelatedArtifact' is required")
	}
	if r.ValueRelatedArtifact != nil {
		if err := r.ValueRelatedArtifact.Validate(); err != nil {
			return fmt.Errorf("ValueRelatedArtifact: %w", err)
		}
	}
	if r.ValueTriggerDefinition == nil {
		return fmt.Errorf("field 'ValueTriggerDefinition' is required")
	}
	if r.ValueTriggerDefinition != nil {
		if err := r.ValueTriggerDefinition.Validate(); err != nil {
			return fmt.Errorf("ValueTriggerDefinition: %w", err)
		}
	}
	if r.ValueUsageContext == nil {
		return fmt.Errorf("field 'ValueUsageContext' is required")
	}
	if r.ValueUsageContext != nil {
		if err := r.ValueUsageContext.Validate(); err != nil {
			return fmt.Errorf("ValueUsageContext: %w", err)
		}
	}
	if r.ValueAvailability == nil {
		return fmt.Errorf("field 'ValueAvailability' is required")
	}
	if r.ValueAvailability != nil {
		if err := r.ValueAvailability.Validate(); err != nil {
			return fmt.Errorf("ValueAvailability: %w", err)
		}
	}
	if r.ValueExtendedContactDetail == nil {
		return fmt.Errorf("field 'ValueExtendedContactDetail' is required")
	}
	if r.ValueExtendedContactDetail != nil {
		if err := r.ValueExtendedContactDetail.Validate(); err != nil {
			return fmt.Errorf("ValueExtendedContactDetail: %w", err)
		}
	}
	if r.ValueVirtualServiceDetail == nil {
		return fmt.Errorf("field 'ValueVirtualServiceDetail' is required")
	}
	if r.ValueVirtualServiceDetail != nil {
		if err := r.ValueVirtualServiceDetail.Validate(); err != nil {
			return fmt.Errorf("ValueVirtualServiceDetail: %w", err)
		}
	}
	if r.ValueDosage == nil {
		return fmt.Errorf("field 'ValueDosage' is required")
	}
	if r.ValueDosage != nil {
		if err := r.ValueDosage.Validate(); err != nil {
			return fmt.Errorf("ValueDosage: %w", err)
		}
	}
	if r.ValueMeta == nil {
		return fmt.Errorf("field 'ValueMeta' is required")
	}
	if r.ValueMeta != nil {
		if err := r.ValueMeta.Validate(); err != nil {
			return fmt.Errorf("ValueMeta: %w", err)
		}
	}
	return nil
}

type ElementDefinitionMapping struct {
	Id       *string `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Identity string  `json:"identity" bson:"identity"`                     // Reference to mapping declaration
	Language *string `json:"language,omitempty" bson:"language,omitempty"` // Computable language of mapping
	Map      string  `json:"map" bson:"map"`                               // Details of the mapping
	Comment  *string `json:"comment,omitempty" bson:"comment,omitempty"`   // Comments about the mapping or its use
}

func (r *ElementDefinitionMapping) Validate() error {
	var emptyString string
	if r.Identity == emptyString {
		return fmt.Errorf("field 'Identity' is required")
	}
	if r.Map == emptyString {
		return fmt.Errorf("field 'Map' is required")
	}
	return nil
}

type ElementDefinitionConstraint struct {
	Id           *string `json:"id,omitempty" bson:"id,omitempty"`                     // Unique id for inter-element referencing
	Key          string  `json:"key" bson:"key"`                                       // Target of 'condition' reference above
	Requirements *string `json:"requirements,omitempty" bson:"requirements,omitempty"` // Why this constraint is necessary or appropriate
	Severity     string  `json:"severity" bson:"severity"`                             // error | warning
	Suppress     bool    `json:"suppress,omitempty" bson:"suppress,omitempty"`         // Suppress warning or hint in profile
	Human        string  `json:"human" bson:"human"`                                   // Human description of constraint
	Expression   *string `json:"expression,omitempty" bson:"expression,omitempty"`     // FHIRPath expression of constraint
	Source       *string `json:"source,omitempty" bson:"source,omitempty"`             // Reference to original source of constraint
}

func (r *ElementDefinitionConstraint) Validate() error {
	var emptyString string
	if r.Key == emptyString {
		return fmt.Errorf("field 'Key' is required")
	}
	if r.Severity == emptyString {
		return fmt.Errorf("field 'Severity' is required")
	}
	if r.Human == emptyString {
		return fmt.Errorf("field 'Human' is required")
	}
	return nil
}

type ElementDefinitionBinding struct {
	Id          *string                              `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Strength    string                               `json:"strength" bson:"strength"`                           // required | extensible | preferred | example | descriptive
	Description *string                              `json:"description,omitempty" bson:"description,omitempty"` // Guidance on the codes to be used
	ValueSet    *string                              `json:"valueSet,omitempty" bson:"value_set,omitempty"`      // Source of value set
	Additional  []ElementDefinitionBindingAdditional `json:"additional,omitempty" bson:"additional,omitempty"`   // Additional Bindings - more rules about the binding
}

func (r *ElementDefinitionBinding) Validate() error {
	var emptyString string
	if r.Strength == emptyString {
		return fmt.Errorf("field 'Strength' is required")
	}
	for i, item := range r.Additional {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Additional[%d]: %w", i, err)
		}
	}
	return nil
}

type ElementDefinitionBindingAdditional struct {
	Id            *string        `json:"id,omitempty" bson:"id,omitempty"`                       // Unique id for inter-element referencing
	Key           *string        `json:"key,omitempty" bson:"key,omitempty"`                     // Unique identifier so additional bindings to be matched across profiles
	Purpose       string         `json:"purpose" bson:"purpose"`                                 // maximum | minimum | required | extensible | candidate | current | current-extensible | best-practice | preferred | ui | starter | component
	ValueSet      string         `json:"valueSet" bson:"value_set"`                              // The value set for the additional binding
	Documentation *string        `json:"documentation,omitempty" bson:"documentation,omitempty"` // Documentation of the purpose of use of the binding
	ShortDoco     *string        `json:"shortDoco,omitempty" bson:"short_doco,omitempty"`        // Concise documentation - for summary tables
	Usage         []UsageContext `json:"usage,omitempty" bson:"usage,omitempty"`                 // Qualifies the usage - jurisdiction, gender, workflow status etc.
	Any           bool           `json:"any,omitempty" bson:"any,omitempty"`                     // Whether binding can applies to all repeats, or just one
}

func (r *ElementDefinitionBindingAdditional) Validate() error {
	var emptyString string
	if r.Purpose == emptyString {
		return fmt.Errorf("field 'Purpose' is required")
	}
	if r.ValueSet == emptyString {
		return fmt.Errorf("field 'ValueSet' is required")
	}
	for i, item := range r.Usage {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Usage[%d]: %w", i, err)
		}
	}
	return nil
}
