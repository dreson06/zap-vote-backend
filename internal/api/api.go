package api

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	"zapvote/internal/services/votestore"
)

type ConfigParams struct {
	DB             *sqlx.DB
	Mode           config.Mode
	UserStore      userstore.Store
	AdminStore     adminstore.Store
	CandidateStore candidatestore.Store
	ElectionStore  electionstore.Store
	VoteStore      votestore.Store
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

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173", "http://192.168.3.3:5173"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	group := e.Group("/api")
	group.Any("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	apiV1(group, conf)
	return e
}

func apiV1(group *echo.Group, conf *ConfigParams) {

	authController := v1.NewAuthController(conf.UserStore)
	userController := v1.NewUserController(conf.UserStore)
	candidateController := v1.NewCandidateController(conf.CandidateStore)
	electionController := v1.NewElectionController(conf.ElectionStore)
	voteController := v1.NewVoteController(conf.VoteStore, conf.UserStore, conf.ElectionStore, conf.DB)
	adminController := v1.NewAdminController(conf.AdminStore)

	group.GET("/election/:id", electionController.ElectionGET, auth.Auth)
	group.GET("/election/presidential", electionController.PresidentialCandidatesGET, auth.Auth)
	group.GET("/election/faculty/:faculty", electionController.FacultyCandidatesGET, auth.Auth)
	group.GET("/election/class/:course", electionController.ClassRepCandidatesGET, auth.Auth)
	group.GET("/election/votes/:id", electionController.TotalVotesGET, auth.Auth)

	group.GET("/results/faculty", electionController.FacultyResultsGET, auth.Auth)
	group.GET("/results/class/:code", electionController.ClassRepResultsGET, auth.Auth)

	//user routes
	group.POST("/user/auth", authController.AuthPOST)
	group.GET("/user/me", userController.MeGET, auth.Auth)

	//admin routes
	group.POST("/admin/auth", adminController.AuthPOST)
	group.POST("/candidate/add", candidateController.AddPOST, auth.AdminAuth)

	//candidates routes

	group.GET("/candidate/get", candidateController.CandidateDepartmentGET, auth.Auth)
	group.GET("/faculty/candidate/:id", candidateController.FacultyCandidateGET, auth.Auth)
	group.GET("/class/candidate/:id", candidateController.ClassRepCandidateGET, auth.Auth)
	group.GET("/presidential/candidate/:id", candidateController.PresidentialOneGET, auth.Auth)

	//vote routes
	group.POST("/vote/presidential", voteController.VotePresidentialPOST, auth.Auth)
	group.POST("/vote/faculty", voteController.FacultyVotePOST, auth.Auth)
	group.POST("/vote/class", voteController.VoteClassPOST, auth.Auth)
	group.GET("/has/voted/:election-id", voteController.HasVotedGET, auth.Auth)

}
