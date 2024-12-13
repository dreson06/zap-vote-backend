package v1

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"zapvote/internal/api/response"
	"zapvote/internal/security/accesstoken.go"
	"zapvote/internal/services/adminstore"
)

type AdminController struct {
	adminStore adminstore.Store
}

func NewAdminController(adminStore adminstore.Store) *AdminController {
	return &AdminController{adminStore: adminStore}
}

type adminBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (ac *AdminController) AuthPOST(e echo.Context) error {
	body := &adminBody{}
	if err := e.Bind(body); err != nil {
		return response.BadRequestError(e, "wrong body data")
	}
	if body.Username == "" || body.Password == "" {
		return response.BadRequestError(e, "missing information")
	}
	me, err := ac.adminStore.GetMe(body.Username)
	if err != nil {
		return response.ServerError(e, err, "server error")
	}

	err = bcrypt.CompareHashAndPassword([]byte(me.Password), []byte(body.Password))
	if err != nil {
		return response.UnauthorizedError(e)
	}

	token, err := accesstoken.GenerateForAdmin(me.ID)
	if err != nil {
		return response.ServerError(e, err, "server error")
	}
	msg := map[string]string{
		"access_token": token,
	}
	return response.JSON(e, msg)
}
