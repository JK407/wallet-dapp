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
	address := req.Address
	// 根据id查询数据库
	err = l.svcCtx.Gdb.WithContext(l.ctx).First(&wallet, walletID).Error
	if err != nil {
		logx.Errorf("get [walletId:%d] Info error:%v", walletID, err)
		return constants.LoginErr, nil
	}
	if wallet.Address != address {
		logx.Error("check address error")
		return constants.LoginErr, nil
	}
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
	logx.Infof("login wallet success,walletInfo: [walletID:%d] [address:%s]", walletInfo.ID, walletInfo.Address)
	return constants.LoginWalletSuccess, nil
}
