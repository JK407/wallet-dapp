package wallet

import (
	"backend/api/internal/constants"
	"backend/api/internal/utils/validator"
	"encoding/json"
	"net/http"

	"backend/api/internal/logic/wallet"
	"backend/api/internal/svc"
	"backend/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateWalletHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateWalletReq
		// 使用decode解析
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields() // 禁止未知字段，提高类型安全性
		if err := decoder.Decode(&req); err != nil {
			constants.ParameterErr.Msg = "param decode error"
			json.NewEncoder(w).Encode(constants.ParameterErr) // 直接使用预定义的错误响应
			return
		}
		// 使用 go-playground/validator 验证请求数据
		if err := validator.ValidateStruct(&req); err != nil {
			constants.ParameterErr.Msg = err.Error()
			json.NewEncoder(w).Encode(constants.ParameterErr) // 直接使用预定义的错误响应
			return
		}

		l := wallet.NewCreateWalletLogic(r.Context(), svcCtx)
		resp, err := l.CreateWallet(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
