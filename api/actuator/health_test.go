package actuator

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http/httptest"
	"testing"
)

func TestHealthController_Status(t *testing.T) {
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	health := new(HealthController)
	health.Status(ctx)

	assert.Equal(t, 200, recorder.Code)
	var h Health
	err := json.Unmarshal(recorder.Body.Bytes(), &h)
	if err != nil {
		t.Fatal()
	}
	assert.Equal(t, "ok", h.Status)
}
