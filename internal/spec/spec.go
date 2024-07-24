package spec

import (
	"log"
	"os"
	"regexp"

	"gopkg.in/yaml.v2"
)

type WebsiteSpec struct {
	Title        string         `yaml:"title"`
	Resume       Resume         `yaml:"resume"`
	Experience   []Experience   `yaml:"experience"`
	Portfolio    []Portfolio    `yaml:"portfolio"`
	Skills       []Skills       `yaml:"skills"`
	Testimonials []Testimonials `yaml:"testimonials"`
}

type Resume struct {
	Name        string `yaml:"name"`
	Tagline     string `yaml:"tagline"`
	ImageURL    string `yaml:"image_url"`
	GithubURL   string `yaml:"github_url"`
	LinkedinURL string `yaml:"linkedin_url"`
}

type Experience struct {
	CompanyName string `yaml:"company_name"`
	JobTitle    string `yaml:"job_title"`
	Date        string `yaml:"date"`
	ImageURL    string `yaml:"image_url"`
	Details     string `yaml:"details"`
	Position    string `yaml:"position"`
}

type Portfolio struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	BlogURL     string `yaml:"blog_url"`
	ImageURL    string `yaml:"image_url"`
	GitURL      string `yaml:"git_url"`
}

type Skills struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	ImageURL    string `yaml:"image_url"`
}

type Testimonials struct {
	From        string `yaml:"from"`
	Testimonial string `yaml:"testimonial"`
	ImageURL    string `yaml:"image_url"`
}

func LoadWebsiteSpec(siteUrl, yamlFile string) (WebsiteSpec, error) {
	// Load the YAML file
	data, err := os.ReadFile(yamlFile)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Perform the regex replacement
	re := regexp.MustCompile(`\$MYSITEURL\$`)
	updatedData := re.ReplaceAllString(string(data), siteUrl) // Load the YAML file
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var spec WebsiteSpec
	err = yaml.Unmarshal([]byte(updatedData), &spec)
	if err != nil {
		return spec, err
	}

	return spec, nil
}
