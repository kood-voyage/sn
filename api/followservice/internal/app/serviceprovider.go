package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"social-network/followservice/internal/api"
	"social-network/followservice/internal/clients"
	"social-network/followservice/internal/config"
	"social-network/followservice/internal/repository"
	"social-network/followservice/internal/service"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	"github.com/golang-migrate/migrate/source/file"
)

type ServiceProvider struct {
	grpcConfig       config.GRPCConfig
	httpConfig       config.HTTPConfig
	config           config.GeneralConfig
	dbClient         *sql.DB
	followRepository repository.FollowRepository
	followService    service.FollowService

	followImpl *api.Implementation
}

func newServiceProvider() *ServiceProvider {
	confGRPC := config.NewGRPCConfig()
	if err := confGRPC.ReadConfig("followservice/config/grpcconfig.json"); err != nil {
		log.Fatalf("failed to parse grpc config: %v", err)
	}
	confHTTP := config.NewHttpConfig()
	if err := confHTTP.ReadConfig("followservice/config/httpconfig.json"); err != nil {
		log.Fatalf("failed to parse http config: %v", err)
	}

	confGeneral := config.NewConfig()
	if err := confGeneral.ReadConfig("followservice/config/config.json"); err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	return &ServiceProvider{
		grpcConfig: confGRPC,
		httpConfig: confHTTP,
		config:     *confGeneral,
	}
}

func InitializeDB(conf config.GeneralConfig) *sql.DB {
	db, err := newDB(conf.DatabaseURL, conf.Migrations, conf.Driver)
	if err != nil {
		log.Fatal("Couldnt connect with database", err)
	}

	return db
}

func newDB(databaseURL, migrationSource, driver string) (*sql.DB, error) {
	db, err := sql.Open(driver, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return nil, fmt.Errorf("instance: %w", err)
	}

	fileSource, err := (&file.File{}).Open(migrationSource)
	if err != nil {
		return nil, fmt.Errorf("fileSource: %w", err)
	}

	m, err := migrate.NewWithInstance("file", fileSource, driver, instance)
	if err != nil {
		return nil, fmt.Errorf("migrations new: %w", err)
	}

	if err = m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return nil, fmt.Errorf("migrations run: %w", err)
		}
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (s *ServiceProvider) FollowRepository(ctx context.Context) repository.FollowRepository {
	if s.followRepository == nil {
		s.followRepository = repository.NewRepository(s.dbClient)
	}

	return s.followRepository
}

func (s *ServiceProvider) FollowService(ctx context.Context) service.FollowService {
	privacyClient, err := clients.NewPrivacyClient(ctx, s.config.PrivacyClientGRPC)
	if err != nil {
		fmt.Println("Can not connect to privacy service --- ", err)
	}
	requestClient, err := clients.NewRequestClient(ctx, s.config.RequestClientGRPC)
	if err != nil {
		fmt.Println("Can not connect to request service --- ", err)
	}

	if s.followService == nil {
		s.followService = service.NewService(s.FollowRepository(ctx), privacyClient, requestClient)
	}

	return s.followService
}

func (s *ServiceProvider) FollowImpl(ctx context.Context) *api.Implementation {
	if s.followImpl == nil {
		s.followImpl = api.NewImplementation(s.FollowService(ctx))
	}
	return s.followImpl
}
