package user

import (
	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(router *echo.Group) {
	router.POST("/register", UserController.Register)
	router.POST("/login", UserController.Login)
}
