package models

import (
	"encoding/json"
	"fmt"
)

// This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefit struct {
	ResourceType          string                                 `json:"resourceType" bson:"resource_type"`                                        // Type of resource
	Id                    *string                                `json:"id,omitempty" bson:"id,omitempty"`                                         // Logical id of this artifact
	Meta                  *Meta                                  `json:"meta,omitempty" bson:"meta,omitempty"`                                     // Metadata about the resource
	ImplicitRules         *string                                `json:"implicitRules,omitempty" bson:"implicit_rules,omitempty"`                  // A set of rules under which this content was created
	Language              *string                                `json:"language,omitempty" bson:"language,omitempty"`                             // Language of the resource content
	Text                  *Narrative                             `json:"text,omitempty" bson:"text,omitempty"`                                     // Text summary of the resource, for human interpretation
	Contained             []json.RawMessage                      `json:"contained,omitempty" bson:"contained,omitempty"`                           // Contained, inline Resources
	Identifier            []Identifier                           `json:"identifier,omitempty" bson:"identifier,omitempty"`                         // Business Identifier for the resource
	TraceNumber           []Identifier                           `json:"traceNumber,omitempty" bson:"trace_number,omitempty"`                      // Number for tracking
	Status                string                                 `json:"status" bson:"status"`                                                     // active | cancelled | draft | entered-in-error
	StatusReason          *string                                `json:"statusReason,omitempty" bson:"status_reason,omitempty"`                    // Reason for status change
	Type                  *CodeableConcept                       `json:"type" bson:"type"`                                                         // Category or discipline
	SubType               *CodeableConcept                       `json:"subType,omitempty" bson:"sub_type,omitempty"`                              // More granular claim type
	Use                   string                                 `json:"use" bson:"use"`                                                           // claim | preauthorization | predetermination
	Subject               *Reference                             `json:"subject" bson:"subject"`                                                   // The recipient(s) of the products and services
	BillablePeriod        *Period                                `json:"billablePeriod,omitempty" bson:"billable_period,omitempty"`                // Relevant time frame for the claim
	Created               string                                 `json:"created" bson:"created"`                                                   // Response creation date
	Enterer               *Reference                             `json:"enterer,omitempty" bson:"enterer,omitempty"`                               // Author of the claim
	Insurer               *Reference                             `json:"insurer,omitempty" bson:"insurer,omitempty"`                               // Party responsible for reimbursement
	Provider              *Reference                             `json:"provider,omitempty" bson:"provider,omitempty"`                             // Party responsible for the claim
	Priority              *CodeableConcept                       `json:"priority,omitempty" bson:"priority,omitempty"`                             // Desired processing urgency
	FundsReserveRequested *CodeableConcept                       `json:"fundsReserveRequested,omitempty" bson:"funds_reserve_requested,omitempty"` // For whom to reserve funds
	FundsReserve          *CodeableConcept                       `json:"fundsReserve,omitempty" bson:"funds_reserve,omitempty"`                    // Funds reserved status
	Related               []ExplanationOfBenefitRelated          `json:"related,omitempty" bson:"related,omitempty"`                               // Prior or corollary claims
	Prescription          *Reference                             `json:"prescription,omitempty" bson:"prescription,omitempty"`                     // Prescription authorizing services or products
	OriginalPrescription  *Reference                             `json:"originalPrescription,omitempty" bson:"original_prescription,omitempty"`    // Original prescription if superceded by fulfiller
	Event                 []ExplanationOfBenefitEvent            `json:"event,omitempty" bson:"event,omitempty"`                                   // Event information
	Payee                 *ExplanationOfBenefitPayee             `json:"payee,omitempty" bson:"payee,omitempty"`                                   // Recipient of benefits payable
	Referral              *Reference                             `json:"referral,omitempty" bson:"referral,omitempty"`                             // Treatment Referral
	Encounter             []Reference                            `json:"encounter,omitempty" bson:"encounter,omitempty"`                           // Encounters associated with the listed treatments
	Facility              *Reference                             `json:"facility,omitempty" bson:"facility,omitempty"`                             // Servicing Facility
	Claim                 *Reference                             `json:"claim,omitempty" bson:"claim,omitempty"`                                   // Claim reference
	ClaimResponse         *Reference                             `json:"claimResponse,omitempty" bson:"claim_response,omitempty"`                  // Claim response reference
	Outcome               string                                 `json:"outcome" bson:"outcome"`                                                   // queued | complete | error | partial
	Decision              *CodeableConcept                       `json:"decision,omitempty" bson:"decision,omitempty"`                             // Result of the adjudication
	Disposition           *string                                `json:"disposition,omitempty" bson:"disposition,omitempty"`                       // Disposition Message
	PreAuthRef            []string                               `json:"preAuthRef,omitempty" bson:"pre_auth_ref,omitempty"`                       // Preauthorization reference
	PreAuthRefPeriod      []Period                               `json:"preAuthRefPeriod,omitempty" bson:"pre_auth_ref_period,omitempty"`          // Preauthorization in-effect period
	DiagnosisRelatedGroup *CodeableConcept                       `json:"diagnosisRelatedGroup,omitempty" bson:"diagnosis_related_group,omitempty"` // Package billing code
	CareTeam              []ExplanationOfBenefitCareTeam         `json:"careTeam,omitempty" bson:"care_team,omitempty"`                            // Care Team members
	SupportingInfo        []ExplanationOfBenefitSupportingInfo   `json:"supportingInfo,omitempty" bson:"supporting_info,omitempty"`                // Supporting information
	Diagnosis             []ExplanationOfBenefitDiagnosis        `json:"diagnosis,omitempty" bson:"diagnosis,omitempty"`                           // Pertinent diagnosis information
	Procedure             []ExplanationOfBenefitProcedure        `json:"procedure,omitempty" bson:"procedure,omitempty"`                           // Clinical procedures performed
	Precedence            *int                                   `json:"precedence,omitempty" bson:"precedence,omitempty"`                         // Precedence (primary, secondary, etc.)
	Insurance             []ExplanationOfBenefitInsurance        `json:"insurance,omitempty" bson:"insurance,omitempty"`                           // Patient insurance information
	Accident              *ExplanationOfBenefitAccident          `json:"accident,omitempty" bson:"accident,omitempty"`                             // Details of the event
	PatientPaid           *Money                                 `json:"patientPaid,omitempty" bson:"patient_paid,omitempty"`                      // Paid by the patient
	Item                  []ExplanationOfBenefitItem             `json:"item,omitempty" bson:"item,omitempty"`                                     // Product or service provided
	AddItem               []ExplanationOfBenefitAddItem          `json:"addItem,omitempty" bson:"add_item,omitempty"`                              // Insurer added line items
	Adjudication          []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty" bson:"adjudication,omitempty"`                     // Header-level adjudication
	Total                 []ExplanationOfBenefitTotal            `json:"total,omitempty" bson:"total,omitempty"`                                   // Adjudication totals
	Payment               *ExplanationOfBenefitPayment           `json:"payment,omitempty" bson:"payment,omitempty"`                               // Payment Details
	FormCode              *CodeableConcept                       `json:"formCode,omitempty" bson:"form_code,omitempty"`                            // Printed form identifier
	Form                  *Attachment                            `json:"form,omitempty" bson:"form,omitempty"`                                     // Printed reference or actual form
	ProcessNote           []ExplanationOfBenefitProcessNote      `json:"processNote,omitempty" bson:"process_note,omitempty"`                      // Note concerning adjudication
	BenefitPeriod         *Period                                `json:"benefitPeriod,omitempty" bson:"benefit_period,omitempty"`                  // When the benefits are applicable
	BenefitBalance        []ExplanationOfBenefitBenefitBalance   `json:"benefitBalance,omitempty" bson:"benefit_balance,omitempty"`                // Balance by Benefit Category
}

