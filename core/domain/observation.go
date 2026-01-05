package domain

import "time"

type Observation struct {
	ID            string
	PatientID     string
	Status        string
	Code          Codeable
	Value         float64
	Unit          string
	EffectiveTime time.Time
	Issued        time.Time
}
