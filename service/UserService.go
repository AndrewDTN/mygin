package service

import(
	"GolangApi/pojo"
	"GolangApi/middlewares"
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

//CreatUserList
func CreatUserList(c *gin.Context){
	users:=pojo.Users{}
	err:=c.BindJSON(&users)
	if err!=nil{
		c.String(400,"Err:%s",err.Error())
		return
	}
	c.JSON(http.StatusOK,users)
}

//LoginUser
func LoginUser(c *gin.Context){
	name:=c.PostForm("name")
	password:=c.PostForm("password")
	user:=pojo.CheckUserPassword(name,password)
	if user.Id==0{
		c.JSON(http.StatusNotFound,"Error No Login")
		return
	}
	middlewares.SaveSession(c, user.Id)
	c.JSON(http.StatusOK,gin.H{
		"message":"Login Successfully",
		"User":user,
		"Sessions":middlewares.GetSession(c),
	})
}

//LogoutUser
func LogoutUser(c *gin.Context){
	middlewares.ClearSession(c)
	c.JSON(http.StatusOK,gin.H{
		"message":"Loggout Successfully",
	})
}

//CheckSession
func CheckUserSession(c *gin.Context){
	/*sessionId:=middlewares.GetSession(c)
	if sessionId==0{
		c.JSON(http.StatusUnauthorized,"Error No Session")
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"Check Session Successfully",
		"User":middlewares.GetSession(c),
	})*/
	if middlewares.CheckSession(c){
		c.JSON(http.StatusOK,gin.H{
			"message":"Check Session Successfully",
			"User":middlewares.GetSession(c),
		})
	}else{
		c.JSON(http.StatusUnauthorized,"Error No Session")
		return
	}
}