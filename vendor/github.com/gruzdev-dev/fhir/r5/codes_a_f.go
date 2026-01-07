package models

// AccountStatus represents codes from http://hl7.org/fhir/ValueSet/account-status
type AccountStatus string

const (
	AccountStatusActive         AccountStatus = "active"
	AccountStatusInactive       AccountStatus = "inactive"
	AccountStatusEnteredInError AccountStatus = "entered-in-error"
	AccountStatusOnHold         AccountStatus = "on-hold"
	AccountStatusUnknown        AccountStatus = "unknown"
)

// ActionApplicabilityBehavior represents codes from http://hl7.org/fhir/ValueSet/action-applicability-behavior
type ActionApplicabilityBehavior string

const (
	ActionApplicabilityBehaviorAll ActionApplicabilityBehavior = "all"
	ActionApplicabilityBehaviorAny ActionApplicabilityBehavior = "any"
)

// ActionCardinalityBehavior represents codes from http://hl7.org/fhir/ValueSet/action-cardinality-behavior
type ActionCardinalityBehavior string

const (
	ActionCardinalityBehaviorSingle   ActionCardinalityBehavior = "single"
	ActionCardinalityBehaviorMultiple ActionCardinalityBehavior = "multiple"
)

// ActionConditionKind represents codes from http://hl7.org/fhir/ValueSet/action-condition-kind
type ActionConditionKind string

const (
	ActionConditionKindApplicability ActionConditionKind = "applicability"
	ActionConditionKindStart         ActionConditionKind = "start"
	ActionConditionKindStop          ActionConditionKind = "stop"
)

// ActionGroupingBehavior represents codes from http://hl7.org/fhir/ValueSet/action-grouping-behavior
type ActionGroupingBehavior string

const (
	ActionGroupingBehaviorVisualGroup   ActionGroupingBehavior = "visual-group"
	ActionGroupingBehaviorLogicalGroup  ActionGroupingBehavior = "logical-group"
	ActionGroupingBehaviorSentenceGroup ActionGroupingBehavior = "sentence-group"
)

// ActionParticipantType represents codes from http://hl7.org/fhir/ValueSet/action-participant-type
type ActionParticipantType string

const (
	ActionParticipantTypeCareteam          ActionParticipantType = "careteam"
	ActionParticipantTypeDevice            ActionParticipantType = "device"
	ActionParticipantTypeGroup             ActionParticipantType = "group"
	ActionParticipantTypeHealthcareservice ActionParticipantType = "healthcareservice"
	ActionParticipantTypeLocation          ActionParticipantType = "location"
	ActionParticipantTypeOrganization      ActionParticipantType = "organization"
	ActionParticipantTypePatient           ActionParticipantType = "patient"
	ActionParticipantTypePractitioner      ActionParticipantType = "practitioner"
	ActionParticipantTypePractitionerrole  ActionParticipantType = "practitionerrole"
	ActionParticipantTypeRelatedperson     ActionParticipantType = "relatedperson"
)

// ActionPrecheckBehavior represents codes from http://hl7.org/fhir/ValueSet/action-precheck-behavior
type ActionPrecheckBehavior string

const (
	ActionPrecheckBehaviorYes ActionPrecheckBehavior = "yes"
	ActionPrecheckBehaviorNo  ActionPrecheckBehavior = "no"
)

// ActionRelationshipType represents codes from http://hl7.org/fhir/ValueSet/action-relationship-type
type ActionRelationshipType string

const (
	ActionRelationshipTypeBefore              ActionRelationshipType = "before"
	ActionRelationshipTypeBeforeStart         ActionRelationshipType = "before-start"
	ActionRelationshipTypeBeforeEnd           ActionRelationshipType = "before-end"
	ActionRelationshipTypeConcurrent          ActionRelationshipType = "concurrent"
	ActionRelationshipTypeConcurrentWithStart ActionRelationshipType = "concurrent-with-start"
	ActionRelationshipTypeConcurrentWithEnd   ActionRelationshipType = "concurrent-with-end"
	ActionRelationshipTypeAfter               ActionRelationshipType = "after"
	ActionRelationshipTypeAfterStart          ActionRelationshipType = "after-start"
	ActionRelationshipTypeAfterEnd            ActionRelationshipType = "after-end"
)

// ActionRequiredBehavior represents codes from http://hl7.org/fhir/ValueSet/action-required-behavior
type ActionRequiredBehavior string

const (
	ActionRequiredBehaviorMust                 ActionRequiredBehavior = "must"
	ActionRequiredBehaviorCould                ActionRequiredBehavior = "could"
	ActionRequiredBehaviorMustUnlessDocumented ActionRequiredBehavior = "must-unless-documented"
)

// ActionSelectionBehavior represents codes from http://hl7.org/fhir/ValueSet/action-selection-behavior
type ActionSelectionBehavior string

const (
	ActionSelectionBehaviorAny        ActionSelectionBehavior = "any"
	ActionSelectionBehaviorAll        ActionSelectionBehavior = "all"
	ActionSelectionBehaviorAllOrNone  ActionSelectionBehavior = "all-or-none"
	ActionSelectionBehaviorExactlyOne ActionSelectionBehavior = "exactly-one"
	ActionSelectionBehaviorAtMostOne  ActionSelectionBehavior = "at-most-one"
	ActionSelectionBehaviorOneOrMore  ActionSelectionBehavior = "one-or-more"
)

// ActorDefinitionActorType represents codes from http://hl7.org/fhir/ValueSet/actordefinition-actor-type
type ActorDefinitionActorType string

const (
	ActorDefinitionActorTypePerson     ActorDefinitionActorType = "person"
	ActorDefinitionActorTypeSystem     ActorDefinitionActorType = "system"
	ActorDefinitionActorTypeCollective ActorDefinitionActorType = "collective"
	ActorDefinitionActorTypeOther      ActorDefinitionActorType = "other"
)

// AdditionalBindingPurposeVS represents codes from http://hl7.org/fhir/ValueSet/additional-binding-purpose
type AdditionalBindingPurposeVS string

const (
	AdditionalBindingPurposeVSMaximum           AdditionalBindingPurposeVS = "maximum"
	AdditionalBindingPurposeVSMinimum           AdditionalBindingPurposeVS = "minimum"
	AdditionalBindingPurposeVSRequired          AdditionalBindingPurposeVS = "required"
	AdditionalBindingPurposeVSExtensible        AdditionalBindingPurposeVS = "extensible"
	AdditionalBindingPurposeVSCandidate         AdditionalBindingPurposeVS = "candidate"
	AdditionalBindingPurposeVSCurrent           AdditionalBindingPurposeVS = "current"
	AdditionalBindingPurposeVSCurrentExtensible AdditionalBindingPurposeVS = "current-extensible"
	AdditionalBindingPurposeVSBestPractice      AdditionalBindingPurposeVS = "best-practice"
	AdditionalBindingPurposeVSOpen              AdditionalBindingPurposeVS = "open"
	AdditionalBindingPurposeVSPreferred         AdditionalBindingPurposeVS = "preferred"
	AdditionalBindingPurposeVSUi                AdditionalBindingPurposeVS = "ui"
	AdditionalBindingPurposeVSStarter           AdditionalBindingPurposeVS = "starter"
	AdditionalBindingPurposeVSComponent         AdditionalBindingPurposeVS = "component"
)

// AddressType represents codes from http://hl7.org/fhir/ValueSet/address-type
type AddressType string

const (
	AddressTypePostal   AddressType = "postal"
	AddressTypePhysical AddressType = "physical"
	AddressTypeBoth     AddressType = "both"
)

// AddressUse represents codes from http://hl7.org/fhir/ValueSet/address-use
type AddressUse string

const (
	AddressUseHome    AddressUse = "home"
	AddressUseWork    AddressUse = "work"
	AddressUseTemp    AddressUse = "temp"
	AddressUseOld     AddressUse = "old"
	AddressUseBilling AddressUse = "billing"
)

// AdministrativeGender represents codes from http://hl7.org/fhir/ValueSet/administrative-gender
type AdministrativeGender string

const (
	AdministrativeGenderMale    AdministrativeGender = "male"
	AdministrativeGenderFemale  AdministrativeGender = "female"
	AdministrativeGenderOther   AdministrativeGender = "other"
	AdministrativeGenderUnknown AdministrativeGender = "unknown"
)

// AdverseEventActuality represents codes from http://hl7.org/fhir/ValueSet/adverse-event-actuality
type AdverseEventActuality string

const (
	AdverseEventActualityActual    AdverseEventActuality = "actual"
	AdverseEventActualityPotential AdverseEventActuality = "potential"
)

// AdverseEventStatus represents codes from http://hl7.org/fhir/ValueSet/adverse-event-status
type AdverseEventStatus string

const (
	AdverseEventStatusPreparation    AdverseEventStatus = "preparation"
	AdverseEventStatusInProgress     AdverseEventStatus = "in-progress"
	AdverseEventStatusNotDone        AdverseEventStatus = "not-done"
	AdverseEventStatusOnHold         AdverseEventStatus = "on-hold"
	AdverseEventStatusStopped        AdverseEventStatus = "stopped"
	AdverseEventStatusCompleted      AdverseEventStatus = "completed"
	AdverseEventStatusEnteredInError AdverseEventStatus = "entered-in-error"
	AdverseEventStatusUnknown        AdverseEventStatus = "unknown"
)

// AggregationMode represents codes from http://hl7.org/fhir/ValueSet/resource-aggregation-mode
type AggregationMode string

const (
	AggregationModeContained  AggregationMode = "contained"
	AggregationModeReferenced AggregationMode = "referenced"
	AggregationModeBundled    AggregationMode = "bundled"
)

// AllergyIntoleranceCategory represents codes from http://hl7.org/fhir/ValueSet/allergy-intolerance-category
type AllergyIntoleranceCategory string

const (
	AllergyIntoleranceCategoryFood        AllergyIntoleranceCategory = "food"
	AllergyIntoleranceCategoryMedication  AllergyIntoleranceCategory = "medication"
	AllergyIntoleranceCategoryEnvironment AllergyIntoleranceCategory = "environment"
	AllergyIntoleranceCategoryBiologic    AllergyIntoleranceCategory = "biologic"
)

// AllergyIntoleranceCriticality represents codes from http://hl7.org/fhir/ValueSet/allergy-intolerance-criticality
type AllergyIntoleranceCriticality string

const (
	AllergyIntoleranceCriticalityLow            AllergyIntoleranceCriticality = "low"
	AllergyIntoleranceCriticalityHigh           AllergyIntoleranceCriticality = "high"
	AllergyIntoleranceCriticalityUnableToAssess AllergyIntoleranceCriticality = "unable-to-assess"
)

// AllergyIntoleranceSeverity represents codes from http://hl7.org/fhir/ValueSet/reaction-event-severity
type AllergyIntoleranceSeverity string

const (
	AllergyIntoleranceSeverityMild     AllergyIntoleranceSeverity = "mild"
	AllergyIntoleranceSeverityModerate AllergyIntoleranceSeverity = "moderate"
	AllergyIntoleranceSeveritySevere   AllergyIntoleranceSeverity = "severe"
)

// AppointmentResponseStatus represents codes from http://hl7.org/fhir/ValueSet/appointmentresponse-status
type AppointmentResponseStatus string

