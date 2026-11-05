package threatservice

import (
	domain_threatDomain "Threat-intelligence-service/internal/core/domain"
	"context"
	"fmt"
)

type ThreatRepository interface {
	AddThreat(ctx context.Context, threat *domain_threatDomain.Threat) error
}

type ThreatService struct {
	threatRepository ThreatRepository
}

func NewThreatService(threatRepository ThreatRepository) *ThreatService {
	return &ThreatService{
		threatRepository: threatRepository,
	}
}

func(s *ThreatService) AddThreat(ctx context.Context,threat *domain_threatDomain.Threat) error  {
	
	if err := s.threatRepository.AddThreat(ctx,threat); err != nil {
		return fmt.Errorf("service AddThreat %w",err)
	}

	return nil 
}