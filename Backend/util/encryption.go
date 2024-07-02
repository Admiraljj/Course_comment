package util

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashPassword 使用 SHA-256 算法将密码转换为加密字符串
func Encryption(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	hashedPassword := hash.Sum(nil)
	return hex.EncodeToString(hashedPassword)
}
