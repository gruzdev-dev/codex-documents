package fhir

type DataRequirement struct {
	Limit                  *Element
	MustSupport            []Element
	Type                   *Element
	CodeFilter             []DataRequirementCodeFilter
	DateFilter             []DataRequirementDateFilter
	Extension              []Extension
	ID                     *String
	Limit_2                *PositiveInt
	MustSupport_2          []String
	Profile                []Canonical
	Sort                   []DataRequirementSort
	SubjectCodeableConcept *CodeableConcept
	SubjectReference       *Reference
	Type_2                 *Code
	ValueFilter            []DataRequirementValueFilter
}

type Instant string

type Integer float64

type DocumentReferenceProfile struct {
	ValueCanonical    *Element
	ValueUri          *Element
	Extension         []Extension
	ID                *String
	ModifierExtension []Extension
	ValueCanonical_2  *string
	ValueCoding       *Coding
	ValueUri_2        *string
}

type EncounterContainedElem interface{}

type Annotation struct {
	AuthorString    *Element
	Text            *Element
	Time            *Element
	AuthorReference *Reference
	AuthorString_2  *string
	Extension       []Extension
	ID              *String
	Text_2          *Markdown
	Time_2          *DateTime
}

type Base64Binary string

type DataRequirementDateFilter struct {
	Path              *Element
	SearchParam       *Element
	ValueDateTime     *Element
	Extension         []Extension
	ID                *String
	ModifierExtension []Extension
	Path_2            *String
	SearchParam_2     *String
	ValueDateTime_2   *string
	ValueDuration     *Duration
	ValuePeriod       *Period
}

type HumanNameUse_2 string

type BundleRequest struct {
	IfMatch           *Element
	IfModifiedSince   *Element
	IfNoneExist       *Element
	IfNoneMatch       *Element
	Method            *Element
	Url               *Element
	Extension         []Extension
	ID                *String
	IfMatch_2         *String
	IfModifiedSince_2 *Instant
	IfNoneExist_2     *String
	IfNoneMatch_2     *String
	Method_2          *Code
	ModifierExtension []Extension
	Url_2             *Uri
}

type IdentifierUse_2 string

type RatioRange struct {
	Denominator   *Quantity
	Extension     []Extension
	HighNumerator *Quantity
	ID            *String
	LowNumerator  *Quantity
}

type DosageSafety struct {
	IfExceeded        *Element
	DoseLimit         []DosageSafetyDoseLimit
	Extension         []Extension
	ID                *String
	IfExceeded_2      *String
	ModifierExtension []Extension
}

type DocumentReferenceAttester struct {
	Time              *Element
	Extension         []Extension
	ID                *String
	Mode              CodeableConcept
	ModifierExtension []Extension
	Party             *Reference
	Time_2            *DateTime
}

type AddressUse_2 string

type Boolean bool

type DosageSafetyDoseLimit struct {
	Scope             *Element
	Text              *Element
	ValueInteger      *Element
	Extension         []Extension
	ID                *String
	ModifierExtension []Extension
	Period            *Duration
	Scope_2           *Code
	Text_2            *String
	ValueExpression   *Expression
	ValueInteger_2    *float64
	ValueQuantity     *Quantity
}

type Coding struct {
	Code           *Element
	Display        *Element
	System         *Element
	UserSelected   *Element
	Version        *Element
	Code_2         *Code
	Display_2      *String
	Extension      []Extension
	ID             *String
	System_2       *Uri
	UserSelected_2 *Boolean
	Version_2      *String
}

type Dosage struct {
	AsNeeded              *Element
	PatientInstruction    *Element
	Text                  *Element
	AdditionalInstruction []CodeableConcept
	AsNeeded_2            *Boolean
	AsNeededFor           []CodeableConcept
	Condition             []DosageCondition
	DoseAndRate           []DosageDoseAndRate
	Extension             []Extension
	ID                    *String
	Method                *CodeableConcept
	ModifierExtension     []Extension
	PatientInstruction_2  *String
	Route                 *CodeableConcept
	Safety                *DosageSafety
	Site                  *CodeableConcept
	Text_2                *String
	Timing                *Timing
}

type Count struct {
	Code         *Element
	Comparator   *Element
	System       *Element
	Unit         *Element
	Value        *Element
	Code_2       *Code
	Comparator_2 *CountComparator_2
	Extension    []Extension
	ID           *String
	System_2     *Uri
	Unit_2       *String
	Value_2      *Decimal
}

