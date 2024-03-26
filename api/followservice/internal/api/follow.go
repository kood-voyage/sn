package api

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"social-network/followservice/pkg/followservice"
)

func (i *Implementation) Follow(ctx context.Context, request *followservice.FollowRequest) (*emptypb.Empty, error) {
	err := i.followService.FollowUser(ctx, request.GetSourceId(), request.GetTargetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
