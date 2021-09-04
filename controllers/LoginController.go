package controllers

import (
	"UserRestApi/helpers"
	"UserRestApi/services"
	"UserRestApi/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	loginService services.LoginService
}

func (ref *LoginController) Login(ctx *gin.Context) string {
	var validator validators.LoginPostValidator
	if err := ctx.ShouldBind(&validator); err != nil {
		errors, _ := helpers.ValidateRequestError(err, validators.DESCRIPTIVE)
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errors})
	}
	token, err := ref.loginService.Login(validator.Username, validator.Password)
	if err != nil {
		//ctx.Error(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return err.Error()
	}
	ctx.JSON(200, gin.H{"token": token})
	return token
}