type QuantityComparator_2 string

type TriggerDefinition struct {
	Name              *Element
	TimingDate        *Element
	TimingDateTime    *Element
	Type              *Element
	Code              *CodeableConcept
	Condition         *Expression
	Data              []DataRequirement
	Extension         []Extension
	ID                *String
	Name_2            *String
	SubscriptionTopic *Canonical
	TimingDate_2      *string
	TimingDateTime_2  *string
	TimingTiming      *Timing
	Type_2            *TriggerDefinitionType_2
}

type PatientLink struct {
	Type              *Element
	Extension         []Extension
	ID                *String
	ModifierExtension []Extension
	Other             Reference
	Type_2            *Code
}

type EncounterBusinessStatus struct {
	EffectiveDate     *Element
	Code              CodeableConcept
	EffectiveDate_2   *DateTime
	Extension         []Extension
	ID                *String
	ModifierExtension []Extension
	Type              *Coding
}

type Extension struct {
	Url                        *Element
	ValueBase64Binary          *Element
	ValueBoolean               *Element
	ValueCanonical             *Element
	ValueCode                  *Element
	ValueDate                  *Element
	ValueDateTime              *Element
	ValueDecimal               *Element
	ValueID                    *Element
	ValueInstant               *Element
	ValueInteger               *Element
	ValueInteger64             *Element
	ValueMarkdown              *Element
	ValueOid                   *Element
	ValuePositiveInt           *Element
	ValueString                *Element
	ValueTime                  *Element
	ValueUnsignedInt           *Element
	ValueUri                   *Element
	ValueUrl                   *Element
	ValueUuid                  *Element
	Extension                  []Extension
	ID                         *String
	Url_2                      *Uri
	ValueAddress               *Address
	ValueAge                   *Age
	ValueAnnotation            *Annotation
	ValueAttachment            *Attachment
	ValueAvailability          *Availability
	ValueBase64Binary_2        *string
	ValueBoolean_2             *bool
	ValueCanonical_2           *string
	ValueCode_2                *string
	ValueCodeableConcept       *CodeableConcept
	ValueCodeableReference     *CodeableReference
	ValueCoding                *Coding
	ValueContactDetail         *ContactDetail
	ValueContactPoint          *ContactPoint
	ValueCount                 *Count
	ValueDataRequirement       *DataRequirement
	ValueDate_2                *string
	ValueDateTime_2            *string
	ValueDecimal_2             *float64
	ValueDistance              *Distance
	ValueDosage                *Dosage
	ValueDuration              *Duration
	ValueExpression            *Expression
	ValueExtendedContactDetail *ExtendedContactDetail
	ValueHumanName             *HumanName
	ValueID_2                  *string
	ValueIdentifier            *Identifier
	ValueInstant_2             *string
	ValueInteger_2             *float64
	ValueInteger64_2           *string
	ValueMarkdown_2            *string
	ValueMeta                  *Meta
	ValueMoney                 *Money
	ValueOid_2                 *string
	ValueParameterDefinition   *ParameterDefinition
	ValuePeriod                *Period
	ValuePositiveInt_2         *float64
	ValueQuantity              *Quantity
	ValueRange                 *Range
	ValueRatio                 *Ratio
	ValueRatioRange            *RatioRange
	ValueReference             *Reference
	ValueRelatedArtifact       *RelatedArtifact
	ValueSampledData           *SampledData
	ValueSignature             *Signature
	ValueString_2              *string
	ValueTime_2                *string
	ValueTiming                *Timing
	ValueTriggerDefinition     *TriggerDefinition
	ValueUnsignedInt_2         *float64
	ValueUri_2                 *string
	ValueUrl_2                 *string
	ValueUsageContext          *UsageContext
	ValueUuid_2                *string
	ValueVirtualServiceDetail  *VirtualServiceDetail
}

type AgeComparator_2 string

type AvailabilityNotAvailableTime struct {
	Description       *Element
	Description_2     *String
	During            *Period
	Extension         []Extension
	ID                *String
	ModifierExtension []Extension
}

type TimingRepeatDurationUnit_2 string

