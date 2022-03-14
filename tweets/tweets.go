package tweets

import (
	"golang-twitter-clone/database"
	"golang-twitter-clone/helpers"
	"golang-twitter-clone/interfaces"
)

func prepareResponse(tweets []*interfaces.ResponseTweet) map[string][]interface{} {

	var response = map[string][]interface{}{}

	for _, tweet := range tweets {
		response["data"] = append(response["data"], tweet)
	}
	//response["data"] = append(response["data"], tweets)

	if len(tweets) == 0 {
		response["message"] = append(response["message"], "No tweets found")
	}

	return response
}

func GetTimeLine(userID uint, jwtToken string) map[string][]interface{} {
	isValid := helpers.ValidateTokenExp(jwtToken)

	if isValid {
		tweets := getAllTweets()
		return prepareResponse(tweets)
	} else {
		return prepareResponse([]*interfaces.ResponseTweet{})
	}

}

func getAllTweets() []*interfaces.ResponseTweet {
	tweets := &[]interfaces.Tweet{}
	tweetsMetaData := []*interfaces.ResponseTweet{}

	database.DB.Find(&tweets)

	for _, tweet := range *tweets {
		user := &interfaces.User{}

		database.DB.Where("id = ?", tweet.UserID).First(&user)

		tweetsMetaData = append(tweetsMetaData, &interfaces.ResponseTweet{
			ID:        tweet.ID,
			UserID:    tweet.UserID,
			Body:      tweet.Body,
			CreatedAt: tweet.CreatedAt,
			User: interfaces.User{
				FullName: user.FullName,
				Username: user.Username},
		})
	}
	return tweetsMetaData
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
