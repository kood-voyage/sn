package server

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"social-network/internal/model"
	"social-network/pkg/validator"

	"github.com/google/uuid"
)

type ValidateStruct struct {
	Privacy string `validate:"lowercase|contains:public,private"`
}

// userCreate Handles the user creation to database. Only userID and privacy state will be stored.
//
// @Summary Create a user with privacy state
// @Tags users
// @Produce json
// @Param privacy_state path string true "Only public, private allowed"
// @Success 201 {object} model.User
// @Failure 401 {object} Error
// @Failure 422 {object} Error
// @Router /api/v1/auth/user/create/{privacy_state} [get]
func (s *Server) userCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		if err := s.decode(r, &user); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		if err := validator.Validate(user); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		privacy, ok := s.types.Privacy.Values[user.Privacy]

		fmt.Println(user)
		if !ok {
			s.error(w, http.StatusUnprocessableEntity, errors.New("public, private, selected states are allowed"))
			return
		}

		u, err := s.store.User().Create(user, privacy)
		if err != nil {

			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusCreated, Response{Data: u})
	}
}

func (s *Server) userLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginDetails struct {
			Login    string `json:"login" validate:"required"`
			Password string `json:"password" validate:"required"`
		}

		if err := s.decode(r, &loginDetails); err != nil {
			s.error(w, http.StatusUnauthorized, err)
			return
		}
		if err := validator.Validate(loginDetails); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}
		user, err := s.store.User().Get(loginDetails.Login)
		if err != nil {
			if err == sql.ErrNoRows {
				s.error(w, http.StatusUnprocessableEntity, errors.New("no user found with specific login details"))
				return
			}
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		if !user.ComparePassword(loginDetails.Password) {
			s.error(w, http.StatusUnauthorized, errors.New("password incorrect"))
			return
		}

		accesstoken_id := uuid.New().String()
		accessToken, err := NewAccessToken(accesstoken_id, user.ID, time.Now().Add(15*time.Minute))
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		refreshToken, err := NewRefreshToken(accesstoken_id, time.Now().Add(24*7*time.Hour))
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		session := model.Session{
			AcessID:   accesstoken_id,
			UserID:    user.ID,
			CreatedAT: time.Now().Add(24 * 7 * time.Hour),
		}

		oldSession, err := s.store.Session().CheckByUserId(user.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				_, err = s.store.Session().Create(session)
				if err != nil {
					s.error(w, http.StatusUnprocessableEntity, err)
					return
				}
				http.SetCookie(w, accessToken)
				http.SetCookie(w, refreshToken)
				user.Sanitize()
				s.respond(w, http.StatusOK, Response{Data: user})
				return
			}
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}
		_, err = s.store.Session().Update(oldSession.AcessID, session)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		http.SetCookie(w, accessToken)
		http.SetCookie(w, refreshToken)
		user.Sanitize()
		s.respond(w, http.StatusOK, Response{Data: user})
	}
}

func (s *Server) userLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		err := s.store.Session().DeleteByUser(userID)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}
		http.SetCookie(w, DeleteAccessToken())
		http.SetCookie(w, DeleteRefreshToken())
		s.respond(w, http.StatusOK, Response{Data: "Successfully logged out"})
	}
}

// userPrivacy Will update the user's privacy.
//
// @Summary Updates user's privacy
// @Tags users
// @Produce json
// @Param privacy_state path string true "Only public, private, selected allowed"
// @Success 200 {object} model.User
// @Failure 400 {object} Error
// @Failure 401 {object} Error
// @Failure 422 {object} Error
// @Router /api/v1/auth/user/privacy/{privacy_state} [get]
func (s *Server) userPrivacy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		user := &model.User{ID: userID}

		privacy_state := r.PathValue("privacy_state")
		privacy, ok := s.types.Privacy.Values[privacy_state]
		if !ok {
			s.error(w, http.StatusBadRequest, errors.New("public, private states are allowed"))
			return
		}

		if err := validator.Validate(ValidateStruct{
			Privacy: privacy_state,
		}); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		if err := s.store.Privacy().Update(user.ID, privacy); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, Response{Data: user})
	}
}