const (
	AppointmentResponseStatusAccepted       AppointmentResponseStatus = "accepted"
	AppointmentResponseStatusDeclined       AppointmentResponseStatus = "declined"
	AppointmentResponseStatusTentative      AppointmentResponseStatus = "tentative"
	AppointmentResponseStatusNeedsAction    AppointmentResponseStatus = "needs-action"
	AppointmentResponseStatusProposed       AppointmentResponseStatus = "proposed"
	AppointmentResponseStatusPending        AppointmentResponseStatus = "pending"
	AppointmentResponseStatusBooked         AppointmentResponseStatus = "booked"
	AppointmentResponseStatusArrived        AppointmentResponseStatus = "arrived"
	AppointmentResponseStatusFulfilled      AppointmentResponseStatus = "fulfilled"
	AppointmentResponseStatusCancelled      AppointmentResponseStatus = "cancelled"
	AppointmentResponseStatusNoshow         AppointmentResponseStatus = "noshow"
	AppointmentResponseStatusEnteredInError AppointmentResponseStatus = "entered-in-error"
	AppointmentResponseStatusCheckedIn      AppointmentResponseStatus = "checked-in"
	AppointmentResponseStatusWaitlist       AppointmentResponseStatus = "waitlist"
)

// AppointmentStatus represents codes from http://hl7.org/fhir/ValueSet/appointmentstatus
type AppointmentStatus string

const (
	AppointmentStatusProposed       AppointmentStatus = "proposed"
	AppointmentStatusPending        AppointmentStatus = "pending"
	AppointmentStatusBooked         AppointmentStatus = "booked"
	AppointmentStatusArrived        AppointmentStatus = "arrived"
	AppointmentStatusFulfilled      AppointmentStatus = "fulfilled"
	AppointmentStatusCancelled      AppointmentStatus = "cancelled"
	AppointmentStatusNoshow         AppointmentStatus = "noshow"
	AppointmentStatusEnteredInError AppointmentStatus = "entered-in-error"
	AppointmentStatusCheckedIn      AppointmentStatus = "checked-in"
	AppointmentStatusWaitlist       AppointmentStatus = "waitlist"
)

// ArtifactAssessmentDisposition represents codes from http://hl7.org/fhir/ValueSet/artifactassessment-disposition
type ArtifactAssessmentDisposition string

const (
	ArtifactAssessmentDispositionUnresolved                    ArtifactAssessmentDisposition = "unresolved"
	ArtifactAssessmentDispositionNotPersuasive                 ArtifactAssessmentDisposition = "not-persuasive"
	ArtifactAssessmentDispositionPersuasive                    ArtifactAssessmentDisposition = "persuasive"
	ArtifactAssessmentDispositionPersuasiveWithModification    ArtifactAssessmentDisposition = "persuasive-with-modification"
	ArtifactAssessmentDispositionNotPersuasiveWithModification ArtifactAssessmentDisposition = "not-persuasive-with-modification"
)

// ArtifactAssessmentWorkflowStatus represents codes from http://hl7.org/fhir/ValueSet/artifactassessment-workflow-status
type ArtifactAssessmentWorkflowStatus string

const (
	ArtifactAssessmentWorkflowStatusSubmitted              ArtifactAssessmentWorkflowStatus = "submitted"
	ArtifactAssessmentWorkflowStatusTriaged                ArtifactAssessmentWorkflowStatus = "triaged"
	ArtifactAssessmentWorkflowStatusWaitingForInput        ArtifactAssessmentWorkflowStatus = "waiting-for-input"
	ArtifactAssessmentWorkflowStatusResolvedNoChange       ArtifactAssessmentWorkflowStatus = "resolved-no-change"
	ArtifactAssessmentWorkflowStatusResolvedChangeRequired ArtifactAssessmentWorkflowStatus = "resolved-change-required"
	ArtifactAssessmentWorkflowStatusDeferred               ArtifactAssessmentWorkflowStatus = "deferred"
	ArtifactAssessmentWorkflowStatusDuplicate              ArtifactAssessmentWorkflowStatus = "duplicate"
	ArtifactAssessmentWorkflowStatusApplied                ArtifactAssessmentWorkflowStatus = "applied"
	ArtifactAssessmentWorkflowStatusPublished              ArtifactAssessmentWorkflowStatus = "published"
	ArtifactAssessmentWorkflowStatusEnteredInError         ArtifactAssessmentWorkflowStatus = "entered-in-error"
)

// AuditEventAction represents codes from http://hl7.org/fhir/ValueSet/audit-event-action
type AuditEventAction string

const (
	AuditEventActionC AuditEventAction = "C"
	AuditEventActionR AuditEventAction = "R"
	AuditEventActionU AuditEventAction = "U"
	AuditEventActionD AuditEventAction = "D"
	AuditEventActionE AuditEventAction = "E"
)

// AuditEventSeverity represents codes from http://hl7.org/fhir/ValueSet/audit-event-severity
type AuditEventSeverity string

const (
	AuditEventSeverityEmergency     AuditEventSeverity = "emergency"
	AuditEventSeverityAlert         AuditEventSeverity = "alert"
	AuditEventSeverityCritical      AuditEventSeverity = "critical"
	AuditEventSeverityError         AuditEventSeverity = "error"
	AuditEventSeverityWarning       AuditEventSeverity = "warning"
	AuditEventSeverityNotice        AuditEventSeverity = "notice"
	AuditEventSeverityInformational AuditEventSeverity = "informational"
	AuditEventSeverityDebug         AuditEventSeverity = "debug"
)

// BindingStrength represents codes from http://hl7.org/fhir/ValueSet/binding-strength
type BindingStrength string

const (
	BindingStrengthRequired    BindingStrength = "required"
	BindingStrengthExtensible  BindingStrength = "extensible"
	BindingStrengthPreferred   BindingStrength = "preferred"
	BindingStrengthExample     BindingStrength = "example"
	BindingStrengthDescriptive BindingStrength = "descriptive"
)

// BundleType represents codes from http://hl7.org/fhir/ValueSet/bundle-type
type BundleType string

const (
	BundleTypeDocument                 BundleType = "document"
	BundleTypeMessage                  BundleType = "message"
	BundleTypeTransaction              BundleType = "transaction"
	BundleTypeTransactionResponse      BundleType = "transaction-response"
	BundleTypeBatch                    BundleType = "batch"
	BundleTypeBatchResponse            BundleType = "batch-response"
	BundleTypeHistory                  BundleType = "history"
	BundleTypeSearchset                BundleType = "searchset"
	BundleTypeCollection               BundleType = "collection"
	BundleTypeSubscriptionNotification BundleType = "subscription-notification"
)

// CapabilityStatementKind represents codes from http://hl7.org/fhir/ValueSet/capability-statement-kind
type CapabilityStatementKind string

const (
	CapabilityStatementKindInstance     CapabilityStatementKind = "instance"
	CapabilityStatementKindCapability   CapabilityStatementKind = "capability"
	CapabilityStatementKindRequirements CapabilityStatementKind = "requirements"
)

// CarePlanIntent represents codes from http://hl7.org/fhir/ValueSet/care-plan-intent
type CarePlanIntent string

const (
	CarePlanIntentProposal      CarePlanIntent = "proposal"
	CarePlanIntentSolicitOffer  CarePlanIntent = "solicit-offer"
	CarePlanIntentOfferResponse CarePlanIntent = "offer-response"
	CarePlanIntentPlan          CarePlanIntent = "plan"
	CarePlanIntentDirective     CarePlanIntent = "directive"
	CarePlanIntentOrder         CarePlanIntent = "order"
	CarePlanIntentOriginalOrder CarePlanIntent = "original-order"
	CarePlanIntentReflexOrder   CarePlanIntent = "reflex-order"
	CarePlanIntentFillerOrder   CarePlanIntent = "filler-order"
	CarePlanIntentInstanceOrder CarePlanIntent = "instance-order"
	CarePlanIntentOption        CarePlanIntent = "option"
)

// CareTeamStatus represents codes from http://hl7.org/fhir/ValueSet/care-team-status
type CareTeamStatus string

const (
	CareTeamStatusProposed       CareTeamStatus = "proposed"
	CareTeamStatusActive         CareTeamStatus = "active"
	CareTeamStatusSuspended      CareTeamStatus = "suspended"
	CareTeamStatusInactive       CareTeamStatus = "inactive"
	CareTeamStatusEnteredInError CareTeamStatus = "entered-in-error"
)

// ClaimProcessingCodes represents codes from http://hl7.org/fhir/ValueSet/claim-outcome
type ClaimProcessingCodes string

const (
	ClaimProcessingCodesQueued   ClaimProcessingCodes = "queued"
	ClaimProcessingCodesComplete ClaimProcessingCodes = "complete"
	ClaimProcessingCodesError    ClaimProcessingCodes = "error"
	ClaimProcessingCodesPartial  ClaimProcessingCodes = "partial"
)

// ClinicalUseDefinitionType represents codes from http://hl7.org/fhir/ValueSet/clinical-use-definition-type
type ClinicalUseDefinitionType string

const (
	ClinicalUseDefinitionTypeIndication        ClinicalUseDefinitionType = "indication"
	ClinicalUseDefinitionTypeContraindication  ClinicalUseDefinitionType = "contraindication"
	ClinicalUseDefinitionTypeInteraction       ClinicalUseDefinitionType = "interaction"
	ClinicalUseDefinitionTypeUndesirableEffect ClinicalUseDefinitionType = "undesirable-effect"
	ClinicalUseDefinitionTypeWarning           ClinicalUseDefinitionType = "warning"
)

// CodeSearchSupport represents codes from http://hl7.org/fhir/ValueSet/code-search-support
type CodeSearchSupport string

const (
	CodeSearchSupportInCompose            CodeSearchSupport = "in-compose"
	CodeSearchSupportInExpansion          CodeSearchSupport = "in-expansion"
	CodeSearchSupportInComposeOrExpansion CodeSearchSupport = "in-compose-or-expansion"
)

// CodeSystemContentMode represents codes from http://hl7.org/fhir/ValueSet/codesystem-content-mode
type CodeSystemContentMode string

const (
	CodeSystemContentModeNotPresent CodeSystemContentMode = "not-present"
	CodeSystemContentModeExample    CodeSystemContentMode = "example"
	CodeSystemContentModeFragment   CodeSystemContentMode = "fragment"
	CodeSystemContentModeComplete   CodeSystemContentMode = "complete"
	CodeSystemContentModeSupplement CodeSystemContentMode = "supplement"
)

// CodeSystemHierarchyMeaning represents codes from http://hl7.org/fhir/ValueSet/codesystem-hierarchy-meaning
type CodeSystemHierarchyMeaning string

const (
	CodeSystemHierarchyMeaningGroupedBy      CodeSystemHierarchyMeaning = "grouped-by"
	CodeSystemHierarchyMeaningIsA            CodeSystemHierarchyMeaning = "is-a"
	CodeSystemHierarchyMeaningPartOf         CodeSystemHierarchyMeaning = "part-of"
	CodeSystemHierarchyMeaningClassifiedWith CodeSystemHierarchyMeaning = "classified-with"
)

// ColorCodesOrRGB represents codes from http://hl7.org/fhir/ValueSet/color-codes
type ColorCodesOrRGB string

