package wallet

import (
	"backend/api/internal/constants"
	"backend/api/internal/models"
	"backend/api/internal/svc"
	"backend/api/internal/types"
	"backend/api/internal/utils/encrypt"
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/tyler-smith/go-bip39"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateWalletLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateWalletLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateWalletLogic {
	return &CreateWalletLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateWalletLogic) CreateWallet(req *types.CreateWalletReq) (resp *types.Response, err error) {
	password := req.Password
	//生成助记词
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	//通过助记词生成种子
	seed := bip39.NewSeed(mnemonic, "")

	//生成私钥
	privateKey, _ := crypto.ToECDSA(seed[:32])
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKey).Hex()
	//使用密码加密助记词
	encryptedMnemonic, err := encrypt.EncryptMnemonic(mnemonic, req.Password)
	if err != nil {
		logx.Errorf("address:%s failed to encrypt mnemonic", address)
	}

	encryptedPass, err := encrypt.HashPassword(password)
	if err != nil {
		logx.Errorf("address:%s failed to encrypt password", address)
	}
	wallet := &models.WalletModel{
		Address:  address,
		Password: encryptedPass,
		Mnemonic: encryptedMnemonic,
	}

	//  存储信息
	if err = l.svcCtx.Gdb.WithContext(l.ctx).Create(wallet).Error; err != nil {
		logx.Errorf("Failed to create wallet: %v", err)
		return constants.CreateErr, nil
	}
	logx.Infof("create wallet success,walletInfo:[address:%s] [walletID:%d]", address, wallet.ID)
	data := &types.CreateWalletData{
		WalletID: int(wallet.ID),
		Address:  address,
		Mnemonic: mnemonic,
	}
	constants.CreateSuccess.Data = data
	return constants.CreateSuccess, nil
}
