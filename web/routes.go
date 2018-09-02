package web

import (
	"fmt"
	"net/http"

	"github.com/alPirates/BitCoinBot/assets"
	magic "github.com/alPirates/Magic"
)

func mainHandler(c *magic.Context) error {
	err := mainTemplate.Execute(c.Writer, serverConf)
	if err != nil {
		return err
	}
	return nil
}

func updateHandler(c *magic.Context) error {

	email, _ := c.PostParams.ParseString("email")
	password, _ := c.PostParams.ParseString("password")
	HTMLURL1, _ := c.PostParams.ParseString("html_url_1")
	HTMLURL2, _ := c.PostParams.ParseString("html_url_2")
	updateTime, _ := c.PostParams.ParseInt("update_time")
	port, _ := c.PostParams.ParseString("port")
	proxy, _ := c.PostParams.ParseString("proxy")
	temperature, _ := c.PostParams.ParseInt("max_temperature")

	fmt.Println(email)
	fmt.Println(password)
	fmt.Println(HTMLURL1)
	fmt.Println(HTMLURL2)
	fmt.Println(updateTime)
	fmt.Println(port)
	fmt.Println(proxy)
	fmt.Println(temperature)

	http.Redirect(c.Writer, c.Request, "/", 301)
	return nil
}

func staticHandler(c *magic.Context) error {
	filename := c.Storage["fileName"].(string)
	file, err := assets.Asset(fmt.Sprintf("data/web/static/%s", filename))
	if err != nil {
		return c.SendError(err)
	}
	str, _ := c.Headers.ParseString("Accept")
	c.Writer.Header().Set("Content-Type", str)
	return c.SendString(string(file))
}
