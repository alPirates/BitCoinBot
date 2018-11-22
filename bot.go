package main

import (
	"log"
	"net/smtp"
)

func SendMessagesByMail(config Config, message string) {
	email := config.Email
	password := config.Password

	auth := smtp.PlainAuth("", email, password, "smtp.yandex.ru")
	to := []string{email}
	msg := []byte(
		"Subject: Предупреждение\r\n" +
			"\r\n" +
			message +
			"\r\n",
	)
	err := smtp.SendMail("smtp.yandex.ru:25", auth, email, to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
