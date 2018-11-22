package main

import (
	"encoding/json"
	"io/ioutil"
    "fmt"
	"os"
)

// Config structure
type Config struct {
	Temperature int    `json:"max_temperature"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	UpdateTime  int    `json:"update_time"`
	HTMLURL1    string `json:"html_url_1"`
	HTMLURL2    string `json:"html_url_2"`
}

func getConfig() *Config {
	file, err := ioutil.ReadFile("data/config.json")
	if err != nil {
        fmt.Println(err)
		os.Exit(0)
	}

	config := &Config{}
	err = json.Unmarshal(file, config)
	if err != nil {
        fmt.Println("kek2")
		os.Exit(0)
	}

	return config
}