func (r *ExplanationOfBenefit) Validate() error {
	if r.ResourceType != "ExplanationOfBenefit" {
		return fmt.Errorf("invalid resourceType: expected 'ExplanationOfBenefit', got '%s'", r.ResourceType)
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
	if r.FundsReserveRequested != nil {
		if err := r.FundsReserveRequested.Validate(); err != nil {
			return fmt.Errorf("FundsReserveRequested: %w", err)
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
	for i, item := range r.Event {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Event[%d]: %w", i, err)
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
	if r.Claim != nil {
		if err := r.Claim.Validate(); err != nil {
			return fmt.Errorf("Claim: %w", err)
		}
	}
	if r.ClaimResponse != nil {
		if err := r.ClaimResponse.Validate(); err != nil {
			return fmt.Errorf("ClaimResponse: %w", err)
		}
	}
	if r.Outcome == emptyString {
		return fmt.Errorf("field 'Outcome' is required")
	}
	if r.Decision != nil {
		if err := r.Decision.Validate(); err != nil {
			return fmt.Errorf("Decision: %w", err)
		}
	}
	for i, item := range r.PreAuthRefPeriod {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("PreAuthRefPeriod[%d]: %w", i, err)
		}
	}
	if r.DiagnosisRelatedGroup != nil {
		if err := r.DiagnosisRelatedGroup.Validate(); err != nil {
			return fmt.Errorf("DiagnosisRelatedGroup: %w", err)
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
	for i, item := range r.AddItem {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("AddItem[%d]: %w", i, err)
		}
	}
	for i, item := range r.Adjudication {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Adjudication[%d]: %w", i, err)
		}
	}
	for i, item := range r.Total {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Total[%d]: %w", i, err)
		}
	}
	if r.Payment != nil {
		if err := r.Payment.Validate(); err != nil {
			return fmt.Errorf("Payment: %w", err)
		}
	}
	if r.FormCode != nil {
		if err := r.FormCode.Validate(); err != nil {
			return fmt.Errorf("FormCode: %w", err)
		}
	}
	if r.Form != nil {
		if err := r.Form.Validate(); err != nil {
			return fmt.Errorf("Form: %w", err)
		}
	}
	for i, item := range r.ProcessNote {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("ProcessNote[%d]: %w", i, err)
		}
	}
	if r.BenefitPeriod != nil {
		if err := r.BenefitPeriod.Validate(); err != nil {
			return fmt.Errorf("BenefitPeriod: %w", err)
		}
	}
	for i, item := range r.BenefitBalance {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("BenefitBalance[%d]: %w", i, err)
		}
	}
	return nil
}

type ExplanationOfBenefitEvent struct {
	Id           *string          `json:"id,omitempty" bson:"id,omitempty"`   // Unique id for inter-element referencing
	Type         *CodeableConcept `json:"type" bson:"type"`                   // Specific event
	WhenDateTime *string          `json:"whenDateTime" bson:"when_date_time"` // Occurance date or period
	WhenPeriod   *Period          `json:"whenPeriod" bson:"when_period"`      // Occurance date or period
}

