package user

import (
	"api_server/global"
	"api_server/model/common/request"
	"api_server/model/common/response"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

func (e *UserApi) Login(c *gin.Context) {
	var user request.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		global.GVA_LOG.Error("Login error:" + err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	token, err := userService.Login(user.Username, user.Password)
	if err != nil {
		global.GVA_LOG.Error("Login error:" + err.Error())
		response.Result(50008,nil, "登录失败", c)
		return
	}
	response.Result(20000, token, "登录成功", c)
}

func (e *UserApi) Search(c *gin.Context) {
	token := c.Query("token")
	info, err := userService.SearchUser(token)
	if err != nil {
		global.GVA_LOG.Error("Search error:" + err.Error())
		response.FailWithDetailed(nil, "用户查询失败", c)
		return
	}
	response.Result(20000, info, "用户查询成功", c)
}
func (e *UserApi) Logout(c *gin.Context) {
	err := userService.Logout()
	if err != nil {
		global.GVA_LOG.Error("Logout error:" + err.Error())
		response.FailWithDetailed(nil, "登出失败", c)
		return
	}
	response.Result(20000, "success", "登出成功", c)
}

