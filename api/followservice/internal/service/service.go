package service

import (
	"fmt"
	"social-network/followservice/internal/repository"
)

type FollowService interface {
	FollowUser(sourceId, targetId string) error
	UnFollowUser(sourceId, targetId string) error
}

type followService struct {
	repository repository.FollowRepository
}

func NewService(rep repository.FollowRepository) *followService {
	return &followService{
		repository: rep,
	}
}

func (f *followService) FollowUser(sourceId, targetId string) error {
	f.repository.Create(sourceId, targetId)
	fmt.Println("Following a user")
	return nil
}

func (f *followService) UnFollowUser(sourceId, targetId string) error {
	f.repository.Delete(sourceId, targetId)
	fmt.Println("Unfollowing a user")
	return nil
}
