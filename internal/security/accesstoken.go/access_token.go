package accesstoken

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"zapvote/config"
)

func Generate(id string) (string, error) {
	tokenSecret := []byte(config.Cfg.AccessTokenSecret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	})
	tokenString, err := token.SignedString(tokenSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Validate(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Cfg.AccessTokenSecret), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["id"].(string), nil
	}
	return "", err
}
