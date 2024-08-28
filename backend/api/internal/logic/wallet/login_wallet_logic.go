package wallet

import (
	"backend/api/internal/constants"
	"backend/api/internal/models"
	"backend/api/internal/utils/encrypt"
	"context"

	"backend/api/internal/svc"
	"backend/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginWalletLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginWalletLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginWalletLogic {
	return &LoginWalletLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginWalletLogic) LoginWallet(req *types.LoginWalletReq) (resp *types.Response, err error) {
	var wallet models.WalletModel
	password := req.Password
	walletID := req.WalletID
	// 根据id查询数据库
	l.svcCtx.Gdb.WithContext(l.ctx).First(&wallet, walletID)
	//  校验密码
	if !encrypt.CheckPasswordHash(password, wallet.Password) {
		logx.Error("check password error")
		return constants.LoginErr, nil
	}
	// 解密助记词
	if !encrypt.DecryptMnemonic(wallet.Mnemonic, password) {
		logx.Error("decrypt mnemonic error")
		return constants.LoginErr, nil
	}
	walletInfo := types.WalletInfoData{
		ID:                int(wallet.ID),
		Address:           wallet.Address,
		EncryptedMnemonic: wallet.Mnemonic,
	}
	constants.LoginWalletSuccess.Data = walletInfo
	return constants.LoginWalletSuccess, nil
}
