package users

import (
	"fmt"
	"time"

	"golang-twitter-clone/database"
	"golang-twitter-clone/helpers"
	"golang-twitter-clone/interfaces"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func prepareResponse(user *interfaces.User, withToken bool) map[string]interface{} {
	responseUser := &interfaces.ResponseUser{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Username: user.Username,
		
	}

	var response = map[string]interface{}{"message": "All is fine"}

	if withToken {
		var token = prepareToken(user)
		response["token"] = token
	}

	response["data"] = responseUser

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

func Login(email string, password string) map[string]interface{} {

	valid := helpers.Validation(
		[]interfaces.Validation{
			{Value: email, Valid: "email"},
			{Value: password, Valid: "password"},
		})

	if valid {
		user := &interfaces.User{}

		if database.DB.Where("email = ?", email).First(&user).RecordNotFound() {
			return map[string]interface{}{"message": "User not found"}
		}

		passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

		if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
			return map[string]interface{}{"message": "Invalid password"}
		}

		var response = prepareResponse(user, true)

		return response
	} else {
		return map[string]interface{}{"message": "not valid values"}
	}
}

func Register(fullName string, username string, email string, password string) map[string]interface{} {
	valid := helpers.Validation(
		[]interfaces.Validation{
			{Value: username, Valid: "username"},
			{Value: email, Valid: "email"},
			{Value: password, Valid: "password"},
			{Value: fullName, Valid: "fullname"},
		})

	if valid {
		generatedPassword := helpers.HashPassword(password)
		user := &interfaces.User{FullName: fullName, Email: email, Password: generatedPassword, Username: username, Following: 0, Followers: 0}
		database.DB.Create(&user)

		var response = prepareResponse(user, false)

		return response

	} else {
		return map[string]interface{}{"message": "not valid values"}
	}
}

func GetUser(id string, jwt string) map[string]interface{} {
	isValid := helpers.ValidateToken(id, jwt)
	if isValid {
		return GetOwnerUser(id)
	} else {
		return getOtherUser(id)
	}
}
func GetOwnerUser(id string) map[string]interface{} {
	
	user := &interfaces.User{}

	if database.DB.Where("id = ?", id).First(&user).RecordNotFound() {
		return map[string]interface{}{"message": "User not found"}
	}

	var response = prepareResponse(user, false)
	return response
}

func getOtherUser(id string) map[string]interface{} {
	
	user := &interfaces.User{}
	userModified := &interfaces.User{}

	if database.DB.Where("id = ?", id).First(&user).RecordNotFound() {
		return map[string]interface{}{"message": "User not found"}
	}

	userModified.ID = user.ID
	userModified.FullName = user.FullName
	userModified.Username = user.Username

	var response = prepareResponse(userModified, false)
	return response
}

func updateUsername(id string, username string) {
	database.DB.Model(&interfaces.User{}).Where("id = ?", id).Update("username", username)
}

func Follow(followerId uint, followeeId uint, jwt string) map[string]interface{} {

	if followerId != followeeId {
		followerIdString := fmt.Sprint(followerId)
		followingIdString := fmt.Sprint(followeeId)

		isValid := helpers.ValidateToken(followerIdString, jwt)
		
		if isValid {

			followerAccount := GetUser(followerIdString, jwt)
			followingAccount := GetUser(followingIdString, jwt)

			if followerAccount["message"] == "User not found" || followingAccount["message"] == "User not found" {
				return map[string]interface{}{"message": "User not found"}
			}

			follow := &interfaces.Follow{}
			if database.DB.Where("follower_id = ? AND followee_id = ?", followerId, followeeId).First(&follow).RecordNotFound() {
				follow := &interfaces.Follow{FollowerID: followerId, FolloweeID: followeeId}
				database.DB.Create(&follow)

				database.DB.Model(&interfaces.User{}).Where("id = ?", followerId).Update("following", gorm.Expr("following + ?", 1))
				database.DB.Model(&interfaces.User{}).Where("id = ?", followeeId).Update("followers", gorm.Expr("followers + ?", 1))

				return map[string]interface{}{"message": "Followed successfully"} 
			} else {
				database.DB.Delete(&follow)
				database.DB.Unscoped().Delete(&follow)

				database.DB.Model(&interfaces.User{}).Where("id = ?", followerId).Update("following", gorm.Expr("following - ?", 1))
				database.DB.Model(&interfaces.User{}).Where("id = ?", followeeId).Update("followers", gorm.Expr("followers - ?", 1))
			
				return map[string]interface{}{"message": "Unfollowed successfully"}
			}

		} else {

			return map[string]interface{}{"message": "Not valid token"}
		}
	} else {
		return map[string]interface{}{"message": "You can't follow yourself"}
	}

}
