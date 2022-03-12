package api

import (
	"encoding/json"
	"golang-twitter-clone/helpers"
	"golang-twitter-clone/users"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Login struct {
	Email string
	Password string
}

type ErrResponse struct {
	Message string `json:"message"`
}

type Register struct {
	FullName string
	Username string
	Email    string
	Password string
	Birthday string
}

func readBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	return body
}

func login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	var formattedBody Login
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	login := users.Login(formattedBody.Email, formattedBody.Password)

	if login["message"] == "All is fine" {
		resp := login
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := ErrResponse{Message: "Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}
}

func InitializeRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/login", login).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Listening on port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

