package middleware

import (
	"github.com/asuncm/vm/service/badger/userInfo"
	"github.com/asuncm/vm/service/config"
)

// 验证host域名权限
func origin(key string) string {
	if key != "" {
		return ""
	} else {
		return "baidu.com"
	}
}

// 生成临时用户签名
func Authorization(options config.ComConf, config userInfo.Authorization) (userInfo.UserInfo, error) {
	users, err := userInfo.Badger(options, config)
	return users, err
}
