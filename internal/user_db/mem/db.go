package my_mem

import (
	"github.com/krizons/rest-api-golang/internal/model"
)

type Db struct {
	user map[int]model.User
}

func New() (*Db, error) {

	return &Db{user: make(map[int]model.User)}, nil
}
