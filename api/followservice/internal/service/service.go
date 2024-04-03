package service

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"social-network/followservice/internal/clients"
	"social-network/followservice/internal/repository"
	"social-network/followservice/model"
	"social-network/privacyservice/pkg/privacyservice"
	models "social-network/requestservice/model"
	"social-network/requestservice/pkg/requestservice"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FollowService interface {
	FollowUser(ctx context.Context, sourceId, targetId string) (*model.Follow, error)
	UnFollowUser(ctx context.Context, sourceId, targetId string) error
}

type followService struct {
	repository    repository.FollowRepository
	PrivacyClient clients.PrivacyClient
	RequestClient clients.RequestClient
}

func NewService(rep repository.FollowRepository, pc *clients.PrivacyClient, rc *clients.RequestClient) FollowService {
	return &followService{
		repository:    rep,
		PrivacyClient: *pc,
		RequestClient: *rc,
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
	var follow *model.Follow
	if privacy.Privacy == model.PUBLIC {
		follow, err = f.repository.Create(ctx, sourceId, targetId)
		if err != nil {
			return nil, err
		}
		fmt.Println("Following a user", follow)
		return follow, nil
	} else if privacy.Privacy == model.PRIVATE {
		existing, err := f.RequestClient.Get(ctx, &requestservice.RequestReq{
			SourceId: sourceId,
			TargetId: targetId,
			TypeId:   models.FOLLOW,
		})
		if err != nil {
			return nil, err
		}
		fmt.Printf("EXISITING --> %+v\n", existing)
		fmt.Println(reflect.ValueOf(existing))
		test := reflect.ValueOf(existing).IsValid()
		fmt.Println("TEST --> ", test)
		fmt.Println("JEPS -> ", reflect.Zero(reflect.TypeOf(existing)).Interface()) 
		if existing != nil {
			return nil, errors.New("request already exists")
		}
		fmt.Println("test3")

		_, err = f.RequestClient.Create(ctx, &requestservice.RequestReq{
			Id:        uuid.New().String(),
			SourceId:  sourceId,
			TargetId:  targetId,
			TypeId:    models.FOLLOW,
			Message:   "has requested to follow you",
			CreatedAt: timestamppb.New(time.Now()),
		})

		if err != nil {
			return nil, err
		}
		return follow, err
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
