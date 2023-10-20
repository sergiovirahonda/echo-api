package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/sergiovirahonda/echo-api/internal/models"
	ec "github.com/sergiovirahonda/echo-api/internal/repositories/echo"
	echos "github.com/sergiovirahonda/echo-api/internal/services/echo"
	"gopkg.in/go-playground/assert.v1"
)

func TestHandlePostCreatesResource(t *testing.T) {
	service := echos.NewDefaultService(
		ec.NewDefaultRepository(database),
	)
	controller := NewController(
		echo.New().Group("/v0"),
		service,
	)
	echoJSON := `{"echo-me":"something"}`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/v0/echo/", strings.NewReader(echoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := controller.handlePost(c)
	assert.Equal(t, err, nil)
	assert.Equal(t, rec.Code, http.StatusCreated)
	assert.Equal(t, rec.Body.String(), "{\"echo-you\":\"something echo\"}\n")
	entities, err := service.GetAll(context.Background())
	assert.Equal(t, err, nil)
	assert.Equal(t, len(*entities), 1)
	assert.Equal(t, (*entities)[0].Value, "something")
	err = service.Delete(context.Background(), (*entities)[0].ID)
	assert.Equal(t, err, nil)
}

func TestHandleGetAllReturnsAllResources(t *testing.T) {
	service := echos.NewDefaultService(
		ec.NewDefaultRepository(database),
	)
	factory := models.NewEchoFactory()
	entity := factory.New(
		"something different",
	)
	err := service.Create(context.Background(), entity)
	assert.Equal(t, err, nil)
	controller := NewController(
		echo.New().Group("/v0"),
		service,
	)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/v0/whats-echoed/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err = controller.handleGetAll(c)
	assert.Equal(t, err, nil)
	assert.Equal(t, rec.Code, http.StatusOK)
	responseStr := rec.Body.Bytes()
	response := &models.EchoResponses{}
	err = json.Unmarshal(responseStr, response)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(response.Echos), 1)
	assert.Equal(t, response.Echos[0].Value, "something different")
	assert.Equal(t, err, nil)
	err = service.Delete(context.Background(), entity.ID)
	assert.Equal(t, err, nil)
}
