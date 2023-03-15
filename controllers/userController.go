package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wkdwilliams/GolangTest/models"
	"github.com/wkdwilliams/GolangTest/services"
)

func UserIndex(c *gin.Context){
	c.JSON(200, gin.H{
		"message":  "get_success",
		"data": 	"",
	})
}

func UserStore(c *gin.Context){
	var body struct{
		Username string
		Password string
		Fname    string
		Lname    string
	}

	c.Bind(&body)

	user, err := services.CreateUser(models.User{
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