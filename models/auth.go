package models

import "gorm.io/gorm"

type AuthToken struct {
	gorm.Model
	UserId int
	User   User
}
