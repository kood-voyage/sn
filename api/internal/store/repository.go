package store

import "social-network/internal/model"

type UserRepository interface {
	Create(user *model.User, privacy int) error
	UpdatePrivacy(user *model.User, privacy int) error
	CheckPrivacy(userID string) (int, error)
	GetFollowers(userID string) ([]model.User, error)
	GetFollowing(userID string) ([]model.User, error)
	IsFollowing(source_id, target_id string) (bool, error)
	GetNotifications(user_id string, req_type int) ([]model.Request, error)
}

type FollowRepository interface {
	Create(follower model.Follower) error
	Delete(follower model.Follower) error
}

type RequestRepository interface {
	Create(request model.Request) error
	Delete(request model.Request) error
	Get(request model.Request) (*model.Request, error)
	DeleteByID(id string) error
}

type PostRepository interface {
	Create(post *model.Post, privacy int) error
	Delete(id string) error
	Get(id string) (*model.Post, error)
}

type CommentRepository interface {
	Create(post *model.Comment) error
	Delete(id string) error
	Get(id string) (*model.Comment, error)
	GetUsers(user_id string) ([]model.Post, error)
	// Update(string) error
}

type GroupRepository interface {
	Create(group model.Group) (*model.Group, error)
	Delete(group_id string) error
	Update(group model.Group) error
	Get(group_id string) (*model.Group, error)
}
