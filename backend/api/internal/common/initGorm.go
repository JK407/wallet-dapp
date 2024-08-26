package common

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func InitGorm(DSN string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(DSN),
		&gorm.Config{
			Logger: &dbLog{},
		})
	if err != nil {
		logx.Errorf("连接mysql数据库失败, error=" + err.Error())
		return nil, err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxOpenConns(100)              // SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // SetConnMaxLifetime 设置了连接可复用的最大时间。不能超过mysql的wait_timeout设置，否则会出现mysql has gone away错误
	return db, nil
}
