package threatservice

import (
	domain_threatDomain "Threat-intelligence-service/internal/core/domain"
	"context"
	"fmt"
)

type ThreatRepository interface {
	AddThreat(ctx context.Context, threat *domain_threatDomain.Threat) error
	GetThreat(ctx context.Context, md5 string) (*domain_threatDomain.Threat, error)
	Update(ctx context.Context, changeThreat *domain_threatDomain.Threat) (
		*domain_threatDomain.Threat,
		error)
}

type ThreatService struct {
	threatRepository ThreatRepository
}

func NewThreatService(threatRepository ThreatRepository) *ThreatService {
	return &ThreatService{
		threatRepository: threatRepository,
	}
}

func (s *ThreatService) AddThreat(ctx context.Context, threat *domain_threatDomain.Threat) error {

	if err := s.threatRepository.AddThreat(ctx, threat); err != nil {
		return fmt.Errorf("service AddThreat %w", err)
	}

	return nil
}

func (s *ThreatService) GetThreat(ctx context.Context, md5 string) (*domain_threatDomain.Threat, error) {

	threat, err := s.threatRepository.GetThreat(ctx, md5)
	if err != nil {
		return nil, fmt.Errorf("getThreat service: %w", err)
	}

	return threat, nil
}

func (s *ThreatService) Change(ctx context.Context, threat *domain_threatDomain.Threat) (*domain_threatDomain.Threat, error) {

	updatedThreat,err := s.threatRepository.Update(ctx,threat)

	if err != nil {
		return nil,fmt.Errorf("change service: %w",err)
	}

	return updatedThreat,nil 
}
