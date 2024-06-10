package main

import (
	"ben/gohtmx/web"
	"log"
	"net/http"
)

type App struct {
	Title string
	CustomAssetPath string
}

func main() {
	devApp := &App{
		Title: "bensmyth",
		CustomAssetPath: "http://localhost:8080/static/",
	}

	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", devApp.IndexHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (a *App) IndexHandler(w http.ResponseWriter, r *http.Request) {
	web.Templ.ExecuteTemplate(w, "home", a)
}
