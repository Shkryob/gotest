package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shkryob/gotest/model"
	"github.com/shkryob/gotest/utils"
)

// AddUser godoc
// @Summary Register a new user
// @Description Register a new user
// @ID sign-up
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body userCreate true "User info for registration"
// @Success 201 {object} userResponse
// @Failure 400 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /users [post]
func (handler *Handler) AddUser(context echo.Context) error {
	var u model.User
	req := &userCreate{}
	if err := req.bind(context, &u); err != nil {
		return utils.ResponseByContentType(context, http.StatusUnprocessableEntity, utils.NewError(err))
	}
	existingUser := handler.userStore.GetByEmail(u.Email)
	if existingUser != nil {
		errorResponse := utils.NewError(errors.New("User with this email already exists"))
		return utils.ResponseByContentType(context, http.StatusUnprocessableEntity, errorResponse)
	}

	handler.userStore.Create(&u)

	return utils.ResponseByContentType(context, http.StatusCreated, newUserResponse(&u))
}

func (handler *Handler) ListUsers(context echo.Context) error {
	users := handler.userStore.ListUsers()
	return utils.ResponseByContentType(context, http.StatusOK, newUserListResponse(users))
}
