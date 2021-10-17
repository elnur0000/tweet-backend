package tweet

import (
	"github.com/labstack/echo/v4"
)

func RegisterTweetRoutes(router *echo.Group) {
	router.GET("", TweetController.GetTweets)
}
