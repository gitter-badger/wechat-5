package wechat

import (
	"encoding/json"
	"github.com/alecthomas/log4go"
	"io/ioutil"
	"net/http"
	"strings"
)

type SubButton struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Key  string `json:"key"`
}

type Button struct {
	Name      string      `json:"name"`
	SubButton []SubButton `json:"sub_button"`
}

type Menu struct {
	Button []Button `json:"button"`
}

type Result struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (m *Menu) Create(menuStr ...string) *Result {

	var data string
	url := "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN"
	res := &Result{Errcode: -1, Errmsg: "create failed."}

	if len(menuStr) > 0 {
		data = menuStr[0]
	} else {
		tmp, err := json.Marshal(m)
		if err != nil {
			log4go.Error(err)
		}
		data = string(tmp)
	}

	at := new(AccessToken).Get()
	url = strings.Replace(url, "ACCESS_TOKEN", at.Token, -1)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data))

	if err != nil {
		log4go.Debug(err)
		return res
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, res)
	return res
}
