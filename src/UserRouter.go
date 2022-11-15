package src

import(
	"github.com/gin-gonic/gin"
	"GolangApi/service"
)

func AddUserRouter(r *gin.RouterGroup){
	user:=r.Group("/users")

	user.GET("/",service.FindAllUsers)
	user.GET("/:id",service.FindByUserId)
	user.POST("/",service.PostUser)
	user.POST("/more",service.CreatUserList)
	user.DELETE("/:id",service.DeleteUser)
	user.PUT("/:id",service.PutUser)
}