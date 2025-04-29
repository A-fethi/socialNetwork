package comments

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-net/db"
	"social-net/session"
	"strconv"
	"time"
)

type Comments struct {
	PostId        string
	Comment       string
	Author        string
	Creation_date time.Time
}

func AddComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")        // Allow frontend origin
	w.Header().Set("Access-Control-Allow-Credentials", "true")                    // Allow cookies
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")          // Allow methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Allow headers
	w.WriteHeader(http.StatusOK)                                                  // Respond with 200 OK for preflight request
	if r.Method == "POST" {
		var comment Comments

		// Decode the incoming JSON data into the `comment` struct
		err := json.NewDecoder(r.Body).Decode(&comment)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			http.Error(w, "Failed to decode comment JSON", http.StatusBadRequest)
			return
		}
		token, _ := r.Cookie("token")
		tokenn := token.Value
		userid, _ := session.GetUserIDFromToken(tokenn)
		username, _ := session.Getusernamefromuserid(userid)

		if userid == 0 || username == "" || tokenn == "" {
			return
		}

		// Insert the comment into the database
		_, err = db.DB.Exec("INSERT INTO comments (post_id, author, content, creation_date) VALUES (?, ?, ?, ?)",
			comment.PostId, username, comment.Comment, time.Now())
		// Handle database insertion errors
		if err != nil {
			fmt.Println("Error inserting comment into database:", err)
			http.Error(w, "Failed to insert comment", http.StatusInternalServerError)
			return
		}

		// Respond with a success message
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Comment created successfully",
		})
	} else {
		// Handle unsupported request methods (e.g., GET, PUT, DELETE)
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func Getcomments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")        // Allow frontend origin
	w.Header().Set("Access-Control-Allow-Credentials", "true")                    // Allow cookies
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")          // Allow methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Allow headers
	qu := r.URL.Query()
	postidd := qu.Get("postId")
	fmt.Println("iddddddd", postidd)
	postid, _ := strconv.Atoi(postidd)
	if postid == 0 {
		http.Error(w, "Missing postid parameter", http.StatusBadRequest)
		return
	}
	rows, err := db.DB.Query("SELECT post_id, author,content,creation_date FROM comments WHERE post_id = ? ORDER BY creation_date DESC", postid)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var comments []Comments
	for rows.Next() {
		var comment Comments
		err := rows.Scan(&comment.PostId, &comment.Author, &comment.Comment, &comment.Creation_date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		comments = append(comments, comment)

	}
	json.NewEncoder(w).Encode(comments)
}
