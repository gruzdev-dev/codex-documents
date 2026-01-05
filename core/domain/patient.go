package domain

import "time"

type Patient struct {
	ID        string
	AuthID    string
	Active    bool
	FullName  string
	Gender    string
	BirthDate time.Time
	Telecom   string
}
