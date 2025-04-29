package auth

import (
	"encoding/json"
	"net/http"
	"social-net/session"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081") // your frontend origin
	w.Header().Set("Access-Control-Allow-Credentials", "true")             // important for cookies
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")   // include all used methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")         // accept JSON headers, etc.
	// Retrieve token from cookie
	cookie, err := r.Cookie("token")
	if err != nil {
		// If cookie is not found or error occurs, return unauthorized error
		http.Error(w, "Unauthorized - No token found", http.StatusUnauthorized)
		return
	}

	token := cookie.Value

	// Get userID from token
	userid, _ := session.GetUserIDFromToken(token)

	// Delete the session
	err = session.Deletesession(userid)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
		Path:    "/",
	})
	if err != nil {
		// If session deletion fails, return an error
		http.Error(w, "Failed to delete session", http.StatusInternalServerError)
		return
	}

	// Send a successful response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logged out successfully",
	})
}
