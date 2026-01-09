package mongodb

import (
	"context"
	"errors"
	"fmt"

	"codex-documents/core/domain"

	models "github.com/gruzdev-dev/fhir/r5"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DocumentRepo struct {
	collection *mongo.Collection
}

func NewDocumentRepo(db *mongo.Database) *DocumentRepo {
	return &DocumentRepo{
		collection: db.Collection("document_references"),
	}
}

func (r *DocumentRepo) Create(ctx context.Context, doc *models.DocumentReference) (*models.DocumentReference, error) {
	_, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return nil, fmt.Errorf("failed to insert document: %w", err)
	}
	return doc, nil
}

func (r *DocumentRepo) GetByID(ctx context.Context, id string) (*models.DocumentReference, error) {
	var doc models.DocumentReference

	filter := bson.M{"id": id}

	err := r.collection.FindOne(ctx, filter).Decode(&doc)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find document: %w", err)
	}

	return &doc, nil
}

func (r *DocumentRepo) Update(ctx context.Context, doc *models.DocumentReference) (*models.DocumentReference, error) {
	if doc.Id == nil {
		return nil, domain.ErrDocumentIDRequired
	}

	filter := bson.M{"id": *doc.Id}
	update := bson.M{"$set": doc}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update document: %w", err)
	}

	return doc, nil
}

func (r *DocumentRepo) Delete(ctx context.Context, id string) error {
	filter := bson.M{"id": id}

	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete document: %w", err)
	}

	return nil
}

func (r *DocumentRepo) Search(ctx context.Context, patientID string, limit, offset int) ([]models.DocumentReference, int64, error) {
	// Фильтруем по subject.reference, который имеет формат "Patient/{patientID}"
	patientRef := fmt.Sprintf("Patient/%s", patientID)
	filter := bson.M{"subject.reference": patientRef}

	// Получаем общее количество для Bundle.total
	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count documents: %w", err)
	}

	// Настраиваем пагинацию
	findOptions := options.Find().
		SetLimit(int64(limit)).
		SetSkip(int64(offset))

	cursor, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to find documents: %w", err)
	}
	defer cursor.Close(ctx)

	var documents []models.DocumentReference
	if err = cursor.All(ctx, &documents); err != nil {
		return nil, 0, fmt.Errorf("failed to decode documents: %w", err)
	}

	// Если ничего не нашли, возвращаем пустой слайс вместо nil
	if documents == nil {
		documents = []models.DocumentReference{}
	}

	return documents, total, nil
}
