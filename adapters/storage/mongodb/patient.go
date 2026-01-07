package mongodb

import (
	"context"
	"errors"
	"fmt"

	"codex-documents/core/domain"

	models "github.com/gruzdev-dev/fhir/r5"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type PatientStore struct {
	collection *mongo.Collection
}

func NewPatientStore(db *mongo.Database) *PatientStore {
	return &PatientStore{
		collection: db.Collection("patients"),
	}
}

func (s *PatientStore) Create(ctx context.Context, patient *models.Patient) error {
	_, err := s.collection.InsertOne(ctx, patient)
	if err != nil {
		return fmt.Errorf("failed to insert patient: %w", err)
	}
	return nil
}

func (s *PatientStore) GetByID(ctx context.Context, id string) (*models.Patient, error) {
	var patient models.Patient

	filter := bson.M{"id": id}

	err := s.collection.FindOne(ctx, filter).Decode(&patient)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find patient: %w", err)
	}

	return &patient, nil
}

func (s *PatientStore) Update(ctx context.Context, patient *models.Patient) error {
	if patient.Id == nil {
		return domain.ErrPatientIDRequired
	}

	filter := bson.M{"id": *patient.Id}
	update := bson.M{"$set": patient}

	_, err := s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update patient: %w", err)
	}

	return nil
}
