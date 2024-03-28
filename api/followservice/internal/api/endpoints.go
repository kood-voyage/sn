package api

import (
	"context"
	"social-network/followservice/pkg/followservice"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Follow(ctx context.Context, request *followservice.FollowRequest) (*followservice.FollowRequest, error) {
	f, err := i.followService.FollowUser(ctx, request.GetSourceId(), request.GetTargetId())
	if err != nil {
		return nil, err
	}

	return &followservice.FollowRequest{
		SourceId: f.SourceId,
		TargetId: f.TargetId,
	}, nil
}

func (i *Implementation) UnFollow(ctx context.Context, request *followservice.FollowRequest) (*emptypb.Empty, error) {
	err := i.followService.UnFollowUser(ctx, request.GetSourceId(), request.GetTargetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
