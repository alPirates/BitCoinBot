package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/alPirates/BitCoinBot/assets"
)

// Config structure
type Config struct {
	Proxy       string `json:"proxy"`
	Port        string `json:"port"`
	Temperature int    `json:"max_temperature"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Token       string `json:"token"`
	UpdateTime  int    `json:"update_time"`
	HTMLURL1    string `json:"html_url_1"`
	HTMLURL2    string `json:"html_url_2"`
	ChatID      int64  `json:"chat_id"`
}

func createConfigFile() ([]byte, error) {
	log.Println("Reading config from default")
	data, err := assets.Asset("data/config.json")
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile("config.json", data, 0644)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetConfig() Config {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		file, err = createConfigFile()
		if err != nil {
			sendError("can't read config : " + err.Error())
			os.Exit(0)
		}
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

func sendError(err string) {
	fmt.Println(time.Now().Format(time.UnixDate) + " : " + err)
}
