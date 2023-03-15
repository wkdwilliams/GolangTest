package main

import (
	"github.com/wkdwilliams/GolangTest/initializers"
	"github.com/wkdwilliams/GolangTest/models"
)

func init(){
	initializers.LoadVariables()
	initializers.ConnectToDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.User{})
}