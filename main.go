package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	initializeRouter()
}

func initializeRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/users/login", loginHandler)
	router.HandleFunc("/users/logout", logoutHandler)
	router.HandleFunc("/users/signup", signupHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Println("Listening on port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to twitter clone"))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login"))
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout"))
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Signup"))
}










