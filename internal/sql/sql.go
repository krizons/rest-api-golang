package sql_user

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MyDb struct {
	Db *gorm.DB
}

func New(url string) (*MyDb, error) {
	db, err := gorm.Open(sqlite.Open(url))
	db.Migrator(&User{})
	if err != nil {
		return nil, err
	}

	return &MyDb{Db: db}, nil
}
