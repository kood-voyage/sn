package store

import "social-network/internal/model"

type UserRepository interface {
	Create(user *model.User) error
}

type FollowRepository interface {
	Create(follower model.Follower) error
	Delete(follower model.Follower) error
}

type RequestRepository interface {
	Create(request model.Request) error
}

type PostRepository interface {
	Create(post *model.Post) error
	Delete(id string) error
	Get(id string) (*model.Post, error)
	// Update(string) error
}

type CommentRepository interface {
	Create(post *model.Comment) error
	Delete(id string) error
	Get(id string) (*model.Comment, error)
}
