package api

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"net/http"
	"zapvote/config"
	"zapvote/internal/api/middleware/ratelimiter"
	"zapvote/internal/api/middleware/simplelog"
	v1 "zapvote/internal/api/v1"
	"zapvote/internal/services/userstore"
)

type ConfigParams struct {
	DB   *sqlx.DB
	Mode config.Mode
}

func Init(conf *ConfigParams) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(ratelimiter.InitEchoLimiter())
	e.Any("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.Use(simplelog.Logger)

	group := e.Group("/api")
	group.Any("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	apiV1(group, conf)
	return e
}

func apiV1(group *echo.Group, conf *ConfigParams) {
	userService := userstore.NewSqlStore(conf.DB)
	userController := v1.NewAuthController(userService)

	group.POST("/auth", userController.AuthPOST)
}