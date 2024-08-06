package router

import (
	"api_server/router/user"
	"api_server/router/api"
)

type RouterGroup struct {
	User user.RouterGroup	
	Api api.RouterGroup	

}

var RouterGroupApp = new(RouterGroup)
