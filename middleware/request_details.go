package middleware

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func RequestLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Infof("[REQUEST] Headers: %v", prettyPrint(c.Request().Header))
		return next(c)
	}
}

func prettyPrint(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return ""
	}
	return string(b)
}
