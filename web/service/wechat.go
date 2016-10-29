package service

import (
	"encoding/json"
	"github.com/alecthomas/log4go"
	"io/ioutil"
	"net/http"
	"nicosoft.org/wechat/utils"
	"nicosoft.org/wechat/web/models"
	"strings"
)

type UserToken struct {
	Token     string `json:"access_token"`
	ExpiresIn int    `json:"expires_in"`
	NewToken  string `json:"refresh_token"`
	OpenId    string `json:"openid"`
	Scope     string `json:"scope"`
}

func GetOAath2Url() string {

	redirect := "https://open.weixin.qq.com/connect/oauth2/authorize?appid=APPID&redirect_uri=REDIRECT_URI&response_type=code&scope=SCOPE&state=STATE#wechat_redirect"
	cache := new(utils.Cache)

	oaUrl := cache.Get(utils.WECHAT_OAUTH2_URL)
	oaType := cache.Get(utils.WECHAT_OAUTH2_TYPE)
	appid := cache.Get(utils.WECHAT_APPID)

	redirect = strings.Replace(redirect, "APPID", appid, -1)
	redirect = strings.Replace(redirect, "REDIRECT_URI", oaUrl, -1)
	redirect = strings.Replace(redirect, "SCOPE", oaType, -1)

	return redirect
}

func (u *UserToken) GetToken(code string) *UserToken {

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

func (u *UserToken) GetUserInfo(openid, token string) *models.UserInfo {

	url := "https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN"
	user := new(models.UserInfo)

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

	return user
}
