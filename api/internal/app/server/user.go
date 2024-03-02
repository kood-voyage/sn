package server

import (
	"errors"
	"net/http"
	"social-network/internal/model"
)

// userCreate Handles the user creation to database. Only userID and privacy state will be stored.
//
// @Summary Create a user with privacy state
// @Tags users
// @Produce json
// @Param privacy_state path string true "Only public, private, selected allowed"
// @Success 201 {object} model.User
// @Failure 400 {object} Error
// @Failure 401 {object} Error
// @Failure 422 {object} Error
// @Router /api/v1/auth/user/create/{privacy_state} [get]
func (s *Server) userCreate() http.HandlerFunc {
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
			s.error(w, http.StatusBadRequest, errors.New("public, private, selected states are allowed"))
			return
		}

		if err := s.store.User().Create(user, privacy); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusCreated, Response{Data: user})
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
			s.error(w, http.StatusBadRequest, errors.New("public, private, selected states are allowed"))
			return
		}

		if err := s.store.User().UpdatePrivacy(user, privacy); err != nil {
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

//NEED TO WRITE TEST FOR THAT FUNCTION AFTER FINISHING WITH POST PRIVACY CHECKING

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
		//firstly check if the request is done by the user itself or some other user
		userID_request, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("unauthorized user"))
			return
		}

		if userID_request != r.PathValue("id") {
			//check requested user profile privacy
			privacy, err := s.store.User().CheckPrivacy(r.PathValue("id"))
			if err != nil {
				s.error(w, http.StatusUnprocessableEntity, err)
				return
			}
			if privacy == s.types.Privacy.Private {
				//check if user follows
				userFollows, err := s.store.User().IsFollowing(userID_request, r.PathValue("id"))
				if err != nil {
					s.error(w, http.StatusUnprocessableEntity, err)
					return
				}
				if !userFollows {
					s.error(w, http.StatusForbidden, errors.New("users profile is private"))
					return
				}
			}
		}

		posts, err := s.store.Post().GetUsers(r.PathValue("id"))
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

		notifications, err := s.store.User().GetNotifications(userID, s.types.Request.Notification)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, Response{Data: notifications})
	}
}
