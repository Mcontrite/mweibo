package main

import (
	"fmt"
	"mweibo2/conf"
	"mweibo2/model"
	"mweibo2/route"
)

func main() {
	conf.InitConfig()
	db := model.InitDB()
	db.AutoMigrate(
		&model.User{},
		&model.Weibo{},
		&model.Comment{},
	)
	defer db.Close()
	e := route.InitRoute()
	e.Run(fmt.Sprintf(":%d", conf.ServerConf.ServerPort))
}
