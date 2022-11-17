package src

import(
	"github.com/gin-gonic/gin"
	"GolangApi/service"
	session"GolangApi/middlewares"
)

func AddUserRouter(r *gin.RouterGroup){
	user:=r.Group("/users",session.SetSession())

	user.GET("/",service.FindAllUsers)
	user.GET("/:id",service.FindByUserId)
	user.POST("/",service.PostUser)
	user.POST("/more",service.CreatUserList)
	
	user.PUT("/:id",service.PutUser)

	//login
	user.POST("/login",service.LoginUser)

	//checkusersession
	user.GET("/check",service.CheckUserSession)

	user.Use(session.AuthSession())
	{
		user.DELETE("/:id",service.DeleteUser)
		user.GET("/logout",service.LogoutUser)
	}
}