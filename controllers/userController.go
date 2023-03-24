package controllers

import (
	"errors"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/wkdwilliams/GolangTest/models"
	"github.com/wkdwilliams/GolangTest/services"
	"gorm.io/gorm"
)

func NewUserController() *UserController{
	service := services.UserService{}

	return &UserController{
		service,
	}
}

type UserController struct{
	UserService services.UserService
}

func (u *UserController) Index(c *gin.Context){
	users, err := u.UserService.GetUsers()

	if err != nil{
		c.JSON(500, gin.H{
			"message": "Error when geting users",
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "get_success",
		"data": 	users,
	})
}

func (u *UserController) Show(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(422, gin.H{
			"message": "Invalid id provided",
		})
		return
	}

	user, err := u.UserService.GetUser(id)

	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			c.JSON(404, gin.H{"message": "Record not found"})
			return
		}
		c.JSON(500, gin.H{"message": "Error when geting user"})
		return
	}

	c.JSON(200, gin.H{
		"message":  "get_success",
		"data": 	user,
	})
}

func (u *UserController) Store(c *gin.Context){
	var body struct{
		Username string
		Password string
		Fname    string
		Lname    string
	}

	c.Bind(&body)

	user, err := u.UserService.CreateUser(models.User{
		Username: body.Username,
		Password: body.Password,
		Fname:    body.Fname,
		Lname:    body.Lname,
	})

	if err != nil{
		c.Status(400)
		return
	}

	c.JSON(201, gin.H{
		"message": "created",
		"data":    user,
	})
}

func (u *UserController) Update(c *gin.Context){
	var body struct{
		Username string
		Password string
		Fname    string
		Lname    string
	}

	c.Bind(&body)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(422, gin.H{
			"message": "Invalid id provided",
		})
		return
	}

	user := u.UserService.UpdateUser(id, models.User{
		Username: body.Username,
		Password: body.Password,
		Fname:    body.Fname,
		Lname:    body.Lname,
	})

	c.JSON(200, gin.H{
		"message": "updated",
		"data":    user,
	})
}

func (u *UserController) Delete(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(422, gin.H{
			"message": "Invalid id provided",
		})
		return
	}

	user := u.UserService.DeleteUser(id)

	c.JSON(200, gin.H{
		"message": "deleted",
		"data": user,
	})
}