func (r *ExplanationOfBenefitEvent) Validate() error {
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

type ExplanationOfBenefitSupportingInfo struct {
	Id                         *string                `json:"id,omitempty" bson:"id,omitempty"`                                                    // Unique id for inter-element referencing
	Sequence                   int                    `json:"sequence" bson:"sequence"`                                                            // Information instance identifier
	Category                   *CodeableConcept       `json:"category" bson:"category"`                                                            // Classification of the supplied information
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
	Reason                     *Coding                `json:"reason,omitempty" bson:"reason,omitempty"`                                            // Explanation for the information
}

func (r *ExplanationOfBenefitSupportingInfo) Validate() error {
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

type ExplanationOfBenefitAccident struct {
	Id                *string          `json:"id,omitempty" bson:"id,omitempty"`                                // Unique id for inter-element referencing
	Date              *string          `json:"date,omitempty" bson:"date,omitempty"`                            // When the incident occurred
	Type              *CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`                            // The nature of the accident
	LocationAddress   *Address         `json:"locationAddress,omitempty" bson:"location_address,omitempty"`     // Where the event occurred
	LocationReference *Reference       `json:"locationReference,omitempty" bson:"location_reference,omitempty"` // Where the event occurred
}

func (r *ExplanationOfBenefitAccident) Validate() error {
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

type ExplanationOfBenefitItemDetailSubDetail struct {
	Id                  *string                                `json:"id,omitempty" bson:"id,omitempty"`                                      // Unique id for inter-element referencing
	Sequence            int                                    `json:"sequence" bson:"sequence"`                                              // Product or service provided
	TraceNumber         []Identifier                           `json:"traceNumber,omitempty" bson:"trace_number,omitempty"`                   // Number for tracking
	Revenue             *CodeableConcept                       `json:"revenue,omitempty" bson:"revenue,omitempty"`                            // Revenue or cost center code
	Category            *CodeableConcept                       `json:"category,omitempty" bson:"category,omitempty"`                          // Benefit classification
	ProductOrService    *CodeableConcept                       `json:"productOrService,omitempty" bson:"product_or_service,omitempty"`        // Billing, service, product, or drug code
	ProductOrServiceEnd *CodeableConcept                       `json:"productOrServiceEnd,omitempty" bson:"product_or_service_end,omitempty"` // End of a range of codes
	Modifier            []CodeableConcept                      `json:"modifier,omitempty" bson:"modifier,omitempty"`                          // Service/Product billing modifiers
	ProgramCode         []CodeableConcept                      `json:"programCode,omitempty" bson:"program_code,omitempty"`                   // Program the product or service is provided under
	PatientPaid         *Money                                 `json:"patientPaid,omitempty" bson:"patient_paid,omitempty"`                   // Paid by the patient
	Quantity            *Quantity                              `json:"quantity,omitempty" bson:"quantity,omitempty"`                          // Count of products or services
	UnitPrice           *Money                                 `json:"unitPrice,omitempty" bson:"unit_price,omitempty"`                       // Fee, charge or cost per item
	Factor              *float64                               `json:"factor,omitempty" bson:"factor,omitempty"`                              // Price scaling factor
	Tax                 *Money                                 `json:"tax,omitempty" bson:"tax,omitempty"`                                    // Total tax
	Net                 *Money                                 `json:"net,omitempty" bson:"net,omitempty"`                                    // Total item cost
	Udi                 []Reference                            `json:"udi,omitempty" bson:"udi,omitempty"`                                    // Unique device identifier
	NoteNumber          []int                                  `json:"noteNumber,omitempty" bson:"note_number,omitempty"`                     // Applicable note numbers
	ReviewOutcome       *ExplanationOfBenefitItemReviewOutcome `json:"reviewOutcome,omitempty" bson:"review_outcome,omitempty"`               // Subdetail level adjudication results
	Adjudication        []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty" bson:"adjudication,omitempty"`                  // Subdetail level adjudication details
}

func (r *ExplanationOfBenefitItemDetailSubDetail) Validate() error {
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
	if r.ReviewOutcome != nil {
		if err := r.ReviewOutcome.Validate(); err != nil {
			return fmt.Errorf("ReviewOutcome: %w", err)
		}
	}
	for i, item := range r.Adjudication {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Adjudication[%d]: %w", i, err)
		}
	}
	return nil
}

type ExplanationOfBenefitAddItemDetail struct {
	Id                  *string                                      `json:"id,omitempty" bson:"id,omitempty"`                                      // Unique id for inter-element referencing
	TraceNumber         []Identifier                                 `json:"traceNumber,omitempty" bson:"trace_number,omitempty"`                   // Number for tracking
	Revenue             *CodeableConcept                             `json:"revenue,omitempty" bson:"revenue,omitempty"`                            // Revenue or cost center code
	ProductOrService    *CodeableConcept                             `json:"productOrService,omitempty" bson:"product_or_service,omitempty"`        // Billing, service, product, or drug code
	ProductOrServiceEnd *CodeableConcept                             `json:"productOrServiceEnd,omitempty" bson:"product_or_service_end,omitempty"` // End of a range of codes
	Modifier            []CodeableConcept                            `json:"modifier,omitempty" bson:"modifier,omitempty"`                          // Service/Product billing modifiers
	PatientPaid         *Money                                       `json:"patientPaid,omitempty" bson:"patient_paid,omitempty"`                   // Paid by the patient
	Quantity            *Quantity                                    `json:"quantity,omitempty" bson:"quantity,omitempty"`                          // Count of products or services
	UnitPrice           *Money                                       `json:"unitPrice,omitempty" bson:"unit_price,omitempty"`                       // Fee, charge or cost per item
	Factor              *float64                                     `json:"factor,omitempty" bson:"factor,omitempty"`                              // Price scaling factor
	Tax                 *Money                                       `json:"tax,omitempty" bson:"tax,omitempty"`                                    // Total tax
	Net                 *Money                                       `json:"net,omitempty" bson:"net,omitempty"`                                    // Total item cost
	NoteNumber          []int                                        `json:"noteNumber,omitempty" bson:"note_number,omitempty"`                     // Applicable note numbers
	ReviewOutcome       *ExplanationOfBenefitItemReviewOutcome       `json:"reviewOutcome,omitempty" bson:"review_outcome,omitempty"`               // Additem detail level adjudication results
	Adjudication        []ExplanationOfBenefitItemAdjudication       `json:"adjudication,omitempty" bson:"adjudication,omitempty"`                  // Added items adjudication
	SubDetail           []ExplanationOfBenefitAddItemDetailSubDetail `json:"subDetail,omitempty" bson:"sub_detail,omitempty"`                       // Insurer added line items
}

func (r *ExplanationOfBenefitAddItemDetail) Validate() error {
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
	if r.ReviewOutcome != nil {
		if err := r.ReviewOutcome.Validate(); err != nil {
			return fmt.Errorf("ReviewOutcome: %w", err)
		}
	}
	for i, item := range r.Adjudication {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Adjudication[%d]: %w", i, err)
		}
	}
	for i, item := range r.SubDetail {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SubDetail[%d]: %w", i, err)
		}
	}
	return nil
}

