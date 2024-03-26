package service

import (
	"context"
	"fmt"
	"social-network/followservice/internal/repository"
)

type FollowService interface {
	FollowUser(ctx context.Context, sourceId, targetId string) error
	UnFollowUser(ctx context.Context, sourceId, targetId string) error
}

type followService struct {
	repository repository.FollowRepository
}

func NewService(rep repository.FollowRepository) FollowService {
	return &followService{
		repository: rep,
	}
}

func (f *followService) FollowUser(ctx context.Context, sourceId, targetId string) error {
	err := f.repository.Create(ctx, sourceId, targetId)
	if err != nil {
		return err
	}
	fmt.Println("Following a user", sourceId, targetId)
	return nil
}

func (f *followService) UnFollowUser(ctx context.Context, sourceId, targetId string) error {
	err := f.repository.Delete(ctx, sourceId, targetId)
	if err != nil {
		return err
	}
	fmt.Println("Unfollowing a user", sourceId, targetId)
	return nil
}
