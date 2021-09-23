package database

import (
	"crud-test/config"
	"crud-test/models"
)

func GetUserByID(id string) (*models.User, error) {
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return &models.User{}, err
	}
	return &user, nil
}

func GetUserAll() (*[]models.User, error){
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil{
		return &[]models.User{}, err
	}
	return &users, nil
}

func PostUser(user models.User) (*models.User, error){
	if err := config.DB.Save(&user).Error; err != nil{
		return &models.User{}, err
	}
	return &user, nil
}

func UpdateUser(id string, userData models.User) (*models.User, error){
	var user models.User

	if err := config.DB.First(&user, id).Updates(&userData).Error; err != nil{
		return &models.User{}, err
	}
	return &user, nil
}

func DeleteUser(id string, user models.User) (string, error){
	if err := config.DB.Unscoped().Delete(&user).Error; err != nil{
		return "", err
	}
	return "deleted", nil
}