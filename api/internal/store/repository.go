package store

import "social-network/internal/model"

type UserRepository interface {
	Create(user *models.User) error
}
