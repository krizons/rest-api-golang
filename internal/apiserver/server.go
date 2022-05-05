package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	sql_user "github.com/krizons/rest-api-golang/internal/sql"
)

type server struct {
	config *Config
	router *mux.Router
	db     *sql_user.MyDb
}

func New(config *Config) *server {
	return &server{
		config: config,
		router: mux.NewRouter(),
	}
}
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
func (s *server) userHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := sql_user.GetAll(s.db)
		if err != nil {
			fmt.Fprint(w, http.StatusBadRequest)
			return
		}
		data, err := json.Marshal(users)
		if err != nil {
			fmt.Fprint(w, http.StatusBadRequest)
			return
		}
		fmt.Fprint(w, data)
		//json_out =: json.ma
	}
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/users/all", s.userHandler())
}
func (s *server) Start() error {
	s.configureRouter()
	if err := s.configureDB(); err != nil {
		return err
	}

	return http.ListenAndServe(s.config.BindAddr, s.router)
}
func (s *server) configureDB() error {
	db, error := sql_user.New(s.config.DB)
	if error != nil {
		return error
	}

	s.db = db
	return nil
}
