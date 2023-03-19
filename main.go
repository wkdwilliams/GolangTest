package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wkdwilliams/GolangTest/controllers"
	"github.com/wkdwilliams/GolangTest/initializers"
)

func init(){
	initializers.LoadVariables()
	initializers.ConnectToDB()
}

func main(){
	r := gin.Default()

	r.POST("/user", controllers.UserStore)
	r.PUT("/user/:id", controllers.UserUpdate)
	r.DELETE("/user/:id", controllers.UserDelete)
	r.GET("/user", controllers.UserIndex)
	r.GET("/user/:id", controllers.UserShow)

	r.Run()
}