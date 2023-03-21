package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wkdwilliams/GolangTest/controllers"
	"github.com/wkdwilliams/GolangTest/services"
)

func RegisterRoutes(router *gin.Engine){
	userController := controllers.UserController{
		UserService: services.UserService{},
	}

	router.POST("/user", userController.Store)
	router.PUT("/user/:id", userController.Update)
	router.DELETE("/user/:id", userController.Delete)
	router.GET("/user", userController.Index)
	router.GET("/user/:id", userController.Show)
}