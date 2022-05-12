package mysql_user

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/krizons/rest-api-golang/internal/model"
)

type UserDB struct {
	Db *sql.DB
}

func (db *UserDB) GetAll() ([]model.User, error) {
	var users []model.User
	results, err := db.Db.Query("SELECT name,age,country FROM users")
	if err != nil {
		return nil, err
	}
	for results.Next() {
		var user model.User

		err = results.Scan(&user.Name, &user.Age, &user.Country)
		users = append(users, user)
		if err != nil {
			return nil, err
		}

	}

	return users, nil
}

func (db *UserDB) Create(user model.User) {
	insert, err := db.Db.Query("INSERT INTO users( name,age,country ) VALUES ( ?, ?, ? )", user.Name, user.Age, user.Country)
	if err != nil {
		return
	}
	insert.Close()
	//db.Db.Create(&user)

}
func (db *UserDB) Filter(user_filtr model.User) ([]model.User, error) {
	var users []model.User
	var qtext string = ""
	if user_filtr.Age != 0 {

		qtext = fmt.Sprintf("SELECT name,age,country FROM users WHERE %s = %d;", "age", user_filtr.Age)
	} else if user_filtr.Country != "" {

		qtext = fmt.Sprintf("SELECT name,age,country FROM users WHERE %s = \"%s\";", "country", user_filtr.Country)
	} else if user_filtr.Name != "" {

		qtext = fmt.Sprintf("SELECT name,age,country FROM users WHERE %s = \"%s\";", "name", user_filtr.Name)
	}

	//results, err := db.Db.Query("SELECT name,age,country FROM users WHERE ? = ?;", colum, value)
	results, err := db.Db.Query(qtext)
	if err != nil {
		return nil, err
	}
	for results.Next() {
		var user model.User

		err = results.Scan(&user.Name, &user.Age, &user.Country)
		users = append(users, user)
		if err != nil {
			return nil, err
		}

	}

	return users, nil

}
func (db *UserDB) Range(age_start int, end_age int) ([]model.User, error) {
	var users []model.User
	results, err := db.Db.Query("SELECT name,age,country FROM users WHERE age >= ? and age <= ?", age_start, end_age)

	if err != nil {
		return nil, err
	}
	for results.Next() {
		var user model.User

		err = results.Scan(&user.Name, &user.Age, &user.Country)
		users = append(users, user)
		if err != nil {
			return nil, err
		}

	}

	return users, nil

}
func (db *UserDB) Order(field string, sorting string) ([]model.User, error) {
	//ORDER BY city DESC
	var users []model.User
	qtext := fmt.Sprintf("SELECT name,age,country FROM users ORDER BY %s %s", field, sorting)
	results, err := db.Db.Query(qtext)

	if err != nil {
		return nil, err
	}
	for results.Next() {
		var user model.User

		err = results.Scan(&user.Name, &user.Age, &user.Country)
		users = append(users, user)
		if err != nil {
			return nil, err
		}

	}

	return users, nil
	/*res := db.Db.Order(field + " " + sorting).Find(&users)
	return users, res.Error*/
}
func (db *UserDB) TextSearch(search string) ([]model.User, error) {
	var users []model.User
	var out_user []model.User
	results, err := db.Db.Query("SELECT name,age,country FROM users")

	if err != nil {
		return nil, err
	}
	for results.Next() {
		var user model.User

		err = results.Scan(&user.Name, &user.Age, &user.Country)
		users = append(users, user)
		if err != nil {
			return nil, err
		}

	}

	for _, person := range users {
		if strings.Contains(person.Name, search) {
			out_user = append(out_user, person)
		}
	}
	return out_user, nil
}

/*func (db *UserDB) CreateTable() {
	insert, _ := db.Db.Query("CREATE TABLE `rest-api`.`users` ( `id` INT NOT NULL AUTO_INCREMENT , `name` TEXT NOT NULL , `age` INT NOT NULL , `country` TEXT NOT NULL , PRIMARY KEY (`id`)) ENGINE = InnoDB;")
	if insert != nil {
		insert.Close()
	}

}*/
