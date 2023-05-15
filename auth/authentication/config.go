package authentication

type authCode struct {
	Code   string `json:"code"`
	Uuid   string `json:"uuid"`
	Verify string `json:"verify"`
}
