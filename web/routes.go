package web

import (
	"fmt"

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
