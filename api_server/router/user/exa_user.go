package user

import (
	v1 "api_server/api/v1"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (e *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	exaUserApi := v1.ApiGroupApp.UserApiGroup.UserApi
	{
		userRouter.POST("login", exaUserApi.Login) // 用户登录
		userRouter.GET("info", exaUserApi.Search) // 用户查询
		userRouter.GET("logout", exaUserApi.Logout) // 用户登出
	}
}