type ExplanationOfBenefitAddItemDetailSubDetail struct {
	Id                  *string                                `json:"id,omitempty" bson:"id,omitempty"`                                      // Unique id for inter-element referencing
	TraceNumber         []Identifier                           `json:"traceNumber,omitempty" bson:"trace_number,omitempty"`                   // Number for tracking
	Revenue             *CodeableConcept                       `json:"revenue,omitempty" bson:"revenue,omitempty"`                            // Revenue or cost center code
	ProductOrService    *CodeableConcept                       `json:"productOrService,omitempty" bson:"product_or_service,omitempty"`        // Billing, service, product, or drug code
	ProductOrServiceEnd *CodeableConcept                       `json:"productOrServiceEnd,omitempty" bson:"product_or_service_end,omitempty"` // End of a range of codes
	Modifier            []CodeableConcept                      `json:"modifier,omitempty" bson:"modifier,omitempty"`                          // Service/Product billing modifiers
	PatientPaid         *Money                                 `json:"patientPaid,omitempty" bson:"patient_paid,omitempty"`                   // Paid by the patient
	Quantity            *Quantity                              `json:"quantity,omitempty" bson:"quantity,omitempty"`                          // Count of products or services
	UnitPrice           *Money                                 `json:"unitPrice,omitempty" bson:"unit_price,omitempty"`                       // Fee, charge or cost per item
	Factor              *float64                               `json:"factor,omitempty" bson:"factor,omitempty"`                              // Price scaling factor
	Tax                 *Money                                 `json:"tax,omitempty" bson:"tax,omitempty"`                                    // Total tax
	Net                 *Money                                 `json:"net,omitempty" bson:"net,omitempty"`                                    // Total item cost
	NoteNumber          []int                                  `json:"noteNumber,omitempty" bson:"note_number,omitempty"`                     // Applicable note numbers
	ReviewOutcome       *ExplanationOfBenefitItemReviewOutcome `json:"reviewOutcome,omitempty" bson:"review_outcome,omitempty"`               // Additem subdetail level adjudication results
	Adjudication        []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty" bson:"adjudication,omitempty"`                  // Added items adjudication
}

func (r *ExplanationOfBenefitAddItemDetailSubDetail) Validate() error {
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
	if r.ReviewOutcome != nil {
		if err := r.ReviewOutcome.Validate(); err != nil {
			return fmt.Errorf("ReviewOutcome: %w", err)
		}
	}
	for i, item := range r.Adjudication {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Adjudication[%d]: %w", i, err)
		}
	}
	return nil
}

type ExplanationOfBenefitBenefitBalanceFinancial struct {
	Id                 *string          `json:"id,omitempty" bson:"id,omitempty"`                                   // Unique id for inter-element referencing
	Type               *CodeableConcept `json:"type" bson:"type"`                                                   // Benefit classification
	AllowedUnsignedInt *int             `json:"allowedUnsignedInt,omitempty" bson:"allowed_unsigned_int,omitempty"` // Benefits allowed
	AllowedString      *string          `json:"allowedString,omitempty" bson:"allowed_string,omitempty"`            // Benefits allowed
	AllowedMoney       *Money           `json:"allowedMoney,omitempty" bson:"allowed_money,omitempty"`              // Benefits allowed
	UsedUnsignedInt    *int             `json:"usedUnsignedInt,omitempty" bson:"used_unsigned_int,omitempty"`       // Benefits used
	UsedMoney          *Money           `json:"usedMoney,omitempty" bson:"used_money,omitempty"`                    // Benefits used
}