const (
	ColorCodesOrRGBAliceblue            ColorCodesOrRGB = "aliceblue"
	ColorCodesOrRGBAntiquewhite         ColorCodesOrRGB = "antiquewhite"
	ColorCodesOrRGBAqua                 ColorCodesOrRGB = "aqua"
	ColorCodesOrRGBAquamarine           ColorCodesOrRGB = "aquamarine"
	ColorCodesOrRGBAzure                ColorCodesOrRGB = "azure"
	ColorCodesOrRGBBeige                ColorCodesOrRGB = "beige"
	ColorCodesOrRGBBisque               ColorCodesOrRGB = "bisque"
	ColorCodesOrRGBBlack                ColorCodesOrRGB = "black"
	ColorCodesOrRGBBlanchedalmond       ColorCodesOrRGB = "blanchedalmond"
	ColorCodesOrRGBBlue                 ColorCodesOrRGB = "blue"
	ColorCodesOrRGBBlueviolet           ColorCodesOrRGB = "blueviolet"
	ColorCodesOrRGBBrown                ColorCodesOrRGB = "brown"
	ColorCodesOrRGBBurlywood            ColorCodesOrRGB = "burlywood"
	ColorCodesOrRGBCadetblue            ColorCodesOrRGB = "cadetblue"
	ColorCodesOrRGBChartreuse           ColorCodesOrRGB = "chartreuse"
	ColorCodesOrRGBChocolate            ColorCodesOrRGB = "chocolate"
	ColorCodesOrRGBCoral                ColorCodesOrRGB = "coral"
	ColorCodesOrRGBCornflowerblue       ColorCodesOrRGB = "cornflowerblue"
	ColorCodesOrRGBCornsilk             ColorCodesOrRGB = "cornsilk"
	ColorCodesOrRGBCrimson              ColorCodesOrRGB = "crimson"
	ColorCodesOrRGBCyan                 ColorCodesOrRGB = "cyan"
	ColorCodesOrRGBDarkblue             ColorCodesOrRGB = "darkblue"
	ColorCodesOrRGBDarkcyan             ColorCodesOrRGB = "darkcyan"
	ColorCodesOrRGBDarkgoldenrod        ColorCodesOrRGB = "darkgoldenrod"
	ColorCodesOrRGBDarkgray             ColorCodesOrRGB = "darkgray"
	ColorCodesOrRGBDarkgreen            ColorCodesOrRGB = "darkgreen"
	ColorCodesOrRGBDarkgrey             ColorCodesOrRGB = "darkgrey"
	ColorCodesOrRGBDarkkhaki            ColorCodesOrRGB = "darkkhaki"
	ColorCodesOrRGBDarkmagenta          ColorCodesOrRGB = "darkmagenta"
	ColorCodesOrRGBDarkolivegreen       ColorCodesOrRGB = "darkolivegreen"
	ColorCodesOrRGBDarkorange           ColorCodesOrRGB = "darkorange"
	ColorCodesOrRGBDarkorchid           ColorCodesOrRGB = "darkorchid"
	ColorCodesOrRGBDarkred              ColorCodesOrRGB = "darkred"
	ColorCodesOrRGBDarksalmon           ColorCodesOrRGB = "darksalmon"
	ColorCodesOrRGBDarkseagreen         ColorCodesOrRGB = "darkseagreen"
	ColorCodesOrRGBDarkslateblue        ColorCodesOrRGB = "darkslateblue"
	ColorCodesOrRGBDarkslategray        ColorCodesOrRGB = "darkslategray"
	ColorCodesOrRGBDarkslategrey        ColorCodesOrRGB = "darkslategrey"
	ColorCodesOrRGBDarkturquoise        ColorCodesOrRGB = "darkturquoise"
	ColorCodesOrRGBDarkviolet           ColorCodesOrRGB = "darkviolet"
	ColorCodesOrRGBDeeppink             ColorCodesOrRGB = "deeppink"
	ColorCodesOrRGBDeepskyblue          ColorCodesOrRGB = "deepskyblue"
	ColorCodesOrRGBDimgray              ColorCodesOrRGB = "dimgray"
	ColorCodesOrRGBDimgrey              ColorCodesOrRGB = "dimgrey"
	ColorCodesOrRGBDodgerblue           ColorCodesOrRGB = "dodgerblue"
	ColorCodesOrRGBFirebrick            ColorCodesOrRGB = "firebrick"
	ColorCodesOrRGBFloralwhite          ColorCodesOrRGB = "floralwhite"
	ColorCodesOrRGBForestgreen          ColorCodesOrRGB = "forestgreen"
	ColorCodesOrRGBFuchsia              ColorCodesOrRGB = "fuchsia"
	ColorCodesOrRGBGainsboro            ColorCodesOrRGB = "gainsboro"
	ColorCodesOrRGBGhostwhite           ColorCodesOrRGB = "ghostwhite"
	ColorCodesOrRGBGold                 ColorCodesOrRGB = "gold"
	ColorCodesOrRGBGoldenrod            ColorCodesOrRGB = "goldenrod"
	ColorCodesOrRGBGray                 ColorCodesOrRGB = "gray"
	ColorCodesOrRGBGreen                ColorCodesOrRGB = "green"
	ColorCodesOrRGBGreenyellow          ColorCodesOrRGB = "greenyellow"
	ColorCodesOrRGBGrey                 ColorCodesOrRGB = "grey"
	ColorCodesOrRGBHoneydew             ColorCodesOrRGB = "honeydew"
	ColorCodesOrRGBHotpink              ColorCodesOrRGB = "hotpink"
	ColorCodesOrRGBIndianred            ColorCodesOrRGB = "indianred"
	ColorCodesOrRGBIndigo               ColorCodesOrRGB = "indigo"
	ColorCodesOrRGBIvory                ColorCodesOrRGB = "ivory"
	ColorCodesOrRGBKhaki                ColorCodesOrRGB = "khaki"
	ColorCodesOrRGBLavender             ColorCodesOrRGB = "lavender"
	ColorCodesOrRGBLavenderblush        ColorCodesOrRGB = "lavenderblush"
	ColorCodesOrRGBLawngreen            ColorCodesOrRGB = "lawngreen"
	ColorCodesOrRGBLemonchiffon         ColorCodesOrRGB = "lemonchiffon"
	ColorCodesOrRGBLightblue            ColorCodesOrRGB = "lightblue"
	ColorCodesOrRGBLightcoral           ColorCodesOrRGB = "lightcoral"
	ColorCodesOrRGBLightcyan            ColorCodesOrRGB = "lightcyan"
	ColorCodesOrRGBLightgoldenrodyellow ColorCodesOrRGB = "lightgoldenrodyellow"
	ColorCodesOrRGBLightgray            ColorCodesOrRGB = "lightgray"
	ColorCodesOrRGBLightgreen           ColorCodesOrRGB = "lightgreen"
	ColorCodesOrRGBLightgrey            ColorCodesOrRGB = "lightgrey"
	ColorCodesOrRGBLightpink            ColorCodesOrRGB = "lightpink"
	ColorCodesOrRGBLightsalmon          ColorCodesOrRGB = "lightsalmon"
	ColorCodesOrRGBLightseagreen        ColorCodesOrRGB = "lightseagreen"
	ColorCodesOrRGBLightskyblue         ColorCodesOrRGB = "lightskyblue"
	ColorCodesOrRGBLightslategray       ColorCodesOrRGB = "lightslategray"
	ColorCodesOrRGBLightslategrey       ColorCodesOrRGB = "lightslategrey"
	ColorCodesOrRGBLightsteelblue       ColorCodesOrRGB = "lightsteelblue"
	ColorCodesOrRGBLightyellow          ColorCodesOrRGB = "lightyellow"
	ColorCodesOrRGBLime                 ColorCodesOrRGB = "lime"
	ColorCodesOrRGBLimegreen            ColorCodesOrRGB = "limegreen"
	ColorCodesOrRGBLinen                ColorCodesOrRGB = "linen"
	ColorCodesOrRGBMagenta              ColorCodesOrRGB = "magenta"
	ColorCodesOrRGBMaroon               ColorCodesOrRGB = "maroon"
	ColorCodesOrRGBMediumaquamarine     ColorCodesOrRGB = "mediumaquamarine"
	ColorCodesOrRGBMediumblue           ColorCodesOrRGB = "mediumblue"
	ColorCodesOrRGBMediumorchid         ColorCodesOrRGB = "mediumorchid"
	ColorCodesOrRGBMediumpurple         ColorCodesOrRGB = "mediumpurple"
	ColorCodesOrRGBMediumseagreen       ColorCodesOrRGB = "mediumseagreen"
	ColorCodesOrRGBMediumslateblue      ColorCodesOrRGB = "mediumslateblue"
	ColorCodesOrRGBMediumspringgreen    ColorCodesOrRGB = "mediumspringgreen"
	ColorCodesOrRGBMediumturquoise      ColorCodesOrRGB = "mediumturquoise"
	ColorCodesOrRGBMediumvioletred      ColorCodesOrRGB = "mediumvioletred"
	ColorCodesOrRGBMidnightblue         ColorCodesOrRGB = "midnightblue"
	ColorCodesOrRGBMintcream            ColorCodesOrRGB = "mintcream"
	ColorCodesOrRGBMistyrose            ColorCodesOrRGB = "mistyrose"
	ColorCodesOrRGBMoccasin             ColorCodesOrRGB = "moccasin"
	ColorCodesOrRGBNavajowhite          ColorCodesOrRGB = "navajowhite"
	ColorCodesOrRGBNavy                 ColorCodesOrRGB = "navy"
	ColorCodesOrRGBOldlace              ColorCodesOrRGB = "oldlace"
	ColorCodesOrRGBOlive                ColorCodesOrRGB = "olive"
	ColorCodesOrRGBOlivedrab            ColorCodesOrRGB = "olivedrab"
	ColorCodesOrRGBOrange               ColorCodesOrRGB = "orange"
	ColorCodesOrRGBOrangered            ColorCodesOrRGB = "orangered"
	ColorCodesOrRGBOrchid               ColorCodesOrRGB = "orchid"
	ColorCodesOrRGBPalegoldenrod        ColorCodesOrRGB = "palegoldenrod"
	ColorCodesOrRGBPalegreen            ColorCodesOrRGB = "palegreen"
	ColorCodesOrRGBPaleturquoise        ColorCodesOrRGB = "paleturquoise"
	ColorCodesOrRGBPalevioletred        ColorCodesOrRGB = "palevioletred"
	ColorCodesOrRGBPapayawhip           ColorCodesOrRGB = "papayawhip"
	ColorCodesOrRGBPeachpuff            ColorCodesOrRGB = "peachpuff"
	ColorCodesOrRGBPeru                 ColorCodesOrRGB = "peru"
	ColorCodesOrRGBPink                 ColorCodesOrRGB = "pink"
	ColorCodesOrRGBPlum                 ColorCodesOrRGB = "plum"
	ColorCodesOrRGBPowderblue           ColorCodesOrRGB = "powderblue"
	ColorCodesOrRGBPurple               ColorCodesOrRGB = "purple"
	ColorCodesOrRGBRebeccapurple        ColorCodesOrRGB = "rebeccapurple"
	ColorCodesOrRGBRed                  ColorCodesOrRGB = "red"
	ColorCodesOrRGBRosybrown            ColorCodesOrRGB = "rosybrown"
	ColorCodesOrRGBRoyalblue            ColorCodesOrRGB = "royalblue"
	ColorCodesOrRGBSaddlebrown          ColorCodesOrRGB = "saddlebrown"
	ColorCodesOrRGBSalmon               ColorCodesOrRGB = "salmon"
	ColorCodesOrRGBSandybrown           ColorCodesOrRGB = "sandybrown"
	ColorCodesOrRGBSeagreen             ColorCodesOrRGB = "seagreen"
	ColorCodesOrRGBSeashell             ColorCodesOrRGB = "seashell"
	ColorCodesOrRGBSienna               ColorCodesOrRGB = "sienna"
	ColorCodesOrRGBSilver               ColorCodesOrRGB = "silver"
	ColorCodesOrRGBSkyblue              ColorCodesOrRGB = "skyblue"
	ColorCodesOrRGBSlateblue            ColorCodesOrRGB = "slateblue"
	ColorCodesOrRGBSlategray            ColorCodesOrRGB = "slategray"
	ColorCodesOrRGBSlategrey            ColorCodesOrRGB = "slategrey"
	ColorCodesOrRGBSnow                 ColorCodesOrRGB = "snow"
	ColorCodesOrRGBSpringgreen          ColorCodesOrRGB = "springgreen"
	ColorCodesOrRGBSteelblue            ColorCodesOrRGB = "steelblue"
	ColorCodesOrRGBTan                  ColorCodesOrRGB = "tan"
	ColorCodesOrRGBTeal                 ColorCodesOrRGB = "teal"
	ColorCodesOrRGBThistle              ColorCodesOrRGB = "thistle"
	ColorCodesOrRGBTomato               ColorCodesOrRGB = "tomato"
	ColorCodesOrRGBTurquoise            ColorCodesOrRGB = "turquoise"
	ColorCodesOrRGBViolet               ColorCodesOrRGB = "violet"
	ColorCodesOrRGBWheat                ColorCodesOrRGB = "wheat"
	ColorCodesOrRGBWhite                ColorCodesOrRGB = "white"
	ColorCodesOrRGBWhitesmoke           ColorCodesOrRGB = "whitesmoke"
	ColorCodesOrRGBYellow               ColorCodesOrRGB = "yellow"
	ColorCodesOrRGBYellowgreen          ColorCodesOrRGB = "yellowgreen"
)

