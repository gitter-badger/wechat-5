package models

type UserInfo struct {
	openid     string   `json:"openid"`
	nickname   string   `json:"nickname"`
	sex        int      `json:"sex"`
	province   string   `json:"province"`
	city       string   `json:"city"`
	country    string   `json:"country"`
	headimgurl string   `json:"headimgurl"`
	privilege  []string `json:"privilege"`
	unionid    string   `json:"unionid"`
}
