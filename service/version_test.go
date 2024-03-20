package service

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"work.ctyun.cn/git/GoStack/gostone/conf"
)

func TestGetVersion3(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Create a request
	req := httptest.NewRequest(http.MethodGet, "/version3", nil)

	// Create a response recorder
	rec := httptest.NewRecorder()

	// Create an Echo context
	ctx := e.NewContext(req, rec)

	// Call the handler
	err := GetVersion3(ctx)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, rec.Code)

	// Parse the response body
	var response map[string]interface{}
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response))

	// Assert the expected version value
	expectedVersion := map[string]interface{}{"version": conf.Version3}
	assert.Equal(t, expectedVersion, response)

	// Assert that there is no error returned
	assert.NoError(t, err)
}

// Similar tests can be written for GetVersion2 and GetAllVersion
