package tests

import (
	"net/http/httptest"
	"testing"
)

func TestPatientIntegration(t *testing.T) {
	env := SetupTestEnv(t)
	defer env.Cleanup()

	ts := httptest.NewServer(env.Handler)
	defer ts.Close()
	t.Run("Create Patient", func(t *testing.T) {

	})

	t.Run("Update Patient", func(t *testing.T) {

	})

	t.Run("Verify FHIR JSON", func(t *testing.T) {

	})
}