// userFollowers Handles the retrieval of provided users followers.
//
// @Summary Returns a list of user followers
// @Tags users
// @Produce json
// @Param id path string true "User id to get followers"
// @Success 200 {object} []model.User
// @Failure 422 {object} Error
// @Router /api/v1/auth/user/{id}/posts [get]
func (s *Server) userFollowers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		followers, err := s.store.User().GetFollowers(r.PathValue("id"))
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, http.StatusOK, Response{Data: followers})
	}
}

// userFollowing Handles user own following.
//
// @Summary Return a list of users who is user following
// @Tags users
// @Produce json
// @Param id path string true "User id to get followers"
// @Success 200 {object} []model.User
// @Failure 422 {object} Error
// @Router /api/v1/auth/user/following/{id} [get]
func (s *Server) userFollowing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		followers, err := s.store.User().GetFollowing(r.PathValue("id"))
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, http.StatusOK, Response{Data: followers})
	}
}

// NEED TO WRITE TEST FOR THAT FUNCTION AFTER FINISHING WITH POST PRIVACY CHECKING

// userPosts Handles getting user profile posts.
//
// @Summary Return a list of posts what user has created
// @Tags users
// @Produce json
// @Param user_id path string true "User id to get user's posts"
// @Success 200 {object} []model.Post
// @Failure 422 {object} Error
// @Router /api/v1/auth/user/posts/{id} [get]
func (s *Server) userPosts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID_request, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("unauthorized user"))
			return
		}

		// if userID_request != r.PathValue("id") {
		// 	//check requested user profile privacy
		// 	privacy, err := s.store.User().CheckPrivacy(r.PathValue("id"))
		// 	if err != nil {
		// 		s.error(w, http.StatusUnprocessableEntity, err)
		// 		return
		// 	}
		// 	if privacy == s.types.Privacy.Private {
		// 		//check if user follows
		// 		userFollows, err := s.store.User().IsFollowing(userID_request, r.PathValue("id"))
		// 		if err != nil {
		// 			s.error(w, http.StatusUnprocessableEntity, err)
		// 			return
		// 		}
		// 		if !userFollows {
		// 			s.error(w, http.StatusForbidden, errors.New("users profile is private"))
		// 			return
		// 		}
		// 	}
		// }

		posts, err := s.store.Post().GetUsers(userID_request, r.PathValue("id"))
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, Response{Data: posts})
	}
}

// userNotifications Handles users notifications.
//
// @Summary Return a list of notifications to user
// @Tags users
// @Produce json
// @Success 200 {object} []model.Request
// @Failure 401 {object} Error
// @Failure 422 {object} Error
// @Router /api/v1/auth/user/notifications [get]
func (s *Server) userNotifications() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		notifications, err := s.store.User().GetNotifications(userID)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, Response{Data: notifications})
	}
}

func (s *Server) userDescription() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		var response struct {
			Description string `json:"description" validate:"required|min_len:4"`
		}
		if err := s.decode(r, &response); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}
		if err := validator.Validate(response); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}
		err := s.store.User().SetDescription(userID, response.Description)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
		}
		s.respond(w, http.StatusOK, Response{Data: "Successfully updated description"})
	}
}

func (s *Server) userCover() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		var response struct {
			Cover string `json:"cover" validate:"required"`
		}
		if err := s.decode(r, &response); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}
		if err := validator.Validate(response); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}
		err := s.store.User().SetCover(userID, response.Cover)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
		}
		s.respond(w, http.StatusOK, Response{Data: "Successfully updated cover"})
	}
}

func (s *Server) userAvatar() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		var response struct {
			Avatar string `json:"avatar" validate:"required"`
		}
		if err := s.decode(r, &response); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}
		if err := validator.Validate(response); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}
		err := s.store.User().SetAvatar(userID, response.Avatar)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
		}
		s.respond(w, http.StatusOK, Response{Data: "Successfully updated avatar"})
	}
}

func (s *Server) userGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := s.store.User().GetAll()
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, Response{Data: users})
	}
}

func (s *Server) userGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := s.store.User().Get(r.PathValue("id"))
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		user.Sanitize()
		s.respond(w, http.StatusOK, Response{Data: user})
	}
}

func (s *Server) currentUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		user, err := s.store.User().Get(userID)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}
		user.Sanitize()
		s.respond(w, http.StatusOK, Response{Data: user})
	}
}
