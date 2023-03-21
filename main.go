package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wkdwilliams/GolangTest/initializers"
	"github.com/wkdwilliams/GolangTest/router"
)

func init(){
	initializers.LoadVariables()
	initializers.ConnectToDB()
	initializers.ConnectToRedis()
}

func main(){
	gin.SetMode(gin.ReleaseMode)
	gin := gin.Default()

	router.RegisterRoutes(gin)

	// get, _ := initializers.Redis.Get(initializers.RedisContext, "test").Result()
	// fmt.Println(get)

	gin.Run()
}