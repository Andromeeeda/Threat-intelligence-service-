package repository_postgres_rep

import (
	domain_threatDomain "Threat-intelligence-service/internal/core/domain"
	common_errors "Threat-intelligence-service/internal/core/errors/common"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *ThreatRepository) Update(ctx context.Context, changeThreat *domain_threatDomain.Threat) (
	*domain_threatDomain.Threat,
	error) {



	query := `
		UPDATE app.threats 
		SET name = $1, severity = $2, description = $3 
		WHERE md5 = $4 
		RETURNING md5, name, severity, description`

	var updatedThreat domain_threatDomain.Threat
	err := r.conn.QueryRow(ctx, query, changeThreat.Name, 
		changeThreat.Severity, 
		changeThreat.Description, 
		changeThreat.Md5,).Scan(
		&updatedThreat.Md5,
		&updatedThreat.Name,
		&updatedThreat.Severity,
		&updatedThreat.Description,
	)

	if err != nil {
		if errors.Is(err,pgx.ErrNoRows) {
			return nil,fmt.Errorf("threat repository update err: %w",common_errors.ErrNotFound)
		} else {
			return nil,fmt.Errorf("threat repository update err: %w",err)
		}
	}

	return &updatedThreat,nil 
}
