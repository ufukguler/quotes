package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"quotes/models"
	"quotes/services"
	"strconv"
)

func Version(c echo.Context) error {
	version := c.QueryParam("version")
	if version == "" {
		return e.ResponseError(errors.New("version parameter is empty"), c)
	}
	versionType := c.QueryParam("device")
	if versionType == "" {
		return e.ResponseError(errors.New("version device type parameter is empty"), c)
	}

	i, err := strconv.ParseInt(version, 10, 64)
	if err != nil {
		return e.ResponseError(errors.New("version parameter is not a number"), c)
	}
	getVersion, err := services.GetVersion(versionType)
	if err != nil {
		return e.ResponseError(err, c)
	}

	ver := models.VersionResponse{
		MinimumVersion: getVersion.Minimum,
		ShouldUpdate:   false,
		ForceUpdate:    false,
	}

	if i < getVersion.Minimum {
		ver.ForceUpdate = true
	}
	if i < getVersion.Current {
		ver.ShouldUpdate = true
	}
	log.Info("Version request responded successfully. Parameters: version:", version, ", device:", versionType)
	return e.ResponseOk(ver, c)
}
