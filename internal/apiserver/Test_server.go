package apiserver

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	sql_user "github.com/krizons/rest-api-golang/internal/sql"
	"github.com/stretchr/testify/assert"
)

func TestUser_handle(t *testing.T) {
	assert := assert.New(t)
	s := New(&Config{DB: &sql_user.Config{DatabaseURL: "db_test.db"}})
	err := s.configureDB()
	assert.NoError(err)
	s.configureRouter()
	b := &bytes.Buffer{}
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/users/all", b)
	s.ServeHTTP(rec, req)
	assert.NotEqual(http.StatusBadRequest, rec.Code)
}
