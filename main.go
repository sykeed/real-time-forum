package main

import (
	"fmt"
	"net/http"
	"os/exec"

	"real-time-forum/backend/routes"
	dataBase "real-time-forum/database"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	Servingfiles()
	dataBase.InitSchema()

	routes.WebRoutes()
	fmt.Println("Starting server on :8000")
	exec.Command("xdg-open", "http://localhost:8000/").Start()
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("err starting the server : ", err)
		return
	}
}

func Servingfiles() {
	// mime.AddExtensionType(".css", "text/css")
	http.Handle("/frontend/", http.StripPrefix("/frontend/", http.FileServer(http.Dir("./frontend/"))))
	// http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("../frontend/js"))))
	// http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("../frontend/public"))))
}
