package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	_ "social-network/docs"
	"social-network/internal/model"
	"social-network/internal/store"
	"social-network/pkg/jwttoken"
	"social-network/pkg/router"
	"time"

	"github.com/swaggo/http-swagger"
)

const (
	jwtKey                 = "JWT_KEY"
	ctxKeyRequestID ctxKey = iota
	ctxUserID
)

type ctxKey int

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
	types  model.Type
}

func newServer(store store.Store) *Server {
	s := &Server{
		router: router.New(),
		logger: log.Default(),
		store:  store,
		types:  model.InitializeTypes(),
	}

	configureRouter(s)

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func configureRouter(s *Server) {
	s.router.Use(s.setRequestID, s.logRequest, s.CORSMiddleware)
	s.router.UseWithPrefix("auth", s.jwtMiddleware)

	s.router.GET("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://ec2-54-91-167-36.compute-1.amazonaws.com:8080/swagger/doc.json"),
	))
	//---------USER---------//
	s.router.GET("/api/v1/auth/user/create/{privacy_state}", s.userCreate())
	s.router.GET("/api/v1/auth/user/privacy/{privacy_state}", s.userPrivacy())
	s.router.GET("/api/v1/auth/user/followers/{id}", s.userFollowers())
	s.router.GET("/api/v1/auth/user/following/{id}", s.userFollowing())
	s.router.GET("/api/v1/auth/user/posts/{id}", s.userPosts())
	s.router.GET("/api/v1/auth/user/notifications", s.userNotifications())
	//---------NOTIFICATION---------//
	s.router.POST("/api/v1/auth/notification/create", s.notificationCreate())
	s.router.DELETE("/api/v1/auth/notification/delete/{id}", s.notificationDelete())
	//---------FOLLOW--------------//
	s.router.GET("/api/v1/auth/follow/{id}", s.handleFollow())
	s.router.GET("/api/v1/auth/unfollow/{id}", s.handleUnfollow())
	s.router.POST("/api/v1/auth/follow/request", s.handleFollowRequest())
	//---------POST---------------//
	s.router.GET("/api/v1/auth/posts/{id}", s.getPost())
	s.router.POST("/api/v1/auth/posts/create", s.createPost())
	s.router.DELETE("/api/v1/auth/posts/delete/{id}", s.deletePost())
	s.router.POST("/api/v1/auth/posts/selected/add", s.addSelected())
	s.router.POST("/api/v1/auth/posts/selected/delete", s.removeSelected())
	//---------COMMENT------------//
	s.router.POST("/api/v1/auth/comment/create", s.createComment())
	s.router.DELETE("/api/v1/auth/comment/delete/{id}", s.deleteComment())
	s.router.GET("/api/v1/auth/comment/{id}", s.getComments())
	//---------GROUP--------------//
	s.router.POST("/api/v1/auth/group/create", s.groupCreate())
	s.router.PUT("/api/v1/auth/group/update", s.groupUpdate())
	s.router.DELETE("/api/v1/auth/group/delete/{id}", s.groupDelete())
	s.router.GET("/api/v1/auth/group/{id}", s.groupGet())
	s.router.POST("/api/v1/auth/group/invite", s.groupInvite())
	s.router.POST("/api/v1/auth/group/request", s.groupInviteRequest())

	s.router.GET("/login/{id}", s.login())

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

//FOR DEVELOPMENT PURPOSES TO GENERATE JWT TOKEN

func (s *Server) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newToken := jwttoken.NewClaims()
		newToken.Set("user_id", r.PathValue("id"))
		newToken.SetTime("exp", time.Now().Add(time.Hour*100))
		a := jwttoken.HmacSha256(os.Getenv(jwtKey))

		token, err := a.Encode(newToken)
		if err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		s.respond(w, http.StatusOK, Response{Data: token})
	}
}
