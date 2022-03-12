package router

// import (
// 	// "strings"
// 	// "time"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"log"
// 	"os"

// 	"golang-twitter-clone/helpers"
// 	// "github.com/bnkamalesh/errors"
// 	"github.com/gorilla/mux"

// )

// type User struct {
// 	FirstName string `json:"firstName,omitempty"`
// 	LastName  string `json:"lastName,omitempty"`
// 	Email     string `json:"email,omitempty"`
// 	Mobile	string `json:"mobile,omitempty"`
// 	Password string `json:"password,omitempty"`
// 	//Username
// 	// CreatedAt *time.Time `json:"createdAt,omitempty"`
// 	// UpdatedAt *time.Time `json:"updatedAt,omitempty"`
// }

// func InitializeRouter() {
// 	router := mux.NewRouter()

// 	router.HandleFunc("/", homeHandler)
// 	router.HandleFunc("/users/login", loginHandler)
// 	router.HandleFunc("/users/logout", logoutHandler)
// 	router.HandleFunc("/users/signup", signupHandler)

// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8001"
// 	}
// 	log.Println("Listening on port: ", port)
// 	log.Fatal(http.ListenAndServe(":"+port, router))
// }

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Welcome to twitter clone"))
// }

// func loginHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Login"))
// }

// func logoutHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Logout"))
// }

// func signupHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: signup") 

// 	w.Header().Set("Content-Type","application/json")   

// 	var user User
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	user.Password = helpers.HashPassword(user.Password)

// 	fmt.Println(user)
// 	w.Write([]byte("User created"))
// }


// func (u *User) setDefaults() {
// 	now := time.Now()
// 	if u.CreatedAt == nil {
// 		u.CreatedAt = &now
// 	}

// 	if u.UpdatedAt == nil {
// 		u.UpdatedAt = &now
// 	}
// }

// func (u *User) Sanitize() {
// 	u.FirstName = strings.TrimSpace(u.FirstName)
// 	u.LastName = strings.TrimSpace(u.LastName)
// 	u.Email = strings.TrimSpace(u.Email)
// 	u.Mobile = strings.TrimSpace(u.Mobile)
// }

// func (u *User) Validate() error {

// 	if u.FirstName == "" {
// 		return errors.Validation("First name is required")
// 	}

// 	if u.LastName == "" {
// 		return errors.Validation("Last name is required")
// 	}

// 	err := validateEmail(u.Email)
// 	if err != nil {
// 		return err
// 	}

// 	err = validateMobile(u.Mobile)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func validateEmail(email string) error {
// 	if email == "" {
// 		return errors.Validation("Email is required")
// 	}
// 	parts := strings.Split(email, "@")
// 	if len(parts) != 2 {
// 		return errors.New("Invalid email")
// 	}

// 	if parts[0] == "" {
// 		return errors.New("Invalid email")
// 	}

// 	if parts[1] == "" {
// 		return errors.New("Invalid email")
// 	}

// 	return nil
// }

// func validateMobile(mobile string) error {
// 	if mobile == "" {
// 		return errors.Validation("Mobile is required")
// 	}

// 	if len(mobile) != 10 {
// 		return errors.New("Invalid mobile number")
// 	}

// 	return nil
// }