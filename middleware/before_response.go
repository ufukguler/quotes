package middleware

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func AfterResponseMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().After(func() {
			log.Info("[RESPONSE] status is: ", c.Response().Status)
		})
		return next(c)
	}
}
