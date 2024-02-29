package server

import (
	"errors"
	"fmt"
	"net/http"
	models "social-network/internal/model"
)

// handleFollow handles the follow action where one user follows another.
//
// @Summary Follow a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "Target user ID to follow"
// @Success 201 {object} Response
// @Failure 500 {object} Error
// @Router /api/v1/follow/{id} [get]
func (s *Server) handleFollow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sourceID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		}
		follow := models.Follower{
			SourceID: sourceID,
			TargetID: r.PathValue("id"),
		}

		//follow user
		if err := s.store.Follow().Create(follow); err != nil {
			fmt.Println(err)
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		//create notification
		notification := models.Request{
			TypeID:   "notification",
			SourceID: sourceID,
			TargetID: r.PathValue("id"),
			Message:  "started following you.",
		}

		if err := s.store.Request().Create(notification); err != nil {
			fmt.Println(err)
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusCreated, Response{Data: nil})
	}
}
