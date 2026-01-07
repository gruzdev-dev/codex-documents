package models

import (
	"encoding/json"
	"fmt"
)

// A provider issued list of professional services and products which have been provided, or are to be provided, to a patient which is sent to an insurer for reimbursement.
type Claim struct {
	Id                    *string               `json:"id,omitempty" bson:"id,omitempty"`                                         // Logical id of this artifact
	Meta                  *Meta                 `json:"meta,omitempty" bson:"meta,omitempty"`                                     // Metadata about the resource
	ImplicitRules         *string               `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                  // A set of rules under which this content was created
	Language              *string               `json:"language,omitempty" bson:"language,omitempty"`                             // Language of the resource content
	Text                  *Narrative            `json:"text,omitempty" bson:"text,omitempty"`                                     // Text summary of the resource, for human interpretation
	Contained             []json.RawMessage     `json:"contained,omitempty" bson:"contained,omitempty"`                           // Contained, inline Resources
	Identifier            []Identifier          `json:"identifier,omitempty" bson:"identifier,omitempty"`                         // Business Identifier for claim
	TraceNumber           []Identifier          `json:"traceNumber,omitempty" bson:"trace_number,omitempty"`                      // Number for tracking
	Status                string                `json:"status" bson:"status"`                                                     // active | cancelled | draft | entered-in-error
	StatusReason          *string               `json:"statusReason,omitempty" bson:"status_reason,omitempty"`                    // Reason for status change
	Type                  *CodeableConcept      `json:"type" bson:"type"`                                                         // Category or discipline
	SubType               *CodeableConcept      `json:"subType,omitempty" bson:"sub_type,omitempty"`                              // More granular claim type
	Use                   string                `json:"use" bson:"use"`                                                           // claim | preauthorization | predetermination
	Subject               *Reference            `json:"subject" bson:"subject"`                                                   // The recipient(s) of the products and services
	BillablePeriod        *Period               `json:"billablePeriod,omitempty" bson:"billable_period,omitempty"`                // Relevant time frame for the claim
	Created               string                `json:"created" bson:"created"`                                                   // Resource creation date
	Enterer               *Reference            `json:"enterer,omitempty" bson:"enterer,omitempty"`                               // Author of the claim
	Insurer               *Reference            `json:"insurer,omitempty" bson:"insurer,omitempty"`                               // Target
	Provider              *Reference            `json:"provider,omitempty" bson:"provider,omitempty"`                             // Party responsible for the claim
	Priority              *CodeableConcept      `json:"priority,omitempty" bson:"priority,omitempty"`                             // Desired processing urgency
	FundsReserve          *CodeableConcept      `json:"fundsReserve,omitempty" bson:"funds_reserve,omitempty"`                    // For whom to reserve funds
	Related               []ClaimRelated        `json:"related,omitempty" bson:"related,omitempty"`                               // Prior or corollary claims
	Prescription          *Reference            `json:"prescription,omitempty" bson:"prescription,omitempty"`                     // Prescription authorizing services and products
	OriginalPrescription  *Reference            `json:"originalPrescription,omitempty" bson:"original_prescription,omitempty"`    // Original prescription if superseded by fulfiller
	Payee                 *ClaimPayee           `json:"payee,omitempty" bson:"payee,omitempty"`                                   // Recipient of benefits payable
	Referral              *Reference            `json:"referral,omitempty" bson:"referral,omitempty"`                             // Treatment referral
	Encounter             []Reference           `json:"encounter,omitempty" bson:"encounter,omitempty"`                           // Encounters associated with the listed treatments
	Facility              *Reference            `json:"facility,omitempty" bson:"facility,omitempty"`                             // Servicing facility
	DiagnosisRelatedGroup *CodeableConcept      `json:"diagnosisRelatedGroup,omitempty" bson:"diagnosis_related_group,omitempty"` // Package billing code
	Event                 []ClaimEvent          `json:"event,omitempty" bson:"event,omitempty"`                                   // Event information
	CareTeam              []ClaimCareTeam       `json:"careTeam,omitempty" bson:"care_team,omitempty"`                            // Members of the care team
	SupportingInfo        []ClaimSupportingInfo `json:"supportingInfo,omitempty" bson:"supporting_info,omitempty"`                // Supporting information
	Diagnosis             []ClaimDiagnosis      `json:"diagnosis,omitempty" bson:"diagnosis,omitempty"`                           // Pertinent diagnosis information
	Procedure             []ClaimProcedure      `json:"procedure,omitempty" bson:"procedure,omitempty"`                           // Clinical procedures performed
	Insurance             []ClaimInsurance      `json:"insurance,omitempty" bson:"insurance,omitempty"`                           // Patient insurance information
	Accident              *ClaimAccident        `json:"accident,omitempty" bson:"accident,omitempty"`                             // Details of the event
	PatientPaid           *Money                `json:"patientPaid,omitempty" bson:"patient_paid,omitempty"`                      // Paid by the patient
	Item                  []ClaimItem           `json:"item,omitempty" bson:"item,omitempty"`                                     // Product or service provided
	Total                 *Money                `json:"total,omitempty" bson:"total,omitempty"`                                   // Total claim cost
}

func (r *Claim) Validate() error {
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
	for i, item := range r.TraceNumber {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("TraceNumber[%d]: %w", i, err)
		}
	}
	var emptyString string
	if r.Status == emptyString {
		return fmt.Errorf("field 'Status' is required")
	}
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.SubType != nil {
		if err := r.SubType.Validate(); err != nil {
			return fmt.Errorf("SubType: %w", err)
		}
	}
	if r.Use == emptyString {
		return fmt.Errorf("field 'Use' is required")
	}
	if r.Subject == nil {
		return fmt.Errorf("field 'Subject' is required")
	}
	if r.Subject != nil {
		if err := r.Subject.Validate(); err != nil {
			return fmt.Errorf("Subject: %w", err)
		}
	}
	if r.BillablePeriod != nil {
		if err := r.BillablePeriod.Validate(); err != nil {
			return fmt.Errorf("BillablePeriod: %w", err)
		}
	}
	if r.Created == emptyString {
		return fmt.Errorf("field 'Created' is required")
	}
	if r.Enterer != nil {
		if err := r.Enterer.Validate(); err != nil {
			return fmt.Errorf("Enterer: %w", err)
		}
	}
	if r.Insurer != nil {
		if err := r.Insurer.Validate(); err != nil {
			return fmt.Errorf("Insurer: %w", err)
		}
	}
	if r.Provider != nil {
		if err := r.Provider.Validate(); err != nil {
			return fmt.Errorf("Provider: %w", err)
		}
	}
	if r.Priority != nil {
		if err := r.Priority.Validate(); err != nil {
			return fmt.Errorf("Priority: %w", err)
		}
	}
	if r.FundsReserve != nil {
		if err := r.FundsReserve.Validate(); err != nil {
			return fmt.Errorf("FundsReserve: %w", err)
		}
	}
	for i, item := range r.Related {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Related[%d]: %w", i, err)
		}
	}
	if r.Prescription != nil {
		if err := r.Prescription.Validate(); err != nil {
			return fmt.Errorf("Prescription: %w", err)
		}
	}
	if r.OriginalPrescription != nil {
		if err := r.OriginalPrescription.Validate(); err != nil {
			return fmt.Errorf("OriginalPrescription: %w", err)
		}
	}
	if r.Payee != nil {
		if err := r.Payee.Validate(); err != nil {
			return fmt.Errorf("Payee: %w", err)
		}
	}
	if r.Referral != nil {
		if err := r.Referral.Validate(); err != nil {
			return fmt.Errorf("Referral: %w", err)
		}
	}
	for i, item := range r.Encounter {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Encounter[%d]: %w", i, err)
		}
	}
	if r.Facility != nil {
		if err := r.Facility.Validate(); err != nil {
			return fmt.Errorf("Facility: %w", err)
		}
	}
	if r.DiagnosisRelatedGroup != nil {
		if err := r.DiagnosisRelatedGroup.Validate(); err != nil {
			return fmt.Errorf("DiagnosisRelatedGroup: %w", err)
		}
	}
	for i, item := range r.Event {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Event[%d]: %w", i, err)
		}
	}
	for i, item := range r.CareTeam {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("CareTeam[%d]: %w", i, err)
		}
	}
	for i, item := range r.SupportingInfo {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SupportingInfo[%d]: %w", i, err)
		}
	}
	for i, item := range r.Diagnosis {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Diagnosis[%d]: %w", i, err)
		}
	}
	for i, item := range r.Procedure {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Procedure[%d]: %w", i, err)
		}
	}
	for i, item := range r.Insurance {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Insurance[%d]: %w", i, err)
		}
	}
	if r.Accident != nil {
		if err := r.Accident.Validate(); err != nil {
			return fmt.Errorf("Accident: %w", err)
		}
	}
	if r.PatientPaid != nil {
		if err := r.PatientPaid.Validate(); err != nil {
			return fmt.Errorf("PatientPaid: %w", err)
		}
	}
	for i, item := range r.Item {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Item[%d]: %w", i, err)
		}
	}
	if r.Total != nil {
		if err := r.Total.Validate(); err != nil {
			return fmt.Errorf("Total: %w", err)
		}
	}
	return nil
}

type ClaimPayee struct {
	Id    *string          `json:"id,omitempty" bson:"id,omitempty"`       // Unique id for inter-element referencing
	Type  *CodeableConcept `json:"type" bson:"type"`                       // Category of recipient
	Party *Reference       `json:"party,omitempty" bson:"party,omitempty"` // Recipient reference
}

func (r *ClaimPayee) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Party != nil {
		if err := r.Party.Validate(); err != nil {
			return fmt.Errorf("Party: %w", err)
		}
	}
	return nil
}

type ClaimEvent struct {
	Id           *string          `json:"id,omitempty" bson:"id,omitempty"`   // Unique id for inter-element referencing
	Type         *CodeableConcept `json:"type" bson:"type"`                   // Specific event
	WhenDateTime *string          `json:"whenDateTime" bson:"when_date_time"` // Occurance date or period
	WhenPeriod   *Period          `json:"whenPeriod" bson:"when_period"`      // Occurance date or period
}

func (r *ClaimEvent) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.WhenDateTime == nil {
		return fmt.Errorf("field 'WhenDateTime' is required")
	}
	if r.WhenPeriod == nil {
		return fmt.Errorf("field 'WhenPeriod' is required")
	}
	if r.WhenPeriod != nil {
		if err := r.WhenPeriod.Validate(); err != nil {
			return fmt.Errorf("WhenPeriod: %w", err)
		}
	}
	return nil
}

type ClaimCareTeam struct {
	Id        *string          `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Sequence  int              `json:"sequence" bson:"sequence"`                       // Order of care team
	Provider  *Reference       `json:"provider" bson:"provider"`                       // Practitioner or organization
	Role      *CodeableConcept `json:"role,omitempty" bson:"role,omitempty"`           // Function within the team
	Specialty *CodeableConcept `json:"specialty,omitempty" bson:"specialty,omitempty"` // Practitioner or provider specialization
}

