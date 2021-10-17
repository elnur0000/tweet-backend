package tweet

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type tweetController struct{}

var TweetController tweetController = tweetController{}

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (*tweetController) GetTweets(c echo.Context) error {
	u := Product{
		Name:        "Laptop",
		Description: "jon@labstack.com",
	}
	// return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
	return c.JSON(http.StatusOK, u)
}
