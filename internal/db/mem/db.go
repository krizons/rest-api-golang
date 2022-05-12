package my_mem

import (
	"github.com/krizons/rest-api-golang/internal/db"
	"github.com/krizons/rest-api-golang/internal/model"
)

type MyDb struct {
	user_db db.UserDB
}

func New() (*MyDb, error) {

	return &MyDb{user_db: &UserDB{users: make(map[int]model.User)}}, nil
}
func (my *MyDb) User() db.UserDB {
	return my.user_db
}
