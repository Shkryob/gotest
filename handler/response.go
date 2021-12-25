package handler

import (
	"time"

	"github.com/shkryob/gotest/model"
)

type userResponse struct {
	User struct {
		ID        uint      `json:"id"`
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"createdAt"`
	} `json:"user"`
}

type userListResponse struct {
	Users []*userResponse `json:"users"`
}

func newUserListResponse(users []*model.User) *userListResponse {
	r := new(userListResponse)
	r.Users = make([]*userResponse, 0)
	for _, a := range users {
		ar := new(userResponse)
		ar.User.ID = a.ID
		ar.User.Username = a.Username
		ar.User.Email = a.Email
		ar.User.CreatedAt = a.CreatedAt

		r.Users = append(r.Users, ar)
	}
	return r
}

func newUserResponse(u *model.User) *userResponse {
	r := new(userResponse)
	r.User.ID = u.ID
	r.User.Username = u.Username
	r.User.Email = u.Email
	r.User.CreatedAt = u.CreatedAt
	return r
}
