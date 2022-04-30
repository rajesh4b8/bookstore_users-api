package app

import (
	"github.com/rajesh4b8/bookstore_users-api/controllers/ping"
	"github.com/rajesh4b8/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
