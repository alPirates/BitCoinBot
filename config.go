package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config structure
type Config struct {
	Temperature int
	Email       string
	Password    string
	UpdateTime  int
	HTMLURL1    string
	HTMLURL2    string
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
