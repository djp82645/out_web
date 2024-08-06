package api

import "api_server/service"

type ApiGroup struct {
	InfoApi
}

var apiService = service.ServiceGroupApp.ApiServiceGroup.ApiService
