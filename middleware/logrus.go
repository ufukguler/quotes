package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"time"
)

func LogrusMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := c.Request()

		fields := map[string]interface{}{
			"time":       time.Now().Format(time.RFC3339),
			"remoteAddr": request.RemoteAddr,
			"uri":        request.RequestURI,
			"method":     request.Method,
		}
		logrus.WithFields(fields).Info("[API] ")

		if err := next(c); err != nil {
			logrus.Error(err)
		}
		return nil
	}
}
