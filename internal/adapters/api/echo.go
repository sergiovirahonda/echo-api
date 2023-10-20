package api

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sergiovirahonda/echo-api/internal/errors"
	"github.com/sergiovirahonda/echo-api/internal/models"
	echos "github.com/sergiovirahonda/echo-api/internal/services/echo"
)

type Controller struct {
	service echos.Service
}

func NewController(
	routes *echo.Group,
	s echos.Service,
) {
	c := &Controller{
		service: s,
	}

	routes.POST("/echo/", c.handlePost)
	routes.GET("/whats-echoed/", c.handleGetAll)
}

// Controllers

// Reservation godoc
// @Summary Creates an Echo resource.
// @Description Creates an Echo resource and returns it echoed.
// @Tags echo
// @Accept json
// @Produce json
// @Param resource body models.EchoRequest true "Echo object"
// @Success 201 {object} models.EchoResponseFromRequest
// @Router /v0/echo/ [POST]
func (c *Controller) handlePost(ctx echo.Context) error {
	payload := &models.EchoRequest{}
	defer ctx.Request().Body.Close()
	err := json.NewDecoder(ctx.Request().Body).Decode(&payload)
	if err != nil {
		return errors.GetHTTPError(err)
	}
	entity, err := c.service.CreateFromRequest(
		ctx.Request().Context(),
		payload,
	)
	if err != nil {
		return errors.GetHTTPError(err)
	}
	response := entity.ToResponseFromRequest()
	return ctx.JSON(http.StatusCreated, response)
}

// Reservation godoc
// @Summary Gets all Echo resources.
// @Description Gets all Echo resources.
// @Tags echo
// @Produce json
// @Success 200 {object} models.EchoResponses
// @Router /v0/whats-echoed/ [GET]
func (c *Controller) handleGetAll(ctx echo.Context) error {
	entities, err := c.service.GetAll(ctx.Request().Context())
	if err != nil {
		return errors.GetHTTPError(err)
	}
	response := entities.ToResponses()
	return ctx.JSON(http.StatusOK, response)
}
