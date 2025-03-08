package tools

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"real-time-forum/database"
)

func Middleware(motherfunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie("session_id")
		if err != nil || session.Value == "" {
			log.Println("Unauthorized. Redirecting to login.")
			w.WriteHeader(http.StatusUnauthorized) // Return 401 status
			json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized"})
			return
		}

		err = ValidateSession(session.Value)
		if err != nil {
			log.Println("Unauthorized. Redirecting to login.")
			w.WriteHeader(http.StatusUnauthorized) // Return 401 status
			json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized"})
			return
		}

		motherfunc(w, r)
	}
}

func ValidateSession(sessionID string) error {
	var userID string
	query := `SELECT user_id FROM sessions WHERE session = ? AND exp_date > ?`
	err := database.DB.QueryRow(query, sessionID, time.Now()).Scan(&userID)
	if err != nil {
		return fmt.Errorf("unauthorized: session not found")
	}

	return nil
}
//aszdzadcze