type Duration struct {
	Code         *Element
	Comparator   *Element
	System       *Element
	Unit         *Element
	Value        *Element
	Code_2       *Code
	Comparator_2 *DurationComparator_2
	Extension    []Extension
	ID           *String
	System_2     *Uri
	Unit_2       *String
	Value_2      *Decimal
}

type AvailabilityAvailableTime struct {
	AllDay               *Element
	AvailableEndTime     *Element
	AvailableStartTime   *Element
	DaysOfWeek           []Element
	AllDay_2             *Boolean
	AvailableEndTime_2   *Time
	AvailableStartTime_2 *Time
	DaysOfWeek_2         []Code
	Extension            []Extension
	ID                   *String
	ModifierExtension    []Extension
}

type DataRequirementValueFilter struct {
	Comparator        *Element
	Path              *Element
	SearchParam       *Element
	ValueDateTime     *Element
	Comparator_2      *Code
	Extension         []Extension
	ID                *String
	ModifierExtension []Extension
	Path_2            *String
	SearchParam_2     *String
	ValueDateTime_2   *string
	ValueDuration     *Duration
	ValuePeriod       *Period
}

type BundleLink struct {
	Relation          *Element
	Url               *Element
	Extension         []Extension
	ID                *String
	ModifierExtension []Extension
	Relation_2        *Code
	Url_2             *Uri
}

type Code string

type DistanceComparator_2 string

type UnsignedInt float64

type String string

type DataRequirementCodeFilter struct {
	Path              *Element
	SearchParam       *Element
	Code              []Coding
	Extension         []Extension
	ID                *String
	ModifierExtension []Extension
	Path_2            *String
	SearchParam_2     *String
	ValueSet          *Canonical
}

type HumanName struct {
	Family    *Element
	Given     []Element
	Prefix    []Element
	Suffix    []Element
	Text      *Element
	Use       *Element
	Extension []Extension
	Family_2  *String
	Given_2   []String
	ID        *String
	Period    *Period
	Prefix_2  []String
	Suffix_2  []String
	Text_2    *String
	Use_2     *HumanNameUse_2
}

type Range struct {
	Extension []Extension
	High      *Quantity
	ID        *String
	Low       *Quantity
}

type Ratio struct {
	Denominator *Quantity
	Extension   []Extension
	ID          *String
	Numerator   *Quantity
}

type TimingRepeatPeriodUnit_2 string

type VirtualServiceDetail struct {
	AdditionalInfo               []Element
	AddressString                *Element
	AddressUrl                   *Element
	MaxParticipants              *Element
	SessionKey                   *Element
	AdditionalInfo_2             []Url
	AddressContactPoint          *ContactPoint
	AddressExtendedContactDetail *ExtendedContactDetail
	AddressString_2              *string
	AddressUrl_2                 *string
	ChannelType                  *Coding
	Extension                    []Extension
	ID                           *String
	MaxParticipants_2            *PositiveInt
	SessionKey_2                 *String
}

type NarrativeStatus_2 string

type AddressType_2 string

type EncounterLocation struct {
	Status            *Element
	Extension         []Extension
	Form              *CodeableConcept
	ID                *String
	Location          Reference
	ModifierExtension []Extension
	Period            *Period
	Status_2          *Code
}

type BundleResponseOutcome interface{}

type PractitionerCommunication struct {
	Preferred         *Element
	Extension         []Extension
	ID                *String
	Language          CodeableConcept
	ModifierExtension []Extension
	Preferred_2       *Boolean
}

type PractitionerContainedElem interface{}

type ContactPointUse_2 string

type Age struct {
	Code         *Element
	Comparator   *Element
	System       *Element
	Unit         *Element
	Value        *Element
	Code_2       *Code
	Comparator_2 *AgeComparator_2
	Extension    []Extension
	ID           *String
	System_2     *Uri
	Unit_2       *String
	Value_2      *Decimal
}

type Decimal float64

type Identifier struct {
	System    *Element
	Use       *Element
	Value     *Element
	Assigner  *Reference
	Extension []Extension
	ID        *String
	Period    *Period
	System_2  *Uri
	Type      *CodeableConcept
	Use_2     *IdentifierUse_2
	Value_2   *String
}

type Integer64 string

type ContactPoint struct {
	Rank      *Element
	System    *Element
	Use       *Element
	Value     *Element
	Extension []Extension
	ID        *String
	Period    *Period
	Rank_2    *PositiveInt
	System_2  *ContactPointSystem_2
	Use_2     *ContactPointUse_2
	Value_2   *String
}

