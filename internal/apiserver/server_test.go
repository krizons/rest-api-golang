package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/gorilla/mux"
	"github.com/krizons/rest-api-golang/internal/model"
	sql_user "github.com/krizons/rest-api-golang/internal/sql"
	"github.com/stretchr/testify/assert"
)

func TestUser_handle(t *testing.T) {
	assert := assert.New(t)
	s := New(&Config{DB: &sql_user.Config{DatabaseURL: "../../db_test.db"}})
	err := s.configureDB()
	assert.NoError(err)
	s.configureRouter()
	b := &bytes.Buffer{}
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/users/all", b)
	s.ServeHTTP(rec, req)
	assert.NotEqual(http.StatusBadRequest, rec.Code)
	var users []model.User
	err = json.Unmarshal(rec.Body.Bytes(), &users)
	assert.NotEqual(len(users), 0)
	assert.NoError(err)

}
func TestFilter_handle(t *testing.T) {
	assert := assert.New(t)
	s := New(&Config{DB: &sql_user.Config{DatabaseURL: "../../db_test.db"}})
	err := s.configureDB()
	assert.NoError(err)
	s.configureRouter()
	test_case_done := []struct {
		colum string
		value string
	}{{colum: "name", value: "Margie Ellison"}, {colum: "name", value: "Gordon Felix"}, {colum: "name", value: "Lillie Perkins"},
		{colum: "country", value: "US"}, {colum: "country", value: "UK"}, {colum: "country", value: "AT"},
		{colum: "age", value: "20"}, {colum: "age", value: "18"}, {colum: "age", value: "45"}}
	for _, val := range test_case_done {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/users/filter/"+val.colum+"/"+val.value, nil)
		s.ServeHTTP(rec, req)
		assert.NotEqual(http.StatusBadRequest, rec.Code)
		var users []model.User
		err = json.Unmarshal(rec.Body.Bytes(), &users)
		assert.NotEqual(len(users), 0)
		assert.NoError(err)
	}
	test_case_zero := []struct {
		colum string
		value string
	}{{colum: "name", value: "TEST"},
		{colum: "country", value: "TEST"},
		{colum: "age", value: "200"}}
	for _, val := range test_case_zero {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/users/filter/"+val.colum+"/"+val.value, nil)
		s.ServeHTTP(rec, req)
		assert.NotEqual(http.StatusBadRequest, rec.Code)
		var users []model.User
		err = json.Unmarshal(rec.Body.Bytes(), &users)
		assert.Equal(len(users), 0)
		assert.NoError(err)
	}
	test_case_bad := []struct {
		colum string
		value string
	}{{colum: "age", value: "TEST"}}
	for _, val := range test_case_bad {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/users/filter/"+val.colum+"/"+val.value, nil)
		s.ServeHTTP(rec, req)
		assert.NotEqual(http.StatusOK, rec.Code)
		var users []model.User
		err = json.Unmarshal(rec.Body.Bytes(), &users)
		assert.Equal(len(users), 0)
		assert.Error(err)
	}
	test_case_not_found := []struct {
		colum string
		value string
	}{{colum: "names", value: "TEST"}}
	for _, val := range test_case_not_found {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/users/filter/"+val.colum+"/"+val.value, nil)
		s.ServeHTTP(rec, req)
		assert.Equal(http.StatusNotFound, rec.Code)
		var users []model.User
		err = json.Unmarshal(rec.Body.Bytes(), &users)
		assert.Equal(len(users), 0)
		assert.Error(err)
	}
}
func TestSorted_handle(t *testing.T) {
	assert := assert.New(t)
	s := New(&Config{DB: &sql_user.Config{DatabaseURL: "../../db_test.db"}})
	err := s.configureDB()
	assert.NoError(err)
	err = s.configureRouter()
	assert.NoError(err)
	test_case_done := []struct {
		colum string
		trend string
	}{{colum: "age", trend: "asc"}, {colum: "age", trend: "desc"}}
	for _, val := range test_case_done {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/users/sorted/"+val.colum+"/"+val.trend, nil)
		s.ServeHTTP(rec, req)
		assert.Equal(http.StatusOK, rec.Code)
		var users []model.User
		err = json.Unmarshal(rec.Body.Bytes(), &users)
		assert.NotEqual(len(users), 0)
		assert.NoError(err)
		if val.trend == "desc" {
			var tmp uint8 = 255
			for _, person := range users {
				assert.GreaterOrEqual(tmp, person.Age)
				tmp = person.Age

			}
		} else if val.trend == "asc" {
			var tmp uint8 = 0
			for _, person := range users {
				assert.LessOrEqual(tmp, person.Age)
				tmp = person.Age

			}
		}

	}
	test_case_bad := []struct {
		colum string
		trend string
	}{{colum: "age", trend: "test"}, {colum: "test", trend: "desc"},
		{colum: "agess", trend: "desc"}, {colum: "age", trend: "descs"},
		{colum: "ege", trend: "dess"}, {colum: "age", trend: "aesc"}}
	for _, val := range test_case_bad {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/users/sorted/"+val.colum+"/"+val.trend, nil)
		s.ServeHTTP(rec, req)
		assert.Equal(http.StatusNotFound, rec.Code)
		var users []model.User
		err = json.Unmarshal(rec.Body.Bytes(), &users)
		assert.Equal(len(users), 0)
		assert.Error(err)

	}
}
func TestSearch_handle(t *testing.T) {
	assert := assert.New(t)
	s := New(&Config{DB: &sql_user.Config{DatabaseURL: "../../db_test.db"}})
	err := s.configureDB()
	assert.NoError(err)
	err = s.configureRouter()
	assert.NoError(err)
	test_case_done := []struct {
		name string
	}{{name: "Margie Ellison"}, {name: "Gordon Felix"}}
	for _, val := range test_case_done {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/users/search/"+val.name, nil)
		s.ServeHTTP(rec, req)
		assert.Equal(http.StatusOK, rec.Code)
		var users []model.User
		err = json.Unmarshal(rec.Body.Bytes(), &users)
		assert.NotEqual(len(users), 0)
		assert.NoError(err)

		for _, person := range users {
			assert.GreaterOrEqual(val.name, person.Name)

		}

	}
	test_case_bad := []struct {
		name string
	}{{name: "Test1"}, {name: "Test2"}}
	for _, val := range test_case_bad {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/users/search/"+val.name, nil)
		s.ServeHTTP(rec, req)
		assert.Equal(http.StatusOK, rec.Code)
		var users []model.User
		err = json.Unmarshal(rec.Body.Bytes(), &users)
		assert.Equal(len(users), 0)
		assert.NoError(err)

	}
}
func TestXxx(t *testing.T) {
	assert := assert.New(t)
	matched, err := regexp.Match("age|name|country", []byte("age"))
	mux.NewRouter().NewRoute().GetPathRegexp()
	if err != nil {
		assert.NoError(err)
	}

	assert.Equal(matched, true)
}
