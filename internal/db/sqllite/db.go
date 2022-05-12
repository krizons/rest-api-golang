package sqlite

import (
	"github.com/krizons/rest-api-golang/internal/db"
	"github.com/krizons/rest-api-golang/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MyDb struct {
	user_db  db.UserDB
	order_db db.OrderDB
}

func New(path string) (*MyDb, error) {
	db, err := gorm.Open(sqlite.Open(path))

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Order{})
	if err != nil {
		return nil, err
	}

	return &MyDb{user_db: &UserDB{Db: db}, order_db: &OrderDB{Db: db}}, nil
}
func (my *MyDb) User() db.UserDB {
	return my.user_db
}
func (my *MyDb) Order() db.OrderDB {
	return my.order_db
}
