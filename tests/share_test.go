//go:build integration

package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

func TestShareIntegration(t *testing.T) {
	env := SetupTestEnv(t)
	defer env.Cleanup()

	env.MockFileProvider.EXPECT().GetPresignedUrls(gomock.Any(), gomock.Any()).Return(&domain.PresignedUrlsResponse{
		FileId:      "test-file-id",
		UploadUrl:   "http://test/upload",
		DownloadUrl: "http://test/download",
	}, nil).AnyTimes()

	client := &nethttp.Client{}

	var patientID string
	var token string
	var doc1ID, doc2ID string
	var obs1ID, obs2ID string
	var tmpToken string

	t.Run("Setup: Create Patient via gRPC", func(t *testing.T) {
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

	t.Run("Step 1: Create Document 1", func(t *testing.T) {
		doc := models.DocumentReference{
			ResourceType: "DocumentReference",
			Status:       "current",
			Content: []models.DocumentReferenceContent{
				{
					Attachment: &models.Attachment{
						ContentType: strPtr("application/pdf"),
						Size:        int64Ptr(1024),
						Title:       strPtr("Document 1"),
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

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var createdDoc models.DocumentReference
		err = json.Unmarshal(body, &createdDoc)
		require.NoError(t, err)
		require.NotNil(t, createdDoc.Id)
		doc1ID = *createdDoc.Id
		require.NotNil(t, createdDoc.Content[0].Attachment.Id)
	})

	t.Run("Step 2: Create Document 2", func(t *testing.T) {
		doc := models.DocumentReference{
			ResourceType: "DocumentReference",
			Status:       "current",
			Content: []models.DocumentReferenceContent{
				{
					Attachment: &models.Attachment{
						ContentType: strPtr("image/jpeg"),
						Size:        int64Ptr(2048),
						Title:       strPtr("Document 2"),
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

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var createdDoc models.DocumentReference
		err = json.Unmarshal(body, &createdDoc)
		require.NoError(t, err)
		require.NotNil(t, createdDoc.Id)
		doc2ID = *createdDoc.Id
	})

	t.Run("Step 3: Create Observation 1 with DerivedFrom to Document 1", func(t *testing.T) {
		doc1Ref := fmt.Sprintf("DocumentReference/%s", doc1ID)
		obs := models.Observation{
			ResourceType: "Observation",
			Status:       "final",
			Code: &models.CodeableConcept{
				Coding: []models.Coding{
					{
						System:  strPtr("http://loinc.org"),
						Code:    strPtr("85354-9"),
						Display: strPtr("Blood pressure"),
					},
				},
			},
			EffectiveDateTime: strPtr("2024-01-15T10:30:00Z"),
			ValueQuantity: &models.Quantity{
				Value:  float64Ptr(120.0),
				Unit:   strPtr("mmHg"),
				System: strPtr("http://unitsofmeasure.org"),
				Code:   strPtr("mm[Hg]"),
			},
			DerivedFrom: []models.Reference{
				{
					Reference: &doc1Ref,
				},
			},
		}

		obsJSON, err := json.Marshal(obs)
		require.NoError(t, err)

		req, err := nethttp.NewRequest("POST", env.ServerURL+"/api/v1/Observation", bytes.NewBuffer(obsJSON))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/fhir+json")
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusCreated, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var createdObs models.Observation
		err = json.Unmarshal(body, &createdObs)
		require.NoError(t, err)
		require.NotNil(t, createdObs.Id)
		obs1ID = *createdObs.Id
	})

	t.Run("Step 4: Create Observation 2 with DerivedFrom to both Documents", func(t *testing.T) {
		doc1Ref := fmt.Sprintf("DocumentReference/%s", doc1ID)
		doc2Ref := fmt.Sprintf("DocumentReference/%s", doc2ID)
		obs := models.Observation{
			ResourceType: "Observation",
			Status:       "preliminary",
			Code: &models.CodeableConcept{
				Coding: []models.Coding{
					{
						System:  strPtr("http://loinc.org"),
						Code:    strPtr("2093-3"),
						Display: strPtr("Cholesterol"),
					},
				},
			},
			EffectiveDateTime: strPtr("2024-01-15T12:00:00Z"),
			ValueQuantity: &models.Quantity{
				Value:  float64Ptr(200.0),
				Unit:   strPtr("mg/dL"),
				System: strPtr("http://unitsofmeasure.org"),
				Code:   strPtr("mg/dL"),
			},
			DerivedFrom: []models.Reference{
				{
					Reference: &doc1Ref,
				},
				{
					Reference: &doc2Ref,
				},
			},
		}

		obsJSON, err := json.Marshal(obs)
		require.NoError(t, err)

		req, err := nethttp.NewRequest("POST", env.ServerURL+"/api/v1/Observation", bytes.NewBuffer(obsJSON))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/fhir+json")
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusCreated, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var createdObs models.Observation
		err = json.Unmarshal(body, &createdObs)
		require.NoError(t, err)
		require.NotNil(t, createdObs.Id)
		obs2ID = *createdObs.Id
	})

	t.Run("Step 5: Create Share for Observation 1", func(t *testing.T) {
		require.NotEmpty(t, obs1ID, "Observation 1 ID should be set")

		shareReq := map[string]interface{}{
			"resource_ids": []string{obs1ID},
			"ttl_seconds":  3600,
		}

		shareReqJSON, err := json.Marshal(shareReq)
		require.NoError(t, err)

		env.MockTmpAccessClient.EXPECT().GenerateTmpToken(
			gomock.Any(),
			gomock.Any(),
		).DoAndReturn(func(ctx context.Context, req domain.GenerateTmpTokenRequest) (*domain.GenerateTmpTokenResponse, error) {
			assert.Contains(t, req.Payload["scopes"], "docs:observation:"+obs1ID+":read")
			assert.Contains(t, req.Payload["scopes"], "docs:document_reference:"+doc1ID+":read")
			assert.Contains(t, req.Payload["scopes"], "files:file:test-file-id:read")
			assert.Equal(t, int64(3600), req.TtlSeconds)

			return &domain.GenerateTmpTokenResponse{
				TmpToken: "mock-tmp-token-12345",
			}, nil
		}).Times(1)

		req, err := nethttp.NewRequest("POST", env.ServerURL+"/api/v1/share", bytes.NewBuffer(shareReqJSON))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var shareResp domain.ShareResponse
		err = json.Unmarshal(body, &shareResp)
		require.NoError(t, err)
		assert.Equal(t, "mock-tmp-token-12345", shareResp.Token)
		assert.Equal(t, "/api/v1/shared", shareResp.ResourceURL)

		tmpToken = shareResp.Token
	})

	t.Run("Step 6: Generate tmp JWT token with scopes", func(t *testing.T) {
		require.NotEmpty(t, tmpToken, "Tmp token should be set from previous step")

		expectedScopes := fmt.Sprintf("docs:observation:%s:read,docs:document_reference:%s:read,files:file:test-file-id:read", obs1ID, doc1ID)

		claims := jwt.MapClaims{
			"scopes": expectedScopes,
		}
		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tmpJWTToken, err := jwtToken.SignedString([]byte("secret-key"))
		require.NoError(t, err)

		tmpToken = tmpJWTToken
	})

	t.Run("Step 7: Get shared resources list", func(t *testing.T) {
		require.NotEmpty(t, tmpToken, "Tmp token should be set")

		req, err := nethttp.NewRequest("GET", env.ServerURL+"/api/v1/shared", nil)
		require.NoError(t, err)
		req.Header.Set("Authorization", "Bearer "+tmpToken)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var sharedResp domain.SharedResourcesResponse
		err = json.Unmarshal(body, &sharedResp)
		require.NoError(t, err)

		assert.Contains(t, sharedResp.Observations, fmt.Sprintf("/api/v1/Observation/%s", obs1ID))
		assert.Contains(t, sharedResp.DocumentReferences, fmt.Sprintf("/api/v1/DocumentReference/%s", doc1ID))
	})

	t.Run("Step 8: Get Observation 1 with tmp token", func(t *testing.T) {
		require.NotEmpty(t, tmpToken, "Tmp token should be set")

		req, err := nethttp.NewRequest("GET", env.ServerURL+"/api/v1/Observation/"+obs1ID, nil)
		require.NoError(t, err)
		req.Header.Set("Authorization", "Bearer "+tmpToken)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var obs models.Observation
		err = json.Unmarshal(body, &obs)
		require.NoError(t, err)
		assert.Equal(t, obs1ID, *obs.Id)
	})

	t.Run("Step 9: Get DocumentReference 1 with tmp token", func(t *testing.T) {
		require.NotEmpty(t, tmpToken, "Tmp token should be set")

		req, err := nethttp.NewRequest("GET", env.ServerURL+"/api/v1/DocumentReference/"+doc1ID, nil)
		require.NoError(t, err)
		req.Header.Set("Authorization", "Bearer "+tmpToken)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var doc models.DocumentReference
		err = json.Unmarshal(body, &doc)
		require.NoError(t, err)
		assert.Equal(t, doc1ID, *doc.Id)
	})

	t.Run("Step 10: Try to get Observation 2 with tmp token (should fail)", func(t *testing.T) {
		require.NotEmpty(t, tmpToken, "Tmp token should be set")

		req, err := nethttp.NewRequest("GET", env.ServerURL+"/api/v1/Observation/"+obs2ID, nil)
		require.NoError(t, err)
		req.Header.Set("Authorization", "Bearer "+tmpToken)

		resp, err := client.Do(req)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusForbidden, resp.StatusCode)
	})

	t.Run("Step 11: Try to create Observation with tmp token (should fail)", func(t *testing.T) {
		require.NotEmpty(t, tmpToken, "Tmp token should be set")

		obs := models.Observation{
			ResourceType: "Observation",
			Status:        "final",
			Code: &models.CodeableConcept{
				Coding: []models.Coding{
					{
						System:  strPtr("http://loinc.org"),
						Code:    strPtr("718-7"),
						Display: strPtr("Hemoglobin"),
					},
				},
			},
		}

		obsJSON, err := json.Marshal(obs)
		require.NoError(t, err)

		req, err := nethttp.NewRequest("POST", env.ServerURL+"/api/v1/Observation", bytes.NewBuffer(obsJSON))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/fhir+json")
		req.Header.Set("Authorization", "Bearer "+tmpToken)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusForbidden, resp.StatusCode)
	})

	t.Run("Step 12: Try to list Observations with tmp token (should fail)", func(t *testing.T) {
		require.NotEmpty(t, tmpToken, "Tmp token should be set")

		req, err := nethttp.NewRequest("GET", env.ServerURL+"/api/v1/Observation?patient="+patientID, nil)
		require.NoError(t, err)
		req.Header.Set("Authorization", "Bearer "+tmpToken)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusForbidden, resp.StatusCode)
	})
}
