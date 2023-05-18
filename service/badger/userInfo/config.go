package userInfo

// 用户信息数据结构
type UserInfo struct {
	Status  bool                        `json:"status"`
	Origin  string                      `json:"origin"`
	Message string                      `json:"message"`
	Users   map[interface{}]interface{} `json:"users"`
}

type Authorization struct {
	Token  string `json:"token"`
	Uid    string `json:"uid"`
	Origin string `json:"origin"`
}
