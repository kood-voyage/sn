package server

import (
	"errors"
	"net/http"
	"social-network/internal/model"
	"social-network/pkg/validator"
)

// createComment handles the post creation
//
// @Summary Create comment
// @Tags comments
// @Accept json
// @Produce json
// @Success 201 {object} model.Comment
// @Failure 500 {object} Error
// @Failure 422 {object} Error
// @Router /api/v1/auth/comment/create [post]
func (s *Server) createComment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		comment := model.NewComment()

		if err := s.decode(r, comment); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		if err := validator.Validate(comment); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		if err := s.store.Comment().Create(comment); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusCreated, Response{Data: comment})
	}
}

// deletePost deletes comment
//
// @Summary Delete comment
// @Tags comments
// @Produce json
// @Success 202
// @Failure 500 {object} Error
// @Router /api/v1/auth/comment/delete/{id} [delete]
func (s *Server) deleteComment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID := r.PathValue("id")
		userID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusInternalServerError, errors.New("unauthorized"))
			return
		}

		if err := s.store.Comment().Delete(postID, userID); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusAccepted, nil)
	}
}

// getComments retrieves comments for a single post
//
// @Summary Retrieves comments for a single post
// @Tags comments
// @Produce json
// @Success 200 {object} []model.Comment
// @Failure 500 {object} Error
// @Router /api/v1/auth/comment/{id} [get]
func (s *Server) getComments() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID := r.PathValue("id")

		comments, err := s.store.Comment().Get(postID)
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusOK, comments)
	}
}
