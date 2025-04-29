package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"social-net/db"
	"social-net/session"
)

type Info struct {
	Username  string
	Email     string
	Firstname string
	Lastname  string
	Date      string
	Bio       string
	Password  string
}

func Getinfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081") // your frontend origin
	w.Header().Set("Access-Control-Allow-Credentials", "true")             // important for cookies
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")   // include all used methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")         // accept JSON headers, etc.
	// Get the token from the cookie
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Unauthorized: Missing token", http.StatusUnauthorized)
		return
	}

	token := cookie.Value
	// Get the user ID from the token
	userID, _ := session.GetUserIDFromToken(token)
	// Get the username from the user ID
	username, _ := session.Getusernamefromuserid(userID)
	// Declare a variable to hold the user information
	var info Info
	fmt.Println("username", username)
	// Query the database to get the user's information
	err = db.DB.QueryRow("SELECT username, email, first_name, last_name, date_of_birth,bio FROM users WHERE username=?", username).Scan(&info.Username, &info.Email, &info.Firstname, &info.Lastname, &info.Date, &info.Bio)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("error", err)
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			fmt.Println("error", err)
			http.Error(w, "Error retrieving user information", http.StatusInternalServerError)
		}
		return
	}
	fmt.Println("info", info)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}
