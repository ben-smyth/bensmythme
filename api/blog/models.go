package blog

type BlogPostList struct {
	Kind  string     `json:"kind"`
	Items []BlogPost `json:"items"`
	Etag  string     `json:"etag"`
}

type BlogPost struct {
	Kind      string  `json:"kind"`
	ID        string  `json:"id"`
	Blog      Blog    `json:"blog"`
	Published string  `json:"published"`
	Updated   string  `json:"updated"`
	URL       string  `json:"url"`
	SelfLink  string  `json:"selfLink"`
	Title     string  `json:"title"`
	Content   string  `json:"content"`
	Author    Author  `json:"author"`
	Replies   Replies `json:"replies"`
	Etag      string  `json:"etag"`
}

type Blog struct {
	ID string `json:"id"`
}

type Author struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	URL         string `json:"url"`
	Image       Image  `json:"image"`
}

type Image struct {
	URL string `json:"url"`
}

type Replies struct {
	TotalItems string `json:"totalItems"`
	SelfLink   string `json:"selfLink"`
}
