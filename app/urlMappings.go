package app

import (
	"github.com/hsimao/bookstore_users-api/controllers/ping"
	"github.com/hsimao/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.Create)
	router.GET("/users/:userId", users.Get)
	router.PUT("/users/:userId", users.Update)
	router.PATCH("/users/:userId", users.Update)
	router.DELETE("/users/:userId", users.Delete)
	router.GET("/internal/users/search", users.Search)
}
