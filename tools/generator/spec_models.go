package generator

type StructureDefinitionBundle struct {
	ResourceType string        `json:"resourceType"`
	ID           string        `json:"id"`
	Type         any           `json:"type,omitempty"`
	Entry        []BundleEntry `json:"entry"`
}

type BundleEntry struct {
	FullUrl  string              `json:"fullUrl"`
	Resource StructureDefinition `json:"resource"`
}

type StructureDefinition struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Description    string   `json:"description,omitempty"`
	Kind           string   `json:"kind"`
	Type           any      `json:"type,omitempty"`
	Abstract       bool     `json:"abstract"`
	BaseDefinition string   `json:"baseDefinition,omitempty"`
	Snapshot       Snapshot `json:"snapshot"`
}

type Snapshot struct {
	Element []ElementDefinition `json:"element"`
}

type ElementDefinition struct {
	ID               string            `json:"id"`
	Path             string            `json:"path"`
	Min              int               `json:"min"`
	Max              string            `json:"max"`
	Type             []ElementDataType `json:"type,omitempty"`
	ContentReference string            `json:"contentReference,omitempty"`
	Short            string            `json:"short"`
	Definition       string            `json:"definition,omitempty"`
	Binding          *Binding          `json:"binding,omitempty"`
	Constraint       []Constraint      `json:"constraint,omitempty"`
	MaxLength        *int              `json:"maxLength,omitempty"`
	Pattern          string            `json:"pattern,omitempty"`
	Fixed            any               `json:"fixed,omitempty"`
}

type Binding struct {
	Strength    string `json:"strength,omitempty"`
	Description string `json:"description,omitempty"`
	ValueSet    string `json:"valueSet,omitempty"`
	Extension   []any  `json:"extension,omitempty"`
}

type Constraint struct {
	Key        string `json:"key"`
	Severity   string `json:"severity"`
	Human      string `json:"human"`
	Expression string `json:"expression"`
	Source     string `json:"source,omitempty"`
}

type ElementDataType struct {
	Code          string   `json:"code"`
	TargetProfile []string `json:"targetProfile,omitempty"`
}
