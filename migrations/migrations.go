package migrations

import (
	"fmt"
	"golang-twitter-clone/helpers"
	"golang-twitter-clone/interfaces"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createTweets() {
	// Todo : Implement this
	fmt.Println("Creating tweets")

	db := helpers.ConnectDB()

	users := &[2]interfaces.User{
		{ FullName: "John Doe", Email: "John@Doe.com", Password: "password"},
		{ FullName: "Jane Doe", Email: "Jane@Doe.com", Password: "password"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashPassword(users[i].Password)
		user := &interfaces.User{ FullName: users[i].FullName, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		tweets := &interfaces.Tweet{Body: "Welcome to twitter", UserID: user.ID}
		db.Create(&tweets)
	}

	defer db.Close()

}

func Migrate() {
	db := helpers.ConnectDB()
	db.AutoMigrate(&interfaces.User{})
	db.AutoMigrate(&interfaces.Tweet{})
	db.AutoMigrate(&interfaces.Follow{})
	defer db.Close()

	//createTweets()
}







