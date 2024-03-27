package api

import (
	"social-network/privacyservice/internal/service"
	"social-network/privacyservice/pkg/privacyservice"
)

type Implementation struct {
	privacyservice.UnimplementedPrivacyServer
	privacyService service.PrivacyService
}

func NewImplementation(privacyService service.PrivacyService) *Implementation {
	return &Implementation{privacyService: privacyService}
}
