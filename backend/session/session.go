package session

import (
	"errors"
	"net/http"
	"sync"
	"time"

	"social-net/db"

	"github.com/gofrs/uuid"
)

var (
	sessions     = make(map[string]int)
	sessionMutex sync.RWMutex
)

// CreateSession creates a new session for a user
func CreateSession(userID int) (string, error) {
	// Generate a new UUID for the session token
	uuid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	token := uuid.String()

	// Store the session in memory
	sessionMutex.Lock()
	sessions[token] = userID
	sessionMutex.Unlock()

	// Store the session in the database for persistence
	database, err := db.GetDB()
	if err != nil {
		return "", err
	}

	_, err = database.Exec(
		"INSERT INTO sessions (token, user_id, created_at) VALUES (?, ?, datetime('now'))",
		token, userID,
	)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetUserIDFromToken retrieves the user ID associated with a session token
func GetUserIDFromToken(token string) (int, bool) {
	// First check in-memory cache
	sessionMutex.RLock()
	userID, exists := sessions[token]
	sessionMutex.RUnlock()

	if exists {
		return userID, true
	}

	// If not in memory, check the database
	database, err := db.GetDB()
	if err != nil {
		return 0, false
	}

	var id int
	err = database.QueryRow(
		"SELECT user_id FROM sessions WHERE token = ? AND created_at > datetime('now', '-30 day')",
		token,
	).Scan(&id)

	if err != nil {
		return 0, false
	}

	// Add to in-memory cache
	sessionMutex.Lock()
	sessions[token] = id
	sessionMutex.Unlock()

	return id, true
}

// GetUserID extracts the user ID from the request's cookie token
func GetUserID(r *http.Request) (int, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return 0, errors.New("unauthorized: no session token")
	}

	userID, ok := GetUserIDFromToken(cookie.Value)
	if !ok {
		return 0, errors.New("unauthorized: invalid session token")
	}

	return userID, nil
}

// DeleteSession removes a session
func DeleteSession(token string) error {
	// Remove from memory
	sessionMutex.Lock()
	delete(sessions, token)
	sessionMutex.Unlock()

	// Remove from database
	database, err := db.GetDB()
	if err != nil {
		return err
	}

	_, err = database.Exec("DELETE FROM sessions WHERE token = ?", token)
	return err
}

// SetSessionCookie sets the session cookie in the response
func SetSessionCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true, // Set to true in production with HTTPS
	})
}

// ClearSessionCookie clears the session cookie
func ClearSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
	})
}
