package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"io"
	
	"GolangApi/src"
	"GolangApi/database"
	"GolangApi/middlewares"
)

func setupLogging(){
	f,_:=os.Create("gin.log")
	gin.DefaultWriter=io.MultiWriter(f,os.Stdout)
}

func main()  {

	setupLogging()

	router:=gin.Default()

	router.Use(gin.BasicAuth(gin.Accounts{"Andrew":"123456"}),middlewares.Logger())

	v1:=router.Group("/v1")
	src.AddUserRouter(v1)
	
	go func(){
		database.DD()
	}()

	router.Run(":8000")
}