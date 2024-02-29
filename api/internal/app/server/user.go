package server

import (
	"net/http"
	models "social-network/internal/model"
	"social-network/pkg/validator"
)

func (s *server) createUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{}

		if err := s.decode(r, user); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		if err := validator.Validate(user); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		if err := s.store.User().Create(user); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, Response{Data: user})
	}
}
