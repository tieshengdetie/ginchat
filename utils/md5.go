package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Md5Encode 小写
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// MD5Encode 大写
func MD5Encode(data string) string {

	return strings.ToUpper(Md5Encode(data))
}

// MakePassword 盐值加密密码
func MakePassword(plainpwd, salt string) string {
	return MD5Encode(plainpwd + salt)
}

// 校验密码
func ValidPassword(plainpwd, salt, password string) bool {

	return MD5Encode(plainpwd+salt) == password
}
