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
	testPatientID = "patient-123"
	testDocID     = "doc-123"
	testFileID    = "file-123"
	testUserID    = "user-123"
	testContentType = "application/pdf"
	testFileSize    = int64(1024)
)

func strPtr(s string) *string {
	return &s
}

func int64Ptr(i int64) *int64 {
	return &i
}

func createTestIdentity(patientID, userID string, scopes []string) domain.Identity {
	return domain.Identity{
		PatientID: patientID,
		UserID:    userID,
		Scopes:    scopes,
	}
}

func createTestDocument(id, patientID string) *models.DocumentReference {
	patientRef := "Patient/" + patientID
	return &models.DocumentReference{
		ResourceType: "DocumentReference",
		Id:           strPtr(id),
		Status:       "current",
		Subject: &models.Reference{
			Reference: &patientRef,
		},
		Content: []models.DocumentReferenceContent{
			{
				Attachment: &models.Attachment{
					ContentType: strPtr(testContentType),
					Size:        int64Ptr(testFileSize),
					Title:       strPtr("Test Document"),
				},
			},
		},
	}
}

func createTestDocumentWithoutFiles(id, patientID string) *models.DocumentReference {
	patientRef := "Patient/" + patientID
	return &models.DocumentReference{
		ResourceType: "DocumentReference",
		Id:           strPtr(id),
		Status:       "current",
		Subject: &models.Reference{
			Reference: &patientRef,
		},
		Content: []models.DocumentReferenceContent{
			{
				Attachment: &models.Attachment{
					ContentType: strPtr(testContentType),
					Size:        int64Ptr(testFileSize),
					Url:         strPtr("https://example.com/file"),
				},
			},
		},
	}
}

