package v1

import (
	"errors"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
	"zapvote/internal/api/response"
	"zapvote/internal/model/user"
	"zapvote/internal/security/accesstoken.go"
	"zapvote/internal/services/userstore"
)

type AuthController struct {
	userService userstore.Store
}

func NewAuthController(userService userstore.Store) *AuthController {
	return &AuthController{
		userService: userService,
	}
}

type authBody struct {
	RegNo    string `json:"reg_no"`
	Password string `json:"password"`
	DeviceID string `json:"device_id"`
}

func (ac *AuthController) AuthPOST(e echo.Context) error {
	body := &authBody{}
	if err := e.Bind(body); err != nil {
		return response.BadRequestError(e, "wrong body")
	}
	if body.RegNo == "" || body.Password == "" {
		return response.BadRequestError(e, "information required")
	}
	ID := strings.ReplaceAll(body.RegNo, "/", "")

	//check if user exists return a token
	u, err := ac.userService.GetUserSimple(ID)
	if err != nil && !errors.Is(err, userstore.ErrorUserNotFound) {
		return response.ServerError(e, err, "")
	}
	if u != nil {
		err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(body.Password))
		if err != nil {
			return response.OtherErrors(e, response.StatusWrongPassword, "wrong password")
		}
		token, err := accesstoken.GenerateForUser(u.ID)
		if err != nil {
			return response.ServerError(e, err, "")
		}
		msg := map[string]string{
			"token": token,
		}
		return response.JSON(e, msg)
	}

	//new user add them to db no device ID return
	c := strings.Split(body.RegNo, "/")
	var courseCode string
	if len(c) > 1 {
		courseCode = c[1]
	}

	if body.DeviceID == "" {
		return response.BadRequestError(e, "device id required")
	}

	hashedP, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.ServerError(e, err, "")
	}
	usr := &user.User{}
	usr.ID = ID
	usr.Password = string(hashedP)
	usr.DeviceID = body.DeviceID
	usr.CourseCode = courseCode
	usr.CreatedAt = time.Now()
	usr.UpdatedAt = usr.CreatedAt

	err = ac.userService.Create(usr)
	if err != nil {
		if errors.Is(err, userstore.ErrorUserExists) {
			return response.OtherErrors(e, response.StatusUserRegistered, "device already registered")
		}
		return response.ServerError(e, err, "")
	}
	token, err := accesstoken.GenerateForUser(ID)
	if err != nil {
		return response.ServerError(e, err, "")
	}
	msg := map[string]string{
		"token": token,
	}
	return response.JSON(e, msg)
}
