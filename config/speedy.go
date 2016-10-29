package config

import (
	"github.com/alecthomas/log4go"
	"github.com/dlintw/goconf"
	"gopkg.in/redis.v5"
	"path"
	"runtime"
)

type WeConf struct {
	Token      string
	Appid      string
	Appsecret  string
	Oauth2Url  string
	Oauth2Type string
}

var (
	Cache    *redis.Client
	WeChat   *WeConf
	filepath string
)

func InitConf() {
	_, filename, _, _ := runtime.Caller(0)
	filepath = path.Dir(filename) + "/speedy.conf"
}

func InitCache() {

	config, _ := goconf.ReadConfigFile(filepath)

	host, _ := config.GetString("redis", "redis.host")
	passwd, _ := config.GetString("redis", "redis.password")
	poolsize, _ := config.GetInt("redis", "redis.poolsize")

	Cache = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: passwd,
		DB:       0,
		PoolSize: poolsize,
	})

	res, err := Cache.Ping().Result()

	if err != nil {
		log4go.Error(err)
	} else {
		log4go.Debug(res)
	}

}

func InitWeConf() {

	config, _ := goconf.ReadConfigFile(filepath)

	token, _ := config.GetString("wechat", "wechat.token")
	appid, _ := config.GetString("wechat", "wechat.appid")
	appsecret, _ := config.GetString("wechat", "wechat.appsecret")
	oauth2Url, _ := config.GetString("wechat", "wechat.oauth2.url")
	oauth2Type, _ := config.GetString("wechat", "wechat.oauth2.type")

	WeChat = &WeConf{Token: token, Appid: appid, Appsecret: appsecret, Oauth2Url: oauth2Url, Oauth2Type: oauth2Type}
}
