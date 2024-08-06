package service

import (
	"api_server/service/user"
	"api_server/service/api"

)

type ServiceGroup struct {
	UserServiceGroup user.ServiceGroup
	ApiServiceGroup api.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
