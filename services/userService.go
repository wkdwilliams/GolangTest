package services

import (
	"github.com/wkdwilliams/GolangTest/initializers"
	"github.com/wkdwilliams/GolangTest/models"
)

// func GetUser(id int) models.User{

// }

func GetUsers() string{
	// users := make([]models.User, 0, 2)
	// users = append(users, models.User{})

	// return users
	return "hello"
}

func UpdateUser(id int, user models.User){

}

func CreateUser(user models.User) (models.User, error){

	result := initializers.DB.Create(&user)

	return user, result.Error
}