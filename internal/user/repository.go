package user

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"plant_identification/internal/database"
)

func saveUser(user User) error {
	if err := database.DB.Create(&user).Error; err != nil {
		return fmt.Errorf("could not save user: %v", err)
	}
	return nil
}

func getUser(username string) (User, error) {
	var user User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, fmt.Errorf("user not found")
		}
		return user, fmt.Errorf("could not retrieve user: %v", err)
	}
	return user, nil
}

func getUserByBatch(condition User) ([]User, error) {
	var users []User
	if err := database.DB.Where(&condition).Find(&users).Error; err != nil {
		return nil, fmt.Errorf("could not retrieve users: %v", err)
	}
	return users, nil
}
