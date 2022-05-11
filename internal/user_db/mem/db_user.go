package my_mem

import (
	"sort"
	"strings"

	"github.com/krizons/rest-api-golang/internal/model"
)

func GetKeyAndValue(myMap map[int]model.User) ([]int, []model.User) {
	keys := make([]int, 0, len(myMap))
	values := make([]model.User, 0, len(myMap))

	for k, v := range myMap {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}
func (db *Db) GetAll() ([]model.User, error) {
	_, users := GetKeyAndValue(db.user)

	return users, nil
}
func (db *Db) Create(user model.User) {
	db.user[len(db.user)] = user
}
func (db *Db) Filter(user_filtr model.User) ([]model.User, error) {
	var users []model.User
	if user_filtr.Age != 0 {
		for _, u := range db.user {
			if u.Age == user_filtr.Age {
				users = append(users, u)
			}
		}
	} else if user_filtr.Country != "" {
		for _, u := range db.user {
			if u.Country == user_filtr.Country {
				users = append(users, u)
			}
		}
	} else if user_filtr.Name != "" {
		for _, u := range db.user {
			if u.Name == user_filtr.Name {
				users = append(users, u)
			}
		}
	}

	return users, nil
}
func (db *Db) Range(age_start int, end_age int) ([]model.User, error) {
	var users []model.User
	for _, user := range db.user {
		if user.Age >= uint8(age_start) && user.Age <= uint8(end_age) {
			users = append(users, user)
		}
	}
	return users, nil
}
func (db *Db) Order(field string, sorting string) ([]model.User, error) {
	_, users := GetKeyAndValue(db.user)

	sort.Slice(users, func(i, j int) (less bool) {
		if field == "name" {
			if sorting == "desc" {
				return users[i].Name > users[j].Name
			} else if sorting == "asc" {
				return users[i].Name < users[j].Name
			}
		} else if field == "age" {
			if sorting == "desc" {
				return users[i].Age > users[j].Age
			} else if sorting == "asc" {
				return users[i].Age < users[j].Age
			}
		} else if field == "country" {
			if sorting == "desc" {
				return users[i].Country > users[j].Country
			} else if sorting == "asc" {
				return users[i].Country < users[j].Country
			}
		}
		return false
	})

	return users, nil
}
func (db *Db) TextSearch(search string) ([]model.User, error) {
	var users []model.User

	for _, person := range db.user {
		if strings.Contains(person.Name, search) {
			users = append(users, person)
		}
	}
	return users, nil
}
