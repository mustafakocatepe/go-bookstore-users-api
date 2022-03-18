package app

import (
	"github.com/mustafakocatepe/go-bookstore-users-api/controllers/ping"
	"github.com/mustafakocatepe/go-bookstore-users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.Get)
	//router.GET("/users/search", controllers.Search)
	router.POST("/users", users.Create)
	router.PUT("/users/users_id", users.Update)
	//PATCH("/users/users_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
}
