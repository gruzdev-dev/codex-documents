package fhir

type Patient struct {
	Active                 *Element
	BirthDate              *Element
	DeceasedBoolean        *Element
	DeceasedDateTime       *Element
	Gender                 *Element
	ImplicitRules          *Element
	Language               *Element
	MultipleBirthBoolean   *Element
	MultipleBirthInteger   *Element
	Active_2               *Boolean
	Address                []Address
	DeceasedBoolean_2      *bool
	DeceasedDateTime_2     *string
	Extension              []Extension
	GeneralPractitioner    []Reference
	Identifier             []Identifier
	ManagingOrganization   *Reference
	MaritalStatus          *CodeableConcept
	Meta                   *Meta
	ModifierExtension      []Extension
	MultipleBirthBoolean_2 *bool
	Name                   []HumanName
	Photo                  []Attachment
	Telecom                []ContactPoint
}
