package database

import (
	"github.com/suryaadi44/go-training-restful/config"
	"github.com/suryaadi44/go-training-restful/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	err := config.DB.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}
