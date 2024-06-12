package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type App struct {
	Title           string
	CustomAssetPath string
}

func main() {
	staticPath := os.Getenv("APP_STATICPATH")
	if staticPath == "" {
		staticPath = "http://localhost/static/"
	}

	title := os.Getenv("APP_TITLE")
	if staticPath == "" {
		staticPath = "bensmyth"
	}

	// port assigned by heroku
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Title: %s\n", title)
	fmt.Printf("Static Path: %s\n", staticPath)

	devApp := &App{
		Title:           title,
		CustomAssetPath: staticPath,
	}

	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", devApp.IndexHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func (a *App) IndexHandler(w http.ResponseWriter, r *http.Request) {
	Templ.ExecuteTemplate(w, "home", a)
}

var Templ = func() *template.Template {
	t := template.New("")
	err := filepath.Walk("web/templates", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			fmt.Println(path)
			_, err = t.ParseFiles(path)
			if err != nil {
				fmt.Println(err)
			}
		}
		return err
	})

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return t
}()
