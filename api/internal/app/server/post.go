package server

import (
	"net/http"
	"social-network/internal/model"
	"social-network/pkg/validator"
)

// createPost handles the post creation
//
// @Summary Create post
// @Tags posts
// @Accept json
// @Produce json
// @Success 201 {object} model.Post
// @Failure 500 {object} Error
// @Failure 422 {object} Error
// @Router /api/v1/posts/create [post]
func (s *Server) createPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		post := model.NewPost()

		if err := s.decode(r, post); err != nil {
			s.error(w, http.StatusInternalServerError, err)
		}

		if err := validator.Validate(post); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
		}

		s.respond(w, http.StatusCreated, Response{Data: post})
	}
}
