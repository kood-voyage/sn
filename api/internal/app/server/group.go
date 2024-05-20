package server

import (
	"database/sql"
	"errors"
	"net/http"
	"social-network/internal/model"
	"social-network/pkg/validator"
	"time"
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

// groupUpdate handles the group information update.
//
// @Summary Update group information
// @Tags group
// @Produce json
// @Success 202
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

		s.respond(w, http.StatusAccepted, nil)
	}
}

// groupDelete handles the deletion of a group.
//
// @Summary Delete a group
// @Tags group
// @Produce json
// @Param id path string true "Group ID to delete"
// @Success 202
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

		s.respond(w, http.StatusAccepted, nil)
	}
}

// groupGet handles the retrieval of a group and its information.
//
// @Summary Returns group information
// @Tags group
// @Produce json
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

		//firstly retrieve the group
		group, err := s.store.Group().Get(r.PathValue("id"))
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		//check group privacy status
		privacy, err := s.store.Privacy().Check(group.ID)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		if privacy == s.types.Privacy.Private {
			t, err := s.store.Group().IsMember(group.ID, sourceID)
			if err != nil {
				s.error(w, http.StatusForbidden, err)
				return
			}

			if !t {
				s.error(w, http.StatusForbidden, errors.New("you don't have access to this group"))
				return
			}
		}

		s.respond(w, http.StatusOK, Response{Data: group})
	}
}

// groupGetAll handles the retrieval of all group and their information.
//
// @Summary Returns groups
// @Tags group
// @Produce json
// @Success 200 {object} []model.Group
// @Failure 422 {object} Error
// @Router /api/v1/auth/group [get]
func (s *Server) groupGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groups, err := s.store.Group().GetAll(s.types)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, Response{Data: groups})
	}
}

// groupInvite handles the invitation of a member to a group.
//
// @Summary Creates a request to invite another user to a group
// @Tags group
// @Produce json
// @Success 200 {object} model.Group
// @Failure 401 {object} Error
// @Failure 403 {object} Error
// @Failure 422 {object} Error
// @Router /api/v1/auth/group/invite [post]
func (s *Server) groupInvite() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type requestBody struct {
			GroupID  string `json:"group_id" validate:"required"`
			TargetID string `json:"target_id" validate:"required"`
			Message  string `json:"message" validate:"required"`
		}
		sourceID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		var req requestBody
		if err := s.decode(r, &req); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		if err := validator.Validate(req); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		//check source id
		t, err := s.store.Group().IsMember(req.GroupID, sourceID)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}
		if !t {
			s.error(w, http.StatusUnprocessableEntity, errors.New("user is not a group member"))
			return
		}

		//check target_id
		t, err = s.store.Group().IsMember(req.GroupID, req.TargetID)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}
		if t {
			s.error(w, http.StatusUnprocessableEntity, errors.New("user is already in a group"))
			return
		}

		//create a request
		request := model.InviteRequest()
		request.TargetID = req.TargetID
		request.SourceID = sourceID
		request.CreatedAt = time.Now()
		request.Message = req.Message
		request.ParentID = req.GroupID
		//check if that request it not already created
		existing, err := s.store.Request().GetGroups(*request)
		if err != nil && err != sql.ErrNoRows {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		if existing != nil {
			s.error(w, http.StatusForbidden, errors.New("already request exists"))
			return
		}
		//if not create a request
		_, err = s.store.Request().Create(*request)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, Response{Data: nil})
	}
}

// groupInviteRequest Handles the group invitation request accept / reject
//
// @Summary Resolves a group invitation request
// @Tags follow
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Failure 401 {object} Error
// @Failure 500 {object} Error
// @Router /api/v1/auth/group/request [post]
func (s *Server) groupInviteRequest() http.HandlerFunc {
	type Req struct {
		TargetID string `json:"target_id" validate:"required"`
		Option   string `json:"option" validate:"lowercase|required|contains:accept,reject"`
		GroupID  string `json:"group_id" validate:"required"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var request Req
		sourceID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		if err := s.decode(r, &request); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		if err := validator.Validate(request); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		//The targetID is source id because first time the request was created from one user(source user) and now the second user(target user) needs to confirm it hence why (target user) sends the request and on the targetID field needs to be the (target user) - getting the user from jwt middleware its naming is just sourceid
		req := model.Request{
			TargetID: sourceID,
			SourceID: request.TargetID,
			TypeID:   s.types.Request.Invite,
		}

		if request.Option == "reject" {
			if err := s.store.Request().Delete(req); err != nil {
				s.error(w, http.StatusUnprocessableEntity, err)
				return
			}
			s.respond(w, http.StatusOK, Response{Data: "Rejected invitation to a group"})
			return
		} else if request.Option == "accept" {
			//check if request exists in first place
			req_exists, err := s.store.Request().Get(req)

			if err != nil {
				s.error(w, http.StatusUnprocessableEntity, err)
				return
			}

			//add member to a group
			if err := s.store.Group().AddMember(request.GroupID, sourceID); err != nil {
				s.error(w, http.StatusUnprocessableEntity, err)
				return
			}

			//delete the request
			if err := s.store.Request().Delete(*req_exists); err != nil {
				s.error(w, http.StatusUnprocessableEntity, err)
				return
			}

			s.respond(w, http.StatusOK, Response{Data: "Accepted group invitation"})
			return
		} else {
			s.error(w, http.StatusBadRequest,
				errors.New("invalid option"))
		}

	}
}

func (s *Server) groupGetPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sourceID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		group_name := r.PathValue("id")

		groupInfo, err := s.store.Group().Get(group_name)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		if s.types.Privacy.Values[groupInfo.Privacy] == s.types.Privacy.Private {
			t, err := s.store.Group().IsMember(groupInfo.ID, sourceID)
			if err != nil {
				s.error(w, http.StatusUnprocessableEntity, err)
				return
			}
			if !t {
				s.error(w, http.StatusForbidden, errors.New("group is private and user is not part of the group"))
				return
			}
		}

		group_posts, err := s.store.Group().GetPosts(groupInfo.ID)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, Response{Data: group_posts})
	}
}

func (s *Server) joinGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sourceID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		//id --> its actually group name
		group_name := r.PathValue("id")

		//check if group is public
		groupInfo, err := s.store.Group().Get(group_name)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		if s.types.Privacy.Public != s.types.Privacy.Values[groupInfo.Privacy] && groupInfo.CreatorID != sourceID {
			s.error(w, http.StatusForbidden, errors.New("can not join to private group"))
			return
		}

		if err := s.store.Group().AddMember(groupInfo.ID, sourceID); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusCreated, Response{Data: nil})
	}
}

func (s *Server) groupInvitedUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		group_id := r.PathValue("id")

		users, err := s.store.Group().GetInvitedUsers(group_id)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, Response{Data: users})
	}
}
