// routes.go
package server

import (
	"github.com/cyctw/go-crud/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// Serve static file
	api := router.Group("api/v1")
	{
		userGroup := api.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.GET("/:id", user.Retrieve)
		}
		authGroup := api.Group("auth")
		{
			auth := new(controllers.AuthController)
			authGroup.POST("/signup", auth.Signup)
			authGroup.POST("/login", auth.Login)
			authGroup.POST("/logout", auth.Logout)
		}

	}
	router.Static("/static", "./dist")

	// Middlewares
	return router
}
