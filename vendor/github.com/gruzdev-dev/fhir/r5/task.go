package models

import (
	"encoding/json"
	"fmt"
)

// A task to be performed as a part of a workflow and the related information like inputs, outputs and execution progress.
type Task struct {
	ResourceType       string              `json:"resourceType" bson:"resource_type"`                                 // Type of resource
	Id                 *string             `json:"id,omitempty" bson:"id,omitempty"`                                  // Logical id of this artifact
	Meta               *Meta               `json:"meta,omitempty" bson:"meta,omitempty"`                              // Metadata about the resource
	ImplicitRules      *string             `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`           // A set of rules under which this content was created
	Language           *string             `json:"language,omitempty" bson:"language,omitempty"`                      // Language of the resource content
	Text               *Narrative          `json:"text,omitempty" bson:"text,omitempty"`                              // Text summary of the resource, for human interpretation
	Contained          []json.RawMessage   `json:"contained,omitempty" bson:"contained,omitempty"`                    // Contained, inline Resources
	Identifier         []Identifier        `json:"identifier,omitempty" bson:"identifier,omitempty"`                  // Task Instance Identifier
	BasedOn            []Reference         `json:"basedOn,omitempty" bson:"based_on,omitempty"`                       // Request fulfilled by this task
	GroupIdentifier    *Identifier         `json:"groupIdentifier,omitempty" bson:"group_identifier,omitempty"`       // Requisition or grouper id
	PartOf             []Reference         `json:"partOf,omitempty" bson:"part_of,omitempty"`                         // Composite task
	Status             string              `json:"status" bson:"status"`                                              // draft | requested | received | accepted | +
	StatusReason       []CodeableReference `json:"statusReason,omitempty" bson:"status_reason,omitempty"`             // Reason for current status
	BusinessStatus     *CodeableConcept    `json:"businessStatus,omitempty" bson:"business_status,omitempty"`         // E.g. "Specimen collected", "IV prepped"
	Intent             string              `json:"intent" bson:"intent"`                                              // unknown | proposal | plan | order | original-order | reflex-order | filler-order | instance-order | option
	Priority           *string             `json:"priority,omitempty" bson:"priority,omitempty"`                      // routine | urgent | asap | stat
	DoNotPerform       *bool               `json:"doNotPerform,omitempty" bson:"do_not_perform,omitempty"`            // True if Task is prohibiting action
	Code               *CodeableConcept    `json:"code,omitempty" bson:"code,omitempty"`                              // Task Type
	Description        *string             `json:"description,omitempty" bson:"description,omitempty"`                // Human-readable explanation of task
	Focus              []TaskFocus         `json:"focus,omitempty" bson:"focus,omitempty"`                            // What task is acting on
	For                *Reference          `json:"for,omitempty" bson:"for,omitempty"`                                // Beneficiary of the Task
	Encounter          *Reference          `json:"encounter,omitempty" bson:"encounter,omitempty"`                    // Healthcare event during which this task originated
	RequestedPeriod    *Period             `json:"requestedPeriod,omitempty" bson:"requested_period,omitempty"`       // When the task should be performed
	ExecutionPeriod    *Period             `json:"executionPeriod,omitempty" bson:"execution_period,omitempty"`       // Start and end time of execution
	AuthoredOn         *string             `json:"authoredOn,omitempty" bson:"authored_on,omitempty"`                 // Task Creation Date
	LastModified       *string             `json:"lastModified,omitempty" bson:"last_modified,omitempty"`             // Task Last Modified Date
	Requester          *Reference          `json:"requester,omitempty" bson:"requester,omitempty"`                    // Who is asking for task to be done
	RequestedPerformer []CodeableReference `json:"requestedPerformer,omitempty" bson:"requested_performer,omitempty"` // Who should perform the Task
	Owner              *Reference          `json:"owner,omitempty" bson:"owner,omitempty"`                            // Responsible individual
	Performer          []TaskPerformer     `json:"performer,omitempty" bson:"performer,omitempty"`                    // Who or what performed the task
	Location           *Reference          `json:"location,omitempty" bson:"location,omitempty"`                      // Where task occurs
	Reason             []CodeableReference `json:"reason,omitempty" bson:"reason,omitempty"`                          // Why task is needed
	Insurance          []Reference         `json:"insurance,omitempty" bson:"insurance,omitempty"`                    // Associated insurance coverage
	Note               []Annotation        `json:"note,omitempty" bson:"note,omitempty"`                              // Comments made about the task
	RelevantHistory    []Reference         `json:"relevantHistory,omitempty" bson:"relevant_history,omitempty"`       // Key events in history of the Task
	Restriction        *TaskRestriction    `json:"restriction,omitempty" bson:"restriction,omitempty"`                // Constraints on fulfillment tasks
	Input              []TaskInput         `json:"input,omitempty" bson:"input,omitempty"`                            // Information used to perform task
	Output             []TaskOutput        `json:"output,omitempty" bson:"output,omitempty"`                          // Information produced as part of task
}

