package tests

import (
	"bytes"
	"context"
	"fmt"
	"io"
	nethttp "net/http"
	"testing"

	"codex-documents/api/proto"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

const UPDATE_JSON_TEMPLATE = `{
		"resourceType": "Patient",
		"id": "%s",
		"active": false,
		"name": [
			{
				"use": "official",
				"family": "Ivanov",
				"given": ["Ivan", "Sergeevich"]
			}
		],
		"telecom": [
			{
				"system": "phone",
				"value": "+79001234567",
				"use": "mobile"
			},
			{
				"system": "email",
				"value": "ivan@example.com"
			}
		],
		"gender": "female",
		"birthDate": "1985-10-25",
		"deceasedBoolean": false,
		"address": [
			{
				"use": "home",
				"line": ["ul. Lenina, 1"],
				"city": "Moscow",
				"postalCode": "101000"
			}
		],
		"maritalStatus": {
			"coding": [
				{
					"system": "http://terminology.hl7.org/CodeSystem/v3-MaritalStatus",
					"code": "M",
					"display": "Married"
				}
			]
		},
		"communication": [
			{
				"language": {
					"coding": [
						{
							"system": "urn:ietf:bcp:47",
							"code": "ru",
							"display": "Russian"
						}
					]
				},
				"preferred": true
			}
		]
	}`

func createTestJWTToken(secret string, patientID string) (string, error) {
	claims := jwt.MapClaims{
		"sub":        "test-user",
		"patient_id": patientID,
		"scope":      "patient/*.read patient/*.write",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func TestPatientIntegration(t *testing.T) {
	env := SetupTestEnv(t)
	defer env.Cleanup()

	client := &nethttp.Client{}

	var patientID string
	var token string

	t.Run("Create Patient via gRPC", func(t *testing.T) {
		md := metadata.Pairs("x-internal-token", "test-secret")
		ctx := metadata.NewOutgoingContext(context.Background(), md)

		// Нам нужен только email, как договорились
		resp, err := env.GRPCClient.CreatePatient(ctx, &proto.CreatePatientRequest{
			Email: "ivan@example.com",
		})
		require.NoError(t, err)

		patientID = resp.PatientId
		require.NotEmpty(t, patientID)

		// Генерируем токен с полученным ID
		token, err = createTestJWTToken("secret-key", patientID)
		require.NoError(t, err)
	})

	t.Run("Update Patient", func(t *testing.T) {
		require.NotEmpty(t, patientID, "Patient ID should be set from Create step")

		updateJSON := fmt.Sprintf(UPDATE_JSON_TEMPLATE, patientID)

		req, err := nethttp.NewRequest("PUT", env.ServerURL+"/api/v1/Patient/"+patientID, bytes.NewBufferString(updateJSON))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/fhir+json")
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, nethttp.StatusOK, resp.StatusCode, "Expected status 200 OK")
	})

	t.Run("Verify FHIR JSON", func(t *testing.T) {
		require.NotEmpty(t, patientID, "Patient ID should be set from Create step")

		req, err := nethttp.NewRequest("GET", env.ServerURL+"/api/v1/Patient/"+patientID, nil)
		require.NoError(t, err)
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, nethttp.StatusOK, resp.StatusCode, "Expected status 200 OK")

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		expectedJSON := fmt.Sprintf(UPDATE_JSON_TEMPLATE, patientID)

		assert.JSONEq(t, expectedJSON, string(body), "Response JSON should match expected structure")
	})
}
