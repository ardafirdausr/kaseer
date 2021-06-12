package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func LogrusLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logrus.SetFormatter(&logrus.TextFormatter{
				ForceColors: true,
			})

			if c == nil {
				logrus.WithFields(logrus.Fields{
					"at": time.Now().Format("2006-01-02 15:04:05"),
				}).Info("Incoming Request")
				return next(c)
			}

			logrus.WithFields(logrus.Fields{
				"at":     time.Now().Format("2006-01-02 15:04:05"),
				"method": c.Request().Method,
				"uri":    c.Request().URL.String(),
				"ip":     c.Request().RemoteAddr,
			}).Info("Incoming Request")
			return next(c)
		}
	}
}

func EchoLogger() echo.MiddlewareFunc {
	config := middleware.LoggerConfig{
		Format:           "[${time_custom}] ${method} ${uri} (${status}) - ${latency_human} \t${error}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}
	return middleware.LoggerWithConfig(config)
}
