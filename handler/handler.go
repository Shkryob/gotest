package handler

import (
	"github.com/shkryob/gotest/model"
)

type Handler struct {
	userStore UserStore
}

type UserStore interface {
	GetByID(uint) *model.User
	GetByEmail(string) *model.User
	GetByUsername(string) *model.User
	Create(*model.User)
	ListUsers() []*model.User
	ClearUsers()
}

func NewHandler(us UserStore) *Handler {
	return &Handler{
		userStore: us,
	}
}