func (r *Task) Validate() error {
	if r.ResourceType != "Task" {
		return fmt.Errorf("invalid resourceType: expected 'Task', got '%s'", r.ResourceType)
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
	for i, item := range r.BasedOn {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("BasedOn[%d]: %w", i, err)
		}
	}
	if r.GroupIdentifier != nil {
		if err := r.GroupIdentifier.Validate(); err != nil {
			return fmt.Errorf("GroupIdentifier: %w", err)
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
	for i, item := range r.StatusReason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("StatusReason[%d]: %w", i, err)
		}
	}
	if r.BusinessStatus != nil {
		if err := r.BusinessStatus.Validate(); err != nil {
			return fmt.Errorf("BusinessStatus: %w", err)
		}
	}
	if r.Intent == emptyString {
		return fmt.Errorf("field 'Intent' is required")
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	for i, item := range r.Focus {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Focus[%d]: %w", i, err)
		}
	}
	if r.For != nil {
		if err := r.For.Validate(); err != nil {
			return fmt.Errorf("For: %w", err)
		}
	}
	if r.Encounter != nil {
		if err := r.Encounter.Validate(); err != nil {
			return fmt.Errorf("Encounter: %w", err)
		}
	}
	if r.RequestedPeriod != nil {
		if err := r.RequestedPeriod.Validate(); err != nil {
			return fmt.Errorf("RequestedPeriod: %w", err)
		}
	}
	if r.ExecutionPeriod != nil {
		if err := r.ExecutionPeriod.Validate(); err != nil {
			return fmt.Errorf("ExecutionPeriod: %w", err)
		}
	}
	if r.Requester != nil {
		if err := r.Requester.Validate(); err != nil {
			return fmt.Errorf("Requester: %w", err)
		}
	}
	for i, item := range r.RequestedPerformer {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RequestedPerformer[%d]: %w", i, err)
		}
	}
	if r.Owner != nil {
		if err := r.Owner.Validate(); err != nil {
			return fmt.Errorf("Owner: %w", err)
		}
	}
	for i, item := range r.Performer {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Performer[%d]: %w", i, err)
		}
	}
	if r.Location != nil {
		if err := r.Location.Validate(); err != nil {
			return fmt.Errorf("Location: %w", err)
		}
	}
	for i, item := range r.Reason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reason[%d]: %w", i, err)
		}
	}
	for i, item := range r.Insurance {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Insurance[%d]: %w", i, err)
		}
	}
	for i, item := range r.Note {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Note[%d]: %w", i, err)
		}
	}
	for i, item := range r.RelevantHistory {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("RelevantHistory[%d]: %w", i, err)
		}
	}
	if r.Restriction != nil {
		if err := r.Restriction.Validate(); err != nil {
			return fmt.Errorf("Restriction: %w", err)
		}
	}
	for i, item := range r.Input {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Input[%d]: %w", i, err)
		}
	}
	for i, item := range r.Output {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Output[%d]: %w", i, err)
		}
	}
	return nil
}

