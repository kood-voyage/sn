package server

import (
	"errors"
	"net/http"
	"social-network/internal/model"
)

func (s *Server) createEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		event := model.NewEvent()
		userID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		if err := s.decode(r, event); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}
		event.UserID = userID

		if err := s.store.Event().Create(event); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusCreated, Response{Data: event})
	}
}

func (s *Server) updateEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Server) deleteEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if err := s.store.Event().Delete(id); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusAccepted, nil)
	}
}

func (s *Server) getEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Server) registerEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
