package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "social-network/docs"
	"social-network/internal/app/config"
	"social-network/internal/model"
	"social-network/internal/store"
	"social-network/pkg/client"
	"social-network/pkg/jwttoken"
	"social-network/pkg/router"

	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	jwtKey                 = "JWT_KEY"
	region          = "us-east-1"
	bucketName      = "profilemediabucket-voyage"
	awsAccessKey    = "AWS_ACCESS_KEY"
	awsSecretKey    = "AWS_SECRET_KEY"
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
	router    *router.Router
	logger    *log.Logger
	store     store.Store
	types     model.Type
	wsClient  client.ChatClient
	wsService *ChatService
}

type Option func(*config.Config)

func newServer(store store.Store, opts ...Option) *Server {
	config := &config.Config{}

	// Apply options
	for _, opt := range opts {
		opt(config)
	}

	s := &Server{
		router:    router.New(),
		logger:    log.New(os.Stdout, "", 0),
		store:     store,
		types:     model.InitializeTypes(),
		wsClient:  client.NewClient(config.ChatServiceURL),
		wsService: NewChatServer(store),
	}

	configureRouter(s)

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func configureRouter(s *Server) {
	// Temporary
	s.router.OPTION("/", s.corsQuickFix())

	///
	s.router.Use(s.setRequestID, s.logRequest, s.CORSMiddleware)
	s.router.UseWithPrefix("auth", s.jwtMiddleware)

	s.router.GET("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://ec2-3-84-51-36.compute-1.amazonaws.com:8080/swagger/doc.json"),
	))
	//---------USER---------//
	s.router.POST("/api/v1/user/create", s.userCreate())
	s.router.POST("/api/v1/user/login", s.userLogin())
	s.router.GET("/api/v1/auth/user/logout", s.userLogout())
	s.router.GET("/api/v1/auth/user/privacy/{privacy_state}", s.userPrivacy())
	s.router.PUT("/api/v1/auth/user/description", s.userDescription())
	s.router.PUT("/api/v1/auth/user/cover", s.userCover())
	s.router.PUT("/api/v1/auth/user/avatar", s.userAvatar())
	s.router.GET("/api/v1/auth/user/followers/{id}", s.userFollowers())
	s.router.GET("/api/v1/auth/user/following/{id}", s.userFollowing())
	s.router.GET("/api/v1/auth/user/posts/{id}", s.userPosts())
	s.router.GET("/api/v1/auth/user/notifications", s.userNotifications())
	s.router.GET("/api/v1/auth/user/all", s.userGetAll())
	s.router.GET("/api/v1/auth/user/get/{id}", s.userGet())
	s.router.GET("/api/v1/auth/user/current", s.currentUser())
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
	s.router.PUT("/api/v1/auth/posts/update", s.updatePost())
	s.router.DELETE("/api/v1/auth/posts/delete/{id}", s.deletePost())
	s.router.POST("/api/v1/auth/posts/selected/add", s.addSelected())
	s.router.POST("/api/v1/auth/posts/selected/delete", s.deleteSelected())
	s.router.GET("/api/v1/auth/posts/feed", s.getFeed())
	//---------COMMENT------------//
	s.router.POST("/api/v1/auth/comment/create", s.createComment())
	s.router.PUT("/api/v1/auth/comment/update", s.updateComment())
	s.router.DELETE("/api/v1/auth/comment/delete/{id}", s.deleteComment())
	s.router.GET("/api/v1/auth/comment/{id}", s.getComments())
	//---------GROUP--------------//
	s.router.POST("/api/v1/auth/group/create", s.groupCreate())
	s.router.PUT("/api/v1/auth/group/update", s.groupUpdate())
	s.router.DELETE("/api/v1/auth/group/delete/{id}", s.groupDelete())
	s.router.GET("/api/v1/auth/group/{id}", s.groupGet())
	s.router.GET("/api/v1/auth/group", s.groupGetAll())
	s.router.GET("/api/v1/auth/group/posts/{id}", s.groupGetPost())
	s.router.POST("/api/v1/auth/group/invite", s.groupInvite())
	s.router.POST("/api/v1/auth/group/request", s.groupInviteRequest())
	s.router.GET("/api/v1/auth/group/join/{id}", s.joinGroup())
	//---------EVENT--------------//
	s.router.POST("/api/v1/auth/group/event/create", s.createEvent())
	s.router.PUT("/api/v1/auth/group/event/update", s.updateEvent())
	s.router.DELETE("/api/v1/auth/group/event/delete/{id}", s.deleteEvent())
	s.router.GET("/api/v1/auth/group/event/{id}", s.getEvent())
	s.router.GET("/api/v1/auth/group/event/{id}/register/{opt}", s.registerEvent())
	s.router.GET("/api/v1/auth/group/{id}/event/all", s.getGroupEvents())
	//---------CHATS--------------//
	s.router.POST("/api/v1/auth/chats/create", s.createChat())
	s.router.POST("/api/v1/auth/chats/add/user", s.addUserChat())
	s.router.POST("/api/v1/auth/chats/add/line", s.addLineChat())
	s.router.GET("/api/v1/auth/chats", s.getAllChats())
	s.router.GET("/api/v1/auth/chats/get/users/{id}", s.getAllChatUsers())
	s.router.GET("/api/v1/auth/chats/{id}", s.getChatLines())
	//----------IMAGES-S3-------------//
	s.router.POST("/api/v1/auth/images/{parent_id}", s.imageUpload())
	//--WEBSOCKET--//
	s.router.GET("/ws", s.wsHandler())
	s.router.GET("/auth/ws", s.wsService.HandleWS)

	s.router.GET("/login/{id}", s.login())
}

func (s *Server) wsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := s.wsClient.Connect(w, r); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}
		s.respond(w, http.StatusOK, nil)
	}
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

// FOR DEVELOPMENT PURPOSES TO GENERATE JWT TOKEN
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
