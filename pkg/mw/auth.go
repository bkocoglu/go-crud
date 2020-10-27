package mw

import (
	"encoding/base64"
	"fmt"
	_const "github.com/bilalkocoglu/go-crud/pkg/const"
	"github.com/bilalkocoglu/go-crud/pkg/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request
		basicToken := req.Header.Get(_const.AuthorizationHeader)
		if basicToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token must be not null.",
			})
			return
		}
		tokenType, token := parseToken(basicToken)
		if tokenType != "Basic" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token type must be Basic.",
			})
			return
		}
		decodeToken, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token decode exception",
			})
			return
		}
		splitToken := strings.Split(string(decodeToken), ":")
		username := splitToken[0]
		password := splitToken[1]

		var user database.User
		err = database.GetUserByUsername(&user, username)
		if err != nil || user.ID == 0 || user.Password != password {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid username or password",
			})
			return
		}
		c.Set(_const.CurrentUser, user)

		c.Next()
	}
}

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request

		bearerToken := req.Header.Get(_const.AuthorizationHeader)

		if bearerToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token must be not null.",
			})
			return
		}

		tokenType, tokenString := parseToken(bearerToken)

		if tokenType != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token type must be Bearer.",
			})
			return
		}

		token, err := VerifyJwtToken(tokenString)
		if err != nil {
			log.Error().Err(err).Msg("Crash verify jwt token.")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {
			log.Error().Msg("Can not create claims map.")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			return
		}

		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			log.Error().Err(err).Msg("user_id claims not found in jwt.")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			return
		}

		var user database.User
		err = database.GetUserById(&user, userId)
		if err != nil || user.ID == 0 {
			log.Error().Err(err).Msg("User not found for jwt user id")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			return
		}
		c.Set(_const.CurrentUser, user)

		c.Next()
	}
}

func VerifyJwtToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func parseToken(token string) (string, string) {
	parsedToken := strings.Split(token, " ")

	return parsedToken[0], parsedToken[1]
}
