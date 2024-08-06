package core

import (
	"fmt"

	"api_server/global"
	"api_server/initialize"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	Router := initialize.Routers()
	address := fmt.Sprintf("0.0.0.0:%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	if global.GVA_DB != nil {
		initialize.RegisterTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
