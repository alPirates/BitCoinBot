package main

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"strconv"

	"github.com/BurntSushi/toml"
)

// Config structure
type Config struct {
	Temperature int    `toml:"Temperature" conf:"TEMP"`
	Email       string `toml:"Email" conf:"MAIL"`
	Password    string `toml:"Password" conf:"PASS"`
	UpdateTime  int    `toml:"UpdateTime" conf:"TIME"`
	Token       string `toml:"Token" conf:"TOKN"`
	GroupId     int    `toml:"GroupId" conf:"GID"`
	UserId      int    `toml:"UserId" conf:"UID"`
	HTMLURL1    string `toml:"Url1" conf:"URL1"`
	HTMLURL2    string `toml:"Url2" conf:"URL2"`
}

func (c *Config) getConfig(ui *UiService) {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		ui.LogError("[CONF_ERR](fg-red) [Не удается прочитать файл config.toml](fg-yellow)")
		ui.SetStatus(false)
		return
	}
	if c.UpdateTime == 0 {
		c.UpdateTime = 99
		ui.LogError("[CONF] [Введите верный UpdateTime!](fg-red)")
		ui.SetStatus(false)
		return
	}
	ui.SetStatus(true)
}

func CreateConfig(ui *UiService) {
	ui.LogError("[MSG] Создаю конфигурационный файл")
	buf := new(bytes.Buffer)
	cfg := Config{
		80, "test@test.com", "password", 15, "vk_token", 0, 0, "url1", "url2",
	}
	if err := toml.NewEncoder(buf).Encode(cfg); err != nil {
		ui.LogError(err.Error())
		return
	}
	err := ioutil.WriteFile("config.toml", buf.Bytes(), 0644)
	if err != nil {
		ui.LogError("[CONF] [Не удается создать файл config.toml](fg-red)")
		return
	}
	ui.LogError("[MSG] [Конфигурационный файл создан успешно](fg-green)")
	config.getConfig(ui)
}

func (config *Config) toStringMas() []string {
	mas := make([]string, 0)
	v := reflect.ValueOf(*config)
	n := v.Type().NumField()
	for i := 0; i < n; i++ {
		g := v.Type().Field(i)
		switch v.Field(i).Kind() {
		case reflect.String:
			mas = append(mas, "["+g.Tag.Get("conf")+"] ["+v.Field(i).String()+"](fg-yellow)")
			break
		case reflect.Int:
			mas = append(mas, "["+g.Tag.Get("conf")+"] ["+strconv.Itoa((int)(v.Field(i).Int()))+"](fg-yellow)")
			break
		}
	}
	return mas
}
