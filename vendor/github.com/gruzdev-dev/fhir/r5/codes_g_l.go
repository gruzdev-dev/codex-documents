package models

// GlobalLangPackSupportVS represents codes from http://hl7.org/fhir/ValueSet/global-langpack-support
type GlobalLangPackSupportVS string

const (
	GlobalLangPackSupportVSNotSupported GlobalLangPackSupportVS = "not-supported"
	GlobalLangPackSupportVSExplicit     GlobalLangPackSupportVS = "explicit"
	GlobalLangPackSupportVSImplicit     GlobalLangPackSupportVS = "implicit"
)

// GoalAcceptStatus represents codes from http://hl7.org/fhir/ValueSet/goal-accept-status
type GoalAcceptStatus string

const (
	GoalAcceptStatusAgree    GoalAcceptStatus = "agree"
	GoalAcceptStatusDisagree GoalAcceptStatus = "disagree"
	GoalAcceptStatusPending  GoalAcceptStatus = "pending"
)

// GoalLifecycleStatus represents codes from http://hl7.org/fhir/ValueSet/goal-status
type GoalLifecycleStatus string

const (
	GoalLifecycleStatusProposed       GoalLifecycleStatus = "proposed"
	GoalLifecycleStatusPlanned        GoalLifecycleStatus = "planned"
	GoalLifecycleStatusAccepted       GoalLifecycleStatus = "accepted"
	GoalLifecycleStatusActive         GoalLifecycleStatus = "active"
	GoalLifecycleStatusOnHold         GoalLifecycleStatus = "on-hold"
	GoalLifecycleStatusCompleted      GoalLifecycleStatus = "completed"
	GoalLifecycleStatusCancelled      GoalLifecycleStatus = "cancelled"
	GoalLifecycleStatusEnteredInError GoalLifecycleStatus = "entered-in-error"
	GoalLifecycleStatusRejected       GoalLifecycleStatus = "rejected"
)

// GroupCharacteristicCombination represents codes from http://hl7.org/fhir/ValueSet/group-characteristic-combination
type GroupCharacteristicCombination string

const (
	GroupCharacteristicCombinationAllOf        GroupCharacteristicCombination = "all-of"
	GroupCharacteristicCombinationAnyOf        GroupCharacteristicCombination = "any-of"
	GroupCharacteristicCombinationAtLeast      GroupCharacteristicCombination = "at-least"
	GroupCharacteristicCombinationAtMost       GroupCharacteristicCombination = "at-most"
	GroupCharacteristicCombinationExceptSubset GroupCharacteristicCombination = "except-subset"
	GroupCharacteristicCombinationStatistical  GroupCharacteristicCombination = "statistical"
	GroupCharacteristicCombinationNetEffect    GroupCharacteristicCombination = "net-effect"
	GroupCharacteristicCombinationDataset      GroupCharacteristicCombination = "dataset"
)

// GroupMembershipBasis represents codes from http://hl7.org/fhir/ValueSet/group-membership-basis
type GroupMembershipBasis string

const (
	GroupMembershipBasisDefinitional GroupMembershipBasis = "definitional"
	GroupMembershipBasisConceptual   GroupMembershipBasis = "conceptual"
	GroupMembershipBasisEnumerated   GroupMembershipBasis = "enumerated"
)

// GroupType represents codes from http://hl7.org/fhir/ValueSet/group-type
type GroupType string

const (
	GroupTypePerson                     GroupType = "person"
	GroupTypeAnimal                     GroupType = "animal"
	GroupTypePractitioner               GroupType = "practitioner"
	GroupTypeDevice                     GroupType = "device"
	GroupTypeCareteam                   GroupType = "careteam"
	GroupTypeHealthcareservice          GroupType = "healthcareservice"
	GroupTypeLocation                   GroupType = "location"
	GroupTypeOrganization               GroupType = "organization"
	GroupTypeRelatedperson              GroupType = "relatedperson"
	GroupTypeSpecimen                   GroupType = "specimen"
	GroupTypeMedication                 GroupType = "medication"
	GroupTypeMedicinalproductdefinition GroupType = "medicinalproductdefinition"
	GroupTypeSubstance                  GroupType = "substance"
	GroupTypeSubstancedefinition        GroupType = "substancedefinition"
	GroupTypeBiologicallyDerivedProduct GroupType = "biologicallyDerivedProduct"
	GroupTypeNutritionProduct           GroupType = "nutritionProduct"
)

