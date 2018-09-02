package web

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/alPirates/BitCoinBot/assets"
	"github.com/alPirates/BitCoinBot/config"
	"github.com/go-chi/chi"
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

	router := chi.NewRouter()
	err := install()
	if err != nil {
		panic(err)
	}

	router.Get("/", mainHandler)
	router.Post("/update", updateHandler)
	router.Route("/static", func(r chi.Router) {
		r.Get("/{filename}", staticHandler)
	})

	http.ListenAndServe(
		fmt.Sprintf(":%s", serverConf.Port),
		router,
	)

}