// ComparisonOperationVS represents codes from http://hl7.org/fhir/ValueSet/comparison-operation
type ComparisonOperationVS string

const (
	ComparisonOperationVSEq  ComparisonOperationVS = "eq"
	ComparisonOperationVSNe  ComparisonOperationVS = "ne"
	ComparisonOperationVSIn  ComparisonOperationVS = "in"
	ComparisonOperationVSNin ComparisonOperationVS = "nin"
	ComparisonOperationVSGt  ComparisonOperationVS = "gt"
	ComparisonOperationVSLt  ComparisonOperationVS = "lt"
	ComparisonOperationVSGe  ComparisonOperationVS = "ge"
	ComparisonOperationVSLe  ComparisonOperationVS = "le"
	ComparisonOperationVSSa  ComparisonOperationVS = "sa"
	ComparisonOperationVSEb  ComparisonOperationVS = "eb"
	ComparisonOperationVSAp  ComparisonOperationVS = "ap"
)

// CompartmentType represents codes from http://hl7.org/fhir/ValueSet/compartment-type
type CompartmentType string

const (
	CompartmentTypePatient       CompartmentType = "Patient"
	CompartmentTypeEncounter     CompartmentType = "Encounter"
	CompartmentTypeRelatedPerson CompartmentType = "RelatedPerson"
	CompartmentTypePractitioner  CompartmentType = "Practitioner"
	CompartmentTypeDevice        CompartmentType = "Device"
	CompartmentTypeEpisodeOfCare CompartmentType = "EpisodeOfCare"
)

// CompositionStatus represents codes from http://hl7.org/fhir/ValueSet/composition-status
type CompositionStatus string

const (
	CompositionStatusRegistered     CompositionStatus = "registered"
	CompositionStatusPartial        CompositionStatus = "partial"
	CompositionStatusPreliminary    CompositionStatus = "preliminary"
	CompositionStatusFinal          CompositionStatus = "final"
	CompositionStatusAmended        CompositionStatus = "amended"
	CompositionStatusCorrected      CompositionStatus = "corrected"
	CompositionStatusAppended       CompositionStatus = "appended"
	CompositionStatusCancelled      CompositionStatus = "cancelled"
	CompositionStatusEnteredInError CompositionStatus = "entered-in-error"
	CompositionStatusDeprecated     CompositionStatus = "deprecated"
	CompositionStatusUnknown        CompositionStatus = "unknown"
)

// ConceptMapAttributeType represents codes from http://hl7.org/fhir/ValueSet/conceptmap-attribute-type
type ConceptMapAttributeType string

const (
	ConceptMapAttributeTypeCode     ConceptMapAttributeType = "code"
	ConceptMapAttributeTypeCoding   ConceptMapAttributeType = "Coding"
	ConceptMapAttributeTypeString   ConceptMapAttributeType = "string"
	ConceptMapAttributeTypeBoolean  ConceptMapAttributeType = "boolean"
	ConceptMapAttributeTypeQuantity ConceptMapAttributeType = "Quantity"
)

// ConceptMapGroupUnmappedMode represents codes from http://hl7.org/fhir/ValueSet/conceptmap-unmapped-mode
type ConceptMapGroupUnmappedMode string

const (
	ConceptMapGroupUnmappedModeUseSourceCode ConceptMapGroupUnmappedMode = "use-source-code"
	ConceptMapGroupUnmappedModeFixed         ConceptMapGroupUnmappedMode = "fixed"
	ConceptMapGroupUnmappedModeOtherMap      ConceptMapGroupUnmappedMode = "other-map"
)

// ConceptMapPropertyType represents codes from http://hl7.org/fhir/ValueSet/conceptmap-property-type
type ConceptMapPropertyType string

const (
	ConceptMapPropertyTypeCoding   ConceptMapPropertyType = "Coding"
	ConceptMapPropertyTypeString   ConceptMapPropertyType = "string"
	ConceptMapPropertyTypeInteger  ConceptMapPropertyType = "integer"
	ConceptMapPropertyTypeBoolean  ConceptMapPropertyType = "boolean"
	ConceptMapPropertyTypeDateTime ConceptMapPropertyType = "dateTime"
	ConceptMapPropertyTypeDecimal  ConceptMapPropertyType = "decimal"
	ConceptMapPropertyTypeCode     ConceptMapPropertyType = "code"
)

// ConceptMapRelationship represents codes from http://hl7.org/fhir/ValueSet/concept-map-relationship
type ConceptMapRelationship string

const (
	ConceptMapRelationshipRelatedTo                  ConceptMapRelationship = "related-to"
	ConceptMapRelationshipEquivalent                 ConceptMapRelationship = "equivalent"
	ConceptMapRelationshipSourceIsNarrowerThanTarget ConceptMapRelationship = "source-is-narrower-than-target"
	ConceptMapRelationshipSourceIsBroaderThanTarget  ConceptMapRelationship = "source-is-broader-than-target"
	ConceptMapRelationshipNotRelatedTo               ConceptMapRelationship = "not-related-to"
)

// ConditionalDeleteStatus represents codes from http://hl7.org/fhir/ValueSet/conditional-delete-status
type ConditionalDeleteStatus string

const (
	ConditionalDeleteStatusNotSupported ConditionalDeleteStatus = "not-supported"
	ConditionalDeleteStatusSingle       ConditionalDeleteStatus = "single"
	ConditionalDeleteStatusMultiple     ConditionalDeleteStatus = "multiple"
)

// ConditionalReadStatus represents codes from http://hl7.org/fhir/ValueSet/conditional-read-status
type ConditionalReadStatus string

const (
	ConditionalReadStatusNotSupported  ConditionalReadStatus = "not-supported"
	ConditionalReadStatusModifiedSince ConditionalReadStatus = "modified-since"
	ConditionalReadStatusNotMatch      ConditionalReadStatus = "not-match"
	ConditionalReadStatusFullSupport   ConditionalReadStatus = "full-support"
)

// ConformanceExpectation represents codes from http://hl7.org/fhir/ValueSet/conformance-expectation
type ConformanceExpectation string

const (
	ConformanceExpectationSHALL     ConformanceExpectation = "SHALL"
	ConformanceExpectationSHOULD    ConformanceExpectation = "SHOULD"
	ConformanceExpectationMAY       ConformanceExpectation = "MAY"
	ConformanceExpectationSHOULDNOT ConformanceExpectation = "SHOULD-NOT"
	ConformanceExpectationSHALLNOT  ConformanceExpectation = "SHALL-NOT"
)

// ConsentDataMeaning represents codes from http://hl7.org/fhir/ValueSet/consent-data-meaning
type ConsentDataMeaning string

const (
	ConsentDataMeaningInstance   ConsentDataMeaning = "instance"
	ConsentDataMeaningRelated    ConsentDataMeaning = "related"
	ConsentDataMeaningDependents ConsentDataMeaning = "dependents"
	ConsentDataMeaningAuthoredby ConsentDataMeaning = "authoredby"
)

// ConsentProvisionType represents codes from http://hl7.org/fhir/ValueSet/consent-provision-type
type ConsentProvisionType string

const (
	ConsentProvisionTypeDeny   ConsentProvisionType = "deny"
	ConsentProvisionTypePermit ConsentProvisionType = "permit"
)

// ConsentState represents codes from http://hl7.org/fhir/ValueSet/consent-state-codes
type ConsentState string

const (
	ConsentStateDraft          ConsentState = "draft"
	ConsentStateActive         ConsentState = "active"
	ConsentStateInactive       ConsentState = "inactive"
	ConsentStateNotDone        ConsentState = "not-done"
	ConsentStateEnteredInError ConsentState = "entered-in-error"
	ConsentStateUnknown        ConsentState = "unknown"
)

// ConstraintSeverity represents codes from http://hl7.org/fhir/ValueSet/constraint-severity
type ConstraintSeverity string

const (
	ConstraintSeverityError   ConstraintSeverity = "error"
	ConstraintSeverityWarning ConstraintSeverity = "warning"
)

// ContactPointSystem represents codes from http://hl7.org/fhir/ValueSet/contact-point-system
type ContactPointSystem string

const (
	ContactPointSystemPhone ContactPointSystem = "phone"
	ContactPointSystemFax   ContactPointSystem = "fax"
	ContactPointSystemEmail ContactPointSystem = "email"
	ContactPointSystemPager ContactPointSystem = "pager"
	ContactPointSystemUrl   ContactPointSystem = "url"
	ContactPointSystemSms   ContactPointSystem = "sms"
	ContactPointSystemOther ContactPointSystem = "other"
)

// ContactPointUse represents codes from http://hl7.org/fhir/ValueSet/contact-point-use
type ContactPointUse string

const (
	ContactPointUseHome   ContactPointUse = "home"
	ContactPointUseWork   ContactPointUse = "work"
	ContactPointUseTemp   ContactPointUse = "temp"
	ContactPointUseOld    ContactPointUse = "old"
	ContactPointUseMobile ContactPointUse = "mobile"
)

