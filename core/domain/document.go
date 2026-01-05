package domain

import "time"

type Document struct {
	ID          string
	PatientID   string
	Title       string
	Status      string
	Category    Codeable
	FileID      string
	ContentType string
	UploadURL   string
	CreatedAt   time.Time
}
