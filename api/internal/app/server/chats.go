package server

import (
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

		s.respond(w, http.StatusCreated, nil)
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
		chatLine := model.NewChatLine()
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