func (r *ExplanationOfBenefitBenefitBalanceFinancial) Validate() error {
	if r.Type == nil {
		return fmt.Errorf("field 'Type' is required")
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.AllowedMoney != nil {
		if err := r.AllowedMoney.Validate(); err != nil {
			return fmt.Errorf("AllowedMoney: %w", err)
		}
	}
	if r.UsedMoney != nil {
		if err := r.UsedMoney.Validate(); err != nil {
			return fmt.Errorf("UsedMoney: %w", err)
		}
	}
	return nil
}

type ExplanationOfBenefitPayee struct {
	Id    *string          `json:"id,omitempty" bson:"id,omitempty"`       // Unique id for inter-element referencing
	Type  *CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`   // Category of recipient
	Party *Reference       `json:"party,omitempty" bson:"party,omitempty"` // Recipient reference
}

func (r *ExplanationOfBenefitPayee) Validate() error {
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

type ExplanationOfBenefitCareTeam struct {
	Id        *string          `json:"id,omitempty" bson:"id,omitempty"`               // Unique id for inter-element referencing
	Sequence  int              `json:"sequence" bson:"sequence"`                       // Order of care team
	Provider  *Reference       `json:"provider" bson:"provider"`                       // Practitioner or organization
	Role      *CodeableConcept `json:"role,omitempty" bson:"role,omitempty"`           // Function within the team
	Specialty *CodeableConcept `json:"specialty,omitempty" bson:"specialty,omitempty"` // Practitioner or provider specialization
}

func (r *ExplanationOfBenefitCareTeam) Validate() error {
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

type ExplanationOfBenefitInsurance struct {
	Id         *string    `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Focal      bool       `json:"focal" bson:"focal"`                                 // Coverage to be used for adjudication
	Coverage   *Reference `json:"coverage" bson:"coverage"`                           // Insurance information
	PreAuthRef []string   `json:"preAuthRef,omitempty" bson:"pre_auth_ref,omitempty"` // Prior authorization reference number
}

func (r *ExplanationOfBenefitInsurance) Validate() error {
	if r.Coverage == nil {
		return fmt.Errorf("field 'Coverage' is required")
	}
	if r.Coverage != nil {
		if err := r.Coverage.Validate(); err != nil {
			return fmt.Errorf("Coverage: %w", err)
		}
	}
	return nil
}

type ExplanationOfBenefitItemAdjudication struct {
	Id           *string          `json:"id,omitempty" bson:"id,omitempty"`                      // Unique id for inter-element referencing
	Category     *CodeableConcept `json:"category" bson:"category"`                              // Type of adjudication information
	Reason       *CodeableConcept `json:"reason,omitempty" bson:"reason,omitempty"`              // Explanation of adjudication outcome
	Amount       *Money           `json:"amount,omitempty" bson:"amount,omitempty"`              // Monetary amount
	Quantity     *Quantity        `json:"quantity,omitempty" bson:"quantity,omitempty"`          // Non-monitary value
	DecisionDate *string          `json:"decisionDate,omitempty" bson:"decision_date,omitempty"` // When was adjudication performed
}

func (r *ExplanationOfBenefitItemAdjudication) Validate() error {
	if r.Category == nil {
		return fmt.Errorf("field 'Category' is required")
	}
	if r.Category != nil {
		if err := r.Category.Validate(); err != nil {
			return fmt.Errorf("Category: %w", err)
		}
	}
	if r.Reason != nil {
		if err := r.Reason.Validate(); err != nil {
			return fmt.Errorf("Reason: %w", err)
		}
	}
	if r.Amount != nil {
		if err := r.Amount.Validate(); err != nil {
			return fmt.Errorf("Amount: %w", err)
		}
	}
	if r.Quantity != nil {
		if err := r.Quantity.Validate(); err != nil {
			return fmt.Errorf("Quantity: %w", err)
		}
	}
	return nil
}

type ExplanationOfBenefitAddItemBodySite struct {
	Id      *string             `json:"id,omitempty" bson:"id,omitempty"`            // Unique id for inter-element referencing
	Site    []CodeableReference `json:"site" bson:"site"`                            // Location
	SubSite []CodeableConcept   `json:"subSite,omitempty" bson:"sub_site,omitempty"` // Sub-location
}

func (r *ExplanationOfBenefitAddItemBodySite) Validate() error {
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

type ExplanationOfBenefitTotal struct {
	Id       *string          `json:"id,omitempty" bson:"id,omitempty"` // Unique id for inter-element referencing
	Category *CodeableConcept `json:"category" bson:"category"`         // Type of adjudication information
	Amount   *Money           `json:"amount" bson:"amount"`             // Financial total for the category
}

func (r *ExplanationOfBenefitTotal) Validate() error {
	if r.Category == nil {
		return fmt.Errorf("field 'Category' is required")
	}
	if r.Category != nil {
		if err := r.Category.Validate(); err != nil {
			return fmt.Errorf("Category: %w", err)
		}
	}
	if r.Amount == nil {
		return fmt.Errorf("field 'Amount' is required")
	}
	if r.Amount != nil {
		if err := r.Amount.Validate(); err != nil {
			return fmt.Errorf("Amount: %w", err)
		}
	}
	return nil
}

type ExplanationOfBenefitProcedure struct {
	Id                       *string           `json:"id,omitempty" bson:"id,omitempty"`                           // Unique id for inter-element referencing
	Sequence                 int               `json:"sequence" bson:"sequence"`                                   // Procedure instance identifier
	Type                     []CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`                       // Category of Procedure
	Date                     *string           `json:"date,omitempty" bson:"date,omitempty"`                       // When the procedure was performed
	ProcedureCodeableConcept *CodeableConcept  `json:"procedureCodeableConcept" bson:"procedure_codeable_concept"` // Specific clinical procedure
	ProcedureReference       *Reference        `json:"procedureReference" bson:"procedure_reference"`              // Specific clinical procedure
	Udi                      []Reference       `json:"udi,omitempty" bson:"udi,omitempty"`                         // Unique device identifier
}

func (r *ExplanationOfBenefitProcedure) Validate() error {
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

type ExplanationOfBenefitItem struct {
	Id                      *string                                `json:"id,omitempty" bson:"id,omitempty"`                                             // Unique id for inter-element referencing
	Sequence                int                                    `json:"sequence" bson:"sequence"`                                                     // Item instance identifier
	CareTeamSequence        []int                                  `json:"careTeamSequence,omitempty" bson:"care_team_sequence,omitempty"`               // Applicable care team members
	DiagnosisSequence       []int                                  `json:"diagnosisSequence,omitempty" bson:"diagnosis_sequence,omitempty"`              // Applicable diagnoses
	ProcedureSequence       []int                                  `json:"procedureSequence,omitempty" bson:"procedure_sequence,omitempty"`              // Applicable procedures
	InformationSequence     []int                                  `json:"informationSequence,omitempty" bson:"information_sequence,omitempty"`          // Applicable exception and supporting information
	TraceNumber             []Identifier                           `json:"traceNumber,omitempty" bson:"trace_number,omitempty"`                          // Number for tracking
	Subject                 *Reference                             `json:"subject,omitempty" bson:"subject,omitempty"`                                   // The recipient of the products and services
	Revenue                 *CodeableConcept                       `json:"revenue,omitempty" bson:"revenue,omitempty"`                                   // Revenue or cost center code
	Category                *CodeableConcept                       `json:"category,omitempty" bson:"category,omitempty"`                                 // Benefit classification
	ProductOrService        *CodeableConcept                       `json:"productOrService,omitempty" bson:"product_or_service,omitempty"`               // Billing, service, product, or drug code
	ProductOrServiceEnd     *CodeableConcept                       `json:"productOrServiceEnd,omitempty" bson:"product_or_service_end,omitempty"`        // End of a range of codes
	Request                 []Reference                            `json:"request,omitempty" bson:"request,omitempty"`                                   // Request or Referral for Service
	Modifier                []CodeableConcept                      `json:"modifier,omitempty" bson:"modifier,omitempty"`                                 // Product or service billing modifiers
	ProgramCode             []CodeableConcept                      `json:"programCode,omitempty" bson:"program_code,omitempty"`                          // Program the product or service is provided under
	ServicedDate            *string                                `json:"servicedDate,omitempty" bson:"serviced_date,omitempty"`                        // Date or dates of service or product delivery
	ServicedPeriod          *Period                                `json:"servicedPeriod,omitempty" bson:"serviced_period,omitempty"`                    // Date or dates of service or product delivery
	LocationCodeableConcept *CodeableConcept                       `json:"locationCodeableConcept,omitempty" bson:"location_codeable_concept,omitempty"` // Place of service or where product was supplied
	LocationAddress         *Address                               `json:"locationAddress,omitempty" bson:"location_address,omitempty"`                  // Place of service or where product was supplied
	LocationReference       *Reference                             `json:"locationReference,omitempty" bson:"location_reference,omitempty"`              // Place of service or where product was supplied
	PatientPaid             *Money                                 `json:"patientPaid,omitempty" bson:"patient_paid,omitempty"`                          // Paid by the patient
	Quantity                *Quantity                              `json:"quantity,omitempty" bson:"quantity,omitempty"`                                 // Count of products or services
	UnitPrice               *Money                                 `json:"unitPrice,omitempty" bson:"unit_price,omitempty"`                              // Fee, charge or cost per item
	Factor                  *float64                               `json:"factor,omitempty" bson:"factor,omitempty"`                                     // Price scaling factor
	Tax                     *Money                                 `json:"tax,omitempty" bson:"tax,omitempty"`                                           // Total tax
	Net                     *Money                                 `json:"net,omitempty" bson:"net,omitempty"`                                           // Total item cost
	Udi                     []Reference                            `json:"udi,omitempty" bson:"udi,omitempty"`                                           // Unique device identifier
	BodySite                []ExplanationOfBenefitItemBodySite     `json:"bodySite,omitempty" bson:"body_site,omitempty"`                                // Anatomical location
	Encounter               []Reference                            `json:"encounter,omitempty" bson:"encounter,omitempty"`                               // Encounters associated with the listed treatments
	NoteNumber              []int                                  `json:"noteNumber,omitempty" bson:"note_number,omitempty"`                            // Applicable note numbers
	ReviewOutcome           *ExplanationOfBenefitItemReviewOutcome `json:"reviewOutcome,omitempty" bson:"review_outcome,omitempty"`                      // Adjudication results
	Adjudication            []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty" bson:"adjudication,omitempty"`                         // Adjudication details
	Detail                  []ExplanationOfBenefitItemDetail       `json:"detail,omitempty" bson:"detail,omitempty"`                                     // Additional items
}

func (r *ExplanationOfBenefitItem) Validate() error {
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
	if r.ReviewOutcome != nil {
		if err := r.ReviewOutcome.Validate(); err != nil {
			return fmt.Errorf("ReviewOutcome: %w", err)
		}
	}
	for i, item := range r.Adjudication {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Adjudication[%d]: %w", i, err)
		}
	}
	for i, item := range r.Detail {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Detail[%d]: %w", i, err)
		}
	}
	return nil
}

