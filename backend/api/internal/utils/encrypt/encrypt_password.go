package encrypt

import "golang.org/x/crypto/bcrypt"

// HashPassword 使用 bcrypt 生成密码哈希
func HashPassword(password string) (string, error) {
	// GenerateFromPassword 第二个参数是 cost，推荐使用 bcrypt.DefaultCost（默认值是10）
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// 返回哈希后的密码
	return string(hashedPassword), nil
}

// CheckPasswordHash 用于验证输入的密码和哈希是否匹配
func CheckPasswordHash(password, hashedPassword string) bool {
	// CompareHashAndPassword 会比较哈希密码和明文密码，并验证是否匹配
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
