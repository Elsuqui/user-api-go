package routes

import (
	"UserRestApi/controllers"

	"github.com/gin-gonic/gin"
)

type LoginRoutes struct {
	loginController controllers.LoginController
}

func (routes *LoginRoutes) InitializeRoutes(server *gin.Engine) {
	group := server.Group("auth")
	{
		group.POST("/token", func(c *gin.Context) { routes.loginController.Login(c) })
	}
}