type ExplanationOfBenefitItemBodySite struct {
	Id      *string             `json:"id,omitempty" bson:"id,omitempty"`            // Unique id for inter-element referencing
	Site    []CodeableReference `json:"site" bson:"site"`                            // Location
	SubSite []CodeableConcept   `json:"subSite,omitempty" bson:"sub_site,omitempty"` // Sub-location
}

func (r *ExplanationOfBenefitItemBodySite) Validate() error {
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

type ExplanationOfBenefitItemReviewOutcome struct {
	Id            *string           `json:"id,omitempty" bson:"id,omitempty"`                         // Unique id for inter-element referencing
	Decision      *CodeableConcept  `json:"decision,omitempty" bson:"decision,omitempty"`             // Result of the adjudication
	Reason        []CodeableConcept `json:"reason,omitempty" bson:"reason,omitempty"`                 // Reason for result of the adjudication
	PreAuthRef    *string           `json:"preAuthRef,omitempty" bson:"pre_auth_ref,omitempty"`       // Preauthorization reference
	PreAuthPeriod *Period           `json:"preAuthPeriod,omitempty" bson:"pre_auth_period,omitempty"` // Preauthorization reference effective period
}

func (r *ExplanationOfBenefitItemReviewOutcome) Validate() error {
	if r.Decision != nil {
		if err := r.Decision.Validate(); err != nil {
			return fmt.Errorf("Decision: %w", err)
		}
	}
	for i, item := range r.Reason {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Reason[%d]: %w", i, err)
		}
	}
	if r.PreAuthPeriod != nil {
		if err := r.PreAuthPeriod.Validate(); err != nil {
			return fmt.Errorf("PreAuthPeriod: %w", err)
		}
	}
	return nil
}

type ExplanationOfBenefitItemDetail struct {
	Id                  *string                                   `json:"id,omitempty" bson:"id,omitempty"`                                      // Unique id for inter-element referencing
	Sequence            int                                       `json:"sequence" bson:"sequence"`                                              // Product or service provided
	TraceNumber         []Identifier                              `json:"traceNumber,omitempty" bson:"trace_number,omitempty"`                   // Number for tracking
	Revenue             *CodeableConcept                          `json:"revenue,omitempty" bson:"revenue,omitempty"`                            // Revenue or cost center code
	Category            *CodeableConcept                          `json:"category,omitempty" bson:"category,omitempty"`                          // Benefit classification
	ProductOrService    *CodeableConcept                          `json:"productOrService,omitempty" bson:"product_or_service,omitempty"`        // Billing, service, product, or drug code
	ProductOrServiceEnd *CodeableConcept                          `json:"productOrServiceEnd,omitempty" bson:"product_or_service_end,omitempty"` // End of a range of codes
	Modifier            []CodeableConcept                         `json:"modifier,omitempty" bson:"modifier,omitempty"`                          // Service/Product billing modifiers
	ProgramCode         []CodeableConcept                         `json:"programCode,omitempty" bson:"program_code,omitempty"`                   // Program the product or service is provided under
	PatientPaid         *Money                                    `json:"patientPaid,omitempty" bson:"patient_paid,omitempty"`                   // Paid by the patient
	Quantity            *Quantity                                 `json:"quantity,omitempty" bson:"quantity,omitempty"`                          // Count of products or services
	UnitPrice           *Money                                    `json:"unitPrice,omitempty" bson:"unit_price,omitempty"`                       // Fee, charge or cost per item
	Factor              *float64                                  `json:"factor,omitempty" bson:"factor,omitempty"`                              // Price scaling factor
	Tax                 *Money                                    `json:"tax,omitempty" bson:"tax,omitempty"`                                    // Total tax
	Net                 *Money                                    `json:"net,omitempty" bson:"net,omitempty"`                                    // Total item cost
	Udi                 []Reference                               `json:"udi,omitempty" bson:"udi,omitempty"`                                    // Unique device identifier
	NoteNumber          []int                                     `json:"noteNumber,omitempty" bson:"note_number,omitempty"`                     // Applicable note numbers
	ReviewOutcome       *ExplanationOfBenefitItemReviewOutcome    `json:"reviewOutcome,omitempty" bson:"review_outcome,omitempty"`               // Detail level adjudication results
	Adjudication        []ExplanationOfBenefitItemAdjudication    `json:"adjudication,omitempty" bson:"adjudication,omitempty"`                  // Detail level adjudication details
	SubDetail           []ExplanationOfBenefitItemDetailSubDetail `json:"subDetail,omitempty" bson:"sub_detail,omitempty"`                       // Additional items
}

