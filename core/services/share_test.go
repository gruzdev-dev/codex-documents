package services

import (
	"context"
	"errors"
	"testing"

	"github.com/gruzdev-dev/codex-documents/core/domain"
	"github.com/gruzdev-dev/codex-documents/core/ports"
	"github.com/gruzdev-dev/codex-documents/pkg/identity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	models "github.com/gruzdev-dev/fhir/r5"
)

func TestShareService_Share(t *testing.T) {
	tests := []struct {
		name           string
		req            domain.ShareRequest
		setupMocks     func(*ports.MockObservationRepository, *ports.MockDocumentRepository, *ports.MockTmpAccessClient)
		setupContext   func() context.Context
		expectedError  error
		validateResult func(*testing.T, *domain.ShareResponse, error)
	}{
		{
			name: "success path - share only observations",
			req: domain.ShareRequest{
				ResourceIDs: []string{"Observation/" + testObsID},
				TTLSeconds:  3600,
			},
			setupMocks: func(obsRepo *ports.MockObservationRepository, docRepo *ports.MockDocumentRepository, client *ports.MockTmpAccessClient) {
				obs := createTestObservation(testObsID, testPatientID)
				obsRepo.EXPECT().
					GetByIDs(gomock.Any(), []string{testObsID}).
					Return([]models.Observation{*obs}, nil)
				docRepo.EXPECT().
					GetByIDs(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, ids []string) ([]models.DocumentReference, error) {
						require.Empty(t, ids)
						return []models.DocumentReference{}, nil
					})
				client.EXPECT().
					GenerateTmpToken(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, req domain.GenerateTmpTokenRequest) (*domain.GenerateTmpTokenResponse, error) {
						require.Contains(t, req.Payload["scopes"], "docs:observation:"+testObsID+":read")
						return &domain.GenerateTmpTokenResponse{
							TmpToken: "tmp-token-123",
						}, nil
					})
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, resp *domain.ShareResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, resp)
				assert.Equal(t, "tmp-token-123", resp.Token)
				assert.Equal(t, "/api/v1/shared", resp.ResourceURL)
			},
		},
		{
			name: "success path - share only documents",
			req: domain.ShareRequest{
				ResourceIDs: []string{"DocumentReference/" + testDocID},
				TTLSeconds:  3600,
			},
			setupMocks: func(obsRepo *ports.MockObservationRepository, docRepo *ports.MockDocumentRepository, client *ports.MockTmpAccessClient) {
				obsRepo.EXPECT().
					GetByIDs(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, ids []string) ([]models.Observation, error) {
						require.Empty(t, ids)
						return []models.Observation{}, nil
					})
				doc := createTestDocument(testDocID, testPatientID)
				docRepo.EXPECT().
					GetByIDs(gomock.Any(), []string{testDocID}).
					Return([]models.DocumentReference{*doc}, nil)
				client.EXPECT().
					GenerateTmpToken(gomock.Any(), gomock.Any()).
					Return(&domain.GenerateTmpTokenResponse{
						TmpToken: "tmp-token-123",
					}, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, resp *domain.ShareResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, resp)
			},
		},
		{
			name: "success path - share observations + documents",
			req: domain.ShareRequest{
				ResourceIDs: []string{"Observation/" + testObsID, "DocumentReference/" + testDocID},
				TTLSeconds:  3600,
			},
			setupMocks: func(obsRepo *ports.MockObservationRepository, docRepo *ports.MockDocumentRepository, client *ports.MockTmpAccessClient) {
				obs := createTestObservation(testObsID, testPatientID)
				obsRepo.EXPECT().
					GetByIDs(gomock.Any(), []string{testObsID}).
					Return([]models.Observation{*obs}, nil)
				doc := createTestDocument(testDocID, testPatientID)
				docRepo.EXPECT().
					GetByIDs(gomock.Any(), []string{testDocID}).
					Return([]models.DocumentReference{*doc}, nil)
				client.EXPECT().
					GenerateTmpToken(gomock.Any(), gomock.Any()).
					Return(&domain.GenerateTmpTokenResponse{
						TmpToken: "tmp-token-123",
					}, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, resp *domain.ShareResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, resp)
			},
		},
		{
			name: "success path - observations with derivedFrom documents",
			req: domain.ShareRequest{
				ResourceIDs: []string{"Observation/" + testObsID},
				TTLSeconds:  3600,
			},
			setupMocks: func(obsRepo *ports.MockObservationRepository, docRepo *ports.MockDocumentRepository, client *ports.MockTmpAccessClient) {
				obs := createTestObservationWithDerivedFrom(testObsID, testPatientID, []string{testDocID})
				obsRepo.EXPECT().
					GetByIDs(gomock.Any(), []string{testObsID}).
					Return([]models.Observation{*obs}, nil)
				docRepo.EXPECT().
					GetByIDs(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, ids []string) ([]models.DocumentReference, error) {
						require.Empty(t, ids)
						return []models.DocumentReference{}, nil
					})
				doc := createTestDocument(testDocID, testPatientID)
				docRepo.EXPECT().
					GetByIDs(gomock.Any(), []string{testDocID}).
					Return([]models.DocumentReference{*doc}, nil)
				client.EXPECT().
					GenerateTmpToken(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, req domain.GenerateTmpTokenRequest) (*domain.GenerateTmpTokenResponse, error) {
						require.Contains(t, req.Payload["scopes"], "docs:observation:"+testObsID+":read")
						require.Contains(t, req.Payload["scopes"], "docs:document_reference:"+testDocID+":read")
						return &domain.GenerateTmpTokenResponse{
							TmpToken: "tmp-token-123",
						}, nil
					})
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, resp *domain.ShareResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, resp)
			},
		},
		{
			name: "success path - documents with files",
			req: domain.ShareRequest{
				ResourceIDs: []string{"DocumentReference/" + testDocID},
				TTLSeconds:  3600,
			},
			setupMocks: func(obsRepo *ports.MockObservationRepository, docRepo *ports.MockDocumentRepository, client *ports.MockTmpAccessClient) {
				obsRepo.EXPECT().
					GetByIDs(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, ids []string) ([]models.Observation, error) {
						require.Empty(t, ids)
						return []models.Observation{}, nil
					})
				doc := createTestDocument(testDocID, testPatientID)
				doc.Content[0].Attachment.Id = strPtr(testFileID)
				docRepo.EXPECT().
					GetByIDs(gomock.Any(), []string{testDocID}).
					Return([]models.DocumentReference{*doc}, nil)
				client.EXPECT().
					GenerateTmpToken(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, req domain.GenerateTmpTokenRequest) (*domain.GenerateTmpTokenResponse, error) {
						require.Contains(t, req.Payload["scopes"], "docs:document_reference:"+testDocID+":read")
						require.Contains(t, req.Payload["scopes"], "files:file:"+testFileID+":read")
						return &domain.GenerateTmpTokenResponse{
							TmpToken: "tmp-token-123",
						}, nil
					})
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, resp *domain.ShareResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, resp)
			},
		},
		{
			name: "success path - resourceIDs with prefixes",
			req: domain.ShareRequest{
				ResourceIDs: []string{"Observation/" + testObsID, "DocumentReference/" + testDocID},
				TTLSeconds:  3600,
			},
			setupMocks: func(obsRepo *ports.MockObservationRepository, docRepo *ports.MockDocumentRepository, client *ports.MockTmpAccessClient) {
				obs := createTestObservation(testObsID, testPatientID)
				obsRepo.EXPECT().
					GetByIDs(gomock.Any(), []string{testObsID}).
					Return([]models.Observation{*obs}, nil)
				doc := createTestDocument(testDocID, testPatientID)
				docRepo.EXPECT().
					GetByIDs(gomock.Any(), []string{testDocID}).
					Return([]models.DocumentReference{*doc}, nil)
				client.EXPECT().
					GenerateTmpToken(gomock.Any(), gomock.Any()).
					Return(&domain.GenerateTmpTokenResponse{
						TmpToken: "tmp-token-123",
					}, nil)
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, resp *domain.ShareResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, resp)
			},
		},
		{
			name: "error - no identity",
			req: domain.ShareRequest{
				ResourceIDs: []string{testObsID},
				TTLSeconds:  3600,
			},
			setupMocks: func(*ports.MockObservationRepository, *ports.MockDocumentRepository, *ports.MockTmpAccessClient) {},
			setupContext: func() context.Context {
				return context.Background()
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, resp *domain.ShareResponse, err error) {
				assert.Error(t, err)
				assert.Nil(t, resp)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name: "error - no read scope",
			req: domain.ShareRequest{
				ResourceIDs: []string{testObsID},
				TTLSeconds:  3600,
			},
			setupMocks: func(*ports.MockObservationRepository, *ports.MockDocumentRepository, *ports.MockTmpAccessClient) {},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, resp *domain.ShareResponse, err error) {
				assert.Error(t, err)
				assert.Nil(t, resp)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name: "error - empty ResourceIDs",
			req: domain.ShareRequest{
				ResourceIDs: []string{},
				TTLSeconds:  3600,
			},
			setupMocks: func(*ports.MockObservationRepository, *ports.MockDocumentRepository, *ports.MockTmpAccessClient) {},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrNoResourcesToShare,
			validateResult: func(t *testing.T, resp *domain.ShareResponse, err error) {
				assert.Error(t, err)
				assert.Nil(t, resp)
				assert.Equal(t, domain.ErrNoResourcesToShare, err)
			},
		},
		{
			name: "error - resource not found",
			req: domain.ShareRequest{
				ResourceIDs: []string{"Observation/" + testObsID},
				TTLSeconds:  3600,
			},
			setupMocks: func(obsRepo *ports.MockObservationRepository, docRepo *ports.MockDocumentRepository, client *ports.MockTmpAccessClient) {
				obsRepo.EXPECT().
					GetByIDs(gomock.Any(), []string{testObsID}).
					Return([]models.Observation{}, nil)
				docRepo.EXPECT().
					GetByIDs(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, ids []string) ([]models.DocumentReference, error) {
						require.Empty(t, ids)
						return []models.DocumentReference{}, nil
					})
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrResourceNotOwned,
			validateResult: func(t *testing.T, resp *domain.ShareResponse, err error) {
				assert.Error(t, err)
				assert.Nil(t, resp)
				assert.Equal(t, domain.ErrResourceNotOwned, err)
			},
		},
		{
			name: "error - resource not owned",
			req: domain.ShareRequest{
				ResourceIDs: []string{"Observation/" + testObsID},
				TTLSeconds:  3600,
			},
			setupMocks: func(obsRepo *ports.MockObservationRepository, docRepo *ports.MockDocumentRepository, client *ports.MockTmpAccessClient) {
				obs := createTestObservation(testObsID, "other-patient")
				obsRepo.EXPECT().
					GetByIDs(gomock.Any(), []string{testObsID}).
					Return([]models.Observation{*obs}, nil)
				docRepo.EXPECT().
					GetByIDs(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, ids []string) ([]models.DocumentReference, error) {
						require.Empty(t, ids)
						return []models.DocumentReference{}, nil
					})
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrResourceNotOwned,
			validateResult: func(t *testing.T, resp *domain.ShareResponse, err error) {
				assert.Error(t, err)
				assert.Nil(t, resp)
				assert.Equal(t, domain.ErrResourceNotOwned, err)
			},
		},
		{
			name: "error - observation repository error",
			req: domain.ShareRequest{
				ResourceIDs: []string{"Observation/" + testObsID},
				TTLSeconds:  3600,
			},
			setupMocks: func(obsRepo *ports.MockObservationRepository, docRepo *ports.MockDocumentRepository, client *ports.MockTmpAccessClient) {
				obsRepo.EXPECT().
					GetByIDs(gomock.Any(), []string{testObsID}).
					Return(nil, errors.New("database error"))
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrInternal,
			validateResult: func(t *testing.T, resp *domain.ShareResponse, err error) {
				assert.Error(t, err)
				assert.Nil(t, resp)
				assert.ErrorIs(t, err, domain.ErrInternal)
			},
		},
		{
			name: "error - document repository error",
			req: domain.ShareRequest{
				ResourceIDs: []string{"DocumentReference/" + testDocID},
				TTLSeconds:  3600,
			},
			setupMocks: func(obsRepo *ports.MockObservationRepository, docRepo *ports.MockDocumentRepository, client *ports.MockTmpAccessClient) {
				obsRepo.EXPECT().
					GetByIDs(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, ids []string) ([]models.Observation, error) {
						require.Empty(t, ids)
						return []models.Observation{}, nil
					})
				docRepo.EXPECT().
					GetByIDs(gomock.Any(), []string{testDocID}).
					Return(nil, errors.New("database error"))
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrInternal,
			validateResult: func(t *testing.T, resp *domain.ShareResponse, err error) {
				assert.Error(t, err)
				assert.Nil(t, resp)
				assert.ErrorIs(t, err, domain.ErrInternal)
			},
		},
		{
			name: "error - tmp access client error",
			req: domain.ShareRequest{
				ResourceIDs: []string{"Observation/" + testObsID},
				TTLSeconds:  3600,
			},
			setupMocks: func(obsRepo *ports.MockObservationRepository, docRepo *ports.MockDocumentRepository, client *ports.MockTmpAccessClient) {
				obs := createTestObservation(testObsID, testPatientID)
				obsRepo.EXPECT().
					GetByIDs(gomock.Any(), []string{testObsID}).
					Return([]models.Observation{*obs}, nil)
				docRepo.EXPECT().
					GetByIDs(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, ids []string) ([]models.DocumentReference, error) {
						require.Empty(t, ids)
						return []models.DocumentReference{}, nil
					})
				client.EXPECT().
					GenerateTmpToken(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("token service error"))
			},
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrInternal,
			validateResult: func(t *testing.T, resp *domain.ShareResponse, err error) {
				assert.Error(t, err)
				assert.Nil(t, resp)
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
			client := ports.NewMockTmpAccessClient(ctrl)

			tt.setupMocks(obsRepo, docRepo, client)

			service := NewShareService(obsRepo, docRepo, client)

			ctx := tt.setupContext()
			result, err := service.Share(ctx, tt.req)

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

func TestShareService_GetSharedResources(t *testing.T) {
	tests := []struct {
		name           string
		setupContext   func() context.Context
		expectedError  error
		validateResult func(*testing.T, *domain.SharedResourcesResponse, error)
	}{
		{
			name: "success path - temporary token with scopes",
			setupContext: func() context.Context {
				id := domain.Identity{
					UserID:    "",
					PatientID: "",
					Scopes: []string{
						"docs:observation:" + testObsID + ":read",
						"docs:document_reference:" + testDocID + ":read",
						"files:file:" + testFileID + ":read",
					},
				}
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, resp *domain.SharedResourcesResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, resp)
				assert.Contains(t, resp.Observations, "/api/v1/Observation/"+testObsID)
				assert.Contains(t, resp.DocumentReferences, "/api/v1/DocumentReference/"+testDocID)
			},
		},
		{
			name: "error - no identity",
			setupContext: func() context.Context {
				return context.Background()
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, resp *domain.SharedResourcesResponse, err error) {
				assert.Error(t, err)
				assert.Nil(t, resp)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name: "error - not temporary token",
			setupContext: func() context.Context {
				id := createTestIdentity(testPatientID, testUserID, []string{"patient/*.read"})
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: domain.ErrAccessDenied,
			validateResult: func(t *testing.T, resp *domain.SharedResourcesResponse, err error) {
				assert.Error(t, err)
				assert.Nil(t, resp)
				assert.Equal(t, domain.ErrAccessDenied, err)
			},
		},
		{
			name: "success path - empty scopes",
			setupContext: func() context.Context {
				id := domain.Identity{
					UserID:    "",
					PatientID: "",
					Scopes:    []string{},
				}
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, resp *domain.SharedResourcesResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, resp)
				assert.Empty(t, resp.Observations)
				assert.Empty(t, resp.DocumentReferences)
			},
		},
		{
			name: "success path - invalid scope format",
			setupContext: func() context.Context {
				id := domain.Identity{
					UserID:    "",
					PatientID: "",
					Scopes: []string{
						"invalid:scope",
						"docs:observation:read",
						"other:service:resource:id:action",
					},
				}
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, resp *domain.SharedResourcesResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, resp)
				assert.Empty(t, resp.Observations)
				assert.Empty(t, resp.DocumentReferences)
			},
		},
		{
			name: "success path - non-read scopes filtered",
			setupContext: func() context.Context {
				id := domain.Identity{
					UserID:    "",
					PatientID: "",
					Scopes: []string{
						"docs:observation:" + testObsID + ":write",
						"docs:document_reference:" + testDocID + ":read",
					},
				}
				return identity.WithCtx(context.Background(), id)
			},
			expectedError: nil,
			validateResult: func(t *testing.T, resp *domain.SharedResourcesResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, resp)
				assert.Empty(t, resp.Observations)
				assert.Contains(t, resp.DocumentReferences, "/api/v1/DocumentReference/"+testDocID)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			obsRepo := ports.NewMockObservationRepository(ctrl)
			docRepo := ports.NewMockDocumentRepository(ctrl)
			client := ports.NewMockTmpAccessClient(ctrl)

			service := NewShareService(obsRepo, docRepo, client)

			ctx := tt.setupContext()
			result, err := service.GetSharedResources(ctx)

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
