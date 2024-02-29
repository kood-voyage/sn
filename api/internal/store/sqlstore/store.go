package sqlstore

import (
	"database/sql"
	"social-network/internal/store"
)

type Store struct {
	Db                     *sql.DB
	userRepository         *UserRepository
	followRepository       *FollowRepository
	requestRepository      *RequestRepository
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
