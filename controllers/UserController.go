package controllers

import (
	"UserRestApi/models"
	"UserRestApi/services"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	IController
}

type UserController struct {
	userControllerService services.UserService
}

func (controller *UserController) Index(ctx *gin.Context) []models.UserPublicInformation {
	users := controller.userControllerService.FindAll()
	ctx.JSON(200, users)
	return users
}

func (controller *UserController) Show(ctx *gin.Context) models.UserPublicInformation {
	param := ctx.Param("id")
	id, _ := strconv.Atoi(param)
	user := controller.userControllerService.Find(id)
	ctx.JSON(200, user)
	return user
}

func (controller *UserController) Store(ctx *gin.Context) models.User {
	var newUser models.User
	ctx.ShouldBind(&newUser)
	newUser, _ = controller.userControllerService.Store(newUser)
	ctx.JSON(200, gin.H{
		"id":         newUser.ID,
		"created_at": newUser.CreatedAt,
		"username":   newUser.Username,
		"status":     newUser.Status,
	})
	return newUser
}

func (controller *UserController) Delete(ctx *gin.Context) bool {
	if id, err := strconv.Atoi(ctx.Param("id")); err == nil {
		fmt.Println()
		ok := controller.userControllerService.Delete(id)
		ctx.JSON(200, ok)
		return ok
	}
	ctx.JSON(500, "Server error, identifier need to be an integer")
	return false
}
