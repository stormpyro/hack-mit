package main

import (
	"hack-mit/config"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.SetAppConfig()
	dbs := config.InitDatabases()
	defer dbs.DeferDatabases()
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	config.InitHandlers(e)
	e.Logger.Fatal(e.StartTLS(":9090", "certs/Certificate.crt", "certs/Private.key"))
}
