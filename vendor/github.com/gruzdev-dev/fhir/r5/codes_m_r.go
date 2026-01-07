package models

// MeasureReportStatus represents codes from http://hl7.org/fhir/ValueSet/measure-report-status
type MeasureReportStatus string

const (
	MeasureReportStatusComplete MeasureReportStatus = "complete"
	MeasureReportStatusPending  MeasureReportStatus = "pending"
	MeasureReportStatusError    MeasureReportStatus = "error"
)

// MeasureReportType represents codes from http://hl7.org/fhir/ValueSet/measure-report-type
type MeasureReportType string

const (
	MeasureReportTypeIndividual   MeasureReportType = "individual"
	MeasureReportTypeSubjectList  MeasureReportType = "subject-list"
	MeasureReportTypeSummary      MeasureReportType = "summary"
	MeasureReportTypeDataExchange MeasureReportType = "data-exchange"
)

// MedicationAdministrationStatusCodes represents codes from http://hl7.org/fhir/ValueSet/medication-admin-status
type MedicationAdministrationStatusCodes string

const (
	MedicationAdministrationStatusCodesInProgress     MedicationAdministrationStatusCodes = "in-progress"
	MedicationAdministrationStatusCodesNotDone        MedicationAdministrationStatusCodes = "not-done"
	MedicationAdministrationStatusCodesOnHold         MedicationAdministrationStatusCodes = "on-hold"
	MedicationAdministrationStatusCodesCompleted      MedicationAdministrationStatusCodes = "completed"
	MedicationAdministrationStatusCodesEnteredInError MedicationAdministrationStatusCodes = "entered-in-error"
	MedicationAdministrationStatusCodesStopped        MedicationAdministrationStatusCodes = "stopped"
	MedicationAdministrationStatusCodesUnknown        MedicationAdministrationStatusCodes = "unknown"
)

// MedicationDispenseStatusCodes represents codes from http://hl7.org/fhir/ValueSet/medicationdispense-status
type MedicationDispenseStatusCodes string

const (
	MedicationDispenseStatusCodesPreparation    MedicationDispenseStatusCodes = "preparation"
	MedicationDispenseStatusCodesInProgress     MedicationDispenseStatusCodes = "in-progress"
	MedicationDispenseStatusCodesCancelled      MedicationDispenseStatusCodes = "cancelled"
	MedicationDispenseStatusCodesOnHold         MedicationDispenseStatusCodes = "on-hold"
	MedicationDispenseStatusCodesCompleted      MedicationDispenseStatusCodes = "completed"
	MedicationDispenseStatusCodesEnteredInError MedicationDispenseStatusCodes = "entered-in-error"
	MedicationDispenseStatusCodesUnfulfilled    MedicationDispenseStatusCodes = "unfulfilled"
	MedicationDispenseStatusCodesDeclined       MedicationDispenseStatusCodes = "declined"
	MedicationDispenseStatusCodesUnknown        MedicationDispenseStatusCodes = "unknown"
)

// MedicationRequestIntent represents codes from http://hl7.org/fhir/ValueSet/medicationrequest-intent
type MedicationRequestIntent string

const (
	MedicationRequestIntentProposal      MedicationRequestIntent = "proposal"
	MedicationRequestIntentPlan          MedicationRequestIntent = "plan"
	MedicationRequestIntentOrder         MedicationRequestIntent = "order"
	MedicationRequestIntentOriginalOrder MedicationRequestIntent = "original-order"
	MedicationRequestIntentReflexOrder   MedicationRequestIntent = "reflex-order"
	MedicationRequestIntentFillerOrder   MedicationRequestIntent = "filler-order"
	MedicationRequestIntentInstanceOrder MedicationRequestIntent = "instance-order"
	MedicationRequestIntentOption        MedicationRequestIntent = "option"
)

// MedicationStatementStatusCodes represents codes from http://hl7.org/fhir/ValueSet/medication-statement-status
type MedicationStatementStatusCodes string

const (
	MedicationStatementStatusCodesRecorded       MedicationStatementStatusCodes = "recorded"
	MedicationStatementStatusCodesEnteredInError MedicationStatementStatusCodes = "entered-in-error"
	MedicationStatementStatusCodesDraft          MedicationStatementStatusCodes = "draft"
)

// MedicationStatusCodes represents codes from http://hl7.org/fhir/ValueSet/medication-status
type MedicationStatusCodes string

const (
	MedicationStatusCodesActive         MedicationStatusCodes = "active"
	MedicationStatusCodesInactive       MedicationStatusCodes = "inactive"
	MedicationStatusCodesEnteredInError MedicationStatusCodes = "entered-in-error"
)

// MedicationrequestStatus represents codes from http://hl7.org/fhir/ValueSet/medicationrequest-status
type MedicationrequestStatus string

const (
	MedicationrequestStatusActive         MedicationrequestStatus = "active"
	MedicationrequestStatusOnHold         MedicationrequestStatus = "on-hold"
	MedicationrequestStatusEnded          MedicationrequestStatus = "ended"
	MedicationrequestStatusStopped        MedicationrequestStatus = "stopped"
	MedicationrequestStatusCompleted      MedicationrequestStatus = "completed"
	MedicationrequestStatusCancelled      MedicationrequestStatus = "cancelled"
	MedicationrequestStatusEnteredInError MedicationrequestStatus = "entered-in-error"
	MedicationrequestStatusDraft          MedicationrequestStatus = "draft"
	MedicationrequestStatusUnknown        MedicationrequestStatus = "unknown"
)

// MessageSignificanceCategory represents codes from http://hl7.org/fhir/ValueSet/message-significance-category
type MessageSignificanceCategory string

const (
	MessageSignificanceCategoryConsequence  MessageSignificanceCategory = "consequence"
	MessageSignificanceCategoryCurrency     MessageSignificanceCategory = "currency"
	MessageSignificanceCategoryNotification MessageSignificanceCategory = "notification"
)

// MessageheaderResponseRequest represents codes from http://hl7.org/fhir/ValueSet/messageheader-response-request
type MessageheaderResponseRequest string

const (
	MessageheaderResponseRequestAlways    MessageheaderResponseRequest = "always"
	MessageheaderResponseRequestOnError   MessageheaderResponseRequest = "on-error"
	MessageheaderResponseRequestNever     MessageheaderResponseRequest = "never"
	MessageheaderResponseRequestOnSuccess MessageheaderResponseRequest = "on-success"
)

// NameUse represents codes from http://hl7.org/fhir/ValueSet/name-use
type NameUse string

const (
	NameUseUsual     NameUse = "usual"
	NameUseOfficial  NameUse = "official"
	NameUseTemp      NameUse = "temp"
	NameUseNickname  NameUse = "nickname"
	NameUseAnonymous NameUse = "anonymous"
	NameUseOld       NameUse = "old"
	NameUseMaiden    NameUse = "maiden"
)

// NamingSystemIdentifierType represents codes from http://hl7.org/fhir/ValueSet/namingsystem-identifier-type
type NamingSystemIdentifierType string

const (
	NamingSystemIdentifierTypeOid          NamingSystemIdentifierType = "oid"
	NamingSystemIdentifierTypeUuid         NamingSystemIdentifierType = "uuid"
	NamingSystemIdentifierTypeUri          NamingSystemIdentifierType = "uri"
	NamingSystemIdentifierTypeIriStem      NamingSystemIdentifierType = "iri-stem"
	NamingSystemIdentifierTypeV2csmnemonic NamingSystemIdentifierType = "v2csmnemonic"
	NamingSystemIdentifierTypeOther        NamingSystemIdentifierType = "other"
)

