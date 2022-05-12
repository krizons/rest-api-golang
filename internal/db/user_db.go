package db

import "github.com/krizons/rest-api-golang/internal/model"

type UserDB interface {
	GetAll() ([]model.User, error)
	Create(user model.User)
	Filter(user_filtr model.User) ([]model.User, error)
	Range(age_start int, end_age int) ([]model.User, error)
	Order(field string, sorting string) ([]model.User, error)
	TextSearch(search string) ([]model.User, error)
}
