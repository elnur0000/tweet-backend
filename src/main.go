package main

import (
	"fmt"

	"github.com/elnur0000/tweet-app/src/config"
	"github.com/elnur0000/tweet-app/src/db/psql"
	"github.com/elnur0000/tweet-app/src/domains/tweet"
	"github.com/elnur0000/tweet-app/src/domains/user"
	"github.com/elnur0000/tweet-app/src/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	e := echo.New()
	e.Validator = utils.NewCustomValidator()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))

	err = psql.OpenDB()
	if err != nil {
		log.Fatal(err)
	}

	tweet.RegisterTweetRoutes(e.Group("/v1/tweets"))
	user.RegisterUserRoutes(e.Group("/v1/users"))

	e.Logger.Fatal(e.Start(fmt.Sprintf("127.0.0.1:%s", config.Config.Port)))
}
