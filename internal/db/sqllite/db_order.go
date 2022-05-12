package sqlite

import (
	"github.com/krizons/rest-api-golang/internal/model"
	"gorm.io/gorm"
)

type OrderDB struct {
	Db *gorm.DB
}

func (db *OrderDB) GetAll() ([]model.Order, error) {
	var orders []model.Order
	res := db.Db.Find(&orders)

	return orders, res.Error
}
func (db *OrderDB) Create(user model.Order) {
	db.Db.Create(&user)

}
func (db *OrderDB) Filter(user_filtr model.Order) ([]model.Order, error) {
	var orders []model.Order
	res := db.Db.Where(user_filtr).Find(&orders)
	return orders, res.Error
}

func (db *OrderDB) Sorted(field string, sorting string) ([]model.Order, error) {
	var orders []model.Order
	res := db.Db.Order(field + " " + sorting).Find(&orders)
	return orders, res.Error
}
