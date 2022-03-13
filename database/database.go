package database

import (
	"golang-twitter-clone/helpers"
	"os"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	DBURL := os.Getenv("DATABASE_URL")
	if DBURL == "" {DBURL = "host=localhost port=5432 user=postgres dbname=Twitter-Clone sslmode=disable"}
	database, err := gorm.Open("postgres", DBURL)
	helpers.HandleErr(err)
	database.DB().SetMaxIdleConns(20)
	database.DB().SetMaxOpenConns(200)
	DB = database
}