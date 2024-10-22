package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type ApiRoute struct {
	w http.ResponseWriter
	r *http.Request
}

func HandleRoutes(w http.ResponseWriter, r *http.Request) {
	api := ApiRoute{
		w: w,
		r: r,
	}

	switch r.Method {
	case http.MethodGet:
		api.getAllPosts()
	case http.MethodPost:
		api.createPost()
	case http.MethodPut:
		api.updatePost()
	case http.MethodDelete:
		api.deletePost()
	}
}

type Post struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

// creating new blog post
func (api ApiRoute) createPost() {
	req := api.r
	writer := api.w

	writer.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(req.Body)
	var post Post
	err := decoder.Decode(&post)

	if err != nil {
		log.Fatal("Error in decoding the request body")
	}

	post.Content = strings.TrimSpace(post.Content)
	post.Title = strings.TrimSpace(post.Title)

	if post.Content == "" || post.Title == "" {
		err := map[string]string{
			"error": `Either "content" or "title" property is missing from the body.`,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(err)
		return
	}

	fmt.Println(post)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(post)

}

// getting all the posts
func (api ApiRoute) getAllPosts() {
	api.w.Header().Set("Content-Type", "application/json")
	api.w.WriteHeader(http.StatusCreated)
	res := map[string]string{
		"msg":   "hello world",
		"value": "10",
	}
	json.NewEncoder(api.w).Encode(res)
}

// updating posts
func (api ApiRoute) updatePost() {}

// deleting post
func (api ApiRoute) deletePost() {}
