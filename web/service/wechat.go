package service

import (
	"nicosoft.org/wechat/utils"
	"strings"
)

func GetOAath2Url() string {

	redirect := "https://open.weixin.qq.com/connect/oauth2/authorize?appid=APPID&redirect_uri=REDIRECT_URI&response_type=code&scope=SCOPE&state=STATE#wechat_redirect"
	cache := new(utils.Cache)

	oaUrl := cache.Get(utils.WECHAT_OAUTH2_URL)
	oaType := cache.Get(utils.WECHAT_OAUTH2_TYPE)
	appid := cache.Get(utils.WECHAT_APPID)

	strings.Replace(redirect, "APPID", appid, -1)
	strings.Replace(redirect, "REDIRECT_URI", oaUrl, -1)
	strings.Replace(redirect, "SCOPE", oaType, -1)

	return redirect
}
