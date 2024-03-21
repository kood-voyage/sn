package server

import (
	"errors"
	"net/http"
	"social-network/internal/model"
	"social-network/pkg/validator"
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
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}
		event.UserID = userID

		if err := validator.Validate(event); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		//check if creator is part of a group
		t, err := s.store.Group().IsMember(event.GroupID, userID)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}
		if !t {
			s.error(w, http.StatusUnprocessableEntity, errors.New("user is not part of the group"))
			return
		}

		if err := s.store.Event().Create(event); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
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
		eventId := r.PathValue("id")

		userID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}
		//if event exists and get groupID
		event, err := s.store.Event().Get(eventId)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		//check if user is part of a group
		t, err := s.store.Group().IsMember(event.GroupID, userID)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}
		if !t {
			s.error(w, http.StatusUnprocessableEntity, errors.New("user is not part of the group to delete an event"))
			return
		}

		if err := s.store.Event().Delete(eventId); err != nil {
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
