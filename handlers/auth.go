package handlers

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"quotes/models"
	"quotes/services"
)

func Login(c echo.Context) error {
	var dto models.LoginDTO
	if err := c.Bind(&dto); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}

	if err := c.Validate(&dto); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}

	token, err := services.GenerateJWT(dto)
	if err != nil {
		return e.ResponseError(err, c)
	}

	log.Info("Login request responded. JWT created for user: ", dto.DeviceId)
	return e.ResponseOk(token, c)
}