// ContractResourcePublicationStatusCodes represents codes from http://hl7.org/fhir/ValueSet/contract-publicationstatus
type ContractResourcePublicationStatusCodes string

const (
	ContractResourcePublicationStatusCodesAmended        ContractResourcePublicationStatusCodes = "amended"
	ContractResourcePublicationStatusCodesAppended       ContractResourcePublicationStatusCodes = "appended"
	ContractResourcePublicationStatusCodesCancelled      ContractResourcePublicationStatusCodes = "cancelled"
	ContractResourcePublicationStatusCodesDisputed       ContractResourcePublicationStatusCodes = "disputed"
	ContractResourcePublicationStatusCodesEnteredInError ContractResourcePublicationStatusCodes = "entered-in-error"
	ContractResourcePublicationStatusCodesExecutable     ContractResourcePublicationStatusCodes = "executable"
	ContractResourcePublicationStatusCodesExecuted       ContractResourcePublicationStatusCodes = "executed"
	ContractResourcePublicationStatusCodesNegotiable     ContractResourcePublicationStatusCodes = "negotiable"
	ContractResourcePublicationStatusCodesOffered        ContractResourcePublicationStatusCodes = "offered"
	ContractResourcePublicationStatusCodesPolicy         ContractResourcePublicationStatusCodes = "policy"
	ContractResourcePublicationStatusCodesRejected       ContractResourcePublicationStatusCodes = "rejected"
	ContractResourcePublicationStatusCodesRenewed        ContractResourcePublicationStatusCodes = "renewed"
	ContractResourcePublicationStatusCodesRevoked        ContractResourcePublicationStatusCodes = "revoked"
	ContractResourcePublicationStatusCodesResolved       ContractResourcePublicationStatusCodes = "resolved"
	ContractResourcePublicationStatusCodesTerminated     ContractResourcePublicationStatusCodes = "terminated"
)

// ContractResourceStatusCodes represents codes from http://hl7.org/fhir/ValueSet/contract-status
type ContractResourceStatusCodes string

const (
	ContractResourceStatusCodesAmended        ContractResourceStatusCodes = "amended"
	ContractResourceStatusCodesAppended       ContractResourceStatusCodes = "appended"
	ContractResourceStatusCodesCancelled      ContractResourceStatusCodes = "cancelled"
	ContractResourceStatusCodesDisputed       ContractResourceStatusCodes = "disputed"
	ContractResourceStatusCodesEnteredInError ContractResourceStatusCodes = "entered-in-error"
	ContractResourceStatusCodesExecutable     ContractResourceStatusCodes = "executable"
	ContractResourceStatusCodesExecuted       ContractResourceStatusCodes = "executed"
	ContractResourceStatusCodesNegotiable     ContractResourceStatusCodes = "negotiable"
	ContractResourceStatusCodesOffered        ContractResourceStatusCodes = "offered"
	ContractResourceStatusCodesPolicy         ContractResourceStatusCodes = "policy"
	ContractResourceStatusCodesRejected       ContractResourceStatusCodes = "rejected"
	ContractResourceStatusCodesRenewed        ContractResourceStatusCodes = "renewed"
	ContractResourceStatusCodesRevoked        ContractResourceStatusCodes = "revoked"
	ContractResourceStatusCodesResolved       ContractResourceStatusCodes = "resolved"
	ContractResourceStatusCodesTerminated     ContractResourceStatusCodes = "terminated"
)

// CriteriaNotExistsBehavior represents codes from http://hl7.org/fhir/ValueSet/subscriptiontopic-cr-behavior
type CriteriaNotExistsBehavior string

const (
	CriteriaNotExistsBehaviorTestPasses CriteriaNotExistsBehavior = "test-passes"
	CriteriaNotExistsBehaviorTestFails  CriteriaNotExistsBehavior = "test-fails"
)

// DaysOfWeek represents codes from http://hl7.org/fhir/ValueSet/days-of-week
type DaysOfWeek string

const (
	DaysOfWeekMon DaysOfWeek = "mon"
	DaysOfWeekTue DaysOfWeek = "tue"
	DaysOfWeekWed DaysOfWeek = "wed"
	DaysOfWeekThu DaysOfWeek = "thu"
	DaysOfWeekFri DaysOfWeek = "fri"
	DaysOfWeekSat DaysOfWeek = "sat"
	DaysOfWeekSun DaysOfWeek = "sun"
)

// DetectedIssueStatus represents codes from http://hl7.org/fhir/ValueSet/detectedissue-status
type DetectedIssueStatus string

const (
	DetectedIssueStatusRegistered        DetectedIssueStatus = "registered"
	DetectedIssueStatusSpecimenInProcess DetectedIssueStatus = "specimen-in-process"
	DetectedIssueStatusPreliminary       DetectedIssueStatus = "preliminary"
	DetectedIssueStatusFinal             DetectedIssueStatus = "final"
	DetectedIssueStatusAmended           DetectedIssueStatus = "amended"
	DetectedIssueStatusCorrected         DetectedIssueStatus = "corrected"
	DetectedIssueStatusAppended          DetectedIssueStatus = "appended"
	DetectedIssueStatusCancelled         DetectedIssueStatus = "cancelled"
	DetectedIssueStatusEnteredInError    DetectedIssueStatus = "entered-in-error"
	DetectedIssueStatusUnknown           DetectedIssueStatus = "unknown"
	DetectedIssueStatusCannotBeObtained  DetectedIssueStatus = "cannot-be-obtained"
	DetectedIssueStatusMitigated         DetectedIssueStatus = "mitigated"
	DetectedIssueStatusProcessingError   DetectedIssueStatus = "processing-error"
)

// DeviceAlertAnnunciationCodes represents codes from http://hl7.org/fhir/ValueSet/devicealert-annunciation
type DeviceAlertAnnunciationCodes string

const (
	DeviceAlertAnnunciationCodesLocal  DeviceAlertAnnunciationCodes = "local"
	DeviceAlertAnnunciationCodesRemote DeviceAlertAnnunciationCodes = "remote"
)

// DeviceAlertStatusCodes represents codes from http://hl7.org/fhir/ValueSet/devicealert-status
type DeviceAlertStatusCodes string

const (
	DeviceAlertStatusCodesInProgress     DeviceAlertStatusCodes = "in-progress"
	DeviceAlertStatusCodesCompleted      DeviceAlertStatusCodes = "completed"
	DeviceAlertStatusCodesEnteredInError DeviceAlertStatusCodes = "entered-in-error"
	DeviceAlertStatusCodesUnknown        DeviceAlertStatusCodes = "unknown"
)

// DeviceAssociationCodes represents codes from http://hl7.org/fhir/ValueSet/deviceassociation-status
type DeviceAssociationCodes string

const (
	DeviceAssociationCodesActive         DeviceAssociationCodes = "active"
	DeviceAssociationCodesInactive       DeviceAssociationCodes = "inactive"
	DeviceAssociationCodesEnteredInError DeviceAssociationCodes = "entered-in-error"
	DeviceAssociationCodesUnknown        DeviceAssociationCodes = "unknown"
)

// DeviceCorrectiveActionScope represents codes from http://hl7.org/fhir/ValueSet/device-correctiveactionscope
type DeviceCorrectiveActionScope string

const (
	DeviceCorrectiveActionScopeModel         DeviceCorrectiveActionScope = "model"
	DeviceCorrectiveActionScopeLotNumbers    DeviceCorrectiveActionScope = "lot-numbers"
	DeviceCorrectiveActionScopeSerialNumbers DeviceCorrectiveActionScope = "serial-numbers"
)

// DeviceDefinitionRegulatoryIdentifierType represents codes from http://hl7.org/fhir/ValueSet/devicedefinition-regulatory-identifier-type
type DeviceDefinitionRegulatoryIdentifierType string

const (
	DeviceDefinitionRegulatoryIdentifierTypeBasic   DeviceDefinitionRegulatoryIdentifierType = "basic"
	DeviceDefinitionRegulatoryIdentifierTypeMaster  DeviceDefinitionRegulatoryIdentifierType = "master"
	DeviceDefinitionRegulatoryIdentifierTypeLicense DeviceDefinitionRegulatoryIdentifierType = "license"
)

// DeviceMetricCalibrationState represents codes from http://hl7.org/fhir/ValueSet/metric-calibration-state
type DeviceMetricCalibrationState string

const (
	DeviceMetricCalibrationStateNotCalibrated       DeviceMetricCalibrationState = "not-calibrated"
	DeviceMetricCalibrationStateCalibrationRequired DeviceMetricCalibrationState = "calibration-required"
	DeviceMetricCalibrationStateCalibrated          DeviceMetricCalibrationState = "calibrated"
	DeviceMetricCalibrationStateUnspecified         DeviceMetricCalibrationState = "unspecified"
)

// DeviceMetricOperationalStatus represents codes from http://hl7.org/fhir/ValueSet/metric-operational-status
type DeviceMetricOperationalStatus string

const (
	DeviceMetricOperationalStatusOn      DeviceMetricOperationalStatus = "on"
	DeviceMetricOperationalStatusOff     DeviceMetricOperationalStatus = "off"
	DeviceMetricOperationalStatusStandby DeviceMetricOperationalStatus = "standby"
	DeviceMetricOperationalStatusUnknown DeviceMetricOperationalStatus = "unknown"
)

// DeviceMetricStatus represents codes from http://hl7.org/fhir/ValueSet/metric-status
type DeviceMetricStatus string

const (
	DeviceMetricStatusActive         DeviceMetricStatus = "active"
	DeviceMetricStatusInactive       DeviceMetricStatus = "inactive"
	DeviceMetricStatusEnteredInError DeviceMetricStatus = "entered-in-error"
	DeviceMetricStatusUnknown        DeviceMetricStatus = "unknown"
)

// DiagnosticReportStatus represents codes from http://hl7.org/fhir/ValueSet/diagnostic-report-status
type DiagnosticReportStatus string

const (
	DiagnosticReportStatusRegistered     DiagnosticReportStatus = "registered"
	DiagnosticReportStatusPartial        DiagnosticReportStatus = "partial"
	DiagnosticReportStatusPreliminary    DiagnosticReportStatus = "preliminary"
	DiagnosticReportStatusModified       DiagnosticReportStatus = "modified"
	DiagnosticReportStatusFinal          DiagnosticReportStatus = "final"
	DiagnosticReportStatusAmended        DiagnosticReportStatus = "amended"
	DiagnosticReportStatusCorrected      DiagnosticReportStatus = "corrected"
	DiagnosticReportStatusAppended       DiagnosticReportStatus = "appended"
	DiagnosticReportStatusCancelled      DiagnosticReportStatus = "cancelled"
	DiagnosticReportStatusEnteredInError DiagnosticReportStatus = "entered-in-error"
	DiagnosticReportStatusUnknown        DiagnosticReportStatus = "unknown"
)

// DiscriminatorType represents codes from http://hl7.org/fhir/ValueSet/discriminator-type
type DiscriminatorType string

