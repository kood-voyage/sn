package service

import (
	"context"
	"errors"
	"fmt"
	"social-network/followservice/internal/repository"
	"social-network/followservice/model"
	"social-network/internal/clients"
	"social-network/privacyservice/pkg/privacyservice"
)

type FollowService interface {
	FollowUser(ctx context.Context, sourceId, targetId string) (*model.Follow, error)
	UnFollowUser(ctx context.Context, sourceId, targetId string) error
}

type followService struct {
	repository    repository.FollowRepository
	PrivacyClient clients.PrivacyClient
}

func NewService(rep repository.FollowRepository, pc *clients.PrivacyClient) FollowService {
	return &followService{
		repository:    rep,
		PrivacyClient: *pc,
	}
}

func (f *followService) FollowUser(ctx context.Context, sourceId, targetId string) (*model.Follow, error) {
	//source and targer cant be the same
	if sourceId == targetId {
		return nil, errors.New("user can not follow itself")
	}
	//check target user privacy
	privacy, err := f.PrivacyClient.Get(ctx, &privacyservice.PrivacyId{
		ParentId: targetId,
	})
	if err != nil {
		return nil, err
	}

	if privacy.Privacy == model.PUBLIC {
		follow, err := f.repository.Create(ctx, sourceId, targetId)
		if err != nil {
			return nil, err
		}
		fmt.Println("Following a user", follow)
		return follow, nil
	} else if privacy.Privacy == model.PRIVATE {
		return nil, fmt.Errorf("user privacy is not public it is - %d", privacy.Privacy)
	} else {
		return nil, fmt.Errorf("unexpected privacy - %d", privacy.Privacy)
	}
}

func (f *followService) UnFollowUser(ctx context.Context, sourceId, targetId string) error {
	err := f.repository.Delete(ctx, sourceId, targetId)
	if err != nil {
		return err
	}
	fmt.Println("Unfollowing a user", sourceId, targetId)
	return nil
}
