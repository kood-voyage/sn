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

	return s.postRepository
}