// GuidanceResponseStatus represents codes from http://hl7.org/fhir/ValueSet/guidance-response-status
type GuidanceResponseStatus string

const (
	GuidanceResponseStatusSuccess        GuidanceResponseStatus = "success"
	GuidanceResponseStatusDataRequested  GuidanceResponseStatus = "data-requested"
	GuidanceResponseStatusDataRequired   GuidanceResponseStatus = "data-required"
	GuidanceResponseStatusInProgress     GuidanceResponseStatus = "in-progress"
	GuidanceResponseStatusFailure        GuidanceResponseStatus = "failure"
	GuidanceResponseStatusEnteredInError GuidanceResponseStatus = "entered-in-error"
)

// GuidePageGeneration represents codes from http://hl7.org/fhir/ValueSet/guide-page-generation
type GuidePageGeneration string

const (
	GuidePageGenerationHtml      GuidePageGeneration = "html"
	GuidePageGenerationMarkdown  GuidePageGeneration = "markdown"
	GuidePageGenerationXml       GuidePageGeneration = "xml"
	GuidePageGenerationGenerated GuidePageGeneration = "generated"
)

// HTTPVerb represents codes from http://hl7.org/fhir/ValueSet/http-verb
type HTTPVerb string

const (
	HTTPVerbGET    HTTPVerb = "GET"
	HTTPVerbHEAD   HTTPVerb = "HEAD"
	HTTPVerbPOST   HTTPVerb = "POST"
	HTTPVerbPUT    HTTPVerb = "PUT"
	HTTPVerbDELETE HTTPVerb = "DELETE"
	HTTPVerbPATCH  HTTPVerb = "PATCH"
)

// IdentifierUse represents codes from http://hl7.org/fhir/ValueSet/identifier-use
type IdentifierUse string

const (
	IdentifierUseUsual     IdentifierUse = "usual"
	IdentifierUseOfficial  IdentifierUse = "official"
	IdentifierUseTemp      IdentifierUse = "temp"
	IdentifierUseSecondary IdentifierUse = "secondary"
	IdentifierUseOld       IdentifierUse = "old"
)

// IdentityAssuranceLevel represents codes from http://hl7.org/fhir/ValueSet/identity-assuranceLevel
type IdentityAssuranceLevel string

const (
	IdentityAssuranceLevelLevel1 IdentityAssuranceLevel = "level1"
	IdentityAssuranceLevelLevel2 IdentityAssuranceLevel = "level2"
	IdentityAssuranceLevelLevel3 IdentityAssuranceLevel = "level3"
	IdentityAssuranceLevelLevel4 IdentityAssuranceLevel = "level4"
)

// ImagingSelection2DGraphicType represents codes from http://hl7.org/fhir/ValueSet/imagingselection-2dgraphictype
type ImagingSelection2DGraphicType string

const (
	ImagingSelection2DGraphicTypePoint      ImagingSelection2DGraphicType = "point"
	ImagingSelection2DGraphicTypePolyline   ImagingSelection2DGraphicType = "polyline"
	ImagingSelection2DGraphicTypeMultipoint ImagingSelection2DGraphicType = "multipoint"
	ImagingSelection2DGraphicTypeCircle     ImagingSelection2DGraphicType = "circle"
	ImagingSelection2DGraphicTypeEllipse    ImagingSelection2DGraphicType = "ellipse"
)

// ImagingSelection3DGraphicType represents codes from http://hl7.org/fhir/ValueSet/imagingselection-3dgraphictype
type ImagingSelection3DGraphicType string

