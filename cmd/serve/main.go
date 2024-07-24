package main

import (
	"ben/gohtmx/internal/spec"
	"ben/gohtmx/web"
	"fmt"
	"log"
	"os"
)

func main() {
	websitePath := os.Getenv("APP_URL")
	if websitePath == "" {
		websitePath = "http://localhost:8080"
	}

	staticPath := websitePath + "/static/"

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

	websiteSpec, err := spec.LoadWebsiteSpec(websitePath, "spec.yaml")
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	app := web.App{
		Content:         websiteSpec,
		CustomAssetPath: staticPath,
		Dev:             dev,
		Port:            port,
	}
	fmt.Printf("Dev: %v\n", dev)
	fmt.Printf("Title: %s\n", app.Content.Title)
	fmt.Printf("Static Path: %s\n", staticPath)

	log.Fatal(web.ServeWebsite(dev, "web/static/", app))
}
