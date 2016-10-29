package test

import (
	"nicosoft.org/wechat/config"
	"nicosoft.org/wechat/utils"
	"nicosoft.org/wechat/wechat"
	"testing"
)

func init() {
	config.InitConf()
	config.InitCache()
	config.InitWeConf()

	utils.LoadCache()
}

func Test_AT(t *testing.T) {

	at := new(wechat.AccessToken)
	data := at.Get()
	t.Log(data)

}
