package blog

import (
	"fmt"
	"net/http"
)

// serve blog posts in list

func GetBlogPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "test")
}
