package handler

import (
	"os"
	"strings"
	"testing"

	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"github.com/shkryob/gotest/model"
	"github.com/shkryob/gotest/router"
	"github.com/shkryob/gotest/store"
	"github.com/stretchr/testify/assert"
)

var (
	us UserStore
	h  *Handler
	e  *echo.Echo
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func TestListUsersCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	e := router.New()
	req := httptest.NewRequest(echo.GET, "/api/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.NoError(t, h.ListUsers(c))
	if assert.Equal(t, http.StatusOK, rec.Code) {
		var aa userListResponse
		err := json.Unmarshal(rec.Body.Bytes(), &aa)
		assert.NoError(t, err)
		assert.Equal(t, "user1", aa.Users[0].User.Username)
	}
}

func TestCreateUsersCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	e := router.New()
	reqJSON := `{"user":{"username":"user2", "email":"email2@mail.test"}}`
	req := httptest.NewRequest(echo.POST, "/api/users", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.NoError(t, h.AddUser(c))

	if assert.Equal(t, http.StatusCreated, rec.Code) {
		var aa userResponse
		err := json.Unmarshal(rec.Body.Bytes(), &aa)
		assert.NoError(t, err)
		assert.Equal(t, "user2", aa.User.Username)
		assert.Equal(t, uint(2), aa.User.ID)
	}
}

func setup() {
	us = store.NewUserStore()
	h = NewHandler(us)
	e = router.New()
	loadFixtures()
}

func responseMap(b []byte, key string) map[string]interface{} {
	var m map[string]interface{}
	json.Unmarshal(b, &m)

	return m[key].(map[string]interface{})
}

func loadFixtures() error {
	u1 := model.User{
		Username: "user1",
		Email:    "user1@test.io",
	}
	us.Create(&u1)

	return nil
}

func tearDown() {
	us.ClearUsers()
}
