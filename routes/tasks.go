package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/lib/pq"
	"github.com/tirlochanarora16/blogging_api_go/db"
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

	// insert the post into the DB
	var postId int
	// sqlQuery := fmt.Sprintf("INSERT INTO posts(title, content, tags) values(%s, %s, %s) RETURNING id", post.Title, post.Content, post.Tags)
	rows, err := db.DB.Query("INSERT INTO posts(title, content, tags) values ($1, $2, $3) RETURNING id", post.Title, post.Content, pq.Array(post.Tags))

	fmt.Println(rows)

	if err != nil {
		errorMsg = map[string]string{
			"error": fmt.Sprintf("Error inserting post into DB. %s", err),
		}
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(errorMsg)
		return
	}

	successResponse := map[string]any{
		"id":   postId,
		"post": post,
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(successResponse)

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
