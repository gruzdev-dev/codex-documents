package fhir

type DocumentReference struct {
	Date              *Element
	Description       *Element
	DocStatus         *Element
	ImplicitRules     *Element
	Language          *Element
	Status            *Element
	Version           *Element
	Author            []Reference
	BasedOn           []Reference
	Category          []CodeableConcept
	Context           []Reference
	Custodian         *Reference
	Date_2            *DateTime
	Extension         []Extension
	FacilityType      *CodeableConcept
	Identifier        []Identifier
	Meta              *Meta
	Modality          []CodeableConcept
	ModifierExtension []Extension
	Period            *Period
	PracticeSetting   *CodeableConcept
	Related           []Reference
	SecurityLabel     []CodeableConcept
	Subject           *Reference
	Type              *CodeableConcept
	Version_2         *String
}
