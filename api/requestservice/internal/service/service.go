package service

import (
	"context"
	"fmt"
	"social-network/requestservice/internal/repository"
	"social-network/requestservice/model"
)

type RequestService interface {
	Create(ctx context.Context, req model.RequestReq) (*model.RequestReq, error)
	Delete(ctx context.Context, req model.RequestReq) (error)
	Get(ctx context.Context, req model.RequestReq) (*model.RequestReq, error)
	GetNotifications(ctx context.Context, user_id model.RequestUserId) (*model.RequestReqs, error)
	GetInvitations(ctx context.Context, user_id model.RequestUserId) (*model.RequestReqs, error)
	GetFollowrequests(ctx context.Context, user_id model.RequestUserId) (*model.RequestReqs, error)
}

type requestService struct {
	repository repository.RequestRepository
}

func NewService(rep repository.RequestRepository) RequestService {
	return &requestService{
		repository: rep,
	}
}

func (r *requestService) Create(ctx context.Context, req model.RequestReq) (*model.RequestReq, error) {
	_, err := r.repository.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	fmt.Println("Creating request", req)
	return &model.RequestReq{}, nil
}

func (r *requestService) Delete(ctx context.Context, req model.RequestReq) (error) {
	err := r.repository.Delete(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println("Deleting request", req)
	return nil
}

func (r *requestService) Get(ctx context.Context, req model.RequestReq) (*model.RequestReq, error) {
	_, err := r.repository.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	fmt.Println("Getting a request", req)
	return &model.RequestReq{}, nil
}

func (r *requestService) GetNotifications(ctx context.Context, user_id model.RequestUserId) (*model.RequestReqs, error) {
	_, err := r.repository.GetNotifications(ctx, user_id)
	if err != nil {
		return nil, err
	}
	fmt.Println("Getting all the notfications for user ", user_id)
	return &model.RequestReqs{}, nil
}

func (r *requestService) GetInvitations(ctx context.Context, user_id model.RequestUserId) (*model.RequestReqs, error) {
	_, err := r.repository.GetInvitations(ctx, user_id)
	if err != nil {
		return nil, err
	}
	fmt.Println("Getting all the invitations for user", user_id)
	return &model.RequestReqs{}, nil
}

func (r *requestService) GetFollowrequests(ctx context.Context, user_id model.RequestUserId) (*model.RequestReqs, error) {
	_, err := r.repository.GetFollowrequests(ctx, user_id)
	if err != nil {
		return nil, err
	}
	fmt.Println("Getting all the invitations for user ", user_id)
	return &model.RequestReqs{}, nil
}
