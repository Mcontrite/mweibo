package mail

import (
	"bytes"
	"crypto/tls"
	"errors"
	"html/template"
	"mweibo/conf"
	"mweibo/middleware/logging"
	"strconv"

	gomail "gopkg.in/gomail.v2"
)

type Mail struct {
	Driver   string   // smtp or log
	Host     string   // 邮箱的服务器地址
	Port     int      // 邮箱的服务器端口
	Username string   // 发送者的 name
	Password string   // 授权码或密码
	Sender   string   // 用来作为邮件的发送者名称
	Receiver []string // 发送目标
	Title    string   // 邮件标题
	Content  string   // 邮件内容
}

func (mail *Mail) sendByLog() error {
	logging.Info(mail.Content)
	return nil
}

func (mail *Mail) sendBySMTP() error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", mail.Sender+"<"+mail.Username+">")
	msg.SetHeader("To", mail.Receiver...)
	msg.SetHeader("Title", mail.Title)
	msg.SetBody("text/html", mail.Content)
	dialer := gomail.NewDialer(mail.Host, mail.Port, mail.Username, mail.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return dialer.DialAndSend(msg)
}

// 读取模板并转换为 string
func TemplateToString(tplName, tplPath string, tplData map[string]interface{}) (string, error) {
	tpl, err := template.New(tplName).ParseFiles(tplPath)
	if err != nil {
		return "", err
	}
	buffer := new(bytes.Buffer)
	if err = tpl.Execute(buffer, tplData); err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func SendMail(receiver []string, title, tplPath string, tplData map[string]interface{}) error {
	filePath := "views/" + tplPath
	port, _ := strconv.Atoi(conf.Mailconfig.MailPort)
	content, _ := TemplateToString(tplPath, filePath, tplData)
	mail := &Mail{
		Driver:   conf.Mailconfig.MailDriver,
		Host:     conf.Mailconfig.MailHost,
		Port:     port,
		Username: conf.Mailconfig.MailUsername,
		Password: conf.Mailconfig.MailPassword,
		Sender:   conf.Mailconfig.MailSender,
		Receiver: receiver,
		Title:    title,
		Content:  content,
	}
	return mail.Send()
}

func (mail *Mail) Send() error {
	if mail.Driver == "log" {
		return mail.sendByLog()
	} else if mail.Driver == "smtp" {
		return mail.sendBySMTP()
	}
	return errors.New("不支持: " + mail.Driver + " 类型的 MailDriver。")
}
