package services

import (
	"context"
	"errors"
	"testing"

	"github.com/gruzdev-dev/codex-documents/core/domain"
	"github.com/gruzdev-dev/codex-documents/core/ports"
	"github.com/gruzdev-dev/codex-documents/core/validator"
	"github.com/gruzdev-dev/codex-documents/pkg/identity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	models "github.com/gruzdev-dev/fhir/r5"
)

const (
	testObsID = "obs-123"
)

func createTestObservation(id, patientID string) *models.Observation {
	patientRef := "Patient/" + patientID
	return &models.Observation{
		ResourceType: "Observation",
		Id:           strPtr(id),
		Status:       "final",
		Subject: &models.Reference{
			Reference: &patientRef,
		},
	}
}

func createTestObservationWithDerivedFrom(id, patientID string, docIDs []string) *models.Observation {
	obs := createTestObservation(id, patientID)
	derivedFrom := make([]models.Reference, 0, len(docIDs))
	for _, docID := range docIDs {
		ref := "DocumentReference/" + docID
		derivedFrom = append(derivedFrom, models.Reference{
			Reference: &ref,
		})
	}
	obs.DerivedFrom = derivedFrom
	return obs
}

func TestObservationService_Create(t *testing.T) {
	tests := []struct {
		name           string
		obs            *models.Observation
		setupMocks     func(*ports.MockObservationRepository, *ports.MockDocumentRepository)
		setupContext   func() context.Context
		expectedError  error
		validateResult func(*testing.T, *models.Observation, error)
	}{
		{
			name: "success path - without derivedFrom",
			obs: &models.Observation{
				ResourceType: "Observation",
				Status:       "final",
			},
			setupMocks: func(obsRepo *ports.MockObservationRepository, docRepo *ports.MockDocumentRepository) {
				obsRepo.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, obs *models.Observation) (*models.Observation, error) {
						require.NotNil(t, obs.Id)
						require.NotEmpty(t, *obs.Id)
						require.NotNil(t, obs.Subject)
						require.Equal(t, "Patient/"+testPatientID, *obs.Subject.Reference)
						return obs, nil
					})
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				require.NoError(t, err)
				require.NotNil(t, obs)
				require.NotNil(t, obs.Id)
				require.NotEmpty(t, *obs.Id)
			},
		},
		{
			name: "success path - with valid derivedFrom",
			obs: createTestObservationWithDerivedFrom("", testPatientID, []string{testDocID}),
			setupMocks: func(obsRepo *ports.MockObservationRepository, docRepo *ports.MockDocumentRepository) {
				doc := createTestDocument(testDocID, testPatientID)
				docRepo.EXPECT().
					GetByID(gomock.Any(), testDocID).
					Return(doc, nil)
				obsRepo.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(&models.Observation{}, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				require.NoError(t, err)
				require.NotNil(t, obs)
			},
		},
		{
			name: "validation error - nil observation",
			obs:  nil,
			setupMocks: func(*ports.MockObservationRepository, *ports.MockDocumentRepository) {},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrInvalidInput,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				assert.Error(t, err)
				assert.Nil(t, obs)
				assert.ErrorIs(t, err, domain.ErrInvalidInput)
			},
		},
		{
			name: "error - ID already provided",
			obs: &models.Observation{
				ResourceType: "Observation",
				Id:           strPtr(testObsID),
				Status:       "final",
			},
			setupMocks: func(*ports.MockObservationRepository, *ports.MockDocumentRepository) {},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrInvalidInput,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				assert.Error(t, err)
				assert.Nil(t, obs)
				assert.ErrorIs(t, err, domain.ErrInvalidInput)
				assert.Contains(t, err.Error(), "observation ID must not be provided")
			},
		},
		{
			name: "error - no identity",
			obs: &models.Observation{
				ResourceType: "Observation",
				Status:       "final",
			},
			setupMocks: func(*ports.MockObservationRepository, *ports.MockDocumentRepository) {},
			setupContext: func() context.Context {
				return context.Background()
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				assert.Error(t, err)
				assert.Nil(t, obs)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name: "error - temporary token",
			obs: &models.Observation{
				ResourceType: "Observation",
				Status:       "final",
			},
			setupMocks: func(*ports.MockObservationRepository, *ports.MockDocumentRepository) {},
			setupContext: func() context.Context {
				id := createTestIdentity("", "", []string{})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				assert.Error(t, err)
				assert.Nil(t, obs)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name: "error - no write scope",
			obs: &models.Observation{
				ResourceType: "Observation",
				Status:       "final",
			},
			setupMocks: func(*ports.MockObservationRepository, *ports.MockDocumentRepository) {},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				assert.Error(t, err)
				assert.Nil(t, obs)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name: "error - invalid derivedFrom - not DocumentReference",
			obs: func() *models.Observation {
				obs := &models.Observation{
					ResourceType: "Observation",
					Status:       "final",
				}
				ref := "Patient/" + testPatientID
				obs.DerivedFrom = []models.Reference{
					{Reference: &ref},
				}
				return obs
			}(),
			setupMocks: func(*ports.MockObservationRepository, *ports.MockDocumentRepository) {},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrInvalidDerivedFromRef,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				assert.Error(t, err)
				assert.Nil(t, obs)
				assert.Equal(t, domain.ErrInvalidDerivedFromRef, err)
			},
		},
		{
			name: "error - derivedFrom document not found",
			obs: createTestObservationWithDerivedFrom("", testPatientID, []string{testDocID}),
			setupMocks: func(obsRepo *ports.MockObservationRepository, docRepo *ports.MockDocumentRepository) {
				docRepo.EXPECT().
					GetByID(gomock.Any(), testDocID).
					Return(nil, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrDerivedFromDocNotFound,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				assert.Error(t, err)
				assert.Nil(t, obs)
				assert.Equal(t, domain.ErrDerivedFromDocNotFound, err)
			},
		},
		{
			name: "error - derivedFrom document belongs to different patient",
			obs: createTestObservationWithDerivedFrom("", testPatientID, []string{testDocID}),
			setupMocks: func(obsRepo *ports.MockObservationRepository, docRepo *ports.MockDocumentRepository) {
				doc := createTestDocument(testDocID, "other-patient")
				docRepo.EXPECT().
					GetByID(gomock.Any(), testDocID).
					Return(doc, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				assert.Error(t, err)
				assert.Nil(t, obs)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name: "error - repository error",
			obs: &models.Observation{
				ResourceType: "Observation",
				Status:       "final",
			},
			setupMocks: func(obsRepo *ports.MockObservationRepository, docRepo *ports.MockDocumentRepository) {
				obsRepo.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("database error"))
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrInternal,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				assert.Error(t, err)
				assert.Nil(t, obs)
				assert.ErrorIs(t, err, domain.ErrInternal)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			obsRepo := ports.NewMockObservationRepository(ctrl)
			docRepo := ports.NewMockDocumentRepository(ctrl)
			validator := validator.NewObservationValidator()

			tt.setupMocks(obsRepo, docRepo)

			service := NewObservationService(obsRepo, docRepo, validator)

			ctx := tt.setupContext()
			result, err := service.Create(ctx, tt.obs)

			if tt.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if tt.validateResult != nil {
				tt.validateResult(t, result, err)
			}
		})
	}
}

func TestObservationService_Get(t *testing.T) {
	tests := []struct {
		name           string
		obsID          string
		setupMocks     func(*ports.MockObservationRepository)
		setupContext   func() context.Context
		expectedError  error
		validateResult func(*testing.T, *models.Observation, error)
	}{
		{
			name: "success path - owner with read scope",
			obsID: testObsID,
			setupMocks: func(repo *ports.MockObservationRepository) {
				obs := createTestObservation(testObsID, testPatientID)
				repo.EXPECT().
					GetByID(gomock.Any(), testObsID).
					Return(obs, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				require.NoError(t, err)
				require.NotNil(t, obs)
				assert.Equal(t, testObsID, *obs.Id)
			},
		},
		{
			name: "success path - access via resource scope",
			obsID: testObsID,
			setupMocks: func(repo *ports.MockObservationRepository) {
				obs := createTestObservation(testObsID, "other-patient")
				repo.EXPECT().
					GetByID(gomock.Any(), testObsID).
					Return(obs, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"docs:observation:" + testObsID + ":read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				require.NoError(t, err)
				require.NotNil(t, obs)
			},
		},
		{
			name: "error - empty ID",
			setupMocks: func(*ports.MockObservationRepository) {},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrObservationIDRequired,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				assert.Error(t, err)
				assert.Nil(t, obs)
				assert.Equal(t, domain.ErrObservationIDRequired, err)
			},
		},
		{
			name: "error - no identity",
			obsID: testObsID,
			setupMocks: func(*ports.MockObservationRepository) {},
			setupContext: func() context.Context {
				return context.Background()
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				assert.Error(t, err)
				assert.Nil(t, obs)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name: "error - not found",
			obsID: testObsID,
			setupMocks: func(repo *ports.MockObservationRepository) {
				repo.EXPECT().
					GetByID(gomock.Any(), testObsID).
					Return(nil, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrObservationNotFound,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				assert.Error(t, err)
				assert.Nil(t, obs)
				assert.Equal(t, domain.ErrObservationNotFound, err)
			},
		},
		{
			name: "error - access denied",
			obsID: testObsID,
			setupMocks: func(repo *ports.MockObservationRepository) {
				obs := createTestObservation(testObsID, "other-patient")
				repo.EXPECT().
					GetByID(gomock.Any(), testObsID).
					Return(obs, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				assert.Error(t, err)
				assert.Nil(t, obs)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name: "error - repository error",
			obsID: testObsID,
			setupMocks: func(repo *ports.MockObservationRepository) {
				repo.EXPECT().
					GetByID(gomock.Any(), testObsID).
					Return(nil, errors.New("database error"))
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrInternal,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				assert.Error(t, err)
				assert.Nil(t, obs)
				assert.ErrorIs(t, err, domain.ErrInternal)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			obsRepo := ports.NewMockObservationRepository(ctrl)
			docRepo := ports.NewMockDocumentRepository(ctrl)
			validator := validator.NewObservationValidator()

			tt.setupMocks(obsRepo)

			service := NewObservationService(obsRepo, docRepo, validator)

			ctx := tt.setupContext()
			result, err := service.Get(ctx, tt.obsID)

			if tt.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if tt.validateResult != nil {
				tt.validateResult(t, result, err)
			}
		})
	}
}

func TestObservationService_Update(t *testing.T) {
	tests := []struct {
		name           string
		obs            *models.Observation
		setupMocks     func(*ports.MockObservationRepository, *ports.MockDocumentRepository)
		setupContext   func() context.Context
		expectedError  error
		validateResult func(*testing.T, *models.Observation, error)
	}{
		{
			name: "success path - without derivedFrom change",
			obs: &models.Observation{
				ResourceType: "Observation",
				Id:           strPtr(testObsID),
				Status:       "final",
			},
			setupMocks: func(obsRepo *ports.MockObservationRepository, docRepo *ports.MockDocumentRepository) {
				existing := createTestObservation(testObsID, testPatientID)
				obsRepo.EXPECT().
					GetByID(gomock.Any(), testObsID).
					Return(existing, nil)
				obsRepo.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Return(&models.Observation{}, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				require.NoError(t, err)
				require.NotNil(t, obs)
			},
		},
		{
			name: "success path - with derivedFrom change",
			obs: createTestObservationWithDerivedFrom(testObsID, testPatientID, []string{testDocID}),
			setupMocks: func(obsRepo *ports.MockObservationRepository, docRepo *ports.MockDocumentRepository) {
				existing := createTestObservation(testObsID, testPatientID)
				obsRepo.EXPECT().
					GetByID(gomock.Any(), testObsID).
					Return(existing, nil)
				doc := createTestDocument(testDocID, testPatientID)
				docRepo.EXPECT().
					GetByID(gomock.Any(), testDocID).
					Return(doc, nil)
				obsRepo.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Return(&models.Observation{}, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				require.NoError(t, err)
				require.NotNil(t, obs)
			},
		},
		{
			name: "error - no ID",
			obs: &models.Observation{
				ResourceType: "Observation",
				Status:       "final",
			},
			setupMocks: func(*ports.MockObservationRepository, *ports.MockDocumentRepository) {},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrObservationIDRequired,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				assert.Error(t, err)
				assert.Nil(t, obs)
				assert.Equal(t, domain.ErrObservationIDRequired, err)
			},
		},
		{
			name: "error - no identity",
			obs: &models.Observation{
				ResourceType: "Observation",
				Id:           strPtr(testObsID),
				Status:       "final",
			},
			setupMocks: func(*ports.MockObservationRepository, *ports.MockDocumentRepository) {},
			setupContext: func() context.Context {
				return context.Background()
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				assert.Error(t, err)
				assert.Nil(t, obs)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name: "error - temporary token",
			obs: &models.Observation{
				ResourceType: "Observation",
				Id:           strPtr(testObsID),
				Status:       "final",
			},
			setupMocks: func(*ports.MockObservationRepository, *ports.MockDocumentRepository) {},
			setupContext: func() context.Context {
				id := createTestIdentity("", "", []string{})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrTmpTokenForbidden,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				assert.Error(t, err)
				assert.Nil(t, obs)
				assert.Equal(t, domain.ErrTmpTokenForbidden, err)
			},
		},
		{
			name: "error - no write scope",
			obs: &models.Observation{
				ResourceType: "Observation",
				Id:           strPtr(testObsID),
				Status:       "final",
			},
			setupMocks: func(*ports.MockObservationRepository, *ports.MockDocumentRepository) {},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				assert.Error(t, err)
				assert.Nil(t, obs)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name: "error - not found",
			obs: &models.Observation{
				ResourceType: "Observation",
				Id:           strPtr(testObsID),
				Status:       "final",
			},
			setupMocks: func(obsRepo *ports.MockObservationRepository, docRepo *ports.MockDocumentRepository) {
				obsRepo.EXPECT().
					GetByID(gomock.Any(), testObsID).
					Return(nil, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrObservationNotFound,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				assert.Error(t, err)
				assert.Nil(t, obs)
				assert.Equal(t, domain.ErrObservationNotFound, err)
			},
		},
		{
			name: "error - not owner",
			obs: &models.Observation{
				ResourceType: "Observation",
				Id:           strPtr(testObsID),
				Status:       "final",
			},
			setupMocks: func(obsRepo *ports.MockObservationRepository, docRepo *ports.MockDocumentRepository) {
				existing := createTestObservation(testObsID, "other-patient")
				obsRepo.EXPECT().
					GetByID(gomock.Any(), testObsID).
					Return(existing, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				assert.Error(t, err)
				assert.Nil(t, obs)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name: "success path - validation passes",
			obs: &models.Observation{
				ResourceType: "Observation",
				Id:           strPtr(testObsID),
				Status:       "final",
			},
			setupMocks: func(obsRepo *ports.MockObservationRepository, docRepo *ports.MockDocumentRepository) {
				existing := createTestObservation(testObsID, testPatientID)
				obsRepo.EXPECT().
					GetByID(gomock.Any(), testObsID).
					Return(existing, nil)
				obsRepo.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Return(&models.Observation{}, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				require.NoError(t, err)
				require.NotNil(t, obs)
			},
		},
		{
			name: "error - invalid new derivedFrom",
			obs: func() *models.Observation {
				obs := &models.Observation{
					ResourceType: "Observation",
					Id:           strPtr(testObsID),
					Status:       "final",
				}
				ref := "Patient/" + testPatientID
				obs.DerivedFrom = []models.Reference{
					{Reference: &ref},
				}
				return obs
			}(),
			setupMocks: func(obsRepo *ports.MockObservationRepository, docRepo *ports.MockDocumentRepository) {
				existing := createTestObservation(testObsID, testPatientID)
				obsRepo.EXPECT().
					GetByID(gomock.Any(), testObsID).
					Return(existing, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrInvalidDerivedFromRef,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				assert.Error(t, err)
				assert.Nil(t, obs)
				assert.Equal(t, domain.ErrInvalidDerivedFromRef, err)
			},
		},
		{
			name: "error - repository update error",
			obs: &models.Observation{
				ResourceType: "Observation",
				Id:           strPtr(testObsID),
				Status:       "final",
			},
			setupMocks: func(obsRepo *ports.MockObservationRepository, docRepo *ports.MockDocumentRepository) {
				existing := createTestObservation(testObsID, testPatientID)
				obsRepo.EXPECT().
					GetByID(gomock.Any(), testObsID).
					Return(existing, nil)
				obsRepo.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("database error"))
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrInternal,
			validateResult: func(t *testing.T, obs *models.Observation, err error) {
				assert.Error(t, err)
				assert.Nil(t, obs)
				assert.ErrorIs(t, err, domain.ErrInternal)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			obsRepo := ports.NewMockObservationRepository(ctrl)
			docRepo := ports.NewMockDocumentRepository(ctrl)
			validator := validator.NewObservationValidator()

			tt.setupMocks(obsRepo, docRepo)

			service := NewObservationService(obsRepo, docRepo, validator)

			ctx := tt.setupContext()
			result, err := service.Update(ctx, tt.obs)

			if tt.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if tt.validateResult != nil {
				tt.validateResult(t, result, err)
			}
		})
	}
}

func TestObservationService_Delete(t *testing.T) {
	tests := []struct {
		name           string
		obsID          string
		setupMocks     func(*ports.MockObservationRepository)
		setupContext   func() context.Context
		expectedError  error
		validateResult func(*testing.T, error)
	}{
		{
			name: "success path",
			obsID: testObsID,
			setupMocks: func(repo *ports.MockObservationRepository) {
				obs := createTestObservation(testObsID, testPatientID)
				repo.EXPECT().
					GetByID(gomock.Any(), testObsID).
					Return(obs, nil)
				repo.EXPECT().
					Delete(gomock.Any(), testObsID).
					Return(nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "error - no identity",
			obsID: testObsID,
			setupMocks: func(*ports.MockObservationRepository) {},
			setupContext: func() context.Context {
				return context.Background()
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, err error) {
				assert.Error(t, err)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name: "error - temporary token",
			obsID: testObsID,
			setupMocks: func(*ports.MockObservationRepository) {},
			setupContext: func() context.Context {
				id := createTestIdentity("", "", []string{})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrTmpTokenForbidden,
			validateResult: func(t *testing.T, err error) {
				assert.Error(t, err)
				assert.Equal(t, domain.ErrTmpTokenForbidden, err)
			},
		},
		{
			name: "error - no write scope",
			obsID: testObsID,
			setupMocks: func(*ports.MockObservationRepository) {},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, err error) {
				assert.Error(t, err)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name: "error - not found",
			obsID: testObsID,
			setupMocks: func(repo *ports.MockObservationRepository) {
				repo.EXPECT().
					GetByID(gomock.Any(), testObsID).
					Return(nil, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrObservationNotFound,
			validateResult: func(t *testing.T, err error) {
				assert.Error(t, err)
				assert.Equal(t, domain.ErrObservationNotFound, err)
			},
		},
		{
			name: "error - not owner",
			obsID: testObsID,
			setupMocks: func(repo *ports.MockObservationRepository) {
				obs := createTestObservation(testObsID, "other-patient")
				repo.EXPECT().
					GetByID(gomock.Any(), testObsID).
					Return(obs, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, err error) {
				assert.Error(t, err)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name: "error - repository error",
			obsID: testObsID,
			setupMocks: func(repo *ports.MockObservationRepository) {
				obs := createTestObservation(testObsID, testPatientID)
				repo.EXPECT().
					GetByID(gomock.Any(), testObsID).
					Return(obs, nil)
				repo.EXPECT().
					Delete(gomock.Any(), testObsID).
					Return(errors.New("database error"))
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrInternal,
			validateResult: func(t *testing.T, err error) {
				assert.Error(t, err)
				assert.ErrorIs(t, err, domain.ErrInternal)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			obsRepo := ports.NewMockObservationRepository(ctrl)
			docRepo := ports.NewMockDocumentRepository(ctrl)
			validator := validator.NewObservationValidator()

			tt.setupMocks(obsRepo)

			service := NewObservationService(obsRepo, docRepo, validator)

			ctx := tt.setupContext()
			err := service.Delete(ctx, tt.obsID)

			if tt.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if tt.validateResult != nil {
				tt.validateResult(t, err)
			}
		})
	}
}

func TestObservationService_List(t *testing.T) {
	tests := []struct {
		name           string
		patientID      string
		limit          int
		offset         int
		setupMocks     func(*ports.MockObservationRepository)
		setupContext   func() context.Context
		expectedError  error
		validateResult func(*testing.T, *domain.ListResponse[models.Observation], error)
	}{
		{
			name:      "success path",
			patientID: testPatientID,
			limit:     10,
			offset:    0,
			setupMocks: func(repo *ports.MockObservationRepository) {
				obs := []models.Observation{
					*createTestObservation("obs-1", testPatientID),
					*createTestObservation("obs-2", testPatientID),
				}
				repo.EXPECT().
					Search(gomock.Any(), testPatientID, 10, 0).
					Return(obs, int64(2), nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, result *domain.ListResponse[models.Observation], err error) {
				require.NoError(t, err)
				require.NotNil(t, result)
				assert.Len(t, result.Items, 2)
				assert.Equal(t, int64(2), result.Total)
			},
		},
		{
			name:      "error - no identity",
			patientID: testPatientID,
			limit:     10,
			offset:    0,
			setupMocks: func(*ports.MockObservationRepository) {},
			setupContext: func() context.Context {
				return context.Background()
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, result *domain.ListResponse[models.Observation], err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name:      "error - temporary token",
			patientID: testPatientID,
			limit:     10,
			offset:    0,
			setupMocks: func(*ports.MockObservationRepository) {},
			setupContext: func() context.Context {
				id := createTestIdentity("", "", []string{})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrTmpTokenForbidden,
			validateResult: func(t *testing.T, result *domain.ListResponse[models.Observation], err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.Equal(t, domain.ErrTmpTokenForbidden, err)
			},
		},
		{
			name:      "error - PatientID mismatch",
			patientID: "other-patient",
			limit:     10,
			offset:    0,
			setupMocks: func(*ports.MockObservationRepository) {},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, result *domain.ListResponse[models.Observation], err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name:      "error - repository error",
			patientID: testPatientID,
			limit:     10,
			offset:    0,
			setupMocks: func(repo *ports.MockObservationRepository) {
				repo.EXPECT().
					Search(gomock.Any(), testPatientID, 10, 0).
					Return(nil, int64(0), errors.New("database error"))
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrInternal,
			validateResult: func(t *testing.T, result *domain.ListResponse[models.Observation], err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.ErrorIs(t, err, domain.ErrInternal)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			obsRepo := ports.NewMockObservationRepository(ctrl)
			docRepo := ports.NewMockDocumentRepository(ctrl)
			validator := validator.NewObservationValidator()

			tt.setupMocks(obsRepo)

			service := NewObservationService(obsRepo, docRepo, validator)

			ctx := tt.setupContext()
			result, err := service.List(ctx, tt.patientID, tt.limit, tt.offset)

			if tt.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if tt.validateResult != nil {
				tt.validateResult(t, result, err)
			}
		})
	}
}