func (r *ClaimCareTeam) Validate() error {
	if r.Sequence == 0 {
		return fmt.Errorf("field 'Sequence' is required")
	}
	if r.Provider == nil {
		return fmt.Errorf("field 'Provider' is required")
	}
	if r.Provider != nil {
		if err := r.Provider.Validate(); err != nil {
			return fmt.Errorf("Provider: %w", err)
		}
	}
	if r.Role != nil {
		if err := r.Role.Validate(); err != nil {
			return fmt.Errorf("Role: %w", err)
		}
	}
	if r.Specialty != nil {
		if err := r.Specialty.Validate(); err != nil {
			return fmt.Errorf("Specialty: %w", err)
		}
	}
	return nil
}

type ClaimSupportingInfo struct {
	Id                         *string                `json:"id,omitempty" bson:"id,omitempty"`                                                    // Unique id for inter-element referencing
	Sequence                   int                    `json:"sequence" bson:"sequence"`                                                            // Information instance identifier
	Category                   *CodeableConcept       `json:"category" bson:"category"`                                                            // Classification of the supplied information
	SubCategory                *CodeableConcept       `json:"subCategory,omitempty" bson:"sub_category,omitempty"`                                 // Finer-grained classification of the supplied information
	Code                       *CodeableConcept       `json:"code,omitempty" bson:"code,omitempty"`                                                // Type of information
	TimingDateTime             *string                `json:"timingDateTime,omitempty" bson:"timing_date_time,omitempty"`                          // When it occurred
	TimingPeriod               *Period                `json:"timingPeriod,omitempty" bson:"timing_period,omitempty"`                               // When it occurred
	TimingTiming               *Timing                `json:"timingTiming,omitempty" bson:"timing_timing,omitempty"`                               // When it occurred
	ValueBase64Binary          *string                `json:"valueBase64Binary,omitempty" bson:"value_base64_binary,omitempty"`                    // Data to be provided
	ValueBoolean               *bool                  `json:"valueBoolean,omitempty" bson:"value_boolean,omitempty"`                               // Data to be provided
	ValueCanonical             *string                `json:"valueCanonical,omitempty" bson:"value_canonical,omitempty"`                           // Data to be provided
	ValueCode                  *string                `json:"valueCode,omitempty" bson:"value_code,omitempty"`                                     // Data to be provided
	ValueDate                  *string                `json:"valueDate,omitempty" bson:"value_date,omitempty"`                                     // Data to be provided
	ValueDateTime              *string                `json:"valueDateTime,omitempty" bson:"value_date_time,omitempty"`                            // Data to be provided
	ValueDecimal               *float64               `json:"valueDecimal,omitempty" bson:"value_decimal,omitempty"`                               // Data to be provided
	ValueId                    *string                `json:"valueId,omitempty" bson:"value_id,omitempty"`                                         // Data to be provided
	ValueInstant               *string                `json:"valueInstant,omitempty" bson:"value_instant,omitempty"`                               // Data to be provided
	ValueInteger               *int                   `json:"valueInteger,omitempty" bson:"value_integer,omitempty"`                               // Data to be provided
	ValueInteger64             *int64                 `json:"valueInteger64,omitempty" bson:"value_integer64,omitempty"`                           // Data to be provided
	ValueMarkdown              *string                `json:"valueMarkdown,omitempty" bson:"value_markdown,omitempty"`                             // Data to be provided
	ValueOid                   *string                `json:"valueOid,omitempty" bson:"value_oid,omitempty"`                                       // Data to be provided
	ValuePositiveInt           *int                   `json:"valuePositiveInt,omitempty" bson:"value_positive_int,omitempty"`                      // Data to be provided
	ValueString                *string                `json:"valueString,omitempty" bson:"value_string,omitempty"`                                 // Data to be provided
	ValueTime                  *string                `json:"valueTime,omitempty" bson:"value_time,omitempty"`                                     // Data to be provided
	ValueUnsignedInt           *int                   `json:"valueUnsignedInt,omitempty" bson:"value_unsigned_int,omitempty"`                      // Data to be provided
	ValueUri                   *string                `json:"valueUri,omitempty" bson:"value_uri,omitempty"`                                       // Data to be provided
	ValueUrl                   *string                `json:"valueUrl,omitempty" bson:"value_url,omitempty"`                                       // Data to be provided
	ValueUuid                  *uuid                  `json:"valueUuid,omitempty" bson:"value_uuid,omitempty"`                                     // Data to be provided
	ValueAddress               *Address               `json:"valueAddress,omitempty" bson:"value_address,omitempty"`                               // Data to be provided
	ValueAge                   *Age                   `json:"valueAge,omitempty" bson:"value_age,omitempty"`                                       // Data to be provided
	ValueAnnotation            *Annotation            `json:"valueAnnotation,omitempty" bson:"value_annotation,omitempty"`                         // Data to be provided
	ValueAttachment            *Attachment            `json:"valueAttachment,omitempty" bson:"value_attachment,omitempty"`                         // Data to be provided
	ValueCodeableConcept       *CodeableConcept       `json:"valueCodeableConcept,omitempty" bson:"value_codeable_concept,omitempty"`              // Data to be provided
	ValueCodeableReference     *CodeableReference     `json:"valueCodeableReference,omitempty" bson:"value_codeable_reference,omitempty"`          // Data to be provided
	ValueCoding                *Coding                `json:"valueCoding,omitempty" bson:"value_coding,omitempty"`                                 // Data to be provided
	ValueContactPoint          *ContactPoint          `json:"valueContactPoint,omitempty" bson:"value_contact_point,omitempty"`                    // Data to be provided
	ValueCount                 *Count                 `json:"valueCount,omitempty" bson:"value_count,omitempty"`                                   // Data to be provided
	ValueDistance              *Distance              `json:"valueDistance,omitempty" bson:"value_distance,omitempty"`                             // Data to be provided
	ValueDuration              *Duration              `json:"valueDuration,omitempty" bson:"value_duration,omitempty"`                             // Data to be provided
	ValueHumanName             *HumanName             `json:"valueHumanName,omitempty" bson:"value_human_name,omitempty"`                          // Data to be provided
	ValueIdentifier            *Identifier            `json:"valueIdentifier,omitempty" bson:"value_identifier,omitempty"`                         // Data to be provided
	ValueMoney                 *Money                 `json:"valueMoney,omitempty" bson:"value_money,omitempty"`                                   // Data to be provided
	ValuePeriod                *Period                `json:"valuePeriod,omitempty" bson:"value_period,omitempty"`                                 // Data to be provided
	ValueQuantity              *Quantity              `json:"valueQuantity,omitempty" bson:"value_quantity,omitempty"`                             // Data to be provided
	ValueRange                 *Range                 `json:"valueRange,omitempty" bson:"value_range,omitempty"`                                   // Data to be provided
	ValueRatio                 *Ratio                 `json:"valueRatio,omitempty" bson:"value_ratio,omitempty"`                                   // Data to be provided
	ValueRatioRange            *RatioRange            `json:"valueRatioRange,omitempty" bson:"value_ratio_range,omitempty"`                        // Data to be provided
	ValueReference             *Reference             `json:"valueReference,omitempty" bson:"value_reference,omitempty"`                           // Data to be provided
	ValueSampledData           *SampledData           `json:"valueSampledData,omitempty" bson:"value_sampled_data,omitempty"`                      // Data to be provided
	ValueSignature             *Signature             `json:"valueSignature,omitempty" bson:"value_signature,omitempty"`                           // Data to be provided
	ValueTiming                *Timing                `json:"valueTiming,omitempty" bson:"value_timing,omitempty"`                                 // Data to be provided
	ValueContactDetail         *ContactDetail         `json:"valueContactDetail,omitempty" bson:"value_contact_detail,omitempty"`                  // Data to be provided
	ValueDataRequirement       *DataRequirement       `json:"valueDataRequirement,omitempty" bson:"value_data_requirement,omitempty"`              // Data to be provided
	ValueExpression            *Expression            `json:"valueExpression,omitempty" bson:"value_expression,omitempty"`                         // Data to be provided
	ValueParameterDefinition   *ParameterDefinition   `json:"valueParameterDefinition,omitempty" bson:"value_parameter_definition,omitempty"`      // Data to be provided
	ValueRelatedArtifact       *RelatedArtifact       `json:"valueRelatedArtifact,omitempty" bson:"value_related_artifact,omitempty"`              // Data to be provided
	ValueTriggerDefinition     *TriggerDefinition     `json:"valueTriggerDefinition,omitempty" bson:"value_trigger_definition,omitempty"`          // Data to be provided
	ValueUsageContext          *UsageContext          `json:"valueUsageContext,omitempty" bson:"value_usage_context,omitempty"`                    // Data to be provided
	ValueAvailability          *Availability          `json:"valueAvailability,omitempty" bson:"value_availability,omitempty"`                     // Data to be provided
	ValueExtendedContactDetail *ExtendedContactDetail `json:"valueExtendedContactDetail,omitempty" bson:"value_extended_contact_detail,omitempty"` // Data to be provided
	ValueVirtualServiceDetail  *VirtualServiceDetail  `json:"valueVirtualServiceDetail,omitempty" bson:"value_virtual_service_detail,omitempty"`   // Data to be provided
	ValueDosage                *Dosage                `json:"valueDosage,omitempty" bson:"value_dosage,omitempty"`                                 // Data to be provided
	ValueMeta                  *Meta                  `json:"valueMeta,omitempty" bson:"value_meta,omitempty"`                                     // Data to be provided
	Reason                     *CodeableConcept       `json:"reason,omitempty" bson:"reason,omitempty"`                                            // Explanation for the information
}

