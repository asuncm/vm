package middleware

import (
	"github.com/asuncm/vm/service/badger"
	"github.com/asuncm/vm/service/badger/userInfo"
	"github.com/vmihailenco/msgpack"
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
func Authorization(config userInfo.Authorization) (userInfo.UserInfo, error) {
	var conf map[string]interface{}
	users, err := badger.Query([]byte(config.Verify), "/basic")
	list := userInfo.UserInfo{
		Origin: "*",
		Status: true,
	}
	if err != nil {
		list.Status = false
		return list, err
	}
	err = msgpack.Unmarshal(users, &conf)
	list.Users = conf
	value := conf["Origin"].(string)
	if value != "" {
		list.Origin = value
	}
	return list, err
}
