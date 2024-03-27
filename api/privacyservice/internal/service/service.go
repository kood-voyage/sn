package service

import (
	"context"
	"fmt"
	"social-network/privacyservice/internal/repository"
	"social-network/privacyservice/model"
)

type PrivacyService interface {
	Create(ctx context.Context, parent_id string, privacy int32) (*model.PrivacyReq, error)
	Update(ctx context.Context, parent_id string, privacy int32) (*model.PrivacyReq, error)
	Delete(ctx context.Context, parent_id string) error
	Get(ctx context.Context, parent_id string) (*model.PrivacyReq, error)
}

type privacyService struct {
	repository repository.PrivacyRepository
}

func NewService(rep repository.PrivacyRepository) PrivacyService {
	return &privacyService{
		repository: rep,
	}
}

func (p *privacyService) Create(ctx context.Context, parent_id string, privacy int32) (*model.PrivacyReq, error) {
	_, err := p.repository.Create(ctx, parent_id, privacy)
	if err != nil {
		return nil, err
	}
	fmt.Println("Setting privacy for id", parent_id, privacy)
	return &model.PrivacyReq{
		ParentId: parent_id,
		Privacy:  privacy,
	}, nil
}

func (p *privacyService) Update(ctx context.Context, parent_id string, privacy int32) (*model.PrivacyReq, error) {
	_, err := p.repository.Update(ctx, parent_id, privacy)
	if err != nil {
		return nil, err
	}
	fmt.Println("Updatng id privacy", parent_id, privacy)
	return &model.PrivacyReq{
		ParentId: parent_id,
		Privacy:  privacy,
	}, nil
}

func (p *privacyService) Delete(ctx context.Context, parent_id string) error {
	err := p.repository.Delete(ctx, parent_id)
	if err != nil {
		return err
	}
	fmt.Println("deleting a privacy", parent_id)
	return nil
}

func (p *privacyService) Get(ctx context.Context, parent_id string) (*model.PrivacyReq, error) {
	_, err := p.repository.Get(ctx, parent_id)
	if err != nil {
		return nil, err
	}
	fmt.Println("Retrieving privacy info for that id ", parent_id)
	return &model.PrivacyReq{
		ParentId: parent_id,
	}, nil
}
