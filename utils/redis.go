package utils

import (
	"github.com/alecthomas/log4go"
	"nicosoft.org/wechat/config"
	"time"
)

const (
	WECHAT_TOKEN       = "config:wechat:token"
	WECHAT_APPID       = "config:wechat:appid"
	WECHAT_APPSECRET   = "config:wechat:appsecret"
	WECHAT_ACCESSTOKEN = "config:wechat:accesstoken"
)

type Cache struct {
	Db int
}

func LoadCache() {
	wechat := config.WeChat
	cache := new(Cache)
	cache.Set(WECHAT_TOKEN, wechat.Token, 0)
	cache.Set(WECHAT_APPID, wechat.Appid, 0)
	cache.Set(WECHAT_APPSECRET, wechat.Appsecret, 0)

}

func (c *Cache) Set(key string, value interface{}, expiration time.Duration) bool {
	config.Cache.Pipeline().Select(c.Db)

	err := config.Cache.Set(key, value, expiration).Err()
	if err != nil {
		log4go.Error(err)
		return false
	}
	return true
}

func (c *Cache) Get(key string) string {
	config.Cache.Pipeline().Select(c.Db)

	res, err := config.Cache.Get(key).Result()
	if err != nil {
		log4go.Error(err)
		return ""
	}
	return res
}

func (c *Cache) Del(key string) int64 {
	config.Cache.Pipeline().Select(c.Db)

	res, err := config.Cache.Del(key).Result()
	if err != nil {
		log4go.Error(err)
		return -1
	}
	return res
}
