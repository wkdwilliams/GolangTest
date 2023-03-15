package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wkdwilliams/GolangTest/models"
	"github.com/wkdwilliams/GolangTest/services"
)

func UserStore(c *gin.Context){
	user := services.CreateUser(models.User{
		Username: "wkdwilliams",
		Password: "pass",
		Fname:    "lewis",
		Lname:    "williams",
	})
	c.JSON(201, gin.H{
		"message": "created",
		"data":    user,
	})
}