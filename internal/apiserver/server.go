package apiserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/krizons/rest-api-golang/internal/db"
	sql_user "github.com/krizons/rest-api-golang/internal/db/sqllite"
	"github.com/krizons/rest-api-golang/internal/model"
)

type server struct {
	config *Config
	router *mux.Router
	db     db.MyDb
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
		users, err := s.db.User().GetAll()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		data, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}
func (s *server) orderHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := s.db.User().GetAll()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		data, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}
func (s *server) FilterHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var user []model.User
		var err error
		switch vars["colum"] {
		case "name":

			user, err = s.db.User().Filter(model.User{Name: vars["value"]})
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

		case "country":

			user, err = s.db.User().Filter(model.User{Country: vars["value"]})
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		case "age":
			i, err := strconv.Atoi(vars["value"])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			user, err = s.db.User().Filter(model.User{Age: uint8(i)})
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		default:
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
func (s *server) SortedHandlerUser() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		users, err := s.db.User().Order(vars["colum"], vars["trend"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
func (s *server) SortedHandlerOrder() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		users, err := s.db.Order().Sorted(vars["colum"], vars["trend"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
func (s *server) SearchHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		users, err := s.db.User().TextSearch(vars["name"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
func (s *server) configureRouter() error {

	if err := s.router.HandleFunc("/users/all", s.userHandler()).Methods("GET").GetError(); err != nil {
		return err
	}
	if err := s.router.HandleFunc("/orders/all", s.orderHandler()).Methods("GET").GetError(); err != nil {
		return err
	}
	if err := s.router.HandleFunc("/users/filter/{colum:(?:age|name|country)}/{value}", s.FilterHandler()).Methods("GET").GetError(); err != nil {
		return err
	}
	if err := s.router.HandleFunc("/users/sorted/{colum:(?:age|name|country)}/{trend:(?:asc|desc)}", s.SortedHandlerUser()).Methods("GET").GetError(); err != nil {
		return err
	}
	if err := s.router.HandleFunc("/order/sorted/{colum:(?:name|price|count)}/{trend:(?:asc|desc)}", s.SortedHandlerOrder()).Methods("GET").GetError(); err != nil {
		return err
	}

	if err := s.router.HandleFunc("/users/search/{name}", s.SearchHandler()).Methods("GET").GetError(); err != nil {
		return err
	}

	return nil
}
func (s *server) Start() error {
	s.configureRouter()
	if err := s.configureDB(); err != nil {
		return err
	}

	return http.ListenAndServe(s.config.BindAddr, s.router)
}
func (s *server) configureDB() error {
	db, error := sql_user.New(s.config.DatabaseURL)
	if error != nil {
		return error
	}

	s.db = db
	return nil
}
