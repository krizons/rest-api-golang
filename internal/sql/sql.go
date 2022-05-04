package sql_user

import (
	"github.com/krizons/rest-api-golang/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MyDb struct {
	Db     *gorm.DB
	config *Config
}

func New(config *Config) (*MyDb, error) {
	db, err := gorm.Open(sqlite.Open(config.DatabaseURL))

	db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}

	return &MyDb{Db: db, config: config}, nil
}
