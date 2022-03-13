package migrations

import (
	"golang-twitter-clone/database"
	"golang-twitter-clone/helpers"
	"golang-twitter-clone/interfaces"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createDefaultAccount() {

	user := &interfaces.User{
		FullName: "Monzer Mahmoud",
		Email:    "Monzer@Mahmoud.com",
		Username: "@monzer97",
		Password: helpers.HashPassword("123456"),
	}
	database.DB.Create(&user)
}

func Migrate() {
	User := &interfaces.User{}
	Tweet := &interfaces.Tweet{}
	Follow := &interfaces.Follow{}
	database.DB.AutoMigrate(&User, &Tweet, &Follow)

	createDefaultAccount()
}







