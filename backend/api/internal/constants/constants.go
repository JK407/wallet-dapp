package constants

import "backend/api/internal/types"

// 数据库表名
const (
	Wallet_Table = "wallet"
)

// 成功状态码
const (
	CreateSuccessCode = 2000 + iota
	LoginWalletSuccessCode
)

// 失败状态码
const (
	CreateErrCode = 5000 + iota
	LoginErrCode

	ParameterErrCode = 4000
)

var (
	CreateErr          = types.NewResponse(CreateErrCode, "create wallet error", nil)
	CreateSuccess      = types.NewResponse(CreateSuccessCode, "create wallet success", nil)
	ParameterErr       = types.NewResponse(ParameterErrCode, "", nil)
	LoginErr           = types.NewResponse(LoginErrCode, "login wallet error", nil)
	LoginWalletSuccess = types.NewResponse(LoginWalletSuccessCode, "login wallet success", nil)
)
