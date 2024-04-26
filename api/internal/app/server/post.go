package server

import (
	"errors"
	"fmt"
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
			fmt.Println("ERROR 1 ", err)
			s.error(w, http.StatusInternalServerError, err)
			return
		}
		post.UserID = userID

		if err := validator.Validate(post); err != nil {
			fmt.Println("ERROR 2 ", err)
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}
		fmt.Printf("post: %+v\n", post)

		if err := s.store.Post().Create(
			post,
			s.types.Privacy.Values[post.Privacy],
		); err != nil {
			fmt.Println("ERROR 3 ", err)

			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusCreated, Response{Data: post})
	}
}

// updatePost update post
//
// @Summary Update post
// @Tags posts
// @Produce json
// @Success 202
// @Failure 401 {object} Error
// @Failure 403 {object} Error
// @Failure 422 {object} Error
// @Router /api/v1/auth/post/update [put]
func (s *Server) updatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var post model.Post
		userID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		if err := s.decode(r, &post); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		if err := validator.Validate(post); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		p, err := s.store.Post().Get(post.ID)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		if p.UserID != userID {
			s.error(w, http.StatusForbidden, errors.New("not allowed to update post"))
			return
		}

		if err = s.store.Post().Update(&post, s.types.Privacy.Values[post.Privacy]); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusAccepted, nil)
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

// getPost Get post
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

// addSelected
//
// @Summary Add selected user to list
// @Tags posts
// @Accept json
// @Produce json
// @Success 201
// @Failure 500 {object} Error
// @Router /api/v1/auth/posts/selected/add [post]
func (s *Server) addSelected() http.HandlerFunc {
	type requestBody struct {
		UserList []model.User `json:"user_list"`
		ParentID string       `json:"parent_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		reqBody := requestBody{}
		if err := s.decode(r, &reqBody); err != nil {
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

// deleteSelected
//
// @Summary Removes selected user from list
// @Tags posts
// @Accept json
// @Produce json
// @Success 204
// @Failure 500 {object} Error
// @Router /api/v1/auth/posts/selected/delete [post]
func (s *Server) deleteSelected() http.HandlerFunc {
	type requestBody struct {
		UserList []model.User `json:"user_list"`
		ParentID string       `json:"parent_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		reqBody := requestBody{}
		if err := s.decode(r, &reqBody); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}
		fmt.Println(reqBody)
		if err := s.store.Post().RemoveSelected(&reqBody.UserList, reqBody.ParentID); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusNoContent, nil)
	}
}

func (s *Server) getFeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		posts, err := s.store.Post().GetUserFeed(userID)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, Response{
			Data: posts,
		})
	}
}
