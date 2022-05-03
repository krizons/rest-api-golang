package sql_user

import (
	"github.com/krizons/rest-api-golang/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MyDb struct {
	Db *gorm.DB
}

func New(url string) (*MyDb, error) {
	db, err := gorm.Open(sqlite.Open(url))

	db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}

	return &MyDb{Db: db}, nil
}
