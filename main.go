package main

import (
	"fmt"
	"time"

	"github.com/alPirates/BitCoinBot/config"
	"github.com/alPirates/BitCoinBot/web"
)

func main() {
	sendError("reading config...")
	config := config.GetConfig()
	sendError("starting bot...")
	bot := startBot(config)
	go web.StartServer(config.Port)
	go parse(bot, config)
	startBotWork(bot, config)
}

func sendError(err string) {
	fmt.Println(time.Now().Format(time.UnixDate) + " : " + err)
}
