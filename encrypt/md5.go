package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// GeneratePassword
// 进行密码加密
// 返回加密结果
func GeneratePassword(pwd, salt string) string {
	hash := md5.New()
	hash.Write([]byte(pwd + salt))
	return hex.EncodeToString(hash.Sum(nil))
}

func GenerateSign(timestamp string) string {
	hash := md5.New()
	hash.Write([]byte(fmt.Sprintf("%s-%s-%s", timestamp, "gtrive", timestamp)))
	return hex.EncodeToString(hash.Sum(nil))
}
