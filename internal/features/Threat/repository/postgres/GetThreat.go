package repository_postgres_rep

import (
	domain_threatDomain "Threat-intelligence-service/internal/core/domain"
	common_errors "Threat-intelligence-service/internal/core/errors/common"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *ThreatRepository) GetThreat(ctx context.Context, md5 string) (*domain_threatDomain.Threat, error) {
	query := `
	SELECT md5,name FROM app.threats
	WHERE md5 = $1`

	var threatModel ThreatModel
	err := r.conn.QueryRow(ctx, query, md5).Scan(&threatModel.Md5, &threatModel.Name)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil,fmt.Errorf("repository getThreat not found: %w",common_errors.ErrNotFound)
		} else {
			return nil,fmt.Errorf("repository getThreat: %w",err)
		}
	}

	threatDomain := &domain_threatDomain.Threat {
		Md5: threatModel.Md5,
		Name: threatModel.Name,
	}	

	return threatDomain,nil 
	
}
