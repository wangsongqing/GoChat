package mail

import (
	"GoChat/pkg/logger"
	"fmt"
	emailPKG "github.com/jordan-wright/email"
	"net/smtp"
	_ "os"
)

// SMTP 实现 email.Driver interface
type SMTP struct {
}

// Send 实现 email.Driver interface 的 Send 方法

func (s *SMTP) Send(email Email, config map[string]string) bool {
	e := emailPKG.NewEmail()

	e.From = fmt.Sprintf("%v <%v>", email.From.Name, email.From.Address)
	e.To = email.To
	e.Bcc = email.Bcc
	e.Cc = email.Cc
	e.Subject = email.Subject
	e.Text = email.Text
	e.HTML = email.HTML

	logger.DebugJSON("发送邮件", "发送详情", e)

	err := e.Send(
		fmt.Sprintf("%v:%v", config["host"], config["port"]),
		smtp.PlainAuth(
			"",
			config["username"],
			config["password"],
			config["host"],
		),
	)

	if err != nil {
		logger.ErrorString("发送邮件", "发送出错", err.Error())
		return false
	}

	logger.DebugString("发送邮件", "发送成功", "")

	return true
}
