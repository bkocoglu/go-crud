package mw

import (
	_const "github.com/bilalkocoglu/go-crud/pkg/const"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

func BasicAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		req := context.Request()

		basicToken := req.Header.Get(_const.AuthorizationHeader)

		if basicToken == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token must be not null.")
		}

		tokenType, token := parseToken(basicToken)

		if tokenType != "Basic" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token type must be Basic.")
		}

		log.Info().Msg("tokenType: " + tokenType + " token: " + token)

		return next(context)
	}
}

func JwtAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		req := context.Request()

		bearerToken := req.Header.Get(_const.AuthorizationHeader)

		if bearerToken == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token must be not null.")
		}

		tokenType, token := parseToken(bearerToken)

		if tokenType != "Bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token type must be Bearer.")
		}

		log.Info().Msg("tokenType: " + tokenType + " token: " + token)

		return next(context)
	}
}

func parseToken(token string) (string, string) {
	parsedToken := strings.Split(token, " ")

	return parsedToken[0], parsedToken[1]
}
