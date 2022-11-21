package src

import(
	"github.com/gin-gonic/gin"
	"GolangApi/service"
	"GolangApi/pojo"
	session"GolangApi/middlewares"
)

func AddUserRouter(r *gin.RouterGroup){
	user:=r.Group("/users",session.SetSession())

	user.GET("/:id",service.CachOneUserDecorator(service.RedisOneUser,"id","user_%s",pojo.User{}))
	user.GET("/",service.CachUserAllDecorator(service.RedisAllUser,"user_all",pojo.User{}))
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