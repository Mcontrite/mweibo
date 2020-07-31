package main

import (
	"flag"
	"mweibo/conf"

	"github.com/lexkong/log"
)

func main() {
	conf.InitLog()
	configuration := flag.String("C", "conf/config.yaml", "Config File Path")
	flag.Parse()
	err := conf.LoadConfiguration(*configuration)
	if err != nil {
		log.Fatalln("Read config file error...")
		return
	}
}
