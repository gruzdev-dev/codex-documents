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
	BirthDate_2        *Date
	Communication      []PractitionerCommunication
	Contained          []PractitionerContainedElem
	DeceasedBoolean_2  *bool
	DeceasedDateTime_2 *string
	Extension          []Extension
	Gender_2           *Code
	ID                 *ID
	Identifier         []Identifier
	ImplicitRules_2    *Uri
	Language_2         *Code
	Meta               *Meta
	ModifierExtension  []Extension
	Name               []HumanName
	Photo              []Attachment
	Qualification      []PractitionerQualification
	ResourceType       interface{}
	Telecom            []ContactPoint
	Text               *Narrative
}
