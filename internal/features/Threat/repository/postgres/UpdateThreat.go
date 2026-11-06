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

	threatModel := &ThreatModel{
		Md5:  changeThreat.Md5,
		Name: changeThreat.Name,
	}

	query := `
		UPDATE app.threats 
		SET name = $1 
		WHERE md5 = $2 
		RETURNING md5, name`

	var updatedThreat domain_threatDomain.Threat
	err := r.conn.QueryRow(ctx, query, threatModel.Name, threatModel.Md5).Scan(
		&updatedThreat.Md5,
		&updatedThreat.Name,
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