// NamingSystemType represents codes from http://hl7.org/fhir/ValueSet/namingsystem-type
type NamingSystemType string

const (
	NamingSystemTypeCodesystem NamingSystemType = "codesystem"
	NamingSystemTypeIdentifier NamingSystemType = "identifier"
	NamingSystemTypeRoot       NamingSystemType = "root"
)

// NarrativeStatus represents codes from http://hl7.org/fhir/ValueSet/narrative-status
type NarrativeStatus string

const (
	NarrativeStatusGenerated  NarrativeStatus = "generated"
	NarrativeStatusExtensions NarrativeStatus = "extensions"
	NarrativeStatusAdditional NarrativeStatus = "additional"
	NarrativeStatusEmpty      NarrativeStatus = "empty"
)

// NoteType represents codes from http://hl7.org/fhir/ValueSet/note-type
type NoteType string

const (
	NoteTypeDisplay   NoteType = "display"
	NoteTypePrint     NoteType = "print"
	NoteTypePrintoper NoteType = "printoper"
)

// NutritionProductStatus represents codes from http://hl7.org/fhir/ValueSet/nutritionproduct-status
type NutritionProductStatus string

const (
	NutritionProductStatusActive         NutritionProductStatus = "active"
	NutritionProductStatusInactive       NutritionProductStatus = "inactive"
	NutritionProductStatusEnteredInError NutritionProductStatus = "entered-in-error"
)

// ObservationDataType represents codes from http://hl7.org/fhir/ValueSet/permitted-data-type
type ObservationDataType string

const (
	ObservationDataTypeQuantity        ObservationDataType = "Quantity"
	ObservationDataTypeCodeableConcept ObservationDataType = "CodeableConcept"
	ObservationDataTypeString          ObservationDataType = "string"
	ObservationDataTypeBoolean         ObservationDataType = "boolean"
	ObservationDataTypeInteger         ObservationDataType = "integer"
	ObservationDataTypeRange           ObservationDataType = "Range"
	ObservationDataTypeRatio           ObservationDataType = "Ratio"
	ObservationDataTypeSampledData     ObservationDataType = "SampledData"
	ObservationDataTypeTime            ObservationDataType = "time"
	ObservationDataTypeDateTime        ObservationDataType = "dateTime"
	ObservationDataTypePeriod          ObservationDataType = "Period"
)

// ObservationRangeCategory represents codes from http://hl7.org/fhir/ValueSet/observation-range-category
type ObservationRangeCategory string

const (
	ObservationRangeCategoryReference ObservationRangeCategory = "reference"
	ObservationRangeCategoryCritical  ObservationRangeCategory = "critical"
	ObservationRangeCategoryAbsolute  ObservationRangeCategory = "absolute"
)

// ObservationStatus represents codes from http://hl7.org/fhir/ValueSet/observation-status
type ObservationStatus string

const (
	ObservationStatusRegistered        ObservationStatus = "registered"
	ObservationStatusSpecimenInProcess ObservationStatus = "specimen-in-process"
	ObservationStatusPreliminary       ObservationStatus = "preliminary"
	ObservationStatusFinal             ObservationStatus = "final"
	ObservationStatusAmended           ObservationStatus = "amended"
	ObservationStatusCorrected         ObservationStatus = "corrected"
	ObservationStatusAppended          ObservationStatus = "appended"
	ObservationStatusCancelled         ObservationStatus = "cancelled"
	ObservationStatusEnteredInError    ObservationStatus = "entered-in-error"
	ObservationStatusUnknown           ObservationStatus = "unknown"
	ObservationStatusCannotBeObtained  ObservationStatus = "cannot-be-obtained"
)

// OperationKind represents codes from http://hl7.org/fhir/ValueSet/operation-kind
type OperationKind string

const (
	OperationKindOperation OperationKind = "operation"
	OperationKindQuery     OperationKind = "query"
)

// OperationParameterScope represents codes from http://hl7.org/fhir/ValueSet/operation-parameter-scope
type OperationParameterScope string

const (
	OperationParameterScopeInstance OperationParameterScope = "instance"
	OperationParameterScopeType     OperationParameterScope = "type"
	OperationParameterScopeSystem   OperationParameterScope = "system"
)

// OperationParameterUse represents codes from http://hl7.org/fhir/ValueSet/operation-parameter-use
type OperationParameterUse string

const (
	OperationParameterUseIn  OperationParameterUse = "in"
	OperationParameterUseOut OperationParameterUse = "out"
)

// OperationSynchronicityControl represents codes from http://hl7.org/fhir/ValueSet/synchronicity-control
type OperationSynchronicityControl string

const (
	OperationSynchronicityControlSynchronous  OperationSynchronicityControl = "synchronous"
	OperationSynchronicityControlAsynchronous OperationSynchronicityControl = "asynchronous"
	OperationSynchronicityControlEither       OperationSynchronicityControl = "either"
)

// ParticipationStatus represents codes from http://hl7.org/fhir/ValueSet/participationstatus
type ParticipationStatus string

const (
	ParticipationStatusAccepted    ParticipationStatus = "accepted"
	ParticipationStatusDeclined    ParticipationStatus = "declined"
	ParticipationStatusTentative   ParticipationStatus = "tentative"
	ParticipationStatusNeedsAction ParticipationStatus = "needs-action"
)

// PatchMimeTypes represents codes from http://hl7.org/fhir/ValueSet/patchmimetypes
type PatchMimeTypes string

const (
	PatchMimeTypesApplicationfhirxml       PatchMimeTypes = "application/fhir+xml"
	PatchMimeTypesApplicationfhirjson      PatchMimeTypes = "application/fhir+json"
	PatchMimeTypesApplicationfhirturtle    PatchMimeTypes = "application/fhir+turtle"
	PatchMimeTypesApplicationjsonPatchjson PatchMimeTypes = "application/json-patch+json"
	PatchMimeTypesApplicationxmlPatchxml   PatchMimeTypes = "application/xml-patch+xml"
)

// PaymentOutcome represents codes from http://hl7.org/fhir/ValueSet/payment-outcome
type PaymentOutcome string

const (
	PaymentOutcomeQueued   PaymentOutcome = "queued"
	PaymentOutcomeComplete PaymentOutcome = "complete"
	PaymentOutcomeError    PaymentOutcome = "error"
	PaymentOutcomePartial  PaymentOutcome = "partial"
)

// PriceComponentType represents codes from http://hl7.org/fhir/ValueSet/price-component-type
type PriceComponentType string

const (
	PriceComponentTypeBase          PriceComponentType = "base"
	PriceComponentTypeSurcharge     PriceComponentType = "surcharge"
	PriceComponentTypeDiscount      PriceComponentType = "discount"
	PriceComponentTypeTax           PriceComponentType = "tax"
	PriceComponentTypeInformational PriceComponentType = "informational"
)

// PropertyRepresentation represents codes from http://hl7.org/fhir/ValueSet/property-representation
type PropertyRepresentation string

const (
	PropertyRepresentationXmlAttr  PropertyRepresentation = "xmlAttr"
	PropertyRepresentationXmlText  PropertyRepresentation = "xmlText"
	PropertyRepresentationTypeAttr PropertyRepresentation = "typeAttr"
	PropertyRepresentationCdaText  PropertyRepresentation = "cdaText"
	PropertyRepresentationXhtml    PropertyRepresentation = "xhtml"
)

