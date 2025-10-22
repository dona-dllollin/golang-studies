package main 

import (
	"fmt"
	"net/smtp"
	"log"
	"strings"
)

const (
	CONFIG_SMTP_HOST = "smtp.gmail.com"
	CONFIG_SMTP_PORT = 587
	CONFIG_SENDER_NAME = "PT. Makmur subur Jaya <donadlollin@gmail.com>"
	CONFIG_AUTH_EMAIL = "donadlollin@gmail.com"
	CONFIG_AUTH_PASSWORD = "yorv gnyd ycsa femb"
)

func main() {
	to := []string{"donadlollin2@gmail.com", "desaloning@gmail.com"}
	cc := []string{"tralalala@gmail.com"}
	subject := "test email golang"
	message := "this is test email using golang smtp package"

	err := sendMail(to, cc, subject, message)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Email sent successfully")
}

func sendMail(to []string, cc []string, subject string, message string) error {

	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message
	
	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}
	return nil
}