func (r *ExplanationOfBenefitItemDetail) Validate() error {
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
	if r.ReviewOutcome != nil {
		if err := r.ReviewOutcome.Validate(); err != nil {
			return fmt.Errorf("ReviewOutcome: %w", err)
		}
	}
	for i, item := range r.Adjudication {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Adjudication[%d]: %w", i, err)
		}
	}
	for i, item := range r.SubDetail {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("SubDetail[%d]: %w", i, err)
		}
	}
	return nil
}

type ExplanationOfBenefitAddItem struct {
	Id                      *string                                `json:"id,omitempty" bson:"id,omitempty"`                                             // Unique id for inter-element referencing
	ItemSequence            []int                                  `json:"itemSequence,omitempty" bson:"item_sequence,omitempty"`                        // Item sequence number
	DetailSequence          []int                                  `json:"detailSequence,omitempty" bson:"detail_sequence,omitempty"`                    // Detail sequence number
	SubDetailSequence       []int                                  `json:"subDetailSequence,omitempty" bson:"sub_detail_sequence,omitempty"`             // Subdetail sequence number
	TraceNumber             []Identifier                           `json:"traceNumber,omitempty" bson:"trace_number,omitempty"`                          // Number for tracking
	Subject                 *Reference                             `json:"subject,omitempty" bson:"subject,omitempty"`                                   // The recipient of the products and services
	InformationSequence     []int                                  `json:"informationSequence,omitempty" bson:"information_sequence,omitempty"`          // Applicable exception and supporting information
	Provider                []Reference                            `json:"provider,omitempty" bson:"provider,omitempty"`                                 // Authorized providers
	Revenue                 *CodeableConcept                       `json:"revenue,omitempty" bson:"revenue,omitempty"`                                   // Revenue or cost center code
	Category                *CodeableConcept                       `json:"category,omitempty" bson:"category,omitempty"`                                 // Benefit classification
	ProductOrService        *CodeableConcept                       `json:"productOrService,omitempty" bson:"product_or_service,omitempty"`               // Billing, service, product, or drug code
	ProductOrServiceEnd     *CodeableConcept                       `json:"productOrServiceEnd,omitempty" bson:"product_or_service_end,omitempty"`        // End of a range of codes
	Request                 []Reference                            `json:"request,omitempty" bson:"request,omitempty"`                                   // Request or Referral for Service
	Modifier                []CodeableConcept                      `json:"modifier,omitempty" bson:"modifier,omitempty"`                                 // Service/Product billing modifiers
	ProgramCode             []CodeableConcept                      `json:"programCode,omitempty" bson:"program_code,omitempty"`                          // Program the product or service is provided under
	ServicedDate            *string                                `json:"servicedDate,omitempty" bson:"serviced_date,omitempty"`                        // Date or dates of service or product delivery
	ServicedPeriod          *Period                                `json:"servicedPeriod,omitempty" bson:"serviced_period,omitempty"`                    // Date or dates of service or product delivery
	LocationCodeableConcept *CodeableConcept                       `json:"locationCodeableConcept,omitempty" bson:"location_codeable_concept,omitempty"` // Place of service or where product was supplied
	LocationAddress         *Address                               `json:"locationAddress,omitempty" bson:"location_address,omitempty"`                  // Place of service or where product was supplied
	LocationReference       *Reference                             `json:"locationReference,omitempty" bson:"location_reference,omitempty"`              // Place of service or where product was supplied
	PatientPaid             *Money                                 `json:"patientPaid,omitempty" bson:"patient_paid,omitempty"`                          // Paid by the patient
	Quantity                *Quantity                              `json:"quantity,omitempty" bson:"quantity,omitempty"`                                 // Count of products or services
	UnitPrice               *Money                                 `json:"unitPrice,omitempty" bson:"unit_price,omitempty"`                              // Fee, charge or cost per item
	Factor                  *float64                               `json:"factor,omitempty" bson:"factor,omitempty"`                                     // Price scaling factor
	Tax                     *Money                                 `json:"tax,omitempty" bson:"tax,omitempty"`                                           // Total tax
	Net                     *Money                                 `json:"net,omitempty" bson:"net,omitempty"`                                           // Total item cost
	BodySite                []ExplanationOfBenefitAddItemBodySite  `json:"bodySite,omitempty" bson:"body_site,omitempty"`                                // Anatomical location
	NoteNumber              []int                                  `json:"noteNumber,omitempty" bson:"note_number,omitempty"`                            // Applicable note numbers
	ReviewOutcome           *ExplanationOfBenefitItemReviewOutcome `json:"reviewOutcome,omitempty" bson:"review_outcome,omitempty"`                      // Additem level adjudication results
	Adjudication            []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty" bson:"adjudication,omitempty"`                         // Added items adjudication
	Detail                  []ExplanationOfBenefitAddItemDetail    `json:"detail,omitempty" bson:"detail,omitempty"`                                     // Insurer added line items
}

func (r *ExplanationOfBenefitAddItem) Validate() error {
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
	for i, item := range r.Provider {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Provider[%d]: %w", i, err)
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
	for i, item := range r.BodySite {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("BodySite[%d]: %w", i, err)
		}
	}
	if r.ReviewOutcome != nil {
		if err := r.ReviewOutcome.Validate(); err != nil {
			return fmt.Errorf("ReviewOutcome: %w", err)
		}
	}
	for i, item := range r.Adjudication {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Adjudication[%d]: %w", i, err)
		}
	}
	for i, item := range r.Detail {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Detail[%d]: %w", i, err)
		}
	}
	return nil
}

