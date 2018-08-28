package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config structure
type Config struct {
	Proxy       string `json:"proxy"`
	Temperature int    `json:"max_temperature"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Token       string `json:"token"`
	UpdateTime  int    `json:"update_time"`
	HTMLURL1    string `json:"html_url_1"`
	HTMLURL2    string `json:"html_url_2"`
	ChatID      int64  `json:"chat_id"`
}

func getConfig() Config {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		sendError("can't read config : " + err.Error())
		os.Exit(0)
	}

	config := &Config{}
	err = json.Unmarshal(file, config)
	if err != nil {
		sendError("can't unmarshal config : " + err.Error())
		os.Exit(0)
	}

	sendError("read config")

	return *config
}
