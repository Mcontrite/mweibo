package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	ServerName     string `yaml:"SERVER_NAME"`
	ServerRunmode  string `yaml:"SERVER_RUNMODE"`
	ServerPort     string `yaml:"SERVER_PORT"`
	ServerURL      string `yaml:"SERVER_URL"`
	ServerKey      string `yaml:"SERVER_KEY"`
	SessionKey     string `yaml:"SERVER_SessionKey"`
	ContextUserKey string `yaml:"SERVER_ContextUserKey"`
	StaticPath     string `yaml:"SERVER_StaticPath"`
	ViewsPath      string `yaml:"SERVER_ViewsPath"`
	EnableCsrf     bool   `yaml:"SERVER_EnableCsrf"`
	CsrfParamName  string `yaml:"SERVER_CsrfParamName"`
	CsrfHeaderName string `yaml:"SERVER_CsrfHeaderName"`

	DBType     string `yaml:"DB_TYPE"`
	DBHost     string `yaml:"DB_HOST"`
	DBPort     string `yaml:"DB_PORT"`
	DBName     string `yaml:"DB_NAME"`
	DBUser     string `yaml:"DB_USER"`
	DBPassword string `yaml:"DB_PASSWORD"`
	DSN        string `yaml:"DSN"`

	MailDriver   string `yaml:"MAIL_DRIVER"`
	MailHost     string `yaml:"MAIL_HOST"`
	MailPort     string `yaml:"MAIL_PORT"`
	MailUsername string `yaml:"MAIL_USERNAME"`
	MailPassword string `yaml:"MAIL_PASSWORD"`
	MailSender   string `yaml:"MAIL_SENDER"`
}

var configuration *Configuration

func GetConfiguration() *Configuration {
	return configuration
}

func LoadConfiguration(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		//log.Fatal("Read config file error: ", err)
		return err
	}
	var config Configuration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil
	}
	configuration = &config
	return err
}
