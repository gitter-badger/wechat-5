package wechat

import (
	"encoding/json"
	"github.com/alecthomas/log4go"
	"io/ioutil"
	"net/http"
	"nicosoft.org/wechat/utils"
	"strings"
	"time"
)

type UserToken struct {
	Token     string `json:"access_token"`
	ExpiresIn int    `json:"expires_in"`
	NewToken  string `json:"refresh_token"`
	OpenId    string `json:"openid"`
	Scope     string `json:"scope"`
}

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

func (u *UserInfo) GetToken(code string) *UserToken {

	url := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=APPID&secret=SECRET&code=CODE&grant_type=authorization_code"
	cache := new(utils.Cache)

	ut := new(UserToken)

	appid := cache.Get(utils.WECHAT_APPID)
	secret := cache.Get(utils.WECHAT_APPSECRET)

	url = strings.Replace(url, "APPID", appid, -1)
	url = strings.Replace(url, "SECRET", secret, -1)
	url = strings.Replace(url, "CODE", code, -1)

	log4go.Debug(url)

	resp, err := http.Get(url)

	if err != nil {
		return ut
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log4go.Debug(string(body))

	json.Unmarshal(body, &ut)

	return ut
}

func (u *UserInfo) GetUserInfo(openid, token string) *UserInfo {

	url := "https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN"
	user := new(UserInfo)
	cache := new(utils.Cache)

	cu := cache.Get(utils.WECHAT_USER_INFO + token)

	if cu != "" {
		log4go.Debug(cu)
		json.Unmarshal([]byte(cu), &user)
		return user
	}

	url = strings.Replace(url, "ACCESS_TOKEN", token, -1)
	url = strings.Replace(url, "OPENID", openid, -1)

	log4go.Debug(url)

	resp, err := http.Get(url)

	if err != nil {
		return user
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log4go.Debug(string(body))

	json.Unmarshal(body, &user)
	cache.Set(utils.WECHAT_USER_INFO+token, string(body), time.Second*7200)

	return user
}
