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
// @Router /api/v1/auth/posts/create [post]
func (s *Server) createPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		post := model.NewPost()

		if err := s.decode(r, post); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		if err := validator.Validate(post); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		if err := s.store.Post().Create(post); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusCreated, Response{Data: post})
	}
}

// deletePost deletes post
//
// @Summary Delete post
// @Tags posts
// @Produce json
// @Success 202
// @Failure 500 {object} Error
// @Router /api/v1/auth/posts/delete/{id} [delete]
func (s *Server) deletePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if err := s.store.Post().Delete(id); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusAccepted, nil)
	}
}

// getPost deletes post
//
// @Summary Get post
// @Tags posts
// @Produce json
// @Success 200 {object} model.Post
// @Failure 500 {object} Error
// @Router /api/v1/auth/posts/{id} [get]
func (s *Server) getPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		post, err := s.store.Post().Get(id)
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusOK, Response{Data: post})
	}
}
