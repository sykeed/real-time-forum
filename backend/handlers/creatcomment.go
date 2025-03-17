package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"real-time-forum/backend/app"
	"real-time-forum/backend/models"
	"real-time-forum/database"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JsonResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	userID, err := app.ValidateCookie(database.DB, w, r)
	if err != nil {
		JsonResponse(w, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	var comment models.Comment
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		fmt.Println(err)
		JsonResponse(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if comment.Content == "" || comment.PostID == "" {
		JsonResponse(w, http.StatusBadRequest, "Post ID and content are required", nil)
		return
	}

	// Get user nickname
	var nickname string
	err = database.DB.QueryRow("SELECT nickname FROM users WHERE id = ?", userID).Scan(&nickname)
	if err != nil {
		JsonResponse(w, http.StatusInternalServerError, "Error getting user information", err)
		return
	}

	// Insert the comment
	res, err := database.DB.Exec("INSERT INTO comments (post_id, user_id, content) VALUES (?, ?, ?)",
		comment.PostID, userID, comment.Content)
	if err != nil {
		fmt.Println("Error creating comment:", err)
		JsonResponse(w, http.StatusInternalServerError, "Error creating comment", err)
		return
	}

	commentID, err := res.LastInsertId()
	if err != nil {
		JsonResponse(w, http.StatusInternalServerError, "Error getting comment ID", err)
		return
	}

	// Get the comment with creation timestamp
	var createdComment models.Comment
	err = database.DB.QueryRow(`
		SELECT id, post_id, user_id, content, created_at 
		FROM comments 
		WHERE id = ?`, 
		commentID).Scan(
			&createdComment.ID, 
			&createdComment.PostID, 
			&createdComment.UserID, 
			&createdComment.Content, 
			&createdComment.CreatedAt,
		)
	if err != nil {
		JsonResponse(w, http.StatusInternalServerError, "Error retrieving created comment", err)
		return
	}

	// Add the author's nickname to the response
	createdComment.Author = nickname

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdComment)
}