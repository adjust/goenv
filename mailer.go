package goenv

import (
	"fmt"
	"net/smtp"
)

func (goenv *Goenv) Mail(recipients []string, subject, body string) {
	smtp.SendMail(
		goenv.Get("mailServer", "")+":25",
		goenv.NewAuth(),
		goenv.Get("mailFrom", ""),
		recipients,
		NewMail(subject, body),
	)
}

func (goenv *Goenv) NewAuth() smtp.Auth {
	return smtp.PlainAuth(
		"",
		goenv.Get("mailFrom", ""),
		goenv.Get("mailPasswd", ""),
		goenv.Get("mailServer", ""),
	)
}

func NewMail(subject, body string) []byte {
	mail := fmt.Sprintf(
		"Subject:%s\n\n%s",
		subject,
		body,
	)
	return []byte(mail)
}
