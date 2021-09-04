package routes

import (
	"UserRestApi/controllers"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userController controllers.UserController
}

func (routes *UserRoutes) InitializeRoutes(server *gin.Engine) {
	group := server.Group("/user")
	{
		group.GET("/", func(c *gin.Context) { routes.userController.Index(c) })
		group.GET("/:id", func(c *gin.Context) { routes.userController.Show(c) })
		group.POST("/", func(c *gin.Context) { routes.userController.Store(c) })
		group.DELETE("/:id", func(c *gin.Context) { routes.userController.Delete(c) })
	}
}
