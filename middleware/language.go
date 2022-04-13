package middleware

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"quotes/util"
	"strings"
)

func SetLanguage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var lang string
		header := c.Request().Header
		acceptLanguage := header["Accept-Language"]

		if len(acceptLanguage) == 0 {
			lang = "EN"
		} else {
			givenLang := strings.ToUpper(acceptLanguage[0])
			if !util.IsLangValid(givenLang) {
				log.Errorln("Given language '" + givenLang + "' is not valid. Setting back to default.")
				givenLang = "EN"
			}
			lang = givenLang
		}
		c.Set("Accept-Language", lang)
		return next(c)
	}
}
