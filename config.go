package main

import (
	"encoding/json"
	"io/ioutil"
    "fmt"
	"os"
    "reflect"
    "strconv"
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

func (config *Config)toStringMas() []string {
    mas := make([]string, 0)
    v := reflect.ValueOf(*config)
    n := v.Type().NumField()
    for i := 0; i < n; i++ {
        g := v.Type().Field(i)
        switch v.Field(i).Kind() {
        case reflect.String:
            mas = append(mas, "[" + g.Name + "] [" + v.Field(i).String() + "](fg-yellow)")
            break
        case reflect.Int:
            mas = append(mas, "[" + g.Name + "] [" + strconv.Itoa((int)(v.Field(i).Int()))+"](fg-yellow)")
            break
        }
    }
    return mas
}
