package server

import (
	"errors"
	"fmt"
	"net/http"
	"social-network/internal/model"
	"social-network/pkg/validator"
)

// groupCreate handles the creation of group.
//
// @Summary Create a group
// @Tags group
// @Produce json
// @Success 201 {object} model.Group
// @Failure 401 {object} Error
// @Failure 422 {object} Error
// @Router /api/v1/auth/group/create [post]
func (s *Server) groupCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		group := model.NewGroup()
		userID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		if err := s.decode(r, group); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		group.CreatorID = userID
		//validate group
		if err := validator.Validate(group); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		g, err := s.store.Group().Create(*group, s.types.Privacy.Values[group.Privacy])
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusCreated, Response{Data: g})
	}
}

// groupUpdate handles the group inforomation update.
//
// @Summary Update group information
// @Tags group
// @Produce json
// @Success 200 {object} model.Group
// @Failure 401 {object} Error
// @Failure 403 {object} Error
// @Failure 422 {object} Error
// @Router /api/v1/auth/group/update [put]
func (s *Server) groupUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var group model.Group
		userID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		if err := s.decode(r, &group); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		if err := validator.Validate(group); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		g, err := s.store.Group().Get(group.ID)
		if err != nil {
			fmt.Println(err)
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		if g.CreatorID != userID {
			s.error(w, http.StatusForbidden, errors.New("not allowed to update group information"))
			return
		}

		if err := s.store.Group().Update(group, s.types.Privacy.Values[group.Privacy]); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, Response{Data: nil})
	}
}

// groupDelete handles the deletion of a group.
//
// @Summary Delete a group
// @Tags group
// @Produce json
// @Param id path string true "Group ID to delete"
// @Success 200 {object} Response
// @Failure 401 {object} Error
// @Failure 403 {object} Error
// @Failure 422 {object} Error
// @Router /api/v1/auth/group/delete/{id} [delete]
func (s *Server) groupDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		g, err := s.store.Group().Get(r.PathValue("id"))
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		if g.CreatorID != userID {
			s.error(w, http.StatusForbidden, errors.New("not allowed to delete a group"))
			return
		}

		if err := s.store.Group().Delete(r.PathValue("id")); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, Response{Data: nil})
	}
}

// groupGet handles the retrieval of a group and its information.
//
// @Summary Returns group information
// @Tags group
// @Produce json
// @Param id path string true "Group ID to get information"
// @Success 200 {object} model.Group
// @Failure 401 {object} Error
// @Failure 403 {object} Error
// @Failure 422 {object} Error
// @Router /api/v1/auth/group/{id} [get]
func (s *Server) groupGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sourceID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		//check group privacy status
		privacy, err := s.store.Privacy().Check(r.PathValue("id"))
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		if privacy == s.types.Privacy.Private {
			if err := s.store.Group().IsMember(r.PathValue("id"), sourceID); err != nil {
				s.error(w, http.StatusForbidden, err)
				return
			}
		}

		group, err := s.store.Group().Get(r.PathValue("id"))
		if err != nil {
			s.error(w, http.StatusUnauthorized, err)
			return
		}

		s.respond(w, http.StatusOK, Response{Data: group})
	}
}
