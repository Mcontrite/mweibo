package conf

import (
	"github.com/go-ini/ini"
)

type DBConfig struct {
	DBType     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

var DBconfig = &DBConfig{}

type MailConfig struct {
	MailDriver   string
	MailHost     string
	MailPort     string
	MailUsername string
	MailPassword string
	MailSender   string
}

var Mailconfig = &MailConfig{}
<<<<<<< HEAD
=======

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
	JWTSecretKey   string
}

var Serverconfig = &ServerConfig{}
>>>>>>> 4f21432... fix ini-config

func InitConfig() {
	var cfg *ini.File
	cfg, _ = ini.Load("conf/config.ini")
	err := cfg.Section("SERVER").MapTo(Serverconfig)
	if err != nil {
		return
	}
	err = cfg.Section("DB").MapTo(DBconfig)
	err = cfg.Section("MAIL").MapTo(Mailconfig)
}
