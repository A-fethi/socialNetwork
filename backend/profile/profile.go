package profile

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"social-net/db"
	"social-net/session"
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Bio       string `json:"bio"`
	IsPrivate bool   `json:"is_private"`
}

type PrivacySettings struct {
	IsPrivate bool `json:"is_private"`
}

// GetProfile handles fetching a user's profile
func GetProfile(w http.ResponseWriter, r *http.Request) {
	// Check if user is authenticated
	userID, err := session.GetUserID(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Extract profile ID from URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 {
		http.Error(w, "Invalid profile ID", http.StatusBadRequest)
		return
	}

	profileID, err := strconv.Atoi(pathParts[3])
	if err != nil {
		http.Error(w, "Invalid profile ID", http.StatusBadRequest)
		return
	}

	// Get database connection
	database, err := db.GetDB()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Fetch user profile
	var user User
	err = database.QueryRow(`
		SELECT id, username, email, first_name, last_name, bio, is_private 
		FROM users 
		WHERE id = ?`, profileID).Scan(
		&user.ID, &user.Username, &user.Email, &user.Firstname, &user.Lastname, &user.Bio, &user.IsPrivate,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	// Check if profile is private and if current user is allowed to view it
	if user.IsPrivate && userID != profileID {
		// Check if current user is following this profile
		var isFollowing bool
		err = database.QueryRow(`
			SELECT EXISTS(
				SELECT 1 FROM followers 
				WHERE follower_id = ? AND following_id = ?
			)`, userID, profileID).Scan(&isFollowing)

		if err != nil || !isFollowing {
			// Return limited profile info
			limitedUser := struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Firstname string `json:"firstname"`
				Lastname  string `json:"lastname"`
				IsPrivate bool   `json:"is_private"`
			}{
				ID:        user.ID,
				Username:  user.Username,
				Firstname: user.Firstname,
				Lastname:  user.Lastname,
				IsPrivate: user.IsPrivate,
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(limitedUser)
			return
		}
	}

	// Return full profile info
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdatePrivacy handles updating a user's privacy settings
func UpdatePrivacy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check if user is authenticated
	userID, err := session.GetUserID(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Parse request body
	var settings PrivacySettings
	err = json.NewDecoder(r.Body).Decode(&settings)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Get database connection
	database, err := db.GetDB()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Update privacy settings
	_, err = database.Exec("UPDATE users SET is_private = ? WHERE id = ?", settings.IsPrivate, userID)
	if err != nil {
		http.Error(w, "Failed to update privacy settings", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}

// GetFollowers returns a list of users who follow the specified user
func GetFollowers(w http.ResponseWriter, r *http.Request) {
	// Check if user is authenticated
	currentUserID, err := session.GetUserID(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Extract user ID from URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var targetUserID int
	if pathParts[3] == "me" {
		targetUserID = currentUserID
	} else {
		var err error
		targetUserID, err = strconv.Atoi(pathParts[3])
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
	}

	// Get database connection
	database, err := db.GetDB()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Check if profile is private and if current user is allowed to view it
	if currentUserID != targetUserID {
		var isPrivate bool
		err = database.QueryRow("SELECT is_private FROM users WHERE id = ?", targetUserID).Scan(&isPrivate)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		if isPrivate {
			// Check if current user is following this profile
			var isFollowing bool
			err = database.QueryRow(`
				SELECT EXISTS(
					SELECT 1 FROM followers 
					WHERE follower_id = ? AND following_id = ?
				)`, currentUserID, targetUserID).Scan(&isFollowing)

			if err != nil || !isFollowing {
				http.Error(w, "Unauthorized to view this private profile", http.StatusForbidden)
				return
			}
		}
	}

	// Fetch followers
	rows, err := database.Query(`
		SELECT u.id, u.username, u.first_name, u.last_name
		FROM users u
		JOIN followers f ON u.id = f.follower_id
		WHERE f.following_id = ?
	`, targetUserID)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	followers := []struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	}{}

	for rows.Next() {
		var follower struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			Firstname string `json:"firstname"`
			Lastname  string `json:"lastname"`
		}
		err := rows.Scan(&follower.ID, &follower.Username, &follower.Firstname, &follower.Lastname)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		followers = append(followers, follower)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(followers)
}

// GetFollowing returns a list of users that the specified user follows
func GetFollowing(w http.ResponseWriter, r *http.Request) {
	// Check if user is authenticated
	currentUserID, err := session.GetUserID(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Extract user ID from URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var targetUserID int
	if pathParts[3] == "me" {
		targetUserID = currentUserID
	} else {
		var err error
		targetUserID, err = strconv.Atoi(pathParts[3])
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
	}

	// Get database connection
	database, err := db.GetDB()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Check if profile is private and if current user is allowed to view it
	if currentUserID != targetUserID {
		var isPrivate bool
		err = database.QueryRow("SELECT is_private FROM users WHERE id = ?", targetUserID).Scan(&isPrivate)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		if isPrivate {
			// Check if current user is following this profile
			var isFollowing bool
			err = database.QueryRow(`
				SELECT EXISTS(
					SELECT 1 FROM followers 
					WHERE follower_id = ? AND following_id = ?
				)`, currentUserID, targetUserID).Scan(&isFollowing)

			if err != nil || !isFollowing {
				http.Error(w, "Unauthorized to view this private profile", http.StatusForbidden)
				return
			}
		}
	}

	// Fetch following users
	rows, err := database.Query(`
		SELECT u.id, u.username, u.first_name, u.last_name
		FROM users u
		JOIN followers f ON u.id = f.following_id
		WHERE f.follower_id = ?
	`, targetUserID)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	following := []struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	}{}

	for rows.Next() {
		var follow struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			Firstname string `json:"firstname"`
			Lastname  string `json:"lastname"`
		}
		err := rows.Scan(&follow.ID, &follow.Username, &follow.Firstname, &follow.Lastname)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		following = append(following, follow)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(following)
}

// Post represents a user's post
type Post struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	Username  string `json:"username"`
	Likes     int    `json:"likes"`
}

// GetUserPosts returns all posts made by a specific user
func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	// Check if user is authenticated
	currentUserID, err := session.GetUserID(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Extract user ID from URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var targetUserID int
	if pathParts[3] == "me" {
		targetUserID = currentUserID
	} else {
		var err error
		targetUserID, err = strconv.Atoi(pathParts[3])
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
	}

	// Get database connection
	database, err := db.GetDB()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Check if profile is private and if current user is allowed to view it
	if currentUserID != targetUserID {
		var isPrivate bool
		err = database.QueryRow("SELECT is_private FROM users WHERE id = ?", targetUserID).Scan(&isPrivate)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		if isPrivate {
			// Check if current user is following this profile
			var isFollowing bool
			err = database.QueryRow(`
				SELECT EXISTS(
					SELECT 1 FROM followers 
					WHERE follower_id = ? AND following_id = ?
				)`, currentUserID, targetUserID).Scan(&isFollowing)

			if err != nil || !isFollowing {
				http.Error(w, "Unauthorized to view this private profile", http.StatusForbidden)
				return
			}
		}
	}

	// Fetch user posts
	rows, err := database.Query(`
		SELECT p.id, p.user_id, p.content, p.created_at, u.username, 
		       (SELECT COUNT(*) FROM post_likes WHERE post_id = p.id) as likes
		FROM posts p
		JOIN users u ON p.user_id = u.id
		WHERE p.user_id = ?
		ORDER BY p.created_at DESC
	`, targetUserID)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	posts := []Post{}
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Content, &post.CreatedAt, &post.Username, &post.Likes)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// Activity represents a user activity event
type Activity struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	TargetID  int    `json:"target_id,omitempty"`
}

// GetUserActivity returns the activity history of a user
func GetUserActivity(w http.ResponseWriter, r *http.Request) {
	// Check if user is authenticated
	currentUserID, err := session.GetUserID(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Extract user ID from URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var targetUserID int
	if pathParts[3] == "me" {
		targetUserID = currentUserID
	} else {
		var err error
		targetUserID, err = strconv.Atoi(pathParts[3])
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
	}

	// Get database connection
	database, err := db.GetDB()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Check if profile is private and if current user is allowed to view it
	if currentUserID != targetUserID {
		var isPrivate bool
		err = database.QueryRow("SELECT is_private FROM users WHERE id = ?", targetUserID).Scan(&isPrivate)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		if isPrivate {
			// Check if current user is following this profile
			var isFollowing bool
			err = database.QueryRow(`
				SELECT EXISTS(
					SELECT 1 FROM followers 
					WHERE follower_id = ? AND following_id = ?
				)`, currentUserID, targetUserID).Scan(&isFollowing)

			if err != nil || !isFollowing {
				http.Error(w, "Unauthorized to view this private profile", http.StatusForbidden)
				return
			}
		}
	}

	// Fetch user activity
	rows, err := database.Query(`
		SELECT id, type, content, created_at, target_id
		FROM user_activity
		WHERE user_id = ?
		ORDER BY created_at DESC
		LIMIT 50
	`, targetUserID)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	activities := []Activity{}
	for rows.Next() {
		var activity Activity
		var targetID sql.NullInt64
		err := rows.Scan(&activity.ID, &activity.Type, &activity.Content, &activity.CreatedAt, &targetID)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		if targetID.Valid {
			activity.TargetID = int(targetID.Int64)
		}

		activities = append(activities, activity)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activities)
}

// ProfileData represents all profile information
type ProfileData struct {
	User      User       `json:"user"`
	Posts     []Post     `json:"posts"`
	Activity  []Activity `json:"activity"`
	Followers []struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	} `json:"followers"`
	Following []struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	} `json:"following"`
	IsOwnProfile bool `json:"is_own_profile"`
}

// GetCompleteProfile returns all profile information for a user
func GetCompleteProfile(w http.ResponseWriter, r *http.Request) {
	// Check if user is authenticated
	currentUserID, err := session.GetUserID(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Extract profile ID from URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 {
		http.Error(w, "Invalid profile ID", http.StatusBadRequest)
		return
	}

	var targetUserID int
	if pathParts[3] == "me" {
		targetUserID = currentUserID
	} else {
		var err error
		targetUserID, err = strconv.Atoi(pathParts[3])
		if err != nil {
			http.Error(w, "Invalid profile ID", http.StatusBadRequest)
			return
		}
	}

	// Get database connection
	database, err := db.GetDB()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Fetch user profile
	var user User
	err = database.QueryRow(`
		SELECT id, username, email, first_name, last_name, bio, is_private 
		FROM users 
		WHERE id = ?`, targetUserID).Scan(
		&user.ID, &user.Username, &user.Email, &user.Firstname, &user.Lastname, &user.Bio, &user.IsPrivate,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	// Check if profile is private and if current user is allowed to view it
	isOwnProfile := currentUserID == targetUserID
	canViewFullProfile := isOwnProfile

	if !canViewFullProfile && user.IsPrivate {
		// Check if current user is following this profile
		var isFollowing bool
		err = database.QueryRow(`
			SELECT EXISTS(
				SELECT 1 FROM followers 
				WHERE follower_id = ? AND following_id = ?
			)`, currentUserID, targetUserID).Scan(&isFollowing)

		canViewFullProfile = err == nil && isFollowing
	}

	// If not own profile and can't view full profile, return limited info
	if !canViewFullProfile {
		limitedProfile := struct {
			User struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Firstname string `json:"firstname"`
				Lastname  string `json:"lastname"`
				IsPrivate bool   `json:"is_private"`
			} `json:"user"`
			IsOwnProfile bool `json:"is_own_profile"`
		}{
			User: struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Firstname string `json:"firstname"`
				Lastname  string `json:"lastname"`
				IsPrivate bool   `json:"is_private"`
			}{
				ID:        user.ID,
				Username:  user.Username,
				Firstname: user.Firstname,
				Lastname:  user.Lastname,
				IsPrivate: user.IsPrivate,
			},
			IsOwnProfile: false,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(limitedProfile)
		return
	}

	// Prepare complete profile data
	profileData := ProfileData{
		User:         user,
		IsOwnProfile: isOwnProfile,
	}

	// Fetch user posts
	postRows, err := database.Query(`
		SELECT p.id, p.user_id, p.content, p.created_at, u.username, 
		       (SELECT COUNT(*) FROM post_likes WHERE post_id = p.id) as likes
		FROM posts p
		JOIN users u ON p.user_id = u.id
		WHERE p.user_id = ?
		ORDER BY p.created_at DESC
	`, targetUserID)
	if err == nil {
		defer postRows.Close()
		for postRows.Next() {
			var post Post
			err := postRows.Scan(&post.ID, &post.UserID, &post.Content, &post.CreatedAt, &post.Username, &post.Likes)
			if err == nil {
				profileData.Posts = append(profileData.Posts, post)
			}
		}
	}

	// Fetch user activity
	activityRows, err := database.Query(`
		SELECT id, type, content, created_at, target_id
		FROM user_activity
		WHERE user_id = ?
		ORDER BY created_at DESC
		LIMIT 50
	`, targetUserID)
	if err == nil {
		defer activityRows.Close()
		for activityRows.Next() {
			var activity Activity
			var targetID sql.NullInt64
			err := activityRows.Scan(&activity.ID, &activity.Type, &activity.Content, &activity.CreatedAt, &targetID)
			if err == nil {
				if targetID.Valid {
					activity.TargetID = int(targetID.Int64)
				}
				profileData.Activity = append(profileData.Activity, activity)
			}
		}
	}

	// Fetch followers
	followerRows, err := database.Query(`
		SELECT u.id, u.username, u.first_name, u.last_name
		FROM users u
		JOIN followers f ON u.id = f.follower_id
		WHERE f.following_id = ?
	`, targetUserID)
	if err == nil {
		defer followerRows.Close()
		for followerRows.Next() {
			var follower struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Firstname string `json:"firstname"`
				Lastname  string `json:"lastname"`
			}
			err := followerRows.Scan(&follower.ID, &follower.Username, &follower.Firstname, &follower.Lastname)
			if err == nil {
				profileData.Followers = append(profileData.Followers, follower)
			}
		}
	}

	// Fetch following
	followingRows, err := database.Query(`
		SELECT u.id, u.username, u.first_name, u.last_name
		FROM users u
		JOIN followers f ON u.id = f.following_id
		WHERE f.follower_id = ?
	`, targetUserID)
	if err == nil {
		defer followingRows.Close()
		for followingRows.Next() {
			var following struct {
				ID        int    `json:"id"`
				Username  string `json:"username"`
				Firstname string `json:"firstname"`
				Lastname  string `json:"lastname"`
			}
			err := followingRows.Scan(&following.ID, &following.Username, &following.Firstname, &following.Lastname)
			if err == nil {
				profileData.Following = append(profileData.Following, following)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profileData)
}

// FollowUser handles a user following another user
func FollowUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check if user is authenticated
	followerID, err := session.GetUserID(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Extract user ID to follow from URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	followingID, err := strconv.Atoi(pathParts[3])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Can't follow yourself
	if followerID == followingID {
		http.Error(w, "Cannot follow yourself", http.StatusBadRequest)
		return
	}

	// Get database connection
	database, err := db.GetDB()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Check if user exists
	var exists bool
	err = database.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", followingID).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Check if already following
	err = database.QueryRow("SELECT EXISTS(SELECT 1 FROM followers WHERE follower_id = ? AND following_id = ?)",
		followerID, followingID).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if exists {
		// Already following, return success
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]bool{"success": true})
		return
	}

	// Add follow relationship
	_, err = database.Exec("INSERT INTO followers (follower_id, following_id, created_at) VALUES (?, ?, datetime('now'))",
		followerID, followingID)
	if err != nil {
		http.Error(w, "Failed to follow user", http.StatusInternalServerError)
		return
	}

	// Add activity record
	_, err = database.Exec(
		"INSERT INTO user_activity (user_id, type, content, target_id, created_at) VALUES (?, ?, ?, ?, datetime('now'))",
		followerID, "follow", "started following a user", followingID,
	)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}

// UnfollowUser handles a user unfollowing another user
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check if user is authenticated
	followerID, err := session.GetUserID(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Extract user ID to unfollow from URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	followingID, err := strconv.Atoi(pathParts[3])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Get database connection
	database, err := db.GetDB()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Delete follow relationship
	_, err = database.Exec("DELETE FROM followers WHERE follower_id = ? AND following_id = ?",
		followerID, followingID)
	if err != nil {
		http.Error(w, "Failed to unfollow user", http.StatusInternalServerError)
		return
	}

	// Add activity record
	_, err = database.Exec(
		"INSERT INTO user_activity (user_id, type, content, target_id, created_at) VALUES (?, ?, ?, ?, datetime('now'))",
		followerID, "unfollow", "unfollowed a user", followingID,
	)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}

// CheckFollowStatus checks if the current user is following another user
func CheckFollowStatus(w http.ResponseWriter, r *http.Request) {
	// Check if user is authenticated
	followerID, err := session.GetUserID(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Extract user ID to check from URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	followingID, err := strconv.Atoi(pathParts[3])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Get database connection
	database, err := db.GetDB()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Check if following
	var isFollowing bool
	err = database.QueryRow(`
		SELECT EXISTS(
			SELECT 1 FROM followers 
			WHERE follower_id = ? AND following_id = ?
		)`, followerID, followingID).Scan(&isFollowing)

	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"is_following": isFollowing})
}
