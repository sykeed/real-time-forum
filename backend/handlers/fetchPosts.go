package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"real-time-forum/backend/models"
	"real-time-forum/database"
)



func FetchPosts(w http.ResponseWriter , r *http.Request){

	if r.Method != http.MethodGet {
		http.Error(w , "Method not allowed" , http.StatusMethodNotAllowed)
	}

	query := "SELECT * FROM posts;"
	
	var post models.Post
	var posts []models.Post

	rows ,err := database.DB.Query(query)
	if err != nil {
		http.Error(w, "internal server error" , http.StatusInternalServerError)
		fmt.Println("HNA1")
	}


	for rows.Next() {
		fmt.Println(post)
	err := 	rows.Scan(&post.UserID,&post.Title,&post.Content,&post.Category,&post.CreatedAt)
	if err != nil {
		fmt.Println("HNA2",err)
		http.Error(w, "internal server error" , http.StatusInternalServerError)
		return
	}
	posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)

	 w.Header().Set("Content-Type","application/json")
	 w.WriteHeader(http.StatusOK)
}