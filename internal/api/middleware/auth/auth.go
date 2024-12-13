package auth

import (
	"github.com/labstack/echo/v4"
	"zapvote/internal/api/response"
	"zapvote/internal/security/accesstoken.go"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		t := e.Request().Header.Get("Access-Token")
		if t == "" {
			return response.UnauthorizedError(e)
		}
		userId, err := accesstoken.ValidateUser(t)
		if err != nil {
			return response.UnauthorizedError(e)
		}
		e.Set("user-id", userId)
		return next(e)
	}
}

func GetID(e echo.Context) string {
	id := e.Get("user-id")
	if id != nil {
		return id.(string)
	}
	panic("user id not found")
	return ""
}
