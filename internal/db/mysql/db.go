package mysql_user

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/krizons/rest-api-golang/internal/db"
)

type MyDb struct {
	user_db db.UserDB
}

func New(DatabaseURL string) (*MyDb, error) {

	db, err := sql.Open("mysql", DatabaseURL)

	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &MyDb{user_db: &UserDB{Db: db}}, nil
}
func (my *MyDb) User() db.UserDB {
	return my.user_db
}
