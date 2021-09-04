package services

import (
	"UserRestApi/helpers"
	"UserRestApi/models"
	"UserRestApi/services/mysql"
)

type UserService struct {
	userModel models.User
}

func (ref *UserService) FindAll() []models.UserPublicInformation {
	users := []models.UserPublicInformation{}
	selectFields := []string{"id", "created_at", "updated_at", "deleted_at", "username", "status"}
	mysql.DB.Table("users").
		Select(selectFields).
		Find(&users)
	return users
}

func (ref *UserService) Find(id int) models.UserPublicInformation {
	var user models.UserPublicInformation
	mysql.DB.Table("users").
		Select("*").
		Where("id = ?", id).
		Scan(&user)
	return user
}

func (ref *UserService) Store(user models.User) (models.User, error) {
	pass, err := helpers.Bcrypt(user.Password)
	if err == nil {
		user.Password = pass
		err = mysql.DB.Create(&user).Error
		return user, err
	}
	return user, err
	/*return mysql.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			// return any error will rollback
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})*/
}

func (ref *UserService) Delete(id int) bool {
	err := mysql.DB.Delete(&models.User{}, id).
		Error
	return err == nil
}

func (ref *UserService) FindByUserName(username string) (models.User, error) {
	var user models.User
	err := mysql.DB.Table("users").Where("username = ?", username).Scan(&user).Error
	return user, err
}
