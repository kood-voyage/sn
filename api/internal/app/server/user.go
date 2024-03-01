package server

import (
	"errors"
	"fmt"
	"net/http"
	"social-network/internal/model"
)

// createUser Handles the user creation to database. Only userID and privacy state will be stored.
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
func (s *Server) createUser() http.HandlerFunc {
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
			fmt.Println(err)
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusCreated, Response{Data: user})
	}
}

// updatePrivacy Will update the user's privacy.
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
func (s *Server) updatePrivacy() http.HandlerFunc {
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

// handleUserFollowers Handles the retrieval of provided users followers.
//
// @Summary Returns a list of user followers
// @Tags users
// @Produce json
// @Param id path string true "User id to get followers"
// @Success 201 {object} []model.User
// @Failure 422 {object} Error
// @Router /api/v1/auth/user/followers/{id} [get]
func (s *Server) handleUserFollowers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		followers, err := s.store.User().GetFollowers(r.PathValue("id"))
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, http.StatusCreated, Response{Data: followers})
	}
}

// handleUserFollowing Handles user own following.
//
// @Summary Return a list of users who is user following
// @Tags users
// @Produce json
// @Param id path string true "User id to get followers"
// @Success 201 {object} []model.User
// @Failure 422 {object} Error
// @Router /api/v1/auth/user/following/{id} [get]
func (s *Server) handleUserFollowing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		followers, err := s.store.User().GetFollowing(r.PathValue("id"))
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, http.StatusCreated, Response{Data: followers})
	}
}
