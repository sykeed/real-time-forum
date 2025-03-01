package handlers

import (
	"net/http"
	"real-time-forum/database"
)

func SessionChecker(w http.ResponseWriter, r *http.Request) {
	var user_id int
	cookie, err := r.Cookie("session")

	if err != nil {
		jsonResponse(w, http.StatusUnauthorized, "No session found", nil)
		return
	}
 
	query := `SELECT user_id FROM sessions WHERE session = ?`
	err = database.DB.QueryRow(query, cookie.Value).Scan(&user_id)
	if err != nil {
		jsonResponse(w, http.StatusUnauthorized, "Invalid session", nil)
		return
	}

	w.WriteHeader(http.StatusOK)
}
