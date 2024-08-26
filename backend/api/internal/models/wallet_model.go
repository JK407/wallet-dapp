package models

import (
	"backend/api/internal/constants"
)

type WalletModel struct {
	ComModel
	Address  string `json:"address" gorm:"size:42;comment:钱包地址 (通常以太坊地址长度是42个字符)"` // Address @Description: 钱包地址 (通常以太坊地址长度是42个字符)
	Password string `json:"password" gorm:"size:100;comment:钱包密码 (已加密)"`           // Password @Description: 钱包密码 (这里假设最大长度为100字符)
	Mnemonic string `json:"mnemonic" gorm:"size:512;comment:钱包助记词 (已加密)"`          // Mnemonic @Description: 钱包助记词 (助记词通常较长，设置256字符足够)
}

// TableName 设置模型对应的表名
func (WalletModel) TableName() string {
	return constants.Wallet_Table
}
