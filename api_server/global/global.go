package global

import (
	// "github.com/qiniu/qmgo"
	"sync"

	"github.com/songzhibin97/gkit/cache/local_cache"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"api_server/config"

	"github.com/spf13/viper"
	
	"gorm.io/gorm"
)

var (
	GVA_DB       *gorm.DB
	GVA_CONFIG              config.Server
	GVA_VP                  *viper.Viper
	GVA_LOG                 *zap.Logger
	GVA_Concurrency_Control = &singleflight.Group{}

	BlackCache local_cache.Cache
	lock       sync.RWMutex
)
