package api

import (
	v1 "api_server/api/v1"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct{}

func (e *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	// userRouter := Router.Group("user")
	exaInfoApi := v1.ApiGroupApp.ApiApiGroup.InfoApi
	{
		Router.GET("list", exaInfoApi.Search) // api接口查询
		Router.POST("add", exaInfoApi.Add) // api接口添加
	}
}
