package sqlite_user

import (
	"strings"

	"github.com/krizons/rest-api-golang/internal/model"
)

func (db *MyDb) GetAll() ([]model.User, error) {
	var users []model.User
	res := db.Db.Find(&users)

	return users, res.Error
}
func (db *MyDb) Create(user model.User) {
	db.Db.Create(&user)

}
func (db *MyDb) Filter(user_filtr model.User) ([]model.User, error) {
	var users []model.User
	res := db.Db.Where(user_filtr).Find(&users)
	return users, res.Error
}
func (db *MyDb) Range(age_start int, end_age int) ([]model.User, error) {
	var users []model.User
	res := db.Db.Where("age >= ? and age <= ?", age_start, end_age).Find(&users)
	return users, res.Error
}
func (db *MyDb) Order(field string, sorting string) ([]model.User, error) {
	var users []model.User
	res := db.Db.Order(field + " " + sorting).Find(&users)
	return users, res.Error
}
func (db *MyDb) TextSearch(search string) ([]model.User, error) {
	var users []model.User
	var out_user []model.User
	res := db.Db.Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	for _, person := range users {
		if strings.Contains(person.Name, search) {
			out_user = append(out_user, person)
		}
	}
	return out_user, nil
}