// PropertyType represents codes from http://hl7.org/fhir/ValueSet/concept-property-type
type PropertyType string

const (
	PropertyTypeCode     PropertyType = "code"
	PropertyTypeCoding   PropertyType = "Coding"
	PropertyTypeString   PropertyType = "string"
	PropertyTypeInteger  PropertyType = "integer"
	PropertyTypeBoolean  PropertyType = "boolean"
	PropertyTypeDateTime PropertyType = "dateTime"
	PropertyTypeDecimal  PropertyType = "decimal"
)

// ProvenanceEntityRole represents codes from http://hl7.org/fhir/ValueSet/provenance-entity-role
type ProvenanceEntityRole string

const (
	ProvenanceEntityRoleRevision     ProvenanceEntityRole = "revision"
	ProvenanceEntityRoleQuotation    ProvenanceEntityRole = "quotation"
	ProvenanceEntityRoleSource       ProvenanceEntityRole = "source"
	ProvenanceEntityRoleInstantiates ProvenanceEntityRole = "instantiates"
	ProvenanceEntityRoleRemoval      ProvenanceEntityRole = "removal"
)

// PublicationStatus represents codes from http://hl7.org/fhir/ValueSet/publication-status
type PublicationStatus string

const (
	PublicationStatusDraft   PublicationStatus = "draft"
	PublicationStatusActive  PublicationStatus = "active"
	PublicationStatusRetired PublicationStatus = "retired"
	PublicationStatusUnknown PublicationStatus = "unknown"
)

// QuantityComparator represents codes from http://hl7.org/fhir/ValueSet/quantity-comparator
type QuantityComparator string

const (
	QuantityComparatorCode   QuantityComparator = "<"
	QuantityComparatorCode_2 QuantityComparator = "<="
	QuantityComparatorCode_3 QuantityComparator = ">="
	QuantityComparatorCode_4 QuantityComparator = ">"
	QuantityComparatorAd     QuantityComparator = "ad"
)

// QuestionnaireAnswerConstraint represents codes from http://hl7.org/fhir/ValueSet/questionnaire-answer-constraint
type QuestionnaireAnswerConstraint string

const (
	QuestionnaireAnswerConstraintOptionsOnly     QuestionnaireAnswerConstraint = "optionsOnly"
	QuestionnaireAnswerConstraintOptionsOrType   QuestionnaireAnswerConstraint = "optionsOrType"
	QuestionnaireAnswerConstraintOptionsOrString QuestionnaireAnswerConstraint = "optionsOrString"
)

// QuestionnaireItemDisabledDisplay represents codes from http://hl7.org/fhir/ValueSet/questionnaire-disabled-display
type QuestionnaireItemDisabledDisplay string

const (
	QuestionnaireItemDisabledDisplayHidden    QuestionnaireItemDisabledDisplay = "hidden"
	QuestionnaireItemDisabledDisplayProtected QuestionnaireItemDisabledDisplay = "protected"
)

// QuestionnaireItemOperator represents codes from http://hl7.org/fhir/ValueSet/questionnaire-enable-operator
type QuestionnaireItemOperator string

const (
	QuestionnaireItemOperatorExists QuestionnaireItemOperator = "exists"
	QuestionnaireItemOperatorCode   QuestionnaireItemOperator = "="
	QuestionnaireItemOperatorCode_2 QuestionnaireItemOperator = "!="
	QuestionnaireItemOperatorCode_3 QuestionnaireItemOperator = ">"
	QuestionnaireItemOperatorCode_4 QuestionnaireItemOperator = "<"
	QuestionnaireItemOperatorCode_5 QuestionnaireItemOperator = ">="
	QuestionnaireItemOperatorCode_6 QuestionnaireItemOperator = "<="
)

// QuestionnaireItemTypeUsable represents codes from http://hl7.org/fhir/ValueSet/item-type-useable
type QuestionnaireItemTypeUsable string

const (
	QuestionnaireItemTypeUsableGroup      QuestionnaireItemTypeUsable = "group"
	QuestionnaireItemTypeUsableDisplay    QuestionnaireItemTypeUsable = "display"
	QuestionnaireItemTypeUsableQuestion   QuestionnaireItemTypeUsable = "question"
	QuestionnaireItemTypeUsableBoolean    QuestionnaireItemTypeUsable = "boolean"
	QuestionnaireItemTypeUsableDecimal    QuestionnaireItemTypeUsable = "decimal"
	QuestionnaireItemTypeUsableInteger    QuestionnaireItemTypeUsable = "integer"
	QuestionnaireItemTypeUsableDate       QuestionnaireItemTypeUsable = "date"
	QuestionnaireItemTypeUsableDateTime   QuestionnaireItemTypeUsable = "dateTime"
	QuestionnaireItemTypeUsableTime       QuestionnaireItemTypeUsable = "time"
	QuestionnaireItemTypeUsableString     QuestionnaireItemTypeUsable = "string"
	QuestionnaireItemTypeUsableText       QuestionnaireItemTypeUsable = "text"
	QuestionnaireItemTypeUsableUrl        QuestionnaireItemTypeUsable = "url"
	QuestionnaireItemTypeUsableCoding     QuestionnaireItemTypeUsable = "coding"
	QuestionnaireItemTypeUsableAttachment QuestionnaireItemTypeUsable = "attachment"
	QuestionnaireItemTypeUsableReference  QuestionnaireItemTypeUsable = "reference"
	QuestionnaireItemTypeUsableQuantity   QuestionnaireItemTypeUsable = "quantity"
)

// QuestionnaireResponseStatus represents codes from http://hl7.org/fhir/ValueSet/questionnaire-answers-status
type QuestionnaireResponseStatus string

const (
	QuestionnaireResponseStatusInProgress     QuestionnaireResponseStatus = "in-progress"
	QuestionnaireResponseStatusCompleted      QuestionnaireResponseStatus = "completed"
	QuestionnaireResponseStatusAmended        QuestionnaireResponseStatus = "amended"
	QuestionnaireResponseStatusEnteredInError QuestionnaireResponseStatus = "entered-in-error"
	QuestionnaireResponseStatusStopped        QuestionnaireResponseStatus = "stopped"
)

// ReferenceHandlingPolicy represents codes from http://hl7.org/fhir/ValueSet/reference-handling-policy
type ReferenceHandlingPolicy string

const (
	ReferenceHandlingPolicyLiteral  ReferenceHandlingPolicy = "literal"
	ReferenceHandlingPolicyLogical  ReferenceHandlingPolicy = "logical"
	ReferenceHandlingPolicyResolves ReferenceHandlingPolicy = "resolves"
	ReferenceHandlingPolicyEnforced ReferenceHandlingPolicy = "enforced"
	ReferenceHandlingPolicyLocal    ReferenceHandlingPolicy = "local"
)

// ReferenceVersionRules represents codes from http://hl7.org/fhir/ValueSet/reference-version-rules
type ReferenceVersionRules string

const (
	ReferenceVersionRulesEither      ReferenceVersionRules = "either"
	ReferenceVersionRulesIndependent ReferenceVersionRules = "independent"
	ReferenceVersionRulesSpecific    ReferenceVersionRules = "specific"
)

