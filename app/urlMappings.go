package app

import (
	"github.com/hsimao/bookstore_users-api/controllers/ping"
	"github.com/hsimao/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:userId", users.GetUser)
	router.PUT("/users/:userId", users.UpdateUser)
}
