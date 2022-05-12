package sqlite_user

import (
	"github.com/krizons/rest-api-golang/internal/db"
	"github.com/krizons/rest-api-golang/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MyDb struct {
	user_db db.UserDB
}

func New(path string) (*MyDb, error) {
	db, err := gorm.Open(sqlite.Open(path))

	db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}

	return &MyDb{user_db: &UserDB{Db: db}}, nil
}
func (my *MyDb) User() db.UserDB {
	return my.user_db
}
