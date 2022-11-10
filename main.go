package main

import (
	"github.com/gin-gonic/gin"
	
	"GolangApi/src"
	"GolangApi/database"
)

func main()  {
	router:=gin.Default()
	v1:=router.Group("/v1")
	src.AddUserRouter(v1)
	
	go func(){
		database.DD()
	}()

	router.Run(":8000")
}