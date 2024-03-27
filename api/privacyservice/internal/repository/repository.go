package repository

import (
	"context"
	"fmt"
	"social-network/privacyservice/model"
)

type PrivacyRepository interface {
	Create(ctx context.Context, parent_id string, privacy int32) (*model.PrivacyReq, error)
	Update(ctx context.Context, parent_id string, privacy int32) (*model.PrivacyReq, error)
	Delete(ctx context.Context, parent_id string) error
	Get(ctx context.Context, parent_id string) (*model.PrivacyReq, error)
}

type privacyRepository struct {
	// db *sql.DB
}

func NewRepository() PrivacyRepository {
	return &privacyRepository{
		// db: db,
	}
}

func (p *privacyRepository) Create(ctx context.Context, parent_id string, privacy int32) (*model.PrivacyReq, error) {
	fmt.Println("Inserting something to db PRIVACY CREATE", parent_id, privacy)
	return &model.PrivacyReq{}, nil
}

func (p *privacyRepository) Update(ctx context.Context, parent_id string, privacy int32) (*model.PrivacyReq, error) {
	fmt.Println("Inserting something to db PRIVACY UPDATE", parent_id, privacy)
	return &model.PrivacyReq{}, nil
}

func (p *privacyRepository) Delete(ctx context.Context, parent_id string) error {
	fmt.Println("Inserting something to db PRIVACY DELETE", parent_id)
	return nil
}

func (p *privacyRepository) Get(ctx context.Context, parent_id string) (*model.PrivacyReq, error) {
	fmt.Println("Retrieveing something from db PRVIACY GET", parent_id)
	return &model.PrivacyReq{}, nil
}
