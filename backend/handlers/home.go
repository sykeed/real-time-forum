package handlers

import (
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "frontend/public/login.html")
		} else {
			jsonResponse(w, http.StatusNotFound, http.StatusText(http.StatusNotFound), nil)
			return
		}
	} else {
		jsonResponse(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed), nil)
		return
	}
	/*
		if r.Method == "GET" {
			// Clean up the path to handle trailing slashes
			path := strings.TrimSuffix(r.URL.Path, "/")
			fmt.Println(path)
			// If it's the root path or /login, serve the login page
			if path == "" || path == "/login" {
				http.ServeFile(w, r, "frontend/public/login.html")
				return
			} else if path == "/register" {
				// For the register path, still serve the login.html
				// Your frontend JS will handle showing the correct form
				http.ServeFile(w, r, "frontend/public/login.html")
				return
			} else if path == "/home" {
				// For the home path, check if user is authenticated
				// This is a placeholder - you should implement proper auth checking
				//http.ServeFile(w, r, "frontend/public/login.html")
				jsonResponse(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed), nil)


			} else {
				// For any other path, return 404
				jsonResponse(w, http.StatusNotFound, http.StatusText(http.StatusNotFound), nil)
				return
			}
		} else {

			jsonResponse(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed), nil)
			return
		}
	*/
}
