package constants

import (
	"hack-mit/models"

	"github.com/labstack/echo/v4"
)

func CustomError(code int, message string) *echo.HTTPError {
	json := models.ResponseError{Code: code, Message: message}
	return echo.NewHTTPError(json.Code, json)
}

func ErrorDatabase() *echo.HTTPError {
	json := models.ResponseError{Code: 500, Message: "Database error"}
	return echo.NewHTTPError(json.Code, json)
}