const (
	ImagingSelection3DGraphicTypePoint      ImagingSelection3DGraphicType = "point"
	ImagingSelection3DGraphicTypeMultipoint ImagingSelection3DGraphicType = "multipoint"
	ImagingSelection3DGraphicTypePolyline   ImagingSelection3DGraphicType = "polyline"
	ImagingSelection3DGraphicTypePolygon    ImagingSelection3DGraphicType = "polygon"
	ImagingSelection3DGraphicTypeEllipse    ImagingSelection3DGraphicType = "ellipse"
	ImagingSelection3DGraphicTypeEllipsoid  ImagingSelection3DGraphicType = "ellipsoid"
)

// ImagingSelectionStatus represents codes from http://hl7.org/fhir/ValueSet/imagingselection-status
type ImagingSelectionStatus string

const (
	ImagingSelectionStatusAvailable      ImagingSelectionStatus = "available"
	ImagingSelectionStatusEnteredInError ImagingSelectionStatus = "entered-in-error"
	ImagingSelectionStatusInactive       ImagingSelectionStatus = "inactive"
	ImagingSelectionStatusUnknown        ImagingSelectionStatus = "unknown"
)

// ImagingStudyStatus represents codes from http://hl7.org/fhir/ValueSet/imagingstudy-status
type ImagingStudyStatus string

const (
	ImagingStudyStatusRegistered     ImagingStudyStatus = "registered"
	ImagingStudyStatusAvailable      ImagingStudyStatus = "available"
	ImagingStudyStatusCancelled      ImagingStudyStatus = "cancelled"
	ImagingStudyStatusEnteredInError ImagingStudyStatus = "entered-in-error"
	ImagingStudyStatusUnknown        ImagingStudyStatus = "unknown"
	ImagingStudyStatusInactive       ImagingStudyStatus = "inactive"
)

// ImmunizationStatusCodes represents codes from http://hl7.org/fhir/ValueSet/immunization-status
type ImmunizationStatusCodes string

const (
	ImmunizationStatusCodesPreparation    ImmunizationStatusCodes = "preparation"
	ImmunizationStatusCodesInProgress     ImmunizationStatusCodes = "in-progress"
	ImmunizationStatusCodesNotDone        ImmunizationStatusCodes = "not-done"
	ImmunizationStatusCodesOnHold         ImmunizationStatusCodes = "on-hold"
	ImmunizationStatusCodesStopped        ImmunizationStatusCodes = "stopped"
	ImmunizationStatusCodesCompleted      ImmunizationStatusCodes = "completed"
	ImmunizationStatusCodesEnteredInError ImmunizationStatusCodes = "entered-in-error"
	ImmunizationStatusCodesUnknown        ImmunizationStatusCodes = "unknown"
)

// IngredientManufacturerRole represents codes from http://hl7.org/fhir/ValueSet/ingredient-manufacturer-role
type IngredientManufacturerRole string

const (
	IngredientManufacturerRoleAllowed  IngredientManufacturerRole = "allowed"
	IngredientManufacturerRolePossible IngredientManufacturerRole = "possible"
	IngredientManufacturerRoleActual   IngredientManufacturerRole = "actual"
)

// InteractionTrigger represents codes from http://hl7.org/fhir/ValueSet/interaction-trigger
type InteractionTrigger string

