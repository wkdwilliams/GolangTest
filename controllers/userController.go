package controllers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/wkdwilliams/GolangTest/models"
	"github.com/wkdwilliams/GolangTest/services"
)

func UserIndex(c *gin.Context){
	users, err := services.GetUsers()

	if err != nil{
		c.JSON(500, gin.H{
			"message": "Error when geting users",
		})
	}

	c.JSON(200, gin.H{
		"message":  "get_success",
		"data": 	users,
	})
}

func UserShow(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(422, gin.H{
			"message": "Invalid id provided",
		})
		return
	}

	user, err := services.GetUser(id)

	if err != nil{
		c.JSON(500, gin.H{
			"message": "Error when geting user",
		})
	}

	c.JSON(200, gin.H{
		"message":  "get_success",
		"data": 	user,
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

func UserUpdate(c *gin.Context){
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

	user := services.UpdateUser(id, models.User{
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

func UserDelete(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(422, gin.H{
			"message": "Invalid id provided",
		})
		return
	}

	user := services.DeleteUser(id)

	c.JSON(200, gin.H{
		"message": "deleted",
		"data": user,
	})
}