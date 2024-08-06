package api

import (
	"api_server/global"
	"api_server/model/common/request"
	"api_server/model/common/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InfoApi struct{}


func (e *InfoApi) Search(c *gin.Context) {
	page := c.Query("page")
	size := c.Query("size")
	sn := c.Query("sn")
	pageInt, _ := strconv.Atoi(page)
	sizeInt, _ := strconv.Atoi(size)
	info, err := apiService.SearchInfo(sn,pageInt,sizeInt)
	if err != nil {
		global.GVA_LOG.Error("SearchInfo error:" + err.Error())
		response.FailWithDetailed(nil, "查询api接口失败", c)
		return
	}
	response.Result(20000, info, "查询api接口成功", c)
}

func (e *InfoApi) Add(c *gin.Context) {
	var user_data request.DataInfo
	err := c.ShouldBindJSON(&user_data)
	if err != nil {
		global.GVA_LOG.Error("Add error:" + err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = apiService.AddInfo(user_data.Data)
	if err != nil {
		global.GVA_LOG.Error("AddInfo error:" + err.Error())
		response.FailWithDetailed(nil, "添加接口失败", c)
		return
	}
	response.Result(20000, "sucess", "添加接口成功", c)
}


