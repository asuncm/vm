package middleware

import (
	"github.com/asuncm/vm/service/badger"
	"github.com/asuncm/vm/service/badger/userInfo"
	"github.com/vmihailenco/msgpack"
)

/*@package Authorization 读取用户信息，判断登录状态
* @param list			 返回相关用户信息
* @param Origin			 是否限制域名
* @param Status			 用户登录状态，0/false未登录，1/true已登录
* @param Users			 用户信息
 */
func Authorization(config userInfo.Authorization) (userInfo.UserInfo, error) {
	var (
		conf  configMap
		users []byte
		err   error
	)
	if config.Uid != "" {
		users, err = badger.Query([]byte(config.Uid), "/basic")
	}
	list := userInfo.UserInfo{
		Origin: "*",
		Status: true,
	}
	if err != nil {
		list.Status = false
		list.Message = "未获取到相关数据"
	} else {
		err = msgpack.Unmarshal(users, &conf)
		if err != nil {
			list.Status = false
			list.Message = "相关数据为空"
		} else {
			list.Users = conf
			value := conf["Origin"].(string)
			if value != "" {
				list.Origin = value
			}
		}
	}
	return list, err
}
