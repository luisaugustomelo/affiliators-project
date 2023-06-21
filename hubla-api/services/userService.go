package services

import (
	"errors"
	"fmt"

	"github.com/luisaugustomelo/hubla-challenge/database/models"
	"github.com/luisaugustomelo/hubla-challenge/interfaces"
	"gorm.io/gorm"
)

func CreateUser(db interfaces.Datastore, user *models.User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByID(db interfaces.Datastore, id uint) (*models.User, error) {
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(db interfaces.Datastore, email string) (*models.User, error) {
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user")
	}
	return &user, nil
}

func UpdateUser(db interfaces.Datastore, id uint, updatedUser *models.User) error {
	var user models.User

	// Find currently user
	if err := db.First(&user, id).Error; err != nil {
		return err
	}

	// Check if email already exists
	var existingUser models.User
	if err := db.Where("email = ? AND id != ?", updatedUser.Email, id).First(&existingUser).Error; err == nil {
		return fmt.Errorf("email already exists")
	}

	// update user
	if updatedUser.Name != "" {
		user.Name = updatedUser.Name
	}
	if updatedUser.Email != "" {
		user.Email = updatedUser.Email
	}
	if updatedUser.Password != "" {
		user.Password = updatedUser.Password
	}

	if err := db.Save(&user).Error; err != nil {
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
