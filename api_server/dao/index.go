package dao

import (
	"api_server/dao/api"

)

type DaoGroup struct {
	ApiDaoGroup     api.DaoGroup
}

var DaoGroupApp = new(DaoGroup)
