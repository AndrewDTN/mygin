package service

import(
	"GolangApi/pojo"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
	"log"
)

var userList = []pojo.User{}

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
	userList = append(userList,user)
	c.JSON(http.StatusOK,"Successfully posted")
	return
}

//delete User
func DeleteUser(c *gin.Context){
	userId,_:=strconv.Atoi(c.Param("id"))
	for _,user:=range userList{
		log.Println(user)
		userList = append(userList[:userId],userList[userId+1:]...)
		c.JSON(http.StatusOK,"Successfully deleted")
		return
	}
	c.JSON(http.StatusNotFound,"Error")
}

//put User(修改)
func PutUser(c *gin.Context){
	beforeUser:=pojo.User{}
	err:=c.BindJSON(&beforeUser)
	if err!=nil{
		c.JSON(http.StatusBadRequest,"Error")
	}
	userId,_:=strconv.Atoi(c.Param("id"))
	for key,user:=range userList{
		if userId==user.Id{
			userList[key]=beforeUser
			log.Println(userList[key])
			c.JSON(http.StatusOK,"Successfully put")
			return
		}
	}
	c.JSON(http.StatusNotFound,"Error")
}