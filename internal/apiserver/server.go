package apiserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/krizons/rest-api-golang/internal/db"
	sql_user "github.com/krizons/rest-api-golang/internal/db/sqllite"
	"github.com/krizons/rest-api-golang/internal/model"
)

var sessionName = "test"

type server struct {
	config *Config
	router *mux.Router
	db     db.MyDb

	sesison sessions.Store
}

var key = []byte("my_secret")

func New(config *Config) *server {
	return &server{
		config:  config,
		router:  mux.NewRouter(),
		sesison: sessions.NewCookieStore(key),
	}
}
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
func (s *server) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sesison.Get(r, sessionName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
func (s *server) Login(w http.ResponseWriter, r *http.Request) {
	session, _ := s.sesison.Get(r, sessionName)
	session.Values["authenticated"] = true
	session.Save(r, w)
	w.Write([]byte("login ok"))
}
func (s *server) userHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := s.db.User().GetAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		data, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
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
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		data, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
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
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

		case "country":

			user, err = s.db.User().Filter(model.User{Country: vars["value"]})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		case "age":
			i, err := strconv.Atoi(vars["value"])
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			user, err = s.db.User().Filter(model.User{Age: uint8(i)})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		default:
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
func (s *server) SortedHandlerUser() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		users, err := s.db.User().Order(vars["colum"], vars["trend"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
func (s *server) SortedHandlerOrder() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		users, err := s.db.Order().Sorted(vars["colum"], vars["trend"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
func (s *server) SearchHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		users, err := s.db.User().TextSearch(vars["name"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
func (s *server) configureRouter() error {
	private := s.router.PathPrefix("/users").Subrouter()
	private.Use(s.authenticate)

	if err := private.HandleFunc("/all", s.userHandler()).Methods("GET").GetError(); err != nil {
		return err
	}
	if err := s.router.HandleFunc("/orders/all", s.orderHandler()).Methods("GET").GetError(); err != nil {
		return err
	}
	if err := private.HandleFunc("/filter/{colum:(?:age|name|country)}/{value}", s.FilterHandler()).Methods("GET").GetError(); err != nil {
		return err
	}
	if err := private.HandleFunc("/sorted/{colum:(?:age|name|country)}/{trend:(?:asc|desc)}", s.SortedHandlerUser()).Methods("GET").GetError(); err != nil {
		return err
	}
	if err := s.router.HandleFunc("/order/sorted/{colum:(?:name|price|count)}/{trend:(?:asc|desc)}", s.SortedHandlerOrder()).Methods("GET").GetError(); err != nil {
		return err
	}

	if err := private.HandleFunc("/search/{name}", s.SearchHandler()).Methods("GET").GetError(); err != nil {
		return err
	}
	if err := s.router.HandleFunc("/login", s.Login).Methods("GET").GetError(); err != nil {
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
