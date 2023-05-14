package userInfo

// 用户信息数据结构
type UserInfo struct {
	Status bool        `json:"status"`
	Origin string      `json:"origin"`
	Users  interface{} `json:"users"`
}

type Authorization struct {
	Token  string `json:"token"`
	Verify string `json:"verify"`
	Origin string `json:"origin"`
}
