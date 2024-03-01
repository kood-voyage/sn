package store

import "social-network/internal/model"

type UserRepository interface {
	Create(user *model.User, privacy int) error
	UpdatePrivacy(user *model.User, privacy int) error
	CheckPrivacy(userID string) (int, error)
	GetFollowers(userID string) ([]model.User, error)
	GetFollowing(userID string) ([]model.User, error)
}

type FollowRepository interface {
	Create(follower model.Follower) error
	Delete(follower model.Follower) error
}

type RequestRepository interface {
	Create(request model.Request) error
	Delete(request model.Request) error
	Get(request model.Request) (*model.Request, error)
}

type PostRepository interface {
	Create(post *model.Post) error
	Delete(id string) error
	Get(id string) (*model.Post, error)
	// Update(string) error
}
