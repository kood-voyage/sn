package api

import (
	"social-network/followservice/internal/service"
	"social-network/followservice/pkg/followservice"
)

type Implementation struct {
	followservice.UnimplementedFollowServer
	followService service.FollowService
}

func NewImplementation(followService service.FollowService) *Implementation {
	return &Implementation{followService:  followService}
}