const (
	DiscriminatorTypeValue    DiscriminatorType = "value"
	DiscriminatorTypeExists   DiscriminatorType = "exists"
	DiscriminatorTypePattern  DiscriminatorType = "pattern"
	DiscriminatorTypeType     DiscriminatorType = "type"
	DiscriminatorTypeProfile  DiscriminatorType = "profile"
	DiscriminatorTypePosition DiscriminatorType = "position"
)

// DocumentMode represents codes from http://hl7.org/fhir/ValueSet/document-mode
type DocumentMode string

const (
	DocumentModeProducer DocumentMode = "producer"
	DocumentModeConsumer DocumentMode = "consumer"
)

// DocumentReferenceStatus represents codes from http://hl7.org/fhir/ValueSet/document-reference-status
type DocumentReferenceStatus string

const (
	DocumentReferenceStatusCurrent        DocumentReferenceStatus = "current"
	DocumentReferenceStatusSuperseded     DocumentReferenceStatus = "superseded"
	DocumentReferenceStatusEnteredInError DocumentReferenceStatus = "entered-in-error"
)

// DoseLimitScopeVS represents codes from http://hl7.org/fhir/ValueSet/dose-limit-scope
type DoseLimitScopeVS string

const (
	DoseLimitScopeVSDosage         DoseLimitScopeVS = "dosage"
	DoseLimitScopeVSPeriod         DoseLimitScopeVS = "period"
	DoseLimitScopeVSAdministration DoseLimitScopeVS = "administration"
	DoseLimitScopeVSLifetime       DoseLimitScopeVS = "lifetime"
)

// EligibilityOutcome represents codes from http://hl7.org/fhir/ValueSet/eligibility-outcome
type EligibilityOutcome string

const (
	EligibilityOutcomeQueued   EligibilityOutcome = "queued"
	EligibilityOutcomeComplete EligibilityOutcome = "complete"
	EligibilityOutcomeError    EligibilityOutcome = "error"
	EligibilityOutcomePartial  EligibilityOutcome = "partial"
)

// EligibilityRequestPurpose represents codes from http://hl7.org/fhir/ValueSet/eligibilityrequest-purpose
type EligibilityRequestPurpose string

const (
	EligibilityRequestPurposeAuthRequirements EligibilityRequestPurpose = "auth-requirements"
	EligibilityRequestPurposeBenefits         EligibilityRequestPurpose = "benefits"
	EligibilityRequestPurposeDiscovery        EligibilityRequestPurpose = "discovery"
	EligibilityRequestPurposeValidation       EligibilityRequestPurpose = "validation"
)

// EligibilityResponsePurpose represents codes from http://hl7.org/fhir/ValueSet/eligibilityresponse-purpose
type EligibilityResponsePurpose string

const (
	EligibilityResponsePurposeAuthRequirements EligibilityResponsePurpose = "auth-requirements"
	EligibilityResponsePurposeBenefits         EligibilityResponsePurpose = "benefits"
	EligibilityResponsePurposeDiscovery        EligibilityResponsePurpose = "discovery"
	EligibilityResponsePurposeValidation       EligibilityResponsePurpose = "validation"
)

// EnableWhenBehavior represents codes from http://hl7.org/fhir/ValueSet/questionnaire-enable-behavior
type EnableWhenBehavior string

const (
	EnableWhenBehaviorAll EnableWhenBehavior = "all"
	EnableWhenBehaviorAny EnableWhenBehavior = "any"
)

// EncounterLocationStatus represents codes from http://hl7.org/fhir/ValueSet/encounter-location-status
type EncounterLocationStatus string

const (
	EncounterLocationStatusPlanned   EncounterLocationStatus = "planned"
	EncounterLocationStatusActive    EncounterLocationStatus = "active"
	EncounterLocationStatusReserved  EncounterLocationStatus = "reserved"
	EncounterLocationStatusCompleted EncounterLocationStatus = "completed"
)

// EncounterStatus represents codes from http://hl7.org/fhir/ValueSet/encounter-status
type EncounterStatus string

const (
	EncounterStatusPlanned        EncounterStatus = "planned"
	EncounterStatusInProgress     EncounterStatus = "in-progress"
	EncounterStatusOnHold         EncounterStatus = "on-hold"
	EncounterStatusDischarged     EncounterStatus = "discharged"
	EncounterStatusCompleted      EncounterStatus = "completed"
	EncounterStatusCancelled      EncounterStatus = "cancelled"
	EncounterStatusDiscontinued   EncounterStatus = "discontinued"
	EncounterStatusEnteredInError EncounterStatus = "entered-in-error"
	EncounterStatusUnknown        EncounterStatus = "unknown"
)

// EndpointStatus represents codes from http://hl7.org/fhir/ValueSet/endpoint-status
type EndpointStatus string

const (
	EndpointStatusActive         EndpointStatus = "active"
	EndpointStatusLimited        EndpointStatus = "limited"
	EndpointStatusSuspended      EndpointStatus = "suspended"
	EndpointStatusError          EndpointStatus = "error"
	EndpointStatusOff            EndpointStatus = "off"
	EndpointStatusEnteredInError EndpointStatus = "entered-in-error"
)

// EnrollmentOutcome represents codes from http://hl7.org/fhir/ValueSet/enrollment-outcome
type EnrollmentOutcome string

const (
	EnrollmentOutcomeQueued   EnrollmentOutcome = "queued"
	EnrollmentOutcomeComplete EnrollmentOutcome = "complete"
	EnrollmentOutcomeError    EnrollmentOutcome = "error"
	EnrollmentOutcomePartial  EnrollmentOutcome = "partial"
)

// EpisodeOfCareStatus represents codes from http://hl7.org/fhir/ValueSet/episode-of-care-status
type EpisodeOfCareStatus string

const (
	EpisodeOfCareStatusPlanned        EpisodeOfCareStatus = "planned"
	EpisodeOfCareStatusWaitlist       EpisodeOfCareStatus = "waitlist"
	EpisodeOfCareStatusActive         EpisodeOfCareStatus = "active"
	EpisodeOfCareStatusOnhold         EpisodeOfCareStatus = "onhold"
	EpisodeOfCareStatusFinished       EpisodeOfCareStatus = "finished"
	EpisodeOfCareStatusCancelled      EpisodeOfCareStatus = "cancelled"
	EpisodeOfCareStatusEnteredInError EpisodeOfCareStatus = "entered-in-error"
)

// EventCapabilityMode represents codes from http://hl7.org/fhir/ValueSet/event-capability-mode
type EventCapabilityMode string

const (
	EventCapabilityModeSender   EventCapabilityMode = "sender"
	EventCapabilityModeReceiver EventCapabilityMode = "receiver"
)

// EventStatus represents codes from http://hl7.org/fhir/ValueSet/event-status
type EventStatus string

const (
	EventStatusPreparation    EventStatus = "preparation"
	EventStatusInProgress     EventStatus = "in-progress"
	EventStatusNotDone        EventStatus = "not-done"
	EventStatusOnHold         EventStatus = "on-hold"
	EventStatusStopped        EventStatus = "stopped"
	EventStatusCompleted      EventStatus = "completed"
	EventStatusEnteredInError EventStatus = "entered-in-error"
	EventStatusUnknown        EventStatus = "unknown"
)

// EventTiming represents codes from http://hl7.org/fhir/ValueSet/event-timing
type EventTiming string

const (
	EventTimingMORN      EventTiming = "MORN"
	EventTimingMORNEarly EventTiming = "MORN.early"
	EventTimingMORNLate  EventTiming = "MORN.late"
	EventTimingNOON      EventTiming = "NOON"
	EventTimingAFT       EventTiming = "AFT"
	EventTimingAFTEarly  EventTiming = "AFT.early"
	EventTimingAFTLate   EventTiming = "AFT.late"
	EventTimingEVE       EventTiming = "EVE"
	EventTimingEVEEarly  EventTiming = "EVE.early"
	EventTimingEVELate   EventTiming = "EVE.late"
	EventTimingNIGHT     EventTiming = "NIGHT"
	EventTimingPHS       EventTiming = "PHS"
	EventTimingIMD       EventTiming = "IMD"
	EventTimingHS        EventTiming = "HS"
	EventTimingWAKE      EventTiming = "WAKE"
	EventTimingC         EventTiming = "C"
	EventTimingCM        EventTiming = "CM"
	EventTimingCD        EventTiming = "CD"
	EventTimingCV        EventTiming = "CV"
	EventTimingAC        EventTiming = "AC"
	EventTimingACM       EventTiming = "ACM"
	EventTimingACD       EventTiming = "ACD"
	EventTimingACV       EventTiming = "ACV"
	EventTimingPC        EventTiming = "PC"
	EventTimingPCM       EventTiming = "PCM"
	EventTimingPCD       EventTiming = "PCD"
	EventTimingPCV       EventTiming = "PCV"
)

// EvidenceVariableRole represents codes from http://hl7.org/fhir/ValueSet/variable-role
type EvidenceVariableRole string

const (
	EvidenceVariableRolePopulation EvidenceVariableRole = "population"
	EvidenceVariableRoleExposure   EvidenceVariableRole = "exposure"
	EvidenceVariableRoleOutcome    EvidenceVariableRole = "outcome"
	EvidenceVariableRoleCovariate  EvidenceVariableRole = "covariate"
)

// ExplanationOfBenefitStatus represents codes from http://hl7.org/fhir/ValueSet/explanationofbenefit-status
type ExplanationOfBenefitStatus string

const (
	ExplanationOfBenefitStatusActive         ExplanationOfBenefitStatus = "active"
	ExplanationOfBenefitStatusCancelled      ExplanationOfBenefitStatus = "cancelled"
	ExplanationOfBenefitStatusDraft          ExplanationOfBenefitStatus = "draft"
	ExplanationOfBenefitStatusEnteredInError ExplanationOfBenefitStatus = "entered-in-error"
)

// ExtensionContextType represents codes from http://hl7.org/fhir/ValueSet/extension-context-type
type ExtensionContextType string

const (
	ExtensionContextTypeFhirpath  ExtensionContextType = "fhirpath"
	ExtensionContextTypeElement   ExtensionContextType = "element"
	ExtensionContextTypeExtension ExtensionContextType = "extension"
)

// FHIRDeviceStatus represents codes from http://hl7.org/fhir/ValueSet/device-status
type FHIRDeviceStatus string

const (
	FHIRDeviceStatusActive         FHIRDeviceStatus = "active"
	FHIRDeviceStatusInactive       FHIRDeviceStatus = "inactive"
	FHIRDeviceStatusEnteredInError FHIRDeviceStatus = "entered-in-error"
	FHIRDeviceStatusUnknown        FHIRDeviceStatus = "unknown"
)

// FHIRSubstanceStatus represents codes from http://hl7.org/fhir/ValueSet/substance-status
type FHIRSubstanceStatus string

const (
	FHIRSubstanceStatusActive         FHIRSubstanceStatus = "active"
	FHIRSubstanceStatusInactive       FHIRSubstanceStatus = "inactive"
	FHIRSubstanceStatusEnteredInError FHIRSubstanceStatus = "entered-in-error"
)

// FHIRTypes represents codes from http://hl7.org/fhir/ValueSet/fhir-types
type FHIRTypes string