const (
	InteractionTriggerRead                      InteractionTrigger = "read"
	InteractionTriggerVread                     InteractionTrigger = "vread"
	InteractionTriggerUpdate                    InteractionTrigger = "update"
	InteractionTriggerUpdateConditional         InteractionTrigger = "update-conditional"
	InteractionTriggerPatch                     InteractionTrigger = "patch"
	InteractionTriggerPatchConditional          InteractionTrigger = "patch-conditional"
	InteractionTriggerDelete                    InteractionTrigger = "delete"
	InteractionTriggerDeleteConditionalSingle   InteractionTrigger = "delete-conditional-single"
	InteractionTriggerDeleteConditionalMultiple InteractionTrigger = "delete-conditional-multiple"
	InteractionTriggerDeleteHistory             InteractionTrigger = "delete-history"
	InteractionTriggerDeleteHistoryVersion      InteractionTrigger = "delete-history-version"
	InteractionTriggerHistory                   InteractionTrigger = "history"
	InteractionTriggerHistoryInstance           InteractionTrigger = "history-instance"
	InteractionTriggerHistoryType               InteractionTrigger = "history-type"
	InteractionTriggerHistorySystem             InteractionTrigger = "history-system"
	InteractionTriggerCreate                    InteractionTrigger = "create"
	InteractionTriggerCreateConditional         InteractionTrigger = "create-conditional"
	InteractionTriggerSearch                    InteractionTrigger = "search"
	InteractionTriggerSearchType                InteractionTrigger = "search-type"
	InteractionTriggerSearchSystem              InteractionTrigger = "search-system"
	InteractionTriggerSearchCompartment         InteractionTrigger = "search-compartment"
	InteractionTriggerCapabilities              InteractionTrigger = "capabilities"
	InteractionTriggerTransaction               InteractionTrigger = "transaction"
	InteractionTriggerBatch                     InteractionTrigger = "batch"
	InteractionTriggerOperation                 InteractionTrigger = "operation"
)

// InvoiceStatus represents codes from http://hl7.org/fhir/ValueSet/invoice-status
type InvoiceStatus string

const (
	InvoiceStatusDraft          InvoiceStatus = "draft"
	InvoiceStatusIssued         InvoiceStatus = "issued"
	InvoiceStatusBalanced       InvoiceStatus = "balanced"
	InvoiceStatusCancelled      InvoiceStatus = "cancelled"
	InvoiceStatusEnteredInError InvoiceStatus = "entered-in-error"
)

// IssueSeverity represents codes from http://hl7.org/fhir/ValueSet/issue-severity
type IssueSeverity string

const (
	IssueSeverityFatal       IssueSeverity = "fatal"
	IssueSeverityError       IssueSeverity = "error"
	IssueSeverityWarning     IssueSeverity = "warning"
	IssueSeverityInformation IssueSeverity = "information"
	IssueSeveritySuccess     IssueSeverity = "success"
)

// IssueType represents codes from http://hl7.org/fhir/ValueSet/issue-type
type IssueType string

const (
	IssueTypeInvalid         IssueType = "invalid"
	IssueTypeStructure       IssueType = "structure"
	IssueTypeRequired        IssueType = "required"
	IssueTypeValue           IssueType = "value"
	IssueTypeInvariant       IssueType = "invariant"
	IssueTypeSecurity        IssueType = "security"
	IssueTypeLogin           IssueType = "login"
	IssueTypeUnknown         IssueType = "unknown"
	IssueTypeExpired         IssueType = "expired"
	IssueTypeForbidden       IssueType = "forbidden"
	IssueTypeSuppressed      IssueType = "suppressed"
	IssueTypeProcessing      IssueType = "processing"
	IssueTypeNotSupported    IssueType = "not-supported"
	IssueTypeDuplicate       IssueType = "duplicate"
	IssueTypeMultipleMatches IssueType = "multiple-matches"
	IssueTypeNotFound        IssueType = "not-found"
	IssueTypeDeleted         IssueType = "deleted"
	IssueTypeTooLong         IssueType = "too-long"
	IssueTypeCodeInvalid     IssueType = "code-invalid"
	IssueTypeExtension       IssueType = "extension"
	IssueTypeTooCostly       IssueType = "too-costly"
	IssueTypeBusinessRule    IssueType = "business-rule"
	IssueTypeConflict        IssueType = "conflict"
	IssueTypeLimitedFilter   IssueType = "limited-filter"
	IssueTypeTransient       IssueType = "transient"
	IssueTypeLockError       IssueType = "lock-error"
	IssueTypeNoStore         IssueType = "no-store"
	IssueTypeException       IssueType = "exception"
	IssueTypeTimeout         IssueType = "timeout"
	IssueTypeIncomplete      IssueType = "incomplete"
	IssueTypeThrottled       IssueType = "throttled"
	IssueTypeInformational   IssueType = "informational"
	IssueTypeSuccess         IssueType = "success"
)