type TimingRepeat struct {
	Count             *Element
	CountMax          *Element
	DayOfWeek         []Element
	Duration          *Element
	DurationMax       *Element
	DurationUnit      *Element
	Frequency         *Element
	FrequencyMax      *Element
	Offset            *Element
	Period            *Element
	PeriodMax         *Element
	PeriodUnit        *Element
	TimeOfDay         []Element
	When              []Element
	BoundsDuration    *Duration
	BoundsPeriod      *Period
	BoundsRange       *Range
	Count_2           *PositiveInt
	CountMax_2        *PositiveInt
	DayOfWeek_2       []Code
	Duration_2        *Decimal
	DurationMax_2     *Decimal
	DurationUnit_2    *TimingRepeatDurationUnit_2
	EndOffset         *Quantity
	Extension         []Extension
	Frequency_2       *PositiveInt
	FrequencyMax_2    *PositiveInt
	ID                *String
	ModifierExtension []Extension
	Offset_2          *UnsignedInt
	Period_2          *Decimal
	PeriodMax_2       *Decimal
	PeriodUnit_2      *TimingRepeatPeriodUnit_2
	StartOffset       *Quantity
	TimeOfDay_2       []Time
	When_2            []TimingRepeatWhen_2Elem
}

type Date string

type DurationComparator_2 string

type EncounterAdmission struct {
	AdmitSource            *CodeableConcept
	Destination            *Reference
	DischargeDisposition   *CodeableConcept
	Extension              []Extension
	ID                     *String
	ModifierExtension      []Extension
	Origin                 *Reference
	PreAdmissionIdentifier *Identifier
	ReAdmission            *CodeableConcept
}

type BundleResponse struct {
	Etag              *Element
	LastModified      *Element
	Location          *Element
	Status            *Element
	Etag_2            *String
	Extension         []Extension
	ID                *String
	LastModified_2    *Instant
	Location_2        *Uri
	ModifierExtension []Extension
	Outcome           BundleResponseOutcome
	Status_2          *String
}

type PractitionerQualification struct {
	Code              CodeableConcept
	Extension         []Extension
	ID                *String
	Identifier        []Identifier
	Issuer            *Reference
	ModifierExtension []Extension
	Period            *Period
	Status            *CodeableConcept
}

type ContactDetail struct {
	Name      *Element
	Extension []Extension
	ID        *String
	Name_2    *String
	Telecom   []ContactPoint
}

type TriggerDefinitionType_2 string

type BundleSearch struct {
	Mode              *Element
	Score             *Element
	Extension         []Extension
	ID                *String
	Mode_2            *Code
	ModifierExtension []Extension
	Score_2           *Decimal
}

type BundleEntry struct {
	FullUrl           *Element
	Extension         []Extension
	FullUrl_2         *Uri
	ID                *String
	Link              []BundleLink
	ModifierExtension []Extension
	Request           *BundleRequest
	Resource          BundleEntryResource
	Response          *BundleResponse
	Search            *BundleSearch
}

type Attachment struct {
	ContentType   *Element
	Creation      *Element
	Data          *Element
	Duration      *Element
	Frames        *Element
	Hash          *Element
	Height        *Element
	Language      *Element
	Pages         *Element
	Size          *Element
	Title         *Element
	Url           *Element
	Width         *Element
	ContentType_2 *Code
	Creation_2    *DateTime
	Data_2        *Base64Binary
	Duration_2    *Decimal
	Extension     []Extension
	Frames_2      *PositiveInt
	Hash_2        *Base64Binary
	Height_2      *PositiveInt
	ID            *String
	Language_2    *Code
	Pages_2       *PositiveInt
	Size_2        *Integer64
	Title_2       *String
	Url_2         *Url
	Width_2       *PositiveInt
}

type Url string

type DataRequirementSortDirection_2 string

type Expression struct {
	Description   *Element
	Expression    *Element
	Language      *Element
	Name          *Element
	Reference     *Element
	Description_2 *String
	Expression_2  *String
	Extension     []Extension
	ID            *String
	Language_2    *Code
	Name_2        *Code
	Reference_2   *Uri
}

