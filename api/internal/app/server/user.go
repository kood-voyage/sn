package server

import (
	"net/http"
	"social-network/internal/model"
	"social-network/pkg/validator"
)

// @Summary Sign up
// @Tags users
// @Accept  json
// @Produce  json
// @Success 201 {string} string "ok"
// @Router /users/create [get]
func (s *Server) createUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &model.User{}

		if err := s.decode(r, user); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		if err := validator.Validate(user); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		if err := s.store.User().Create(user); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusCreated, Response{Data: user})
	}
}
