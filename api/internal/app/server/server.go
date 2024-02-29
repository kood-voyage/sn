package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"social-network/internal/store"
	"social-network/pkg/router"
)

const (
	sessionName     = "session"
	jwtKey          = "JWT_KEY"
	ctxKeyRequestID = iota
	ctxUserID
)

type Response struct {
	Data interface{} `json:"data"`
}

type server struct {
	router *router.Router
	logger *log.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: router.New(),
		logger: log.Default(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(s.setRequestID, s.logRequest, s.CORSMiddleware)

	s.router.GET("/", s.testRoute())
	s.router.POST("/api/v1/users/create", s.createUser())
}

func (s *server) testRoute() http.HandlerFunc {
	type respond struct {
		Message string `json:"message"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, respond{
			Message: "Hello World!",
		})
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *server) decode(r *http.Request, data interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return fmt.Errorf("decode json: %w", err)
	}
	return nil
}