func (r *ClaimSupportingInfo) Validate() error {
	if r.Sequence == 0 {
		return fmt.Errorf("field 'Sequence' is required")
	}
	if r.Category == nil {
		return fmt.Errorf("field 'Category' is required")
	}
	if r.Category != nil {
		if err := r.Category.Validate(); err != nil {
			return fmt.Errorf("Category: %w", err)
		}
	}
	if r.SubCategory != nil {
		if err := r.SubCategory.Validate(); err != nil {
			return fmt.Errorf("SubCategory: %w", err)
		}
	}
	if r.Code != nil {
		if err := r.Code.Validate(); err != nil {
			return fmt.Errorf("Code: %w", err)
		}
	}
	if r.TimingPeriod != nil {
		if err := r.TimingPeriod.Validate(); err != nil {
			return fmt.Errorf("TimingPeriod: %w", err)
		}
	}
	if r.TimingTiming != nil {
		if err := r.TimingTiming.Validate(); err != nil {
			return fmt.Errorf("TimingTiming: %w", err)
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
	if r.Reason != nil {
		if err := r.Reason.Validate(); err != nil {
			return fmt.Errorf("Reason: %w", err)
		}
	}
	return nil
}

type ClaimProcedure struct {
	Id                       *string           `json:"id,omitempty" bson:"id,omitempty"`                           // Unique id for inter-element referencing
	Sequence                 int               `json:"sequence" bson:"sequence"`                                   // Procedure instance identifier
	Type                     []CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`                       // Category of Procedure
	Date                     *string           `json:"date,omitempty" bson:"date,omitempty"`                       // When the procedure was performed
	ProcedureCodeableConcept *CodeableConcept  `json:"procedureCodeableConcept" bson:"procedure_codeable_concept"` // Specific clinical procedure
	ProcedureReference       *Reference        `json:"procedureReference" bson:"procedure_reference"`              // Specific clinical procedure
	Udi                      []Reference       `json:"udi,omitempty" bson:"udi,omitempty"`                         // Unique device identifier
}

func (r *ClaimProcedure) Validate() error {
	if r.Sequence == 0 {
		return fmt.Errorf("field 'Sequence' is required")
	}
	for i, item := range r.Type {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Type[%d]: %w", i, err)
		}
	}
	if r.ProcedureCodeableConcept == nil {
		return fmt.Errorf("field 'ProcedureCodeableConcept' is required")
	}
	if r.ProcedureCodeableConcept != nil {
		if err := r.ProcedureCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("ProcedureCodeableConcept: %w", err)
		}
	}
	if r.ProcedureReference == nil {
		return fmt.Errorf("field 'ProcedureReference' is required")
	}
	if r.ProcedureReference != nil {
		if err := r.ProcedureReference.Validate(); err != nil {
			return fmt.Errorf("ProcedureReference: %w", err)
		}
	}
	for i, item := range r.Udi {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Udi[%d]: %w", i, err)
		}
	}
	return nil
}

type ClaimAccident struct {
	Id                *string          `json:"id,omitempty" bson:"id,omitempty"`                                // Unique id for inter-element referencing
	Date              string           `json:"date" bson:"date"`                                                // When the incident occurred
	Type              *CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`                            // The nature of the accident
	LocationAddress   *Address         `json:"locationAddress,omitempty" bson:"location_address,omitempty"`     // Where the event occurred
	LocationReference *Reference       `json:"locationReference,omitempty" bson:"location_reference,omitempty"` // Where the event occurred
}

func (r *ClaimAccident) Validate() error {
	var emptyString string
	if r.Date == emptyString {
		return fmt.Errorf("field 'Date' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.LocationAddress != nil {
		if err := r.LocationAddress.Validate(); err != nil {
			return fmt.Errorf("LocationAddress: %w", err)
		}
	}
	if r.LocationReference != nil {
		if err := r.LocationReference.Validate(); err != nil {
			return fmt.Errorf("LocationReference: %w", err)
		}
	}
	return nil
}

type ClaimItem struct {
	Id                      *string             `json:"id,omitempty" bson:"id,omitempty"`                                             // Unique id for inter-element referencing
	Sequence                int                 `json:"sequence" bson:"sequence"`                                                     // Item instance identifier
	TraceNumber             []Identifier        `json:"traceNumber,omitempty" bson:"trace_number,omitempty"`                          // Number for tracking
	Subject                 *Reference          `json:"subject,omitempty" bson:"subject,omitempty"`                                   // The recipient of the products and services
	CareTeamSequence        []int               `json:"careTeamSequence,omitempty" bson:"care_team_sequence,omitempty"`               // Applicable careTeam members
	DiagnosisSequence       []int               `json:"diagnosisSequence,omitempty" bson:"diagnosis_sequence,omitempty"`              // Applicable diagnoses
	ProcedureSequence       []int               `json:"procedureSequence,omitempty" bson:"procedure_sequence,omitempty"`              // Applicable procedures
	InformationSequence     []int               `json:"informationSequence,omitempty" bson:"information_sequence,omitempty"`          // Applicable exception and supporting information
	Revenue                 *CodeableConcept    `json:"revenue,omitempty" bson:"revenue,omitempty"`                                   // Revenue or cost center code
	Category                *CodeableConcept    `json:"category,omitempty" bson:"category,omitempty"`                                 // Benefit classification
	ProductOrService        *CodeableConcept    `json:"productOrService,omitempty" bson:"product_or_service,omitempty"`               // Billing, service, product, or drug code
	ProductOrServiceEnd     *CodeableConcept    `json:"productOrServiceEnd,omitempty" bson:"product_or_service_end,omitempty"`        // End of a range of codes
	Request                 []Reference         `json:"request,omitempty" bson:"request,omitempty"`                                   // Request or Referral for Service
	Modifier                []CodeableConcept   `json:"modifier,omitempty" bson:"modifier,omitempty"`                                 // Product or service billing modifiers
	ProgramCode             []CodeableConcept   `json:"programCode,omitempty" bson:"program_code,omitempty"`                          // Program the product or service is provided under
	ServicedDate            *string             `json:"servicedDate,omitempty" bson:"serviced_date,omitempty"`                        // Date or dates of service or product delivery
	ServicedPeriod          *Period             `json:"servicedPeriod,omitempty" bson:"serviced_period,omitempty"`                    // Date or dates of service or product delivery
	LocationCodeableConcept *CodeableConcept    `json:"locationCodeableConcept,omitempty" bson:"location_codeable_concept,omitempty"` // Place of service or where product was supplied
	LocationAddress         *Address            `json:"locationAddress,omitempty" bson:"location_address,omitempty"`                  // Place of service or where product was supplied
	LocationReference       *Reference          `json:"locationReference,omitempty" bson:"location_reference,omitempty"`              // Place of service or where product was supplied
	PatientPaid             *Money              `json:"patientPaid,omitempty" bson:"patient_paid,omitempty"`                          // Paid by the patient
	Quantity                *Quantity           `json:"quantity,omitempty" bson:"quantity,omitempty"`                                 // Count of products or services
	UnitPrice               *Money              `json:"unitPrice,omitempty" bson:"unit_price,omitempty"`                              // Fee, charge or cost per item
	Factor                  *float64            `json:"factor,omitempty" bson:"factor,omitempty"`                                     // Price scaling factor
	Tax                     *Money              `json:"tax,omitempty" bson:"tax,omitempty"`                                           // Total tax
	Net                     *Money              `json:"net,omitempty" bson:"net,omitempty"`                                           // Total item cost
	Udi                     []Reference         `json:"udi,omitempty" bson:"udi,omitempty"`                                           // Unique device identifier
	BodySite                []ClaimItemBodySite `json:"bodySite,omitempty" bson:"body_site,omitempty"`                                // Anatomical location
	Encounter               []Reference         `json:"encounter,omitempty" bson:"encounter,omitempty"`                               // Encounters associated with the listed treatments
	Detail                  []ClaimItemDetail   `json:"detail,omitempty" bson:"detail,omitempty"`                                     // Product or service provided
}

func (r *ClaimItem) Validate() error {
	if r.Sequence == 0 {
		return fmt.Errorf("field 'Sequence' is required")
	}
	for i, item := range r.TraceNumber {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("TraceNumber[%d]: %w", i, err)
		}
	}
	if r.Subject != nil {
		if err := r.Subject.Validate(); err != nil {
			return fmt.Errorf("Subject: %w", err)
		}
	}
	if r.Revenue != nil {
		if err := r.Revenue.Validate(); err != nil {
			return fmt.Errorf("Revenue: %w", err)
		}
	}
	if r.Category != nil {
		if err := r.Category.Validate(); err != nil {
			return fmt.Errorf("Category: %w", err)
		}
	}
	if r.ProductOrService != nil {
		if err := r.ProductOrService.Validate(); err != nil {
			return fmt.Errorf("ProductOrService: %w", err)
		}
	}
	if r.ProductOrServiceEnd != nil {
		if err := r.ProductOrServiceEnd.Validate(); err != nil {
			return fmt.Errorf("ProductOrServiceEnd: %w", err)
		}
	}
	for i, item := range r.Request {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Request[%d]: %w", i, err)
		}
	}
	for i, item := range r.Modifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Modifier[%d]: %w", i, err)
		}
	}
	for i, item := range r.ProgramCode {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ProgramCode[%d]: %w", i, err)
		}
	}
	if r.ServicedPeriod != nil {
		if err := r.ServicedPeriod.Validate(); err != nil {
			return fmt.Errorf("ServicedPeriod: %w", err)
		}
	}
	if r.LocationCodeableConcept != nil {
		if err := r.LocationCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("LocationCodeableConcept: %w", err)
		}
	}
	if r.LocationAddress != nil {
		if err := r.LocationAddress.Validate(); err != nil {
			return fmt.Errorf("LocationAddress: %w", err)
		}
	}
	if r.LocationReference != nil {
		if err := r.LocationReference.Validate(); err != nil {
			return fmt.Errorf("LocationReference: %w", err)
		}
	}
	if r.PatientPaid != nil {
		if err := r.PatientPaid.Validate(); err != nil {
			return fmt.Errorf("PatientPaid: %w", err)
		}
	}
	if r.Quantity != nil {
		if err := r.Quantity.Validate(); err != nil {
			return fmt.Errorf("Quantity: %w", err)
		}
	}
	if r.UnitPrice != nil {
		if err := r.UnitPrice.Validate(); err != nil {
			return fmt.Errorf("UnitPrice: %w", err)
		}
	}
	if r.Tax != nil {
		if err := r.Tax.Validate(); err != nil {
			return fmt.Errorf("Tax: %w", err)
		}
	}
	if r.Net != nil {
		if err := r.Net.Validate(); err != nil {
			return fmt.Errorf("Net: %w", err)
		}
	}
	for i, item := range r.Udi {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Udi[%d]: %w", i, err)
		}
	}
	for i, item := range r.BodySite {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("BodySite[%d]: %w", i, err)
		}
	}
	for i, item := range r.Encounter {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Encounter[%d]: %w", i, err)
		}
	}
	for i, item := range r.Detail {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Detail[%d]: %w", i, err)
		}
	}
	return nil
}

