package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ben-smyth/bensmythme/api/blog"
	"github.com/ben-smyth/bensmythme/internal/spec"
	"github.com/ben-smyth/bensmythme/web"
)

func main() {
	// ENVIRONMENT Configuration
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// WEBSITE SPEC Configuration Unmarshall
	websiteSpec, err := spec.LoadWebsiteSpec(websitePath, "spec.yaml")
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	// BLOG  Configuration
	if websiteSpec.SectionSelection.Blog {
		blogspotApiKey := os.Getenv(websiteSpec.Integrations.Blogger.APIKeyEnvVar)
		if blogspotApiKey == "" {
			fmt.Println(fmt.Errorf("blogspot api key not set. Env var: %v", websiteSpec.Integrations.Blogger.APIKeyEnvVar))
			return
		}
		blog.InitBlogSettings(blogspotApiKey, websiteSpec.Integrations.Blogger.BlogId)
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

	// START Web Server
	log.Fatalln(web.ServeWebsite(dev, "web/static/", app))
}
