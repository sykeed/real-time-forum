package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"real-time-forum/database"
)

func CheckSessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		JsonResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	cookie, err := r.Cookie("session")
	if err != nil {
		JsonResponse(w, http.StatusUnauthorized, "No session found", nil)
		return
	}

	var userID int
	var expiry time.Time
	err = database.DB.QueryRow("SELECT user_id, exp_date FROM sessions WHERE session = ?", cookie.Value).Scan(&userID, &expiry)
	if err != nil {
		JsonResponse(w, http.StatusOK, "Session valid", nil)
		if err == sql.ErrNoRows {
			JsonResponse(w, http.StatusUnauthorized, "Invalid session", nil)
		}
		return
	}
	//defer database.DB.Close()

	if time.Now().After(expiry) {

		database.DB.Exec("DELETE FROM sessions WHERE session = ?", cookie.Value)
		JsonResponse(w, http.StatusUnauthorized, "Session expired", nil)
		return
	}

	JsonResponse(w, http.StatusOK, "Session valid", nil)
}
