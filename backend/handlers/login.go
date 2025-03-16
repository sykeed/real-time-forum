package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"real-time-forum/backend/app"
	"real-time-forum/backend/models"
	"real-time-forum/backend/tools"
	"real-time-forum/database"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var pasw string
	var user_id int

	if r.Method != http.MethodPost {
		JsonResponse(w, http.StatusMethodNotAllowed, "Invalid request method", nil)
		return
	}

	var user models.LoginCredentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		JsonResponse(w, http.StatusBadRequest, "error :", err)
		fmt.Println(err)
		return
	}

	err = database.DB.QueryRow("SELECT password, id FROM users WHERE nickname = ? OR email = ?", user.Email, user.Email).Scan(&pasw, &user_id)
	if err != nil {
		// fmt.Println("pasw:", pasw, "  user id :", user_id, "error : ", err)
		JsonResponse(w, http.StatusUnauthorized, "User not found or incorrect email", nil)
		return
	}
	defer database.DB.Close()
	err = tools.CheckPassword(pasw, user.Password)
	if err != nil {
		JsonResponse(w, http.StatusUnauthorized, "Invalid password", nil)
		return
	}

	// handling one session at a time
	deleteQuery := "DELETE FROM sessions WHERE user_id = ?"
	_, err = database.DB.Exec(deleteQuery, user_id)
	if err != nil {
		fmt.Println("Error cleaning old sessions", err)
		return
	}

	cookie := app.CookieMaker(w)
	err = app.InsretCookie(database.DB, user_id, cookie, time.Now().Add(time.Hour*24))
	if err != nil {
		fmt.Println(err)
		return
	}

	JsonResponse(w, http.StatusOK, "Login successful", nil)
}
