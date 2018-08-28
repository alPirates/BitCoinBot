package main

import (
	"fmt"
	"time"

	magicFramework "github.com/Gittoks/Magic/project"
)

func main() {
	go upHTML()
	sendError("reading config...")
	config := getConfig()
	sendError("starting bot...")
	bot := startBot(config)
	go parse(bot, config)
	startBotWork(bot, config)
}

func upHTML() {
	magic := magicFramework.NewMagic("8081")
	magic.FILE("/file1", "html1.html")
	magic.FILE("/file2", "html2.html")
	magic.ListenAndServe()
}

func sendError(err string) {
	fmt.Println(time.Now().Format(time.UnixDate) + " : " + err)
}