// RelatedArtifactType represents codes from http://hl7.org/fhir/ValueSet/related-artifact-type
type RelatedArtifactType string

const (
	RelatedArtifactTypeDocumentation RelatedArtifactType = "documentation"
	RelatedArtifactTypeJustification RelatedArtifactType = "justification"
	RelatedArtifactTypeCitation      RelatedArtifactType = "citation"
	RelatedArtifactTypePredecessor   RelatedArtifactType = "predecessor"
	RelatedArtifactTypeSuccessor     RelatedArtifactType = "successor"
	RelatedArtifactTypeDerivedFrom   RelatedArtifactType = "derived-from"
	RelatedArtifactTypeDependsOn     RelatedArtifactType = "depends-on"
	RelatedArtifactTypeComposedOf    RelatedArtifactType = "composed-of"
	RelatedArtifactTypePartOf        RelatedArtifactType = "part-of"
)

// RequestIntent represents codes from http://hl7.org/fhir/ValueSet/request-intent
type RequestIntent string

const (
	RequestIntentProposal      RequestIntent = "proposal"
	RequestIntentSolicitOffer  RequestIntent = "solicit-offer"
	RequestIntentOfferResponse RequestIntent = "offer-response"
	RequestIntentPlan          RequestIntent = "plan"
	RequestIntentDirective     RequestIntent = "directive"
	RequestIntentOrder         RequestIntent = "order"
	RequestIntentOriginalOrder RequestIntent = "original-order"
	RequestIntentReflexOrder   RequestIntent = "reflex-order"
	RequestIntentFillerOrder   RequestIntent = "filler-order"
	RequestIntentInstanceOrder RequestIntent = "instance-order"
	RequestIntentOption        RequestIntent = "option"
)

// RequestPriority represents codes from http://hl7.org/fhir/ValueSet/request-priority
type RequestPriority string

const (
	RequestPriorityRoutine RequestPriority = "routine"
	RequestPriorityUrgent  RequestPriority = "urgent"
	RequestPriorityAsap    RequestPriority = "asap"
	RequestPriorityStat    RequestPriority = "stat"
)

// RequestResourceTypes represents codes from http://hl7.org/fhir/ValueSet/request-resource-types
type RequestResourceTypes string

