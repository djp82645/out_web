package api

import (

	"api_server/global"
	"api_server/model/common/pojo"
	"strings"
)

var api pojo.ApiInfo

type ApiDao struct{}
func (ApiDao *ApiDao) Insert(data []string) error{
	var api_info []pojo.ApiInfo
	for _, v := range data {
		info := strings.Split(v, "----")
		msg := pojo.ApiInfo{
			Sn:     info[0],
			Api:info[1],
		}
		api_info = append(api_info, msg)
	}
	result := global.GVA_DB.Create(&api_info) // 通过数据的指针来创建
	if result.Error != nil {
		return  result.Error
	}
	return nil
}
// 列表
func (ApiDao *ApiDao) Pagenate(
	sn string,	page_num, page_size int,
) ([]pojo.ApiInfo, int, int, error) {
	/*
		返回 列表，行数，分页 错误
	*/
	var total int64
	var api_list []pojo.ApiInfo
	if sn != "" {
		global.GVA_DB.Model(&api).Where("sn =?", sn).Count(&total)
	}else{
		global.GVA_DB.Model(&api).Count(&total)
	}
	if total == 0 {
		return api_list, 0, 0,  nil
	}
	pageNum := int(total) / page_size
	if int(total)%page_size != 0 {
		pageNum++
	}
	tmpDB := global.GVA_DB.Model(&api)
	if sn != "" {
		result := tmpDB.Where("sn =?", sn).Limit(page_size).Offset((page_num - 1) * page_size).Find(&api_list)
		err := result.Error // returns error
		if err != nil {
			return nil, 0, 0, err
		}
	}else{
		result := tmpDB.Limit(page_size).Offset((page_num - 1) * page_size).Find(&api_list)
		err := result.Error // returns error
		if err != nil {
			return nil, 0, 0, err
		}
	}
	return api_list, int(total), pageNum,  nil
}

