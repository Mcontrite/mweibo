package main

import (
	"fmt"
	"mweibo/conf"
	"mweibo/model"
	"mweibo/router"
)

func main() {
	// configpath := flag.String("C", "conf/config.yaml", "Config File Path")
	// flag.Parse()
	// conf.InitLog()
	// err := conf.LoadConfiguration(*configpath)
	// if err != nil {
	// 	log.Fatal("Read config file error: ", err)
	// 	panic("Read config file error...")
	// }
	conf.InitConfig()
	db := model.InitDB()
	db.AutoMigrate(
		&model.User{},
		&model.Follower{},
		&model.Weibo{},
		&model.Comment{},
		&model.Tag{},
		&model.TagWeibo{},
		&model.PwdReset{},
		&model.Attach{},
	)
	defer db.Close()
	g := router.InitRouter()
	//g.Run(conf.GetConfiguration().ServerPort)

	g.Run(fmt.Sprintf(":%d", conf.Serverconfig.ServerPort))
	//g.Run(":8080")
}

// log mail csrf passwordreset pagination
// group image upload attach favorite ip
// myweibos captcha auth admin collections likes
// ROUTE tips
// create weibo while user not exsit
// editor
