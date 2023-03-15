package services

import "github.com/wkdwilliams/GolangTest/models"

// func GetUser(id int) models.User{

// }

func GetUsers() []models.User{
	users := make([]models.User, 0, 2)
	users = append(users, models.User{})

	return users
}

func UpdateUser(id int, user models.User){

}

func CreateUser(user models.User) models.User{
	return user
}