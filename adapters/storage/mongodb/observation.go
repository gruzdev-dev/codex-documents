package mongodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/gruzdev-dev/codex-documents/core/domain"

	models "github.com/gruzdev-dev/fhir/r5"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type ObservationRepo struct {
	collection *mongo.Collection
}

func NewObservationRepo(db *mongo.Database) *ObservationRepo {
	return &ObservationRepo{
		collection: db.Collection("observations"),
	}
}

func (r *ObservationRepo) Create(ctx context.Context, obs *models.Observation) (*models.Observation, error) {
	_, err := r.collection.InsertOne(ctx, obs)
	if err != nil {
		return nil, fmt.Errorf("failed to insert observation: %w", err)
	}
	return obs, nil
}

func (r *ObservationRepo) GetByID(ctx context.Context, id string) (*models.Observation, error) {
	var obs models.Observation

	filter := bson.M{"id": id}

	err := r.collection.FindOne(ctx, filter).Decode(&obs)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find observation: %w", err)
	}

	return &obs, nil
}

func (r *ObservationRepo) GetByIDs(ctx context.Context, ids []string) ([]models.Observation, error) {
	if len(ids) == 0 {
		return []models.Observation{}, nil
	}

	filter := bson.M{"id": bson.M{"$in": ids}}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find observations: %w", err)
	}
	defer func() { _ = cursor.Close(ctx) }()

	var observations []models.Observation
	if err = cursor.All(ctx, &observations); err != nil {
		return nil, fmt.Errorf("failed to decode observations: %w", err)
	}

	if observations == nil {
		observations = []models.Observation{}
	}

	return observations, nil
}

func (r *ObservationRepo) Update(ctx context.Context, obs *models.Observation) (*models.Observation, error) {
	if obs.Id == nil {
		return nil, domain.ErrObservationIDRequired
	}

	filter := bson.M{"id": *obs.Id}
	update := bson.M{"$set": obs}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update observation: %w", err)
	}

	return obs, nil
}

func (r *ObservationRepo) Delete(ctx context.Context, id string) error {
	filter := bson.M{"id": id}

	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete observation: %w", err)
	}

	return nil
}

func (r *ObservationRepo) Search(ctx context.Context, patientID string, limit, offset int) ([]models.Observation, int64, error) {
	patientRef := fmt.Sprintf("Patient/%s", patientID)
	filter := bson.M{"subject.reference": patientRef}

	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count observations: %w", err)
	}

	findOptions := options.Find().
		SetLimit(int64(limit)).
		SetSkip(int64(offset))

	cursor, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to find observations: %w", err)
	}
	defer func() { _ = cursor.Close(ctx) }()

	var observations []models.Observation
	if err = cursor.All(ctx, &observations); err != nil {
		return nil, 0, fmt.Errorf("failed to decode observations: %w", err)
	}

	if observations == nil {
		observations = []models.Observation{}
	}

	return observations, total, nil
}
