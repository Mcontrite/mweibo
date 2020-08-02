package main

import (
	"flag"
	"mweibo/conf"
	"mweibo/model"
	"mweibo/router"

	"github.com/lexkong/log"
)

func main() {
	configpath := flag.String("C", "conf/config.yaml", "Config File Path")
	flag.Parse()
	conf.InitLog()
	err := conf.LoadConfiguration(*configpath)
	if err != nil {
		log.Fatal("Read config file error: ", err)
		panic("Read config file error...")
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
	g := router.InitRouter()
	//g.Run(conf.GetConfiguration().ServerPort)
	g.Run(":8080")
}
