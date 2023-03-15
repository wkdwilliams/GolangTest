package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB(){
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/main?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatal("Cannot connect to database")
	}

	DB = d
}