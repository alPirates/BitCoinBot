package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func loopParse(u *UiService, config *Config) {
    for true {
        startTime := time.Now()
        parse(u, config)
        time.Sleep(time.Duration(config.UpdateTime)*time.Minute - time.Now().Sub(startTime))
    }
}

func parse(u *UiService, config *Config) {
		messages := []string{}

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
							u.LogError("[ERR] [температура не является числом : " + err.Error() + "](fg-red)")
						} else if temperatura > config.Temperature {
							messages = append(messages, `температура превышена (`+text+")")
						}
					}
				})
			})

		} else {
            u.LogError("[ERR] [can't open your website : " + err.Error() + "](fg-red)")
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
            u.LogError("[ERR] [can't open your website : " + err.Error() + "](fg-red)")
		}

		if len(messages) != 0 {
            for _, message := range messages {
                go SendMessagesByMail(*config, message)
            }
		}
}
