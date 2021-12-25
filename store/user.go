package store

import (
	"time"

	"github.com/shkryob/gotest/model"
)

type UserStore struct {
	usersList []*model.User
}

func NewUserStore() *UserStore {
	return &UserStore{usersList: make([]*model.User, 0)}
}

func (us *UserStore) GetByID(id uint) *model.User {
	for _, v := range us.usersList {
		if v.ID == id {
			return v
		}
	}

	return nil
}

func (us *UserStore) Create(u *model.User) {
	u.ID = uint(len(us.usersList) + 1)
	u.CreatedAt = time.Now()

	us.usersList = append(us.usersList, u)
}

func (us *UserStore) GetByEmail(e string) *model.User {
	for _, v := range us.usersList {
		if v.Email == e {
			return v
		}
	}

	return nil
}

func (us *UserStore) GetByUsername(username string) *model.User {
	for _, v := range us.usersList {
		if v.Username == username {
			return v
		}
	}

	return nil
}

func (us *UserStore) ListUsers() []*model.User {
	return us.usersList
}

func (us *UserStore) ClearUsers() {
	us.usersList = make([]*model.User, 0)
}
