package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"net/url"
	"os"

	"github.com/alPirates/BitCoinBot/config"
	tgbot "gopkg.in/telegram-bot-api.v4"
)

func startBot(config config.Config) *tgbot.BotAPI {

	var err error
	proxyURL, _ := url.Parse(config.Proxy)
	myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	bot, err := tgbot.NewBotAPIWithClient(config.Token, myClient)
	if err != nil {
		sendError("can't start bot : " + err.Error())
		os.Exit(0)
	}

	sendError("start " + bot.Self.UserName)

	return bot
}

func startBotWork(bot *tgbot.BotAPI, config config.Config) {
	channel := tgbot.NewUpdate(0)
	channel.Timeout = 30
	updates, _ := bot.GetUpdatesChan(channel)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Command() != "" {
			switch update.Message.Command() {
			case "start":
				msg := tgbot.NewMessage(update.Message.Chat.ID, "/id - return your id")
				bot.Send(msg)
				break
			case "help":
				msg := tgbot.NewMessage(update.Message.Chat.ID, "/id - return your id")
				bot.Send(msg)
				break
			case "id":
				msg := tgbot.NewMessage(
					update.Message.Chat.ID,
					fmt.Sprint(update.Message.Chat.ID),
				)
				bot.Send(msg)
				break
			}
		}

	}

}

func sendMessages(bot *tgbot.BotAPI, config config.Config, messages []string) {
	mes := ""
	for i, str := range messages {
		mes += str
		if i != len(messages)-1 {
			mes += "\n"
		}
	}
	err := sendMessagesByTelegram(bot, config, mes)
	if err != nil {
		sendMessagesByMail(config, mes)
	}
}

func sendMessagesByMail(config config.Config, message string) {
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

func sendMessagesByTelegram(bot *tgbot.BotAPI, config config.Config, message string) error {
	if config.ChatID == 0 {
		return errors.New("no chat id")
	}
	msg := tgbot.NewMessage(config.ChatID, message)
	_, err := bot.Send(msg)
	return err
}
