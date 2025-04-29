package main

import (
	"net/http"

	"social-net/auth"
	"social-net/comments"
	"social-net/db"
	"social-net/posts"
	"social-net/profile"
	"social-net/session"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")        // Allow your frontend's origin
		w.Header().Set("Access-Control-Allow-Credentials", "true")                    // Allow credentials (cookies)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")          // Allowed methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Allowed headers

		// If the request method is OPTIONS, just return OK (preflight request)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	db.Initdb()
	http.HandleFunc("/api/login", auth.Login)
	http.HandleFunc("/middle", session.Middleware)
	http.HandleFunc("/api/info", auth.Getinfo)
	http.HandleFunc("/api/logout", auth.Logout)
	http.HandleFunc("/api/posts", posts.Post)
	http.HandleFunc("/api/getposts", posts.Getposts)
	http.HandleFunc("/api/getcomments", comments.Getcomments)
	http.HandleFunc("/api/addcomments", comments.AddComments)
	
	// New profile endpoints
	http.HandleFunc("/api/profile/", profile.GetProfile)
	http.HandleFunc("/api/profile/privacy", profile.UpdatePrivacy)
	http.HandleFunc("/api/followers/", profile.GetFollowers)
	http.HandleFunc("/api/following/", profile.GetFollowing)
	http.HandleFunc("/api/follow/", profile.FollowUser)
	http.HandleFunc("/api/unfollow/", profile.UnfollowUser)
	http.HandleFunc("/api/follow/status/", profile.CheckFollowStatus)
	http.HandleFunc("/api/posts/user/", profile.GetUserPosts)
	
	http.Handle("/api/", corsMiddleware(http.DefaultServeMux))
	http.ListenAndServe(":8080", nil)
}
