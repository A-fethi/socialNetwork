package session

import (
	"encoding/json"
	"net/http"

	"social-net/db"
)

func Middleware(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081") // your frontend origin
	w.Header().Set("Access-Control-Allow-Credentials", "true")             // important for cookies
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")   // include all used methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")         // accept JSON headers, etc.
	re, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var sessionCount int
	err = db.DB.QueryRow("SELECT COUNT(*) FROM sessions WHERE token=?", re.Value).Scan(&sessionCount)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Database query failed",
		})
		http.Error(w, "Database query failed", http.StatusInternalServerError)
		return
	}

	// Check if the token exists in the session table
	if sessionCount > 0 {
		// Token is valid, return success
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Login successful",
		})
		return
	}

	// Token not found in the database
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login failed",
	})
}
