package tweets

import (
	"golang-twitter-clone/database"
	"golang-twitter-clone/interfaces"
)

func getAllTweets() *[]interfaces.Tweet {
	tweets := &[]interfaces.Tweet{}
	database.DB.Find(&tweets)
	return tweets
}

func getAllTweetsOfUser(userID uint) *[]interfaces.Tweet {
	tweets := &[]interfaces.Tweet{}
	database.DB.Where("user_id = ?", userID).Find(&tweets)
	return tweets
}

// create function to get all tweets from users who are followed by the user
func getAllTweetsOfFollowedUsers(userID uint) *[]interfaces.Tweet {
	tweets := &[]interfaces.Tweet{}
	database.DB.Table("tweets").
		Joins("JOIN follows ON follows.followed_user_id = tweets.user_id").
		Where("follows.user_id = ?", userID).
		Find(&tweets)
	return tweets
}

// create a function to get all tweets from a specific user
func getAllTweetsOfSpecificUser(userID uint) *[]interfaces.Tweet {
	tweets := &[]interfaces.Tweet{}
	database.DB.Where("user_id = ?", userID).Find(&tweets)
	return tweets
}
