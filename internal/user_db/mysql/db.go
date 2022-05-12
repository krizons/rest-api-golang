package mysql_user

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Db struct {
	Db          *sql.DB
	DatabaseURL string
}

func New(DatabaseURL string) (*Db, error) {

	db, err := sql.Open("mysql", DatabaseURL)

	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	my_db := &Db{Db: db, DatabaseURL: DatabaseURL}
	my_db.CreateTable()
	return my_db, nil
}
func Close(db *Db) {
	db.Db.Close()
}
