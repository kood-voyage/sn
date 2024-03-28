package server

import (
	"database/sql"
	"errors"
	"net/http"
	"social-network/followservice/pkg/followservice"
	"social-network/internal/model"
	"social-network/pkg/validator"
)

// handleFollow handles the follow action where one user follows another. If another user profile is private, it creates a follow request instead.
//
// @Summary Follow a user or create follow request
// @Tags follow
// @Accept json
// @Produce json
// @Param id path string true "Target user ID to follow"
// @Success 201 {object} Response
// @Failure 400 {object} Error
// @Failure 401 {object} Error
// @Failure 422 {object} Error
// @Router /api/v1/auth/follow/{id} [get]
func (s *Server) handleFollow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sourceID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		followClient, err := s.FollowClient.Follow(r.Context(), &followservice.FollowRequest{
			SourceId: sourceID,
			TargetId: r.PathValue("id"),
		})
		if err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		s.respond(w, http.StatusCreated, followClient)
		//WHOLE LOGIC UP

		//this still needs implementing on follow service side
		// if privacyCode == s.types.Privacy.Private {
		// 	request := model.Request{
		// 		TypeID:   s.types.Request.Follow,
		// 		SourceID: sourceID,
		// 		TargetID: r.PathValue("id"),
		// 	}

		// 	existing, err := s.store.Request().Get(request)
		// 	if err != nil && err != sql.ErrNoRows {
		// 		s.error(w, http.StatusUnprocessableEntity, err)
		// 		return
		// 	}

		// 	if existing != nil {
		// 		s.error(w, http.StatusForbidden, errors.New("already request exists"))
		// 		return
		// 	}

		// 	if err := s.store.Request().Create(request); err != nil {
		// 		s.error(w, http.StatusUnprocessableEntity, err)
		// 		return
		// 	}
		// 	s.respond(w, http.StatusCreated, Response{Data: "Successfully created a follow request"})
		// 	return
		// }

		// s.error(w, http.StatusBadRequest, errors.New("Invalid"))

	}
}

// handleUnfollow handles the unfollow action where one user unfollows another.
//
// @Summary Unfollow a user
// @Tags follow
// @Accept json
// @Produce json
// @Param id path string true "Target user ID to Unfollow"
// @Success 201 {object} Response
// @Failure 401 {object} Error
// @Failure 500 {object} Error
// @Router /api/v1/auth/unfollow/{id} [get]
func (s *Server) handleUnfollow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sourceID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		_, err := s.FollowClient.UnFollow(r.Context(), &followservice.FollowRequest{
			SourceId: sourceID,
			TargetId: r.PathValue("id"),
		})
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, Response{Data: err})
	}
}

// handleFollowRequest Handles the follow request accept / reject
//
// @Summary Resolves a follow request
// @Tags follow
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Failure 401 {object} Error
// @Failure 500 {object} Error
// @Router /api/v1/auth/follow/request [post]
func (s *Server) handleFollowRequest() http.HandlerFunc {
	type Req struct {
		TargetID string `json:"target_id"`
		Option   string `json:"option" validate:"lowercase"`
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
			TypeID:   s.types.Request.Follow,
		}

		if request.Option == "reject" {
			if err := s.store.Request().Delete(req); err != nil {
				s.error(w, http.StatusUnprocessableEntity, err)
				return
			}
			s.respond(w, http.StatusOK, Response{Data: "Rejected user request"})
			return
		} else if request.Option == "accept" {
			//check if request exists in first place
			req_exists, err := s.store.Request().Get(req)
			if err != nil {
				if err == sql.ErrNoRows {
					//create a follow link
					if err := s.store.Follow().Create(model.Follower{
						SourceID: request.TargetID,
						TargetID: sourceID,
					}); err != nil {
						s.error(w, http.StatusUnprocessableEntity, err)
						return
					}
				}
				s.error(w, http.StatusUnprocessableEntity, err)
				return
			}

			//delete the request
			if err := s.store.Request().Delete(*req_exists); err != nil {
				s.error(w, http.StatusUnprocessableEntity, err)
				return
			}

			s.respond(w, http.StatusOK, Response{Data: "Accepted user request"})
			return
		} else {
			s.error(w, http.StatusBadRequest,
				errors.New("invalid option"))
		}

	}
}
