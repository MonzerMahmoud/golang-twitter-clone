package users

import (
 	"time"

    "golang-twitter-clone/helpers"
    "golang-twitter-clone/interfaces"
    "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) map[string]interface{} {

	db := helpers.ConnectDB()
	user := &interfaces.User{}

	if db.Where("email = ?", email).First(&user).RecordNotFound() {
		return map[string]interface{}{"message": "User not found"}
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return map[string]interface{}{"message": "Invalid password"}
	}

	responseUser := &interfaces.ResponseUser{
		ID: user.ID,
		FullName: user.FullName,
		Email: user.Email,
		Username: user.Username,
	}

	defer db.Close()

	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("secret-key"))

	helpers.HandleErr(err)

	var response = map[string]interface{}{"message": "All is fine"}
	response["data"] = responseUser
	response["token"] = token

	return response
}