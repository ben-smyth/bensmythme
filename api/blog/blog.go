package blog

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"text/template"
)

var (
	blogSpotApiKey string
	client         http.Client
	blogId         int
)

func InitBlogSettings(bsApiKey string, blogID int) {
	blogSpotApiKey = bsApiKey
	blogId = blogID
	client = *http.DefaultClient
}

// serve blog posts in list
func GetBlogPosts(w http.ResponseWriter, r *http.Request) {
	blogs, err := GetBlogPostsFromBlogspot()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/html")

	var tmpl = template.Must(template.ParseFiles("web/templates/components/blog_post.html"))

	err = tmpl.Execute(w, blogs)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

func GetBlogPostsFromBlogspot() (BlogPostList, error) {
	var PostList BlogPostList

	resp, err := client.Get(fmt.Sprintf("https://www.googleapis.com/blogger/v3/blogs/%v/posts?key=%v", blogId, blogSpotApiKey))
	if err != nil {
		return PostList, fmt.Errorf("Failed to get blogs from Blogger: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PostList, fmt.Errorf("Failed to read response body: %v", err)
	}
	err = json.Unmarshal(body, &PostList)
	if err != nil {
		return PostList, fmt.Errorf("Failed to unmarshall Blog Post JSON: %v", err)
	}

	return PostList, nil
}
