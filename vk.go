package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

type VkService struct {
	Notifier
	link     string
	group_id string
	user_id  string
	token    string
}

func NewVkService(group_id, user_id int, token string) VkService {
	vk := VkService{
		link:     "https://api.vk.com/method/",
		group_id: fmt.Sprint(group_id),
		user_id:  fmt.Sprint(user_id),
		token:    token,
	}
	return vk
}

func (vk VkService) Send(message string) error {
	rand.Seed(time.Now().Unix())
	resp, err := http.Get(vk.link + "messages.send?group_id=" + vk.group_id + "&user_id=" + vk.user_id + "&random_id=" + fmt.Sprint(rand.Int63()) + "&message=" + url.PathEscape(message) + "&access_token=" + vk.token + "&v=5.92")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
