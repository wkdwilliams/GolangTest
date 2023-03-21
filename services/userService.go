package services

import (
	"github.com/wkdwilliams/GolangTest/initializers"
	"github.com/wkdwilliams/GolangTest/models"
)

type UserService struct{

}

func (u *UserService) GetUsers() ([]models.User, error){
	var users []models.User
	result := initializers.DB.Find(&users)

	return users, result.Error
}

func (u *UserService) GetUser(id int) (*models.User, error){
	var user models.User

	if err := initializers.DB.Where("id = ?", id).First(&user).Error; err != nil{
		return nil, err
	}

	return &user, nil
}

func (u *UserService) DeleteUser(id int) (models.User){
	var user models.User
	initializers.DB.First(&user, id)

	initializers.DB.Delete(&user)

	return user
}

func (u *UserService) UpdateUser(id int, user models.User) (models.User){
	var userBeforeUpdate models.User
	initializers.DB.First(&userBeforeUpdate, id)

	initializers.DB.Model(&userBeforeUpdate).Updates(user)

	return user
}

func (u *UserService) CreateUser(user models.User) (models.User, error){
	result := initializers.DB.Create(&user)

	return user, result.Error
}