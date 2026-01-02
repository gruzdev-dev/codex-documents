package fhir

type Bundle struct {
	ImplicitRules *Element
	Language      *Element
	Timestamp     *Element
	Total         *Element
	Type          *Element
	Identifier    *Identifier
	Meta          *Meta
	Timestamp_2   *Instant
}
