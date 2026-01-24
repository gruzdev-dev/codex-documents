package domain

import (
	models "github.com/gruzdev-dev/fhir/r5"
)

type ListResponse[T any] struct {
	Items []T
	Total int64
}

type CreateDocumentResult struct {
	Document   *models.DocumentReference
	UploadUrls map[string]string
}
