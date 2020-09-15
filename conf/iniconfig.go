package conf

import (
	"time"

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

type RedisConfig struct {
	RedisHost        string
	RedisPassword    string
	RedisMaxidle     int
	RedisMaxActive   int
	RedisIdleTimeout time.Duration
}

var Redisconfig = &RedisConfig{}

type ImageConfig struct {
	ImageSavePath   string
	ImageMaxSize    int
	ImageAlloweXts  string
	ImageAllowExts  []string
	RuntimeRootPath string
}

var Imageconfig = &ImageConfig{}

type MailConfig struct {
	MailDriver   string
	MailHost     string
	MailPort     string
	MailUsername string
	MailPassword string
	MailSender   string
}

var Mailconfig = &MailConfig{}

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

func InitConfig() {
	var cfg *ini.File
	cfg, _ = ini.Load("conf/config.ini")
	err := cfg.Section("SERVER").MapTo(Serverconfig)
	if err != nil {
		return
	}
	err = cfg.Section("DB").MapTo(DBconfig)
	err = cfg.Section("REDIS").MapTo(Redisconfig)
	err = cfg.Section("IMAGE").MapTo(Imageconfig)
	err = cfg.Section("MAIL").MapTo(Mailconfig)
}