func TestDocumentService_CreateDocument(t *testing.T) {
	tests := []struct {
		name           string
		doc            *models.DocumentReference
		setupMocks     func(*ports.MockDocumentRepository, *ports.MockFileProvider)
		setupContext   func() context.Context
		expectedError  error
		validateResult func(*testing.T, *domain.CreateDocumentResult, error)
	}{
		{
			name: "success path - with files",
			doc: &models.DocumentReference{
				ResourceType: "DocumentReference",
				Status:       "current",
				Content: []models.DocumentReferenceContent{
					{
						Attachment: &models.Attachment{
							ContentType: strPtr(testContentType),
							Size:        int64Ptr(testFileSize),
						},
					},
				},
			},
			setupMocks: func(repo *ports.MockDocumentRepository, provider *ports.MockFileProvider) {
				provider.EXPECT().
					GetPresignedUrls(gomock.Any(), domain.GetPresignedUrlsRequest{
						UserId:      testPatientID,
						ContentType: testContentType,
						Size:        testFileSize,
					}).
					Return(&domain.PresignedUrlsResponse{
						FileId:      testFileID,
						UploadUrl:   "https://s3.example.com/upload",
						DownloadUrl: "https://s3.example.com/download",
					}, nil)
				repo.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, doc *models.DocumentReference) (*models.DocumentReference, error) {
						require.NotNil(t, doc.Id)
						require.NotEmpty(t, *doc.Id)
						require.NotNil(t, doc.Subject)
						require.Equal(t, "Patient/"+testPatientID, *doc.Subject.Reference)
						require.NotNil(t, doc.Content[0].Attachment.Id)
						require.NotNil(t, doc.Content[0].Attachment.Url)
						return doc, nil
					})
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, result *domain.CreateDocumentResult, err error) {
				require.NoError(t, err)
				require.NotNil(t, result)
				require.NotNil(t, result.Document)
				require.NotEmpty(t, result.UploadUrls)
				assert.Equal(t, "https://s3.example.com/upload", result.UploadUrls[testFileID])
			},
		},
		{
			name: "success path - without files",
			doc: &models.DocumentReference{
				ResourceType: "DocumentReference",
				Status:       "current",
				Content: []models.DocumentReferenceContent{
					{
						Attachment: &models.Attachment{
							ContentType: strPtr(testContentType),
							Size:        int64Ptr(testFileSize),
							Url:         strPtr("https://example.com/file"),
						},
					},
				},
			},
			setupMocks: func(repo *ports.MockDocumentRepository, provider *ports.MockFileProvider) {
				repo.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, doc *models.DocumentReference) (*models.DocumentReference, error) {
						return doc, nil
					})
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, result *domain.CreateDocumentResult, err error) {
				require.NoError(t, err)
				require.NotNil(t, result)
				require.Empty(t, result.UploadUrls)
			},
		},
		{
			name: "validation error - nil document",
			doc:  nil,
			setupMocks: func(*ports.MockDocumentRepository, *ports.MockFileProvider) {},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrInvalidInput,
			validateResult: func(t *testing.T, result *domain.CreateDocumentResult, err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.ErrorIs(t, err, domain.ErrInvalidInput)
			},
		},
		{
			name: "validation error - empty content",
			doc: &models.DocumentReference{
				ResourceType: "DocumentReference",
				Status:       "current",
				Content:      []models.DocumentReferenceContent{},
			},
			setupMocks: func(*ports.MockDocumentRepository, *ports.MockFileProvider) {},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrInvalidInput,
			validateResult: func(t *testing.T, result *domain.CreateDocumentResult, err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.ErrorIs(t, err, domain.ErrInvalidInput)
			},
		},
		{
			name: "error - ID already provided",
			doc: &models.DocumentReference{
				ResourceType: "DocumentReference",
				Id:           strPtr(testDocID),
				Status:       "current",
				Content: []models.DocumentReferenceContent{
					{
						Attachment: &models.Attachment{
							ContentType: strPtr(testContentType),
							Size:        int64Ptr(testFileSize),
						},
					},
				},
			},
			setupMocks: func(*ports.MockDocumentRepository, *ports.MockFileProvider) {},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrInvalidInput,
			validateResult: func(t *testing.T, result *domain.CreateDocumentResult, err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.ErrorIs(t, err, domain.ErrInvalidInput)
				assert.Contains(t, err.Error(), "document ID must not be provided")
			},
		},
		{
			name: "error - no identity in context",
			doc: &models.DocumentReference{
				ResourceType: "DocumentReference",
				Status:       "current",
				Content: []models.DocumentReferenceContent{
					{
						Attachment: &models.Attachment{
							ContentType: strPtr(testContentType),
							Size:        int64Ptr(testFileSize),
						},
					},
				},
			},
			setupMocks: func(*ports.MockDocumentRepository, *ports.MockFileProvider) {},
			setupContext: func() context.Context {
				return context.Background()
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, result *domain.CreateDocumentResult, err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name: "error - temporary token",
			doc: &models.DocumentReference{
				ResourceType: "DocumentReference",
				Status:       "current",
				Content: []models.DocumentReferenceContent{
					{
						Attachment: &models.Attachment{
							ContentType: strPtr(testContentType),
							Size:        int64Ptr(testFileSize),
						},
					},
				},
			},
			setupMocks: func(*ports.MockDocumentRepository, *ports.MockFileProvider) {},
			setupContext: func() context.Context {
				id := createTestIdentity("", "", []string{})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrTmpTokenForbidden,
			validateResult: func(t *testing.T, result *domain.CreateDocumentResult, err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.Equal(t, domain.ErrTmpTokenForbidden, err)
			},
		},
		{
			name: "error - no write scope",
			doc: &models.DocumentReference{
				ResourceType: "DocumentReference",
				Status:       "current",
				Content: []models.DocumentReferenceContent{
					{
						Attachment: &models.Attachment{
							ContentType: strPtr(testContentType),
							Size:        int64Ptr(testFileSize),
						},
					},
				},
			},
			setupMocks: func(*ports.MockDocumentRepository, *ports.MockFileProvider) {},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, result *domain.CreateDocumentResult, err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name: "error - empty PatientID",
			doc: &models.DocumentReference{
				ResourceType: "DocumentReference",
				Status:       "current",
				Content: []models.DocumentReferenceContent{
					{
						Attachment: &models.Attachment{
							ContentType: strPtr(testContentType),
							Size:        int64Ptr(testFileSize),
						},
					},
				},
			},
			setupMocks: func(*ports.MockDocumentRepository, *ports.MockFileProvider) {},
			setupContext: func() context.Context {
				id := createTestIdentity("", testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, result *domain.CreateDocumentResult, err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name: "error - file provider error",
			doc: &models.DocumentReference{
				ResourceType: "DocumentReference",
				Status:       "current",
				Content: []models.DocumentReferenceContent{
					{
						Attachment: &models.Attachment{
							ContentType: strPtr(testContentType),
							Size:        int64Ptr(testFileSize),
						},
					},
				},
			},
			setupMocks: func(repo *ports.MockDocumentRepository, provider *ports.MockFileProvider) {
				provider.EXPECT().
					GetPresignedUrls(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("s3 error"))
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrInternal,
			validateResult: func(t *testing.T, result *domain.CreateDocumentResult, err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.ErrorIs(t, err, domain.ErrInternal)
			},
		},
		{
			name: "error - repository create error",
			doc: &models.DocumentReference{
				ResourceType: "DocumentReference",
				Status:       "current",
				Content: []models.DocumentReferenceContent{
					{
						Attachment: &models.Attachment{
							ContentType: strPtr(testContentType),
							Size:        int64Ptr(testFileSize),
							Url:         strPtr("https://example.com/file"),
						},
					},
				},
			},
			setupMocks: func(repo *ports.MockDocumentRepository, provider *ports.MockFileProvider) {
				repo.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("database error"))
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrInternal,
			validateResult: func(t *testing.T, result *domain.CreateDocumentResult, err error) {
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

			repo := ports.NewMockDocumentRepository(ctrl)
			provider := ports.NewMockFileProvider(ctrl)
			validator := validator.NewDocumentValidator()

			tt.setupMocks(repo, provider)

			service := NewDocumentService(repo, provider, validator)

			ctx := tt.setupContext()
			result, err := service.CreateDocument(ctx, tt.doc)

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

func TestDocumentService_GetDocument(t *testing.T) {
	tests := []struct {
		name           string
		docID          string
		setupMocks     func(*ports.MockDocumentRepository)
		setupContext   func() context.Context
		expectedError  error
		validateResult func(*testing.T, *models.DocumentReference, error)
	}{
		{
			name:  "success path - owner with read scope",
			docID: testDocID,
			setupMocks: func(repo *ports.MockDocumentRepository) {
				doc := createTestDocument(testDocID, testPatientID)
				repo.EXPECT().
					GetByID(gomock.Any(), testDocID).
					Return(doc, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, doc *models.DocumentReference, err error) {
				require.NoError(t, err)
				require.NotNil(t, doc)
				assert.Equal(t, testDocID, *doc.Id)
			},
		},
		{
			name:  "success path - access via resource scope",
			docID: testDocID,
			setupMocks: func(repo *ports.MockDocumentRepository) {
				doc := createTestDocument(testDocID, "other-patient")
				repo.EXPECT().
					GetByID(gomock.Any(), testDocID).
					Return(doc, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"docs:document_reference:" + testDocID + ":read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, doc *models.DocumentReference, err error) {
				require.NoError(t, err)
				require.NotNil(t, doc)
			},
		},
		{
			name: "error - empty ID",
			setupMocks: func(*ports.MockDocumentRepository) {},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrDocumentIDRequired,
			validateResult: func(t *testing.T, doc *models.DocumentReference, err error) {
				assert.Error(t, err)
				assert.Nil(t, doc)
				assert.Equal(t, domain.ErrDocumentIDRequired, err)
			},
		},
		{
			name:  "error - no identity",
			docID: testDocID,
			setupMocks: func(*ports.MockDocumentRepository) {},
			setupContext: func() context.Context {
				return context.Background()
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, doc *models.DocumentReference, err error) {
				assert.Error(t, err)
				assert.Nil(t, doc)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name:  "error - document not found",
			docID: testDocID,
			setupMocks: func(repo *ports.MockDocumentRepository) {
				repo.EXPECT().
					GetByID(gomock.Any(), testDocID).
					Return(nil, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrDocumentNotFound,
			validateResult: func(t *testing.T, doc *models.DocumentReference, err error) {
				assert.Error(t, err)
				assert.Nil(t, doc)
				assert.Equal(t, domain.ErrDocumentNotFound, err)
			},
		},
		{
			name:  "error - access denied",
			docID: testDocID,
			setupMocks: func(repo *ports.MockDocumentRepository) {
				doc := createTestDocument(testDocID, "other-patient")
				repo.EXPECT().
					GetByID(gomock.Any(), testDocID).
					Return(doc, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, doc *models.DocumentReference, err error) {
				assert.Error(t, err)
				assert.Nil(t, doc)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name:  "error - repository error",
			docID: testDocID,
			setupMocks: func(repo *ports.MockDocumentRepository) {
				repo.EXPECT().
					GetByID(gomock.Any(), testDocID).
					Return(nil, errors.New("database error"))
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrInternal,
			validateResult: func(t *testing.T, doc *models.DocumentReference, err error) {
				assert.Error(t, err)
				assert.Nil(t, doc)
				assert.ErrorIs(t, err, domain.ErrInternal)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := ports.NewMockDocumentRepository(ctrl)
			provider := ports.NewMockFileProvider(ctrl)
			validator := validator.NewDocumentValidator()

			tt.setupMocks(repo)

			service := NewDocumentService(repo, provider, validator)

			ctx := tt.setupContext()
			result, err := service.GetDocument(ctx, tt.docID)

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

func TestDocumentService_DeleteDocument(t *testing.T) {
	tests := []struct {
		name           string
		docID          string
		setupMocks     func(*ports.MockDocumentRepository, *ports.MockFileProvider)
		setupContext   func() context.Context
		expectedError  error
		validateResult func(*testing.T, error)
	}{
		{
			name: "success path - with files",
			docID: testDocID,
			setupMocks: func(repo *ports.MockDocumentRepository, provider *ports.MockFileProvider) {
				doc := createTestDocument(testDocID, testPatientID)
				doc.Content[0].Attachment.Id = strPtr(testFileID)
				repo.EXPECT().
					GetByID(gomock.Any(), testDocID).
					Return(doc, nil)
				provider.EXPECT().
					DeleteFile(gomock.Any(), testFileID).
					Return(nil)
				repo.EXPECT().
					Delete(gomock.Any(), testDocID).
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
			name: "success path - without files",
			docID: testDocID,
			setupMocks: func(repo *ports.MockDocumentRepository, provider *ports.MockFileProvider) {
				doc := createTestDocumentWithoutFiles(testDocID, testPatientID)
				repo.EXPECT().
					GetByID(gomock.Any(), testDocID).
					Return(doc, nil)
				repo.EXPECT().
					Delete(gomock.Any(), testDocID).
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
			docID: testDocID,
			setupMocks: func(*ports.MockDocumentRepository, *ports.MockFileProvider) {},
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
			docID: testDocID,
			setupMocks: func(*ports.MockDocumentRepository, *ports.MockFileProvider) {},
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
			docID: testDocID,
			setupMocks: func(*ports.MockDocumentRepository, *ports.MockFileProvider) {},
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
			name: "error - document not found",
			docID: testDocID,
			setupMocks: func(repo *ports.MockDocumentRepository, provider *ports.MockFileProvider) {
				repo.EXPECT().
					GetByID(gomock.Any(), testDocID).
					Return(nil, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.write"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrDocumentNotFound,
			validateResult: func(t *testing.T, err error) {
				assert.Error(t, err)
				assert.Equal(t, domain.ErrDocumentNotFound, err)
			},
		},
		{
			name: "error - not owner",
			docID: testDocID,
			setupMocks: func(repo *ports.MockDocumentRepository, provider *ports.MockFileProvider) {
				doc := createTestDocument(testDocID, "other-patient")
				repo.EXPECT().
					GetByID(gomock.Any(), testDocID).
					Return(doc, nil)
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
			name: "error - repository delete error",
			docID: testDocID,
			setupMocks: func(repo *ports.MockDocumentRepository, provider *ports.MockFileProvider) {
				doc := createTestDocumentWithoutFiles(testDocID, testPatientID)
				repo.EXPECT().
					GetByID(gomock.Any(), testDocID).
					Return(doc, nil)
				repo.EXPECT().
					Delete(gomock.Any(), testDocID).
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

			repo := ports.NewMockDocumentRepository(ctrl)
			provider := ports.NewMockFileProvider(ctrl)
			validator := validator.NewDocumentValidator()

			tt.setupMocks(repo, provider)

			service := NewDocumentService(repo, provider, validator)

			ctx := tt.setupContext()
			err := service.DeleteDocument(ctx, tt.docID)

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

func TestDocumentService_ListDocuments(t *testing.T) {
	tests := []struct {
		name           string
		patientID      string
		limit          int
		offset         int
		setupMocks     func(*ports.MockDocumentRepository)
		setupContext   func() context.Context
		expectedError  error
		validateResult func(*testing.T, *domain.ListResponse[models.DocumentReference], error)
	}{
		{
			name:      "success path",
			patientID: testPatientID,
			limit:     10,
			offset:    0,
			setupMocks: func(repo *ports.MockDocumentRepository) {
				docs := []models.DocumentReference{
					*createTestDocument("doc-1", testPatientID),
					*createTestDocument("doc-2", testPatientID),
				}
				repo.EXPECT().
					Search(gomock.Any(), testPatientID, 10, 0).
					Return(docs, int64(2), nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, result *domain.ListResponse[models.DocumentReference], err error) {
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
			setupMocks: func(*ports.MockDocumentRepository) {},
			setupContext: func() context.Context {
				return context.Background()
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, result *domain.ListResponse[models.DocumentReference], err error) {
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
			setupMocks: func(*ports.MockDocumentRepository) {},
			setupContext: func() context.Context {
				id := createTestIdentity("", "", []string{})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrTmpTokenForbidden,
			validateResult: func(t *testing.T, result *domain.ListResponse[models.DocumentReference], err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.Equal(t, domain.ErrTmpTokenForbidden, err)
			},
		},
		{
			name:      "error - no read scope",
			patientID: testPatientID,
			limit:     10,
			offset:    0,
			setupMocks: func(*ports.MockDocumentRepository) {},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, result *domain.ListResponse[models.DocumentReference], err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name:      "error - PatientID mismatch",
			patientID: "other-patient",
			limit:     10,
			offset:    0,
			setupMocks: func(*ports.MockDocumentRepository) {},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, result *domain.ListResponse[models.DocumentReference], err error) {
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
			setupMocks: func(repo *ports.MockDocumentRepository) {
				repo.EXPECT().
					Search(gomock.Any(), testPatientID, 10, 0).
					Return(nil, int64(0), errors.New("database error"))
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrInternal,
			validateResult: func(t *testing.T, result *domain.ListResponse[models.DocumentReference], err error) {
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

			repo := ports.NewMockDocumentRepository(ctrl)
			provider := ports.NewMockFileProvider(ctrl)
			validator := validator.NewDocumentValidator()

			tt.setupMocks(repo)

			service := NewDocumentService(repo, provider, validator)

			ctx := tt.setupContext()
			result, err := service.ListDocuments(ctx, tt.patientID, tt.limit, tt.offset)

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
