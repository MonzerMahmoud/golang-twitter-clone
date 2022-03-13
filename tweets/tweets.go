package tweets

import (
	//"fmt"
	//"fmt"
	"golang-twitter-clone/helpers"
	"golang-twitter-clone/interfaces"
	//"golang-twitter-clone/users"
)

func getTweets() *[]interfaces.Tweet {
	db := helpers.ConnectDB()
	tweets := &[]interfaces.Tweet{}
	db.Find(&tweets)
	defer db.Close()

	return tweets
}

