package main

import (
	"github.com/shkryob/gotest/handler"
	"github.com/shkryob/gotest/router"
	"github.com/shkryob/gotest/store"
)

// @title GoTest API
// @version 1.0
// @description GoTest API
// @title GoTest API

// @host 127.0.0.1:1323
// @BasePath /api

// @schemes http https
// @produce	application/json
// @produce	application/xml
// @consumes application/json

func main() {
	r := router.New()

	v1 := r.Group("/api")

	us := store.NewUserStore()
	h := handler.NewHandler(us)
	h.Register(v1)
	r.Logger.Fatal(r.Start(":1323"))
}
