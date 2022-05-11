package sqlite_user_test

import (
	"testing"

	"github.com/krizons/rest-api-golang/internal/model"
	sql_user "github.com/krizons/rest-api-golang/internal/user_db/sqllite"
	"github.com/stretchr/testify/assert"
)

func TestCreateGetSelect(t *testing.T) {
	assert := assert.New(t)
	test_db, err := sql_user.New("../../../db_test.db")
	assert.NoError(err)
	test_db.Create(model.User{Name: "Margie Ellison", Age: 18, Country: "UK"})
	test_db.Create(model.User{Name: "Gordon Felix", Age: 20, Country: "US"})
	test_db.Create(model.User{Name: "Johnnie Burton", Age: 45, Country: "US"})
	test_db.Create(model.User{Name: "Lillie Perkins", Age: 18, Country: "AT"})
	test_db.Create(model.User{Name: "Moses Jarvis", Age: 64, Country: "AT"})
	test_db.Create(model.User{Name: "Margie Ellison", Age: 18, Country: "UK"})
	test_db.Create(model.User{Name: "Gordon Felix", Age: 20, Country: "US"})
	test_db.Create(model.User{Name: "Johnnie Burton", Age: 45, Country: "US"})
	test_db.Create(model.User{Name: "Lillie Perkins", Age: 18, Country: "AT"})
	test_db.Create(model.User{Name: "Moses Jarvis", Age: 64, Country: "AT"})
	users, err := test_db.GetAll()
	assert.NoError(err)
	assert.NotEqual(len(users), 0)

}
func TestFindParam(t *testing.T) {
	assert := assert.New(t)
	test_db, err := sql_user.New("../../../db_test.db")
	assert.NoError(err)
	filter_out, err := test_db.Filter(model.User{Country: "AT"})
	assert.NoError(err)
	assert.NotEqual(len(filter_out), 0)
	filter_out, err = test_db.Filter(model.User{Country: "UK"})
	assert.NoError(err)
	assert.NotEqual(len(filter_out), 0)
	filter_out, err = test_db.Filter(model.User{Age: 18})
	assert.NoError(err)
	assert.NotEqual(len(filter_out), 0)

}
func TestRange(t *testing.T) {
	assert := assert.New(t)
	test_db, err := sql_user.New("../../../db_test.db")
	assert.NoError(err)
	range_out, err := test_db.Range(18, 30)
	assert.NoError(err)
	for _, person := range range_out {
		assert.LessOrEqual(uint8(18), person.Age)
		assert.GreaterOrEqual(uint8(30), person.Age)

	}
}
func TestOrder(t *testing.T) {
	assert := assert.New(t)
	test_db, err := sql_user.New("../../../db_test.db")
	assert.NoError(err)
	range_out, err := test_db.Order("Age", "desc")
	assert.NoError(err)
	var tmp uint8 = 255
	for _, person := range range_out {
		assert.GreaterOrEqual(tmp, person.Age)
		tmp = person.Age

	}
	range_out, err = test_db.Order("Age", "ASC")
	assert.NoError(err)
	tmp = 0
	for _, person := range range_out {
		assert.LessOrEqual(tmp, person.Age)
		tmp = person.Age

	}
}
func TestTextFins(t *testing.T) {
	assert := assert.New(t)
	test_db, err := sql_user.New("../../../db_test.db")
	assert.NoError(err)
	range_out, err := test_db.TextSearch("Margie")
	assert.NoError(err)

	for _, person := range range_out {
		assert.Contains(person.Name, "Margie")

	}

}
