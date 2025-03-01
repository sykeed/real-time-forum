package tools

import (
	"net/http"
	"real-time-forum/database"

	"golang.org/x/crypto/bcrypt"
)

//func IsEmail(s string) bool {
//	re := regexp.MustCompile(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,6}$`)
//  return re.MatchString(s)
//}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}


func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		cookie, err := r.Cookie("session")
		if err != nil {
			http.Error(w, "Unauthorized - No session", http.StatusUnauthorized)
			return
		}

		
		var userID int
		err = database.DB.QueryRow("SELECT user_id FROM sessions WHERE session = ?", cookie.Value).Scan(&userID)
		if err != nil {
			http.Error(w, "Unauthorized - Invalid session", http.StatusUnauthorized)
			return
		}

		// If session is valid, proceed to the next handler
		next(w, r)
	}
}
