package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

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

type SinglePost struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
	_, err = db.DB.Query("INSERT INTO posts(title, content, tags) values ($1, $2, $3) RETURNING id", post.Title, post.Content, pq.Array(post.Tags))

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
	writer := api.w
	writer.Header().Set("Content-type", "application/json")
	rows, err := db.DB.Query("SELECT id, title, content, tags, created_at, updated_at FROM posts")

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(err)
		return
	}

	var posts []SinglePost

	for rows.Next() {
		var post SinglePost

		err := rows.Scan(&post.Id, &post.Title, &post.Content, pq.Array(&post.Tags), &post.CreatedAt, &post.UpdatedAt)

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			panic(err)
		}

		posts = append(posts, post)
	}

	json.NewEncoder(api.w).Encode(posts)
}

// updating posts
func (api ApiRoute) updatePost() {}

// deleting post
func (api ApiRoute) deletePost() {}
