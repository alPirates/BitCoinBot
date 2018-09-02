package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	err := mainTemplate.Execute(w, serverConf)
	if err != nil {
		log.Println(err)
	}
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("erro rparsing")
	}
	fmt.Println("email is ", r.FormValue("email"))
	http.Redirect(w, r, "/", 301)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	filename := chi.URLParam(r, "filename")
	fmt.Println("filename", filename)

	// file, err := assets.Asset(fmt.Sprintf("data/web/static/%s", filename))
	// if err != nil {
	// 	return c.SendError(err)
	// }
}
