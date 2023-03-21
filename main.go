package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wkdwilliams/GolangTest/initializers"
	"github.com/wkdwilliams/GolangTest/router"
)

func init(){
	initializers.LoadVariables()
	initializers.ConnectToDB()
}

func main(){
	gin.SetMode(gin.ReleaseMode)
	gin := gin.Default()

	router.RegisterRoutes(gin)

	gin.Run()
}