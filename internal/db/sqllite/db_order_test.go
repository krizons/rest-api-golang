package sqlite_test

import (
	"testing"

	sql_user "github.com/krizons/rest-api-golang/internal/db/sqllite"
	"github.com/krizons/rest-api-golang/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateGetSelectOrder(t *testing.T) {
	assert := assert.New(t)
	test_db, err := sql_user.New("../../../db_test.db")
	assert.NoError(err)
	test_db.Order().Create(model.Order{Name: "Test1", Price: 1000, Count: 10})
	test_db.Order().Create(model.Order{Name: "Test2", Price: 200, Count: 20})
	test_db.Order().Create(model.Order{Name: "Test3", Price: 300, Count: 30})
	test_db.Order().Create(model.Order{Name: "Test4", Price: 400, Count: 40})

	users, err := test_db.Order().GetAll()
	assert.NoError(err)
	assert.NotEqual(len(users), 0)

}
func TestFindParamOrder(t *testing.T) {
	assert := assert.New(t)
	test_db, err := sql_user.New("../../../db_test.db")
	assert.NoError(err)
	filter_out, err := test_db.Order().Filter(model.Order{Name: "Test4"})
	assert.NoError(err)
	assert.NotEqual(len(filter_out), 0)
	filter_out, err = test_db.Order().Filter(model.Order{Price: 400})
	assert.NoError(err)
	assert.NotEqual(len(filter_out), 0)
	filter_out, err = test_db.Order().Filter(model.Order{Count: 40})
	assert.NoError(err)
	assert.NotEqual(len(filter_out), 0)

}

func Sorted(t *testing.T) {
	assert := assert.New(t)
	test_db, err := sql_user.New("../../../db_test.db")
	assert.NoError(err)
	range_out, err := test_db.Order().Sorted("price", "desc")
	assert.NoError(err)
	var tmp int = 255
	for _, person := range range_out {
		assert.GreaterOrEqual(tmp, person.Price)
		tmp = person.Price

	}
	range_out, err = test_db.Order().Sorted("price", "ASC")
	assert.NoError(err)
	tmp = 0
	for _, person := range range_out {
		assert.LessOrEqual(tmp, person.Price)
		tmp = person.Price

	}
}
