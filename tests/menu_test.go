package test

import (
	"encoding/json"
	"github.com/dlintw/goconf"
	"nicosoft.org/wechat/config"
	"nicosoft.org/wechat/utils"
	"nicosoft.org/wechat/wechat"
	"testing"
)

func init() {
	config.InitCache()
	config.InitWeConf()

	utils.LoadCache()
}

func Test_Create(t *testing.T) {

	conf, _ := goconf.ReadConfigFile("/Users/Nico/Documents/software/gopath/src/nicosoft.org/speedy/config/speedy.conf")
	data, _ := conf.GetString("wechat", "wechat.test.data")

	menu := new(wechat.Menu)
	err := json.Unmarshal([]byte(data), &menu)

	if err != nil {
		t.Log(err)
	}

	res := menu.Create()

	t.Log(res)

}
