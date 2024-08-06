package initialize

import (
	"fmt"
	// "log"
	"database/sql"
	"time"

	"api_server/global"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GormMysql 初始化Mysql数据库
func GormMysql() *gorm.DB {
	dbHost := global.GVA_CONFIG.Mysql.Host
	dbPort := global.GVA_CONFIG.Mysql.Port
	dbName := global.GVA_CONFIG.Mysql.DbName
	dbParams := global.GVA_CONFIG.Mysql.Config
	dbUser := global.GVA_CONFIG.Mysql.Username
	dbPasswd := global.GVA_CONFIG.Mysql.Password
	dbURL := fmt.Sprintf("%s:%s@(%s:%d)/%s?%s", dbUser, dbPasswd, dbHost, dbPort, dbName, dbParams)
	// ----------------------- 连接数据库 -----------------------
	var err error
	DB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dbURL, // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: true,  // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		// dev Logger: newLogger,
		Logger: logger.Default.LogMode(logger.LogLevel(global.GVA_CONFIG.Mysql.GormLog)),
		// pro
		// Logger: logger.Default.LogMode(logger.Info),
		PrepareStmt: true, // 全局模式，所有的操作都会创建并缓存预编译语句，以加速后续执行速度
		// SkipDefaultTransaction: true, //禁用事务
		CreateBatchSize: 1000, // 分页批量插入设置
	})
	if err != nil {
		global.GVA_LOG.Error("Open mysql failed,err:"+err.Error())
		// panic(err)
	}
	sqlDB, err := DB.DB()
	sqlDB.SetConnMaxLifetime(600 * time.Second)
	sqlDB.SetMaxIdleConns(global.GVA_CONFIG.Mysql.MaxIdleConns) // 最大打开的连接数
	sqlDB.SetMaxOpenConns(global.GVA_CONFIG.Mysql.MaxOpenConns) // 设置最大闲置个数
	return DB
}

// 创建数据库
func AutoCreateDB() {
	dbType := global.GVA_CONFIG.System.DbType
	dbHost := global.GVA_CONFIG.Mysql.Host
	dbPort := global.GVA_CONFIG.Mysql.Port
	dbName := global.GVA_CONFIG.Mysql.DbName
	dbUser := global.GVA_CONFIG.Mysql.Username
	dbPasswd := global.GVA_CONFIG.Mysql.Password
	// sql := "CREATE DATABASE IF NOT EXISTS ? DEFAULT CHARACTER SET = 'utf8';"
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/", dbUser, dbPasswd, dbHost, dbPort)
	db, err := sql.Open(dbType, dbURL)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		global.GVA_LOG.Error(err.Error())
	}
	global.GVA_LOG.Info("数据库连接成功")

	createTable := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;", dbName)
	_, err = db.Exec(createTable)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
	}
	global.GVA_LOG.Info("数据库创建成功")
}
