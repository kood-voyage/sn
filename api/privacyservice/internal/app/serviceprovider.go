package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"social-network/privacyservice/internal/api"
	"social-network/privacyservice/internal/config"
	"social-network/privacyservice/internal/repository"
	"social-network/privacyservice/internal/service"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
)

type ServiceProvider struct {
	grpcConfig        config.GRPCConfig
	httpConfig        config.HTTPConfig
	config            config.GeneralConfig
	dbClient          *sql.DB
	privacyRepository repository.PrivacyRepository
	privacyService    service.PrivacyService
	privacyImpl       *api.Implementation
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
	confGeneral := config.NewConfig()
	if err := confGeneral.ReadConfig("privacyservice/config/config.json"); err != nil {
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

func (s *ServiceProvider) PrivacyRepository(ctx context.Context) repository.PrivacyRepository {
	if s.privacyRepository == nil {
		s.privacyRepository = repository.NewRepository(s.dbClient)
	}

	return s.privacyRepository
}

func (s *ServiceProvider) PrivacyService(ctx context.Context) service.PrivacyService {
	if s.privacyService == nil {
		s.privacyService = service.NewService(s.PrivacyRepository(ctx))
	}
	return s.privacyService
}

func (s *ServiceProvider) PrivacyImpl(ctx context.Context) *api.Implementation {
	if s.privacyImpl == nil {
		s.privacyImpl = api.NewImplementation(s.PrivacyService(ctx))
	}
	return s.privacyImpl
}
