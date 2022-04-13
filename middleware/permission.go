package middleware

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"quotes/models"
	. "quotes/services"
	"quotes/util"
	"strings"
)

var e models.Response

func CheckJwt(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Path()
		if strings.HasPrefix("/api/auth/login", path) {
			return next(c)
		}

		bearer := c.Request().Header.Get("Authorization")
		if bearer == "" {
			log.Error("authorization header not valid: ", bearer)
			return e.ResponseUnauthorized(c)
		}

		token := bearer[7:]
		claims, err := util.ExtractClaims(token)
		if err != nil {
			log.Error("token not valid: ", err)
			return e.ResponseError(err, c)
		}

		if strings.Contains(path, "/admin") {
			roles := claims["role"].([]interface{})
			if !contains(roles, RoleAdmin) {
				return e.ResponseForbidden(c)
			}
			log.Info("Current admin user is: ", claims["email"].(string))
		} else {
			c.Set("deviceId", claims["deviceId"].(string))
			log.Info("Current user is: ", claims["deviceId"].(string))
		}
		return next(c)
	}
}

func contains(s []interface{}, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
