package main

import (
	"flag"
	"mweibo/conf"
	"mweibo/model"
)

func main() {
	conf.InitLog()
	configuration := flag.String("C", "conf/config.yaml", "Config File Path")
	flag.Parse()
	err := conf.LoadConfiguration(*configuration)
	if err != nil {
		panic("Read config file error...")
		return
	}
	db := model.InitDB()
	db.AutoMigrate(
		&model.User{},
		&model.Follower{},
		&model.Weibo{},
		&model.Comment{},
		&model.Tag{},
		&model.TagWeibo{},
		&model.PwdReset{},
	)
	defer db.Close()
}
