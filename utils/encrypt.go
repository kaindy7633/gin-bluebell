package utils

import (
	"crypto/md5"
	"encoding/hex"
)

const secret = "kaindy7633"

// 使用 md5 加密字符串
func EncryptString(str string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(str)))
}