const (
	RequestResourceTypesBase                           RequestResourceTypes = "Base"
	RequestResourceTypesElement                        RequestResourceTypes = "Element"
	RequestResourceTypesBackboneElement                RequestResourceTypes = "BackboneElement"
	RequestResourceTypesDataType                       RequestResourceTypes = "DataType"
	RequestResourceTypesAddress                        RequestResourceTypes = "Address"
	RequestResourceTypesAnnotation                     RequestResourceTypes = "Annotation"
	RequestResourceTypesAttachment                     RequestResourceTypes = "Attachment"
	RequestResourceTypesAvailability                   RequestResourceTypes = "Availability"
	RequestResourceTypesBackboneType                   RequestResourceTypes = "BackboneType"
	RequestResourceTypesDosage                         RequestResourceTypes = "Dosage"
	RequestResourceTypesDosageCondition                RequestResourceTypes = "DosageCondition"
	RequestResourceTypesDosageDetails                  RequestResourceTypes = "DosageDetails"
	RequestResourceTypesDosageSafety                   RequestResourceTypes = "DosageSafety"
	RequestResourceTypesElementDefinition              RequestResourceTypes = "ElementDefinition"
	RequestResourceTypesMarketingStatus                RequestResourceTypes = "MarketingStatus"
	RequestResourceTypesProductShelfLife               RequestResourceTypes = "ProductShelfLife"
	RequestResourceTypesRelativeTime                   RequestResourceTypes = "RelativeTime"
	RequestResourceTypesTiming                         RequestResourceTypes = "Timing"
	RequestResourceTypesCodeableConcept                RequestResourceTypes = "CodeableConcept"
	RequestResourceTypesCodeableReference              RequestResourceTypes = "CodeableReference"
	RequestResourceTypesCoding                         RequestResourceTypes = "Coding"
	RequestResourceTypesContactDetail                  RequestResourceTypes = "ContactDetail"
	RequestResourceTypesContactPoint                   RequestResourceTypes = "ContactPoint"
	RequestResourceTypesDataRequirement                RequestResourceTypes = "DataRequirement"
	RequestResourceTypesExpression                     RequestResourceTypes = "Expression"
	RequestResourceTypesExtendedContactDetail          RequestResourceTypes = "ExtendedContactDetail"
	RequestResourceTypesExtension                      RequestResourceTypes = "Extension"
	RequestResourceTypesHumanName                      RequestResourceTypes = "HumanName"
	RequestResourceTypesIdentifier                     RequestResourceTypes = "Identifier"
	RequestResourceTypesMeta                           RequestResourceTypes = "Meta"
	RequestResourceTypesMonetaryComponent              RequestResourceTypes = "MonetaryComponent"
	RequestResourceTypesMoney                          RequestResourceTypes = "Money"
	RequestResourceTypesNarrative                      RequestResourceTypes = "Narrative"
	RequestResourceTypesParameterDefinition            RequestResourceTypes = "ParameterDefinition"
	RequestResourceTypesPeriod                         RequestResourceTypes = "Period"
	RequestResourceTypesPrimitiveType                  RequestResourceTypes = "PrimitiveType"
	RequestResourceTypesBase64Binary                   RequestResourceTypes = "base64Binary"
	RequestResourceTypesBoolean                        RequestResourceTypes = "boolean"
	RequestResourceTypesDate                           RequestResourceTypes = "date"
	RequestResourceTypesDateTime                       RequestResourceTypes = "dateTime"
	RequestResourceTypesDecimal                        RequestResourceTypes = "decimal"
	RequestResourceTypesInstant                        RequestResourceTypes = "instant"
	RequestResourceTypesInteger                        RequestResourceTypes = "integer"
	RequestResourceTypesPositiveInt                    RequestResourceTypes = "positiveInt"
	RequestResourceTypesUnsignedInt                    RequestResourceTypes = "unsignedInt"
	RequestResourceTypesInteger64                      RequestResourceTypes = "integer64"
	RequestResourceTypesString                         RequestResourceTypes = "string"
	RequestResourceTypesCode                           RequestResourceTypes = "code"
	RequestResourceTypesId                             RequestResourceTypes = "id"
	RequestResourceTypesMarkdown                       RequestResourceTypes = "markdown"
	RequestResourceTypesTime                           RequestResourceTypes = "time"
	RequestResourceTypesUri                            RequestResourceTypes = "uri"
	RequestResourceTypesCanonical                      RequestResourceTypes = "canonical"
	RequestResourceTypesOid                            RequestResourceTypes = "oid"
	RequestResourceTypesUrl                            RequestResourceTypes = "url"
	RequestResourceTypesUuid                           RequestResourceTypes = "uuid"
	RequestResourceTypesQuantity                       RequestResourceTypes = "Quantity"
	RequestResourceTypesAge                            RequestResourceTypes = "Age"
	RequestResourceTypesCount                          RequestResourceTypes = "Count"
	RequestResourceTypesDistance                       RequestResourceTypes = "Distance"
	RequestResourceTypesDuration                       RequestResourceTypes = "Duration"
	RequestResourceTypesRange                          RequestResourceTypes = "Range"
	RequestResourceTypesRatio                          RequestResourceTypes = "Ratio"
	RequestResourceTypesRatioRange                     RequestResourceTypes = "RatioRange"
	RequestResourceTypesReference                      RequestResourceTypes = "Reference"
	RequestResourceTypesRelatedArtifact                RequestResourceTypes = "RelatedArtifact"
	RequestResourceTypesSampledData                    RequestResourceTypes = "SampledData"
	RequestResourceTypesSignature                      RequestResourceTypes = "Signature"
	RequestResourceTypesTriggerDefinition              RequestResourceTypes = "TriggerDefinition"
	RequestResourceTypesUsageContext                   RequestResourceTypes = "UsageContext"
	RequestResourceTypesVirtualServiceDetail           RequestResourceTypes = "VirtualServiceDetail"
	RequestResourceTypesXhtml                          RequestResourceTypes = "xhtml"
	RequestResourceTypesResource                       RequestResourceTypes = "Resource"
	RequestResourceTypesBinary                         RequestResourceTypes = "Binary"
	RequestResourceTypesBundle                         RequestResourceTypes = "Bundle"
	RequestResourceTypesDomainResource                 RequestResourceTypes = "DomainResource"
	RequestResourceTypesAccount                        RequestResourceTypes = "Account"
	RequestResourceTypesActivityDefinition             RequestResourceTypes = "ActivityDefinition"
	RequestResourceTypesActorDefinition                RequestResourceTypes = "ActorDefinition"
	RequestResourceTypesAdministrableProductDefinition RequestResourceTypes = "AdministrableProductDefinition"
	RequestResourceTypesAdverseEvent                   RequestResourceTypes = "AdverseEvent"
	RequestResourceTypesAllergyIntolerance             RequestResourceTypes = "AllergyIntolerance"
	RequestResourceTypesAppointment                    RequestResourceTypes = "Appointment"
	RequestResourceTypesAppointmentResponse            RequestResourceTypes = "AppointmentResponse"
	RequestResourceTypesArtifactAssessment             RequestResourceTypes = "ArtifactAssessment"
	RequestResourceTypesAuditEvent                     RequestResourceTypes = "AuditEvent"
	RequestResourceTypesBasic                          RequestResourceTypes = "Basic"
	RequestResourceTypesBiologicallyDerivedProduct     RequestResourceTypes = "BiologicallyDerivedProduct"
	RequestResourceTypesBodyStructure                  RequestResourceTypes = "BodyStructure"
	RequestResourceTypesCanonicalResource              RequestResourceTypes = "CanonicalResource"
	RequestResourceTypesCapabilityStatement            RequestResourceTypes = "CapabilityStatement"
	RequestResourceTypesCarePlan                       RequestResourceTypes = "CarePlan"
	RequestResourceTypesCareTeam                       RequestResourceTypes = "CareTeam"
	RequestResourceTypesClaim                          RequestResourceTypes = "Claim"
	RequestResourceTypesClaimResponse                  RequestResourceTypes = "ClaimResponse"
	RequestResourceTypesClinicalUseDefinition          RequestResourceTypes = "ClinicalUseDefinition"
	RequestResourceTypesCodeSystem                     RequestResourceTypes = "CodeSystem"
	RequestResourceTypesCommunication                  RequestResourceTypes = "Communication"
	RequestResourceTypesCommunicationRequest           RequestResourceTypes = "CommunicationRequest"
	RequestResourceTypesCompartmentDefinition          RequestResourceTypes = "CompartmentDefinition"
	RequestResourceTypesComposition                    RequestResourceTypes = "Composition"
	RequestResourceTypesConceptMap                     RequestResourceTypes = "ConceptMap"
	RequestResourceTypesCondition                      RequestResourceTypes = "Condition"
	RequestResourceTypesConsent                        RequestResourceTypes = "Consent"
	RequestResourceTypesContract                       RequestResourceTypes = "Contract"
	RequestResourceTypesCoverage                       RequestResourceTypes = "Coverage"
	RequestResourceTypesCoverageEligibilityRequest     RequestResourceTypes = "CoverageEligibilityRequest"
	RequestResourceTypesCoverageEligibilityResponse    RequestResourceTypes = "CoverageEligibilityResponse"
	RequestResourceTypesDetectedIssue                  RequestResourceTypes = "DetectedIssue"
	RequestResourceTypesDevice                         RequestResourceTypes = "Device"
	RequestResourceTypesDeviceAlert                    RequestResourceTypes = "DeviceAlert"
	RequestResourceTypesDeviceAssociation              RequestResourceTypes = "DeviceAssociation"
	RequestResourceTypesDeviceDefinition               RequestResourceTypes = "DeviceDefinition"
	RequestResourceTypesDeviceMetric                   RequestResourceTypes = "DeviceMetric"
	RequestResourceTypesDeviceRequest                  RequestResourceTypes = "DeviceRequest"
	RequestResourceTypesDiagnosticReport               RequestResourceTypes = "DiagnosticReport"
	RequestResourceTypesDocumentReference              RequestResourceTypes = "DocumentReference"
	RequestResourceTypesEncounter                      RequestResourceTypes = "Encounter"
	RequestResourceTypesEndpoint                       RequestResourceTypes = "Endpoint"
	RequestResourceTypesEnrollmentRequest              RequestResourceTypes = "EnrollmentRequest"
	RequestResourceTypesEnrollmentResponse             RequestResourceTypes = "EnrollmentResponse"
	RequestResourceTypesEpisodeOfCare                  RequestResourceTypes = "EpisodeOfCare"
	RequestResourceTypesEventDefinition                RequestResourceTypes = "EventDefinition"
	RequestResourceTypesEvidence                       RequestResourceTypes = "Evidence"
	RequestResourceTypesEvidenceVariable               RequestResourceTypes = "EvidenceVariable"
	RequestResourceTypesExampleScenario                RequestResourceTypes = "ExampleScenario"
	RequestResourceTypesExplanationOfBenefit           RequestResourceTypes = "ExplanationOfBenefit"
	RequestResourceTypesFamilyMemberHistory            RequestResourceTypes = "FamilyMemberHistory"
	RequestResourceTypesFlag                           RequestResourceTypes = "Flag"
	RequestResourceTypesGoal                           RequestResourceTypes = "Goal"
	RequestResourceTypesGroup                          RequestResourceTypes = "Group"
	RequestResourceTypesGuidanceResponse               RequestResourceTypes = "GuidanceResponse"
	RequestResourceTypesHealthcareService              RequestResourceTypes = "HealthcareService"
	RequestResourceTypesImagingSelection               RequestResourceTypes = "ImagingSelection"
	RequestResourceTypesImagingStudy                   RequestResourceTypes = "ImagingStudy"
	RequestResourceTypesImmunization                   RequestResourceTypes = "Immunization"
	RequestResourceTypesImplementationGuide            RequestResourceTypes = "ImplementationGuide"
	RequestResourceTypesIngredient                     RequestResourceTypes = "Ingredient"
	RequestResourceTypesInsurancePlan                  RequestResourceTypes = "InsurancePlan"
	RequestResourceTypesInsuranceProduct               RequestResourceTypes = "InsuranceProduct"
	RequestResourceTypesInvoice                        RequestResourceTypes = "Invoice"
	RequestResourceTypesLibrary                        RequestResourceTypes = "Library"
	RequestResourceTypesList                           RequestResourceTypes = "List"
	RequestResourceTypesLocation                       RequestResourceTypes = "Location"
	RequestResourceTypesManufacturedItemDefinition     RequestResourceTypes = "ManufacturedItemDefinition"
	RequestResourceTypesMeasure                        RequestResourceTypes = "Measure"
	RequestResourceTypesMeasureReport                  RequestResourceTypes = "MeasureReport"
	RequestResourceTypesMedication                     RequestResourceTypes = "Medication"
	RequestResourceTypesMedicationAdministration       RequestResourceTypes = "MedicationAdministration"
	RequestResourceTypesMedicationDispense             RequestResourceTypes = "MedicationDispense"
	RequestResourceTypesMedicationRequest              RequestResourceTypes = "MedicationRequest"
	RequestResourceTypesMedicationStatement            RequestResourceTypes = "MedicationStatement"
	RequestResourceTypesMedicinalProductDefinition     RequestResourceTypes = "MedicinalProductDefinition"
	RequestResourceTypesMessageDefinition              RequestResourceTypes = "MessageDefinition"
	RequestResourceTypesMessageHeader                  RequestResourceTypes = "MessageHeader"
	RequestResourceTypesMetadataResource               RequestResourceTypes = "MetadataResource"
	RequestResourceTypesNamingSystem                   RequestResourceTypes = "NamingSystem"
	RequestResourceTypesNutritionIntake                RequestResourceTypes = "NutritionIntake"
	RequestResourceTypesNutritionOrder                 RequestResourceTypes = "NutritionOrder"
	RequestResourceTypesNutritionProduct               RequestResourceTypes = "NutritionProduct"
	RequestResourceTypesObservation                    RequestResourceTypes = "Observation"
	RequestResourceTypesObservationDefinition          RequestResourceTypes = "ObservationDefinition"
	RequestResourceTypesOperationDefinition            RequestResourceTypes = "OperationDefinition"
	RequestResourceTypesOperationOutcome               RequestResourceTypes = "OperationOutcome"
	RequestResourceTypesOrganization                   RequestResourceTypes = "Organization"
	RequestResourceTypesOrganizationAffiliation        RequestResourceTypes = "OrganizationAffiliation"
	RequestResourceTypesPackagedProductDefinition      RequestResourceTypes = "PackagedProductDefinition"
	RequestResourceTypesPatient                        RequestResourceTypes = "Patient"
	RequestResourceTypesPaymentNotice                  RequestResourceTypes = "PaymentNotice"
	RequestResourceTypesPaymentReconciliation          RequestResourceTypes = "PaymentReconciliation"
	RequestResourceTypesPerson                         RequestResourceTypes = "Person"
	RequestResourceTypesPlanDefinition                 RequestResourceTypes = "PlanDefinition"
	RequestResourceTypesPractitioner                   RequestResourceTypes = "Practitioner"
	RequestResourceTypesPractitionerRole               RequestResourceTypes = "PractitionerRole"
	RequestResourceTypesProcedure                      RequestResourceTypes = "Procedure"
	RequestResourceTypesProvenance                     RequestResourceTypes = "Provenance"
	RequestResourceTypesQuestionnaire                  RequestResourceTypes = "Questionnaire"
	RequestResourceTypesQuestionnaireResponse          RequestResourceTypes = "QuestionnaireResponse"
	RequestResourceTypesRegulatedAuthorization         RequestResourceTypes = "RegulatedAuthorization"
	RequestResourceTypesRelatedPerson                  RequestResourceTypes = "RelatedPerson"
	RequestResourceTypesRequestOrchestration           RequestResourceTypes = "RequestOrchestration"
	RequestResourceTypesRequirements                   RequestResourceTypes = "Requirements"
	RequestResourceTypesResearchStudy                  RequestResourceTypes = "ResearchStudy"
	RequestResourceTypesResearchSubject                RequestResourceTypes = "ResearchSubject"
	RequestResourceTypesRiskAssessment                 RequestResourceTypes = "RiskAssessment"
	RequestResourceTypesSchedule                       RequestResourceTypes = "Schedule"
	RequestResourceTypesSearchParameter                RequestResourceTypes = "SearchParameter"
	RequestResourceTypesServiceRequest                 RequestResourceTypes = "ServiceRequest"
	RequestResourceTypesSlot                           RequestResourceTypes = "Slot"
	RequestResourceTypesSpecimen                       RequestResourceTypes = "Specimen"
	RequestResourceTypesSpecimenDefinition             RequestResourceTypes = "SpecimenDefinition"
	RequestResourceTypesStructureDefinition            RequestResourceTypes = "StructureDefinition"
	RequestResourceTypesStructureMap                   RequestResourceTypes = "StructureMap"
	RequestResourceTypesSubscription                   RequestResourceTypes = "Subscription"
	RequestResourceTypesSubscriptionStatus             RequestResourceTypes = "SubscriptionStatus"
	RequestResourceTypesSubscriptionTopic              RequestResourceTypes = "SubscriptionTopic"
	RequestResourceTypesSubstance                      RequestResourceTypes = "Substance"
	RequestResourceTypesSubstanceDefinition            RequestResourceTypes = "SubstanceDefinition"
	RequestResourceTypesTask                           RequestResourceTypes = "Task"
	RequestResourceTypesTerminologyCapabilities        RequestResourceTypes = "TerminologyCapabilities"
	RequestResourceTypesValueSet                       RequestResourceTypes = "ValueSet"
	RequestResourceTypesVisionPrescription             RequestResourceTypes = "VisionPrescription"
	RequestResourceTypesParameters                     RequestResourceTypes = "Parameters"
)

