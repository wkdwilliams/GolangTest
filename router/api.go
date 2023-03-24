package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wkdwilliams/GolangTest/controllers"
)

func RegisterRoutes(router *gin.Engine){
	userController := controllers.NewUserController()

	router.POST("/user", userController.Store)
	router.PUT("/user/:id", userController.Update)
	router.DELETE("/user/:id", userController.Delete)
	router.GET("/user", userController.Index)
	router.GET("/user/:id", userController.Show)
}