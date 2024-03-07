package server

import (
	"errors"
	"fmt"
	"net/http"
	"social-network/internal/model"
	"social-network/pkg/validator"
)

// notificationCreate handles the creation of notification.
//
// @Summary Create a notification from source id to target id
// @Tags notification
// @Produce json
// @Success 200 {object} model.Request
// @Failure 401 {object} Error
// @Failure 406 {object} Error
// @Failure 422 {object} Error
// @Router /api/v1/auth/notification/create [post]
func (s *Server) notificationCreate() http.HandlerFunc {
	Notification := model.NotificationRequest()
	return func(w http.ResponseWriter, r *http.Request) {
		if err := s.decode(r, Notification); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		user_id, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		//Random user could not make notifications for everybody
		if user_id != Notification.SourceID {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		if err := validator.Validate(Notification); err != nil {
			s.error(w, http.StatusNotAcceptable, err)
			return
		}

		if err := s.store.Request().Create(*Notification); err != nil {
			fmt.Println(err)
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, Response{Data: Notification})
	}
}

// notificationDelete handles the deletion of notification.
//
// @Summary Delete a notification by ID
// @Tags notification
// @Produce json
// @Param id path string true "Notification ID to delete"
// @Success 200 {object} Response
// @Failure 422 {object} Error
// @Router /api/v1/auth/notification/delete/{id} [delete]
func (s *Server) notificationDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if err := s.store.Request().DeleteByID(r.PathValue("id")); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, Response{Data: nil})
	}
}
