package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/elnur0000/tweet-app/src/types"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type userController struct{}

var UserController = userController{}

func (*userController) Register(c echo.Context) error {
	registerDTO := new(RegisterDTO)
	if err := c.Bind(registerDTO); err != nil {
		return err
	}
	if err := c.Validate(registerDTO); err != nil {
		return err
	}
	existingUser, err := UserModel.FindByEmail(registerDTO.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return echo.NewHTTPError(http.StatusConflict, "User with that email address already exists")
	}
	user := User{
		Avatar:   registerDTO.Avatar,
		Email:    registerDTO.Email,
		Password: registerDTO.Password,
	}

	if err := UserModel.Create(&user); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (*userController) Login(c echo.Context) error {
	loginDTO := new(LoginDTO)
	if err := c.Bind(loginDTO); err != nil {
		return err
	}
	if err := c.Validate(loginDTO); err != nil {
		return err
	}
	user, err := UserModel.FindByEmail(loginDTO.Email)
	if err != nil {
		return err
	}
	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Invalid Email")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDTO.Password))
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "Invalid password")
	}
	claims := &types.JWTCustomClaims{
		ID:    user.ID,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	c.Request().Header.Set("Authorization", fmt.Sprintf("Bearer %v", t))
	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
