package models

import "time"

type ComModel struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键，自增"`   // ID: 主键，自增
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;comment:创建时间，自动维护"` // CreatedAt: 创建时间，自动维护
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;comment:更新时间，自动维护"` // UpdatedAt: 更新时间，自动维护
}
