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
	r.GET("/user", controllers.UserIndex)

	r.Run()
}