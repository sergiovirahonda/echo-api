package errors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewHTTPBadRequest(msg string) error {
	return echo.NewHTTPError(http.StatusBadRequest, msg)
}

func NewHTTPInternalServerError(msg string) error {
	return echo.NewHTTPError(http.StatusInternalServerError, msg)
}

func NewHTTPForbidden(msg string) error {
	return echo.NewHTTPError(http.StatusForbidden, msg)
}

func NewHTTPUnauthenticated(msg string) error {
	return echo.NewHTTPError(http.StatusUnauthorized, msg)
}

func NewHTTPInvalidJSON() error {
	return echo.NewHTTPError(http.StatusBadRequest, "invalid JSON")
}

func GetHTTPError(err error) error {
	switch err.(type) {
	case *BadRequest:
		return NewHTTPBadRequest(err.Error())
	default:
		return NewHTTPInternalServerError(err.Error())

	}
}
