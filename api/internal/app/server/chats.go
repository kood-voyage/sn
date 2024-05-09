package server

import (
	"errors"
	"fmt"
	"net/http"

	"social-network/internal/model"
	"social-network/pkg/validator"
)

// createChat handles chat creation
//
// @Summary Create chat
// @Tags chats
// @Accept json
// @Produce json
// @Success 201 {object} model.Chat
// @Failure 422 {object} Error
// @Router /api/v1/auth/chats/create [post]
func (s *Server) createChat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		var chat model.Chat
		if err := s.decode(r, &chat); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}
		if err := validator.Validate(chat); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		c, err := s.store.Chat().Create(chat)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		if err := s.store.Chat().AddUser(model.User{ID: userID}, chat); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusCreated, c)
	}
}

// addUserChat Add user to chat for groups
//
// @Summary Add user to chat
// @Tags chats
// @Accept json
// @Produce json
// @Success 201 {object} model.ChatUser
// @Failure 422 {object} Error
// @Router /api/v1/auth/chats/add/user [post]
func (s *Server) addUserChat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var chatUser model.ChatUser
		if err := s.decode(r, &chatUser); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}
		if err := validator.Validate(chatUser); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		if err := s.store.Chat().AddUser(model.User{ID: chatUser.UserID}, model.Chat{ID: chatUser.ChatID}); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusCreated, Response{
			Data: nil,
		})
	}
}

// addLineChat Add line to chat
//
// @Summary Add line to chat
// @Tags chats
// @Accept json
// @Produce json
// @Success 201 {object} model.ChatLine
// @Failure 422 {object} Error
// @Router /api/v1/auth/chats/add/user [post]
func (s *Server) addLineChat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sourceID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		chatLine := model.NewChatLine()
		chatLine.UserID = sourceID
		if err := s.decode(r, &chatLine); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}
		if err := validator.Validate(chatLine); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		cLine, err := s.store.Chat().AddLine(chatLine)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusCreated, cLine)
	}
}

// getAllChats Retrieve all chats
//
// @Summary Get all chats
// @Tags chats
// @Accept json
// @Produce json
// @Success 201 {object} []model.Chat
// @Failure 422 {object} Error
// @Router /api/v1/auth/chats [get]
func (s *Server) getAllChats() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sourceID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		chats, err := s.store.Chat().GetAll(sourceID)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, Response{
			Data: chats,
		})
	}
}

// getChatLines Retrieve all chatlines of a specific chat
//
// @Summary Get all chatlines of a specific chat
// @Tags chats
// @Accept json
// @Produce json
// @Success 201 {object} []model.ChatLine
// @Failure 422 {object} Error
// @Router /api/v1/auth/chats/{id} [get]
func (s *Server) getChatLines() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sourceID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		chatLines, err := s.store.Chat().Load(r.PathValue("id"), sourceID)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, Response{
			Data: chatLines,
		})
	}
}

// getChatLines Retrieve all chatlines of a specific chat
//
// @Summary Get all chatlines of a specific chat
// @Tags chats
// @Accept json
// @Produce json
// @Success 201 {object} []model.ChatLine
// @Failure 422 {object} Error
// @Router /api/v1/auth/chats/{id} [get]
func (s *Server) getAllChatUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		chat := model.Chat{ID: r.PathValue("id")}
		chatUsers, err := s.store.Chat().GetUsers(chat)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}
		fmt.Println("CHAT USERS >>> ", chatUsers)
		s.respond(w, http.StatusOK, Response{
			Data: chatUsers,
		})
	}
}
