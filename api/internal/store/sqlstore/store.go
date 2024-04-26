package sqlstore

import (
	"database/sql"

	"social-network/internal/store"
)

type Store struct {
	Db                *sql.DB
	userRepository    *UserRepository
	followRepository  *FollowRepository
	requestRepository *RequestRepository
	postRepository    *PostRepository
	commentRepository *CommentRepository
	groupRepository   *GroupRepository
	privacyRepository *PrivacyRepository
	eventRepository   *EventRepository
	chatRepository    *ChatRepository
	imageRepository   *ImageRepository
	sessionRepository *SessionRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		Db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

func (s *Store) Session() store.SessionRepository {
	if s.sessionRepository != nil {
		return s.sessionRepository
	}
	s.sessionRepository = &SessionRepository{
		store: s,
	}
	return s.sessionRepository
}

func (s *Store) Follow() store.FollowRepository {
	if s.followRepository != nil {
		return s.followRepository
	}

	s.followRepository = &FollowRepository{
		store: s,
	}

	return s.followRepository
}

func (s *Store) Request() store.RequestRepository {
	if s.requestRepository != nil {
		return s.requestRepository
	}

	s.requestRepository = &RequestRepository{
		store: s,
	}

	return s.requestRepository
}

func (s *Store) Post() store.PostRepository {
	if s.postRepository != nil {
		return s.postRepository
	}

	s.postRepository = &PostRepository{
		store: s,
	}
	s.imageRepository = &ImageRepository{
		store: s,
	}

	return s.postRepository
}

func (s *Store) Comment() store.CommentRepository {
	if s.commentRepository != nil {
		return s.commentRepository
	}

	s.commentRepository = &CommentRepository{
		store: s,
	}

	return s.commentRepository
}

func (s *Store) Group() store.GroupRepository {
	if s.groupRepository != nil {
		return s.groupRepository
	}

	s.groupRepository = &GroupRepository{
		store: s,
	}

	return s.groupRepository
}

func (s *Store) Privacy() store.PrivacyRepository {
	if s.privacyRepository != nil {
		return s.privacyRepository
	}

	s.privacyRepository = &PrivacyRepository{
		store: s,
	}

	return s.privacyRepository
}

func (s *Store) Event() store.EventRepository {
	if s.eventRepository != nil {
		return s.eventRepository
	}

	s.eventRepository = &EventRepository{
		store: s,
	}

	return s.eventRepository
}

func (s *Store) Chat() store.ChatRepository {
	if s.chatRepository != nil {
		return s.chatRepository
	}

	s.chatRepository = &ChatRepository{
		store: s,
	}

	return s.chatRepository
}

func (s *Store) Image() store.ImageRepository {
	if s.imageRepository != nil {
		return s.imageRepository
	}

	s.imageRepository = &ImageRepository{
		store: s,
	}

	return s.imageRepository
}
