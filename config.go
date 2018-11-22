package main

import (
	"reflect"
	"strconv"

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

func (config *Config) toStringMas() []string {
	mas := make([]string, 0)
	v := reflect.ValueOf(*config)
	n := v.Type().NumField()
	for i := 0; i < n; i++ {
		g := v.Type().Field(i)
		switch v.Field(i).Kind() {
		case reflect.String:
			mas = append(mas, "["+g.Name+"] ["+v.Field(i).String()+"](fg-yellow)")
			break
		case reflect.Int:
			mas = append(mas, "["+g.Name+"] ["+strconv.Itoa((int)(v.Field(i).Int()))+"](fg-yellow)")
			break
		}
	}
	return mas
}
