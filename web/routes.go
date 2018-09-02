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

	fmt.Println(c.RawJSON)

	err = c.Request.ParseForm()
	fmt.Println(err)
	fmt.Println(c.Request.FormValue("email"))
	fmt.Println(c.Request.Form)

	email := c.Request.FormValue("email")
	password := c.Request.Form.Get("password")
	HTMLURL1 := c.Request.Form.Get("html_url_1")
	HTMLURL2 := c.Request.Form.Get("html_url_2")
	updateTime := c.Request.Form.Get("update_time")
	port := c.Request.Form.Get("port")
	proxy := c.Request.Form.Get("proxy")
	temperature := c.Request.Form.Get("max_temperature")

	fmt.Println(email, password, HTMLURL1, HTMLURL2, updateTime, port, proxy, temperature)

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
