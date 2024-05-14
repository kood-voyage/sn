package server

import (
	"errors"
	"net/http"
	"social-network/internal/model"
	"social-network/pkg/validator"
)

// createEvent Creates an event for a group
//
// @Summary Create a group event
// @Tags events
// @Accept json
// @Produce json
// @Success 200 {object} model.Event
// @Failure 401 {object} Error
// @Failure 422 {object} Error
// @Failure 500 {object} Error
// @Router /api/v1/auth/group/event/create [post]
func (s *Server) createEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		event := model.NewEvent()
		userID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		if err := s.decode(r, &event); err != nil {
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

// updateEvent Updates event details
//
// @Summary Update a group event
// @Tags events
// @Accept json
// @Produce json
// @Success 200 {object} model.Event
// @Failure 401 {object} Error
// @Failure 403 {object} Error
// @Failure 422 {object} Error
// @Router /api/v1/auth/group/event/update [put]
func (s *Server) updateEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var event model.Event

		userID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		if err := s.decode(r, &event); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		if err := validator.Validate(event); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		e, err := s.store.Event().Get(event.ID)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		if e.UserID != userID {
			s.error(w, http.StatusForbidden, errors.New("this user has no edit access to that event"))
			return
		}

		if err := s.store.Event().Update(&event); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, Response{Data: event})
	}
}

// deleteEvent Deletes an event
//
// @Summary Delete an event
// @Tags events
// @Accept json
// @Produce json
// @Param id path string true "event id to delete"
// @Success 200 {object} Response
// @Failure 401 {object} Error
// @Failure 403 {object} Error
// @Failure 422 {object} Error
// @Router /api/v1/auth/group/event/delete/{id} [delete]
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

// getEvent Information about single event
//
// @Summary Retrieve information about single event
// @Tags events
// @Accept json
// @Produce json
// @Param id path string true "event id"
// @Success 200 {object} model.Event
// @Failure 401 {object} Error
// @Failure 403 {object} Error
// @Failure 422 {object} Error
// @Router /api/v1/auth/group/event/{id} [delete]
func (s *Server) getEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		event, err := s.store.Event().Get(r.PathValue("id"))
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		//check group id
		t, err := s.store.Group().IsMember(event.GroupID, userID)
		if err != nil {
			s.error(w, http.StatusForbidden, err)
			return
		}
		if !t {
			s.error(w, http.StatusForbidden, errors.New("user is not part of the group with that event"))
			return
		}

		s.respond(w, http.StatusOK, Response{Data: event})
	}
}

// registerEvent Register to a specific event
//
// @Summary Register to an event
// @Tags events
// @Accept json
// @Produce json
// @Param id path string true "event id"
// @Param opt path string true "option"
// @Success 200 {object} Response
// @Failure 401 {object} Error
// @Failure 403 {object} Error
// @Failure 422 {object} Error
// @Router /api/v1/auth/group/event/{id} [delete]
func (s *Server) registerEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}
		eventID := r.PathValue("id")
		option := r.PathValue("opt")
		if option != "interested" && option != "going" && option != "notgoing" && option != "maybe" {
			s.error(w, http.StatusUnprocessableEntity, errors.New("only options allowed - interested, going, notgoing, maybe"))
			return
		}
		event, err := s.store.Event().Get(eventID)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		t, err := s.store.Group().IsMember(event.GroupID, userID)
		if err != nil {
			s.error(w, http.StatusForbidden, err)
			return
		}
		if !t {
			s.error(w, http.StatusForbidden, errors.New("user is not part of the group to register to that event"))
			return
		}

		if err := s.store.Event().Register(userID, eventID, s.types.Event.Values[option]); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, Response{Data: nil})

	}
}

func (s *Server) getGroupEvents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		group_id := r.PathValue("id")
		userID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		events, err := s.store.Group().GetAllEvents(group_id, userID)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, Response{
			Data: events,
		})
	}
}
