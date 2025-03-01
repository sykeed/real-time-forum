package routes

import (
	"net/http"

	"real-time-forum/backend/handlers"
//	"real-time-forum/backend/tools"
)

func WebRoutes() {
	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/login-submit", handlers.LoginSubmit)
	http.HandleFunc("/register-submit", handlers.RegisterHandler)
	http.HandleFunc("/lougout", handlers.LogOutHandler)
	//	http.HandleFunc("/login", handlers.LoginHandler)
//	http.HandleFunc("/home", tools.AuthMiddleware(handlers.HandleHome))
//	http.HandleFunc("/sessionChecker", handlers.SessionChecker)

	// http.HandleFunc("/login",handlers.LoginHandler)
}
