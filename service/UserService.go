package service

import(
	"GolangApi/pojo"
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
)

//var userList = []pojo.User{}

//Get User
func FindAllUsers(c *gin.Context){
	//c.JSON(http.StatusOK,userList)
	users:=pojo.FindAllUsers()
	c.JSON(http.StatusOK,users)
}

//Get User By Id
func FindByUserId(c *gin.Context){
	user:=pojo.FindByUserId(c.Param("id"))
	if user.Id==0{
		c.JSON(http.StatusNotFound,"Error")
		return
	}
	log.Println("User ->",user)
	c.JSON(http.StatusOK,user)
}

//Post User
func PostUser(c *gin.Context){
	user:=pojo.User{}
	err:=c.BindJSON(&user)
	if err!=nil{
		c.JSON(http.StatusNotAcceptable,"Error: "+err.Error())
		return
	}
	newUser:=pojo.CreatUser(user)
	c.JSON(http.StatusOK, newUser)
}

//delete User
func DeleteUser(c *gin.Context){
	user:=pojo.DeleteUser(c.Param("id"))
	if !user{
		c.JSON(http.StatusNotFound,"Error")
		return
	}
	c.JSON(http.StatusOK,"Successfully")
}

//put User(修改)
func PutUser(c *gin.Context){
	user:=pojo.User{}
	err:=c.BindJSON(&user)
	if err!=nil{
		c.JSON(http.StatusBadRequest,"Error")
		return
	}
	user=pojo.UpdateUser(c.Param("id"),user)
	if user.Id==0{
		c.JSON(http.StatusNotFound,"Error")
		return
	}
	c.JSON(http.StatusOK,user)
}