package repository

import (
	"context"
	"database/sql"
	"social-network/requestservice/model"
)

type RequestRepository interface {
	Create(ctx context.Context, req model.RequestReq) (*model.RequestReq, error)
	Delete(ctx context.Context, req model.RequestReq) error
	Get(ctx context.Context, req model.RequestReq) (*model.RequestReq, error)
	GetNotifications(ctx context.Context, user_id model.RequestUserId) (*model.RequestReqs, error)
	GetInvitations(ctx context.Context, user_id model.RequestUserId) (*model.RequestReqs, error)
	GetFollowrequests(ctx context.Context, user_id model.RequestUserId) (*model.RequestReqs, error)
}

type requestRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) RequestRepository {
	return &requestRepository{
		db: db,
	}
}

func (r *requestRepository) Create(ctx context.Context, req model.RequestReq) (*model.RequestReq, error) {
	return nil, nil
}

func (r *requestRepository) Delete(ctx context.Context, req model.RequestReq) error {
	return nil
}

func (r *requestRepository) Get(ctx context.Context, req model.RequestReq) (*model.RequestReq, error) {
	return nil, nil
}

func (r *requestRepository) GetNotifications(ctx context.Context, user_id model.RequestUserId) (*model.RequestReqs, error) {
	return nil, nil
}

func (r *requestRepository) GetInvitations(ctx context.Context, user_id model.RequestUserId) (*model.RequestReqs, error) {
	return nil, nil
}

func (r *requestRepository) GetFollowrequests(ctx context.Context, user_id model.RequestUserId) (*model.RequestReqs, error) {
	return nil, nil
}
