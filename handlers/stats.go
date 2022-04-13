package handlers

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"quotes/services"
)

func GetStats(c echo.Context) error {
	stats := services.GetStats()
	log.Info("Get stats request responded")
	return e.ResponseOk(stats, c)
}
