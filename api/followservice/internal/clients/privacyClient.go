package clients

import (
	"context"
	"social-network/privacyservice/pkg/privacyservice"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PrivacyClient struct {
	privacy privacyservice.PrivacyClient
}

func NewPrivacyClient(ctx context.Context, addr string) (*PrivacyClient, error) {
	cc, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &PrivacyClient{
		privacy: privacyservice.NewPrivacyClient(cc),
	}, nil
}

func (p *PrivacyClient) Create(ctx context.Context, request *privacyservice.PrivacyReq) (*privacyservice.PrivacyReq, error) {
	return p.privacy.Create(ctx, request)
}

func (p *PrivacyClient) Update(ctx context.Context, request *privacyservice.PrivacyReq) (*privacyservice.PrivacyReq, error) {
	return p.privacy.Update(ctx, request)
}

func (p *PrivacyClient) Delete(ctx context.Context, request *privacyservice.PrivacyId) (*emptypb.Empty, error) {
	return p.privacy.Delete(ctx, request)
}

func (p *PrivacyClient) Get(ctx context.Context, request *privacyservice.PrivacyId) (*privacyservice.PrivacyReq, error) {
	return p.privacy.Get(ctx, request)
}
