package initialize

import (
	"net/http"
	"os"

	"api_server/global"
	"api_server/middleware"
	"api_server/router"
	"github.com/gin-gonic/gin"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrPermission
	}

	return f, nil
}

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	switch global.GVA_CONFIG.System.Env {
	case "local":
		gin.SetMode(gin.DebugMode)
		Router.Use(gin.Logger())
	case "public":
		gin.SetMode(gin.ReleaseMode)
		Router.Use(gin.Logger())
	default:
		Router.Use(gin.Logger())
	}
	userRouter := router.RouterGroupApp.User.UserRouter
	apiRouter := router.RouterGroupApp.Api.ApiRouter

	// Router.StaticFS(global.GVA_CONFIG.Local.StorePath, justFilesFilesystem{http.Dir(global.GVA_CONFIG.Local.StorePath)}) // Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	// 跨域，如需跨域可以打开下面的注释
	Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	// global.GVA_LOG.Info("use middleware cors")
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
		PublicGroup.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "欢迎使用api-server服务")
		})
	}
	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{
		userRouter.InitUserRouter(PrivateGroup)   // 用户路由
		apiRouter.InitApiRouter(PrivateGroup)     // 接口路由
	}

	global.GVA_LOG.Info("router register success")
	return Router
}