type ExtendedContactDetail struct {
	Address      *Address
	Extension    []Extension
	ID           *String
	Name         []HumanName
	Organization *Reference
	Period       *Period
	Purpose      *CodeableConcept
	Telecom      []ContactPoint
}

type Quantity struct {
	Code         *Element
	Comparator   *Element
	System       *Element
	Unit         *Element
	Value        *Element
	Code_2       *Code
	Comparator_2 *QuantityComparator_2
	Extension    []Extension
	ID           *String
	System_2     *Uri
	Unit_2       *String
	Value_2      *Decimal
}

type TimingRepeatWhen_2Elem string

type BundleIssues interface{}

type ObservationReferenceRange struct {
	Text              *Element
	Age               *Range
	AppliesTo         []CodeableConcept
	Extension         []Extension
	High              *Quantity
	ID                *String
	Low               *Quantity
	ModifierExtension []Extension
	NormalValue       *CodeableConcept
	Text_2            *Markdown
	Type              *CodeableConcept
}

type EncounterReason struct {
	Extension         []Extension
	ID                *String
	ModifierExtension []Extension
	Use               []CodeableConcept
	Value             []CodeableReference
}

type ID string

type DataRequirementSort struct {
	Direction         *Element
	Path              *Element
	Direction_2       *DataRequirementSortDirection_2
	Extension         []Extension
	ID                *String
	ModifierExtension []Extension
	Path_2            *String
}

type Canonical string

type UsageContext struct {
	Code                 Coding
	Extension            []Extension
	ID                   *String
	ValueCodeableConcept *CodeableConcept
	ValueQuantity        *Quantity
	ValueRange           *Range
	ValueReference       *Reference
}

type ObservationContainedElem interface{}

type ObservationTriggeredBy struct {
	Reason            *Element
	Type              *Element
	Extension         []Extension
	ID                *String
	ModifierExtension []Extension
	Observation       Reference
	Reason_2          *String
	Type_2            *Code
}

type Element struct {
	Extension []Extension
	ID        *String
}

type PatientContact struct {
	Gender            *Element
	AdditionalAddress []Address
	AdditionalName    []HumanName
	Address           *Address
	Extension         []Extension
	Gender_2          *Code
	ID                *String
	ModifierExtension []Extension
	Name              *HumanName
	Organization      *Reference
	Period            *Period
	Relationship      []CodeableConcept
	Role              []CodeableConcept
	Telecom           []ContactPoint
}

type Period struct {
	End       *Element
	Start     *Element
	End_2     *DateTime
	Extension []Extension
	ID        *String
	Start_2   *DateTime
}

type CodeableConcept struct {
	Text      *Element
	Coding    []Coding
	Extension []Extension
	ID        *String
	Text_2    *String
}

type Markdown string

type Meta struct {
	LastUpdated   *Element
	Source        *Element
	VersionID     *Element
	Extension     []Extension
	ID            *String
	LastUpdated_2 *Instant
	Profile       []Canonical
	Security      []Coding
	Source_2      *Uri
	Tag           []Coding
	VersionID_2   *ID
}

type ParameterDefinition struct {
	Documentation   *Element
	Max             *Element
	Min             *Element
	Name            *Element
	Type            *Element
	Use             *Element
	Documentation_2 *String
	Extension       []Extension
	ID              *String
	Max_2           *String
	Min_2           *Integer
	Name_2          *Code
	Profile         *Canonical
	Type_2          *Code
	Use_2           *Code
}

type RelatedArtifact struct {
	ArtifactCanonical   *Element
	ArtifactMarkdown    *Element
	Citation            *Element
	Display             *Element
	Label               *Element
	Type                *Element
	ArtifactAttachment  *Attachment
	ArtifactCanonical_2 *string
	ArtifactMarkdown_2  *string
	ArtifactReference   *Reference
	Citation_2          *Markdown
	Display_2           *String
	Document            *Attachment
	Extension           []Extension
	ID                  *String
	Label_2             *String
	Resource            *Canonical
	ResourceReference   *Reference
	Type_2              *RelatedArtifactType_2
}

type PatientContainedElem interface{}

type EncounterDiagnosis struct {
	Condition         []CodeableReference
	Extension         []Extension
	ID                *String
	ModifierExtension []Extension
	Use               []CodeableConcept
}

