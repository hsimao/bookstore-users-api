package app

import "github.com/hsimao/bookstore_users-api/controllers"

func mapUrls() {
	router.GET("/", controllers.Ping)
}
