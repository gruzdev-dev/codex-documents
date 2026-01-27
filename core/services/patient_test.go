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

func createTestPatient(id string) *models.Patient {
	email := "test@example.com"
	patient := &models.Patient{
		ResourceType: "Patient",
		Telecom: []models.ContactPoint{
			{
				System: strPtr("email"),
				Value:  &email,
			},
		},
	}
	if id != "" {
		patient.Id = strPtr(id)
	}
	return patient
}

func TestPatientService_Create(t *testing.T) {
	tests := []struct {
		name           string
		patient        *models.Patient
		setupMocks     func(*ports.MockPatientRepository)
		expectedError  error
		validateResult func(*testing.T, *models.Patient, error)
	}{
		{
			name:    "success path - with email",
			patient: createTestPatient(""),
			setupMocks: func(repo *ports.MockPatientRepository) {
				repo.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, patient *models.Patient) (*models.Patient, error) {
						require.NotNil(t, patient.Id)
						require.NotEmpty(t, *patient.Id)
						return patient, nil
					})
			},
			expectedError: nil,
			validateResult: func(t *testing.T, patient *models.Patient, err error) {
				require.NoError(t, err)
				require.NotNil(t, patient)
				require.NotNil(t, patient.Id)
				require.NotEmpty(t, *patient.Id)
			},
		},
		{
			name:    "validation error - nil patient",
			patient: nil,
			setupMocks: func(*ports.MockPatientRepository) {},
			expectedError: domain.ErrInvalidInput,
			validateResult: func(t *testing.T, patient *models.Patient, err error) {
				assert.Error(t, err)
				assert.Nil(t, patient)
				assert.ErrorIs(t, err, domain.ErrInvalidInput)
			},
		},
		{
			name: "validation error - no email",
			patient: &models.Patient{
				ResourceType: "Patient",
			},
			setupMocks: func(*ports.MockPatientRepository) {},
			expectedError: domain.ErrInvalidInput,
			validateResult: func(t *testing.T, patient *models.Patient, err error) {
				assert.Error(t, err)
				assert.Nil(t, patient)
				assert.ErrorIs(t, err, domain.ErrInvalidInput)
				assert.Contains(t, err.Error(), "email")
			},
		},
		{
			name: "error - ID already provided",
			patient: &models.Patient{
				ResourceType: "Patient",
				Id:           strPtr(testPatientID),
				Telecom: []models.ContactPoint{
					{
						System: strPtr("email"),
						Value:  strPtr("test@example.com"),
					},
				},
			},
			setupMocks: func(*ports.MockPatientRepository) {},
			expectedError: domain.ErrInvalidInput,
			validateResult: func(t *testing.T, patient *models.Patient, err error) {
				assert.Error(t, err)
				assert.Nil(t, patient)
				assert.ErrorIs(t, err, domain.ErrInvalidInput)
				assert.Contains(t, err.Error(), "patient ID must not be provided")
			},
		},
		{
			name:    "error - repository error",
			patient: createTestPatient(""),
			setupMocks: func(repo *ports.MockPatientRepository) {
				repo.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("database error"))
			},
			expectedError: domain.ErrInternal,
			validateResult: func(t *testing.T, patient *models.Patient, err error) {
				assert.Error(t, err)
				assert.Nil(t, patient)
				assert.ErrorIs(t, err, domain.ErrInternal)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := ports.NewMockPatientRepository(ctrl)
			validator := validator.NewPatientValidator()

			tt.setupMocks(repo)

			service := NewPatientService(repo, validator)

			result, err := service.Create(context.Background(), tt.patient)

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

func TestPatientService_Get(t *testing.T) {
	tests := []struct {
		name           string
		patientID      string
		setupMocks     func(*ports.MockPatientRepository)
		setupContext   func() context.Context
		expectedError  error
		validateResult func(*testing.T, *models.Patient, error)
	}{
		{
			name:      "success path - owner",
			patientID: testPatientID,
			setupMocks: func(repo *ports.MockPatientRepository) {
				patient := createTestPatient(testPatientID)
				repo.EXPECT().
					GetByID(gomock.Any(), testPatientID).
					Return(patient, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, patient *models.Patient, err error) {
				require.NoError(t, err)
				require.NotNil(t, patient)
				assert.Equal(t, testPatientID, *patient.Id)
			},
		},
		{
			name:      "success path - access via read scope",
			patientID: testPatientID,
			setupMocks: func(repo *ports.MockPatientRepository) {
				patient := createTestPatient(testPatientID)
				repo.EXPECT().
					GetByID(gomock.Any(), testPatientID).
					Return(patient, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity("other-patient", testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, patient *models.Patient, err error) {
				require.NoError(t, err)
				require.NotNil(t, patient)
			},
		},
		{
			name:      "error - no identity",
			patientID: testPatientID,
			setupMocks: func(*ports.MockPatientRepository) {},
			setupContext: func() context.Context {
				return context.Background()
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, patient *models.Patient, err error) {
				assert.Error(t, err)
				assert.Nil(t, patient)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name:      "error - access denied",
			patientID: testPatientID,
			setupMocks: func(*ports.MockPatientRepository) {},
			setupContext: func() context.Context {
				id := createTestIdentity("other-patient", testUserID, []string{})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, patient *models.Patient, err error) {
				assert.Error(t, err)
				assert.Nil(t, patient)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name:      "error - not found",
			patientID: testPatientID,
			setupMocks: func(repo *ports.MockPatientRepository) {
				repo.EXPECT().
					GetByID(gomock.Any(), testPatientID).
					Return(nil, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrPatientNotFound,
			validateResult: func(t *testing.T, patient *models.Patient, err error) {
				assert.Error(t, err)
				assert.Nil(t, patient)
				assert.Equal(t, domain.ErrPatientNotFound, err)
			},
		},
		{
			name:      "error - repository error",
			patientID: testPatientID,
			setupMocks: func(repo *ports.MockPatientRepository) {
				repo.EXPECT().
					GetByID(gomock.Any(), testPatientID).
					Return(nil, errors.New("database error"))
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrInternal,
			validateResult: func(t *testing.T, patient *models.Patient, err error) {
				assert.Error(t, err)
				assert.Nil(t, patient)
				assert.ErrorIs(t, err, domain.ErrInternal)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := ports.NewMockPatientRepository(ctrl)
			validator := validator.NewPatientValidator()

			tt.setupMocks(repo)

			service := NewPatientService(repo, validator)

			ctx := tt.setupContext()
			result, err := service.Get(ctx, tt.patientID)

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

func TestPatientService_Update(t *testing.T) {
	tests := []struct {
		name           string
		patient        *models.Patient
		setupMocks     func(*ports.MockPatientRepository)
		setupContext   func() context.Context
		expectedError  error
		validateResult func(*testing.T, *models.Patient, error)
	}{
		{
			name:    "success path - owner updates",
			patient: createTestPatient(testPatientID),
			setupMocks: func(repo *ports.MockPatientRepository) {
				existing := createTestPatient(testPatientID)
				repo.EXPECT().
					GetByID(gomock.Any(), testPatientID).
					Return(existing, nil)
				repo.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Return(&models.Patient{}, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, patient *models.Patient, err error) {
				require.NoError(t, err)
				require.NotNil(t, patient)
			},
		},
		{
			name:    "success path - update via write scope",
			patient: createTestPatient(testPatientID),
			setupMocks: func(repo *ports.MockPatientRepository) {
				existing := createTestPatient(testPatientID)
				repo.EXPECT().
					GetByID(gomock.Any(), testPatientID).
					Return(existing, nil)
				repo.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Return(&models.Patient{}, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity("other-patient", testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, patient *models.Patient, err error) {
				require.NoError(t, err)
				require.NotNil(t, patient)
			},
		},
		{
			name: "error - no ID",
			patient: &models.Patient{
				ResourceType: "Patient",
				Telecom: []models.ContactPoint{
					{
						System: strPtr("email"),
						Value:  strPtr("test@example.com"),
					},
				},
			},
			setupMocks: func(*ports.MockPatientRepository) {},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrPatientIDRequired,
			validateResult: func(t *testing.T, patient *models.Patient, err error) {
				assert.Error(t, err)
				assert.Nil(t, patient)
				assert.Equal(t, domain.ErrPatientIDRequired, err)
			},
		},
		{
			name:    "error - no identity",
			patient: createTestPatient(testPatientID),
			setupMocks: func(*ports.MockPatientRepository) {},
			setupContext: func() context.Context {
				return context.Background()
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, patient *models.Patient, err error) {
				assert.Error(t, err)
				assert.Nil(t, patient)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name:    "error - access denied",
			patient: createTestPatient(testPatientID),
			setupMocks: func(*ports.MockPatientRepository) {},
			setupContext: func() context.Context {
				id := createTestIdentity("other-patient", testUserID, []string{})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, patient *models.Patient, err error) {
				assert.Error(t, err)
				assert.Nil(t, patient)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name: "error - validation error",
			patient: &models.Patient{
				ResourceType: "Patient",
				Id:           strPtr(testPatientID),
			},
			setupMocks: func(repo *ports.MockPatientRepository) {
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrInvalidInput,
			validateResult: func(t *testing.T, patient *models.Patient, err error) {
				assert.Error(t, err)
				assert.Nil(t, patient)
				assert.ErrorIs(t, err, domain.ErrInvalidInput)
			},
		},
		{
			name:    "error - not found",
			patient: createTestPatient(testPatientID),
			setupMocks: func(repo *ports.MockPatientRepository) {
				repo.EXPECT().
					GetByID(gomock.Any(), testPatientID).
					Return(nil, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrPatientNotFound,
			validateResult: func(t *testing.T, patient *models.Patient, err error) {
				assert.Error(t, err)
				assert.Nil(t, patient)
				assert.Equal(t, domain.ErrPatientNotFound, err)
			},
		},
		{
			name:    "error - repository error",
			patient: createTestPatient(testPatientID),
			setupMocks: func(repo *ports.MockPatientRepository) {
				existing := createTestPatient(testPatientID)
				repo.EXPECT().
					GetByID(gomock.Any(), testPatientID).
					Return(existing, nil)
				repo.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("database error"))
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrInternal,
			validateResult: func(t *testing.T, patient *models.Patient, err error) {
				assert.Error(t, err)
				assert.Nil(t, patient)
				assert.ErrorIs(t, err, domain.ErrInternal)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := ports.NewMockPatientRepository(ctrl)
			validator := validator.NewPatientValidator()

			tt.setupMocks(repo)

			service := NewPatientService(repo, validator)

			ctx := tt.setupContext()
			result, err := service.Update(ctx, tt.patient)

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