type Reference struct {
	Display     *Element
	Reference   *Element
	Type        *Element
	Display_2   *String
	Extension   []Extension
	ID          *String
	Identifier  *Identifier
	Reference_2 *String
	Type_2      *Uri
}

type OperationOutcomeContainedElem interface{}

type Time string

type Signature struct {
	Data           *Element
	SigFormat      *Element
	TargetFormat   *Element
	When           *Element
	Data_2         *Base64Binary
	Extension      []Extension
	ID             *String
	OnBehalfOf     *Reference
	SigFormat_2    *Code
	TargetFormat_2 *Code
	Type           []Coding
	When_2         *Instant
	Who            *Reference
}

type Availability struct {
	AvailableTime    []AvailabilityAvailableTime
	Extension        []Extension
	ID               *String
	NotAvailableTime []AvailabilityNotAvailableTime
	Period           *Period
}

type ContactPointSystem_2 string

type DocumentReferenceRelatesTo struct {
	Code              CodeableConcept
	Extension         []Extension
	ID                *String
	ModifierExtension []Extension
	Target            Reference
}

type BundleEntryResource interface{}

type DosageCondition struct {
	Operation                  *Element
	Text                       *Element
	ValueBase64Binary          *Element
	ValueBoolean               *Element
	ValueCanonical             *Element
	ValueCode                  *Element
	ValueDate                  *Element
	ValueDateTime              *Element
	ValueDecimal               *Element
	ValueID                    *Element
	ValueInstant               *Element
	ValueInteger               *Element
	ValueInteger64             *Element
	ValueMarkdown              *Element
	ValueOid                   *Element
	ValuePositiveInt           *Element
	ValueString                *Element
	ValueTime                  *Element
	ValueUnsignedInt           *Element
	ValueUri                   *Element
	ValueUrl                   *Element
	ValueUuid                  *Element
	Code                       CodeableConcept
	Details                    *CodeableConcept
	Extension                  []Extension
	ID                         *String
	ModifierExtension          []Extension
	Operation_2                *Code
	Text_2                     *String
	ValueAddress               *Address
	ValueAge                   *Age
	ValueAnnotation            *Annotation
	ValueAttachment            *Attachment
	ValueAvailability          *Availability
	ValueBase64Binary_2        *string
	ValueBoolean_2             *bool
	ValueCanonical_2           *string
	ValueCode_2                *string
	ValueCodeableConcept       *CodeableConcept
	ValueCodeableReference     *CodeableReference
	ValueCoding                *Coding
	ValueContactDetail         *ContactDetail
	ValueContactPoint          *ContactPoint
	ValueCount                 *Count
	ValueDataRequirement       *DataRequirement
	ValueDate_2                *string
	ValueDateTime_2            *string
	ValueDecimal_2             *float64
	ValueDistance              *Distance
	ValueDosage                *Dosage
	ValueDuration              *Duration
	ValueExpression            *Expression
	ValueExtendedContactDetail *ExtendedContactDetail
	ValueHumanName             *HumanName
	ValueID_2                  *string
	ValueIdentifier            *Identifier
	ValueInstant_2             *string
	ValueInteger_2             *float64
	ValueInteger64_2           *string
	ValueMarkdown_2            *string
	ValueMeta                  *Meta
	ValueMoney                 *Money
	ValueOid_2                 *string
	ValueParameterDefinition   *ParameterDefinition
	ValuePeriod                *Period
	ValuePositiveInt_2         *float64
	ValueQuantity              *Quantity
	ValueRange                 *Range
	ValueRatio                 *Ratio
	ValueRatioRange            *RatioRange
	ValueReference             *Reference
	ValueRelatedArtifact       *RelatedArtifact
	ValueSampledData           *SampledData
	ValueSignature             *Signature
	ValueString_2              *string
	ValueTime_2                *string
	ValueTiming                *Timing
	ValueTriggerDefinition     *TriggerDefinition
	ValueUnsignedInt_2         *float64
	ValueUri_2                 *string
	ValueUrl_2                 *string
	ValueUsageContext          *UsageContext
	ValueUuid_2                *string
	ValueVirtualServiceDetail  *VirtualServiceDetail
}

type PatientCommunication struct {
	Preferred         *Element
	Extension         []Extension
	ID                *String
	Language          CodeableConcept
	ModifierExtension []Extension
	Preferred_2       *Boolean
}

