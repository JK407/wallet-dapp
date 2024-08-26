package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
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