type TaskOutput struct {
	Id                         *string                `json:"id,omitempty" bson:"id,omitempty"`                                // Unique id for inter-element referencing
	Type                       *CodeableConcept       `json:"type" bson:"type"`                                                // Label for output
	ValueBase64Binary          *string                `json:"valueBase64Binary" bson:"value_base64_binary"`                    // Result of output
	ValueBoolean               *bool                  `json:"valueBoolean" bson:"value_boolean"`                               // Result of output
	ValueCanonical             *string                `json:"valueCanonical" bson:"value_canonical"`                           // Result of output
	ValueCode                  *string                `json:"valueCode" bson:"value_code"`                                     // Result of output
	ValueDate                  *string                `json:"valueDate" bson:"value_date"`                                     // Result of output
	ValueDateTime              *string                `json:"valueDateTime" bson:"value_date_time"`                            // Result of output
	ValueDecimal               *float64               `json:"valueDecimal" bson:"value_decimal"`                               // Result of output
	ValueId                    *string                `json:"valueId" bson:"value_id"`                                         // Result of output
	ValueInstant               *string                `json:"valueInstant" bson:"value_instant"`                               // Result of output
	ValueInteger               *int                   `json:"valueInteger" bson:"value_integer"`                               // Result of output
	ValueInteger64             *int64                 `json:"valueInteger64" bson:"value_integer64"`                           // Result of output
	ValueMarkdown              *string                `json:"valueMarkdown" bson:"value_markdown"`                             // Result of output
	ValueOid                   *string                `json:"valueOid" bson:"value_oid"`                                       // Result of output
	ValuePositiveInt           *int                   `json:"valuePositiveInt" bson:"value_positive_int"`                      // Result of output
	ValueString                *string                `json:"valueString" bson:"value_string"`                                 // Result of output
	ValueTime                  *string                `json:"valueTime" bson:"value_time"`                                     // Result of output
	ValueUnsignedInt           *int                   `json:"valueUnsignedInt" bson:"value_unsigned_int"`                      // Result of output
	ValueUri                   *string                `json:"valueUri" bson:"value_uri"`                                       // Result of output
	ValueUrl                   *string                `json:"valueUrl" bson:"value_url"`                                       // Result of output
	ValueUuid                  *uuid                  `json:"valueUuid" bson:"value_uuid"`                                     // Result of output
	ValueAddress               *Address               `json:"valueAddress" bson:"value_address"`                               // Result of output
	ValueAge                   *Age                   `json:"valueAge" bson:"value_age"`                                       // Result of output
	ValueAnnotation            *Annotation            `json:"valueAnnotation" bson:"value_annotation"`                         // Result of output
	ValueAttachment            *Attachment            `json:"valueAttachment" bson:"value_attachment"`                         // Result of output
	ValueCodeableConcept       *CodeableConcept       `json:"valueCodeableConcept" bson:"value_codeable_concept"`              // Result of output
	ValueCodeableReference     *CodeableReference     `json:"valueCodeableReference" bson:"value_codeable_reference"`          // Result of output
	ValueCoding                *Coding                `json:"valueCoding" bson:"value_coding"`                                 // Result of output
	ValueContactPoint          *ContactPoint          `json:"valueContactPoint" bson:"value_contact_point"`                    // Result of output
	ValueCount                 *Count                 `json:"valueCount" bson:"value_count"`                                   // Result of output
	ValueDistance              *Distance              `json:"valueDistance" bson:"value_distance"`                             // Result of output
	ValueDuration              *Duration              `json:"valueDuration" bson:"value_duration"`                             // Result of output
	ValueHumanName             *HumanName             `json:"valueHumanName" bson:"value_human_name"`                          // Result of output
	ValueIdentifier            *Identifier            `json:"valueIdentifier" bson:"value_identifier"`                         // Result of output
	ValueMoney                 *Money                 `json:"valueMoney" bson:"value_money"`                                   // Result of output
	ValuePeriod                *Period                `json:"valuePeriod" bson:"value_period"`                                 // Result of output
	ValueQuantity              *Quantity              `json:"valueQuantity" bson:"value_quantity"`                             // Result of output
	ValueRange                 *Range                 `json:"valueRange" bson:"value_range"`                                   // Result of output
	ValueRatio                 *Ratio                 `json:"valueRatio" bson:"value_ratio"`                                   // Result of output
	ValueRatioRange            *RatioRange            `json:"valueRatioRange" bson:"value_ratio_range"`                        // Result of output
	ValueReference             *Reference             `json:"valueReference" bson:"value_reference"`                           // Result of output
	ValueSampledData           *SampledData           `json:"valueSampledData" bson:"value_sampled_data"`                      // Result of output
	ValueSignature             *Signature             `json:"valueSignature" bson:"value_signature"`                           // Result of output
	ValueTiming                *Timing                `json:"valueTiming" bson:"value_timing"`                                 // Result of output
	ValueContactDetail         *ContactDetail         `json:"valueContactDetail" bson:"value_contact_detail"`                  // Result of output
	ValueDataRequirement       *DataRequirement       `json:"valueDataRequirement" bson:"value_data_requirement"`              // Result of output
	ValueExpression            *Expression            `json:"valueExpression" bson:"value_expression"`                         // Result of output
	ValueParameterDefinition   *ParameterDefinition   `json:"valueParameterDefinition" bson:"value_parameter_definition"`      // Result of output
	ValueRelatedArtifact       *RelatedArtifact       `json:"valueRelatedArtifact" bson:"value_related_artifact"`              // Result of output
	ValueTriggerDefinition     *TriggerDefinition     `json:"valueTriggerDefinition" bson:"value_trigger_definition"`          // Result of output
	ValueUsageContext          *UsageContext          `json:"valueUsageContext" bson:"value_usage_context"`                    // Result of output
	ValueAvailability          *Availability          `json:"valueAvailability" bson:"value_availability"`                     // Result of output
	ValueExtendedContactDetail *ExtendedContactDetail `json:"valueExtendedContactDetail" bson:"value_extended_contact_detail"` // Result of output
	ValueVirtualServiceDetail  *VirtualServiceDetail  `json:"valueVirtualServiceDetail" bson:"value_virtual_service_detail"`   // Result of output
	ValueDosage                *Dosage                `json:"valueDosage" bson:"value_dosage"`                                 // Result of output
	ValueMeta                  *Meta                  `json:"valueMeta" bson:"value_meta"`                                     // Result of output
}