type Narrative struct {
	Status    *Element
	Div       NarrativeDiv
	Extension []Extension
	ID        *String
	Status_2  *NarrativeStatus_2
}

type OperationOutcomeIssue struct {
	Code              *Element
	Diagnostics       *Element
	Expression        []Element
	Location          []Element
	Severity          *Element
	Code_2            *Code
	Details           *CodeableConcept
	Diagnostics_2     *String
	Expression_2      []String
	Extension         []Extension
	ID                *String
	Location_2        []String
	ModifierExtension []Extension
	Severity_2        *Code
}

type Uri string

type CountComparator_2 string

type NarrativeDiv interface{}

type EncounterParticipant struct {
	Actor             *Reference
	Extension         []Extension
	ID                *String
	ModifierExtension []Extension
	Period            *Period
	Type              []CodeableConcept
}

type Distance struct {
	Code         *Element
	Comparator   *Element
	System       *Element
	Unit         *Element
	Value        *Element
	Code_2       *Code
	Comparator_2 *DistanceComparator_2
	Extension    []Extension
	ID           *String
	System_2     *Uri
	Unit_2       *String
	Value_2      *Decimal
}

type DateTime string

type Address struct {
	City         *Element
	Country      *Element
	District     *Element
	Line         []Element
	PostalCode   *Element
	State        *Element
	Text         *Element
	Type         *Element
	Use          *Element
	City_2       *String
	Country_2    *String
	District_2   *String
	Extension    []Extension
	ID           *String
	Line_2       []String
	Period       *Period
	PostalCode_2 *String
	State_2      *String
	Text_2       *String
	Type_2       *AddressType_2
	Use_2        *AddressUse_2
}

type RelatedArtifactType_2 string

type DocumentReferenceContent struct {
	Attachment        Attachment
	Extension         []Extension
	ID                *String
	ModifierExtension []Extension
	Profile           []DocumentReferenceProfile
}

type PositiveInt float64

type SampledData struct {
	Data           *Element
	Dimensions     *Element
	Factor         *Element
	Interval       *Element
	IntervalUnit   *Element
	LowerLimit     *Element
	Offsets        *Element
	UpperLimit     *Element
	CodeMap        *Canonical
	Data_2         *String
	Dimensions_2   *PositiveInt
	Extension      []Extension
	Factor_2       *Decimal
	ID             *String
	Interval_2     *Decimal
	IntervalUnit_2 *Code
	LowerLimit_2   *Decimal
	Offsets_2      *String
	Origin         Quantity
	UpperLimit_2   *Decimal
}

type DosageDoseAndRate struct {
	DoseExpression    *Expression
	DoseQuantity      *Quantity
	DoseRange         *Range
	Extension         []Extension
	ID                *String
	ModifierExtension []Extension
	RateExpression    *Expression
	RateQuantity      *Quantity
	RateRange         *Range
	RateRatio         *Ratio
	Type              *CodeableConcept
}

type DocumentReferenceContainedElem interface{}

type CodeableReference struct {
	Concept   *CodeableConcept
	Extension []Extension
	ID        *String
	Reference *Reference
}

type Money struct {
	Currency   *Element
	Value      *Element
	Currency_2 *Code
	Extension  []Extension
	ID         *String
	Value_2    *Decimal
}

type Timing struct {
	Event             []Element
	Code              *CodeableConcept
	Event_2           []DateTime
	Extension         []Extension
	ID                *String
	ModifierExtension []Extension
	Repeat            *TimingRepeat
}

type ObservationComponent struct {
	ValueBoolean         *Element
	ValueDateTime        *Element
	ValueInteger         *Element
	ValueString          *Element
	ValueTime            *Element
	Code                 CodeableConcept
	DataAbsentReason     *CodeableConcept
	Extension            []Extension
	ID                   *String
	Interpretation       []CodeableConcept
	ModifierExtension    []Extension
	ReferenceRange       []ObservationReferenceRange
	ValueAttachment      *Attachment
	ValueBoolean_2       *bool
	ValueCodeableConcept *CodeableConcept
	ValueDateTime_2      *string
	ValueInteger_2       *float64
	ValuePeriod          *Period
	ValueQuantity        *Quantity
	ValueRange           *Range
	ValueRatio           *Ratio
	ValueSampledData     *SampledData
	ValueString_2        *string
	ValueTime_2          *string
}
