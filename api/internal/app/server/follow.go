package server

import (
	"errors"
	"net/http"
	"social-network/internal/model"
)

// handleFollow handles the follow action where one user follows another. If another user profile is private, it creates a follow request instead.
//
// @Summary Follow a user
// @Tags follow
// @Accept json
// @Produce json
// @Param id path string true "Target user ID to follow"
// @Success 201 {object} Response
// @Failure 401 {object} Error
// @Failure 500 {object} Error
// @Router /api/v1/auth/follow/{id} [get]
func (s *Server) handleFollow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sourceID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		follow := model.Follower{
			SourceID: sourceID,
			TargetID: r.PathValue("id"),
		}

		//follow user
		if err := s.store.Follow().Create(follow); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		//create notification
		notification := model.Request{
			TypeID:   s.types.Request.Notification,
			SourceID: sourceID,
			TargetID: r.PathValue("id"),
			Message:  "started following you.",
		}

		if err := s.store.Request().Create(notification); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusCreated, Response{Data: nil})
	}
}

// handleUnfollow handles the unfollow action where one user unfollows another.
//
// @Summary Unfollow a user
// @Tags follow
// @Accept json
// @Produce json
// @Param id path string true "Target user ID to Unfollow"
// @Success 201 {object} Response
// @Failure 401 {object} Error
// @Failure 500 {object} Error
// @Router /api/v1/auth/unfollow/{id} [get]
func (s *Server) handleUnfollow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sourceID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		}
		follow := model.Follower{
			SourceID: sourceID,
			TargetID: r.PathValue("id"),
		}

		//unfollow a user
		if err := s.store.Follow().Delete(follow); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusCreated, Response{Data: nil})
	}
}

// handleFollowRequest handles the follow action where one user profile is private and creates a request to view their profile. If request is accepted create another request to api/v1/follow/{id} endpoint
//
// @Summary Request a follow for user
// @Tags follow
// @Accept json
// @Produce json
// @Param id path string true "Target user ID to request follow"
// @Success 201 {object} Response
// @Failure 401 {object} Error
// @Failure 500 {object} Error
// @Router /api/v1/auth/follow/request/{id} [get]
func (s *Server) handleFollowRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sourceID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		}
		request := model.Request{
			TypeID:   s.types.Request.Follow,
			SourceID: sourceID,
			TargetID: r.PathValue("id"),
		}

		if err := s.store.Request().Create(request); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusCreated, Response{Data: nil})
	}
}
