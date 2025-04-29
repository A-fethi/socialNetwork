package posts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-net/db"
	"social-net/session"
	"time"
)

type Posts struct {
	Author        string `json:"author"`
	Content       string `json:"content"`
	Categories    string `json:"categories"`
	Title         string `json:"title"`
	Creation_date string `json:"creation_date"`
}

func Post(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")        // Allow frontend origin
	w.Header().Set("Access-Control-Allow-Credentials", "true")                    // Allow cookies
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")          // Allow methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Allow headers

	// Handle OPTIONS preflight request
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Handle POST request
	if r.Method == "POST" {
		tokene, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Unauthorized: Missing token", http.StatusUnauthorized)
			return
		}

		token := tokene.Value

		var post Posts
		err = json.NewDecoder(r.Body).Decode(&post)
		fmt.Println(post.Categories)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if post.Title == "" || post.Content == "" || post.Categories == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		// Get user info from the token
		userid, _ := session.GetUserIDFromToken(token)
		if err != nil || userid == 0 {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		author, _ := session.Getusernamefromuserid(userid)
		if err != nil || author == "" {
			http.Error(w, "Unauthorized: User not found", http.StatusUnauthorized)
			return
		}

		// Insert post into the database
		_, err = db.DB.Exec("INSERT INTO posts (title, content, categories, user_id, author, creation_date) VALUES (?, ?, ?, ?, ?, ?)",
			post.Title, post.Content, post.Categories, userid, author, time.Now())
		if err != nil {
			http.Error(w, fmt.Sprintf("Error inserting post: %v", err), http.StatusInternalServerError)
			return
		}

		// Return success message
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Post created successfully"})
	}
}
