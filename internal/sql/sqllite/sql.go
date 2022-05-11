package sqlite_user

import (
	"github.com/krizons/rest-api-golang/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MyDb struct {
	Db          *gorm.DB
	DatabaseURL string
}

func New(path string) (*MyDb, error) {
	db, err := gorm.Open(sqlite.Open(path))

	db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}

	return &MyDb{Db: db, DatabaseURL: path}, nil
}