type ClaimItemBodySite struct {
	Id      *string             `json:"id,omitempty" bson:"id,omitempty"`            // Unique id for inter-element referencing
	Site    []CodeableReference `json:"site" bson:"site"`                            // Location
	SubSite []CodeableConcept   `json:"subSite,omitempty" bson:"sub_site,omitempty"` // Sub-location
}

func (r *ClaimItemBodySite) Validate() error {
	if len(r.Site) < 1 {
		return fmt.Errorf("field 'Site' must have at least 1 elements")
	}
	for i, item := range r.Site {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Site[%d]: %w", i, err)
		}
	}
	for i, item := range r.SubSite {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SubSite[%d]: %w", i, err)
		}
	}
	return nil
}

type ClaimRelated struct {
	Id           *string          `json:"id,omitempty" bson:"id,omitempty"`                     // Unique id for inter-element referencing
	Claim        *Reference       `json:"claim,omitempty" bson:"claim,omitempty"`               // Reference to the related claim
	Relationship *CodeableConcept `json:"relationship,omitempty" bson:"relationship,omitempty"` // How the reference claim is related
	Reference    *Identifier      `json:"reference,omitempty" bson:"reference,omitempty"`       // File or case reference
}

func (r *ClaimRelated) Validate() error {
	if r.Claim != nil {
		if err := r.Claim.Validate(); err != nil {
			return fmt.Errorf("Claim: %w", err)
		}
	}
	if r.Relationship != nil {
		if err := r.Relationship.Validate(); err != nil {
			return fmt.Errorf("Relationship: %w", err)
		}
	}
	if r.Reference != nil {
		if err := r.Reference.Validate(); err != nil {
			return fmt.Errorf("Reference: %w", err)
		}
	}
	return nil
}

