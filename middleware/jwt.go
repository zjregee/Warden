package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

const secretKey = "secretKey"

func GetJWTInfo(c echo.Context) jwt.MapClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims
}

func GetJWTToken(data map[string]interface{}) string {
	t := jwt.New(jwt.SigningMethodHS256)
	claims := t.Claims.(jwt.MapClaims)
	for key, value := range data {
		claims[key] = value
	}
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	jwtToken, _ := t.SignedString([]byte(secretKey))
	return jwtToken
}

func GetJWTTokenWithClaims(claims jwt.Claims) string {
	t := jwt.New(jwt.SigningMethodHS256)
	t.Claims = claims
	jwtToken, _ := t.SignedString([]byte(secretKey))
	return jwtToken
}

func CustomJWTConfig(skipperPaths []string, authScheme string) middleware.JWTConfig {
	if authScheme == "" {
		authScheme = "Bearer"
	}
	return middleware.JWTConfig{
		SigningKey:  []byte(secretKey),
		TokenLookup: "header:" + echo.HeaderAuthorization,
		AuthScheme:  authScheme,
		Claims:      jwt.MapClaims{},
		ContextKey:  "user",
		Skipper:     CustomSkipper(skipperPaths),
	}
}

func CustomSkipper(skipperPaths []string) func(c echo.Context) bool {
	return func(c echo.Context) bool {
		for _, skipperPath := range skipperPaths {
			if c.Path() == skipperPath {
				return true
			}
		}
		return false
	}
}