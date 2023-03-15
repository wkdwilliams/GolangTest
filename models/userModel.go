package models

import "gorm.io/gorm"

type User struct{
	ID       uint `gorm:"primarykey"`
	Username string
	Password string
	Fname 	 string
	Lname 	 string
	gorm.Model
}