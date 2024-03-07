package store

import "social-network/internal/model"

type UserRepository interface {
	Create(user *model.User, privacy int) error
	GetFollowers(userID string) ([]model.User, error)
	GetFollowing(userID string) ([]model.User, error)
	IsFollowing(source_id, target_id string) (bool, error)
	GetNotifications(user_id string) ([]model.Request, error)
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
	GetUsers(source_id, target_id string) ([]model.Post, error)
}

type CommentRepository interface {
	Create(post *model.Comment) error
	Delete(id string) error
	Get(id string) (*model.Comment, error)
	// Update(string) error
}

type GroupRepository interface {
	Create(group model.Group, privacy int) (*model.Group, error)
	Delete(group_id string) error
	Update(group model.Group, privacy int) error
	Get(group_id string) (*model.Group, error)
	Members(group_id string) (*[]model.User, error)
	IsMember(group_id, user_id string) error
}

type PrivacyRepository interface {
	Set(parent_id string, privacy int) error
	Update(parent_id string, privacy int) error
	Delete(parent_id string) error
	Check(parent_id string) (int, error)
}
