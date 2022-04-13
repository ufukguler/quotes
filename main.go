package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"quotes/config"
	"quotes/database"
	m "quotes/middleware"
	"quotes/models"
	"quotes/routes"
	"quotes/services"
	"time"
)

func main() {
	err := os.Setenv("TZ", "Europe/Istanbul")
	if err != nil {
		panic(err)
	}
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Recover())
	e.Use(m.LogrusMiddleware)
	e.Use(m.CheckJwt)
	e.Use(m.SetLanguage)
	e.Use(m.RequestLogger)
	e.Use(m.AfterResponseMiddleware)

	e.Any("*", func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, models.Response{
			Code:    http.StatusNotFound,
			Message: "Not Found",
			Data:    nil,
		})
	})
	e.Validator = &config.CustomValidator{Validator: validator.New()}

	log.SetFormatter(&log.TextFormatter{
		EnvironmentOverrideColors: true,
		ForceColors:               true,
		FullTimestamp:             true,
		TimestampFormat:           time.RFC1123Z,
	})
	log.SetLevel(log.DebugLevel)

	config.LoadEnv()
	database.Database.Connect()
	routes.UseRoute(e)
	services.InitVersion()

	port := config.GetEnv("SERVER_PORT")
	if err := e.Start(port); err != nil {
		log.Panic(err)
	}
}
