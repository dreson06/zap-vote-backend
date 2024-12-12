package v1

import (
	"github.com/labstack/echo/v4"
	"zapvote/internal/api/middleware/auth"
	"zapvote/internal/api/response"
	"zapvote/internal/services/userstore"
)

type UserController struct {
	userStore userstore.Store
}

func NewUserController(userStore userstore.Store) *UserController {
	return &UserController{
		userStore: userStore,
	}
}

func (uc *UserController) MeGET(e echo.Context) error {
	id := auth.GetID(e)
	u, err := uc.userStore.GetMe(id)
	if err != nil {
		return response.ServerError(e, err, "")
	}
	return response.JSON(e, u)
}
