package user

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/elnur0000/tweet-app/src/db/psql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Avatar    string    `json:"avatar"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Version   int32     `json:"-"`
}

type UserModeler interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
}

type userModel struct{}

var UserModel UserModeler = &userModel{}

func (m *userModel) Create(user *User) error {
	query := `
		INSERT INTO "user" (avatar, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, version
	`
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return err
	}

	args := []interface{}{user.Avatar, user.Email, string(password)}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return psql.DB.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt, &user.Version)
}

func (m *userModel) FindByEmail(email string) (*User, error) {
	query := `
	SELECT id, created_at, avatar, email, version, password from "user"
	WHERE email = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user User
	err := psql.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Avatar,
		&user.Email,
		&user.Version,
		&user.Password,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, nil
		default:
			return nil, err
		}
	}

	return &user, nil
}
