package api

import (
	"encoding/json"
	"fmt"

	"strconv"

	"golang-twitter-clone/helpers"
	"golang-twitter-clone/interfaces"
	"golang-twitter-clone/tweets"
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

type Register struct {
	FullName string
	Username string
	Email    string
	Password string
}

func readBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	return body
}

func apiResponse(call map[string]interface{}, w http.ResponseWriter) {
	if call["message"] == "All is fine" {
		resp := call
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := interfaces.ErrResponse{Message: "Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body := readBody(r)

	var formattedBody Login
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	login := users.Login(formattedBody.Email, formattedBody.Password)

	apiResponse(login, w)
}

func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body := readBody(r)

	var formattedBody Register
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	register := users.Register(formattedBody.FullName, formattedBody.Username, formattedBody.Email, formattedBody.Password)

	apiResponse(register, w)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userId := params["id"]
	auth := r.Header.Get("Authorization")
	fmt.Println(helpers.GetUserIdFromToken(auth))
	user := users.GetUser(userId, auth)
	apiResponse(user, w)
}

func follow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	
	auth := r.Header.Get("Authorization")

	followerId, err := strconv.ParseUint(helpers.GetUserIdFromToken(auth), 10, 64)
	helpers.HandleErr(err)
	followeeId, err := strconv.ParseUint(params["id"], 10, 64)
	helpers.HandleErr(err)

	follow := users.Follow(uint(followerId), uint(followeeId), auth)
	apiResponse(follow, w)
}

func getTimeLine(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	auth := r.Header.Get("Authorization")
	userId, err := strconv.ParseUint(helpers.GetUserIdFromToken(auth), 10, 64)
	helpers.HandleErr(err)

	tweets := tweets.GetTimeLine(uint(userId), auth)
	
	json.NewEncoder(w).Encode(tweets)
}

func InitializeRouter() {
	fmt.Println("Initializing router...")
	router := mux.NewRouter()

	router.Use(helpers.PanicHandler)
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users/{id}/follow", follow).Methods("POST")
	router.HandleFunc("/tweets/timeline", getTimeLine).Methods("GET")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Listening on port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

