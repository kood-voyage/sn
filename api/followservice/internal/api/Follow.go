package api

import (
	"context"
	"social-network/followservice/pkg/followservice"
)

func (i *Implementation) Follow(ctx context.Context, request *followservice.FollowRequest) 