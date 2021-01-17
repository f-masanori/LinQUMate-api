package api

import (
	"linqumate/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUsecase *usecase.UserUsecase
}

func NewUserController(userUsecase *usecase.UserUsecase) *UserController {
	return &UserController{
		UserUsecase: userUsecase,
	}
}

func (u *UserController) UserCreate(context echo.Context) error {
	return context.String(http.StatusOK, "Create User!")
}
