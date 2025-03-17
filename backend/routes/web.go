package routes

import (
	"net/http"

	"real-time-forum/backend/handlers"
)

func WebRoutes() {
	 http.HandleFunc("/ws", handlers.HandleConnections)
	http.HandleFunc("/api/login", handlers.LoginHandler)
	http.HandleFunc("/api/register", handlers.RegisterHandler)
	http.HandleFunc("/api/logout", handlers.LogOutHandler)  
	http.HandleFunc("/api/fetchposts", handlers.FetchPosts)  
	http.HandleFunc("/api/creatpost", handlers.CreatePost)  
	http.HandleFunc("/checksession", handlers.CheckSessionHandler)
	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/api/users", handlers.GetUsers)

	http.HandleFunc("/api/comments", handlers.FetchComments)       
	http.HandleFunc("/api/createcomment", handlers.CreateComment)

}