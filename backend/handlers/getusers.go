package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"real-time-forum/backend/models"
	"real-time-forum/database"
)


func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ana hnaaa")
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	if !strings.HasPrefix(r.URL.Path, "/api/users") {
		http.NotFound(w, r)
		return
	}
	cookie, err := r.Cookie("session")
	if err != nil {
		JsonResponse(w, http.StatusUnauthorized, "No session found", nil)
		return
	}
	var userID int

	err = database.DB.QueryRow("SELECT user_id FROM sessions WHERE session = ?", cookie.Value).Scan(&userID)
	if err != nil {
		JsonResponse(w, http.StatusOK, "Session valid", nil)
		if err == sql.ErrNoRows {
			JsonResponse(w, http.StatusUnauthorized, "Invalid session", nil)
		}
		return
	}
	var nickname string
	err = database.DB.QueryRow("SELECT nickname FROM users WHERE id =?", userID).Scan(&nickname)
	if err != nil {
		JsonResponse(w, http.StatusOK, "this is probleme in geting nickname from DB", nil)
		if err == sql.ErrNoRows {
			JsonResponse(w, http.StatusUnauthorized, "this is probleme in geting nickname from DB", nil)
		}
		return
	}
	query := `
    WITH LastMessages AS (
        SELECT 
            u.nickname,
			MAX(m.created_at) as last_message_time
        FROM users u
        LEFT JOIN (
            SELECT sender, receiver, created_at
            FROM messages 
            WHERE sender = ?
               OR receiver = ?
        ) m ON (u.nickname = m.sender OR u.nickname = m.receiver)
        WHERE u.id != ?
        GROUP BY u.nickname
    )
    SELECT 
        nickname
    FROM LastMessages
    ORDER BY 
        CASE 
            WHEN last_message_time IS NOT NULL THEN 1
            ELSE 2
        END,
        last_message_time DESC NULLS LAST,
        nickname ASC;
    `
	rows, err := database.DB.Query(query, nickname, nickname, userID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var nicknames []models.Nickname
	for rows.Next() {
		var name models.Nickname
		if err := rows.Scan(&name.Username); err != nil {
			fmt.Println("internal server error")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		fmt.Println(name)
		nicknames = append(nicknames, name)
	}

	responseData, err := json.Marshal(nicknames)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}
