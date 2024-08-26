// Code generated by goctl. DO NOT EDIT.
package types

type CreateWalletReq struct {
	Password string `json:"password" validate:"required"`
}

type CreateWalletData struct {
	Address  string `json:"address"`
	Mnemonic string `json:"mnemonic"`
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}
