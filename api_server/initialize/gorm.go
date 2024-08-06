package initialize

import (
	"errors"
	"os"

	// "log"

	"api_server/global"
	"api_server/model/common/pojo"

	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		AutoCreateDB()
		return GormMysql()
	default:
		panic(errors.New("only support mysql"))
	}
}

// RegisterTables 注册数据库表专用
// Author SliverHorn
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		pojo.ApiInfo{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed:"+err.Error())
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
