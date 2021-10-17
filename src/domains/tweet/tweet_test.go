package tweet

import (
	"fmt"
	"testing"
)

type testTweetModel struct{}

func (*testTweetModel) CreateTweet() {
	fmt.Println("hello324")
}

func TestCreateTest(t *testing.T) {
	TweetModel = &testTweetModel{}
	TweetService.Create()
	// fmt.Println("hello world")
}
