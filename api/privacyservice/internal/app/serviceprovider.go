package app

import (
	"context"
	"database/sql"
	"log"
	"social-network/privacyservice/internal/api"
	"social-network/privacyservice/internal/config"
	"social-network/privacyservice/internal/repository"
	"social-network/privacyservice/internal/service"
)

type ServiceProvider struct {
	grpcConfig config.GRPCConfig
	httpConfig config.HTTPConfig

	dbClient          *sql.DB
	privacyRepository repository.PrivacyRepository
	privacyService    service.PrivacyService

	privacyImpl *api.Implementation
}

func newServiceProvider() *ServiceProvider {
	confGRPC := config.NewGRPCConfig()
	if err := confGRPC.ReadConfig("privacyservice/config/grpcconfig.json"); err != nil {
		log.Fatalf("failed to parse grpc config: %v", err)
	}
	confHTTP := config.NewHttpConfig()
	if err := confHTTP.ReadConfig("privacyservice/config/httpconfig.json"); err != nil {
		log.Fatalf("failed to parse http config: %v", err)
	}

	return &ServiceProvider{
		grpcConfig: confGRPC,
		httpConfig: confHTTP,
	}
}

func (s *ServiceProvider) PrivacyRepository(ctx context.Context) repository.PrivacyRepository {
	if s.privacyRepository == nil {
		s.privacyRepository = repository.NewRepository()
	}

	return s.privacyRepository
}

func (s *ServiceProvider) PrivacyService(ctx context.Context) service.PrivacyService {
	if s.privacyRepository == nil {
		s.privacyRepository = service.NewService(s.PrivacyRepository(ctx))
	}
	return s.privacyRepository
}

func (s *ServiceProvider) PrivacyImpl(ctx context.Context) *api.Implementation {
	if s.privacyImpl == nil {
		s.privacyImpl = api.NewImplementation(s.PrivacyService(ctx))
	}
	return s.privacyImpl
}