// Kind represents codes from http://hl7.org/fhir/ValueSet/coverage-kind
type Kind string

const (
	KindInsurance Kind = "insurance"
	KindSelfPay   Kind = "self-pay"
	KindOther     Kind = "other"
)

// LinkRelationTypes represents codes from http://hl7.org/fhir/ValueSet/iana-link-relations
type LinkRelationTypes string

const (
	LinkRelationTypesAbout                  LinkRelationTypes = "about"
	LinkRelationTypesAcl                    LinkRelationTypes = "acl"
	LinkRelationTypesAlternate              LinkRelationTypes = "alternate"
	LinkRelationTypesAmphtml                LinkRelationTypes = "amphtml"
	LinkRelationTypesAppendix               LinkRelationTypes = "appendix"
	LinkRelationTypesAppleTouchIcon         LinkRelationTypes = "apple-touch-icon"
	LinkRelationTypesAppleTouchStartupImage LinkRelationTypes = "apple-touch-startup-image"
	LinkRelationTypesArchives               LinkRelationTypes = "archives"
	LinkRelationTypesAuthor                 LinkRelationTypes = "author"
	LinkRelationTypesBlockedBy              LinkRelationTypes = "blocked-by"
	LinkRelationTypesBookmark               LinkRelationTypes = "bookmark"
	LinkRelationTypesCanonical              LinkRelationTypes = "canonical"
	LinkRelationTypesChapter                LinkRelationTypes = "chapter"
	LinkRelationTypesCiteAs                 LinkRelationTypes = "cite-as"
	LinkRelationTypesCollection             LinkRelationTypes = "collection"
	LinkRelationTypesContents               LinkRelationTypes = "contents"
	LinkRelationTypesConvertedFrom          LinkRelationTypes = "convertedFrom"
	LinkRelationTypesCopyright              LinkRelationTypes = "copyright"
	LinkRelationTypesCreateForm             LinkRelationTypes = "create-form"
	LinkRelationTypesCurrent                LinkRelationTypes = "current"
	LinkRelationTypesDescribedby            LinkRelationTypes = "describedby"
	LinkRelationTypesDescribes              LinkRelationTypes = "describes"
	LinkRelationTypesDisclosure             LinkRelationTypes = "disclosure"
	LinkRelationTypesDnsPrefetch            LinkRelationTypes = "dns-prefetch"
	LinkRelationTypesDuplicate              LinkRelationTypes = "duplicate"
	LinkRelationTypesEdit                   LinkRelationTypes = "edit"
	LinkRelationTypesEditForm               LinkRelationTypes = "edit-form"
	LinkRelationTypesEditMedia              LinkRelationTypes = "edit-media"
	LinkRelationTypesEnclosure              LinkRelationTypes = "enclosure"
	LinkRelationTypesExternal               LinkRelationTypes = "external"
	LinkRelationTypesFirst                  LinkRelationTypes = "first"
	LinkRelationTypesGlossary               LinkRelationTypes = "glossary"
	LinkRelationTypesHelp                   LinkRelationTypes = "help"
	LinkRelationTypesHosts                  LinkRelationTypes = "hosts"
	LinkRelationTypesHub                    LinkRelationTypes = "hub"
	LinkRelationTypesIcon                   LinkRelationTypes = "icon"
	LinkRelationTypesIndex                  LinkRelationTypes = "index"
	LinkRelationTypesIntervalAfter          LinkRelationTypes = "intervalAfter"
	LinkRelationTypesIntervalBefore         LinkRelationTypes = "intervalBefore"
	LinkRelationTypesIntervalContains       LinkRelationTypes = "intervalContains"
	LinkRelationTypesIntervalDisjoint       LinkRelationTypes = "intervalDisjoint"
	LinkRelationTypesIntervalDuring         LinkRelationTypes = "intervalDuring"
	LinkRelationTypesIntervalEquals         LinkRelationTypes = "intervalEquals"
	LinkRelationTypesIntervalFinishedBy     LinkRelationTypes = "intervalFinishedBy"
	LinkRelationTypesIntervalFinishes       LinkRelationTypes = "intervalFinishes"
	LinkRelationTypesIntervalIn             LinkRelationTypes = "intervalIn"
	LinkRelationTypesIntervalMeets          LinkRelationTypes = "intervalMeets"
	LinkRelationTypesIntervalMetBy          LinkRelationTypes = "intervalMetBy"
	LinkRelationTypesIntervalOverlappedBy   LinkRelationTypes = "intervalOverlappedBy"
	LinkRelationTypesIntervalOverlaps       LinkRelationTypes = "intervalOverlaps"
	LinkRelationTypesIntervalStartedBy      LinkRelationTypes = "intervalStartedBy"
	LinkRelationTypesIntervalStarts         LinkRelationTypes = "intervalStarts"
	LinkRelationTypesItem                   LinkRelationTypes = "item"
	LinkRelationTypesLast                   LinkRelationTypes = "last"
	LinkRelationTypesLatestVersion          LinkRelationTypes = "latest-version"
	LinkRelationTypesLicense                LinkRelationTypes = "license"
	LinkRelationTypesLinkset                LinkRelationTypes = "linkset"
	LinkRelationTypesLrdd                   LinkRelationTypes = "lrdd"
	LinkRelationTypesManifest               LinkRelationTypes = "manifest"
	LinkRelationTypesMaskIcon               LinkRelationTypes = "mask-icon"
	LinkRelationTypesMediaFeed              LinkRelationTypes = "media-feed"
	LinkRelationTypesMemento                LinkRelationTypes = "memento"
	LinkRelationTypesMicropub               LinkRelationTypes = "micropub"
	LinkRelationTypesModulepreload          LinkRelationTypes = "modulepreload"
	LinkRelationTypesMonitor                LinkRelationTypes = "monitor"
	LinkRelationTypesMonitorGroup           LinkRelationTypes = "monitor-group"
	LinkRelationTypesNext                   LinkRelationTypes = "next"
	LinkRelationTypesNextArchive            LinkRelationTypes = "next-archive"
	LinkRelationTypesNofollow               LinkRelationTypes = "nofollow"
	LinkRelationTypesNoopener               LinkRelationTypes = "noopener"
	LinkRelationTypesNoreferrer             LinkRelationTypes = "noreferrer"
	LinkRelationTypesOpener                 LinkRelationTypes = "opener"
	LinkRelationTypesOpenid2LocalId         LinkRelationTypes = "openid2.local_id"
	LinkRelationTypesOpenid2Provider        LinkRelationTypes = "openid2.provider"
	LinkRelationTypesOriginal               LinkRelationTypes = "original"
	LinkRelationTypesP3Pv1                  LinkRelationTypes = "P3Pv1"
	LinkRelationTypesPayment                LinkRelationTypes = "payment"
	LinkRelationTypesPingback               LinkRelationTypes = "pingback"
	LinkRelationTypesPreconnect             LinkRelationTypes = "preconnect"
	LinkRelationTypesPredecessorVersion     LinkRelationTypes = "predecessor-version"
	LinkRelationTypesPrefetch               LinkRelationTypes = "prefetch"
	LinkRelationTypesPreload                LinkRelationTypes = "preload"
	LinkRelationTypesPrerender              LinkRelationTypes = "prerender"
	LinkRelationTypesPrev                   LinkRelationTypes = "prev"
	LinkRelationTypesPreview                LinkRelationTypes = "preview"
	LinkRelationTypesPrevious               LinkRelationTypes = "previous"
	LinkRelationTypesPrevArchive            LinkRelationTypes = "prev-archive"
	LinkRelationTypesPrivacyPolicy          LinkRelationTypes = "privacy-policy"
	LinkRelationTypesProfile                LinkRelationTypes = "profile"
	LinkRelationTypesPublication            LinkRelationTypes = "publication"
	LinkRelationTypesRelated                LinkRelationTypes = "related"
	LinkRelationTypesRestconf               LinkRelationTypes = "restconf"
	LinkRelationTypesReplies                LinkRelationTypes = "replies"
	LinkRelationTypesRuleinput              LinkRelationTypes = "ruleinput"
	LinkRelationTypesSearch                 LinkRelationTypes = "search"
	LinkRelationTypesSection                LinkRelationTypes = "section"
	LinkRelationTypesSelf                   LinkRelationTypes = "self"
	LinkRelationTypesService                LinkRelationTypes = "service"
	LinkRelationTypesServiceDesc            LinkRelationTypes = "service-desc"
	LinkRelationTypesServiceDoc             LinkRelationTypes = "service-doc"
	LinkRelationTypesServiceMeta            LinkRelationTypes = "service-meta"
	LinkRelationTypesSponsored              LinkRelationTypes = "sponsored"
	LinkRelationTypesStart                  LinkRelationTypes = "start"
	LinkRelationTypesStatus                 LinkRelationTypes = "status"
	LinkRelationTypesStylesheet             LinkRelationTypes = "stylesheet"
	LinkRelationTypesSubsection             LinkRelationTypes = "subsection"
	LinkRelationTypesSuccessorVersion       LinkRelationTypes = "successor-version"
	LinkRelationTypesSunset                 LinkRelationTypes = "sunset"
	LinkRelationTypesTag                    LinkRelationTypes = "tag"
	LinkRelationTypesTermsOfService         LinkRelationTypes = "terms-of-service"
	LinkRelationTypesTimegate               LinkRelationTypes = "timegate"
	LinkRelationTypesTimemap                LinkRelationTypes = "timemap"
	LinkRelationTypesType                   LinkRelationTypes = "type"
	LinkRelationTypesUgc                    LinkRelationTypes = "ugc"
	LinkRelationTypesUp                     LinkRelationTypes = "up"
	LinkRelationTypesVersionHistory         LinkRelationTypes = "version-history"
	LinkRelationTypesVia                    LinkRelationTypes = "via"
	LinkRelationTypesWebmention             LinkRelationTypes = "webmention"
	LinkRelationTypesWorkingCopy            LinkRelationTypes = "working-copy"
	LinkRelationTypesWorkingCopyOf          LinkRelationTypes = "working-copy-of"
)

