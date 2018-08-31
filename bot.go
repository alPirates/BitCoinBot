package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"

	tgbot "gopkg.in/telegram-bot-api.v4"
)

func startBot(config Config) *tgbot.BotAPI {

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

func startBotWork(bot *tgbot.BotAPI, config Config) {
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

func sendMessages(bot *tgbot.BotAPI, config Config, messages []string) {
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

func sendMessagesByMail(config Config, message string) {

}

func sendMessagesByTelegram(bot *tgbot.BotAPI, config Config, message string) error {
	if config.ChatID == 0 {
		return errors.New("no chat id")
	}
	msg := tgbot.NewMessage(config.ChatID, message)
	_, err := bot.Send(msg)
	return err
}
