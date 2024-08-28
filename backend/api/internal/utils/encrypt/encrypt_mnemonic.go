package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

// 加密助记词的函数
func EncryptMnemonic(mnemonic, password string) (string, error) {
	block, err := aes.NewCipher([]byte(password)[:16]) // 使用前16字节作为AES密钥
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(mnemonic), nil)
	return hex.EncodeToString(ciphertext), nil
}

// DecryptMnemonic
// @Description 解密助记词
// @Author Oberl-Fitzgerald 2024-08-28 16:51:29
// @Param  encryptedMnemonic string
// @Param  password string
// @Return bool
func DecryptMnemonic(encryptedMnemonic, password string) bool {
	data, err := hex.DecodeString(encryptedMnemonic)
	if err != nil {
		logx.Error(err)
		return false
	}

	key := []byte(password) // 同样，确保使用相同逻辑生成密钥
	block, err := aes.NewCipher(key[:16])
	if err != nil {
		logx.Error(err)
		return false
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		logx.Error(err)
		return false
	}

	if len(data) < gcm.NonceSize() {
		logx.Error(errors.New("ciphertext too short"))
		return false
	}

	nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]
	_, err = gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		logx.Error(err)
		return false
	}

	return true
}
