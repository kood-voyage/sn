package server

import (
	"errors"
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
		userID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		if err := s.decode(r, post); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}
		post.UserID = userID

		if err := validator.Validate(post); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		if err := s.store.Post().Create(
			post,
			s.types.Privacy.Values[post.Privacy],
		); err != nil {
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

func (s *Server) addSelected() http.HandlerFunc {
	type requestBody struct {
		UserList []model.User `json:"user_list"`
		ParentID string       `json:"parent_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		reqBody := requestBody{}
		if err := s.decode(r, reqBody); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		if err := s.store.Post().AddSelected(&reqBody.UserList, reqBody.ParentID); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusCreated, Response{Data: reqBody})
	}
}

func (s *Server) removeSelected() http.HandlerFunc {
	type requestBody struct {
		UserList []model.User `json:"user_list"`
		ParentID string       `json:"parent_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		reqBody := requestBody{}
		if err := s.decode(r, reqBody); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		if err := s.store.Post().RemoveSelected(&reqBody.UserList, reqBody.ParentID); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusNoContent, nil)
	}
}
