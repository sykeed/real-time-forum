package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"real-time-forum/backend/models"
	"real-time-forum/database"
)

func FetchPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Updated query to include the author's nickname from users table
	query := `
		SELECT p.id, p.user_id, p.title, p.content, p.category, p.created_at, u.nickname 
		FROM posts p
		JOIN users u ON p.user_id = u.id
		ORDER BY p.created_at DESC;
	`

	rows, err := database.DB.Query(query)
	if err != nil {
		fmt.Println("err2 : ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

 

var posts []models.Post

	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.ID, 
			&post.UserID, 
			&post.Title, 
			&post.Content, 
			&post.Category, 
			&post.CreatedAt,
			&post.Author,  
		)
		if err != nil {
			fmt.Println("err1 : ", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("err3 : ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// Set headers first, before writing any response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	// Then encode the response
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		// Can't change headers or status after writing, but we can log the error
		fmt.Println("Error encoding JSON response:", err)
	}
}