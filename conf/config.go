package conf

import "github.com/go-ini/ini"

type DBConfig struct {
	DBType string
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBName string
}

var DBConf = &DBConfig{}

type ServerConfig struct {
	ServerName    string
	ServerPort    int
	ServerRunmode string
	ServerURL     string
	ServerKey     string
	SerssionKey   string
	StaticPath    string
	ViewPath      string
}

var ServerConf = &ServerConfig{}

func InitConfig() {
	var cfg *ini.File
	cfg, _ = ini.Load("conf/config.ini")
	err := cfg.Section("SERVER").MapTo(ServerConf)
	err = cfg.Section("DB").MapTo(DBConf)
	if err != nil {
		return
	}
}
