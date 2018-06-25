package dbinit

import (
	"github.com/jinzhu/gorm"
	"github.com/xpp666/go.orm/model"
	"github.com/xpp666/go.orm/util"
	"strconv"
)

var Db *gorm.DB

func init() {
	var err error
	conf := util.ConfigUtil()
	mysqlStr := conf.Mysql_user + ":" + conf.Mysql_password + "@tcp(" + conf.Mysql_host + ":" +conf.Mysql_port + ")/" + conf.Database_name + "?parseTime=true&&charset=utf8"
	Db, err = gorm.Open("mysql", mysqlStr)
	if err != nil {
		panic("连接出错:")
	}
	if !Db.HasTable(&model.User{}) {
		if err := Db.CreateTable(&model.User{}).Error; err != nil {
			panic(err)
		}
	}
	idle,err:=strconv.Atoi(conf.Max_idle_conns)
	if err!= nil {
		panic(err)
	}
	open,err:=strconv.Atoi(conf.Max_open_conns)
	if err!= nil {
		panic(err)
	}
	Db.DB().SetMaxIdleConns(idle)
	Db.DB().SetMaxOpenConns(open)
}

