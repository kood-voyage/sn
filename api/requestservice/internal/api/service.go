package api

import (
	"social-network/requestservice/internal/service"
	"social-network/requestservice/pkg/requestservice"
)

type Implementation struct {
	requestservice.UnimplementedRequestServer
	requestService service.RequestService
}

func NewImplementation(requestService service.RequestService) *Implementation {
	return &Implementation{requestService: requestService}
}
