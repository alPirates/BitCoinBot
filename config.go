package main

import (
	"github.com/BurntSushi/toml"
)

// Config structure
type Config struct {
	Temperature int    `toml:"Temperature"`
	Email       string `toml:"Email"`
	Password    string `toml:"Password"`
	UpdateTime  int    `toml:"UpdateTime"`
	HTMLURL1    string `toml:"Url1"`
	HTMLURL2    string `toml:"Url2"`
}

func (c *Config) getConfig(ui *UiService) {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		ui.LogError("[CONF_ERR](fg-red) [Не удается прочитать файл config.toml](fg-yellow)")
	}
}
