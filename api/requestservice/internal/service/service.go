package service

import (
	"context"
	"fmt"
	"social-network/requestservice/internal/repository"
	models "social-network/requestservice/model"
)

type RequestService interface {
	Create(ctx context.Context, req models.RequestReq) (*models.RequestReq, error)
	Delete(ctx context.Context, req models.RequestReq) error
	Get(ctx context.Context, req models.RequestReq) (*models.RequestReq, error)
	GetNotifications(ctx context.Context, user_id models.RequestUserId) (*models.RequestReqs, error)
	GetInvitations(ctx context.Context, user_id models.RequestUserId) (*models.RequestReqs, error)
	GetFollowrequests(ctx context.Context, user_id models.RequestUserId) (*models.RequestReqs, error)
}

type requestService struct {
	repository repository.RequestRepository
}

func NewService(rep repository.RequestRepository) RequestService {
	return &requestService{
		repository: rep,
	}
}

func (r *requestService) Create(ctx context.Context, req models.RequestReq) (*models.RequestReq, error) {
	rr, err := r.repository.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	fmt.Println("Creating request", req)
	return rr, nil
}

func (r *requestService) Delete(ctx context.Context, req models.RequestReq) error {
	err := r.repository.Delete(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println("Deleting request", req)
	return nil
}

func (r *requestService) Get(ctx context.Context, req models.RequestReq) (*models.RequestReq, error) {
	rr, err := r.repository.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	fmt.Println("Getting a request", req)
	return rr, nil
}

func (r *requestService) GetNotifications(ctx context.Context, user_id models.RequestUserId) (*models.RequestReqs, error) {
	rr, err := r.repository.GetNotifications(ctx, user_id)
	if err != nil {
		return nil, err
	}
	fmt.Println("Getting all the notfications for user ", user_id)
	return rr, nil
}

func (r *requestService) GetInvitations(ctx context.Context, user_id models.RequestUserId) (*models.RequestReqs, error) {
	rr, err := r.repository.GetInvitations(ctx, user_id)
	if err != nil {
		return nil, err
	}
	fmt.Println("Getting all the invitations for user", user_id)
	return rr, nil
}

func (r *requestService) GetFollowrequests(ctx context.Context, user_id models.RequestUserId) (*models.RequestReqs, error) {
	rr, err := r.repository.GetFollowrequests(ctx, user_id)
	if err != nil {
		return nil, err
	}
	fmt.Println("Getting all the invitations for user ", user_id)
	return rr, nil
}
