package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type App struct {
	Title           string
	CustomAssetPath string
	Dev             bool
}

func main() {
	staticPath := os.Getenv("APP_STATICPATH")
	if staticPath == "" {
		staticPath = "http://localhost:8080/static/"
	}

	title := os.Getenv("APP_TITLE")
	if title == "" {
		title = "Ben Smyth"
	}

	dev := false

	envfromenv := os.Getenv("APP_DEV")
	if envfromenv == "" {
		dev = true
	}

	// port assigned by heroku
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Title: %s\n", title)
	fmt.Printf("Static Path: %s\n", staticPath)
	fmt.Printf("Dev: %v\n", dev)

	devApp := &App{
		Title:           title,
		CustomAssetPath: staticPath,
		Dev:             dev,
	}

	// HTTP
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("web/static"))

	// hot reloading for dev environment
	if dev {
		r.HandleFunc("/dev", handleWebSocket)
	}

	r.PathPrefix("/static/").Handler(cacheControlMiddleware(http.StripPrefix("/static/", fs)))
	r.HandleFunc("/", devApp.IndexHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade failed:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read failed:", err)
			break
		}

		fmt.Printf("Received: %s\n", message)

		if err := conn.WriteMessage(messageType, message); err != nil {
			fmt.Println("Write failed:", err)
			break
		}
	}
}

func cacheControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "max-age=14400, must-revalidate")

		next.ServeHTTP(w, r)
	})
}
