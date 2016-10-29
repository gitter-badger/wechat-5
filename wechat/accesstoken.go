package wechat

import (
	"encoding/json"
	"github.com/alecthomas/log4go"
	"io/ioutil"
	"net/http"
	"nicosoft.org/wechat/utils"
	"strings"
)

type AccessToken struct {
	Token     string `json:"access_token"`
	ExpiresIn int    `json:"expires_in"`
}

func (a *AccessToken) Get() *AccessToken {

	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET"
	result := new(AccessToken)

	cache := new(utils.Cache)

	at := cache.Get(utils.WECHAT_ACCESSTOKEN)

	if at != "" {
		log4go.Debug(at)
		json.Unmarshal([]byte(at), &result)
		return result
	}

	appid := cache.Get(utils.WECHAT_APPID)
	appsecret := cache.Get(utils.WECHAT_APPSECRET)

	url = strings.Replace(url, "APPID", appid, -1)
	url = strings.Replace(url, "APPSECRET", appsecret, -1)

	log4go.Debug(url)

	resp, err := http.Get(url)

	if err != nil {
		return result
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log4go.Debug(string(body))

	json.Unmarshal(body, &result)
	cache.Set(utils.WECHAT_ACCESSTOKEN, string(body), 0)

	return result
}
