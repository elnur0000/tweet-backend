package tweet

import (
	"database/sql"
	"fmt"

	"github.com/elnur0000/tweet-app/src/db/psql"
)

type Tweet struct {
	content    string
	likeCount  int
	replyCount int
}

type TweetModeler interface {
	CreateTweet()
}

type tweetModel struct {
	DB *sql.DB
}

var TweetModel TweetModeler = &tweetModel{
	DB: psql.DB,
}

func (tm *tweetModel) CreateTweet() {
	fmt.Println("hello")
}
