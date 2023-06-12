package services

import (
	"github.com/luisaugustomelo/hubla-challenge/database/models"
	"github.com/luisaugustomelo/hubla-challenge/interfaces"
)

func CreateUser(db interfaces.Datastore, user *models.User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUser(db interfaces.Datastore, id uint) (*models.User, error) {
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(db interfaces.Datastore, user *models.User) error {
	if err := db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(db interfaces.Datastore, id uint) error {
	if err := db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
