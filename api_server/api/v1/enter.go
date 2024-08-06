package v1

import (
	"api_server/api/v1/user"
	"api_server/api/v1/api"

)

type ApiGroup struct {
	UserApiGroup user.ApiGroup
	ApiApiGroup api.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
