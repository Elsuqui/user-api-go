package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string `gorm:"<-"`
	Status   bool   `gorm:"default:true"`
}

type UserPublicInformation struct {
	gorm.Model
	Username string `json:"username"`
	Status   bool   `json:"status"`
}
