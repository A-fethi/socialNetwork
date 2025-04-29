package posts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-net/db"
)

type GetPost struct {
	Id            int
	User_id       int
	Author        string
	Content       string
	Categories    string
	Title         string
	Creation_date string
}

func Getposts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")        // Allow frontend origin
	w.Header().Set("Access-Control-Allow-Credentials", "true")                    // Allow cookies
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")          // Allow methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Allow headers
	rows, err := db.DB.Query("SELECT  id, author, content, categories,title, user_id, creation_date FROM posts ORDER BY creation_date DESC")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var posts []GetPost
	for rows.Next() {
		var post GetPost
		err := rows.Scan(&post.Id, &post.Author, &post.Content, &post.Categories, &post.Title, &post.User_id, &post.Creation_date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		posts = append(posts, post)

	}
	json.NewEncoder(w).Encode(posts)
}
