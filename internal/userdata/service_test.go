package userdata

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/artprocessors/nsm-microservice-golang-userdata/restapi/models"

	"github.com/go-openapi/runtime"

	"github.com/artprocessors/nsm-microservice-golang-userdata/restapi/operations"
	"github.com/stretchr/testify/assert"
)

var userDataService Service

func TestMain(m *testing.M) {
	userDataService = New()
	os.Exit(m.Run())
}

func TestGetHealth(t *testing.T) {
	t.Run("test OK response", func(t *testing.T) {
		response := userDataService.GetHealth(operations.GetHealthParams{})
		assert.NotNil(t, response)

		recorder := httptest.NewRecorder()
		response.WriteResponse(recorder, runtime.JSONProducer())
		assert.Equal(t, http.StatusOK, recorder.Code)

		var healthResponse models.HealthResponse
		if err := json.Unmarshal(recorder.Body.Bytes(), &healthResponse); err != nil {
			assert.Fail(t, err.Error())
		}
		assert.Equal(t, "UP", *healthResponse.Status)
	})
}
