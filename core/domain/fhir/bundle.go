package fhir

type Bundle struct {
	ImplicitRules   *Element
	Language        *Element
	Timestamp       *Element
	Total           *Element
	Type            *Element
	Entry           []BundleEntry
	ID              *ID
	Identifier      *Identifier
	ImplicitRules_2 *Uri
	Issues          BundleIssues
	Language_2      *Code
	Link            []BundleLink
	Meta            *Meta
	ResourceType    interface{}
	Signature       *Signature
	Timestamp_2     *Instant
	Total_2         *UnsignedInt
	Type_2          *Code
}
