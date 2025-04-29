package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"social-net/db"
	"social-net/session"
)

type User struct {
	Username  string
	Email     string
	Password  string
	FirstName string
	LastName  string
	Bio       string
	Date      string
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081") // your frontend origin
	w.Header().Set("Access-Control-Allow-Credentials", "true")             // important for cookies
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")   // include all used methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")         // accept JSON headers, etc.

	log.Println(w.Header())
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {

		var user User
		json.NewDecoder(r.Body).Decode(&user)

		if user.Username != "" && user.Password != "" && user.FirstName == "" {
			fmt.Println("user.Username", user.Username)
			fmt.Println("user.Password", user.Password)
			var pass string
			err := db.DB.QueryRow("SELECT password FROM users WHERE username = ?", user.Username).Scan(&pass)
			fmt.Println("pass", pass)
			if !Validate(pass, user.Password) {
				fmt.Println("error", err)
				Senddata(w, 2, "Invalid password", "Error Password")
				return
			}
			if err != nil {

				if err == sql.ErrNoRows {
					http.Error(w, "Invalid username or password", http.StatusUnauthorized)
				} else {
					http.Error(w, "Database error", http.StatusInternalServerError)
				}
				Senddata(w, 0, "Invalid password", "Error Informaation")
				return
			}

			user_id, _ := session.GetUserIDFromUsername(user.Username)
			session.Setsession(w, r, user_id)
			Senddata(w, 0, "Login Success", "Success")

		} else if user.Username != "" && user.Email != "" && user.Password != "" && user.Date != "" && user.FirstName != "" && user.LastName != "" {
		
			newpss := Hashpwd(user.Password)
			_, err := db.DB.Exec("INSERT INTO users (username, email,password, first_name, last_name, date_of_birth,bio) VALUES (?, ?, ?, ?, ?, ?, ?)",
				user.Username, user.Email, newpss, user.FirstName, user.LastName, user.Date, user.Bio)
			if err != nil {
				http.Error(w, "Database error", http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode("register successful")
		} else {
			fmt.Println("ss")
			http.Error(w, "Invalid input", http.StatusBadRequest)
		}

	}
}
