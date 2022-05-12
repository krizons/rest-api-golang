package db

import "github.com/krizons/rest-api-golang/internal/model"

type OrderDB interface {
	GetAll() ([]model.Order, error)
	Create(user model.Order)
	Filter(user_filtr model.Order) ([]model.Order, error)
	Sorted(field string, sorting string) ([]model.Order, error)
}
