package api

import (
	"context"
	"social-network/privacyservice/pkg/privacyservice"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Create(ctx context.Context, request *privacyservice.PrivacyReq) (*privacyservice.PrivacyReq, error) {
	r, err := i.privacyService.Create(ctx, request.GetParentId(), request.GetPrivacy())
	if err != nil {
		return nil, err
	}

	return &privacyservice.PrivacyReq{
		ParentId: r.ParentId,
		Privacy:  r.Privacy,
	}, nil
}

func (i *Implementation) Update(ctx context.Context, request *privacyservice.PrivacyReq) (*privacyservice.PrivacyReq, error) {
	r, err := i.privacyService.Update(ctx, request.GetParentId(), request.GetPrivacy())
	if err != nil {
		return nil, err
	}

	return &privacyservice.PrivacyReq{
		ParentId: r.ParentId,
		Privacy:  r.Privacy,
	}, nil
}

func (i *Implementation) Delete(ctx context.Context, request *privacyservice.PrivacyId) (*emptypb.Empty, error) {
	err := i.privacyService.Delete(ctx, request.GetParentId())
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (i *Implementation) Get(ctx context.Context, request *privacyservice.PrivacyId) (*privacyservice.PrivacyReq, error) {
	r, err := i.privacyService.Get(ctx, request.GetParentId())
	if err != nil {
		return nil, err
	}
	return &privacyservice.PrivacyReq{
		ParentId: r.ParentId,
		Privacy:  r.Privacy,
	}, nil
}
