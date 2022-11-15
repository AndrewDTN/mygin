package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"io"
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin/binding"
	
	"GolangApi/src"
	"GolangApi/database"
	"GolangApi/middlewares"
	"GolangApi/pojo"
)

func setupLogging(){
	f,_:=os.Create("gin.log")
	gin.DefaultWriter=io.MultiWriter(f,os.Stdout)
}

func main()  {

	setupLogging()

	router:=gin.Default()

	//註冊
	if v,ok:=binding.Validator.Engine().(*validator.Validate);ok{
		v.RegisterValidation("userpasd",middlewares.UserPasd)
		v.RegisterStructValidation(middlewares.UserList,pojo.Users{})
	}

	router.Use(gin.Recovery(),middlewares.Logger())

	v1:=router.Group("/v1")
	src.AddUserRouter(v1)
	
	go func(){
		database.DD()
	}()

	router.Run(":8000")
}