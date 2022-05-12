package my_mem

import (
	"sort"
	"strings"

	"github.com/krizons/rest-api-golang/internal/model"
)

type UserDB struct {
	users map[int]model.User
}

func GetKeyAndValue(myMap map[int]model.User) ([]int, []model.User) {
	keys := make([]int, 0, len(myMap))
	values := make([]model.User, 0, len(myMap))

	for k, v := range myMap {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}
func (db *UserDB) GetAll() ([]model.User, error) {
	_, users := GetKeyAndValue(db.users)

	return users, nil
}
func (db *UserDB) Create(user model.User) {
	db.users[len(db.users)] = user
}
func (db *UserDB) Filter(user_filtr model.User) ([]model.User, error) {
	var users []model.User
	if user_filtr.Age != 0 {
		for _, u := range db.users {
			if u.Age == user_filtr.Age {
				users = append(users, u)
			}
		}
	} else if user_filtr.Country != "" {
		for _, u := range db.users {
			if u.Country == user_filtr.Country {
				users = append(users, u)
			}
		}
	} else if user_filtr.Name != "" {
		for _, u := range db.users {
			if u.Name == user_filtr.Name {
				users = append(users, u)
			}
		}
	}

	return users, nil
}
func (db *UserDB) Range(age_start int, end_age int) ([]model.User, error) {
	var users []model.User
	for _, user := range db.users {
		if user.Age >= uint8(age_start) && user.Age <= uint8(end_age) {
			users = append(users, user)
		}
	}
	return users, nil
}
func (db *UserDB) Order(field string, sorting string) ([]model.User, error) {
	_, users := GetKeyAndValue(db.users)

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
func (db *UserDB) TextSearch(search string) ([]model.User, error) {
	var users []model.User

	for _, person := range db.users {
		if strings.Contains(person.Name, search) {
			users = append(users, person)
		}
	}
	return users, nil
}