type ExplanationOfBenefitPayment struct {
	Id               *string          `json:"id,omitempty" bson:"id,omitempty"`                              // Unique id for inter-element referencing
	Type             *CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`                          // Partial or complete payment
	Adjustment       *Money           `json:"adjustment,omitempty" bson:"adjustment,omitempty"`              // Payment adjustment for non-claim issues
	AdjustmentReason *CodeableConcept `json:"adjustmentReason,omitempty" bson:"adjustment_reason,omitempty"` // Explanation for the variance
	Date             *string          `json:"date,omitempty" bson:"date,omitempty"`                          // Expected date of payment
	Amount           *Money           `json:"amount,omitempty" bson:"amount,omitempty"`                      // Payable amount after adjustment
	Identifier       *Identifier      `json:"identifier,omitempty" bson:"identifier,omitempty"`              // Business identifier for the payment
}

func (r *ExplanationOfBenefitPayment) Validate() error {
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Adjustment != nil {
		if err := r.Adjustment.Validate(); err != nil {
			return fmt.Errorf("Adjustment: %w", err)
		}
	}
	if r.AdjustmentReason != nil {
		if err := r.AdjustmentReason.Validate(); err != nil {
			return fmt.Errorf("AdjustmentReason: %w", err)
		}
	}
	if r.Amount != nil {
		if err := r.Amount.Validate(); err != nil {
			return fmt.Errorf("Amount: %w", err)
		}
	}
	if r.Identifier != nil {
		if err := r.Identifier.Validate(); err != nil {
			return fmt.Errorf("Identifier: %w", err)
		}
	}
	return nil
}

type ExplanationOfBenefitBenefitBalance struct {
	Id          *string                                       `json:"id,omitempty" bson:"id,omitempty"`                   // Unique id for inter-element referencing
	Category    *CodeableConcept                              `json:"category" bson:"category"`                           // Benefit classification
	Excluded    *bool                                         `json:"excluded,omitempty" bson:"excluded,omitempty"`       // Excluded from the plan
	Name        *string                                       `json:"name,omitempty" bson:"name,omitempty"`               // Short name for the benefit
	Description *string                                       `json:"description,omitempty" bson:"description,omitempty"` // Description of the benefit or services covered
	Network     *CodeableConcept                              `json:"network,omitempty" bson:"network,omitempty"`         // In or out of network
	Unit        *CodeableConcept                              `json:"unit,omitempty" bson:"unit,omitempty"`               // Individual or family
	Term        *CodeableConcept                              `json:"term,omitempty" bson:"term,omitempty"`               // Annual or lifetime
	Financial   []ExplanationOfBenefitBenefitBalanceFinancial `json:"financial,omitempty" bson:"financial,omitempty"`     // Benefit Summary
}

func (r *ExplanationOfBenefitBenefitBalance) Validate() error {
	if r.Category == nil {
		return fmt.Errorf("field 'Category' is required")
	}
	if r.Category != nil {
		if err := r.Category.Validate(); err != nil {
			return fmt.Errorf("Category: %w", err)
		}
	}
	if r.Network != nil {
		if err := r.Network.Validate(); err != nil {
			return fmt.Errorf("Network: %w", err)
		}
	}
	if r.Unit != nil {
		if err := r.Unit.Validate(); err != nil {
			return fmt.Errorf("Unit: %w", err)
		}
	}
	if r.Term != nil {
		if err := r.Term.Validate(); err != nil {
			return fmt.Errorf("Term: %w", err)
		}
	}
	for i, item := range r.Financial {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("Financial[%d]: %w", i, err)
		}
	}
	return nil
}

type ExplanationOfBenefitRelated struct {
	Id           *string          `json:"id,omitempty" bson:"id,omitempty"`                     // Unique id for inter-element referencing
	Claim        *Reference       `json:"claim,omitempty" bson:"claim,omitempty"`               // Reference to the related claim
	Relationship *CodeableConcept `json:"relationship,omitempty" bson:"relationship,omitempty"` // How the reference claim is related
	Reference    *Identifier      `json:"reference,omitempty" bson:"reference,omitempty"`       // File or case reference
}

func (r *ExplanationOfBenefitRelated) Validate() error {
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

type ExplanationOfBenefitDiagnosis struct {
	Id                       *string           `json:"id,omitempty" bson:"id,omitempty"`                           // Unique id for inter-element referencing
	Sequence                 int               `json:"sequence" bson:"sequence"`                                   // Diagnosis instance identifier
	DiagnosisCodeableConcept *CodeableConcept  `json:"diagnosisCodeableConcept" bson:"diagnosis_codeable_concept"` // Nature of illness or problem
	DiagnosisReference       *Reference        `json:"diagnosisReference" bson:"diagnosis_reference"`              // Nature of illness or problem
	Type                     []CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`                       // Timing or nature of the diagnosis
	OnAdmission              *CodeableConcept  `json:"onAdmission,omitempty" bson:"on_admission,omitempty"`        // Present on admission
}

func (r *ExplanationOfBenefitDiagnosis) Validate() error {
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

type ExplanationOfBenefitProcessNote struct {
	Id       *string          `json:"id,omitempty" bson:"id,omitempty"`             // Unique id for inter-element referencing
	Class    *CodeableConcept `json:"class,omitempty" bson:"class,omitempty"`       // Business kind of note
	Number   *int             `json:"number,omitempty" bson:"number,omitempty"`     // Note instance identifier
	Type     *CodeableConcept `json:"type,omitempty" bson:"type,omitempty"`         // Note purpose
	Text     *string          `json:"text,omitempty" bson:"text,omitempty"`         // Note explanatory text
	Language *CodeableConcept `json:"language,omitempty" bson:"language,omitempty"` // Language of the text
}

func (r *ExplanationOfBenefitProcessNote) Validate() error {
	if r.Class != nil {
		if err := r.Class.Validate(); err != nil {
			return fmt.Errorf("Class: %w", err)
		}
	}
	if r.Type != nil {
		if err := r.Type.Validate(); err != nil {
			return fmt.Errorf("Type: %w", err)
		}
	}
	if r.Language != nil {
		if err := r.Language.Validate(); err != nil {
			return fmt.Errorf("Language: %w", err)
		}
	}
	return nil
}
