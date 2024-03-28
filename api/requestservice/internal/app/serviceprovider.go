package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"social-network/requestservice/internal/api"
	"social-network/requestservice/internal/config"
	"social-network/requestservice/internal/repository"
	"social-network/requestservice/internal/service"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
)

type ServiceProvider struct {
	grpcConfig        config.GRPCConfig
	httpConfig        config.HTTPConfig
	config            config.GeneralConfig
	dbClient          *sql.DB
	requestRepository repository.RequestRepository
	requestService    service.RequestService
	requestImpl       *api.Implementation
}

func newServiceProvider() *ServiceProvider {
	confGRPC := config.NewGRPCConfig()
	if err := confGRPC.ReadConfig("requestservice/config/grpcconfig.json"); err != nil {
		log.Fatalf("failed to parse grpc config: %v", err)
	}
	confHTTP := config.NewHttpConfig()
	if err := confHTTP.ReadConfig("requestservice/config/httpconfig.json"); err != nil {
		log.Fatalf("failed to parse http config: %v", err)
	}
	confGeneral := config.NewConfig()
	if err := confGeneral.ReadConfig("requestservice/config/config.json"); err != nil {
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

	//NEED TO THINK OF A SOLUTION LATER
	// defer db.Close()

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

func (s *ServiceProvider) RequestRepository(ctx context.Context) repository.RequestRepository {
	if s.requestRepository == nil {
		s.requestRepository = repository.NewRepository(s.dbClient)
	}

	return s.requestRepository
}

func (s *ServiceProvider) RequestService(ctx context.Context) service.RequestService {
	if s.requestService == nil {
		s.requestService = service.NewService(s.RequestRepository(ctx))
	}
	return s.requestService
}

func (s *ServiceProvider) RequestImpl(ctx context.Context) *api.Implementation {
	if s.requestImpl == nil {
		s.requestImpl = api.NewImplementation(s.RequestService(ctx))
	}
	return s.requestImpl
}
