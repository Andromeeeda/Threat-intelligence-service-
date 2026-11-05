package transport_threat_handlers

import (
	domain_threatDomain "Threat-intelligence-service/internal/core/domain"
	"context"
)

type ThreatService interface {
	AddThreat(ctx context.Context, threat *domain_threatDomain.Threat) error
}

type ThreatHandler struct {
	threatService ThreatService
}

func NewThreatHandler(threatService ThreatService) *ThreatHandler {
	return &ThreatHandler{
		threatService: threatService,
	}
}
