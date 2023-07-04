package my_utils

import (
	"fmt"
	"math/rand"

	"gopkg.in/gomail.v2"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func SendMail(to string, subject string, message string) {
	mailer := gomail.NewDialer("sandbox.smtp.mailtrap.io", 2525, "8fab621c0c38f8", "adbfaccdb1f897")

	msg := gomail.NewMessage()

	msg.SetHeader("To", to)
	msg.SetHeader("subject", subject)
	msg.SetHeader("From", "admin@go-fintech.com")
	msg.SetBody("text/html", message)

	if err := mailer.DialAndSend(msg); err != nil {
		fmt.Println(err)
		fmt.Println("Email not sent...")
	}
}
