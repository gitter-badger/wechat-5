package main

import (
	"nicosoft.org/wechat/config"
	"nicosoft.org/wechat/jobs"
	"nicosoft.org/wechat/utils"
	"nicosoft.org/wechat/web"
)

func init() {
	config.InitCache()
	config.InitWeConf()

	utils.LoadCache()

	jobs.Run()
}

func main() {
	r := new(web.Router)
	r.Start()
}
