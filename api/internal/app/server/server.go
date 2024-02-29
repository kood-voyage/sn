package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "social-network/docs"
	"social-network/internal/store"
	"social-network/pkg/router"

	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	sessionName     = "session"
	jwtKey          = "JWT_KEY"
	ctxKeyRequestID = iota
)

type ctxKey int

const ctxUserID ctxKey = 1

type Response struct {
	Data interface{} `json:"data"`
}

type Error struct {
	Error interface{} `json:"error"`
}

type Server struct {
	router *router.Router
	logger *log.Logger
	store  store.Store
}

func newServer(store store.Store) *Server {
	s := &Server{
		router: router.New(),
		logger: log.Default(),
		store:  store,
	}

	configureRouter(s)

	return s
}


func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func configureRouter(s *Server) {
	s.router.Use(s.setRequestID, s.logRequest, s.CORSMiddleware)

	s.router.GET("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))
	s.router.POST("/api/v1/users/create", s.createUser())
	s.router.GET("/api/v1/follow/{id}", s.handleFollow())
}

func (s *Server) error(w http.ResponseWriter, code int, err error) {
	s.respond(w, code, Error{err.Error()})
}

func (s *Server) respond(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *Server) decode(r *http.Request, data interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return fmt.Errorf("decode json: %w", err)
	}
	return nil
}
