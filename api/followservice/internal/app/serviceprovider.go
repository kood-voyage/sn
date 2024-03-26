package app

import (
	"context"
	"database/sql"
	"log"
	"social-network/followservice/internal/api"
	"social-network/followservice/internal/config"
	"social-network/followservice/internal/repository"
	"social-network/followservice/internal/service"
)

type ServiceProvider struct {
	grpcConfig config.GRPCConfig
	httpConfig config.HTTPConfig

	dbClient         *sql.DB
	followRepository repository.FollowRepository
	followService    service.FollowService

	followImpl *api.Implementation
}

func newServiceProvider() *ServiceProvider {
	confGRPC := config.NewGRPCConfig()
	if err := confGRPC.ReadConfig("followservice/config/grpcconfig.json"); err != nil {
		log.Fatal(err)
	}
	confHTTP := config.NewHttpConfig()
	if err := confHTTP.ReadConfig("followservice/config/httpconfig.json"); err != nil {
		log.Fatal(err)
	}

	return &ServiceProvider{
		grpcConfig: confGRPC,
		httpConfig: confHTTP,
	}
}

func (s *ServiceProvider) FollowRepository(ctx context.Context) repository.FollowRepository {
	if s.followRepository == nil {
		s.followRepository = repository.NewRepository()
	}

	return s.followRepository
}

func (s *ServiceProvider) FollowService(ctx context.Context) service.FollowService {
	if s.followService == nil {
		s.followService = service.NewService(s.FollowRepository(ctx))
	}
	return s.followService
}

func (s *ServiceProvider) FollowImpl(ctx context.Context) *api.Implementation {
	if s.followImpl == nil {
		s.followImpl = api.NewImplementation(s.FollowService(ctx))
	}
	return s.followImpl
}
