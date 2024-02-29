package store

type Store interface {
	User() UserRepository
	Follow() FollowRepository
	Request() RequestRepository
}
