package constants

import "backend/api/internal/types"

// 数据库表名
const (
	Wallet_Table = "wallet"
)

// 成功状态码
const (
	CreateSuccessCode = 2000 + iota
)

// 失败状态码
const (
	ParameterErrCode = 4000
	CreateErrCode    = 5000 + iota
)

var (
	CreateErr     = types.NewResponse(CreateErrCode, "create wallet error", nil)
	CreateSuccess = types.NewResponse(CreateSuccessCode, "create wallet success", nil)
	ParameterErr  = types.NewResponse(ParameterErrCode, "", nil)
)
