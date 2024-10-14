package routes

import (
	"encoding/json"
	"net/http"
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

	if r.Method == "GET" {
		api.getAllPosts()
	}
}

// creating new blog post
func (api ApiRoute) createPost() {

}

// getting all the posts
func (api ApiRoute) getAllPosts() {
	api.w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(api.w).Encode("hello world")
}

// updating posts
func (api ApiRoute) updatePost() {}

// deleting post
func (api ApiRoute) deletePost() {}
