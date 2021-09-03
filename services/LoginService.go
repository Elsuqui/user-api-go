package services

import (
	"UserRestApi/helpers"
	"UserRestApi/models"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type LoginService struct {
	userService UserService
}

type CustomClaim struct {
	username string
	email    string
	jwt.StandardClaims
}

func (ref *LoginService) Login(username, password string) (string, error) {
	var token string
	user, err := ref.userService.FindByUserName(username)
	if err != nil {
		return token, err
	}
	fmt.Println(user.Password)
	isAuthenticated := helpers.CheckBcrypt(user.Password, password)
	fmt.Println(isAuthenticated)
	if isAuthenticated {
		return token, errors.New("incorrect password")
	}
	return getJwtToken(user)
}

func Logout() {

}

func getJwtToken(user models.User) (string, error) {
	generator := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaim{
		username: user.Username,
		email:    "mail@mail.com",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 1500,
			Issuer:    "Authentication service",
		},
	})
	secret := helpers.GetEnvParam("JWT_SECRET")
	return generator.SignedString(secret)
}
