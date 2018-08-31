package main

import (
	"fmt"
	"time"
)

func main() {
	sendError("reading config...")
	config := getConfig()
	sendError("starting bot...")
	bot := startBot(config)
	go parse(bot, config)
	startBotWork(bot, config)
}

func sendError(err string) {
	fmt.Println(time.Now().Format(time.UnixDate) + " : " + err)
}
