package server

import (
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
		id := r.PathValue("id")
		if err := s.store.Comment().Delete(id); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusAccepted, nil)
	}
}
