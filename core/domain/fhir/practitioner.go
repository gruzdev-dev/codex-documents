package fhir

type Practitioner struct {
	Active             *Element
	BirthDate          *Element
	DeceasedBoolean    *Element
	DeceasedDateTime   *Element
	Gender             *Element
	ImplicitRules      *Element
	Language           *Element
	Active_2           *Boolean
	Address            []Address
	DeceasedBoolean_2  *bool
	DeceasedDateTime_2 *string
	Extension          []Extension
	Identifier         []Identifier
	Meta               *Meta
	ModifierExtension  []Extension
	Name               []HumanName
	Photo              []Attachment
	Telecom            []ContactPoint
}
