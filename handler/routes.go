package handler

import (
	"github.com/labstack/echo/v4"
)

func (handler *Handler) Register(v1 *echo.Group) {
	guestUsers := v1.Group("/users")
	guestUsers.GET("", handler.ListUsers)
	guestUsers.POST("", handler.AddUser)
}
