package fhir

type Attachment struct {
	ContentType *Element
	Creation    *Element
	Data        *Element
	Duration    *Element
	Frames      *Element
	Hash        *Element
	Height      *Element
	Language    *Element
	Pages       *Element
	Size        *Element
	Title       *Element
	Url         *Element
	Width       *Element
	Creation_2  *DateTime
	Extension   []Extension
	Frames_2    *PositiveInt
	Height_2    *PositiveInt
	ID          *String
	Pages_2     *PositiveInt
	Title_2     *String
	Width_2     *PositiveInt
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
}

type Range struct {
	Extension []Extension
	High      *Quantity
	ID        *String
	Low       *Quantity
}

type Meta struct {
	LastUpdated   *Element
	Source        *Element
	VersionID     *Element
	Extension     []Extension
	ID            *String
	LastUpdated_2 *Instant
	Security      []Coding
	Tag           []Coding
}

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
}

type Boolean bool

type ContactPoint struct {
	Rank      *Element
	System    *Element
	Use       *Element
	Value     *Element
	Extension []Extension
	ID        *String
	Period    *Period
	Rank_2    *PositiveInt
	Value_2   *String
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
}

type Money struct {
	Currency  *Element
	Value     *Element
	Extension []Extension
	ID        *String
}

type Quantity struct {
	Code       *Element
	Comparator *Element
	System     *Element
	Unit       *Element
	Value      *Element
	Extension  []Extension
	ID         *String
	Unit_2     *String
}

type PositiveInt float64

type Element struct {
	Extension []Extension
	ID        *String
}

type Extension struct {
	Url                  *Element
	ValueBase64Binary    *Element
	ValueBoolean         *Element
	ValueCanonical       *Element
	ValueCode            *Element
	ValueDate            *Element
	ValueDateTime        *Element
	ValueDecimal         *Element
	ValueID              *Element
	ValueInstant         *Element
	ValueInteger         *Element
	ValueInteger64       *Element
	ValueMarkdown        *Element
	ValueOid             *Element
	ValuePositiveInt     *Element
	ValueString          *Element
	ValueTime            *Element
	ValueUnsignedInt     *Element
	ValueUri             *Element
	ValueUrl             *Element
	ValueUuid            *Element
	Extension            []Extension
	ID                   *String
	ValueAddress         *Address
	ValueAttachment      *Attachment
	ValueBase64Binary_2  *string
	ValueBoolean_2       *bool
	ValueCanonical_2     *string
	ValueCode_2          *string
	ValueCodeableConcept *CodeableConcept
	ValueCoding          *Coding
	ValueContactPoint    *ContactPoint
	ValueDate_2          *string
	ValueDateTime_2      *string
	ValueHumanName       *HumanName
	ValueID_2            *string
	ValueIdentifier      *Identifier
	ValueInstant_2       *string
	ValueInteger64_2     *string
	ValueMarkdown_2      *string
	ValueMeta            *Meta
	ValueMoney           *Money
	ValueOid_2           *string
	ValuePeriod          *Period
	ValueQuantity        *Quantity
	ValueRange           *Range
	ValueRatio           *Ratio
	ValueReference       *Reference
	ValueString_2        *string
	ValueTime_2          *string
	ValueUri_2           *string
	ValueUrl_2           *string
	ValueUuid_2          *string
}

type Instant string

type String string

type Period struct {
	End       *Element
	Start     *Element
	End_2     *DateTime
	Extension []Extension
	ID        *String
	Start_2   *DateTime
}

type Coding struct {
	Code           *Element
	Display        *Element
	System         *Element
	UserSelected   *Element
	Version        *Element
	Display_2      *String
	Extension      []Extension
	ID             *String
	UserSelected_2 *Boolean
	Version_2      *String
}

type Ratio struct {
	Denominator *Quantity
	Extension   []Extension
	ID          *String
	Numerator   *Quantity
}

type DateTime string

type Identifier struct {
	System    *Element
	Use       *Element
	Value     *Element
	Assigner  *Reference
	Extension []Extension
	ID        *String
	Period    *Period
	Type      *CodeableConcept
	Value_2   *String
}

type CodeableConcept struct {
	Text      *Element
	Coding    []Coding
	Extension []Extension
	ID        *String
	Text_2    *String
}
