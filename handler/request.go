package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/shkryob/gotest/model"
)

type userCreate struct {
	User struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
	} `json:"user"`
}

func (request *userCreate) bind(context echo.Context, user *model.User) error {
	if err := context.Bind(request); err != nil {
		return err
	}
	if err := context.Validate(request); err != nil {
		return err
	}
	user.Username = request.User.Username
	user.Email = request.User.Email
	return nil
}