func (r *TaskOutput) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
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

type TaskFocus struct {
	Id             *string    `json:"id,omitempty" bson:"id,omitempty"`      // Unique id for inter-element referencing
	ValueReference *Reference `json:"valueReference" bson:"value_reference"` // What task is acting on
	ValueCanonical *string    `json:"valueCanonical" bson:"value_canonical"` // What task is acting on
}

func (r *TaskFocus) Validate() error {
	if r.ValueReference == nil {
		return fmt.Errorf("field 'ValueReference' is required")
	}
	if r.ValueReference != nil {
		if err := r.ValueReference.Validate(); err != nil {
			return fmt.Errorf("ValueReference: %w", err)
		}
	}
	if r.ValueCanonical == nil {
		return fmt.Errorf("field 'ValueCanonical' is required")
	}
	return nil
}

type TaskPerformer struct {
	Id       *string          `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Function *CodeableConcept `json:"function,omitempty" bson:"function,omitempty"` // Type of performance
	Actor    *Reference       `json:"actor" bson:"actor"`                           // Who performed the task
}

func (r *TaskPerformer) Validate() error {
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

type TaskRestriction struct {
	Id          *string     `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Repetitions *int        `json:"repetitions,omitempty" bson:"repetitions,omitempty"` // How many times to repeat
	Period      *Period     `json:"period,omitempty" bson:"period,omitempty"`           // When fulfillment is sought
	Recipient   []Reference `json:"recipient,omitempty" bson:"recipient,omitempty"`     // Individual or entity from whom fulfillment is being sought
}

