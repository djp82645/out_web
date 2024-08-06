package user

import "api_server/service"

type ApiGroup struct {
	UserApi
}

var userService = service.ServiceGroupApp.UserServiceGroup.UserService
