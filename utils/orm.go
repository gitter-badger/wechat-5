package utils

import (
	"fmt"
	"github.com/alecthomas/log4go"
	"github.com/dlintw/goconf"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"nicosoft.org/wechat/config"
)

var Orm *xorm.Engine

func InitOrm() {

	config, _ := goconf.ReadConfigFile(config.Filepath)

	driver, _ := config.GetString("database", "database.driver")
	dbname, _ := config.GetString("database", "database.dbname")
	user, _ := config.GetString("database", "database.user")
	password, _ := config.GetString("database", "database.password")
	host, _ := config.GetString("database", "database.host")
	prefix, _ := config.GetString("database", "database.prefix")

	var err error
	Orm, err = xorm.NewEngine(driver, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user, password, host, dbname))
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, prefix)
	Orm.SetTableMapper(tbMapper)

	if err != nil {
		log4go.Error(err)
	} else {
		Orm.Ping()
	}
}
