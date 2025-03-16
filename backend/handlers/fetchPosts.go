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

	query := "SELECT * FROM posts;"

	var post models.Post
	var posts []models.Post

	rows, err := database.DB.Query(query)
	if err != nil {
		fmt.Println("err2 : ",err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	for rows.Next() {

		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Category, &post.CreatedAt)
		if err != nil {
			fmt.Println("err1 : ",err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}
 

	// Set headers first, before writing any response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	// Then encode the response
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		// Can't change headers or status after writing, but we can log the error
		fmt.Println("Error encoding JSON response:", err)}
	
}
