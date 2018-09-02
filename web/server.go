package web

import (
	"html/template"

	"github.com/alPirates/BitCoinBot/assets"
	"github.com/alPirates/BitCoinBot/config"
	magic "github.com/alPirates/Magic"
)

var (
	serverConf   config.Config
	mainTemplate *template.Template
)

func install() error {
	var (
		err   error
		index []byte
	)

	index, err = assets.Asset("data/web/index.html")
	if err != nil {
		return err
	}

	mainTemplate, err = template.New("index.html").Parse(string(index))
	if err != nil {
		return err
	}

	return nil
}

func StartServer(port string) {
	serverConf = config.GetConfig()

	server := magic.NewMagic(port)
	err := install()
	if err != nil {
		panic(err)
	}

	server.GET("/", mainHandler)
	server.CUSTOM("/static", "STATIC", staticHandler)

	server.ListenAndServe()

}
