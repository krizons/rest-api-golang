package sql_user_test

import (
	"testing"

	"github.com/krizons/rest-api-golang/internal/model"
	sql_user "github.com/krizons/rest-api-golang/internal/sql"
	"github.com/stretchr/testify/assert"
)

func TestCreateGetSelect(t *testing.T) {
	assert := assert.New(t)
	test_db, err := sql_user.New("tmp.db")
	assert.NoError(err)
	sql_user.Create(test_db, model.User{Name: "Margie Ellison", Age: 18, Country: "UK"})
	sql_user.Create(test_db, model.User{Name: "Gordon Felix", Age: 20, Country: "US"})
	sql_user.Create(test_db, model.User{Name: "Johnnie Burton", Age: 45, Country: "US"})
	sql_user.Create(test_db, model.User{Name: "Lillie Perkins", Age: 18, Country: "AT"})
	sql_user.Create(test_db, model.User{Name: "Moses Jarvis", Age: 64, Country: "AT"})
	sql_user.Create(test_db, model.User{Name: "Margie Ellison", Age: 18, Country: "UK"})
	sql_user.Create(test_db, model.User{Name: "Gordon Felix", Age: 20, Country: "US"})
	sql_user.Create(test_db, model.User{Name: "Johnnie Burton", Age: 45, Country: "US"})
	sql_user.Create(test_db, model.User{Name: "Lillie Perkins", Age: 18, Country: "AT"})
	sql_user.Create(test_db, model.User{Name: "Moses Jarvis", Age: 64, Country: "AT"})
	users, err := sql_user.GetAll(test_db)
	assert.NoError(err)
	assert.NotEqual(len(users), 0)

}
func TestFindParam(t *testing.T) {
	assert := assert.New(t)
	test_db, err := sql_user.New("tmp.db")
	assert.NoError(err)
	filter_out, err := sql_user.Filter(test_db, model.User{Country: "AT"})
	assert.NoError(err)
	assert.NotEqual(len(filter_out), 0)
	filter_out, err = sql_user.Filter(test_db, model.User{Country: "UK"})
	assert.NoError(err)
	assert.NotEqual(len(filter_out), 0)
	filter_out, err = sql_user.Filter(test_db, model.User{Age: 18})
	assert.NoError(err)
	assert.NotEqual(len(filter_out), 0)

}
func TestRange(t *testing.T) {
	assert := assert.New(t)
	test_db, err := sql_user.New("tmp.db")
	assert.NoError(err)
	range_out, err := sql_user.Range(test_db, 18, 30)
	assert.NoError(err)
	for _, person := range range_out {
		assert.LessOrEqual(int8(18), person.Age)
		assert.GreaterOrEqual(int8(30), person.Age)

	}
}
