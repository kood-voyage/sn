package clients

import (
	"context"
	"social-network/requestservice/pkg/requestservice"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RequestClient struct {
	request requestservice.RequestClient
}

func NewRequestClient(ctx context.Context, addr string) (*RequestClient, error) {
	cc, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &RequestClient{
		request: requestservice.NewRequestClient(cc),
	}, nil
}

func (r *RequestClient) Create(ctx context.Context, req *requestservice.RequestReq) (*requestservice.RequestReq, error) {
	return r.request.Create(ctx, req)
}

func (r *RequestClient) Delete(ctx context.Context, req *requestservice.RequestReq) (*emptypb.Empty, error) {
	return r.request.Delete(ctx, req)
}
func (r *RequestClient) Get(ctx context.Context, req *requestservice.RequestReq) (*requestservice.RequestReq, error) {
	return r.request.Get(ctx, req)
}
func (r *RequestClient) GetNotifications(ctx context.Context, req *requestservice.RequestUserId) (*requestservice.RequestReqs, error) {
	return r.request.GetNotifications(ctx, req)
}
func (r *RequestClient) GetInvitations(ctx context.Context, req *requestservice.RequestUserId) (*requestservice.RequestReqs, error) {
	return r.request.GetInvitations(ctx, req)
}
func (r *RequestClient) GetFollowrequests(ctx context.Context, req *requestservice.RequestUserId) (*requestservice.RequestReqs, error) {
	return r.request.GetFollowrequests(ctx, req)
}
