package config

import (
	"hack-mit/services/fue"
	"hack-mit/usecase"

	"github.com/labstack/echo/v4"
)

func InitHandlers(e *echo.Echo) {
	api := e.Group("/api/v1")
	api.GET("/test", func(c echo.Context) error { return c.String(200, "Hello World") })
	fueUC := usecase.NewFueUsecase()
	fue.NewFueHandler(api, fueUC)
}
