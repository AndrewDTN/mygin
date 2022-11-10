package pojo

import (
	"GolangApi/database"
)

type User struct{
	Id int `json:"Userid"`
	Name string `json:"UserName"`
	Password string `json:"UserPassword"`
	Email string `json:"UserEmail"`
}

func FindAllUsers() []User{
	var users []User
	database.DBconnecct.Find(&users)
	return users
}

func FindByUserId(userId string)User{
	var user User
	database.DBconnecct.Where("id = ?",userId).First(&user)
	return user
}
//test