const (
	FHIRTypesBase                           FHIRTypes = "Base"
	FHIRTypesElement                        FHIRTypes = "Element"
	FHIRTypesBackboneElement                FHIRTypes = "BackboneElement"
	FHIRTypesDataType                       FHIRTypes = "DataType"
	FHIRTypesAddress                        FHIRTypes = "Address"
	FHIRTypesAnnotation                     FHIRTypes = "Annotation"
	FHIRTypesAttachment                     FHIRTypes = "Attachment"
	FHIRTypesAvailability                   FHIRTypes = "Availability"
	FHIRTypesBackboneType                   FHIRTypes = "BackboneType"
	FHIRTypesDosage                         FHIRTypes = "Dosage"
	FHIRTypesDosageCondition                FHIRTypes = "DosageCondition"
	FHIRTypesDosageDetails                  FHIRTypes = "DosageDetails"
	FHIRTypesDosageSafety                   FHIRTypes = "DosageSafety"
	FHIRTypesElementDefinition              FHIRTypes = "ElementDefinition"
	FHIRTypesMarketingStatus                FHIRTypes = "MarketingStatus"
	FHIRTypesProductShelfLife               FHIRTypes = "ProductShelfLife"
	FHIRTypesRelativeTime                   FHIRTypes = "RelativeTime"
	FHIRTypesTiming                         FHIRTypes = "Timing"
	FHIRTypesCodeableConcept                FHIRTypes = "CodeableConcept"
	FHIRTypesCodeableReference              FHIRTypes = "CodeableReference"
	FHIRTypesCoding                         FHIRTypes = "Coding"
	FHIRTypesContactDetail                  FHIRTypes = "ContactDetail"
	FHIRTypesContactPoint                   FHIRTypes = "ContactPoint"
	FHIRTypesDataRequirement                FHIRTypes = "DataRequirement"
	FHIRTypesExpression                     FHIRTypes = "Expression"
	FHIRTypesExtendedContactDetail          FHIRTypes = "ExtendedContactDetail"
	FHIRTypesExtension                      FHIRTypes = "Extension"
	FHIRTypesHumanName                      FHIRTypes = "HumanName"
	FHIRTypesIdentifier                     FHIRTypes = "Identifier"
	FHIRTypesMeta                           FHIRTypes = "Meta"
	FHIRTypesMonetaryComponent              FHIRTypes = "MonetaryComponent"
	FHIRTypesMoney                          FHIRTypes = "Money"
	FHIRTypesNarrative                      FHIRTypes = "Narrative"
	FHIRTypesParameterDefinition            FHIRTypes = "ParameterDefinition"
	FHIRTypesPeriod                         FHIRTypes = "Period"
	FHIRTypesPrimitiveType                  FHIRTypes = "PrimitiveType"
	FHIRTypesBase64Binary                   FHIRTypes = "base64Binary"
	FHIRTypesBoolean                        FHIRTypes = "boolean"
	FHIRTypesDate                           FHIRTypes = "date"
	FHIRTypesDateTime                       FHIRTypes = "dateTime"
	FHIRTypesDecimal                        FHIRTypes = "decimal"
	FHIRTypesInstant                        FHIRTypes = "instant"
	FHIRTypesInteger                        FHIRTypes = "integer"
	FHIRTypesPositiveInt                    FHIRTypes = "positiveInt"
	FHIRTypesUnsignedInt                    FHIRTypes = "unsignedInt"
	FHIRTypesInteger64                      FHIRTypes = "integer64"
	FHIRTypesString                         FHIRTypes = "string"
	FHIRTypesCode                           FHIRTypes = "code"
	FHIRTypesId                             FHIRTypes = "id"
	FHIRTypesMarkdown                       FHIRTypes = "markdown"
	FHIRTypesTime                           FHIRTypes = "time"
	FHIRTypesUri                            FHIRTypes = "uri"
	FHIRTypesCanonical                      FHIRTypes = "canonical"
	FHIRTypesOid                            FHIRTypes = "oid"
	FHIRTypesUrl                            FHIRTypes = "url"
	FHIRTypesUuid                           FHIRTypes = "uuid"
	FHIRTypesQuantity                       FHIRTypes = "Quantity"
	FHIRTypesAge                            FHIRTypes = "Age"
	FHIRTypesCount                          FHIRTypes = "Count"
	FHIRTypesDistance                       FHIRTypes = "Distance"
	FHIRTypesDuration                       FHIRTypes = "Duration"
	FHIRTypesRange                          FHIRTypes = "Range"
	FHIRTypesRatio                          FHIRTypes = "Ratio"
	FHIRTypesRatioRange                     FHIRTypes = "RatioRange"
	FHIRTypesReference                      FHIRTypes = "Reference"
	FHIRTypesRelatedArtifact                FHIRTypes = "RelatedArtifact"
	FHIRTypesSampledData                    FHIRTypes = "SampledData"
	FHIRTypesSignature                      FHIRTypes = "Signature"
	FHIRTypesTriggerDefinition              FHIRTypes = "TriggerDefinition"
	FHIRTypesUsageContext                   FHIRTypes = "UsageContext"
	FHIRTypesVirtualServiceDetail           FHIRTypes = "VirtualServiceDetail"
	FHIRTypesXhtml                          FHIRTypes = "xhtml"
	FHIRTypesResource                       FHIRTypes = "Resource"
	FHIRTypesBinary                         FHIRTypes = "Binary"
	FHIRTypesBundle                         FHIRTypes = "Bundle"
	FHIRTypesDomainResource                 FHIRTypes = "DomainResource"
	FHIRTypesAccount                        FHIRTypes = "Account"
	FHIRTypesActivityDefinition             FHIRTypes = "ActivityDefinition"
	FHIRTypesActorDefinition                FHIRTypes = "ActorDefinition"
	FHIRTypesAdministrableProductDefinition FHIRTypes = "AdministrableProductDefinition"
	FHIRTypesAdverseEvent                   FHIRTypes = "AdverseEvent"
	FHIRTypesAllergyIntolerance             FHIRTypes = "AllergyIntolerance"
	FHIRTypesAppointment                    FHIRTypes = "Appointment"
	FHIRTypesAppointmentResponse            FHIRTypes = "AppointmentResponse"
	FHIRTypesArtifactAssessment             FHIRTypes = "ArtifactAssessment"
	FHIRTypesAuditEvent                     FHIRTypes = "AuditEvent"
	FHIRTypesBasic                          FHIRTypes = "Basic"
	FHIRTypesBiologicallyDerivedProduct     FHIRTypes = "BiologicallyDerivedProduct"
	FHIRTypesBodyStructure                  FHIRTypes = "BodyStructure"
	FHIRTypesCanonicalResource              FHIRTypes = "CanonicalResource"
	FHIRTypesCapabilityStatement            FHIRTypes = "CapabilityStatement"
	FHIRTypesCarePlan                       FHIRTypes = "CarePlan"
	FHIRTypesCareTeam                       FHIRTypes = "CareTeam"
	FHIRTypesClaim                          FHIRTypes = "Claim"
	FHIRTypesClaimResponse                  FHIRTypes = "ClaimResponse"
	FHIRTypesClinicalUseDefinition          FHIRTypes = "ClinicalUseDefinition"
	FHIRTypesCodeSystem                     FHIRTypes = "CodeSystem"
	FHIRTypesCommunication                  FHIRTypes = "Communication"
	FHIRTypesCommunicationRequest           FHIRTypes = "CommunicationRequest"
	FHIRTypesCompartmentDefinition          FHIRTypes = "CompartmentDefinition"
	FHIRTypesComposition                    FHIRTypes = "Composition"
	FHIRTypesConceptMap                     FHIRTypes = "ConceptMap"
	FHIRTypesCondition                      FHIRTypes = "Condition"
	FHIRTypesConsent                        FHIRTypes = "Consent"
	FHIRTypesContract                       FHIRTypes = "Contract"
	FHIRTypesCoverage                       FHIRTypes = "Coverage"
	FHIRTypesCoverageEligibilityRequest     FHIRTypes = "CoverageEligibilityRequest"
	FHIRTypesCoverageEligibilityResponse    FHIRTypes = "CoverageEligibilityResponse"
	FHIRTypesDetectedIssue                  FHIRTypes = "DetectedIssue"
	FHIRTypesDevice                         FHIRTypes = "Device"
	FHIRTypesDeviceAlert                    FHIRTypes = "DeviceAlert"
	FHIRTypesDeviceAssociation              FHIRTypes = "DeviceAssociation"
	FHIRTypesDeviceDefinition               FHIRTypes = "DeviceDefinition"
	FHIRTypesDeviceMetric                   FHIRTypes = "DeviceMetric"
	FHIRTypesDeviceRequest                  FHIRTypes = "DeviceRequest"
	FHIRTypesDiagnosticReport               FHIRTypes = "DiagnosticReport"
	FHIRTypesDocumentReference              FHIRTypes = "DocumentReference"
	FHIRTypesEncounter                      FHIRTypes = "Encounter"
	FHIRTypesEndpoint                       FHIRTypes = "Endpoint"
	FHIRTypesEnrollmentRequest              FHIRTypes = "EnrollmentRequest"
	FHIRTypesEnrollmentResponse             FHIRTypes = "EnrollmentResponse"
	FHIRTypesEpisodeOfCare                  FHIRTypes = "EpisodeOfCare"
	FHIRTypesEventDefinition                FHIRTypes = "EventDefinition"
	FHIRTypesEvidence                       FHIRTypes = "Evidence"
	FHIRTypesEvidenceVariable               FHIRTypes = "EvidenceVariable"
	FHIRTypesExampleScenario                FHIRTypes = "ExampleScenario"
	FHIRTypesExplanationOfBenefit           FHIRTypes = "ExplanationOfBenefit"
	FHIRTypesFamilyMemberHistory            FHIRTypes = "FamilyMemberHistory"
	FHIRTypesFlag                           FHIRTypes = "Flag"
	FHIRTypesGoal                           FHIRTypes = "Goal"
	FHIRTypesGroup                          FHIRTypes = "Group"
	FHIRTypesGuidanceResponse               FHIRTypes = "GuidanceResponse"
	FHIRTypesHealthcareService              FHIRTypes = "HealthcareService"
	FHIRTypesImagingSelection               FHIRTypes = "ImagingSelection"
	FHIRTypesImagingStudy                   FHIRTypes = "ImagingStudy"
	FHIRTypesImmunization                   FHIRTypes = "Immunization"
	FHIRTypesImplementationGuide            FHIRTypes = "ImplementationGuide"
	FHIRTypesIngredient                     FHIRTypes = "Ingredient"
	FHIRTypesInsurancePlan                  FHIRTypes = "InsurancePlan"
	FHIRTypesInsuranceProduct               FHIRTypes = "InsuranceProduct"
	FHIRTypesInvoice                        FHIRTypes = "Invoice"
	FHIRTypesLibrary                        FHIRTypes = "Library"
	FHIRTypesList                           FHIRTypes = "List"
	FHIRTypesLocation                       FHIRTypes = "Location"
	FHIRTypesManufacturedItemDefinition     FHIRTypes = "ManufacturedItemDefinition"
	FHIRTypesMeasure                        FHIRTypes = "Measure"
	FHIRTypesMeasureReport                  FHIRTypes = "MeasureReport"
	FHIRTypesMedication                     FHIRTypes = "Medication"
	FHIRTypesMedicationAdministration       FHIRTypes = "MedicationAdministration"
	FHIRTypesMedicationDispense             FHIRTypes = "MedicationDispense"
	FHIRTypesMedicationRequest              FHIRTypes = "MedicationRequest"
	FHIRTypesMedicationStatement            FHIRTypes = "MedicationStatement"
	FHIRTypesMedicinalProductDefinition     FHIRTypes = "MedicinalProductDefinition"
	FHIRTypesMessageDefinition              FHIRTypes = "MessageDefinition"
	FHIRTypesMessageHeader                  FHIRTypes = "MessageHeader"
	FHIRTypesMetadataResource               FHIRTypes = "MetadataResource"
	FHIRTypesNamingSystem                   FHIRTypes = "NamingSystem"
	FHIRTypesNutritionIntake                FHIRTypes = "NutritionIntake"
	FHIRTypesNutritionOrder                 FHIRTypes = "NutritionOrder"
	FHIRTypesNutritionProduct               FHIRTypes = "NutritionProduct"
	FHIRTypesObservation                    FHIRTypes = "Observation"
	FHIRTypesObservationDefinition          FHIRTypes = "ObservationDefinition"
	FHIRTypesOperationDefinition            FHIRTypes = "OperationDefinition"
	FHIRTypesOperationOutcome               FHIRTypes = "OperationOutcome"
	FHIRTypesOrganization                   FHIRTypes = "Organization"
	FHIRTypesOrganizationAffiliation        FHIRTypes = "OrganizationAffiliation"
	FHIRTypesPackagedProductDefinition      FHIRTypes = "PackagedProductDefinition"
	FHIRTypesPatient                        FHIRTypes = "Patient"
	FHIRTypesPaymentNotice                  FHIRTypes = "PaymentNotice"
	FHIRTypesPaymentReconciliation          FHIRTypes = "PaymentReconciliation"
	FHIRTypesPerson                         FHIRTypes = "Person"
	FHIRTypesPlanDefinition                 FHIRTypes = "PlanDefinition"
	FHIRTypesPractitioner                   FHIRTypes = "Practitioner"
	FHIRTypesPractitionerRole               FHIRTypes = "PractitionerRole"
	FHIRTypesProcedure                      FHIRTypes = "Procedure"
	FHIRTypesProvenance                     FHIRTypes = "Provenance"
	FHIRTypesQuestionnaire                  FHIRTypes = "Questionnaire"
	FHIRTypesQuestionnaireResponse          FHIRTypes = "QuestionnaireResponse"
	FHIRTypesRegulatedAuthorization         FHIRTypes = "RegulatedAuthorization"
	FHIRTypesRelatedPerson                  FHIRTypes = "RelatedPerson"
	FHIRTypesRequestOrchestration           FHIRTypes = "RequestOrchestration"
	FHIRTypesRequirements                   FHIRTypes = "Requirements"
	FHIRTypesResearchStudy                  FHIRTypes = "ResearchStudy"
	FHIRTypesResearchSubject                FHIRTypes = "ResearchSubject"
	FHIRTypesRiskAssessment                 FHIRTypes = "RiskAssessment"
	FHIRTypesSchedule                       FHIRTypes = "Schedule"
	FHIRTypesSearchParameter                FHIRTypes = "SearchParameter"
	FHIRTypesServiceRequest                 FHIRTypes = "ServiceRequest"
	FHIRTypesSlot                           FHIRTypes = "Slot"
	FHIRTypesSpecimen                       FHIRTypes = "Specimen"
	FHIRTypesSpecimenDefinition             FHIRTypes = "SpecimenDefinition"
	FHIRTypesStructureDefinition            FHIRTypes = "StructureDefinition"
	FHIRTypesStructureMap                   FHIRTypes = "StructureMap"
	FHIRTypesSubscription                   FHIRTypes = "Subscription"
	FHIRTypesSubscriptionStatus             FHIRTypes = "SubscriptionStatus"
	FHIRTypesSubscriptionTopic              FHIRTypes = "SubscriptionTopic"
	FHIRTypesSubstance                      FHIRTypes = "Substance"
	FHIRTypesSubstanceDefinition            FHIRTypes = "SubstanceDefinition"
	FHIRTypesTask                           FHIRTypes = "Task"
	FHIRTypesTerminologyCapabilities        FHIRTypes = "TerminologyCapabilities"
	FHIRTypesValueSet                       FHIRTypes = "ValueSet"
	FHIRTypesVisionPrescription             FHIRTypes = "VisionPrescription"
	FHIRTypesParameters                     FHIRTypes = "Parameters"
)

