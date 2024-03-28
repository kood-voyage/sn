package api

import (
	"context"
	"social-network/requestservice/model"
	"social-network/requestservice/pkg/requestservice"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i *Implementation) Create(ctx context.Context, req *requestservice.RequestReq) (*requestservice.RequestReq, error) {
	r, err := i.requestService.Create(ctx, model.RequestReq{
		Id:        req.GetId(),
		TypeId:    req.GetTypeId(),
		SourceId:  req.GetSourceId(),
		TargetId:  req.GetTargetId(),
		ParentId:  req.GetParentId(),
		Message:   req.GetMessage(),
		CreatedAt: req.GetCreatedAt().AsTime(),
	})
	if err != nil {
		return nil, err
	}

	//convert timestamp
	c := timestamppb.New(r.CreatedAt)
	return &requestservice.RequestReq{
		Id:        r.Id,
		TypeId:    r.TypeId,
		SourceId:  r.SourceId,
		TargetId:  r.TargetId,
		ParentId:  r.ParentId,
		Message:   r.Message,
		CreatedAt: c,
	}, nil
}

func (i *Implementation) Delete(ctx context.Context, req *requestservice.RequestReq) (*emptypb.Empty, error) {
	if err := i.requestService.Delete(ctx, model.RequestReq{
		Id:        req.GetId(),
		TypeId:    req.GetTypeId(),
		SourceId:  req.GetSourceId(),
		TargetId:  req.GetTargetId(),
		ParentId:  req.GetParentId(),
		Message:   req.GetMessage(),
		CreatedAt: req.GetCreatedAt().AsTime(),
	}); err != nil {
		return nil, err
	}

	return nil, nil
}

func (i *Implementation) Get(ctx context.Context, req *requestservice.RequestReq) (*requestservice.RequestReq, error) {
	r, err := i.requestService.Get(ctx, model.RequestReq{
		Id:        req.GetId(),
		TypeId:    req.GetTypeId(),
		SourceId:  req.GetSourceId(),
		TargetId:  req.GetTargetId(),
		ParentId:  req.GetParentId(),
		Message:   req.GetMessage(),
		CreatedAt: req.GetCreatedAt().AsTime(),
	})
	if err != nil {
		return nil, err
	}

	//convert timestamp
	c := timestamppb.New(r.CreatedAt)
	return &requestservice.RequestReq{
		Id:        r.Id,
		TypeId:    r.TypeId,
		SourceId:  r.SourceId,
		TargetId:  r.TargetId,
		ParentId:  r.ParentId,
		Message:   r.Message,
		CreatedAt: c,
	}, nil
}

func (i *Implementation) GetNotifications(ctx context.Context, req *requestservice.RequestUserId) (*requestservice.RequestReqs, error) {
	r, err := i.requestService.GetNotifications(ctx, model.RequestUserId{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	requestReqs := &requestservice.RequestReqs{}
	for _, requestReq := range r.Requests {
		requestReqs.Requests = append(requestReqs.Requests, &requestservice.RequestReq{
			Id:        requestReq.Id,
			TypeId:    requestReq.TypeId,
			SourceId:  requestReq.SourceId,
			TargetId:  requestReq.TargetId,
			ParentId:  requestReq.ParentId,
			Message:   requestReq.Message,
			CreatedAt: timestamppb.New(requestReq.CreatedAt),
		})
	}

	return requestReqs, nil
}

func (i *Implementation) GetInvitations(ctx context.Context, req *requestservice.RequestUserId) (*requestservice.RequestReqs, error) {
	r, err := i.requestService.GetInvitations(ctx, model.RequestUserId{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	requestReqs := &requestservice.RequestReqs{}
	for _, requestReq := range r.Requests {
		requestReqs.Requests = append(requestReqs.Requests, &requestservice.RequestReq{
			Id:        requestReq.Id,
			TypeId:    requestReq.TypeId,
			SourceId:  requestReq.SourceId,
			TargetId:  requestReq.TargetId,
			ParentId:  requestReq.ParentId,
			Message:   requestReq.Message,
			CreatedAt: timestamppb.New(requestReq.CreatedAt),
		})
	}

	return requestReqs, nil
}

func (i *Implementation) GetFollowrequests(ctx context.Context, req *requestservice.RequestUserId) (*requestservice.RequestReqs, error) {
	r, err := i.requestService.GetFollowrequests(ctx, model.RequestUserId{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	requestReqs := &requestservice.RequestReqs{}
	for _, requestReq := range r.Requests {
		requestReqs.Requests = append(requestReqs.Requests, &requestservice.RequestReq{
			Id:        requestReq.Id,
			TypeId:    requestReq.TypeId,
			SourceId:  requestReq.SourceId,
			TargetId:  requestReq.TargetId,
			ParentId:  requestReq.ParentId,
			Message:   requestReq.Message,
			CreatedAt: timestamppb.New(requestReq.CreatedAt),
		})
	}

	return requestReqs, nil
}
