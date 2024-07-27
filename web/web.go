package web

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ben-smyth/bensmythme/api/blog"
	"github.com/ben-smyth/bensmythme/internal/spec"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type App struct {
	Content         spec.WebsiteSpec
	Port            string
	CustomAssetPath string
	Dev             bool
}

func ServeWebsite(dev bool, relativeStaticLocation string, app App) error {
	// HTTP
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir(relativeStaticLocation))

	// hot reloading for dev environment
	if dev {
		r.HandleFunc("/dev", handleWebSocket)
	}

	r.PathPrefix("/static/").Handler(cacheControlMiddleware(http.StripPrefix("/static/", fs)))
	r.HandleFunc("/blog/posts", blog.GetBlogPosts)
	r.HandleFunc("/", app.IndexHandler)

	return http.ListenAndServe(fmt.Sprintf(":%s", app.Port), r)
}

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

// cacheControlMiddleware - prevent caching static files for too long
func cacheControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "max-age=7200, must-revalidate")

		next.ServeHTTP(w, r)
	})
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
