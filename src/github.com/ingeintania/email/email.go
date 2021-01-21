package email

import (
	"fmt"
	"log"
	"net/smtp"
)

const ConfigSmtpHost = "smtp.gmail.com"
const ConfigSmtpPort = 587
const ConfigEmail = "pesutmahakam168@gmail.com"
const ConfigPassword = ""

func sendMail(to []string, subject string, message string) error {

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n";
	msg := subject + mime + message

	auth := smtp.PlainAuth("", ConfigEmail, ConfigPassword, ConfigSmtpHost)
	smtpAddr := fmt.Sprintf("%s:%d", ConfigSmtpHost, ConfigSmtpPort)

	err := smtp.SendMail(smtpAddr, auth, ConfigEmail, append(to), []byte(msg))
	if err != nil {
		return err
	}

	return nil
}

func CodeEmail(destination []string, code int) {

	subject := "Subject: TPAWEBEI Company\n"
	message := "Kode Verifikasi anda adalah mantappumantoel"

	err := sendMail(destination, subject, message)

	if err != nil {
		log.Println("New Mail Failed to Send !\n" + err.Error())
		return
	}

	log.Println("New Mail Sent to " + destination[0] + "!")
}

func SubscriptionEmail(destination []string) {

	subject := "Subject: Karcis Company\n"
	message := "Selamat anda sudah bergabung!"
	err := sendMail(destination, subject, message)

	if err != nil {
		log.Println("New Mail Failed to Send !\n" + err.Error())
		return
	}

	log.Println("New Mail Sent to " + destination[0] + "!")
}
