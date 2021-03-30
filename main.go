package main

import (
	"hack-mit/config"
	"os"

	_ "hack-mit/docs"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Hack MIT API
// @version 1.0
// @description This is a sample server for hack MIT.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email renattominayarojas@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @schemes https
func main() {
	config.SetAppConfig()
	dbs := config.InitDatabases()
	defer dbs.DeferDatabases()
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	config.InitHandlers(e)
	e.Logger.Fatal(e.StartTLS(os.Getenv("PORT"), "certs/Certificate.crt", "certs/Private.key"))
}
