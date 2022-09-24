package database

import (
	"errors"

	"github.com/suryaadi44/go-training-restful/config"
	"github.com/suryaadi44/go-training-restful/models"
)

var (
	// ErrInvalidID is returned when an invalid ID is provided to a method like Delete.
	ErrInvalidID = errors.New("provided id was invalid")
)

func GetUsers() (interface{}, error) {
	var users []models.User

	err := config.DB.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func GetUserByID(id int) (interface{}, error) {
	var user models.User

	err := config.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(user models.User) (interface{}, error) {
	err := config.DB.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func DeleteUser(id int) error {
	var user models.User

	result := config.DB.Where("id = ?", id).Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrInvalidID
	}

	return nil
}

func UpdateUser(id int, user models.User) (interface{}, error) {
	result := config.DB.Model(&models.User{}).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return user, result.Error
	}

	if result.RowsAffected == 0 {
		return user, ErrInvalidID
	}

	return user, nil
}
