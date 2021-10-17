package psql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/elnur0000/tweet-app/src/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func OpenDB() error {

	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		config.Config.Psql.User,
		config.Config.Psql.Password,
		config.Config.Psql.Host,
		config.Config.Psql.Port,
		config.Config.Psql.DB)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return err
	}

	DB = db

	return nil
}
