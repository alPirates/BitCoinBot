package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	tgbot "gopkg.in/telegram-bot-api.v4"
)

func parse(bot *tgbot.BotAPI, config Config) {
	for true {
		messages := []string{}
		startTime := time.Now()

		doc, err := goquery.NewDocument(config.HTMLURL1)
		if err == nil {
			doc.Find("#cbi-table-table").Each(func(i int, s1 *goquery.Selection) {
				s1.Find("#cbi-table-1-status").Each(func(arg1 int, s2 *goquery.Selection) {

					text := s2.Text()
					switch i {
					case 1:
						if text != "" && text != "Alive" {
							messages = append(messages, `поле "Alive" изменено`)
						}
						break
					case 2:
						if text != "" {
							statuses := strings.Split(s2.Text(), " ")
							for i := 1; i < len(statuses)-1; i++ {
								if statuses[i] != "oooooooo" {
									messages = append(messages, `поле "oooooooo" изменено`)
								}
							}
							if statuses[len(statuses)-1] != "ooooooo" {
								messages = append(messages, `поле "oooooooo" изменено`)
							}
						}
						break
					default:
						break
					}

				})
				s1.Find("#cbi-table-1-temp2").Each(func(arg1 int, s2 *goquery.Selection) {
					text := s2.Text()
					if text != "" {
						temperatura, errT := strconv.Atoi(text)
						if errT != nil {
							sendError("температура не является числом : " + err.Error())
						} else if temperatura > config.Temperature {
							messages = append(messages, `температура превышена (`+text+")")
						}
					}
				})
			})

		} else {
			sendError("can't open your website : " + err.Error())
		}

		doc, err = goquery.NewDocument(config.HTMLURL2)
		if err == nil {
			doc.Find("#cbi-table-table").Each(func(i int, s1 *goquery.Selection) {
				s1.Find("#cbi-table-1-status").Each(func(arg1 int, s2 *goquery.Selection) {

					text := s2.Text()
					switch i {
					case 1:
						if text != "" && text != "Alive" {
							messages = append(messages, `поле "Alive" изменено`)
						}
						break
					case 2:
						if text != "" {
							statuses := strings.Split(s2.Text(), " ")
							for i := 1; i < len(statuses); i++ {
								if statuses[i] != "oooooooo" {
									messages = append(messages, `поле "oooooooo" изменено`)
								}
							}
						}
						break
					default:
						break
					}

				})
			})

		} else {
			sendError("can't open your website : " + err.Error())
		}

		if len(messages) != 0 {
			go sendMessages(bot, config, messages)
		}

		time.Sleep(time.Duration(config.UpdateTime)*time.Minute - time.Now().Sub(startTime))
	}
}
