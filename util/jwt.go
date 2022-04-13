package util

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"time"
)

var JwtKeyString = "verycomplicatedjwtsecret"
var JwtKey = []byte(JwtKeyString)

func ExtractClaims(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid jwt token")
}

func GenerateJwt(deviceId string) (string, int64, error) {
	expire := time.Now().AddDate(0, 1, 0).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"deviceId":   deviceId,
		"iat":        time.Now().Unix(),
		"exp":        expire,
	})
	tokenString, err := token.SignedString(JwtKey)
	return tokenString, expire, err
}

func GetCurrentUserDeviceId(c echo.Context) (string, error) {
	bearer := c.Request().Header.Get("Authorization")
	if bearer == "" {
		log.Error("authorization header not valid: ", bearer)
		return "", errors.New("authorization header not valid")
	}

	token := bearer[7:]
	claims, err := ExtractClaims(token)
	if err != nil {
		log.Error("token not valid: ", err)
		return "", errors.New("token not valid")
	}
	return claims["deviceId"].(string), nil
}
