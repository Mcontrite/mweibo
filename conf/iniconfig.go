package conf

import (
	"github.com/go-ini/ini"
)

type ServerConfig struct {
	ServerName     string
	ServerRunmode  string
	ServerPort     int
	ServerURL      string
	ServerKey      string
	SessionKey     string
	ContextUserKey string
	StaticPath     string
	ViewsPath      string
	EnableCsrf     bool
	CsrfParamName  string
	CsrfHeaderName string
}

var Serverconfig = &ServerConfig{}

type DBConfig struct {
	DBType     string
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	DSN        string
}

var dbconfig = &DBConfig{}

type MailConfig struct {
	MailDriver   string
	MailHost     string
	MailPort     string
	MailUsername string
	MailPassword string
	MailSender   string
}

var mailconfig = &MailConfig{}

func InitConfig() {
	// cfg, err := ini.Load("conf/config.ini")
	// if err != nil {
	// 	return
	// }
	var cfg *ini.File
	cfg, _ = ini.Load("conf/config.ini")
	err := cfg.Section("SERVER").MapTo(Serverconfig)
	if err != nil {
		return
	}
	err = cfg.Section("DB").MapTo(dbconfig)
	err = cfg.Section("MAIL").MapTo(mailconfig)
}
