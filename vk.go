package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type VkService struct {
	Notifier
	link string
}

func NewVkService() VkService {
	vk := VkService{
		link: "https://api.vk.com/method/",
	}
	return vk
}

func (vk VkService) Notify(ui *UiService, message string) error {
	rand.Seed(time.Now().Unix())
	resp, err := http.Get(vk.link + "messages.send?group_id=" + fmt.Sprint(config.GroupId) + "&user_id=" + fmt.Sprint(config.UserId) + "&random_id=" + fmt.Sprint(rand.Int63()) + "&message=" + url.PathEscape(message) + "&access_token=" + config.Token + "&v=5.92")
	if err != nil {
		ui.LogError(err.Error())
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ui.LogError(err.Error())
		return err
	}
	if strings.Contains(string(body), "error") {
		ui.LogError(string(body))       // TODO: error
		return errors.New(string(body)) // TODO: error
	}
	return nil
}