// RequestStatus represents codes from http://hl7.org/fhir/ValueSet/request-status
type RequestStatus string

const (
	RequestStatusDraft          RequestStatus = "draft"
	RequestStatusActive         RequestStatus = "active"
	RequestStatusOnHold         RequestStatus = "on-hold"
	RequestStatusEnteredInError RequestStatus = "entered-in-error"
	RequestStatusEnded          RequestStatus = "ended"
	RequestStatusCompleted      RequestStatus = "completed"
	RequestStatusRevoked        RequestStatus = "revoked"
	RequestStatusUnknown        RequestStatus = "unknown"
)

// ResourceType represents codes from http://hl7.org/fhir/ValueSet/resource-types
type ResourceType string

const (
	ResourceTypeBase                           ResourceType = "Base"
	ResourceTypeElement                        ResourceType = "Element"
	ResourceTypeBackboneElement                ResourceType = "BackboneElement"
	ResourceTypeDataType                       ResourceType = "DataType"
	ResourceTypeAddress                        ResourceType = "Address"
	ResourceTypeAnnotation                     ResourceType = "Annotation"
	ResourceTypeAttachment                     ResourceType = "Attachment"
	ResourceTypeAvailability                   ResourceType = "Availability"
	ResourceTypeBackboneType                   ResourceType = "BackboneType"
	ResourceTypeDosage                         ResourceType = "Dosage"
	ResourceTypeDosageCondition                ResourceType = "DosageCondition"
	ResourceTypeDosageDetails                  ResourceType = "DosageDetails"
	ResourceTypeDosageSafety                   ResourceType = "DosageSafety"
	ResourceTypeElementDefinition              ResourceType = "ElementDefinition"
	ResourceTypeMarketingStatus                ResourceType = "MarketingStatus"
	ResourceTypeProductShelfLife               ResourceType = "ProductShelfLife"
	ResourceTypeRelativeTime                   ResourceType = "RelativeTime"
	ResourceTypeTiming                         ResourceType = "Timing"
	ResourceTypeCodeableConcept                ResourceType = "CodeableConcept"
	ResourceTypeCodeableReference              ResourceType = "CodeableReference"
	ResourceTypeCoding                         ResourceType = "Coding"
	ResourceTypeContactDetail                  ResourceType = "ContactDetail"
	ResourceTypeContactPoint                   ResourceType = "ContactPoint"
	ResourceTypeDataRequirement                ResourceType = "DataRequirement"
	ResourceTypeExpression                     ResourceType = "Expression"
	ResourceTypeExtendedContactDetail          ResourceType = "ExtendedContactDetail"
	ResourceTypeExtension                      ResourceType = "Extension"
	ResourceTypeHumanName                      ResourceType = "HumanName"
	ResourceTypeIdentifier                     ResourceType = "Identifier"
	ResourceTypeMeta                           ResourceType = "Meta"
	ResourceTypeMonetaryComponent              ResourceType = "MonetaryComponent"
	ResourceTypeMoney                          ResourceType = "Money"
	ResourceTypeNarrative                      ResourceType = "Narrative"
	ResourceTypeParameterDefinition            ResourceType = "ParameterDefinition"
	ResourceTypePeriod                         ResourceType = "Period"
	ResourceTypePrimitiveType                  ResourceType = "PrimitiveType"
	ResourceTypeBase64Binary                   ResourceType = "base64Binary"
	ResourceTypeBoolean                        ResourceType = "boolean"
	ResourceTypeDate                           ResourceType = "date"
	ResourceTypeDateTime                       ResourceType = "dateTime"
	ResourceTypeDecimal                        ResourceType = "decimal"
	ResourceTypeInstant                        ResourceType = "instant"
	ResourceTypeInteger                        ResourceType = "integer"
	ResourceTypePositiveInt                    ResourceType = "positiveInt"
	ResourceTypeUnsignedInt                    ResourceType = "unsignedInt"
	ResourceTypeInteger64                      ResourceType = "integer64"
	ResourceTypeString                         ResourceType = "string"
	ResourceTypeCode                           ResourceType = "code"
	ResourceTypeId                             ResourceType = "id"
	ResourceTypeMarkdown                       ResourceType = "markdown"
	ResourceTypeTime                           ResourceType = "time"
	ResourceTypeUri                            ResourceType = "uri"
	ResourceTypeCanonical                      ResourceType = "canonical"
	ResourceTypeOid                            ResourceType = "oid"
	ResourceTypeUrl                            ResourceType = "url"
	ResourceTypeUuid                           ResourceType = "uuid"
	ResourceTypeQuantity                       ResourceType = "Quantity"
	ResourceTypeAge                            ResourceType = "Age"
	ResourceTypeCount                          ResourceType = "Count"
	ResourceTypeDistance                       ResourceType = "Distance"
	ResourceTypeDuration                       ResourceType = "Duration"
	ResourceTypeRange                          ResourceType = "Range"
	ResourceTypeRatio                          ResourceType = "Ratio"
	ResourceTypeRatioRange                     ResourceType = "RatioRange"
	ResourceTypeReference                      ResourceType = "Reference"
	ResourceTypeRelatedArtifact                ResourceType = "RelatedArtifact"
	ResourceTypeSampledData                    ResourceType = "SampledData"
	ResourceTypeSignature                      ResourceType = "Signature"
	ResourceTypeTriggerDefinition              ResourceType = "TriggerDefinition"
	ResourceTypeUsageContext                   ResourceType = "UsageContext"
	ResourceTypeVirtualServiceDetail           ResourceType = "VirtualServiceDetail"
	ResourceTypeXhtml                          ResourceType = "xhtml"
	ResourceTypeResource                       ResourceType = "Resource"
	ResourceTypeBinary                         ResourceType = "Binary"
	ResourceTypeBundle                         ResourceType = "Bundle"
	ResourceTypeDomainResource                 ResourceType = "DomainResource"
	ResourceTypeAccount                        ResourceType = "Account"
	ResourceTypeActivityDefinition             ResourceType = "ActivityDefinition"
	ResourceTypeActorDefinition                ResourceType = "ActorDefinition"
	ResourceTypeAdministrableProductDefinition ResourceType = "AdministrableProductDefinition"
	ResourceTypeAdverseEvent                   ResourceType = "AdverseEvent"
	ResourceTypeAllergyIntolerance             ResourceType = "AllergyIntolerance"
	ResourceTypeAppointment                    ResourceType = "Appointment"
	ResourceTypeAppointmentResponse            ResourceType = "AppointmentResponse"
	ResourceTypeArtifactAssessment             ResourceType = "ArtifactAssessment"
	ResourceTypeAuditEvent                     ResourceType = "AuditEvent"
	ResourceTypeBasic                          ResourceType = "Basic"
	ResourceTypeBiologicallyDerivedProduct     ResourceType = "BiologicallyDerivedProduct"
	ResourceTypeBodyStructure                  ResourceType = "BodyStructure"
	ResourceTypeCanonicalResource              ResourceType = "CanonicalResource"
	ResourceTypeCapabilityStatement            ResourceType = "CapabilityStatement"
	ResourceTypeCarePlan                       ResourceType = "CarePlan"
	ResourceTypeCareTeam                       ResourceType = "CareTeam"
	ResourceTypeClaim                          ResourceType = "Claim"
	ResourceTypeClaimResponse                  ResourceType = "ClaimResponse"
	ResourceTypeClinicalUseDefinition          ResourceType = "ClinicalUseDefinition"
	ResourceTypeCodeSystem                     ResourceType = "CodeSystem"
	ResourceTypeCommunication                  ResourceType = "Communication"
	ResourceTypeCommunicationRequest           ResourceType = "CommunicationRequest"
	ResourceTypeCompartmentDefinition          ResourceType = "CompartmentDefinition"
	ResourceTypeComposition                    ResourceType = "Composition"
	ResourceTypeConceptMap                     ResourceType = "ConceptMap"
	ResourceTypeCondition                      ResourceType = "Condition"
	ResourceTypeConsent                        ResourceType = "Consent"
	ResourceTypeContract                       ResourceType = "Contract"
	ResourceTypeCoverage                       ResourceType = "Coverage"
	ResourceTypeCoverageEligibilityRequest     ResourceType = "CoverageEligibilityRequest"
	ResourceTypeCoverageEligibilityResponse    ResourceType = "CoverageEligibilityResponse"
	ResourceTypeDetectedIssue                  ResourceType = "DetectedIssue"
	ResourceTypeDevice                         ResourceType = "Device"
	ResourceTypeDeviceAlert                    ResourceType = "DeviceAlert"
	ResourceTypeDeviceAssociation              ResourceType = "DeviceAssociation"
	ResourceTypeDeviceDefinition               ResourceType = "DeviceDefinition"
	ResourceTypeDeviceMetric                   ResourceType = "DeviceMetric"
	ResourceTypeDeviceRequest                  ResourceType = "DeviceRequest"
	ResourceTypeDiagnosticReport               ResourceType = "DiagnosticReport"
	ResourceTypeDocumentReference              ResourceType = "DocumentReference"
	ResourceTypeEncounter                      ResourceType = "Encounter"
	ResourceTypeEndpoint                       ResourceType = "Endpoint"
	ResourceTypeEnrollmentRequest              ResourceType = "EnrollmentRequest"
	ResourceTypeEnrollmentResponse             ResourceType = "EnrollmentResponse"
	ResourceTypeEpisodeOfCare                  ResourceType = "EpisodeOfCare"
	ResourceTypeEventDefinition                ResourceType = "EventDefinition"
	ResourceTypeEvidence                       ResourceType = "Evidence"
	ResourceTypeEvidenceVariable               ResourceType = "EvidenceVariable"
	ResourceTypeExampleScenario                ResourceType = "ExampleScenario"
	ResourceTypeExplanationOfBenefit           ResourceType = "ExplanationOfBenefit"
	ResourceTypeFamilyMemberHistory            ResourceType = "FamilyMemberHistory"
	ResourceTypeFlag                           ResourceType = "Flag"
	ResourceTypeGoal                           ResourceType = "Goal"
	ResourceTypeGroup                          ResourceType = "Group"
	ResourceTypeGuidanceResponse               ResourceType = "GuidanceResponse"
	ResourceTypeHealthcareService              ResourceType = "HealthcareService"
	ResourceTypeImagingSelection               ResourceType = "ImagingSelection"
	ResourceTypeImagingStudy                   ResourceType = "ImagingStudy"
	ResourceTypeImmunization                   ResourceType = "Immunization"
	ResourceTypeImplementationGuide            ResourceType = "ImplementationGuide"
	ResourceTypeIngredient                     ResourceType = "Ingredient"
	ResourceTypeInsurancePlan                  ResourceType = "InsurancePlan"
	ResourceTypeInsuranceProduct               ResourceType = "InsuranceProduct"
	ResourceTypeInvoice                        ResourceType = "Invoice"
	ResourceTypeLibrary                        ResourceType = "Library"
	ResourceTypeList                           ResourceType = "List"
	ResourceTypeLocation                       ResourceType = "Location"
	ResourceTypeManufacturedItemDefinition     ResourceType = "ManufacturedItemDefinition"
	ResourceTypeMeasure                        ResourceType = "Measure"
	ResourceTypeMeasureReport                  ResourceType = "MeasureReport"
	ResourceTypeMedication                     ResourceType = "Medication"
	ResourceTypeMedicationAdministration       ResourceType = "MedicationAdministration"
	ResourceTypeMedicationDispense             ResourceType = "MedicationDispense"
	ResourceTypeMedicationRequest              ResourceType = "MedicationRequest"
	ResourceTypeMedicationStatement            ResourceType = "MedicationStatement"
	ResourceTypeMedicinalProductDefinition     ResourceType = "MedicinalProductDefinition"
	ResourceTypeMessageDefinition              ResourceType = "MessageDefinition"
	ResourceTypeMessageHeader                  ResourceType = "MessageHeader"
	ResourceTypeMetadataResource               ResourceType = "MetadataResource"
	ResourceTypeNamingSystem                   ResourceType = "NamingSystem"
	ResourceTypeNutritionIntake                ResourceType = "NutritionIntake"
	ResourceTypeNutritionOrder                 ResourceType = "NutritionOrder"
	ResourceTypeNutritionProduct               ResourceType = "NutritionProduct"
	ResourceTypeObservation                    ResourceType = "Observation"
	ResourceTypeObservationDefinition          ResourceType = "ObservationDefinition"
	ResourceTypeOperationDefinition            ResourceType = "OperationDefinition"
	ResourceTypeOperationOutcome               ResourceType = "OperationOutcome"
	ResourceTypeOrganization                   ResourceType = "Organization"
	ResourceTypeOrganizationAffiliation        ResourceType = "OrganizationAffiliation"
	ResourceTypePackagedProductDefinition      ResourceType = "PackagedProductDefinition"
	ResourceTypePatient                        ResourceType = "Patient"
	ResourceTypePaymentNotice                  ResourceType = "PaymentNotice"
	ResourceTypePaymentReconciliation          ResourceType = "PaymentReconciliation"
	ResourceTypePerson                         ResourceType = "Person"
	ResourceTypePlanDefinition                 ResourceType = "PlanDefinition"
	ResourceTypePractitioner                   ResourceType = "Practitioner"
	ResourceTypePractitionerRole               ResourceType = "PractitionerRole"
	ResourceTypeProcedure                      ResourceType = "Procedure"
	ResourceTypeProvenance                     ResourceType = "Provenance"
	ResourceTypeQuestionnaire                  ResourceType = "Questionnaire"
	ResourceTypeQuestionnaireResponse          ResourceType = "QuestionnaireResponse"
	ResourceTypeRegulatedAuthorization         ResourceType = "RegulatedAuthorization"
	ResourceTypeRelatedPerson                  ResourceType = "RelatedPerson"
	ResourceTypeRequestOrchestration           ResourceType = "RequestOrchestration"
	ResourceTypeRequirements                   ResourceType = "Requirements"
	ResourceTypeResearchStudy                  ResourceType = "ResearchStudy"
	ResourceTypeResearchSubject                ResourceType = "ResearchSubject"
	ResourceTypeRiskAssessment                 ResourceType = "RiskAssessment"
	ResourceTypeSchedule                       ResourceType = "Schedule"
	ResourceTypeSearchParameter                ResourceType = "SearchParameter"
	ResourceTypeServiceRequest                 ResourceType = "ServiceRequest"
	ResourceTypeSlot                           ResourceType = "Slot"
	ResourceTypeSpecimen                       ResourceType = "Specimen"
	ResourceTypeSpecimenDefinition             ResourceType = "SpecimenDefinition"
	ResourceTypeStructureDefinition            ResourceType = "StructureDefinition"
	ResourceTypeStructureMap                   ResourceType = "StructureMap"
	ResourceTypeSubscription                   ResourceType = "Subscription"
	ResourceTypeSubscriptionStatus             ResourceType = "SubscriptionStatus"
	ResourceTypeSubscriptionTopic              ResourceType = "SubscriptionTopic"
	ResourceTypeSubstance                      ResourceType = "Substance"
	ResourceTypeSubstanceDefinition            ResourceType = "SubstanceDefinition"
	ResourceTypeTask                           ResourceType = "Task"
	ResourceTypeTerminologyCapabilities        ResourceType = "TerminologyCapabilities"
	ResourceTypeValueSet                       ResourceType = "ValueSet"
	ResourceTypeVisionPrescription             ResourceType = "VisionPrescription"
	ResourceTypeParameters                     ResourceType = "Parameters"
)

// ResourceVersionPolicy represents codes from http://hl7.org/fhir/ValueSet/versioning-policy
type ResourceVersionPolicy string

const (
	ResourceVersionPolicyNoVersion       ResourceVersionPolicy = "no-version"
	ResourceVersionPolicyVersioned       ResourceVersionPolicy = "versioned"
	ResourceVersionPolicyVersionedUpdate ResourceVersionPolicy = "versioned-update"
)

// ResponseType represents codes from http://hl7.org/fhir/ValueSet/response-code
type ResponseType string

const (
	ResponseTypeOk             ResponseType = "ok"
	ResponseTypeTransientError ResponseType = "transient-error"
	ResponseTypeFatalError     ResponseType = "fatal-error"
)

// RestfulCapabilityMode represents codes from http://hl7.org/fhir/ValueSet/restful-capability-mode
type RestfulCapabilityMode string

const (
	RestfulCapabilityModeClient RestfulCapabilityMode = "client"
	RestfulCapabilityModeServer RestfulCapabilityMode = "server"
)