func (r *TaskRestriction) Validate() error {
	if r.Period != nil {
		if err := r.Period.Validate(); err != nil {
			return fmt.Errorf("Period: %w", err)
		}
	}
	for i, item := range r.Recipient {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Recipient[%d]: %w", i, err)
		}
	}
	return nil
}

type TaskInput struct {
	Id                         *string                `json:"id,omitempty" bson:"id,omitempty"`                                // Unique id for inter-element referencing
	Type                       *CodeableConcept       `json:"type" bson:"type"`                                                // Label for the input
	ValueBase64Binary          *string                `json:"valueBase64Binary" bson:"value_base64_binary"`                    // Content to use in performing the task
	ValueBoolean               *bool                  `json:"valueBoolean" bson:"value_boolean"`                               // Content to use in performing the task
	ValueCanonical             *string                `json:"valueCanonical" bson:"value_canonical"`                           // Content to use in performing the task
	ValueCode                  *string                `json:"valueCode" bson:"value_code"`                                     // Content to use in performing the task
	ValueDate                  *string                `json:"valueDate" bson:"value_date"`                                     // Content to use in performing the task
	ValueDateTime              *string                `json:"valueDateTime" bson:"value_date_time"`                            // Content to use in performing the task
	ValueDecimal               *float64               `json:"valueDecimal" bson:"value_decimal"`                               // Content to use in performing the task
	ValueId                    *string                `json:"valueId" bson:"value_id"`                                         // Content to use in performing the task
	ValueInstant               *string                `json:"valueInstant" bson:"value_instant"`                               // Content to use in performing the task
	ValueInteger               *int                   `json:"valueInteger" bson:"value_integer"`                               // Content to use in performing the task
	ValueInteger64             *int64                 `json:"valueInteger64" bson:"value_integer64"`                           // Content to use in performing the task
	ValueMarkdown              *string                `json:"valueMarkdown" bson:"value_markdown"`                             // Content to use in performing the task
	ValueOid                   *string                `json:"valueOid" bson:"value_oid"`                                       // Content to use in performing the task
	ValuePositiveInt           *int                   `json:"valuePositiveInt" bson:"value_positive_int"`                      // Content to use in performing the task
	ValueString                *string                `json:"valueString" bson:"value_string"`                                 // Content to use in performing the task
	ValueTime                  *string                `json:"valueTime" bson:"value_time"`                                     // Content to use in performing the task
	ValueUnsignedInt           *int                   `json:"valueUnsignedInt" bson:"value_unsigned_int"`                      // Content to use in performing the task
	ValueUri                   *string                `json:"valueUri" bson:"value_uri"`                                       // Content to use in performing the task
	ValueUrl                   *string                `json:"valueUrl" bson:"value_url"`                                       // Content to use in performing the task
	ValueUuid                  *uuid                  `json:"valueUuid" bson:"value_uuid"`                                     // Content to use in performing the task
	ValueAddress               *Address               `json:"valueAddress" bson:"value_address"`                               // Content to use in performing the task
	ValueAge                   *Age                   `json:"valueAge" bson:"value_age"`                                       // Content to use in performing the task
	ValueAnnotation            *Annotation            `json:"valueAnnotation" bson:"value_annotation"`                         // Content to use in performing the task
	ValueAttachment            *Attachment            `json:"valueAttachment" bson:"value_attachment"`                         // Content to use in performing the task
	ValueCodeableConcept       *CodeableConcept       `json:"valueCodeableConcept" bson:"value_codeable_concept"`              // Content to use in performing the task
	ValueCodeableReference     *CodeableReference     `json:"valueCodeableReference" bson:"value_codeable_reference"`          // Content to use in performing the task
	ValueCoding                *Coding                `json:"valueCoding" bson:"value_coding"`                                 // Content to use in performing the task
	ValueContactPoint          *ContactPoint          `json:"valueContactPoint" bson:"value_contact_point"`                    // Content to use in performing the task
	ValueCount                 *Count                 `json:"valueCount" bson:"value_count"`                                   // Content to use in performing the task
	ValueDistance              *Distance              `json:"valueDistance" bson:"value_distance"`                             // Content to use in performing the task
	ValueDuration              *Duration              `json:"valueDuration" bson:"value_duration"`                             // Content to use in performing the task
	ValueHumanName             *HumanName             `json:"valueHumanName" bson:"value_human_name"`                          // Content to use in performing the task
	ValueIdentifier            *Identifier            `json:"valueIdentifier" bson:"value_identifier"`                         // Content to use in performing the task
	ValueMoney                 *Money                 `json:"valueMoney" bson:"value_money"`                                   // Content to use in performing the task
	ValuePeriod                *Period                `json:"valuePeriod" bson:"value_period"`                                 // Content to use in performing the task
	ValueQuantity              *Quantity              `json:"valueQuantity" bson:"value_quantity"`                             // Content to use in performing the task
	ValueRange                 *Range                 `json:"valueRange" bson:"value_range"`                                   // Content to use in performing the task
	ValueRatio                 *Ratio                 `json:"valueRatio" bson:"value_ratio"`                                   // Content to use in performing the task
	ValueRatioRange            *RatioRange            `json:"valueRatioRange" bson:"value_ratio_range"`                        // Content to use in performing the task
	ValueReference             *Reference             `json:"valueReference" bson:"value_reference"`                           // Content to use in performing the task
	ValueSampledData           *SampledData           `json:"valueSampledData" bson:"value_sampled_data"`                      // Content to use in performing the task
	ValueSignature             *Signature             `json:"valueSignature" bson:"value_signature"`                           // Content to use in performing the task
	ValueTiming                *Timing                `json:"valueTiming" bson:"value_timing"`                                 // Content to use in performing the task
	ValueContactDetail         *ContactDetail         `json:"valueContactDetail" bson:"value_contact_detail"`                  // Content to use in performing the task
	ValueDataRequirement       *DataRequirement       `json:"valueDataRequirement" bson:"value_data_requirement"`              // Content to use in performing the task
	ValueExpression            *Expression            `json:"valueExpression" bson:"value_expression"`                         // Content to use in performing the task
	ValueParameterDefinition   *ParameterDefinition   `json:"valueParameterDefinition" bson:"value_parameter_definition"`      // Content to use in performing the task
	ValueRelatedArtifact       *RelatedArtifact       `json:"valueRelatedArtifact" bson:"value_related_artifact"`              // Content to use in performing the task
	ValueTriggerDefinition     *TriggerDefinition     `json:"valueTriggerDefinition" bson:"value_trigger_definition"`          // Content to use in performing the task
	ValueUsageContext          *UsageContext          `json:"valueUsageContext" bson:"value_usage_context"`                    // Content to use in performing the task
	ValueAvailability          *Availability          `json:"valueAvailability" bson:"value_availability"`                     // Content to use in performing the task
	ValueExtendedContactDetail *ExtendedContactDetail `json:"valueExtendedContactDetail" bson:"value_extended_contact_detail"` // Content to use in performing the task
	ValueVirtualServiceDetail  *VirtualServiceDetail  `json:"valueVirtualServiceDetail" bson:"value_virtual_service_detail"`   // Content to use in performing the task
	ValueDosage                *Dosage                `json:"valueDosage" bson:"value_dosage"`                                 // Content to use in performing the task
	ValueMeta                  *Meta                  `json:"valueMeta" bson:"value_meta"`                                     // Content to use in performing the task
}

func (r *TaskInput) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
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