type ClaimDiagnosis struct {
	Id                       *string           `json:"id,omitempty" bson:"id,omitempty"`                           // Unique id for inter-element referencing
	Sequence                 int               `json:"sequence" bson:"sequence"`                                   // Diagnosis instance identifier
	DiagnosisCodeableConcept *CodeableConcept  `json:"diagnosisCodeableConcept" bson:"diagnosis_codeable_concept"` // Nature of illness or problem
	DiagnosisReference       *Reference        `json:"diagnosisReference" bson:"diagnosis_reference"`              // Nature of illness or problem
	Type                     []CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`                       // Timing or nature of the diagnosis
	OnAdmission              *CodeableConcept  `json:"onAdmission,omitempty" bson:"on_admission,omitempty"`        // Present on admission
}

func (r *ClaimDiagnosis) Validate() error {
	if r.Sequence == 0 {
		return fmt.Errorf("field 'Sequence' is required")
	}
	if r.DiagnosisCodeableConcept == nil {
		return fmt.Errorf("field 'DiagnosisCodeableConcept' is required")
	}
	if r.DiagnosisCodeableConcept != nil {
		if err := r.DiagnosisCodeableConcept.Validate(); err != nil {
			return fmt.Errorf("DiagnosisCodeableConcept: %w", err)
		}
	}
	if r.DiagnosisReference == nil {
		return fmt.Errorf("field 'DiagnosisReference' is required")
	}
	if r.DiagnosisReference != nil {
		if err := r.DiagnosisReference.Validate(); err != nil {
			return fmt.Errorf("DiagnosisReference: %w", err)
		}
	}
	for i, item := range r.Type {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Type[%d]: %w", i, err)
		}
	}
	if r.OnAdmission != nil {
		if err := r.OnAdmission.Validate(); err != nil {
			return fmt.Errorf("OnAdmission: %w", err)
		}
	}
	return nil
}

type ClaimInsurance struct {
	Id                  *string     `json:"id,omitempty" bson:"id,omitempty"`                                    // Unique id for inter-element referencing
	Sequence            int         `json:"sequence" bson:"sequence"`                                            // Insurance instance identifier
	Focal               bool        `json:"focal" bson:"focal"`                                                  // Coverage to be used for adjudication
	Identifier          *Identifier `json:"identifier,omitempty" bson:"identifier,omitempty"`                    // Pre-assigned Claim number
	Coverage            *Reference  `json:"coverage" bson:"coverage"`                                            // Insurance information
	BusinessArrangement *string     `json:"businessArrangement,omitempty" bson:"business_arrangement,omitempty"` // Additional provider contract number
	PreAuthRef          []string    `json:"preAuthRef,omitempty" bson:"pre_auth_ref,omitempty"`                  // Prior authorization reference number
	ClaimResponse       *Reference  `json:"claimResponse,omitempty" bson:"claim_response,omitempty"`             // Adjudication results
}

func (r *ClaimInsurance) Validate() error {
	if r.Sequence == 0 {
		return fmt.Errorf("field 'Sequence' is required")
	}
	if r.Identifier != nil {
		if err := r.Identifier.Validate(); err != nil {
			return fmt.Errorf("Identifier: %w", err)
		}
	}
	if r.Coverage == nil {
		return fmt.Errorf("field 'Coverage' is required")
	}
	if r.Coverage != nil {
		if err := r.Coverage.Validate(); err != nil {
			return fmt.Errorf("Coverage: %w", err)
		}
	}
	if r.ClaimResponse != nil {
		if err := r.ClaimResponse.Validate(); err != nil {
			return fmt.Errorf("ClaimResponse: %w", err)
		}
	}
	return nil
}

type ClaimItemDetail struct {
	Id                  *string                    `json:"id,omitempty" bson:"id,omitempty"`                                      // Unique id for inter-element referencing
	Sequence            int                        `json:"sequence" bson:"sequence"`                                              // Item instance identifier
	TraceNumber         []Identifier               `json:"traceNumber,omitempty" bson:"trace_number,omitempty"`                   // Number for tracking
	Revenue             *CodeableConcept           `json:"revenue,omitempty" bson:"revenue,omitempty"`                            // Revenue or cost center code
	Category            *CodeableConcept           `json:"category,omitempty" bson:"category,omitempty"`                          // Benefit classification
	ProductOrService    *CodeableConcept           `json:"productOrService,omitempty" bson:"product_or_service,omitempty"`        // Billing, service, product, or drug code
	ProductOrServiceEnd *CodeableConcept           `json:"productOrServiceEnd,omitempty" bson:"product_or_service_end,omitempty"` // End of a range of codes
	Modifier            []CodeableConcept          `json:"modifier,omitempty" bson:"modifier,omitempty"`                          // Service/Product billing modifiers
	ProgramCode         []CodeableConcept          `json:"programCode,omitempty" bson:"program_code,omitempty"`                   // Program the product or service is provided under
	PatientPaid         *Money                     `json:"patientPaid,omitempty" bson:"patient_paid,omitempty"`                   // Paid by the patient
	Quantity            *Quantity                  `json:"quantity,omitempty" bson:"quantity,omitempty"`                          // Count of products or services
	UnitPrice           *Money                     `json:"unitPrice,omitempty" bson:"unit_price,omitempty"`                       // Fee, charge or cost per item
	Factor              *float64                   `json:"factor,omitempty" bson:"factor,omitempty"`                              // Price scaling factor
	Tax                 *Money                     `json:"tax,omitempty" bson:"tax,omitempty"`                                    // Total tax
	Net                 *Money                     `json:"net,omitempty" bson:"net,omitempty"`                                    // Total item cost
	Udi                 []Reference                `json:"udi,omitempty" bson:"udi,omitempty"`                                    // Unique device identifier
	SubDetail           []ClaimItemDetailSubDetail `json:"subDetail,omitempty" bson:"sub_detail,omitempty"`                       // Product or service provided
}

func (r *ClaimItemDetail) Validate() error {
	if r.Sequence == 0 {
		return fmt.Errorf("field 'Sequence' is required")
	}
	for i, item := range r.TraceNumber {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("TraceNumber[%d]: %w", i, err)
		}
	}
	if r.Revenue != nil {
		if err := r.Revenue.Validate(); err != nil {
			return fmt.Errorf("Revenue: %w", err)
		}
	}
	if r.Category != nil {
		if err := r.Category.Validate(); err != nil {
			return fmt.Errorf("Category: %w", err)
		}
	}
	if r.ProductOrService != nil {
		if err := r.ProductOrService.Validate(); err != nil {
			return fmt.Errorf("ProductOrService: %w", err)
		}
	}
	if r.ProductOrServiceEnd != nil {
		if err := r.ProductOrServiceEnd.Validate(); err != nil {
			return fmt.Errorf("ProductOrServiceEnd: %w", err)
		}
	}
	for i, item := range r.Modifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Modifier[%d]: %w", i, err)
		}
	}
	for i, item := range r.ProgramCode {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ProgramCode[%d]: %w", i, err)
		}
	}
	if r.PatientPaid != nil {
		if err := r.PatientPaid.Validate(); err != nil {
			return fmt.Errorf("PatientPaid: %w", err)
		}
	}
	if r.Quantity != nil {
		if err := r.Quantity.Validate(); err != nil {
			return fmt.Errorf("Quantity: %w", err)
		}
	}
	if r.UnitPrice != nil {
		if err := r.UnitPrice.Validate(); err != nil {
			return fmt.Errorf("UnitPrice: %w", err)
		}
	}
	if r.Tax != nil {
		if err := r.Tax.Validate(); err != nil {
			return fmt.Errorf("Tax: %w", err)
		}
	}
	if r.Net != nil {
		if err := r.Net.Validate(); err != nil {
			return fmt.Errorf("Net: %w", err)
		}
	}
	for i, item := range r.Udi {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Udi[%d]: %w", i, err)
		}
	}
	for i, item := range r.SubDetail {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SubDetail[%d]: %w", i, err)
		}
	}
	return nil
}

type ClaimItemDetailSubDetail struct {
	Id                  *string           `json:"id,omitempty" bson:"id,omitempty"`                                      // Unique id for inter-element referencing
	Sequence            int               `json:"sequence" bson:"sequence"`                                              // Item instance identifier
	TraceNumber         []Identifier      `json:"traceNumber,omitempty" bson:"trace_number,omitempty"`                   // Number for tracking
	Revenue             *CodeableConcept  `json:"revenue,omitempty" bson:"revenue,omitempty"`                            // Revenue or cost center code
	Category            *CodeableConcept  `json:"category,omitempty" bson:"category,omitempty"`                          // Benefit classification
	ProductOrService    *CodeableConcept  `json:"productOrService,omitempty" bson:"product_or_service,omitempty"`        // Billing, service, product, or drug code
	ProductOrServiceEnd *CodeableConcept  `json:"productOrServiceEnd,omitempty" bson:"product_or_service_end,omitempty"` // End of a range of codes
	Modifier            []CodeableConcept `json:"modifier,omitempty" bson:"modifier,omitempty"`                          // Service/Product billing modifiers
	ProgramCode         []CodeableConcept `json:"programCode,omitempty" bson:"program_code,omitempty"`                   // Program the product or service is provided under
	PatientPaid         *Money            `json:"patientPaid,omitempty" bson:"patient_paid,omitempty"`                   // Paid by the patient
	Quantity            *Quantity         `json:"quantity,omitempty" bson:"quantity,omitempty"`                          // Count of products or services
	UnitPrice           *Money            `json:"unitPrice,omitempty" bson:"unit_price,omitempty"`                       // Fee, charge or cost per item
	Factor              *float64          `json:"factor,omitempty" bson:"factor,omitempty"`                              // Price scaling factor
	Tax                 *Money            `json:"tax,omitempty" bson:"tax,omitempty"`                                    // Total tax
	Net                 *Money            `json:"net,omitempty" bson:"net,omitempty"`                                    // Total item cost
	Udi                 []Reference       `json:"udi,omitempty" bson:"udi,omitempty"`                                    // Unique device identifier
}

func (r *ClaimItemDetailSubDetail) Validate() error {
	if r.Sequence == 0 {
		return fmt.Errorf("field 'Sequence' is required")
	}
	for i, item := range r.TraceNumber {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("TraceNumber[%d]: %w", i, err)
		}
	}
	if r.Revenue != nil {
		if err := r.Revenue.Validate(); err != nil {
			return fmt.Errorf("Revenue: %w", err)
		}
	}
	if r.Category != nil {
		if err := r.Category.Validate(); err != nil {
			return fmt.Errorf("Category: %w", err)
		}
	}
	if r.ProductOrService != nil {
		if err := r.ProductOrService.Validate(); err != nil {
			return fmt.Errorf("ProductOrService: %w", err)
		}
	}
	if r.ProductOrServiceEnd != nil {
		if err := r.ProductOrServiceEnd.Validate(); err != nil {
			return fmt.Errorf("ProductOrServiceEnd: %w", err)
		}
	}
	for i, item := range r.Modifier {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Modifier[%d]: %w", i, err)
		}
	}
	for i, item := range r.ProgramCode {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ProgramCode[%d]: %w", i, err)
		}
	}
	if r.PatientPaid != nil {
		if err := r.PatientPaid.Validate(); err != nil {
			return fmt.Errorf("PatientPaid: %w", err)
		}
	}
	if r.Quantity != nil {
		if err := r.Quantity.Validate(); err != nil {
			return fmt.Errorf("Quantity: %w", err)
		}
	}
	if r.UnitPrice != nil {
		if err := r.UnitPrice.Validate(); err != nil {
			return fmt.Errorf("UnitPrice: %w", err)
		}
	}
	if r.Tax != nil {
		if err := r.Tax.Validate(); err != nil {
			return fmt.Errorf("Tax: %w", err)
		}
	}
	if r.Net != nil {
		if err := r.Net.Validate(); err != nil {
			return fmt.Errorf("Net: %w", err)
		}
	}
	for i, item := range r.Udi {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Udi[%d]: %w", i, err)
		}
	}
	return nil
}
