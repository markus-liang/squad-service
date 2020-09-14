package helpers

import (
	"net/smtp"
)

func SendMail(to []string, subject string, msg string) (bool, error) {
	host := Env("MAIL_HOST")
	port := Env("MAIL_PORT")
	addr := host + ":" + port
	from := Env("MAIL_USER")
	pass := Env("MAIL_PASS")
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject = "Subject: " + subject + " \n"
	msg = subject + mime + "\n" + msg
	auth := smtp.PlainAuth("", from, pass, host)

	if err := smtp.SendMail(addr, auth, from, to, []byte(msg)); err != nil {
		return false, err
	}
	return true, nil
}

