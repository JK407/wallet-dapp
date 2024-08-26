package svc

import (
	"backend/api/internal/common"
	"backend/api/internal/config"
	"backend/api/internal/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Gdb    *gorm.DB
	Rdb    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	GDb, dbErr := common.InitGorm(c.Mysql.DSN)
	if dbErr != nil {
		panic("连接mysql数据库失败, error=" + dbErr.Error())
	}
	Rdb, rdbErr := common.InitRedis(c.Redis.Addr, c.Redis.Password, c.Redis.DB)
	if rdbErr != nil {
		panic("连接redis失败, error=" + rdbErr.Error())

	}
	models.Migrate(GDb)
	return &ServiceContext{
		Config: c,
		Gdb:    GDb,
		Rdb:    Rdb,
	}
}
