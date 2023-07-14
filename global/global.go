package global

import (
	"demo-gin-1/config"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/timer"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"sync"
)

var (
	GVA_DB                  *gorm.DB
	GVA_DBList              map[string]*gorm.DB
	GVA_CONFIG              config.Server
	GVA_VP                  *viper.Viper
	GVA_LOG                 *zap.Logger
	GVA_Timer               timer.Timer = timer.NewTimerTask()
	GVA_Concurrency_Control             = &singleflight.Group{}

	BlackCache local_cache.Cache
	lock       sync.RWMutex
)

// 通过名称获取dblist中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.Lock()
	defer lock.Unlock()
	return GVA_DBList[dbname]
}

func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.Lock()
	defer lock.Unlock()
	db, ok := GVA_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