// FHIRVersion represents codes from http://hl7.org/fhir/ValueSet/FHIR-version
type FHIRVersion string

const (
	FHIRVersionCode001           FHIRVersion = "0.01"
	FHIRVersionCode005           FHIRVersion = "0.05"
	FHIRVersionCode006           FHIRVersion = "0.06"
	FHIRVersionCode011           FHIRVersion = "0.11"
	FHIRVersionCode00            FHIRVersion = "0.0"
	FHIRVersionCode0080          FHIRVersion = "0.0.80"
	FHIRVersionCode0081          FHIRVersion = "0.0.81"
	FHIRVersionCode0082          FHIRVersion = "0.0.82"
	FHIRVersionCode04            FHIRVersion = "0.4"
	FHIRVersionCode040           FHIRVersion = "0.4.0"
	FHIRVersionCode05            FHIRVersion = "0.5"
	FHIRVersionCode050           FHIRVersion = "0.5.0"
	FHIRVersionCode10            FHIRVersion = "1.0"
	FHIRVersionCode100           FHIRVersion = "1.0.0"
	FHIRVersionCode101           FHIRVersion = "1.0.1"
	FHIRVersionCode102           FHIRVersion = "1.0.2"
	FHIRVersionCode11            FHIRVersion = "1.1"
	FHIRVersionCode110           FHIRVersion = "1.1.0"
	FHIRVersionCode14            FHIRVersion = "1.4"
	FHIRVersionCode140           FHIRVersion = "1.4.0"
	FHIRVersionCode16            FHIRVersion = "1.6"
	FHIRVersionCode160           FHIRVersion = "1.6.0"
	FHIRVersionCode18            FHIRVersion = "1.8"
	FHIRVersionCode180           FHIRVersion = "1.8.0"
	FHIRVersionCode30            FHIRVersion = "3.0"
	FHIRVersionCode300           FHIRVersion = "3.0.0"
	FHIRVersionCode301           FHIRVersion = "3.0.1"
	FHIRVersionCode302           FHIRVersion = "3.0.2"
	FHIRVersionCode33            FHIRVersion = "3.3"
	FHIRVersionCode330           FHIRVersion = "3.3.0"
	FHIRVersionCode35            FHIRVersion = "3.5"
	FHIRVersionCode350           FHIRVersion = "3.5.0"
	FHIRVersionCode40            FHIRVersion = "4.0"
	FHIRVersionCode400           FHIRVersion = "4.0.0"
	FHIRVersionCode401           FHIRVersion = "4.0.1"
	FHIRVersionCode41            FHIRVersion = "4.1"
	FHIRVersionCode410           FHIRVersion = "4.1.0"
	FHIRVersionCode42            FHIRVersion = "4.2"
	FHIRVersionCode420           FHIRVersion = "4.2.0"
	FHIRVersionCode43            FHIRVersion = "4.3"
	FHIRVersionCode430           FHIRVersion = "4.3.0"
	FHIRVersionCode430Cibuild    FHIRVersion = "4.3.0-cibuild"
	FHIRVersionCode430Snapshot1  FHIRVersion = "4.3.0-snapshot1"
	FHIRVersionCode44            FHIRVersion = "4.4"
	FHIRVersionCode440           FHIRVersion = "4.4.0"
	FHIRVersionCode45            FHIRVersion = "4.5"
	FHIRVersionCode450           FHIRVersion = "4.5.0"
	FHIRVersionCode46            FHIRVersion = "4.6"
	FHIRVersionCode460           FHIRVersion = "4.6.0"
	FHIRVersionCode50            FHIRVersion = "5.0"
	FHIRVersionCode500           FHIRVersion = "5.0.0"
	FHIRVersionCode500Cibuild    FHIRVersion = "5.0.0-cibuild"
	FHIRVersionCode500Snapshot1  FHIRVersion = "5.0.0-snapshot1"
	FHIRVersionCode500Snapshot2  FHIRVersion = "5.0.0-snapshot2"
	FHIRVersionCode500Ballot     FHIRVersion = "5.0.0-ballot"
	FHIRVersionCode500Snapshot3  FHIRVersion = "5.0.0-snapshot3"
	FHIRVersionCode500DraftFinal FHIRVersion = "5.0.0-draft-final"
	FHIRVersionCode60            FHIRVersion = "6.0"
	FHIRVersionCode600           FHIRVersion = "6.0.0"
	FHIRVersionCode600Ballo1     FHIRVersion = "6.0.0-ballo1"
	FHIRVersionCode600Ballot2    FHIRVersion = "6.0.0-ballot2"
	FHIRVersionCode600Ballot3    FHIRVersion = "6.0.0-ballot3"
)

// FamilyHistoryStatus represents codes from http://hl7.org/fhir/ValueSet/history-status
type FamilyHistoryStatus string

const (
	FamilyHistoryStatusPartial        FamilyHistoryStatus = "partial"
	FamilyHistoryStatusCompleted      FamilyHistoryStatus = "completed"
	FamilyHistoryStatusEnteredInError FamilyHistoryStatus = "entered-in-error"
	FamilyHistoryStatusHealthUnknown  FamilyHistoryStatus = "health-unknown"
)

// FilterOperator represents codes from http://hl7.org/fhir/ValueSet/filter-operator
type FilterOperator string

const (
	FilterOperatorCode           FilterOperator = "="
	FilterOperatorIsA            FilterOperator = "is-a"
	FilterOperatorDescendentOf   FilterOperator = "descendent-of"
	FilterOperatorIsNotA         FilterOperator = "is-not-a"
	FilterOperatorRegex          FilterOperator = "regex"
	FilterOperatorIn             FilterOperator = "in"
	FilterOperatorNotIn          FilterOperator = "not-in"
	FilterOperatorGeneralizes    FilterOperator = "generalizes"
	FilterOperatorChildOf        FilterOperator = "child-of"
	FilterOperatorDescendentLeaf FilterOperator = "descendent-leaf"
	FilterOperatorExists         FilterOperator = "exists"
)

// FinancialResourceStatusCodes represents codes from http://hl7.org/fhir/ValueSet/fm-status
type FinancialResourceStatusCodes string

const (
	FinancialResourceStatusCodesActive         FinancialResourceStatusCodes = "active"
	FinancialResourceStatusCodesCancelled      FinancialResourceStatusCodes = "cancelled"
	FinancialResourceStatusCodesDraft          FinancialResourceStatusCodes = "draft"
	FinancialResourceStatusCodesEnteredInError FinancialResourceStatusCodes = "entered-in-error"
)

// FlagStatus represents codes from http://hl7.org/fhir/ValueSet/flag-status
type FlagStatus string

const (
	FlagStatusActive         FlagStatus = "active"
	FlagStatusInactive       FlagStatus = "inactive"
	FlagStatusEnteredInError FlagStatus = "entered-in-error"
)
