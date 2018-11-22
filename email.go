package main

import (
	"errors"
	"log"
	"net/smtp"
	"regexp"
)

func getSmtpName(name string) (map[string]string, error) {
	servers := map[string]map[string]string{
		"yandex": map[string]string{
			"smtp": "smtp.yandex.ru",
			"port": "smtp.yandex.ru:25",
		},
		"ya": map[string]string{
			"smtp": "smtp.yandex.ru",
			"port": "smtp.yandex.ru:25",
		},
		"gmail": map[string]string{
			"smtp": "smtp.gmail.com",
			"port": "smtp.gmail.com:587",
		},
	}
	rgx := regexp.MustCompile(`@(\w+)`)
	output := rgx.FindStringSubmatch(name)
	if len(output) > 1 {
		r := servers[output[1]]
		if r != nil {
			return r, nil
		}
		return nil, errors.New("Cant find smtp server in list")
	}
	return nil, errors.New("Cant find email addres regexp in string")
}

func sendMessageByEmail(message string) {
	email := config.Email
	password := config.Password
	serverSMTP, err := getSmtpName(email)

	if err != nil {
		log.Println(err)
		return
	}

	auth := smtp.PlainAuth("", email, password, serverSMTP["smtp"])
	to := []string{email}
	msg := []byte(
		"Subject: Новая заявка!\r\n" +
			"\r\n" +
			message +
			"\r\n",
	)
	err = smtp.SendMail(serverSMTP["port"], auth, email, to, msg)
	if err != nil {
		log.Fatal(err)
	}

}
