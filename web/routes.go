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
	fmt.Println("filename is ", filename)
	if err != nil {
		return c.SendError(err)
	}
	fmt.Println("returning", file)
	return c.SendString(string(file))
}
