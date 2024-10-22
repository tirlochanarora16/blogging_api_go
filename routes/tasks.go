package routes

import (
	"encoding/json"
	"fmt"
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

	var errorMsg map[string]string = make(map[string]string)

	writer.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(req.Body)
	var post Post
	err := decoder.Decode(&post)

	if err != nil {
		errorMsg = map[string]string{
			"error": "Error in decoding the request body",
		}
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorMsg)
		return
	}

	post.Content = strings.TrimSpace(post.Content)
	post.Title = strings.TrimSpace(post.Title)

	if post.Content == "" || post.Title == "" {
		errorMsg = map[string]string{
			"error": `"content" or "title" cannot be empty.`,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(errorMsg)
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
