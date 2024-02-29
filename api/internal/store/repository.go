package store

import models "social-network/internal/model"

type UserRepository interface {
	Create(user *models.User) error
}

type FollowRepository interface {
	Create(follower models.Follower) error
	Delete(follower models.Follower) error
}

type RequestRepository interface {
	Create(request models.Request) error
}
