package models

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&WalletModel{},
	)
	if err != nil {
		return fmt.Errorf("auto migrate error:%s ", err)
	}
	logx.Info("auto migrate success")
	return nil
}
