package services

import (
	"github.com/wkdwilliams/GolangTest/initializers"
	"github.com/wkdwilliams/GolangTest/models"
)

func GetUsers() ([]models.User, error){
	var users []models.User
	result := initializers.DB.Find(&users)

	return users, result.Error
}

func GetUser(id int) (models.User, error){
	var user models.User
	result := initializers.DB.First(&user, id)

	return user, result.Error
}

func DeleteUser(id int) (models.User){
	var user models.User
	initializers.DB.First(&user, id)

	initializers.DB.Delete(&models.User{}, id)

	return user
}

func UpdateUser(id int, user models.User) (models.User){
	var userBeforeUpdate models.User
	initializers.DB.First(&userBeforeUpdate, id)

	initializers.DB.Model(&userBeforeUpdate).Updates(user)

	return user
}

func CreateUser(user models.User) (models.User, error){
	result := initializers.DB.Create(&user)

	return user, result.Error
}