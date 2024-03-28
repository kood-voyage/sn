package clients

import (
	"context"
	"social-network/followservice/pkg/followservice"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FollowClient struct {
	follow followservice.FollowClient
}

func NewFollowClient(ctx context.Context, addr string) (*FollowClient, error) {
	cc, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &FollowClient{
		follow: followservice.NewFollowClient(cc),
	}, nil
}

func (c *FollowClient) Follow(ctx context.Context, req *followservice.FollowRequest) (*followservice.FollowRequest, error) {
	return c.follow.Follow(ctx, req)
}

func (c *FollowClient) UnFollow(ctx context.Context, req *followservice.FollowRequest) (*emptypb.Empty, error) {
	return c.follow.UnFollow(ctx, req)
}
