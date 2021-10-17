package tweet

type tweetService struct{}

var TweetService tweetService = tweetService{}

func (*tweetService) Create() {
	TweetModel.CreateTweet()
}
