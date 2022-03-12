package users

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang-twitter-clone/helpers"
	"golang-twitter-clone/interfaces"
	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) map[string]interface{} {

	valid := helpers.Validation(
		[]interfaces.Validation{
			{Value: email, Valid: "email"},
			{Value: password, Valid: "password"},
		})

	if valid {
		db := helpers.ConnectDB()
		user := &interfaces.User{}

		if db.Where("email = ?", email).First(&user).RecordNotFound() {
			return map[string]interface{}{"message": "User not found"}
		}

		passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

		if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
			return map[string]interface{}{"message": "Invalid password"}
		}

		defer db.Close()

		var response = prepareResponse(user)

		return response
	} else {
		return map[string]interface{}{"message": "not valid values"}
	}
}

func prepareResponse(user *interfaces.User) map[string]interface{} {
	responseUser := &interfaces.ResponseUser{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Username: user.Username,
	}

	var token = prepareToken(user)

	var response = map[string]interface{}{"message": "All is fine"}
	response["data"] = responseUser
	response["token"] = token

	return response
}

func prepareToken(user *interfaces.User) string {
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("secret-key"))

	helpers.HandleErr(err)

	return token
}

func Register(fullName string, username string, email string, password string) map[string]interface{} {
	valid := helpers.Validation(
		[]interfaces.Validation{
			{Value: username, Valid: "username"},
			{Value: email, Valid: "email"},
			{Value: password, Valid: "password"},
		})

	if valid {
		db := helpers.ConnectDB()

		generatedPassword := helpers.HashPassword(password)
		user := &interfaces.User{FullName: fullName, Email: email, Password: generatedPassword, Username: username}
		db.Create(&user)

		defer db.Close()

		var response = prepareResponse(user)

		return response

	} else {
		return map[string]interface{}{"message": "not valid values"}
	}
}
