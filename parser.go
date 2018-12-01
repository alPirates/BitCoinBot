package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func loopParse(u *UiService, config *Config) {
	time.Sleep(1 * time.Second)
	for true {
		startTime := time.Now()
		parse(u, config)
		time.Sleep(time.Duration(config.UpdateTime)*time.Minute - time.Now().Sub(startTime))
	}
}

func parse(u *UiService, config *Config) {
	messages := []string{}
	temps := make([]int, 0)
	u.LogError("[PARSE] [начинаю парсить](fg-green)")

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
					}
					break
				}

			})
			s1.Find("#cbi-table-1-temp2").Each(func(arg1 int, s2 *goquery.Selection) {
				text := s2.Text()
				if text != "" {
					text = text[6:]
					temperatura, errT := strconv.Atoi(text)
					temps = append(temps, temperatura)
					if errT != nil {
						u.LogError("[ERR] [температура не является числом : " + errT.Error() + "](fg-red)")
					} else if temperatura > config.Temperature {
						messages = append(messages, `температура превышена (`+text+")")
					}
				}
			})
		})

	} else {
		u.LogError("[ERR] [не удалось открыть url1 : " + err.Error() + "](fg-red)")
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
						for i := 1; i < len(statuses)-1; i++ {
							if statuses[i] != "oooooooo" {
								messages = append(messages, `поле "oooooooo" изменено`)
							}
						}
						if statuses[len(statuses)-1] != "ooooooo" && len(messages) == 0 {
							messages = append(messages, `поле "oooooooo" изменено`)
						}
					}
					break
				}
			})
			s1.Find("#cbi-table-1-temp2").Each(func(arg1 int, s2 *goquery.Selection) {
				text := s2.Text()
				if text != "" {
					temperatura, errT := strconv.Atoi(text)
					temps = append(temps, temperatura)
					if errT != nil {
						u.LogError("[ERR] [температура не является числом : " + errT.Error() + "](fg-red)")
					} else if temperatura > config.Temperature {
						messages = append(messages, `температура превышена (`+text+")")
					}
				}
			})
		})

	} else {
		u.LogError("[ERR] [не удалось открыть url2 : " + err.Error() + "](fg-red)")
	}

	u.SetCharts(temps)

	if len(messages) != 0 {
		for _, message := range messages {
			// go sendMessageByEmail(u, message)
			u.NotifyServ.Notify(u, message)
			u.LogError("[ERR] [" + message + "](fg-red)")
		}
	}

}
