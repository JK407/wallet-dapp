package common

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
)

// InitRedis 初始化redis连接
func InitRedis(add, password string, db int) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     add,
		Password: password,
		DB:       db,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logx.Error("连接redis失败, error=" + err.Error())
		return nil, err
	}
	return rdb, nil
}