// LinkType represents codes from http://hl7.org/fhir/ValueSet/link-type
type LinkType string

const (
	LinkTypeReplacedBy LinkType = "replaced-by"
	LinkTypeReplaces   LinkType = "replaces"
	LinkTypeRefer      LinkType = "refer"
	LinkTypeSeealso    LinkType = "seealso"
)

// ListMode represents codes from http://hl7.org/fhir/ValueSet/list-mode
type ListMode string

const (
	ListModeWorking  ListMode = "working"
	ListModeSnapshot ListMode = "snapshot"
	ListModeChanges  ListMode = "changes"
)

// ListStatus represents codes from http://hl7.org/fhir/ValueSet/list-status
type ListStatus string

const (
	ListStatusCurrent        ListStatus = "current"
	ListStatusRetired        ListStatus = "retired"
	ListStatusEnteredInError ListStatus = "entered-in-error"
)

// LocationMode represents codes from http://hl7.org/fhir/ValueSet/location-mode
type LocationMode string

const (
	LocationModeInstance LocationMode = "instance"
	LocationModeKind     LocationMode = "kind"
)

// LocationStatus represents codes from http://hl7.org/fhir/ValueSet/location-status
type LocationStatus string

const (
	LocationStatusActive    LocationStatus = "active"
	LocationStatusSuspended LocationStatus = "suspended"
	LocationStatusInactive  LocationStatus = "inactive"
)
