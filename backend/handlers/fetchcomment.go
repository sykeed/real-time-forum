package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"real-time-forum/backend/models"
	"real-time-forum/database"
)

func FetchComments(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		JsonResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	// Get post ID from query parameter
	postIDStr := r.URL.Query().Get("post_id")
	if postIDStr == "" {
		JsonResponse(w, http.StatusBadRequest, "Post ID is required", nil)
		return
	}

	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		JsonResponse(w, http.StatusBadRequest, "Invalid post ID", nil)
		return
	}

	// Query to get comments along with user nicknames
	query := `
		SELECT c.id, c.post_id, c.user_id, c.content, c.created_at, u.nickname
		FROM comments c
		JOIN users u ON c.user_id = u.id
		WHERE c.post_id = ?
		ORDER BY c.created_at ASC
	`

	rows, err := database.DB.Query(query, postID)
	if err != nil {
		fmt.Println("Error querying comments:", err)
		JsonResponse(w, http.StatusInternalServerError, "Error fetching comments", nil)
		return
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(
			&comment.ID,
			&comment.PostID,
			&comment.UserID,
			&comment.Content,
			&comment.CreatedAt,
			&comment.Author,
		)
		if err != nil {
			fmt.Println("Error scanning comment row:", err)
			continue
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error iterating through rows:", err)
		JsonResponse(w, http.StatusInternalServerError, "Error processing comments", nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}