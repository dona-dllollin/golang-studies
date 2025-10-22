package main

import (
	"gopkg.in/gomail.v2"
	"log"
)

const (
	CONFIG_SMTP_HOST     = "smtp.gmail.com"
	CONFIG_SMTP_PORT     = 587
	CONFIG_SENDER_NAME   = "PT. Makmur Subur Jaya <donadlollin@gmail.com>"
	CONFIG_AUTH_EMAIL    = "donadlollin@gmail.com"
	CONFIG_AUTH_PASSWORD = "yorv gnyd ycsa femb" // gunakan App Password, bukan password biasa!
)

func main() {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", "donadlollin2@gmail.com", "emaillain@gmail.com")
	mailer.SetAddressHeader("Cc", "tralalala@gmail.com", "Tra Lala La")
	mailer.SetHeader("Subject", "Test mail")
	mailer.SetBody("text/html", "Hello, <b>have a nice day</b>")
	mailer.Attach("./sample.png")

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Email sent successfully")
}
