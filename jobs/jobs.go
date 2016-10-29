package jobs

import (
	"encoding/json"
	"github.com/alecthomas/log4go"
	"nicosoft.org/wechat/utils"
	"nicosoft.org/wechat/wechat"
	"time"
)

func Run() {
	go AccessTokenJob()
}

func AccessTokenJob() {

	log4go.Debug("start jobs::AccessTokenJob")

	ticker := time.NewTicker(time.Second * 7200)
	cache := new(utils.Cache)

	for {
		<-ticker.C

		at := new(wechat.AccessToken).Get()
		atByte, _ := json.Marshal(at)
		atStr := string(atByte)

		cache.Del(utils.WECHAT_ACCESSTOKEN)
		cache.Set(utils.WECHAT_ACCESSTOKEN, atStr, 0)
	}

}
