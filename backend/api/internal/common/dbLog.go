package common

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm/logger"
	"time"
)

// todo caller
type dbLog struct {
	LogLevel logger.LogLevel
}

func (l *dbLog) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *dbLog) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < logger.Info {
		return
	}
	logx.WithContext(ctx).Debugf(msg, data)
}
func (l *dbLog) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < logger.Warn {
		return
	}
	logx.WithContext(ctx).Infof(msg, data)
}

func (l *dbLog) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < logger.Error {
		return
	}
	logx.WithContext(ctx).Errorf(msg, data)
}

func (l *dbLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	//这块的逻辑可以自己根据业务情况修改
	//fmt.Println(l.LogLevel)
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		logx.WithContext(ctx).WithDuration(elapsed).Errorf("Trace sql: %v  row： %v  err: %v", sql, rows, err)
	} else {
		logx.WithContext(ctx).WithDuration(elapsed).Infof("Trace sql: %v  row： %v", sql, rows)
	}
}
