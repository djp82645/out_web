package api

import (
	// "fmt"
	"api_server/dao"

	// "api_server/model/common/request"
	// "api_server/utils"
)
var apiDao = dao.DaoGroupApp.ApiDaoGroup.ApiDao
type ApiService struct{}

func (exa *ApiService) SearchInfo(sn string,page ,size int) (map[string]interface{}, error) {
	info,total ,pageNum,err := apiDao.Pagenate(sn,page,size)
	if err != nil {
		return nil, err
	} 
	data := make(map[string]interface{},3)
	data["data"] = info
	data["totalCount"] = total
	data["totalPage"] = pageNum
	return data, nil
}

func (exa *ApiService) AddInfo(data []string) error {
	err := apiDao.Insert(data)
	if err != nil {
		return err
	} 
	return nil
}
