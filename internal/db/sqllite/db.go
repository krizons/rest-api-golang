package sqlite

import (
	"context"

	"github.com/krizons/rest-api-golang/internal/db"
	"github.com/krizons/rest-api-golang/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MyDb struct {
	user_db  db.UserDB
	order_db db.OrderDB
	db       *gorm.DB
}

func New(path string) (*MyDb, error) {
	db, err := gorm.Open(sqlite.Open(path))

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Order{})
	if err != nil {
		return nil, err
	}

	return &MyDb{db: db, user_db: &UserDB{Db: db}, order_db: &OrderDB{Db: db}}, nil
}
func (my *MyDb) User() db.UserDB {
	return my.user_db
}
func (my *MyDb) Order() db.OrderDB {
	return my.order_db
}
func (my *MyDb) Close(ctx context.Context) error {

	sql_DB, err := my.db.DB()
	if err != nil {
		return err
	}
	return sql_DB.Close()
}
