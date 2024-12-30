package api

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"net/http"
	"zapvote/config"
	"zapvote/internal/api/middleware/auth"
	"zapvote/internal/api/middleware/ratelimiter"
	"zapvote/internal/api/middleware/simplelog"
	v1 "zapvote/internal/api/v1"
	"zapvote/internal/services/adminstore"
	"zapvote/internal/services/candidatestore"
	"zapvote/internal/services/electionstore"
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

	authController := v1.NewAuthController(userService)
	userController := v1.NewUserController(userService)

	adminService := adminstore.NewSQLStore(conf.DB)
	adminController := v1.NewAdminController(adminService)

	candidateService := candidatestore.NewSQLStore(conf.DB)
	candidateController := v1.NewCandidateController(candidateService)

	electionService := electionstore.NewSQLStore(conf.DB)
	electionController := v1.NewElectionController(electionService)

	group.GET("/election/presidential", electionController.PresidentialCandidatesGET, auth.Auth)
	group.GET("/election/faculty/:faculty", electionController.FacultyCandidatesGET, auth.Auth)
	group.GET("/election/class/:course", electionController.ClassRepCandidatesGET, auth.Auth)

	group.GET("/results/faculty", electionController.FacultyResultsGET, auth.Auth)
	group.GET("/results/class/:code", electionController.ClassRepResultsGET, auth.Auth)

	//user routes
	group.POST("/user/auth", authController.AuthPOST)
	group.GET("/user/me", userController.MeGET, auth.Auth)

	//admin routes
	group.POST("/admin/auth", adminController.AuthPOST)
	group.POST("/candidate/add", candidateController.AddPOST, auth.AdminAuth)

	//candidates routes
	group.GET("/candidate/get", candidateController.GetFromDepartment, auth.Auth)

}
