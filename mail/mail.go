package mail

import (
	"crypto/tls"
	"fmt"
	gomail "gopkg.in/mail.v2"
)

func SendMail(name, vacancy string) {
	from := "sovcom.people@mail.ru"
	password := "JGRCPSKPeEhQdf1QiSJY"

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", from)

	// Set E-Mail receivers
	m.SetHeader("To", "myfactoryisgood@gmail.com")

	// Set E-Mail subject
	m.SetHeader("Subject", "Gomail test subject")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "This is Gomail test body")

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.mail.ru", 465, from, password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return
}
