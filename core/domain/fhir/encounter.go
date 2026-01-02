package fhir

type Encounter struct {
	ImplicitRules      *Element
	Language           *Element
	PlannedEndDate     *Element
	PlannedStartDate   *Element
	Status             *Element
	Account            []Reference
	ActualPeriod       *Period
	Appointment        []Reference
	BasedOn            []Reference
	CareTeam           []Reference
	Class              []CodeableConcept
	DietPreference     []CodeableConcept
	EpisodeOfCare      []Reference
	Extension          []Extension
	Identifier         []Identifier
	Meta               *Meta
	ModifierExtension  []Extension
	PartOf             *Reference
	PlannedEndDate_2   *DateTime
	PlannedStartDate_2 *DateTime
	Priority           *CodeableConcept
	ServiceProvider    *Reference
	SpecialArrangement []CodeableConcept
	SpecialCourtesy    []CodeableConcept
	Subject            *Reference
	SubjectStatus      *CodeableConcept
	Type               []CodeableConcept
}
