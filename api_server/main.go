package main

import (
	"go.uber.org/zap"

	"api_server/core"
	"api_server/global"
)

func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	global.GVA_LOG = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	core.RunWindowsServer()
}
