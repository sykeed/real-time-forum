package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"real-time-forum/backend/models"
	"real-time-forum/database"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("rani f crercomment")
	if r.Method != http.MethodPost {
		fmt.Println("error 1")
		JsonResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	 
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("error 2")
		JsonResponse(w, http.StatusUnauthorized, "No session found", nil)
		return
	}

	 
	var userID int

	fmt.Println("kooki",cookie.Value)
	err = database.DB.QueryRow("SELECT user_id FROM sessions WHERE session = ?", cookie.Value).Scan(&userID)
	if err != nil {
		fmt.Println("error 3")
		JsonResponse(w, http.StatusUnauthorized, "Invalid session", nil)
		return
	}

	var comment models.Comment
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		fmt.Println("error 4")
		fmt.Println("Error decoding request body:", err)
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
 
	postID, err := strconv.Atoi(comment.PostID)
	if err != nil {
		JsonResponse(w, http.StatusBadRequest, "Invalid post ID format", nil)
		return
	}

	// Check if post exists
	var postExists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM posts WHERE id = ?)", postID).Scan(&postExists)
	if err != nil {
		JsonResponse(w, http.StatusInternalServerError, "Database error", err)
		return
	}

	if !postExists {
		JsonResponse(w, http.StatusNotFound, "Post not found", nil)
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

	 
	createdComment.Author = nickname

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdComment)
}
