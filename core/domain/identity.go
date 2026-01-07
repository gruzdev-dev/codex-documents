package domain

type Identity struct {
	UserID    string
	PatientID string
	Scopes    []string
}

func (i *Identity) HasScope(scope string) bool {
	for _, s := range i.Scopes {
		if s == scope {
			return true
		}
	}
	return false
}

func (i *Identity) IsPatient(id string) bool {
	return i.PatientID == id
}