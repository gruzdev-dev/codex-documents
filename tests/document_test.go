//go:build integration

package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	nethttp "net/http"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gruzdev-dev/codex-documents/core/domain"
	"github.com/gruzdev-dev/codex-documents/proto"
	models "github.com/gruzdev-dev/fhir/r5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/metadata"
)

func TestDocumentIntegration(t *testing.T) {
	env := SetupTestEnv(t)
	defer env.Cleanup()

	client := &nethttp.Client{}

	var patientID string
	var token string
	var doc1ID, doc2ID, doc3ID string

	t.Run("Step 1: Create Patient via gRPC", func(t *testing.T) {
		md := metadata.Pairs("x-internal-token", "test-secret")
		ctx := metadata.NewOutgoingContext(context.Background(), md)

		resp, err := env.GRPCClient.CreatePatient(ctx, &proto.CreatePatientRequest{
			Email: "test@example.com",
		})
		require.NoError(t, err)

		patientID = resp.PatientId
		require.NotEmpty(t, patientID)

		claims := jwt.MapClaims{
			"sub":        "test-user",
			"patient_id": patientID,
			"scope":      "patient/*.read patient/*.write",
		}
		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token, err = jwtToken.SignedString([]byte("secret-key"))
		require.NoError(t, err)
	})

	t.Run("Step 2: Create Document 1 - mixed attachments", func(t *testing.T) {
		env.MockFileProvider.EXPECT().GetPresignedUrls(gomock.Any(), gomock.Any()).
			Return(&domain.PresignedUrlsResponse{
				FileId:      "file-id-doc1",
				UploadUrl:   "http://test/upload/doc1",
				DownloadUrl: "http://test/download/doc1",
			}, nil).Times(1)

		externalUrl := "http://external.com/file.pdf"
		doc := models.DocumentReference{
			ResourceType: "DocumentReference",
			Status:       "current",
			Content: []models.DocumentReferenceContent{
				{
					Attachment: &models.Attachment{
						ContentType: strPtr("application/pdf"),
						Size:        int64Ptr(1024),
						Title:       strPtr("External document"),
						Url:         &externalUrl,
					},
				},
				{
					Attachment: &models.Attachment{
						ContentType: strPtr("image/jpeg"),
						Size:        int64Ptr(2048),
						Title:       strPtr("Document to upload"),
					},
				},
			},
		}

		docJSON, err := json.Marshal(doc)
		require.NoError(t, err)

		req, err := nethttp.NewRequest("POST", env.ServerURL+"/api/v1/DocumentReference", bytes.NewBuffer(docJSON))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/fhir+json")
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusCreated, resp.StatusCode)

		uploadUrlsHeader := resp.Header.Get("X-Upload-Urls")
		require.NotEmpty(t, uploadUrlsHeader, "X-Upload-Urls header should be present")

		var uploadUrls map[string]string
		err = json.Unmarshal([]byte(uploadUrlsHeader), &uploadUrls)
		require.NoError(t, err)
		assert.Contains(t, uploadUrls, "file-id-doc1")
		assert.Equal(t, "http://test/upload/doc1", uploadUrls["file-id-doc1"])

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var createdDoc models.DocumentReference
		err = json.Unmarshal(body, &createdDoc)
		require.NoError(t, err)
		require.NotNil(t, createdDoc.Id)
		doc1ID = *createdDoc.Id

		assert.Len(t, createdDoc.Content, 2)
		assert.Equal(t, externalUrl, *createdDoc.Content[0].Attachment.Url)
		assert.Equal(t, "file-id-doc1", *createdDoc.Content[1].Attachment.Id)
		assert.Equal(t, "http://test/download/doc1", *createdDoc.Content[1].Attachment.Url)
	})

	t.Run("Step 3: Create Document 2 - single attachment for upload", func(t *testing.T) {
		env.MockFileProvider.EXPECT().GetPresignedUrls(gomock.Any(), gomock.Any()).
			Return(&domain.PresignedUrlsResponse{
				FileId:      "file-id-doc2",
				UploadUrl:   "http://test/upload/doc2",
				DownloadUrl: "http://test/download/doc2",
			}, nil).Times(1)

		doc := models.DocumentReference{
			ResourceType: "DocumentReference",
			Status:       "current",
			Content: []models.DocumentReferenceContent{
				{
					Attachment: &models.Attachment{
						ContentType: strPtr("text/plain"),
						Size:        int64Ptr(512),
						Title:       strPtr("Text file"),
					},
				},
			},
		}

		docJSON, err := json.Marshal(doc)
		require.NoError(t, err)

		req, err := nethttp.NewRequest("POST", env.ServerURL+"/api/v1/DocumentReference", bytes.NewBuffer(docJSON))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/fhir+json")
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusCreated, resp.StatusCode)

		uploadUrlsHeader := resp.Header.Get("X-Upload-Urls")
		require.NotEmpty(t, uploadUrlsHeader)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var createdDoc models.DocumentReference
		err = json.Unmarshal(body, &createdDoc)
		require.NoError(t, err)
		require.NotNil(t, createdDoc.Id)
		doc2ID = *createdDoc.Id

		assert.Len(t, createdDoc.Content, 1)
		assert.Equal(t, "file-id-doc2", *createdDoc.Content[0].Attachment.Id)
	})

	t.Run("Step 4: Create Document 3 - empty content", func(t *testing.T) {
		doc := models.DocumentReference{
			ResourceType: "DocumentReference",
			Status:       "current",
			Content: []models.DocumentReferenceContent{
				{
					Attachment: &models.Attachment{
						Title: strPtr("No file attached"),
					},
				},
			},
		}

		docJSON, err := json.Marshal(doc)
		require.NoError(t, err)

		req, err := nethttp.NewRequest("POST", env.ServerURL+"/api/v1/DocumentReference", bytes.NewBuffer(docJSON))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/fhir+json")
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusCreated, resp.StatusCode)

		uploadUrlsHeader := resp.Header.Get("X-Upload-Urls")
		assert.Empty(t, uploadUrlsHeader, "No upload URLs should be present for document without files")

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var createdDoc models.DocumentReference
		err = json.Unmarshal(body, &createdDoc)
		require.NoError(t, err)
		require.NotNil(t, createdDoc.Id)
		doc3ID = *createdDoc.Id
	})

	t.Run("Step 5: Delete Document 2", func(t *testing.T) {
		env.MockFileProvider.EXPECT().DeleteFile(gomock.Any(), "file-id-doc2").
			Return(nil).Times(1)

		req, err := nethttp.NewRequest("DELETE", env.ServerURL+"/api/v1/DocumentReference/"+doc2ID, nil)
		require.NoError(t, err)
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusOK, resp.StatusCode)
	})

	t.Run("Step 6: List Documents - should have 2 documents", func(t *testing.T) {
		req, err := nethttp.NewRequest("GET", env.ServerURL+"/api/v1/DocumentReference?patient="+patientID, nil)
		require.NoError(t, err)
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var bundle models.Bundle
		err = json.Unmarshal(body, &bundle)
		require.NoError(t, err)

		assert.Equal(t, 2, *bundle.Total, "Should have 2 documents after deleting one")
		assert.Len(t, bundle.Entry, 2)

		var foundDoc1, foundDoc3 bool
		for _, entry := range bundle.Entry {
			var doc models.DocumentReference
			err := json.Unmarshal(entry.Resource, &doc)
			require.NoError(t, err)

			if *doc.Id == doc1ID {
				foundDoc1 = true
				assert.Len(t, doc.Content, 2, "Document 1 should have 2 attachments")
				assert.NotNil(t, doc.Subject)
				assert.Contains(t, *doc.Subject.Reference, patientID)
			}
			if *doc.Id == doc3ID {
				foundDoc3 = true
				assert.NotNil(t, doc.Subject)
				assert.Contains(t, *doc.Subject.Reference, patientID)
			}
		}

		assert.True(t, foundDoc1, "Document 1 should be in the list")
		assert.True(t, foundDoc3, "Document 3 should be in the list")
	})
}

func strPtr(s string) *string {
	return &s
}

func int64Ptr(i int64) *int64 {
	return &i
}
