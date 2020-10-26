package mw

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
	"time"
)

func SetInterceptors(g *echo.Group) {
	g.Use(RequestID)
	g.Use(ErrorLog)
	g.Use(RequestLog)
}

func RequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		req := context.Request()
		res := context.Response()

		requestId := req.Header.Get("req-id")

		if requestId == "" {
			requestId = uuid.NewV4().String()
		}

		res.Header().Set("req-id", requestId)

		return next(context)
	}
}

func ErrorLog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		if err = next(c); err != nil {
			log.Info().Str("Request-Id", c.Response().Header().Get("req-id")).Msgf("error occurred: %+v", err)
			c.Error(err)
		}

		return
	}
}

func RequestLog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		start := time.Now()

		if err = next(c); err != nil {
			c.Error(err)
		}

		stop := time.Now()

		requestId := c.Response().Header().Get("req-id")
		path := c.Request().URL.Path
		method := c.Request().Method

		log.Info().Str("RequestId", requestId).Str("RealIP", c.RealIP()).Str("method", method).Str("path", path).Int("status", c.Response().Status).Int64("length", c.Response().Size).Str("latency", stop.Sub(start).String()).Msg("request")

		return
	}
}
