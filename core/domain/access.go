package domain

import "time"

type SharedAccess struct {
	ID        string
	PatientID string
	DoctorID  string
	Token     string
	Expiry    time.Time
	CreatedAt time.Time
}
